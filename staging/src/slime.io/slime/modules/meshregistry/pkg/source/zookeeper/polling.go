package zookeeper

import (
	"time"

	cmap "github.com/orcaman/concurrent-map"
	networking "istio.io/api/networking/v1alpha3"
	"istio.io/libistio/pkg/config/event"

	"slime.io/slime/modules/meshregistry/pkg/monitoring"
)

func (s *Source) Polling() {
	go func() {
		ticker := time.NewTicker(time.Duration(s.args.RefreshPeriod))
		defer ticker.Stop()
		for {
			s.refresh()

			forceUpdateTrigger := s.forceUpdateTrigger.Load().(chan struct{})
			select {
			case <-s.stop:
				return
			case <-ticker.C:
			case <-forceUpdateTrigger:
			}
		}
	}()
}

func (s *Source) refresh() {
	t0 := time.Now()
	log.Infof("zk refresh start : %d", t0.UnixNano())
	children, err := s.Con.Children(s.args.RegistryRootNode)
	monitoring.RecordSourceClientRequest(SourceName, err == nil)
	if err != nil {
		monitoring.RecordPolling(SourceName, t0, time.Now(), false)
		log.Errorf("zk path %s get child error: %s", s.args.RegistryRootNode, err.Error())
		return
	}
	for _, child := range children {
		s.iface(child)
	}
	s.handleNodeDelete(children)
	t1 := time.Now()
	log.Infof("zk refresh finish : %d", t1.UnixNano())
	monitoring.RecordPolling(SourceName, t0, t1, true)
	s.markServiceEntryInitDone()
}

func (s *Source) iface(service string) {
	providers, err := s.Con.Children(s.args.RegistryRootNode + "/" + service + "/" + ProviderNode)
	monitoring.RecordSourceClientRequest(SourceName, err == nil)
	if err != nil {
		log.Errorf("zk %s get provider error: %s", service, err.Error())
		return
	}

	var consumers []string
	if s.args.GatewayModel {
		consumers = make([]string, 0)
	} else {
		consumers, err = s.Con.Children(s.args.RegistryRootNode + "/" + service + "/" + ConsumerNode)
		monitoring.RecordSourceClientRequest(SourceName, err == nil)
		if err != nil {
			log.Debugf("zk %s get consumer error: %s", service, err.Error())
		}
	}

	var configurators []string
	if s.args.EnableConfiguratorMeta {
		configurators, err = s.Con.Children(s.args.RegistryRootNode + "/" + service + "/" + ConfiguratorNode)
		if err != nil {
			log.Debugf("zk %s get configurator error: %s", service, err.Error())
		}
	}

	s.handleServiceData(s.pollingCache, providers, consumers, configurators, service)
}

func (s *Source) handleNodeDelete(childrens []string) {
	existMap := make(map[string]string)
	for _, child := range childrens {
		existMap[child] = child
	}
	deleteKey := make([]string, 0)
	for service := range s.pollingCache.Items() {
		if _, exist := existMap[service]; !exist {
			deleteKey = append(deleteKey, service)
		}
	}

	for _, service := range deleteKey {
		if seCache, ok := s.pollingCache.Get(service); ok {
			if ses, castok := seCache.(cmap.ConcurrentMap); castok {
				for k, v := range ses.Items() {
					seValue, ok := v.(*ServiceEntryWithMeta)
					if !ok {
						log.Errorf("cast se failed, key: %s", k)
						continue
					}

					if len(seValue.ServiceEntry.Endpoints) == 0 {
						continue
					}

					// DELETE ==> empty endpoints

					seValueCopy := *seValue
					seCopy := *seValue.ServiceEntry
					seCopy.Endpoints = make([]*networking.WorkloadEntry, 0)
					seValueCopy.ServiceEntry = &seCopy
					ses.Set(k, &seValueCopy)
					event, err := buildServiceEntryEvent(event.Updated, seValue.ServiceEntry, seValue.Meta, nil)
					if err == nil {
						log.Infof("delete(update) zk se, hosts: %s, ep size: %d ", seValue.ServiceEntry.Hosts[0], len(seValue.ServiceEntry.Endpoints))
						for _, h := range s.handlers {
							h.Handle(event)
						}
					} else {
						log.Errorf("delete(update) svc %s failed, case: %v", k, err.Error())
					}
					monitoring.RecordServiceEntryDeletion(SourceName, false, err == nil)
				}
			}
		}
	}
}

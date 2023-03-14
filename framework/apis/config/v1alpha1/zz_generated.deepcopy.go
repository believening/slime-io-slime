//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import ()

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bundle) DeepCopyInto(out *Bundle) {
	*out = *in
	if in.Modules != nil {
		in, out := &in.Modules, &out.Modules
		*out = make([]*Bundle_Item, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Bundle_Item)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bundle.
func (in *Bundle) DeepCopy() *Bundle {
	if in == nil {
		return nil
	}
	out := new(Bundle)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bundle_Item) DeepCopyInto(out *Bundle_Item) {
	*out = *in
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bundle_Item.
func (in *Bundle_Item) DeepCopy() *Bundle_Item {
	if in == nil {
		return nil
	}
	out := new(Bundle_Item)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientGoTokenBucket) DeepCopyInto(out *ClientGoTokenBucket) {
	*out = *in
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientGoTokenBucket.
func (in *ClientGoTokenBucket) DeepCopy() *ClientGoTokenBucket {
	if in == nil {
		return nil
	}
	out := new(ClientGoTokenBucket)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Component) DeepCopyInto(out *Component) {
	*out = *in
	if in.GlobalSidecarNamespace != nil {
		in, out := &in.GlobalSidecarNamespace, &out.GlobalSidecarNamespace
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Component.
func (in *Component) DeepCopy() *Component {
	if in == nil {
		return nil
	}
	out := new(Component)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Config) DeepCopyInto(out *Config) {
	*out = *in
	if in.Global != nil {
		in, out := &in.Global, &out.Global
		*out = new(Global)
		(*in).DeepCopyInto(*out)
	}
	if in.Metric != nil {
		in, out := &in.Metric, &out.Metric
		*out = new(Metric)
		(*in).DeepCopyInto(*out)
	}
	if in.General != nil {
		in, out := &in.General, &out.General
		*out = new(General)
		(*in).DeepCopyInto(*out)
	}
	if in.Bundle != nil {
		in, out := &in.Bundle, &out.Bundle
		*out = new(Bundle)
		(*in).DeepCopyInto(*out)
	}
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Config.
func (in *Config) DeepCopy() *Config {
	if in == nil {
		return nil
	}
	out := new(Config)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigSource) DeepCopyInto(out *ConfigSource) {
	*out = *in
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigSource.
func (in *ConfigSource) DeepCopy() *ConfigSource {
	if in == nil {
		return nil
	}
	out := new(ConfigSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *General) DeepCopyInto(out *General) {
	*out = *in
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new General.
func (in *General) DeepCopy() *General {
	if in == nil {
		return nil
	}
	out := new(General)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Global) DeepCopyInto(out *Global) {
	*out = *in
	if in.Log != nil {
		in, out := &in.Log, &out.Log
		*out = new(Log)
		(*in).DeepCopyInto(*out)
	}
	if in.Misc != nil {
		in, out := &in.Misc, &out.Misc
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ConfigSources != nil {
		in, out := &in.ConfigSources, &out.ConfigSources
		*out = make([]*ConfigSource, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ConfigSource)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.ClientGoTokenBucket != nil {
		in, out := &in.ClientGoTokenBucket, &out.ClientGoTokenBucket
		*out = new(ClientGoTokenBucket)
		(*in).DeepCopyInto(*out)
	}
	if in.IstioConfigSource != nil {
		in, out := &in.IstioConfigSource, &out.IstioConfigSource
		*out = new(ConfigSource)
		(*in).DeepCopyInto(*out)
	}
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Global.
func (in *Global) DeepCopy() *Global {
	if in == nil {
		return nil
	}
	out := new(Global)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *K8S_Source) DeepCopyInto(out *K8S_Source) {
	*out = *in
	if in.Handlers != nil {
		in, out := &in.Handlers, &out.Handlers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K8S_Source.
func (in *K8S_Source) DeepCopy() *K8S_Source {
	if in == nil {
		return nil
	}
	out := new(K8S_Source)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Log) DeepCopyInto(out *Log) {
	*out = *in
	if in.LogRotateConfig != nil {
		in, out := &in.LogRotateConfig, &out.LogRotateConfig
		*out = new(LogRotateConfig)
		(*in).DeepCopyInto(*out)
	}
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Log.
func (in *Log) DeepCopy() *Log {
	if in == nil {
		return nil
	}
	out := new(Log)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogRotateConfig) DeepCopyInto(out *LogRotateConfig) {
	*out = *in
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogRotateConfig.
func (in *LogRotateConfig) DeepCopy() *LogRotateConfig {
	if in == nil {
		return nil
	}
	out := new(LogRotateConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Metric) DeepCopyInto(out *Metric) {
	*out = *in
	if in.Prometheus != nil {
		in, out := &in.Prometheus, &out.Prometheus
		*out = new(Prometheus_Source)
		(*in).DeepCopyInto(*out)
	}
	if in.K8S != nil {
		in, out := &in.K8S, &out.K8S
		*out = new(K8S_Source)
		(*in).DeepCopyInto(*out)
	}
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Metric.
func (in *Metric) DeepCopy() *Metric {
	if in == nil {
		return nil
	}
	out := new(Metric)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Prometheus_Source) DeepCopyInto(out *Prometheus_Source) {
	*out = *in
	if in.Handlers != nil {
		in, out := &in.Handlers, &out.Handlers
		*out = make(map[string]*Prometheus_Source_Handler, len(*in))
		for key, val := range *in {
			var outVal *Prometheus_Source_Handler
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(Prometheus_Source_Handler)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Prometheus_Source.
func (in *Prometheus_Source) DeepCopy() *Prometheus_Source {
	if in == nil {
		return nil
	}
	out := new(Prometheus_Source)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Prometheus_Source_Handler) DeepCopyInto(out *Prometheus_Source_Handler) {
	*out = *in
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Prometheus_Source_Handler.
func (in *Prometheus_Source_Handler) DeepCopy() *Prometheus_Source_Handler {
	if in == nil {
		return nil
	}
	out := new(Prometheus_Source_Handler)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlimeBootSpec) DeepCopyInto(out *SlimeBootSpec) {
	*out = *in
	if in.Module != nil {
		in, out := &in.Module, &out.Module
		*out = make([]*Config, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Config)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Component != nil {
		in, out := &in.Component, &out.Component
		*out = new(Component)
		(*in).DeepCopyInto(*out)
	}
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlimeBootSpec.
func (in *SlimeBootSpec) DeepCopy() *SlimeBootSpec {
	if in == nil {
		return nil
	}
	out := new(SlimeBootSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlimeBootStatus) DeepCopyInto(out *SlimeBootStatus) {
	*out = *in
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlimeBootStatus.
func (in *SlimeBootStatus) DeepCopy() *SlimeBootStatus {
	if in == nil {
		return nil
	}
	out := new(SlimeBootStatus)
	in.DeepCopyInto(out)
	return out
}

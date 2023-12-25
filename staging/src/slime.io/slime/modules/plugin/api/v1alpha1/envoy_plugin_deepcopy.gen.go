// Code generated by protoc-gen-deepcopy. DO NOT EDIT.
package v1alpha1

import (
	proto "google.golang.org/protobuf/proto"
)

// DeepCopyInto supports using WorkloadSelector within kubernetes types, where deepcopy-gen is used.
func (in *WorkloadSelector) DeepCopyInto(out *WorkloadSelector) {
	p := proto.Clone(in).(*WorkloadSelector)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadSelector. Required by controller-gen.
func (in *WorkloadSelector) DeepCopy() *WorkloadSelector {
	if in == nil {
		return nil
	}
	out := new(WorkloadSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadSelector. Required by controller-gen.
func (in *WorkloadSelector) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using EnvoyPluginSpec within kubernetes types, where deepcopy-gen is used.
func (in *EnvoyPluginSpec) DeepCopyInto(out *EnvoyPluginSpec) {
	p := proto.Clone(in).(*EnvoyPluginSpec)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvoyPluginSpec. Required by controller-gen.
func (in *EnvoyPluginSpec) DeepCopy() *EnvoyPluginSpec {
	if in == nil {
		return nil
	}
	out := new(EnvoyPluginSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new EnvoyPluginSpec. Required by controller-gen.
func (in *EnvoyPluginSpec) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using EnvoyPluginSpec_Listener within kubernetes types, where deepcopy-gen is used.
func (in *EnvoyPluginSpec_Listener) DeepCopyInto(out *EnvoyPluginSpec_Listener) {
	p := proto.Clone(in).(*EnvoyPluginSpec_Listener)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvoyPluginSpec_Listener. Required by controller-gen.
func (in *EnvoyPluginSpec_Listener) DeepCopy() *EnvoyPluginSpec_Listener {
	if in == nil {
		return nil
	}
	out := new(EnvoyPluginSpec_Listener)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new EnvoyPluginSpec_Listener. Required by controller-gen.
func (in *EnvoyPluginSpec_Listener) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

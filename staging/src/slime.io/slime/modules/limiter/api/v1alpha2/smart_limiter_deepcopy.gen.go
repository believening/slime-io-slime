// Code generated by protoc-gen-deepcopy. DO NOT EDIT.
package v1alpha2

import (
	proto "google.golang.org/protobuf/proto"
)

// DeepCopyInto supports using SmartLimiterSpec within kubernetes types, where deepcopy-gen is used.
func (in *SmartLimiterSpec) DeepCopyInto(out *SmartLimiterSpec) {
	p := proto.Clone(in).(*SmartLimiterSpec)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SmartLimiterSpec. Required by controller-gen.
func (in *SmartLimiterSpec) DeepCopy() *SmartLimiterSpec {
	if in == nil {
		return nil
	}
	out := new(SmartLimiterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new SmartLimiterSpec. Required by controller-gen.
func (in *SmartLimiterSpec) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using SmartLimiterStatus within kubernetes types, where deepcopy-gen is used.
func (in *SmartLimiterStatus) DeepCopyInto(out *SmartLimiterStatus) {
	p := proto.Clone(in).(*SmartLimiterStatus)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SmartLimiterStatus. Required by controller-gen.
func (in *SmartLimiterStatus) DeepCopy() *SmartLimiterStatus {
	if in == nil {
		return nil
	}
	out := new(SmartLimiterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new SmartLimiterStatus. Required by controller-gen.
func (in *SmartLimiterStatus) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using SmartLimitDescriptor within kubernetes types, where deepcopy-gen is used.
func (in *SmartLimitDescriptor) DeepCopyInto(out *SmartLimitDescriptor) {
	p := proto.Clone(in).(*SmartLimitDescriptor)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SmartLimitDescriptor. Required by controller-gen.
func (in *SmartLimitDescriptor) DeepCopy() *SmartLimitDescriptor {
	if in == nil {
		return nil
	}
	out := new(SmartLimitDescriptor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new SmartLimitDescriptor. Required by controller-gen.
func (in *SmartLimitDescriptor) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using SmartLimitDescriptor_Matcher within kubernetes types, where deepcopy-gen is used.
func (in *SmartLimitDescriptor_Matcher) DeepCopyInto(out *SmartLimitDescriptor_Matcher) {
	p := proto.Clone(in).(*SmartLimitDescriptor_Matcher)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SmartLimitDescriptor_Matcher. Required by controller-gen.
func (in *SmartLimitDescriptor_Matcher) DeepCopy() *SmartLimitDescriptor_Matcher {
	if in == nil {
		return nil
	}
	out := new(SmartLimitDescriptor_Matcher)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new SmartLimitDescriptor_Matcher. Required by controller-gen.
func (in *SmartLimitDescriptor_Matcher) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using SmartLimitDescriptor_Action within kubernetes types, where deepcopy-gen is used.
func (in *SmartLimitDescriptor_Action) DeepCopyInto(out *SmartLimitDescriptor_Action) {
	p := proto.Clone(in).(*SmartLimitDescriptor_Action)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SmartLimitDescriptor_Action. Required by controller-gen.
func (in *SmartLimitDescriptor_Action) DeepCopy() *SmartLimitDescriptor_Action {
	if in == nil {
		return nil
	}
	out := new(SmartLimitDescriptor_Action)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new SmartLimitDescriptor_Action. Required by controller-gen.
func (in *SmartLimitDescriptor_Action) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using SmartLimitDescriptors within kubernetes types, where deepcopy-gen is used.
func (in *SmartLimitDescriptors) DeepCopyInto(out *SmartLimitDescriptors) {
	p := proto.Clone(in).(*SmartLimitDescriptors)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SmartLimitDescriptors. Required by controller-gen.
func (in *SmartLimitDescriptors) DeepCopy() *SmartLimitDescriptors {
	if in == nil {
		return nil
	}
	out := new(SmartLimitDescriptors)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new SmartLimitDescriptors. Required by controller-gen.
func (in *SmartLimitDescriptors) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Duration within kubernetes types, where deepcopy-gen is used.
func (in *Duration) DeepCopyInto(out *Duration) {
	p := proto.Clone(in).(*Duration)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Duration. Required by controller-gen.
func (in *Duration) DeepCopy() *Duration {
	if in == nil {
		return nil
	}
	out := new(Duration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Duration. Required by controller-gen.
func (in *Duration) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Target within kubernetes types, where deepcopy-gen is used.
func (in *Target) DeepCopyInto(out *Target) {
	p := proto.Clone(in).(*Target)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Target. Required by controller-gen.
func (in *Target) DeepCopy() *Target {
	if in == nil {
		return nil
	}
	out := new(Target)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Target. Required by controller-gen.
func (in *Target) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Header within kubernetes types, where deepcopy-gen is used.
func (in *Header) DeepCopyInto(out *Header) {
	p := proto.Clone(in).(*Header)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Header. Required by controller-gen.
func (in *Header) DeepCopy() *Header {
	if in == nil {
		return nil
	}
	out := new(Header)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Header. Required by controller-gen.
func (in *Header) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

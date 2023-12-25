//
// @Author: yangdihang
// @Date: 2020/5/21

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.20.1
// source: plugin_manager.proto

package v1alpha1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Plugin_ListenerType int32

const (
	Plugin_Outbound Plugin_ListenerType = 0
	Plugin_Inbound  Plugin_ListenerType = 1
	Plugin_Gateway  Plugin_ListenerType = 2
)

// Enum value maps for Plugin_ListenerType.
var (
	Plugin_ListenerType_name = map[int32]string{
		0: "Outbound",
		1: "Inbound",
		2: "Gateway",
	}
	Plugin_ListenerType_value = map[string]int32{
		"Outbound": 0,
		"Inbound":  1,
		"Gateway":  2,
	}
)

func (x Plugin_ListenerType) Enum() *Plugin_ListenerType {
	p := new(Plugin_ListenerType)
	*p = x
	return p
}

func (x Plugin_ListenerType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Plugin_ListenerType) Descriptor() protoreflect.EnumDescriptor {
	return file_plugin_manager_proto_enumTypes[0].Descriptor()
}

func (Plugin_ListenerType) Type() protoreflect.EnumType {
	return &file_plugin_manager_proto_enumTypes[0]
}

func (x Plugin_ListenerType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Plugin_ListenerType.Descriptor instead.
func (Plugin_ListenerType) EnumDescriptor() ([]byte, []int) {
	return file_plugin_manager_proto_rawDescGZIP(), []int{1, 0}
}

type Plugin_Protocol int32

const (
	Plugin_HTTP    Plugin_Protocol = 0
	Plugin_Dubbo   Plugin_Protocol = 1
	Plugin_Generic Plugin_Protocol = 2
)

// Enum value maps for Plugin_Protocol.
var (
	Plugin_Protocol_name = map[int32]string{
		0: "HTTP",
		1: "Dubbo",
		2: "Generic",
	}
	Plugin_Protocol_value = map[string]int32{
		"HTTP":    0,
		"Dubbo":   1,
		"Generic": 2,
	}
)

func (x Plugin_Protocol) Enum() *Plugin_Protocol {
	p := new(Plugin_Protocol)
	*p = x
	return p
}

func (x Plugin_Protocol) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Plugin_Protocol) Descriptor() protoreflect.EnumDescriptor {
	return file_plugin_manager_proto_enumTypes[1].Descriptor()
}

func (Plugin_Protocol) Type() protoreflect.EnumType {
	return &file_plugin_manager_proto_enumTypes[1]
}

func (x Plugin_Protocol) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Plugin_Protocol.Descriptor instead.
func (Plugin_Protocol) EnumDescriptor() ([]byte, []int) {
	return file_plugin_manager_proto_rawDescGZIP(), []int{1, 1}
}

type PluginManagerSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Zero or more labels that indicate a specific set of pods/VMs whose
	// proxies should be configured to use these additional filters.  The
	// scope of label search is platform dependent. On Kubernetes, for
	// example, the scope includes pods running in all reachable
	// namespaces. Omitting the selector applies the filter to all proxies in
	// the mesh.
	WorkloadLabels map[string]string `protobuf:"bytes,1,rep,name=workload_labels,json=workloadLabels,proto3" json:"workload_labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Plugin         []*Plugin         `protobuf:"bytes,2,rep,name=plugin,proto3" json:"plugin,omitempty"`
	// Names of gateways where the rule should be applied to. Gateway names
	// at the top of the VirtualService (if any) are overridden. The gateway
	// match is independent of sourceLabels.
	Gateways []string `protobuf:"bytes,3,rep,name=gateways,proto3" json:"gateways,omitempty"`
	// priority defines the order in which patch sets are applied within a context.
	Priority int32 `protobuf:"varint,4,opt,name=priority,proto3" json:"priority,omitempty"`
}

func (x *PluginManagerSpec) Reset() {
	*x = PluginManagerSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_manager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PluginManagerSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PluginManagerSpec) ProtoMessage() {}

func (x *PluginManagerSpec) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_manager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PluginManagerSpec.ProtoReflect.Descriptor instead.
func (*PluginManagerSpec) Descriptor() ([]byte, []int) {
	return file_plugin_manager_proto_rawDescGZIP(), []int{0}
}

func (x *PluginManagerSpec) GetWorkloadLabels() map[string]string {
	if x != nil {
		return x.WorkloadLabels
	}
	return nil
}

func (x *PluginManagerSpec) GetPlugin() []*Plugin {
	if x != nil {
		return x.Plugin
	}
	return nil
}

func (x *PluginManagerSpec) GetGateways() []string {
	if x != nil {
		return x.Gateways
	}
	return nil
}

func (x *PluginManagerSpec) GetPriority() int32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

type Plugin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enable bool   `protobuf:"varint,1,opt,name=enable,proto3" json:"enable,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Deprecated
	Settings     *structpb.Struct    `protobuf:"bytes,3,opt,name=settings,proto3" json:"settings,omitempty"`
	ListenerType Plugin_ListenerType `protobuf:"varint,4,opt,name=listenerType,proto3,enum=slime.microservice.plugin.v1alpha1.Plugin_ListenerType" json:"listenerType,omitempty"`
	TypeUrl      string              `protobuf:"bytes,5,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	// Types that are assignable to PluginSettings:
	//
	//	*Plugin_Wasm
	//	*Plugin_Inline
	//	*Plugin_Rider
	PluginSettings isPlugin_PluginSettings `protobuf_oneof:"plugin_settings"`
	Port           uint32                  `protobuf:"varint,8,opt,name=port,proto3" json:"port,omitempty"`
	// rawPatch will patch to the generated final envoy filter config patch (EnvoyFilter_EnvoyConfigObjectPatch)
	RawPatch             *structpb.Struct `protobuf:"bytes,10,opt,name=rawPatch,proto3" json:"rawPatch,omitempty"`
	Protocol             Plugin_Protocol  `protobuf:"varint,11,opt,name=protocol,proto3,enum=slime.microservice.plugin.v1alpha1.Plugin_Protocol" json:"protocol,omitempty"`
	GenericAppProtocol   string           `protobuf:"bytes,12,opt,name=generic_app_protocol,json=genericAppProtocol,proto3" json:"generic_app_protocol,omitempty"`
	DisableOnFilterLevel bool             `protobuf:"varint,13,opt,name=disable_on_filter_level,json=disableOnFilterLevel,proto3" json:"disable_on_filter_level,omitempty"`
}

func (x *Plugin) Reset() {
	*x = Plugin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_manager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Plugin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Plugin) ProtoMessage() {}

func (x *Plugin) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_manager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Plugin.ProtoReflect.Descriptor instead.
func (*Plugin) Descriptor() ([]byte, []int) {
	return file_plugin_manager_proto_rawDescGZIP(), []int{1}
}

func (x *Plugin) GetEnable() bool {
	if x != nil {
		return x.Enable
	}
	return false
}

func (x *Plugin) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Plugin) GetSettings() *structpb.Struct {
	if x != nil {
		return x.Settings
	}
	return nil
}

func (x *Plugin) GetListenerType() Plugin_ListenerType {
	if x != nil {
		return x.ListenerType
	}
	return Plugin_Outbound
}

func (x *Plugin) GetTypeUrl() string {
	if x != nil {
		return x.TypeUrl
	}
	return ""
}

func (m *Plugin) GetPluginSettings() isPlugin_PluginSettings {
	if m != nil {
		return m.PluginSettings
	}
	return nil
}

func (x *Plugin) GetWasm() *Wasm {
	if x, ok := x.GetPluginSettings().(*Plugin_Wasm); ok {
		return x.Wasm
	}
	return nil
}

func (x *Plugin) GetInline() *Inline {
	if x, ok := x.GetPluginSettings().(*Plugin_Inline); ok {
		return x.Inline
	}
	return nil
}

func (x *Plugin) GetRider() *Rider {
	if x, ok := x.GetPluginSettings().(*Plugin_Rider); ok {
		return x.Rider
	}
	return nil
}

func (x *Plugin) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *Plugin) GetRawPatch() *structpb.Struct {
	if x != nil {
		return x.RawPatch
	}
	return nil
}

func (x *Plugin) GetProtocol() Plugin_Protocol {
	if x != nil {
		return x.Protocol
	}
	return Plugin_HTTP
}

func (x *Plugin) GetGenericAppProtocol() string {
	if x != nil {
		return x.GenericAppProtocol
	}
	return ""
}

func (x *Plugin) GetDisableOnFilterLevel() bool {
	if x != nil {
		return x.DisableOnFilterLevel
	}
	return false
}

type isPlugin_PluginSettings interface {
	isPlugin_PluginSettings()
}

type Plugin_Wasm struct {
	Wasm *Wasm `protobuf:"bytes,6,opt,name=wasm,proto3,oneof"`
}

type Plugin_Inline struct {
	// plugin compiled inside envoy
	Inline *Inline `protobuf:"bytes,7,opt,name=inline,proto3,oneof"`
}

type Plugin_Rider struct {
	Rider *Rider `protobuf:"bytes,9,opt,name=rider,proto3,oneof"`
}

func (*Plugin_Wasm) isPlugin_PluginSettings() {}

func (*Plugin_Inline) isPlugin_PluginSettings() {}

func (*Plugin_Rider) isPlugin_PluginSettings() {}

type Wasm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Settings   *structpb.Struct `protobuf:"bytes,1,opt,name=settings,proto3" json:"settings,omitempty"`
	PluginName string           `protobuf:"bytes,2,opt,name=plugin_name,json=pluginName,proto3" json:"plugin_name,omitempty"`
	Url        string           `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	Sha256     string           `protobuf:"bytes,4,opt,name=sha256,proto3" json:"sha256,omitempty"`
	// Types that are assignable to ImagePullSecret:
	//
	//	*Wasm_ImagePullSecretName
	//	*Wasm_ImagePullSecretContent
	ImagePullSecret isWasm_ImagePullSecret `protobuf_oneof:"image_pull_secret"`
}

func (x *Wasm) Reset() {
	*x = Wasm{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_manager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Wasm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Wasm) ProtoMessage() {}

func (x *Wasm) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_manager_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Wasm.ProtoReflect.Descriptor instead.
func (*Wasm) Descriptor() ([]byte, []int) {
	return file_plugin_manager_proto_rawDescGZIP(), []int{2}
}

func (x *Wasm) GetSettings() *structpb.Struct {
	if x != nil {
		return x.Settings
	}
	return nil
}

func (x *Wasm) GetPluginName() string {
	if x != nil {
		return x.PluginName
	}
	return ""
}

func (x *Wasm) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Wasm) GetSha256() string {
	if x != nil {
		return x.Sha256
	}
	return ""
}

func (m *Wasm) GetImagePullSecret() isWasm_ImagePullSecret {
	if m != nil {
		return m.ImagePullSecret
	}
	return nil
}

func (x *Wasm) GetImagePullSecretName() string {
	if x, ok := x.GetImagePullSecret().(*Wasm_ImagePullSecretName); ok {
		return x.ImagePullSecretName
	}
	return ""
}

func (x *Wasm) GetImagePullSecretContent() string {
	if x, ok := x.GetImagePullSecret().(*Wasm_ImagePullSecretContent); ok {
		return x.ImagePullSecretContent
	}
	return ""
}

type isWasm_ImagePullSecret interface {
	isWasm_ImagePullSecret()
}

type Wasm_ImagePullSecretName struct {
	ImagePullSecretName string `protobuf:"bytes,5,opt,name=image_pull_secret_name,json=imagePullSecretName,proto3,oneof"`
}

type Wasm_ImagePullSecretContent struct {
	ImagePullSecretContent string `protobuf:"bytes,6,opt,name=image_pull_secret_content,json=imagePullSecretContent,proto3,oneof"`
}

func (*Wasm_ImagePullSecretName) isWasm_ImagePullSecret() {}

func (*Wasm_ImagePullSecretContent) isWasm_ImagePullSecret() {}

type Rider struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Settings   *structpb.Struct `protobuf:"bytes,1,opt,name=settings,proto3" json:"settings,omitempty"`
	PluginName string           `protobuf:"bytes,2,opt,name=plugin_name,json=pluginName,proto3" json:"plugin_name,omitempty"`
	Url        string           `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	Sha256     string           `protobuf:"bytes,4,opt,name=sha256,proto3" json:"sha256,omitempty"`
	// Types that are assignable to ImagePullSecret:
	//
	//	*Rider_ImagePullSecretName
	//	*Rider_ImagePullSecretContent
	ImagePullSecret isRider_ImagePullSecret `protobuf_oneof:"image_pull_secret"`
}

func (x *Rider) Reset() {
	*x = Rider{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_manager_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rider) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rider) ProtoMessage() {}

func (x *Rider) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_manager_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rider.ProtoReflect.Descriptor instead.
func (*Rider) Descriptor() ([]byte, []int) {
	return file_plugin_manager_proto_rawDescGZIP(), []int{3}
}

func (x *Rider) GetSettings() *structpb.Struct {
	if x != nil {
		return x.Settings
	}
	return nil
}

func (x *Rider) GetPluginName() string {
	if x != nil {
		return x.PluginName
	}
	return ""
}

func (x *Rider) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Rider) GetSha256() string {
	if x != nil {
		return x.Sha256
	}
	return ""
}

func (m *Rider) GetImagePullSecret() isRider_ImagePullSecret {
	if m != nil {
		return m.ImagePullSecret
	}
	return nil
}

func (x *Rider) GetImagePullSecretName() string {
	if x, ok := x.GetImagePullSecret().(*Rider_ImagePullSecretName); ok {
		return x.ImagePullSecretName
	}
	return ""
}

func (x *Rider) GetImagePullSecretContent() string {
	if x, ok := x.GetImagePullSecret().(*Rider_ImagePullSecretContent); ok {
		return x.ImagePullSecretContent
	}
	return ""
}

type isRider_ImagePullSecret interface {
	isRider_ImagePullSecret()
}

type Rider_ImagePullSecretName struct {
	ImagePullSecretName string `protobuf:"bytes,5,opt,name=image_pull_secret_name,json=imagePullSecretName,proto3,oneof"`
}

type Rider_ImagePullSecretContent struct {
	ImagePullSecretContent string `protobuf:"bytes,6,opt,name=image_pull_secret_content,json=imagePullSecretContent,proto3,oneof"`
}

func (*Rider_ImagePullSecretName) isRider_ImagePullSecret() {}

func (*Rider_ImagePullSecretContent) isRider_ImagePullSecret() {}

type Inline struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Settings     *structpb.Struct `protobuf:"bytes,1,opt,name=settings,proto3" json:"settings,omitempty"`
	DirectPatch  bool             `protobuf:"varint,2,opt,name=directPatch,proto3" json:"directPatch,omitempty"`
	FieldPatchTo string           `protobuf:"bytes,3,opt,name=fieldPatchTo,proto3" json:"fieldPatchTo,omitempty"`
}

func (x *Inline) Reset() {
	*x = Inline{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_manager_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Inline) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Inline) ProtoMessage() {}

func (x *Inline) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_manager_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Inline.ProtoReflect.Descriptor instead.
func (*Inline) Descriptor() ([]byte, []int) {
	return file_plugin_manager_proto_rawDescGZIP(), []int{4}
}

func (x *Inline) GetSettings() *structpb.Struct {
	if x != nil {
		return x.Settings
	}
	return nil
}

func (x *Inline) GetDirectPatch() bool {
	if x != nil {
		return x.DirectPatch
	}
	return false
}

func (x *Inline) GetFieldPatchTo() string {
	if x != nil {
		return x.FieldPatchTo
	}
	return ""
}

var File_plugin_manager_proto protoreflect.FileDescriptor

var file_plugin_manager_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x73, 0x6c, 0x69, 0x6d, 0x65, 0x2e, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc6, 0x02, 0x0a, 0x11, 0x50, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x53, 0x70, 0x65, 0x63, 0x12, 0x72,
	0x0a, 0x0f, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x49, 0x2e, 0x73, 0x6c, 0x69, 0x6d, 0x65, 0x2e,
	0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x57,
	0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0e, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x12, 0x42, 0x0a, 0x06, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x73, 0x6c, 0x69, 0x6d, 0x65, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x06,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x1a, 0x41,
	0x0a, 0x13, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0xa6, 0x06, 0x0a, 0x06, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x16, 0x0a, 0x06,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x08, 0x73, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x52, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x5b, 0x0a,
	0x0c, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x37, 0x2e, 0x73, 0x6c, 0x69, 0x6d, 0x65, 0x2e, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x6c, 0x69,
	0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x79,
	0x70, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x79,
	0x70, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x3e, 0x0a, 0x04, 0x77, 0x61, 0x73, 0x6d, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x73, 0x6c, 0x69, 0x6d, 0x65, 0x2e, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x57, 0x61, 0x73, 0x6d, 0x48, 0x00, 0x52,
	0x04, 0x77, 0x61, 0x73, 0x6d, 0x12, 0x44, 0x0a, 0x06, 0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x73, 0x6c, 0x69, 0x6d, 0x65, 0x2e, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x49, 0x6e, 0x6c, 0x69, 0x6e,
	0x65, 0x48, 0x00, 0x52, 0x06, 0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x41, 0x0a, 0x05, 0x72,
	0x69, 0x64, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x73, 0x6c, 0x69,
	0x6d, 0x65, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x52, 0x69, 0x64, 0x65, 0x72, 0x48, 0x00, 0x52, 0x05, 0x72, 0x69, 0x64, 0x65, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f,
	0x72, 0x74, 0x12, 0x33, 0x0a, 0x08, 0x72, 0x61, 0x77, 0x50, 0x61, 0x74, 0x63, 0x68, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x08, 0x72,
	0x61, 0x77, 0x50, 0x61, 0x74, 0x63, 0x68, 0x12, 0x4f, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x33, 0x2e, 0x73, 0x6c, 0x69, 0x6d,
	0x65, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x52, 0x08,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x30, 0x0a, 0x14, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x69, 0x63, 0x5f, 0x61, 0x70, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x41,
	0x70, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x35, 0x0a, 0x17, 0x64, 0x69,
	0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6f, 0x6e, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f,
	0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x64, 0x69, 0x73,
	0x61, 0x62, 0x6c, 0x65, 0x4f, 0x6e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x22, 0x36, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0c, 0x0a, 0x08, 0x4f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x10, 0x00, 0x12,
	0x0b, 0x0a, 0x07, 0x49, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07,
	0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x10, 0x02, 0x22, 0x2c, 0x0a, 0x08, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x54, 0x54, 0x50, 0x10, 0x00, 0x12,
	0x09, 0x0a, 0x05, 0x44, 0x75, 0x62, 0x62, 0x6f, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x69, 0x63, 0x10, 0x02, 0x42, 0x11, 0x0a, 0x0f, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x8f, 0x02, 0x0a, 0x04, 0x57,
	0x61, 0x73, 0x6d, 0x12, 0x33, 0x0a, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x08,
	0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x68, 0x61, 0x32, 0x35, 0x36, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x68, 0x61,
	0x32, 0x35, 0x36, 0x12, 0x35, 0x0a, 0x16, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x75, 0x6c,
	0x6c, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x13, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x50, 0x75, 0x6c, 0x6c,
	0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x19, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x5f, 0x70, 0x75, 0x6c, 0x6c, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x16, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x50, 0x75, 0x6c, 0x6c, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x13, 0x0a, 0x11, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x5f, 0x70, 0x75, 0x6c, 0x6c, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x22, 0x90, 0x02, 0x0a,
	0x05, 0x52, 0x69, 0x64, 0x65, 0x72, 0x12, 0x33, 0x0a, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x52, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x68, 0x61, 0x32, 0x35, 0x36, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x68, 0x61, 0x32, 0x35, 0x36, 0x12, 0x35, 0x0a, 0x16, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f,
	0x70, 0x75, 0x6c, 0x6c, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x13, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x50,
	0x75, 0x6c, 0x6c, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a,
	0x19, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x75, 0x6c, 0x6c, 0x5f, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x16, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x50, 0x75, 0x6c, 0x6c, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x13, 0x0a, 0x11, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x5f, 0x70, 0x75, 0x6c, 0x6c, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x22,
	0x83, 0x01, 0x0a, 0x06, 0x49, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x33, 0x0a, 0x08, 0x73, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x50, 0x61, 0x74, 0x63, 0x68, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x50, 0x61, 0x74, 0x63,
	0x68, 0x12, 0x22, 0x0a, 0x0c, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x61, 0x74, 0x63, 0x68, 0x54,
	0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x61,
	0x74, 0x63, 0x68, 0x54, 0x6f, 0x42, 0x2c, 0x5a, 0x2a, 0x73, 0x6c, 0x69, 0x6d, 0x65, 0x2e, 0x69,
	0x6f, 0x2f, 0x73, 0x6c, 0x69, 0x6d, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_plugin_manager_proto_rawDescOnce sync.Once
	file_plugin_manager_proto_rawDescData = file_plugin_manager_proto_rawDesc
)

func file_plugin_manager_proto_rawDescGZIP() []byte {
	file_plugin_manager_proto_rawDescOnce.Do(func() {
		file_plugin_manager_proto_rawDescData = protoimpl.X.CompressGZIP(file_plugin_manager_proto_rawDescData)
	})
	return file_plugin_manager_proto_rawDescData
}

var file_plugin_manager_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_plugin_manager_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_plugin_manager_proto_goTypes = []interface{}{
	(Plugin_ListenerType)(0),  // 0: slime.microservice.plugin.v1alpha1.Plugin.ListenerType
	(Plugin_Protocol)(0),      // 1: slime.microservice.plugin.v1alpha1.Plugin.Protocol
	(*PluginManagerSpec)(nil), // 2: slime.microservice.plugin.v1alpha1.PluginManagerSpec
	(*Plugin)(nil),            // 3: slime.microservice.plugin.v1alpha1.Plugin
	(*Wasm)(nil),              // 4: slime.microservice.plugin.v1alpha1.Wasm
	(*Rider)(nil),             // 5: slime.microservice.plugin.v1alpha1.Rider
	(*Inline)(nil),            // 6: slime.microservice.plugin.v1alpha1.Inline
	nil,                       // 7: slime.microservice.plugin.v1alpha1.PluginManagerSpec.WorkloadLabelsEntry
	(*structpb.Struct)(nil),   // 8: google.protobuf.Struct
}
var file_plugin_manager_proto_depIdxs = []int32{
	7,  // 0: slime.microservice.plugin.v1alpha1.PluginManagerSpec.workload_labels:type_name -> slime.microservice.plugin.v1alpha1.PluginManagerSpec.WorkloadLabelsEntry
	3,  // 1: slime.microservice.plugin.v1alpha1.PluginManagerSpec.plugin:type_name -> slime.microservice.plugin.v1alpha1.Plugin
	8,  // 2: slime.microservice.plugin.v1alpha1.Plugin.settings:type_name -> google.protobuf.Struct
	0,  // 3: slime.microservice.plugin.v1alpha1.Plugin.listenerType:type_name -> slime.microservice.plugin.v1alpha1.Plugin.ListenerType
	4,  // 4: slime.microservice.plugin.v1alpha1.Plugin.wasm:type_name -> slime.microservice.plugin.v1alpha1.Wasm
	6,  // 5: slime.microservice.plugin.v1alpha1.Plugin.inline:type_name -> slime.microservice.plugin.v1alpha1.Inline
	5,  // 6: slime.microservice.plugin.v1alpha1.Plugin.rider:type_name -> slime.microservice.plugin.v1alpha1.Rider
	8,  // 7: slime.microservice.plugin.v1alpha1.Plugin.rawPatch:type_name -> google.protobuf.Struct
	1,  // 8: slime.microservice.plugin.v1alpha1.Plugin.protocol:type_name -> slime.microservice.plugin.v1alpha1.Plugin.Protocol
	8,  // 9: slime.microservice.plugin.v1alpha1.Wasm.settings:type_name -> google.protobuf.Struct
	8,  // 10: slime.microservice.plugin.v1alpha1.Rider.settings:type_name -> google.protobuf.Struct
	8,  // 11: slime.microservice.plugin.v1alpha1.Inline.settings:type_name -> google.protobuf.Struct
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_plugin_manager_proto_init() }
func file_plugin_manager_proto_init() {
	if File_plugin_manager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_plugin_manager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PluginManagerSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_plugin_manager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Plugin); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_plugin_manager_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Wasm); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_plugin_manager_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rider); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_plugin_manager_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Inline); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_plugin_manager_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Plugin_Wasm)(nil),
		(*Plugin_Inline)(nil),
		(*Plugin_Rider)(nil),
	}
	file_plugin_manager_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*Wasm_ImagePullSecretName)(nil),
		(*Wasm_ImagePullSecretContent)(nil),
	}
	file_plugin_manager_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*Rider_ImagePullSecretName)(nil),
		(*Rider_ImagePullSecretContent)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_plugin_manager_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_plugin_manager_proto_goTypes,
		DependencyIndexes: file_plugin_manager_proto_depIdxs,
		EnumInfos:         file_plugin_manager_proto_enumTypes,
		MessageInfos:      file_plugin_manager_proto_msgTypes,
	}.Build()
	File_plugin_manager_proto = out.File
	file_plugin_manager_proto_rawDesc = nil
	file_plugin_manager_proto_goTypes = nil
	file_plugin_manager_proto_depIdxs = nil
}

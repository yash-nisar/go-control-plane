// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/network/redis_proxy/v3alpha/redis_proxy.proto

package envoy_config_filter_network_redis_proxy_v3alpha

import (
	fmt "fmt"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v3alpha/core"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type RedisProxy_ConnPoolSettings_ReadPolicy int32

const (
	RedisProxy_ConnPoolSettings_MASTER         RedisProxy_ConnPoolSettings_ReadPolicy = 0
	RedisProxy_ConnPoolSettings_PREFER_MASTER  RedisProxy_ConnPoolSettings_ReadPolicy = 1
	RedisProxy_ConnPoolSettings_REPLICA        RedisProxy_ConnPoolSettings_ReadPolicy = 2
	RedisProxy_ConnPoolSettings_PREFER_REPLICA RedisProxy_ConnPoolSettings_ReadPolicy = 3
	RedisProxy_ConnPoolSettings_ANY            RedisProxy_ConnPoolSettings_ReadPolicy = 4
)

var RedisProxy_ConnPoolSettings_ReadPolicy_name = map[int32]string{
	0: "MASTER",
	1: "PREFER_MASTER",
	2: "REPLICA",
	3: "PREFER_REPLICA",
	4: "ANY",
}

var RedisProxy_ConnPoolSettings_ReadPolicy_value = map[string]int32{
	"MASTER":         0,
	"PREFER_MASTER":  1,
	"REPLICA":        2,
	"PREFER_REPLICA": 3,
	"ANY":            4,
}

func (x RedisProxy_ConnPoolSettings_ReadPolicy) String() string {
	return proto.EnumName(RedisProxy_ConnPoolSettings_ReadPolicy_name, int32(x))
}

func (RedisProxy_ConnPoolSettings_ReadPolicy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8d5ad9127b0653d7, []int{0, 0, 0}
}

type RedisProxy struct {
	StatPrefix             string                       `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	Cluster                string                       `protobuf:"bytes,2,opt,name=cluster,proto3" json:"cluster,omitempty"` // Deprecated: Do not use.
	Settings               *RedisProxy_ConnPoolSettings `protobuf:"bytes,3,opt,name=settings,proto3" json:"settings,omitempty"`
	LatencyInMicros        bool                         `protobuf:"varint,4,opt,name=latency_in_micros,json=latencyInMicros,proto3" json:"latency_in_micros,omitempty"`
	PrefixRoutes           *RedisProxy_PrefixRoutes     `protobuf:"bytes,5,opt,name=prefix_routes,json=prefixRoutes,proto3" json:"prefix_routes,omitempty"`
	DownstreamAuthPassword *core.DataSource             `protobuf:"bytes,6,opt,name=downstream_auth_password,json=downstreamAuthPassword,proto3" json:"downstream_auth_password,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                     `json:"-"`
	XXX_unrecognized       []byte                       `json:"-"`
	XXX_sizecache          int32                        `json:"-"`
}

func (m *RedisProxy) Reset()         { *m = RedisProxy{} }
func (m *RedisProxy) String() string { return proto.CompactTextString(m) }
func (*RedisProxy) ProtoMessage()    {}
func (*RedisProxy) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d5ad9127b0653d7, []int{0}
}

func (m *RedisProxy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedisProxy.Unmarshal(m, b)
}
func (m *RedisProxy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedisProxy.Marshal(b, m, deterministic)
}
func (m *RedisProxy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedisProxy.Merge(m, src)
}
func (m *RedisProxy) XXX_Size() int {
	return xxx_messageInfo_RedisProxy.Size(m)
}
func (m *RedisProxy) XXX_DiscardUnknown() {
	xxx_messageInfo_RedisProxy.DiscardUnknown(m)
}

var xxx_messageInfo_RedisProxy proto.InternalMessageInfo

func (m *RedisProxy) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

// Deprecated: Do not use.
func (m *RedisProxy) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *RedisProxy) GetSettings() *RedisProxy_ConnPoolSettings {
	if m != nil {
		return m.Settings
	}
	return nil
}

func (m *RedisProxy) GetLatencyInMicros() bool {
	if m != nil {
		return m.LatencyInMicros
	}
	return false
}

func (m *RedisProxy) GetPrefixRoutes() *RedisProxy_PrefixRoutes {
	if m != nil {
		return m.PrefixRoutes
	}
	return nil
}

func (m *RedisProxy) GetDownstreamAuthPassword() *core.DataSource {
	if m != nil {
		return m.DownstreamAuthPassword
	}
	return nil
}

type RedisProxy_ConnPoolSettings struct {
	OpTimeout                     *duration.Duration                     `protobuf:"bytes,1,opt,name=op_timeout,json=opTimeout,proto3" json:"op_timeout,omitempty"`
	EnableHashtagging             bool                                   `protobuf:"varint,2,opt,name=enable_hashtagging,json=enableHashtagging,proto3" json:"enable_hashtagging,omitempty"`
	EnableRedirection             bool                                   `protobuf:"varint,3,opt,name=enable_redirection,json=enableRedirection,proto3" json:"enable_redirection,omitempty"`
	MaxBufferSizeBeforeFlush      uint32                                 `protobuf:"varint,4,opt,name=max_buffer_size_before_flush,json=maxBufferSizeBeforeFlush,proto3" json:"max_buffer_size_before_flush,omitempty"`
	BufferFlushTimeout            *duration.Duration                     `protobuf:"bytes,5,opt,name=buffer_flush_timeout,json=bufferFlushTimeout,proto3" json:"buffer_flush_timeout,omitempty"`
	MaxUpstreamUnknownConnections *wrappers.UInt32Value                  `protobuf:"bytes,6,opt,name=max_upstream_unknown_connections,json=maxUpstreamUnknownConnections,proto3" json:"max_upstream_unknown_connections,omitempty"`
	ReadPolicy                    RedisProxy_ConnPoolSettings_ReadPolicy `protobuf:"varint,7,opt,name=read_policy,json=readPolicy,proto3,enum=envoy.config.filter.network.redis_proxy.v3alpha.RedisProxy_ConnPoolSettings_ReadPolicy" json:"read_policy,omitempty"`
	XXX_NoUnkeyedLiteral          struct{}                               `json:"-"`
	XXX_unrecognized              []byte                                 `json:"-"`
	XXX_sizecache                 int32                                  `json:"-"`
}

func (m *RedisProxy_ConnPoolSettings) Reset()         { *m = RedisProxy_ConnPoolSettings{} }
func (m *RedisProxy_ConnPoolSettings) String() string { return proto.CompactTextString(m) }
func (*RedisProxy_ConnPoolSettings) ProtoMessage()    {}
func (*RedisProxy_ConnPoolSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d5ad9127b0653d7, []int{0, 0}
}

func (m *RedisProxy_ConnPoolSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedisProxy_ConnPoolSettings.Unmarshal(m, b)
}
func (m *RedisProxy_ConnPoolSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedisProxy_ConnPoolSettings.Marshal(b, m, deterministic)
}
func (m *RedisProxy_ConnPoolSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedisProxy_ConnPoolSettings.Merge(m, src)
}
func (m *RedisProxy_ConnPoolSettings) XXX_Size() int {
	return xxx_messageInfo_RedisProxy_ConnPoolSettings.Size(m)
}
func (m *RedisProxy_ConnPoolSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_RedisProxy_ConnPoolSettings.DiscardUnknown(m)
}

var xxx_messageInfo_RedisProxy_ConnPoolSettings proto.InternalMessageInfo

func (m *RedisProxy_ConnPoolSettings) GetOpTimeout() *duration.Duration {
	if m != nil {
		return m.OpTimeout
	}
	return nil
}

func (m *RedisProxy_ConnPoolSettings) GetEnableHashtagging() bool {
	if m != nil {
		return m.EnableHashtagging
	}
	return false
}

func (m *RedisProxy_ConnPoolSettings) GetEnableRedirection() bool {
	if m != nil {
		return m.EnableRedirection
	}
	return false
}

func (m *RedisProxy_ConnPoolSettings) GetMaxBufferSizeBeforeFlush() uint32 {
	if m != nil {
		return m.MaxBufferSizeBeforeFlush
	}
	return 0
}

func (m *RedisProxy_ConnPoolSettings) GetBufferFlushTimeout() *duration.Duration {
	if m != nil {
		return m.BufferFlushTimeout
	}
	return nil
}

func (m *RedisProxy_ConnPoolSettings) GetMaxUpstreamUnknownConnections() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxUpstreamUnknownConnections
	}
	return nil
}

func (m *RedisProxy_ConnPoolSettings) GetReadPolicy() RedisProxy_ConnPoolSettings_ReadPolicy {
	if m != nil {
		return m.ReadPolicy
	}
	return RedisProxy_ConnPoolSettings_MASTER
}

type RedisProxy_PrefixRoutes struct {
	Routes               []*RedisProxy_PrefixRoutes_Route `protobuf:"bytes,1,rep,name=routes,proto3" json:"routes,omitempty"`
	CaseInsensitive      bool                             `protobuf:"varint,2,opt,name=case_insensitive,json=caseInsensitive,proto3" json:"case_insensitive,omitempty"`
	CatchAllCluster      string                           `protobuf:"bytes,3,opt,name=catch_all_cluster,json=catchAllCluster,proto3" json:"catch_all_cluster,omitempty"` // Deprecated: Do not use.
	CatchAllRoute        *RedisProxy_PrefixRoutes_Route   `protobuf:"bytes,4,opt,name=catch_all_route,json=catchAllRoute,proto3" json:"catch_all_route,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *RedisProxy_PrefixRoutes) Reset()         { *m = RedisProxy_PrefixRoutes{} }
func (m *RedisProxy_PrefixRoutes) String() string { return proto.CompactTextString(m) }
func (*RedisProxy_PrefixRoutes) ProtoMessage()    {}
func (*RedisProxy_PrefixRoutes) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d5ad9127b0653d7, []int{0, 1}
}

func (m *RedisProxy_PrefixRoutes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedisProxy_PrefixRoutes.Unmarshal(m, b)
}
func (m *RedisProxy_PrefixRoutes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedisProxy_PrefixRoutes.Marshal(b, m, deterministic)
}
func (m *RedisProxy_PrefixRoutes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedisProxy_PrefixRoutes.Merge(m, src)
}
func (m *RedisProxy_PrefixRoutes) XXX_Size() int {
	return xxx_messageInfo_RedisProxy_PrefixRoutes.Size(m)
}
func (m *RedisProxy_PrefixRoutes) XXX_DiscardUnknown() {
	xxx_messageInfo_RedisProxy_PrefixRoutes.DiscardUnknown(m)
}

var xxx_messageInfo_RedisProxy_PrefixRoutes proto.InternalMessageInfo

func (m *RedisProxy_PrefixRoutes) GetRoutes() []*RedisProxy_PrefixRoutes_Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

func (m *RedisProxy_PrefixRoutes) GetCaseInsensitive() bool {
	if m != nil {
		return m.CaseInsensitive
	}
	return false
}

// Deprecated: Do not use.
func (m *RedisProxy_PrefixRoutes) GetCatchAllCluster() string {
	if m != nil {
		return m.CatchAllCluster
	}
	return ""
}

func (m *RedisProxy_PrefixRoutes) GetCatchAllRoute() *RedisProxy_PrefixRoutes_Route {
	if m != nil {
		return m.CatchAllRoute
	}
	return nil
}

type RedisProxy_PrefixRoutes_Route struct {
	Prefix               string                                               `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	RemovePrefix         bool                                                 `protobuf:"varint,2,opt,name=remove_prefix,json=removePrefix,proto3" json:"remove_prefix,omitempty"`
	Cluster              string                                               `protobuf:"bytes,3,opt,name=cluster,proto3" json:"cluster,omitempty"`
	RequestMirrorPolicy  []*RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy `protobuf:"bytes,4,rep,name=request_mirror_policy,json=requestMirrorPolicy,proto3" json:"request_mirror_policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                             `json:"-"`
	XXX_unrecognized     []byte                                               `json:"-"`
	XXX_sizecache        int32                                                `json:"-"`
}

func (m *RedisProxy_PrefixRoutes_Route) Reset()         { *m = RedisProxy_PrefixRoutes_Route{} }
func (m *RedisProxy_PrefixRoutes_Route) String() string { return proto.CompactTextString(m) }
func (*RedisProxy_PrefixRoutes_Route) ProtoMessage()    {}
func (*RedisProxy_PrefixRoutes_Route) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d5ad9127b0653d7, []int{0, 1, 0}
}

func (m *RedisProxy_PrefixRoutes_Route) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedisProxy_PrefixRoutes_Route.Unmarshal(m, b)
}
func (m *RedisProxy_PrefixRoutes_Route) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedisProxy_PrefixRoutes_Route.Marshal(b, m, deterministic)
}
func (m *RedisProxy_PrefixRoutes_Route) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedisProxy_PrefixRoutes_Route.Merge(m, src)
}
func (m *RedisProxy_PrefixRoutes_Route) XXX_Size() int {
	return xxx_messageInfo_RedisProxy_PrefixRoutes_Route.Size(m)
}
func (m *RedisProxy_PrefixRoutes_Route) XXX_DiscardUnknown() {
	xxx_messageInfo_RedisProxy_PrefixRoutes_Route.DiscardUnknown(m)
}

var xxx_messageInfo_RedisProxy_PrefixRoutes_Route proto.InternalMessageInfo

func (m *RedisProxy_PrefixRoutes_Route) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *RedisProxy_PrefixRoutes_Route) GetRemovePrefix() bool {
	if m != nil {
		return m.RemovePrefix
	}
	return false
}

func (m *RedisProxy_PrefixRoutes_Route) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *RedisProxy_PrefixRoutes_Route) GetRequestMirrorPolicy() []*RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy {
	if m != nil {
		return m.RequestMirrorPolicy
	}
	return nil
}

type RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy struct {
	Cluster              string                         `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	RuntimeFraction      *core.RuntimeFractionalPercent `protobuf:"bytes,2,opt,name=runtime_fraction,json=runtimeFraction,proto3" json:"runtime_fraction,omitempty"`
	ExcludeReadCommands  bool                           `protobuf:"varint,3,opt,name=exclude_read_commands,json=excludeReadCommands,proto3" json:"exclude_read_commands,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy) Reset() {
	*m = RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy{}
}
func (m *RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy) String() string {
	return proto.CompactTextString(m)
}
func (*RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy) ProtoMessage() {}
func (*RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d5ad9127b0653d7, []int{0, 1, 0, 0}
}

func (m *RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy.Unmarshal(m, b)
}
func (m *RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy.Marshal(b, m, deterministic)
}
func (m *RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy.Merge(m, src)
}
func (m *RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy) XXX_Size() int {
	return xxx_messageInfo_RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy.Size(m)
}
func (m *RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy proto.InternalMessageInfo

func (m *RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy) GetRuntimeFraction() *core.RuntimeFractionalPercent {
	if m != nil {
		return m.RuntimeFraction
	}
	return nil
}

func (m *RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy) GetExcludeReadCommands() bool {
	if m != nil {
		return m.ExcludeReadCommands
	}
	return false
}

type RedisProtocolOptions struct {
	AuthPassword         *core.DataSource `protobuf:"bytes,1,opt,name=auth_password,json=authPassword,proto3" json:"auth_password,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *RedisProtocolOptions) Reset()         { *m = RedisProtocolOptions{} }
func (m *RedisProtocolOptions) String() string { return proto.CompactTextString(m) }
func (*RedisProtocolOptions) ProtoMessage()    {}
func (*RedisProtocolOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d5ad9127b0653d7, []int{1}
}

func (m *RedisProtocolOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedisProtocolOptions.Unmarshal(m, b)
}
func (m *RedisProtocolOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedisProtocolOptions.Marshal(b, m, deterministic)
}
func (m *RedisProtocolOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedisProtocolOptions.Merge(m, src)
}
func (m *RedisProtocolOptions) XXX_Size() int {
	return xxx_messageInfo_RedisProtocolOptions.Size(m)
}
func (m *RedisProtocolOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_RedisProtocolOptions.DiscardUnknown(m)
}

var xxx_messageInfo_RedisProtocolOptions proto.InternalMessageInfo

func (m *RedisProtocolOptions) GetAuthPassword() *core.DataSource {
	if m != nil {
		return m.AuthPassword
	}
	return nil
}

func init() {
	proto.RegisterEnum("envoy.config.filter.network.redis_proxy.v3alpha.RedisProxy_ConnPoolSettings_ReadPolicy", RedisProxy_ConnPoolSettings_ReadPolicy_name, RedisProxy_ConnPoolSettings_ReadPolicy_value)
	proto.RegisterType((*RedisProxy)(nil), "envoy.config.filter.network.redis_proxy.v3alpha.RedisProxy")
	proto.RegisterType((*RedisProxy_ConnPoolSettings)(nil), "envoy.config.filter.network.redis_proxy.v3alpha.RedisProxy.ConnPoolSettings")
	proto.RegisterType((*RedisProxy_PrefixRoutes)(nil), "envoy.config.filter.network.redis_proxy.v3alpha.RedisProxy.PrefixRoutes")
	proto.RegisterType((*RedisProxy_PrefixRoutes_Route)(nil), "envoy.config.filter.network.redis_proxy.v3alpha.RedisProxy.PrefixRoutes.Route")
	proto.RegisterType((*RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy)(nil), "envoy.config.filter.network.redis_proxy.v3alpha.RedisProxy.PrefixRoutes.Route.RequestMirrorPolicy")
	proto.RegisterType((*RedisProtocolOptions)(nil), "envoy.config.filter.network.redis_proxy.v3alpha.RedisProtocolOptions")
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/redis_proxy/v3alpha/redis_proxy.proto", fileDescriptor_8d5ad9127b0653d7)
}

var fileDescriptor_8d5ad9127b0653d7 = []byte{
	// 979 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x51, 0x6f, 0x1b, 0x45,
	0x10, 0xe6, 0xec, 0xc4, 0x49, 0xc7, 0x71, 0xec, 0x6c, 0xda, 0x62, 0xac, 0x80, 0xdc, 0xf0, 0x62,
	0x90, 0x38, 0x23, 0xe7, 0x19, 0x24, 0x3b, 0x4d, 0x68, 0x44, 0x53, 0xac, 0x4d, 0x03, 0x42, 0x20,
	0xad, 0xd6, 0xe7, 0xb5, 0xbd, 0xea, 0xdd, 0xee, 0xb1, 0xbb, 0xe7, 0x38, 0x7d, 0xac, 0x78, 0xe2,
	0x07, 0x20, 0x7e, 0x03, 0xff, 0x04, 0x89, 0x7f, 0xc1, 0xbf, 0xe8, 0x13, 0xba, 0xdd, 0x3d, 0xdb,
	0xa1, 0x44, 0x50, 0x35, 0x4f, 0xf6, 0xce, 0x37, 0xfb, 0xcd, 0xcc, 0xce, 0xcc, 0x77, 0xd0, 0x67,
	0x62, 0x2e, 0xaf, 0xbb, 0x91, 0x14, 0x13, 0x3e, 0xed, 0x4e, 0x78, 0x6c, 0x98, 0xea, 0x0a, 0x66,
	0xae, 0xa4, 0x7a, 0xd1, 0x55, 0x6c, 0xcc, 0x35, 0x49, 0x95, 0x5c, 0x5c, 0x77, 0xe7, 0x47, 0x34,
	0x4e, 0x67, 0x74, 0xdd, 0x16, 0xa6, 0x4a, 0x1a, 0x89, 0xba, 0x96, 0x22, 0x74, 0x14, 0xa1, 0xa3,
	0x08, 0x3d, 0x45, 0xb8, 0xee, 0xee, 0x29, 0x5a, 0x8f, 0x5c, 0x4c, 0x9a, 0xf2, 0x25, 0x6b, 0x24,
	0x15, 0xeb, 0x8e, 0xa8, 0x66, 0x8e, 0xb3, 0xf5, 0xd1, 0x54, 0xca, 0x69, 0xcc, 0xba, 0xf6, 0x34,
	0xca, 0x26, 0xdd, 0x71, 0xa6, 0xa8, 0xe1, 0x52, 0xdc, 0x86, 0x5f, 0x29, 0x9a, 0xa6, 0x4c, 0x69,
	0x8f, 0xbf, 0x3f, 0xa7, 0x31, 0x1f, 0x53, 0xc3, 0xba, 0xc5, 0x1f, 0x07, 0x1c, 0xfe, 0xb5, 0x0b,
	0x80, 0xf3, 0x9c, 0x86, 0x79, 0x4a, 0xa8, 0x03, 0x55, 0x6d, 0xa8, 0x21, 0xa9, 0x62, 0x13, 0xbe,
	0x68, 0x06, 0xed, 0xa0, 0x73, 0x6f, 0xb0, 0xf5, 0x7a, 0xb0, 0xa1, 0x4a, 0xed, 0x00, 0x43, 0x8e,
	0x0d, 0x2d, 0x84, 0x0e, 0x60, 0x2b, 0x8a, 0x33, 0x6d, 0x98, 0x6a, 0x96, 0xac, 0x57, 0xa9, 0x19,
	0xe0, 0xc2, 0x84, 0x14, 0x6c, 0x6b, 0x66, 0x0c, 0x17, 0x53, 0xdd, 0x2c, 0xb7, 0x83, 0x4e, 0xb5,
	0xf7, 0x34, 0x7c, 0xcb, 0x67, 0x09, 0x57, 0x69, 0x85, 0xc7, 0x52, 0x88, 0xa1, 0x94, 0xf1, 0x85,
	0xe7, 0x1c, 0x6c, 0xbf, 0x1e, 0x6c, 0xfe, 0x12, 0x94, 0x1a, 0x01, 0x5e, 0xc6, 0x41, 0x9f, 0xc2,
	0x5e, 0x4c, 0x0d, 0x13, 0xd1, 0x35, 0xe1, 0x82, 0x24, 0x3c, 0x52, 0x52, 0x37, 0x37, 0xda, 0x41,
	0x67, 0x1b, 0xd7, 0x3d, 0x70, 0x26, 0xce, 0xad, 0x19, 0x25, 0x50, 0x73, 0x25, 0x12, 0x25, 0x33,
	0xc3, 0x74, 0x73, 0xd3, 0x26, 0xf9, 0xe4, 0x5d, 0x92, 0x74, 0x0f, 0x83, 0x2d, 0x1f, 0xde, 0x49,
	0xd7, 0x4e, 0xe8, 0x47, 0x68, 0x8e, 0xe5, 0x95, 0xd0, 0x46, 0x31, 0x9a, 0x10, 0x9a, 0x99, 0x19,
	0x49, 0xa9, 0xd6, 0x57, 0x52, 0x8d, 0x9b, 0x15, 0x1b, 0xf9, 0xd0, 0x47, 0xa6, 0x29, 0x5f, 0x72,
	0xe7, 0x43, 0x10, 0x3e, 0xa6, 0x86, 0x5e, 0xc8, 0x4c, 0x45, 0x0c, 0x3f, 0x5c, 0x71, 0xf4, 0x33,
	0x33, 0x1b, 0x7a, 0x86, 0xd6, 0xcf, 0x9b, 0xd0, 0xf8, 0xe7, 0x0b, 0xa1, 0x01, 0x80, 0x4c, 0x89,
	0xe1, 0x09, 0x93, 0x99, 0xb1, 0x8d, 0xac, 0xf6, 0x3e, 0x08, 0xdd, 0x98, 0x84, 0xc5, 0x98, 0x84,
	0x8f, 0xfd, 0x18, 0xd9, 0x07, 0xfd, 0x3d, 0x28, 0x6d, 0x07, 0xf8, 0x9e, 0x4c, 0x9f, 0xbb, 0x5b,
	0xe8, 0x33, 0x40, 0x4c, 0xd0, 0x51, 0xcc, 0xc8, 0x8c, 0xea, 0x99, 0xa1, 0xd3, 0x29, 0x17, 0x53,
	0xdb, 0xee, 0x6d, 0xbc, 0xe7, 0x90, 0x27, 0x2b, 0x60, 0xcd, 0x3d, 0x7f, 0x29, 0xc5, 0xa2, 0x9c,
	0xd9, 0xb6, 0x7f, 0xe9, 0x8e, 0x57, 0x00, 0xfa, 0x12, 0x0e, 0x12, 0xba, 0x20, 0xa3, 0x6c, 0x32,
	0x61, 0x8a, 0x68, 0xfe, 0x92, 0x91, 0x11, 0x9b, 0x48, 0xc5, 0xc8, 0x24, 0xce, 0xf4, 0xcc, 0xb6,
	0xae, 0x86, 0x9b, 0x09, 0x5d, 0x0c, 0xac, 0xcb, 0x05, 0x7f, 0xc9, 0x06, 0xd6, 0xe1, 0x34, 0xc7,
	0xd1, 0xd7, 0x70, 0xdf, 0xdf, 0xb5, 0xfe, 0xcb, 0x5a, 0x37, 0xff, 0xa3, 0x56, 0x8c, 0xdc, 0x35,
	0xcb, 0x52, 0x94, 0xca, 0xa0, 0x9d, 0x27, 0x93, 0xa5, 0xbe, 0x47, 0x99, 0x78, 0x21, 0xe4, 0x95,
	0x20, 0x91, 0x14, 0xc2, 0xe5, 0xab, 0x7d, 0xa7, 0x0e, 0xde, 0x20, 0xbe, 0x3c, 0x13, 0xe6, 0xa8,
	0xf7, 0x2d, 0x8d, 0x33, 0x86, 0x3f, 0x4c, 0xe8, 0xe2, 0xd2, 0x93, 0x5c, 0x3a, 0x8e, 0xe3, 0x15,
	0x05, 0x7a, 0x15, 0x40, 0x55, 0x31, 0x3a, 0x26, 0xa9, 0x8c, 0x79, 0x74, 0xdd, 0xdc, 0x6a, 0x07,
	0x9d, 0xdd, 0xde, 0x77, 0x77, 0xb9, 0x1b, 0x21, 0x66, 0x74, 0x3c, 0xb4, 0xf4, 0xb6, 0xab, 0xaf,
	0xec, 0x9a, 0x80, 0x5a, 0x5a, 0x0f, 0x2f, 0xf3, 0x95, 0x2f, 0x4e, 0x08, 0xa0, 0x72, 0xde, 0xbf,
	0x78, 0x7e, 0x82, 0x1b, 0xef, 0xa1, 0x3d, 0xa8, 0x0d, 0xf1, 0xc9, 0xe9, 0x09, 0x26, 0xde, 0x14,
	0xa0, 0x2a, 0x6c, 0xe1, 0x93, 0xe1, 0xd3, 0xb3, 0xe3, 0x7e, 0xa3, 0x84, 0x10, 0xec, 0x7a, 0xbc,
	0xb0, 0x95, 0xd1, 0x16, 0x94, 0xfb, 0xcf, 0xbe, 0x6f, 0x6c, 0xb4, 0x7e, 0xab, 0xc0, 0xce, 0xfa,
	0x0e, 0xa0, 0x09, 0x54, 0xfc, 0x76, 0x05, 0xed, 0x72, 0xa7, 0xda, 0x7b, 0x76, 0x57, 0xdb, 0x15,
	0xda, 0x1f, 0xec, 0xd9, 0xd1, 0x27, 0xd0, 0x88, 0xa8, 0x66, 0x84, 0x0b, 0xcd, 0x84, 0xe6, 0x86,
	0xcf, 0x99, 0x1f, 0xd2, 0x7a, 0x6e, 0x3f, 0x5b, 0x99, 0x51, 0x08, 0x7b, 0x11, 0x35, 0xd1, 0x8c,
	0xd0, 0x38, 0x26, 0x85, 0x7e, 0x95, 0x97, 0xfa, 0x55, 0xb7, 0x60, 0x3f, 0x8e, 0x8f, 0xbd, 0x8e,
	0xcd, 0xa1, 0xbe, 0xf2, 0xb7, 0xe1, 0xec, 0x58, 0xde, 0x7d, 0x2d, 0xb5, 0x22, 0xb2, 0x3d, 0xb6,
	0xfe, 0x2c, 0xc3, 0xa6, 0xfd, 0x87, 0x1e, 0x42, 0x65, 0x5d, 0x8c, 0xb1, 0x3f, 0xa1, 0x8f, 0xa1,
	0xa6, 0x58, 0x22, 0xe7, 0xac, 0xd0, 0x6a, 0x57, 0xf1, 0x8e, 0x33, 0x7a, 0x91, 0x7e, 0xb4, 0x12,
	0xe9, 0xf2, 0x4d, 0x29, 0x5f, 0x2a, 0xf5, 0xaf, 0x01, 0x3c, 0x50, 0xec, 0xa7, 0x8c, 0x69, 0x43,
	0x12, 0xae, 0x94, 0x54, 0xc5, 0x6c, 0x6e, 0xd8, 0xa6, 0x8d, 0xee, 0xb6, 0xd0, 0x10, 0xbb, 0x58,
	0xe7, 0x36, 0x94, 0x1b, 0x41, 0xbc, 0xaf, 0xde, 0x34, 0xb6, 0xfe, 0x08, 0x60, 0xff, 0x5f, 0x9c,
	0xd7, 0x6b, 0x0a, 0x6e, 0xa9, 0xe9, 0x07, 0x68, 0xa8, 0x4c, 0xe4, 0x7a, 0x40, 0x26, 0x8a, 0x3a,
	0x19, 0x2a, 0xd9, 0xb6, 0x7d, 0x7e, 0x9b, 0xcc, 0x62, 0xe7, 0x7f, 0xea, 0xdd, 0x69, 0x3c, 0x64,
	0x2a, 0x62, 0xc2, 0xe0, 0xba, 0xba, 0x89, 0xa0, 0x1e, 0x3c, 0x60, 0x8b, 0x28, 0xce, 0xc6, 0xb9,
	0xcc, 0xd1, 0x31, 0x89, 0x64, 0x92, 0x50, 0x31, 0xd6, 0x5e, 0xe8, 0xf6, 0x3d, 0x98, 0x6f, 0xd8,
	0xb1, 0x87, 0x0e, 0x09, 0xdc, 0x2f, 0x5e, 0xc5, 0xc8, 0x48, 0xc6, 0xdf, 0xa4, 0x4e, 0x0e, 0xbe,
	0x82, 0xda, 0xcd, 0x8f, 0x41, 0xf0, 0xbf, 0x3f, 0x06, 0x3b, 0x74, 0xed, 0x13, 0x30, 0xc0, 0xf0,
	0x05, 0x97, 0xee, 0x96, 0x6b, 0xc6, 0x5b, 0x36, 0x6d, 0x50, 0x5f, 0x75, 0xcd, 0x26, 0x39, 0x0c,
	0x46, 0x15, 0x2b, 0x70, 0x47, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0xfd, 0xf8, 0x22, 0x38, 0x13,
	0x09, 0x00, 0x00,
}
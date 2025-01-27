// Code generated by protoc-gen-go. DO NOT EDIT.
// source: orderer/smartbft/configuration.proto

package smartbft

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Options_Rotation int32

const (
	Options_UNDEFINED Options_Rotation = 0
	Options_OFF       Options_Rotation = 1
	Options_ON        Options_Rotation = 2
)

var Options_Rotation_name = map[int32]string{
	0: "UNDEFINED",
	1: "OFF",
	2: "ON",
}

var Options_Rotation_value = map[string]int32{
	"UNDEFINED": 0,
	"OFF":       1,
	"ON":        2,
}

func (x Options_Rotation) String() string {
	return proto.EnumName(Options_Rotation_name, int32(x))
}

func (Options_Rotation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a8a81ac5a2771ff3, []int{2, 0}
}

// ConfigMetadata is serialized and set as the value of ConsensusType.Metadata in
// a channel configuration when the ConsensusType.Type is set "smartbft".
type ConfigMetadata struct {
	Consenters           []*Consenter `protobuf:"bytes,1,rep,name=consenters,proto3" json:"consenters,omitempty"`
	Options              *Options     `protobuf:"bytes,2,opt,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ConfigMetadata) Reset()         { *m = ConfigMetadata{} }
func (m *ConfigMetadata) String() string { return proto.CompactTextString(m) }
func (*ConfigMetadata) ProtoMessage()    {}
func (*ConfigMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8a81ac5a2771ff3, []int{0}
}

func (m *ConfigMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigMetadata.Unmarshal(m, b)
}
func (m *ConfigMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigMetadata.Marshal(b, m, deterministic)
}
func (m *ConfigMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigMetadata.Merge(m, src)
}
func (m *ConfigMetadata) XXX_Size() int {
	return xxx_messageInfo_ConfigMetadata.Size(m)
}
func (m *ConfigMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigMetadata proto.InternalMessageInfo

func (m *ConfigMetadata) GetConsenters() []*Consenter {
	if m != nil {
		return m.Consenters
	}
	return nil
}

func (m *ConfigMetadata) GetOptions() *Options {
	if m != nil {
		return m.Options
	}
	return nil
}

// Consenter represents a consenting node (i.e. replica).
type Consenter struct {
	ConsenterId          uint64   `protobuf:"varint,1,opt,name=consenter_id,json=consenterId,proto3" json:"consenter_id,omitempty"`
	Host                 string   `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	Port                 uint32   `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	MspId                string   `protobuf:"bytes,4,opt,name=msp_id,json=mspId,proto3" json:"msp_id,omitempty"`
	Identity             []byte   `protobuf:"bytes,5,opt,name=identity,proto3" json:"identity,omitempty"`
	ClientTlsCert        []byte   `protobuf:"bytes,6,opt,name=client_tls_cert,json=clientTlsCert,proto3" json:"client_tls_cert,omitempty"`
	ServerTlsCert        []byte   `protobuf:"bytes,7,opt,name=server_tls_cert,json=serverTlsCert,proto3" json:"server_tls_cert,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Consenter) Reset()         { *m = Consenter{} }
func (m *Consenter) String() string { return proto.CompactTextString(m) }
func (*Consenter) ProtoMessage()    {}
func (*Consenter) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8a81ac5a2771ff3, []int{1}
}

func (m *Consenter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Consenter.Unmarshal(m, b)
}
func (m *Consenter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Consenter.Marshal(b, m, deterministic)
}
func (m *Consenter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Consenter.Merge(m, src)
}
func (m *Consenter) XXX_Size() int {
	return xxx_messageInfo_Consenter.Size(m)
}
func (m *Consenter) XXX_DiscardUnknown() {
	xxx_messageInfo_Consenter.DiscardUnknown(m)
}

var xxx_messageInfo_Consenter proto.InternalMessageInfo

func (m *Consenter) GetConsenterId() uint64 {
	if m != nil {
		return m.ConsenterId
	}
	return 0
}

func (m *Consenter) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *Consenter) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Consenter) GetMspId() string {
	if m != nil {
		return m.MspId
	}
	return ""
}

func (m *Consenter) GetIdentity() []byte {
	if m != nil {
		return m.Identity
	}
	return nil
}

func (m *Consenter) GetClientTlsCert() []byte {
	if m != nil {
		return m.ClientTlsCert
	}
	return nil
}

func (m *Consenter) GetServerTlsCert() []byte {
	if m != nil {
		return m.ServerTlsCert
	}
	return nil
}

// Options to be specified for all the smartbft nodes. These can be modified on a
// per-channel basis.
type Options struct {
	RequestBatchMaxCount      uint64           `protobuf:"varint,2,opt,name=request_batch_max_count,json=requestBatchMaxCount,proto3" json:"request_batch_max_count,omitempty"`
	RequestBatchMaxBytes      uint64           `protobuf:"varint,3,opt,name=request_batch_max_bytes,json=requestBatchMaxBytes,proto3" json:"request_batch_max_bytes,omitempty"`
	RequestBatchMaxInterval   string           `protobuf:"bytes,4,opt,name=request_batch_max_interval,json=requestBatchMaxInterval,proto3" json:"request_batch_max_interval,omitempty"`
	IncomingMessageBufferSize uint64           `protobuf:"varint,5,opt,name=incoming_message_buffer_size,json=incomingMessageBufferSize,proto3" json:"incoming_message_buffer_size,omitempty"`
	RequestPoolSize           uint64           `protobuf:"varint,6,opt,name=request_pool_size,json=requestPoolSize,proto3" json:"request_pool_size,omitempty"`
	RequestForwardTimeout     string           `protobuf:"bytes,7,opt,name=request_forward_timeout,json=requestForwardTimeout,proto3" json:"request_forward_timeout,omitempty"`
	RequestComplainTimeout    string           `protobuf:"bytes,8,opt,name=request_complain_timeout,json=requestComplainTimeout,proto3" json:"request_complain_timeout,omitempty"`
	RequestAutoRemoveTimeout  string           `protobuf:"bytes,9,opt,name=request_auto_remove_timeout,json=requestAutoRemoveTimeout,proto3" json:"request_auto_remove_timeout,omitempty"`
	ViewChangeResendInterval  string           `protobuf:"bytes,10,opt,name=view_change_resend_interval,json=viewChangeResendInterval,proto3" json:"view_change_resend_interval,omitempty"`
	ViewChangeTimeout         string           `protobuf:"bytes,11,opt,name=view_change_timeout,json=viewChangeTimeout,proto3" json:"view_change_timeout,omitempty"`
	LeaderHeartbeatTimeout    string           `protobuf:"bytes,12,opt,name=leader_heartbeat_timeout,json=leaderHeartbeatTimeout,proto3" json:"leader_heartbeat_timeout,omitempty"`
	LeaderHeartbeatCount      uint64           `protobuf:"varint,13,opt,name=leader_heartbeat_count,json=leaderHeartbeatCount,proto3" json:"leader_heartbeat_count,omitempty"`
	CollectTimeout            string           `protobuf:"bytes,14,opt,name=collect_timeout,json=collectTimeout,proto3" json:"collect_timeout,omitempty"`
	SyncOnStart               bool             `protobuf:"varint,15,opt,name=sync_on_start,json=syncOnStart,proto3" json:"sync_on_start,omitempty"`
	SpeedUpViewChange         bool             `protobuf:"varint,16,opt,name=speed_up_view_change,json=speedUpViewChange,proto3" json:"speed_up_view_change,omitempty"`
	LeaderRotation            Options_Rotation `protobuf:"varint,17,opt,name=leader_rotation,json=leaderRotation,proto3,enum=smartbft.Options_Rotation" json:"leader_rotation,omitempty"`
	DecisionsPerLeader        uint64           `protobuf:"varint,18,opt,name=decisions_per_leader,json=decisionsPerLeader,proto3" json:"decisions_per_leader,omitempty"`
	RequestMaxBytes           uint64           `protobuf:"varint,19,opt,name=request_max_bytes,json=requestMaxBytes,proto3" json:"request_max_bytes,omitempty"`
	RequestPoolSubmitTimeout  string           `protobuf:"bytes,20,opt,name=request_pool_submit_timeout,json=requestPoolSubmitTimeout,proto3" json:"request_pool_submit_timeout,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}         `json:"-"`
	XXX_unrecognized          []byte           `json:"-"`
	XXX_sizecache             int32            `json:"-"`
}

func (m *Options) Reset()         { *m = Options{} }
func (m *Options) String() string { return proto.CompactTextString(m) }
func (*Options) ProtoMessage()    {}
func (*Options) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8a81ac5a2771ff3, []int{2}
}

func (m *Options) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Options.Unmarshal(m, b)
}
func (m *Options) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Options.Marshal(b, m, deterministic)
}
func (m *Options) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Options.Merge(m, src)
}
func (m *Options) XXX_Size() int {
	return xxx_messageInfo_Options.Size(m)
}
func (m *Options) XXX_DiscardUnknown() {
	xxx_messageInfo_Options.DiscardUnknown(m)
}

var xxx_messageInfo_Options proto.InternalMessageInfo

func (m *Options) GetRequestBatchMaxCount() uint64 {
	if m != nil {
		return m.RequestBatchMaxCount
	}
	return 0
}

func (m *Options) GetRequestBatchMaxBytes() uint64 {
	if m != nil {
		return m.RequestBatchMaxBytes
	}
	return 0
}

func (m *Options) GetRequestBatchMaxInterval() string {
	if m != nil {
		return m.RequestBatchMaxInterval
	}
	return ""
}

func (m *Options) GetIncomingMessageBufferSize() uint64 {
	if m != nil {
		return m.IncomingMessageBufferSize
	}
	return 0
}

func (m *Options) GetRequestPoolSize() uint64 {
	if m != nil {
		return m.RequestPoolSize
	}
	return 0
}

func (m *Options) GetRequestForwardTimeout() string {
	if m != nil {
		return m.RequestForwardTimeout
	}
	return ""
}

func (m *Options) GetRequestComplainTimeout() string {
	if m != nil {
		return m.RequestComplainTimeout
	}
	return ""
}

func (m *Options) GetRequestAutoRemoveTimeout() string {
	if m != nil {
		return m.RequestAutoRemoveTimeout
	}
	return ""
}

func (m *Options) GetViewChangeResendInterval() string {
	if m != nil {
		return m.ViewChangeResendInterval
	}
	return ""
}

func (m *Options) GetViewChangeTimeout() string {
	if m != nil {
		return m.ViewChangeTimeout
	}
	return ""
}

func (m *Options) GetLeaderHeartbeatTimeout() string {
	if m != nil {
		return m.LeaderHeartbeatTimeout
	}
	return ""
}

func (m *Options) GetLeaderHeartbeatCount() uint64 {
	if m != nil {
		return m.LeaderHeartbeatCount
	}
	return 0
}

func (m *Options) GetCollectTimeout() string {
	if m != nil {
		return m.CollectTimeout
	}
	return ""
}

func (m *Options) GetSyncOnStart() bool {
	if m != nil {
		return m.SyncOnStart
	}
	return false
}

func (m *Options) GetSpeedUpViewChange() bool {
	if m != nil {
		return m.SpeedUpViewChange
	}
	return false
}

func (m *Options) GetLeaderRotation() Options_Rotation {
	if m != nil {
		return m.LeaderRotation
	}
	return Options_UNDEFINED
}

func (m *Options) GetDecisionsPerLeader() uint64 {
	if m != nil {
		return m.DecisionsPerLeader
	}
	return 0
}

func (m *Options) GetRequestMaxBytes() uint64 {
	if m != nil {
		return m.RequestMaxBytes
	}
	return 0
}

func (m *Options) GetRequestPoolSubmitTimeout() string {
	if m != nil {
		return m.RequestPoolSubmitTimeout
	}
	return ""
}

func init() {
	proto.RegisterEnum("smartbft.Options_Rotation", Options_Rotation_name, Options_Rotation_value)
	proto.RegisterType((*ConfigMetadata)(nil), "smartbft.ConfigMetadata")
	proto.RegisterType((*Consenter)(nil), "smartbft.Consenter")
	proto.RegisterType((*Options)(nil), "smartbft.Options")
}

func init() {
	proto.RegisterFile("orderer/smartbft/configuration.proto", fileDescriptor_a8a81ac5a2771ff3)
}

var fileDescriptor_a8a81ac5a2771ff3 = []byte{
	// 795 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x95, 0x6d, 0x8f, 0xdb, 0x44,
	0x10, 0xc7, 0xf1, 0x5d, 0x9a, 0x4b, 0x36, 0x97, 0xa7, 0xbd, 0xb4, 0x35, 0x07, 0x2f, 0x42, 0x84,
	0x20, 0x2a, 0x92, 0x83, 0xae, 0x14, 0x21, 0x21, 0x84, 0x48, 0xda, 0x88, 0x20, 0xee, 0x52, 0xb9,
	0x2d, 0x2f, 0x78, 0xb3, 0x5a, 0xdb, 0x13, 0x67, 0x25, 0xdb, 0x6b, 0x76, 0xd7, 0x69, 0xd3, 0xaf,
	0xc8, 0x27, 0xe0, 0xdb, 0x20, 0xef, 0xfa, 0xe9, 0x2e, 0x7d, 0x67, 0xcf, 0xff, 0xff, 0x9b, 0xf5,
	0x8c, 0x66, 0xc7, 0xe8, 0x6b, 0x2e, 0x02, 0x10, 0x20, 0x16, 0x32, 0xa6, 0x42, 0x79, 0x3b, 0xb5,
	0xf0, 0x79, 0xb2, 0x63, 0x61, 0x26, 0xa8, 0x62, 0x3c, 0x71, 0x52, 0xc1, 0x15, 0xc7, 0x9d, 0x52,
	0x9d, 0x09, 0x34, 0x58, 0x69, 0xc3, 0x2d, 0x28, 0x1a, 0x50, 0x45, 0xf1, 0x73, 0x84, 0x7c, 0x9e,
	0x48, 0x48, 0x14, 0x08, 0x69, 0x5b, 0xd3, 0xf3, 0x79, 0xef, 0xe6, 0xca, 0x29, 0x01, 0x67, 0x55,
	0x6a, 0x6e, 0xc3, 0x86, 0xbf, 0x43, 0x17, 0x3c, 0xcd, 0x0f, 0x90, 0xf6, 0xd9, 0xd4, 0x9a, 0xf7,
	0x6e, 0xc6, 0x35, 0xb1, 0x35, 0x82, 0x5b, 0x3a, 0x66, 0xff, 0x59, 0xa8, 0x5b, 0xa5, 0xc1, 0x5f,
	0xa1, 0xcb, 0x2a, 0x11, 0x61, 0x81, 0x6d, 0x4d, 0xad, 0x79, 0xcb, 0xed, 0x55, 0xb1, 0x4d, 0x80,
	0x31, 0x6a, 0xed, 0xb9, 0x54, 0x3a, 0x75, 0xd7, 0xd5, 0xcf, 0x79, 0x2c, 0xe5, 0x42, 0xd9, 0xe7,
	0x53, 0x6b, 0xde, 0x77, 0xf5, 0x33, 0x7e, 0x8c, 0xda, 0xb1, 0x4c, 0xf3, 0x24, 0x2d, 0xed, 0x7c,
	0x14, 0xcb, 0x74, 0x13, 0xe0, 0x6b, 0xd4, 0x61, 0x01, 0x24, 0x8a, 0xa9, 0xa3, 0xfd, 0x68, 0x6a,
	0xcd, 0x2f, 0xdd, 0xea, 0x1d, 0x7f, 0x83, 0x86, 0x7e, 0xc4, 0x20, 0x51, 0x44, 0x45, 0x92, 0xf8,
	0x20, 0x94, 0xdd, 0xd6, 0x96, 0xbe, 0x09, 0xbf, 0x8d, 0xe4, 0x0a, 0x84, 0xca, 0x7d, 0x12, 0xc4,
	0x01, 0x44, 0xed, 0xbb, 0x30, 0x3e, 0x13, 0x2e, 0x7c, 0xb3, 0x7f, 0x3b, 0xe8, 0xa2, 0x28, 0x18,
	0xbf, 0x40, 0x4f, 0x05, 0xfc, 0x93, 0x81, 0x54, 0xc4, 0xa3, 0xca, 0xdf, 0x93, 0x98, 0x7e, 0x20,
	0x3e, 0xcf, 0x12, 0x53, 0x49, 0xcb, 0x9d, 0x14, 0xf2, 0x32, 0x57, 0x6f, 0xe9, 0x87, 0x55, 0xae,
	0x7d, 0x1a, 0xf3, 0x8e, 0x0a, 0xa4, 0x2e, 0xf6, 0x14, 0x5b, 0xe6, 0x1a, 0xfe, 0x19, 0x5d, 0x9f,
	0x62, 0x2c, 0xef, 0xe0, 0x81, 0x46, 0x45, 0x43, 0x9e, 0x3e, 0x20, 0x37, 0x85, 0x8c, 0x7f, 0x45,
	0x5f, 0xb2, 0xc4, 0xe7, 0x31, 0x4b, 0x42, 0x12, 0x83, 0x94, 0x34, 0x04, 0xe2, 0x65, 0xbb, 0x1d,
	0x08, 0x22, 0xd9, 0x47, 0xd0, 0x6d, 0x6b, 0xb9, 0x9f, 0x97, 0x9e, 0x5b, 0x63, 0x59, 0x6a, 0xc7,
	0x1b, 0xf6, 0x11, 0xf0, 0x33, 0x34, 0x2e, 0x4f, 0x4f, 0x39, 0x8f, 0x0c, 0xd5, 0xd6, 0xd4, 0xb0,
	0x10, 0x5e, 0x73, 0x1e, 0x69, 0xef, 0x8f, 0x75, 0x81, 0x3b, 0x2e, 0xde, 0x53, 0x11, 0x10, 0xc5,
	0x62, 0xe0, 0x99, 0xe9, 0x69, 0xd7, 0x7d, 0x5c, 0xc8, 0x6b, 0xa3, 0xbe, 0x35, 0x22, 0xfe, 0x09,
	0xd9, 0x25, 0xe7, 0xf3, 0x38, 0x8d, 0x28, 0x4b, 0x2a, 0xb0, 0xa3, 0xc1, 0x27, 0x85, 0xbe, 0x2a,
	0xe4, 0x92, 0xfc, 0x05, 0x7d, 0x51, 0x92, 0x34, 0x53, 0x9c, 0x08, 0x88, 0xf9, 0x01, 0x2a, 0xb8,
	0xab, 0xe1, 0x32, 0xf9, 0x6f, 0x99, 0xe2, 0xae, 0x36, 0x34, 0xf0, 0x03, 0x83, 0xf7, 0xc4, 0xdf,
	0xd3, 0x24, 0x04, 0x22, 0x40, 0x42, 0x12, 0xd4, 0xbd, 0x45, 0x06, 0xcf, 0x2d, 0x2b, 0xed, 0x70,
	0xb5, 0xa1, 0x6a, 0xae, 0x83, 0xae, 0x9a, 0x78, 0x79, 0x6a, 0x4f, 0x63, 0xe3, 0x1a, 0x6b, 0xd4,
	0x19, 0x01, 0x0d, 0x40, 0x90, 0x3d, 0xe4, 0x77, 0x08, 0xa8, 0xaa, 0xa0, 0x4b, 0x53, 0xa7, 0xd1,
	0x7f, 0x2f, 0xe5, 0x92, 0xfc, 0x01, 0x3d, 0x39, 0x21, 0xcd, 0xc0, 0xf5, 0xcd, 0xe4, 0x3c, 0xe0,
	0xcc, 0xc0, 0x7d, 0x8b, 0x86, 0x3e, 0x8f, 0x22, 0xf0, 0xeb, 0x63, 0x06, 0xfa, 0x98, 0x41, 0x11,
	0x2e, 0xd3, 0xcf, 0x50, 0x5f, 0x1e, 0x13, 0x9f, 0xf0, 0x84, 0x48, 0x45, 0x85, 0xb2, 0x87, 0x53,
	0x6b, 0xde, 0x71, 0x7b, 0x79, 0x70, 0x9b, 0xbc, 0xc9, 0x43, 0x78, 0x81, 0x26, 0x32, 0x05, 0x08,
	0x48, 0x96, 0x92, 0x46, 0xd5, 0xf6, 0x48, 0x5b, 0xc7, 0x5a, 0x7b, 0x97, 0xfe, 0x55, 0x15, 0x8d,
	0x57, 0x68, 0x58, 0x7c, 0xb3, 0xe0, 0x4a, 0x2f, 0x29, 0x7b, 0x3c, 0xb5, 0xe6, 0x83, 0x9b, 0xeb,
	0x93, 0x15, 0xe2, 0xb8, 0x85, 0xc3, 0x1d, 0x18, 0xa4, 0x7c, 0xc7, 0xdf, 0xa3, 0x49, 0x00, 0x3e,
	0x93, 0xb9, 0x8b, 0xa4, 0x20, 0x88, 0xd1, 0x6d, 0xac, 0xcb, 0xc6, 0x95, 0xf6, 0x1a, 0xc4, 0x9f,
	0x5a, 0x69, 0x0e, 0x6c, 0x7d, 0xbf, 0xae, 0xee, 0x0d, 0x6c, 0x75, 0xb5, 0x1a, 0xe3, 0x63, 0x86,
	0x3b, 0xf3, 0x62, 0x56, 0x37, 0x6b, 0x72, 0x6f, 0x7c, 0xf4, 0x98, 0x6b, 0x43, 0xd1, 0xb6, 0xd9,
	0x33, 0xd4, 0xa9, 0x3e, 0xb4, 0x8f, 0xba, 0xef, 0xee, 0x5e, 0xbe, 0x5a, 0x6f, 0xee, 0x5e, 0xbd,
	0x1c, 0x7d, 0x86, 0x2f, 0xd0, 0xf9, 0x76, 0xbd, 0x1e, 0x59, 0xb8, 0x8d, 0xce, 0xb6, 0x77, 0xa3,
	0xb3, 0x3f, 0x5a, 0x1d, 0x6b, 0x74, 0xe6, 0xb6, 0xcd, 0xd2, 0x5e, 0x86, 0xc8, 0xe1, 0x22, 0x74,
	0xf6, 0xc7, 0x14, 0x44, 0x04, 0x41, 0x08, 0xc2, 0xd9, 0x51, 0x4f, 0x30, 0xdf, 0xec, 0x71, 0xe9,
	0x14, 0xdb, 0xbe, 0xea, 0xd4, 0xdf, 0x2f, 0x42, 0xa6, 0xf6, 0x99, 0xe7, 0xf8, 0x3c, 0x5e, 0x34,
	0xb0, 0x85, 0xc1, 0x16, 0x06, 0x5b, 0x3c, 0xfc, 0x49, 0x78, 0x6d, 0x2d, 0x3c, 0xff, 0x3f, 0x00,
	0x00, 0xff, 0xff, 0x01, 0x0e, 0xd6, 0x36, 0x3f, 0x06, 0x00, 0x00,
}

// Copyright the Hyperledger Fabric contributors. All rights reserved.
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: peer/snapshot.proto

package peer

import (
	common "github.com/hyperledger/fabric-protos-go/common"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// SnapshotRequest contains information for a generate/cancel snapshot request
type SnapshotRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The signature header that contains creator identity and nonce
	SignatureHeader *common.SignatureHeader `protobuf:"bytes,1,opt,name=signature_header,json=signatureHeader,proto3" json:"signature_header,omitempty"`
	// The channel ID
	ChannelId string `protobuf:"bytes,2,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	// The block number to generate a snapshot
	BlockNumber uint64 `protobuf:"varint,3,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
}

func (x *SnapshotRequest) Reset() {
	*x = SnapshotRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_snapshot_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapshotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapshotRequest) ProtoMessage() {}

func (x *SnapshotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_peer_snapshot_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapshotRequest.ProtoReflect.Descriptor instead.
func (*SnapshotRequest) Descriptor() ([]byte, []int) {
	return file_peer_snapshot_proto_rawDescGZIP(), []int{0}
}

func (x *SnapshotRequest) GetSignatureHeader() *common.SignatureHeader {
	if x != nil {
		return x.SignatureHeader
	}
	return nil
}

func (x *SnapshotRequest) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *SnapshotRequest) GetBlockNumber() uint64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

// SnapshotQuery contains information for a query snapshot request
type SnapshotQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The signature header that contains creator identity and nonce
	SignatureHeader *common.SignatureHeader `protobuf:"bytes,1,opt,name=signature_header,json=signatureHeader,proto3" json:"signature_header,omitempty"`
	// The channel ID
	ChannelId string `protobuf:"bytes,2,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
}

func (x *SnapshotQuery) Reset() {
	*x = SnapshotQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_snapshot_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapshotQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapshotQuery) ProtoMessage() {}

func (x *SnapshotQuery) ProtoReflect() protoreflect.Message {
	mi := &file_peer_snapshot_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapshotQuery.ProtoReflect.Descriptor instead.
func (*SnapshotQuery) Descriptor() ([]byte, []int) {
	return file_peer_snapshot_proto_rawDescGZIP(), []int{1}
}

func (x *SnapshotQuery) GetSignatureHeader() *common.SignatureHeader {
	if x != nil {
		return x.SignatureHeader
	}
	return nil
}

func (x *SnapshotQuery) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

// SignedSnapshotRequest contains marshalled request bytes and signature
type SignedSnapshotRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The bytes of SnapshotRequest or SnapshotQuery
	Request []byte `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	// Signaure over request bytes; this signature is to be verified against the client identity
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *SignedSnapshotRequest) Reset() {
	*x = SignedSnapshotRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_snapshot_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignedSnapshotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignedSnapshotRequest) ProtoMessage() {}

func (x *SignedSnapshotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_peer_snapshot_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignedSnapshotRequest.ProtoReflect.Descriptor instead.
func (*SignedSnapshotRequest) Descriptor() ([]byte, []int) {
	return file_peer_snapshot_proto_rawDescGZIP(), []int{2}
}

func (x *SignedSnapshotRequest) GetRequest() []byte {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *SignedSnapshotRequest) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

// QueryPendingSnapshotsResponse specifies the response payload of a query pending snapshots request
type QueryPendingSnapshotsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockNumbers []uint64 `protobuf:"varint,1,rep,packed,name=block_numbers,json=blockNumbers,proto3" json:"block_numbers,omitempty"`
}

func (x *QueryPendingSnapshotsResponse) Reset() {
	*x = QueryPendingSnapshotsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_snapshot_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryPendingSnapshotsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryPendingSnapshotsResponse) ProtoMessage() {}

func (x *QueryPendingSnapshotsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_peer_snapshot_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryPendingSnapshotsResponse.ProtoReflect.Descriptor instead.
func (*QueryPendingSnapshotsResponse) Descriptor() ([]byte, []int) {
	return file_peer_snapshot_proto_rawDescGZIP(), []int{3}
}

func (x *QueryPendingSnapshotsResponse) GetBlockNumbers() []uint64 {
	if x != nil {
		return x.BlockNumbers
	}
	return nil
}

var File_peer_snapshot_proto protoreflect.FileDescriptor

var file_peer_snapshot_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x65, 0x65, 0x72, 0x2f, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x97, 0x01, 0x0a, 0x0f, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x42, 0x0a, 0x10, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x0f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x72, 0x0a, 0x0d, 0x53, 0x6e, 0x61,
	0x70, 0x73, 0x68, 0x6f, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x42, 0x0a, 0x10, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x0f, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x22, 0x4f, 0x0a,
	0x15, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x44,
	0x0a, 0x1d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x6e,
	0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x23, 0x0a, 0x0d, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x04, 0x52, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x32, 0xeb, 0x01, 0x0a, 0x08, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f,
	0x74, 0x12, 0x43, 0x0a, 0x08, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x53, 0x6e, 0x61,
	0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x06, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c,
	0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64,
	0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x0d, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68,
	0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x53,
	0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x52, 0x0a, 0x22, 0x6f, 0x72, 0x67, 0x2e, 0x68, 0x79, 0x70, 0x65, 0x72, 0x6c,
	0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x70, 0x65, 0x65, 0x72, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x79, 0x70, 0x65, 0x72, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72,
	0x2f, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2d, 0x67,
	0x6f, 0x2f, 0x70, 0x65, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_peer_snapshot_proto_rawDescOnce sync.Once
	file_peer_snapshot_proto_rawDescData = file_peer_snapshot_proto_rawDesc
)

func file_peer_snapshot_proto_rawDescGZIP() []byte {
	file_peer_snapshot_proto_rawDescOnce.Do(func() {
		file_peer_snapshot_proto_rawDescData = protoimpl.X.CompressGZIP(file_peer_snapshot_proto_rawDescData)
	})
	return file_peer_snapshot_proto_rawDescData
}

var file_peer_snapshot_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_peer_snapshot_proto_goTypes = []interface{}{
	(*SnapshotRequest)(nil),               // 0: protos.SnapshotRequest
	(*SnapshotQuery)(nil),                 // 1: protos.SnapshotQuery
	(*SignedSnapshotRequest)(nil),         // 2: protos.SignedSnapshotRequest
	(*QueryPendingSnapshotsResponse)(nil), // 3: protos.QueryPendingSnapshotsResponse
	(*common.SignatureHeader)(nil),        // 4: common.SignatureHeader
	(*emptypb.Empty)(nil),                 // 5: google.protobuf.Empty
}
var file_peer_snapshot_proto_depIdxs = []int32{
	4, // 0: protos.SnapshotRequest.signature_header:type_name -> common.SignatureHeader
	4, // 1: protos.SnapshotQuery.signature_header:type_name -> common.SignatureHeader
	2, // 2: protos.Snapshot.Generate:input_type -> protos.SignedSnapshotRequest
	2, // 3: protos.Snapshot.Cancel:input_type -> protos.SignedSnapshotRequest
	2, // 4: protos.Snapshot.QueryPendings:input_type -> protos.SignedSnapshotRequest
	5, // 5: protos.Snapshot.Generate:output_type -> google.protobuf.Empty
	5, // 6: protos.Snapshot.Cancel:output_type -> google.protobuf.Empty
	3, // 7: protos.Snapshot.QueryPendings:output_type -> protos.QueryPendingSnapshotsResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_peer_snapshot_proto_init() }
func file_peer_snapshot_proto_init() {
	if File_peer_snapshot_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_peer_snapshot_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnapshotRequest); i {
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
		file_peer_snapshot_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnapshotQuery); i {
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
		file_peer_snapshot_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignedSnapshotRequest); i {
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
		file_peer_snapshot_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryPendingSnapshotsResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_peer_snapshot_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_peer_snapshot_proto_goTypes,
		DependencyIndexes: file_peer_snapshot_proto_depIdxs,
		MessageInfos:      file_peer_snapshot_proto_msgTypes,
	}.Build()
	File_peer_snapshot_proto = out.File
	file_peer_snapshot_proto_rawDesc = nil
	file_peer_snapshot_proto_goTypes = nil
	file_peer_snapshot_proto_depIdxs = nil
}

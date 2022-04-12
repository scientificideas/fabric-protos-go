// Copyright the Hyperledger Fabric contributors. All rights reserved.
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: peer/proposal_response.proto

package peer

import (
	common "github.com/hyperledger/fabric-protos-go/common"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A ProposalResponse is returned from an endorser to the proposal submitter.
// The idea is that this message contains the endorser's response to the
// request of a client to perform an action over a chaincode (or more
// generically on the ledger); the response might be success/error (conveyed in
// the Response field) together with a description of the action and a
// signature over it by that endorser.  If a sufficient number of distinct
// endorsers agree on the same action and produce signature to that effect, a
// transaction can be generated and sent for ordering.
type ProposalResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Version indicates message protocol version
	Version int32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	// Timestamp is the time that the message
	// was created as  defined by the sender
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// A response message indicating whether the
	// endorsement of the action was successful
	Response *Response `protobuf:"bytes,4,opt,name=response,proto3" json:"response,omitempty"`
	// The payload of response. It is the bytes of ProposalResponsePayload
	Payload []byte `protobuf:"bytes,5,opt,name=payload,proto3" json:"payload,omitempty"`
	// The endorsement of the proposal, basically
	// the endorser's signature over the payload
	Endorsement *Endorsement `protobuf:"bytes,6,opt,name=endorsement,proto3" json:"endorsement,omitempty"`
	// The chaincode interest derived from simulating the proposal.
	Interest *ChaincodeInterest `protobuf:"bytes,7,opt,name=interest,proto3" json:"interest,omitempty"`
}

func (x *ProposalResponse) Reset() {
	*x = ProposalResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_proposal_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProposalResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProposalResponse) ProtoMessage() {}

func (x *ProposalResponse) ProtoReflect() protoreflect.Message {
	mi := &file_peer_proposal_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProposalResponse.ProtoReflect.Descriptor instead.
func (*ProposalResponse) Descriptor() ([]byte, []int) {
	return file_peer_proposal_response_proto_rawDescGZIP(), []int{0}
}

func (x *ProposalResponse) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *ProposalResponse) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *ProposalResponse) GetResponse() *Response {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *ProposalResponse) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *ProposalResponse) GetEndorsement() *Endorsement {
	if x != nil {
		return x.Endorsement
	}
	return nil
}

func (x *ProposalResponse) GetInterest() *ChaincodeInterest {
	if x != nil {
		return x.Interest
	}
	return nil
}

// A response with a representation similar to an HTTP response that can
// be used within another message.
type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A status code that should follow the HTTP status codes.
	Status int32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	// A message associated with the response code.
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// A payload that can be used to include metadata with this response.
	Payload []byte `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_proposal_response_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_peer_proposal_response_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_peer_proposal_response_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Response) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

// ProposalResponsePayload is the payload of a proposal response.  This message
// is the "bridge" between the client's request and the endorser's action in
// response to that request. Concretely, for chaincodes, it contains a hashed
// representation of the proposal (proposalHash) and a representation of the
// chaincode state changes and events inside the extension field.
type ProposalResponsePayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Hash of the proposal that triggered this response. The hash is used to
	// link a response with its proposal, both for bookeeping purposes on an
	// asynchronous system and for security reasons (accountability,
	// non-repudiation). The hash usually covers the entire Proposal message
	// (byte-by-byte).
	ProposalHash []byte `protobuf:"bytes,1,opt,name=proposal_hash,json=proposalHash,proto3" json:"proposal_hash,omitempty"`
	// Extension should be unmarshaled to a type-specific message. The type of
	// the extension in any proposal response depends on the type of the proposal
	// that the client selected when the proposal was initially sent out.  In
	// particular, this information is stored in the type field of a Header.  For
	// chaincode, it's a ChaincodeAction message
	Extension []byte `protobuf:"bytes,2,opt,name=extension,proto3" json:"extension,omitempty"`
}

func (x *ProposalResponsePayload) Reset() {
	*x = ProposalResponsePayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_proposal_response_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProposalResponsePayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProposalResponsePayload) ProtoMessage() {}

func (x *ProposalResponsePayload) ProtoReflect() protoreflect.Message {
	mi := &file_peer_proposal_response_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProposalResponsePayload.ProtoReflect.Descriptor instead.
func (*ProposalResponsePayload) Descriptor() ([]byte, []int) {
	return file_peer_proposal_response_proto_rawDescGZIP(), []int{2}
}

func (x *ProposalResponsePayload) GetProposalHash() []byte {
	if x != nil {
		return x.ProposalHash
	}
	return nil
}

func (x *ProposalResponsePayload) GetExtension() []byte {
	if x != nil {
		return x.Extension
	}
	return nil
}

// An endorsement is a signature of an endorser over a proposal response.  By
// producing an endorsement message, an endorser implicitly "approves" that
// proposal response and the actions contained therein. When enough
// endorsements have been collected, a transaction can be generated out of a
// set of proposal responses.  Note that this message only contains an identity
// and a signature but no signed payload. This is intentional because
// endorsements are supposed to be collected in a transaction, and they are all
// expected to endorse a single proposal response/action (many endorsements
// over a single proposal response)
type Endorsement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identity of the endorser (e.g. its certificate)
	Endorser []byte `protobuf:"bytes,1,opt,name=endorser,proto3" json:"endorser,omitempty"`
	// Signature of the payload included in ProposalResponse concatenated with
	// the endorser's certificate; ie, sign(ProposalResponse.payload + endorser)
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *Endorsement) Reset() {
	*x = Endorsement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_proposal_response_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Endorsement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Endorsement) ProtoMessage() {}

func (x *Endorsement) ProtoReflect() protoreflect.Message {
	mi := &file_peer_proposal_response_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Endorsement.ProtoReflect.Descriptor instead.
func (*Endorsement) Descriptor() ([]byte, []int) {
	return file_peer_proposal_response_proto_rawDescGZIP(), []int{3}
}

func (x *Endorsement) GetEndorser() []byte {
	if x != nil {
		return x.Endorser
	}
	return nil
}

func (x *Endorsement) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

// ChaincodeInterest defines an interest about an endorsement
// for a specific single chaincode invocation.
// Multiple chaincodes indicate chaincode to chaincode invocations.
type ChaincodeInterest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chaincodes []*ChaincodeCall `protobuf:"bytes,1,rep,name=chaincodes,proto3" json:"chaincodes,omitempty"`
}

func (x *ChaincodeInterest) Reset() {
	*x = ChaincodeInterest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_proposal_response_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChaincodeInterest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChaincodeInterest) ProtoMessage() {}

func (x *ChaincodeInterest) ProtoReflect() protoreflect.Message {
	mi := &file_peer_proposal_response_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChaincodeInterest.ProtoReflect.Descriptor instead.
func (*ChaincodeInterest) Descriptor() ([]byte, []int) {
	return file_peer_proposal_response_proto_rawDescGZIP(), []int{4}
}

func (x *ChaincodeInterest) GetChaincodes() []*ChaincodeCall {
	if x != nil {
		return x.Chaincodes
	}
	return nil
}

// ChaincodeCall defines a call to a chaincode.
// It may have collections that are related to the chaincode
type ChaincodeCall struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name            string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CollectionNames []string `protobuf:"bytes,2,rep,name=collection_names,json=collectionNames,proto3" json:"collection_names,omitempty"`
	NoPrivateReads  bool     `protobuf:"varint,3,opt,name=no_private_reads,json=noPrivateReads,proto3" json:"no_private_reads,omitempty"` // Indicates we do not need to read from private data
	NoPublicWrites  bool     `protobuf:"varint,4,opt,name=no_public_writes,json=noPublicWrites,proto3" json:"no_public_writes,omitempty"` // Indicates we do not need to write to the chaincode namespace
	// The set of signature policies associated with states in the write-set
	// that have state-based endorsement policies.
	KeyPolicies []*common.SignaturePolicyEnvelope `protobuf:"bytes,5,rep,name=key_policies,json=keyPolicies,proto3" json:"key_policies,omitempty"`
	// Indicates we wish to ignore the namespace endorsement policy
	DisregardNamespacePolicy bool `protobuf:"varint,6,opt,name=disregard_namespace_policy,json=disregardNamespacePolicy,proto3" json:"disregard_namespace_policy,omitempty"`
}

func (x *ChaincodeCall) Reset() {
	*x = ChaincodeCall{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_proposal_response_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChaincodeCall) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChaincodeCall) ProtoMessage() {}

func (x *ChaincodeCall) ProtoReflect() protoreflect.Message {
	mi := &file_peer_proposal_response_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChaincodeCall.ProtoReflect.Descriptor instead.
func (*ChaincodeCall) Descriptor() ([]byte, []int) {
	return file_peer_proposal_response_proto_rawDescGZIP(), []int{5}
}

func (x *ChaincodeCall) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ChaincodeCall) GetCollectionNames() []string {
	if x != nil {
		return x.CollectionNames
	}
	return nil
}

func (x *ChaincodeCall) GetNoPrivateReads() bool {
	if x != nil {
		return x.NoPrivateReads
	}
	return false
}

func (x *ChaincodeCall) GetNoPublicWrites() bool {
	if x != nil {
		return x.NoPublicWrites
	}
	return false
}

func (x *ChaincodeCall) GetKeyPolicies() []*common.SignaturePolicyEnvelope {
	if x != nil {
		return x.KeyPolicies
	}
	return nil
}

func (x *ChaincodeCall) GetDisregardNamespacePolicy() bool {
	if x != nil {
		return x.DisregardNamespacePolicy
	}
	return false
}

var File_peer_proposal_response_proto protoreflect.FileDescriptor

var file_peer_proposal_response_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x65, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x5f,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c,
	0x02, 0x0a, 0x10, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x2c, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12,
	0x35, 0x0a, 0x0b, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x45, 0x6e,
	0x64, 0x6f, 0x72, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x65, 0x6e, 0x64, 0x6f, 0x72,
	0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x35, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x65,
	0x73, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x65, 0x73, 0x74, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x22, 0x56, 0x0a,
	0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x5c, 0x0a, 0x17, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x5f, 0x68, 0x61, 0x73,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61,
	0x6c, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x22, 0x47, 0x0a, 0x0b, 0x45, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x65, 0x72, 0x12, 0x1c,
	0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x4a, 0x0a, 0x11,
	0x43, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73,
	0x74, 0x12, 0x35, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43,
	0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x0a, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x22, 0xa4, 0x02, 0x0a, 0x0d, 0x43, 0x68, 0x61,
	0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x29,
	0x0a, 0x10, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x10, 0x6e, 0x6f, 0x5f,
	0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0e, 0x6e, 0x6f, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x61, 0x64, 0x73, 0x12, 0x28, 0x0a, 0x10, 0x6e, 0x6f, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x5f, 0x77, 0x72, 0x69, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x6e,
	0x6f, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x57, 0x72, 0x69, 0x74, 0x65, 0x73, 0x12, 0x42, 0x0a,
	0x0c, 0x6b, 0x65, 0x79, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x45, 0x6e, 0x76, 0x65,
	0x6c, 0x6f, 0x70, 0x65, 0x52, 0x0b, 0x6b, 0x65, 0x79, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65,
	0x73, 0x12, 0x3c, 0x0a, 0x1a, 0x64, 0x69, 0x73, 0x72, 0x65, 0x67, 0x61, 0x72, 0x64, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x18, 0x64, 0x69, 0x73, 0x72, 0x65, 0x67, 0x61, 0x72, 0x64,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x42,
	0x6b, 0x0a, 0x22, 0x6f, 0x72, 0x67, 0x2e, 0x68, 0x79, 0x70, 0x65, 0x72, 0x6c, 0x65, 0x64, 0x67,
	0x65, 0x72, 0x2e, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x70, 0x65, 0x65, 0x72, 0x42, 0x17, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5a, 0x2c,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x79, 0x70, 0x65, 0x72,
	0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x65, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_peer_proposal_response_proto_rawDescOnce sync.Once
	file_peer_proposal_response_proto_rawDescData = file_peer_proposal_response_proto_rawDesc
)

func file_peer_proposal_response_proto_rawDescGZIP() []byte {
	file_peer_proposal_response_proto_rawDescOnce.Do(func() {
		file_peer_proposal_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_peer_proposal_response_proto_rawDescData)
	})
	return file_peer_proposal_response_proto_rawDescData
}

var file_peer_proposal_response_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_peer_proposal_response_proto_goTypes = []interface{}{
	(*ProposalResponse)(nil),               // 0: protos.ProposalResponse
	(*Response)(nil),                       // 1: protos.Response
	(*ProposalResponsePayload)(nil),        // 2: protos.ProposalResponsePayload
	(*Endorsement)(nil),                    // 3: protos.Endorsement
	(*ChaincodeInterest)(nil),              // 4: protos.ChaincodeInterest
	(*ChaincodeCall)(nil),                  // 5: protos.ChaincodeCall
	(*timestamppb.Timestamp)(nil),          // 6: google.protobuf.Timestamp
	(*common.SignaturePolicyEnvelope)(nil), // 7: common.SignaturePolicyEnvelope
}
var file_peer_proposal_response_proto_depIdxs = []int32{
	6, // 0: protos.ProposalResponse.timestamp:type_name -> google.protobuf.Timestamp
	1, // 1: protos.ProposalResponse.response:type_name -> protos.Response
	3, // 2: protos.ProposalResponse.endorsement:type_name -> protos.Endorsement
	4, // 3: protos.ProposalResponse.interest:type_name -> protos.ChaincodeInterest
	5, // 4: protos.ChaincodeInterest.chaincodes:type_name -> protos.ChaincodeCall
	7, // 5: protos.ChaincodeCall.key_policies:type_name -> common.SignaturePolicyEnvelope
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_peer_proposal_response_proto_init() }
func file_peer_proposal_response_proto_init() {
	if File_peer_proposal_response_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_peer_proposal_response_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProposalResponse); i {
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
		file_peer_proposal_response_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_peer_proposal_response_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProposalResponsePayload); i {
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
		file_peer_proposal_response_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Endorsement); i {
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
		file_peer_proposal_response_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChaincodeInterest); i {
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
		file_peer_proposal_response_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChaincodeCall); i {
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
			RawDescriptor: file_peer_proposal_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_peer_proposal_response_proto_goTypes,
		DependencyIndexes: file_peer_proposal_response_proto_depIdxs,
		MessageInfos:      file_peer_proposal_response_proto_msgTypes,
	}.Build()
	File_peer_proposal_response_proto = out.File
	file_peer_proposal_response_proto_rawDesc = nil
	file_peer_proposal_response_proto_goTypes = nil
	file_peer_proposal_response_proto_depIdxs = nil
}

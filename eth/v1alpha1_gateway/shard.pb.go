// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.4
// source: eth/v1alpha1/shard.proto

package eth

import (
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ShardBlock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShardParentRoot  []byte `protobuf:"bytes,1,opt,name=shard_parent_root,json=shardParentRoot,proto3" json:"shard_parent_root,omitempty"`
	BeaconParentRoot []byte `protobuf:"bytes,2,opt,name=beacon_parent_root,json=beaconParentRoot,proto3" json:"beacon_parent_root,omitempty"`
	Slot             uint64 `protobuf:"varint,3,opt,name=slot,proto3" json:"slot,omitempty"`
	Shard            uint64 `protobuf:"varint,4,opt,name=shard,proto3" json:"shard,omitempty"`
	ProposerIndex    uint64 `protobuf:"varint,5,opt,name=proposer_index,json=proposerIndex,proto3" json:"proposer_index,omitempty"`
	Body             []byte `protobuf:"bytes,6,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *ShardBlock) Reset() {
	*x = ShardBlock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eth_v1alpha1_shard_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShardBlock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShardBlock) ProtoMessage() {}

func (x *ShardBlock) ProtoReflect() protoreflect.Message {
	mi := &file_eth_v1alpha1_shard_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShardBlock.ProtoReflect.Descriptor instead.
func (*ShardBlock) Descriptor() ([]byte, []int) {
	return file_eth_v1alpha1_shard_proto_rawDescGZIP(), []int{0}
}

func (x *ShardBlock) GetShardParentRoot() []byte {
	if x != nil {
		return x.ShardParentRoot
	}
	return nil
}

func (x *ShardBlock) GetBeaconParentRoot() []byte {
	if x != nil {
		return x.BeaconParentRoot
	}
	return nil
}

func (x *ShardBlock) GetSlot() uint64 {
	if x != nil {
		return x.Slot
	}
	return 0
}

func (x *ShardBlock) GetShard() uint64 {
	if x != nil {
		return x.Shard
	}
	return 0
}

func (x *ShardBlock) GetProposerIndex() uint64 {
	if x != nil {
		return x.ProposerIndex
	}
	return 0
}

func (x *ShardBlock) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

type SignedShardBlock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message   *ShardBlock `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Signature []byte      `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *SignedShardBlock) Reset() {
	*x = SignedShardBlock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eth_v1alpha1_shard_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignedShardBlock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignedShardBlock) ProtoMessage() {}

func (x *SignedShardBlock) ProtoReflect() protoreflect.Message {
	mi := &file_eth_v1alpha1_shard_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignedShardBlock.ProtoReflect.Descriptor instead.
func (*SignedShardBlock) Descriptor() ([]byte, []int) {
	return file_eth_v1alpha1_shard_proto_rawDescGZIP(), []int{1}
}

func (x *SignedShardBlock) GetMessage() *ShardBlock {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *SignedShardBlock) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type ShardBlockHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShardParentRoot  []byte `protobuf:"bytes,1,opt,name=shard_parent_root,json=shardParentRoot,proto3" json:"shard_parent_root,omitempty"`
	BeaconParentRoot []byte `protobuf:"bytes,2,opt,name=beacon_parent_root,json=beaconParentRoot,proto3" json:"beacon_parent_root,omitempty"`
	Slot             uint64 `protobuf:"varint,3,opt,name=slot,proto3" json:"slot,omitempty"`
	Shard            uint64 `protobuf:"varint,4,opt,name=shard,proto3" json:"shard,omitempty"`
	ProposerIndex    uint64 `protobuf:"varint,5,opt,name=proposer_index,json=proposerIndex,proto3" json:"proposer_index,omitempty"`
	BodyRoot         []byte `protobuf:"bytes,6,opt,name=body_root,json=bodyRoot,proto3" json:"body_root,omitempty"`
}

func (x *ShardBlockHeader) Reset() {
	*x = ShardBlockHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eth_v1alpha1_shard_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShardBlockHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShardBlockHeader) ProtoMessage() {}

func (x *ShardBlockHeader) ProtoReflect() protoreflect.Message {
	mi := &file_eth_v1alpha1_shard_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShardBlockHeader.ProtoReflect.Descriptor instead.
func (*ShardBlockHeader) Descriptor() ([]byte, []int) {
	return file_eth_v1alpha1_shard_proto_rawDescGZIP(), []int{2}
}

func (x *ShardBlockHeader) GetShardParentRoot() []byte {
	if x != nil {
		return x.ShardParentRoot
	}
	return nil
}

func (x *ShardBlockHeader) GetBeaconParentRoot() []byte {
	if x != nil {
		return x.BeaconParentRoot
	}
	return nil
}

func (x *ShardBlockHeader) GetSlot() uint64 {
	if x != nil {
		return x.Slot
	}
	return 0
}

func (x *ShardBlockHeader) GetShard() uint64 {
	if x != nil {
		return x.Shard
	}
	return 0
}

func (x *ShardBlockHeader) GetProposerIndex() uint64 {
	if x != nil {
		return x.ProposerIndex
	}
	return 0
}

func (x *ShardBlockHeader) GetBodyRoot() []byte {
	if x != nil {
		return x.BodyRoot
	}
	return nil
}

type ShardState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slot            uint64 `protobuf:"varint,1,opt,name=slot,proto3" json:"slot,omitempty"`
	GasPrice        uint64 `protobuf:"varint,2,opt,name=gas_price,json=gasPrice,proto3" json:"gas_price,omitempty"`
	LatestBlockRoot []byte `protobuf:"bytes,3,opt,name=latest_block_root,json=latestBlockRoot,proto3" json:"latest_block_root,omitempty"`
}

func (x *ShardState) Reset() {
	*x = ShardState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eth_v1alpha1_shard_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShardState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShardState) ProtoMessage() {}

func (x *ShardState) ProtoReflect() protoreflect.Message {
	mi := &file_eth_v1alpha1_shard_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShardState.ProtoReflect.Descriptor instead.
func (*ShardState) Descriptor() ([]byte, []int) {
	return file_eth_v1alpha1_shard_proto_rawDescGZIP(), []int{3}
}

func (x *ShardState) GetSlot() uint64 {
	if x != nil {
		return x.Slot
	}
	return 0
}

func (x *ShardState) GetGasPrice() uint64 {
	if x != nil {
		return x.GasPrice
	}
	return 0
}

func (x *ShardState) GetLatestBlockRoot() []byte {
	if x != nil {
		return x.LatestBlockRoot
	}
	return nil
}

type ShardTransition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartSlot                  uint64        `protobuf:"varint,1,opt,name=start_slot,json=startSlot,proto3" json:"start_slot,omitempty"`
	ShardBlockLengths          []uint64      `protobuf:"varint,2,rep,packed,name=shard_block_lengths,json=shardBlockLengths,proto3" json:"shard_block_lengths,omitempty"`
	ShardDataRoots             [][]byte      `protobuf:"bytes,3,rep,name=shard_data_roots,json=shardDataRoots,proto3" json:"shard_data_roots,omitempty"`
	ShardStates                []*ShardState `protobuf:"bytes,4,rep,name=shard_states,json=shardStates,proto3" json:"shard_states,omitempty"`
	ProposerSignatureAggregate []byte        `protobuf:"bytes,5,opt,name=proposer_signature_aggregate,json=proposerSignatureAggregate,proto3" json:"proposer_signature_aggregate,omitempty"`
}

func (x *ShardTransition) Reset() {
	*x = ShardTransition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eth_v1alpha1_shard_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShardTransition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShardTransition) ProtoMessage() {}

func (x *ShardTransition) ProtoReflect() protoreflect.Message {
	mi := &file_eth_v1alpha1_shard_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShardTransition.ProtoReflect.Descriptor instead.
func (*ShardTransition) Descriptor() ([]byte, []int) {
	return file_eth_v1alpha1_shard_proto_rawDescGZIP(), []int{4}
}

func (x *ShardTransition) GetStartSlot() uint64 {
	if x != nil {
		return x.StartSlot
	}
	return 0
}

func (x *ShardTransition) GetShardBlockLengths() []uint64 {
	if x != nil {
		return x.ShardBlockLengths
	}
	return nil
}

func (x *ShardTransition) GetShardDataRoots() [][]byte {
	if x != nil {
		return x.ShardDataRoots
	}
	return nil
}

func (x *ShardTransition) GetShardStates() []*ShardState {
	if x != nil {
		return x.ShardStates
	}
	return nil
}

func (x *ShardTransition) GetProposerSignatureAggregate() []byte {
	if x != nil {
		return x.ProposerSignatureAggregate
	}
	return nil
}

type CompactCommittee struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PubKeys           [][]byte `protobuf:"bytes,1,rep,name=pub_keys,json=pubKeys,proto3" json:"pub_keys,omitempty"`
	CompactValidators []uint64 `protobuf:"varint,2,rep,packed,name=compact_validators,json=compactValidators,proto3" json:"compact_validators,omitempty"`
}

func (x *CompactCommittee) Reset() {
	*x = CompactCommittee{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eth_v1alpha1_shard_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompactCommittee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompactCommittee) ProtoMessage() {}

func (x *CompactCommittee) ProtoReflect() protoreflect.Message {
	mi := &file_eth_v1alpha1_shard_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompactCommittee.ProtoReflect.Descriptor instead.
func (*CompactCommittee) Descriptor() ([]byte, []int) {
	return file_eth_v1alpha1_shard_proto_rawDescGZIP(), []int{5}
}

func (x *CompactCommittee) GetPubKeys() [][]byte {
	if x != nil {
		return x.PubKeys
	}
	return nil
}

func (x *CompactCommittee) GetCompactValidators() []uint64 {
	if x != nil {
		return x.CompactValidators
	}
	return nil
}

var File_eth_v1alpha1_shard_proto protoreflect.FileDescriptor

var file_eth_v1alpha1_shard_proto_rawDesc = []byte{
	0x0a, 0x18, 0x65, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x73,
	0x68, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x65, 0x74, 0x68, 0x65,
	0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x1a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f,
	0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x6f, 0x67, 0x6f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x88, 0x02, 0x0a, 0x0a, 0x53, 0x68, 0x61, 0x72, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12,
	0x3d, 0x0a, 0x11, 0x73, 0x68, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f,
	0x72, 0x6f, 0x6f, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x11, 0xf2, 0xde, 0x1f, 0x0d,
	0x73, 0x73, 0x7a, 0x2d, 0x73, 0x69, 0x7a, 0x65, 0x3a, 0x22, 0x33, 0x32, 0x22, 0x52, 0x0f, 0x73,
	0x68, 0x61, 0x72, 0x64, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x3f,
	0x0a, 0x12, 0x62, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f,
	0x72, 0x6f, 0x6f, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x11, 0xf2, 0xde, 0x1f, 0x0d,
	0x73, 0x73, 0x7a, 0x2d, 0x73, 0x69, 0x7a, 0x65, 0x3a, 0x22, 0x33, 0x32, 0x22, 0x52, 0x10, 0x62,
	0x65, 0x61, 0x63, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x6f, 0x6f, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73,
	0x6c, 0x6f, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x68, 0x61, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x05, 0x73, 0x68, 0x61, 0x72, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x72, 0x6f,
	0x70, 0x6f, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x29, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x15,
	0xf2, 0xde, 0x1f, 0x11, 0x73, 0x73, 0x7a, 0x2d, 0x6d, 0x61, 0x78, 0x3a, 0x22, 0x31, 0x30, 0x34,
	0x38, 0x35, 0x37, 0x36, 0x22, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x80, 0x01, 0x0a, 0x10,
	0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x53, 0x68, 0x61, 0x72, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x12, 0x3b, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x64, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2f, 0x0a,
	0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x42, 0x11, 0xf2, 0xde, 0x1f, 0x0d, 0x73, 0x73, 0x7a, 0x2d, 0x73, 0x69, 0x7a, 0x65, 0x3a, 0x22,
	0x39, 0x36, 0x22, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x93,
	0x02, 0x0a, 0x10, 0x53, 0x68, 0x61, 0x72, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x12, 0x3d, 0x0a, 0x11, 0x73, 0x68, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x11,
	0xf2, 0xde, 0x1f, 0x0d, 0x73, 0x73, 0x7a, 0x2d, 0x73, 0x69, 0x7a, 0x65, 0x3a, 0x22, 0x33, 0x32,
	0x22, 0x52, 0x0f, 0x73, 0x68, 0x61, 0x72, 0x64, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x6f,
	0x6f, 0x74, 0x12, 0x3f, 0x0a, 0x12, 0x62, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x11,
	0xf2, 0xde, 0x1f, 0x0d, 0x73, 0x73, 0x7a, 0x2d, 0x73, 0x69, 0x7a, 0x65, 0x3a, 0x22, 0x33, 0x32,
	0x22, 0x52, 0x10, 0x62, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x52,
	0x6f, 0x6f, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x68, 0x61, 0x72, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x73, 0x68, 0x61, 0x72, 0x64, 0x12, 0x25, 0x0a,
	0x0e, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x12, 0x2e, 0x0a, 0x09, 0x62, 0x6f, 0x64, 0x79, 0x5f, 0x72, 0x6f, 0x6f,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x11, 0xf2, 0xde, 0x1f, 0x0d, 0x73, 0x73, 0x7a,
	0x2d, 0x73, 0x69, 0x7a, 0x65, 0x3a, 0x22, 0x33, 0x32, 0x22, 0x52, 0x08, 0x62, 0x6f, 0x64, 0x79,
	0x52, 0x6f, 0x6f, 0x74, 0x22, 0x7c, 0x0a, 0x0a, 0x53, 0x68, 0x61, 0x72, 0x64, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x61, 0x73, 0x5f, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x67, 0x61, 0x73, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x11, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x11,
	0xf2, 0xde, 0x1f, 0x0d, 0x73, 0x73, 0x7a, 0x2d, 0x73, 0x69, 0x7a, 0x65, 0x3a, 0x22, 0x33, 0x32,
	0x22, 0x52, 0x0f, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x6f,
	0x6f, 0x74, 0x22, 0xdf, 0x02, 0x0a, 0x0f, 0x53, 0x68, 0x61, 0x72, 0x64, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x73, 0x6c, 0x6f, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x53, 0x6c, 0x6f, 0x74, 0x12, 0x40, 0x0a, 0x13, 0x73, 0x68, 0x61, 0x72, 0x64, 0x5f, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x04, 0x42, 0x10, 0xf2, 0xde, 0x1f, 0x0c, 0x73, 0x73, 0x7a, 0x2d, 0x6d, 0x61, 0x78, 0x3a,
	0x22, 0x31, 0x32, 0x22, 0x52, 0x11, 0x73, 0x68, 0x61, 0x72, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x73, 0x12, 0x3e, 0x0a, 0x10, 0x73, 0x68, 0x61, 0x72, 0x64,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0c, 0x42, 0x14, 0xf2, 0xde, 0x1f, 0x10, 0x73, 0x73, 0x7a, 0x2d, 0x73, 0x69, 0x7a, 0x65, 0x3a,
	0x22, 0x31, 0x32, 0x2c, 0x33, 0x32, 0x22, 0x52, 0x0e, 0x73, 0x68, 0x61, 0x72, 0x64, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x6f, 0x6f, 0x74, 0x73, 0x12, 0x56, 0x0a, 0x0c, 0x73, 0x68, 0x61, 0x72, 0x64,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x42, 0x10, 0xf2, 0xde, 0x1f, 0x0c, 0x73, 0x73, 0x7a, 0x2d, 0x6d, 0x61, 0x78, 0x3a, 0x22, 0x31,
	0x32, 0x22, 0x52, 0x0b, 0x73, 0x68, 0x61, 0x72, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x12,
	0x53, 0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x11, 0xf2, 0xde, 0x1f, 0x0d, 0x73, 0x73, 0x7a, 0x2d, 0x73,
	0x69, 0x7a, 0x65, 0x3a, 0x22, 0x39, 0x36, 0x22, 0x52, 0x1a, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73,
	0x65, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x41, 0x67, 0x67, 0x72, 0x65,
	0x67, 0x61, 0x74, 0x65, 0x22, 0x88, 0x01, 0x0a, 0x10, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74,
	0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x65, 0x12, 0x31, 0x0a, 0x08, 0x70, 0x75, 0x62,
	0x5f, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x42, 0x16, 0xf2, 0xde, 0x1f,
	0x12, 0x73, 0x73, 0x7a, 0x2d, 0x73, 0x69, 0x7a, 0x65, 0x3a, 0x22, 0x32, 0x30, 0x34, 0x38, 0x2c,
	0x34, 0x38, 0x22, 0x52, 0x07, 0x70, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x41, 0x0a, 0x12,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x04, 0x42, 0x12, 0xf2, 0xde, 0x1f, 0x0e, 0x73, 0x73,
	0x7a, 0x2d, 0x6d, 0x61, 0x78, 0x3a, 0x22, 0x32, 0x30, 0x34, 0x38, 0x22, 0x52, 0x11, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x63, 0x74, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x42,
	0x50, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72,
	0x79, 0x73, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x65, 0x74, 0x68, 0x65,
	0x72, 0x65, 0x75, 0x6d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x65, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x65, 0x74, 0x68, 0xaa, 0x02, 0x15, 0x45, 0x74, 0x68, 0x65,
	0x72, 0x65, 0x75, 0x6d, 0x2e, 0x45, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eth_v1alpha1_shard_proto_rawDescOnce sync.Once
	file_eth_v1alpha1_shard_proto_rawDescData = file_eth_v1alpha1_shard_proto_rawDesc
)

func file_eth_v1alpha1_shard_proto_rawDescGZIP() []byte {
	file_eth_v1alpha1_shard_proto_rawDescOnce.Do(func() {
		file_eth_v1alpha1_shard_proto_rawDescData = protoimpl.X.CompressGZIP(file_eth_v1alpha1_shard_proto_rawDescData)
	})
	return file_eth_v1alpha1_shard_proto_rawDescData
}

var file_eth_v1alpha1_shard_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_eth_v1alpha1_shard_proto_goTypes = []interface{}{
	(*ShardBlock)(nil),       // 0: ethereum.eth.v1alpha1.ShardBlock
	(*SignedShardBlock)(nil), // 1: ethereum.eth.v1alpha1.SignedShardBlock
	(*ShardBlockHeader)(nil), // 2: ethereum.eth.v1alpha1.ShardBlockHeader
	(*ShardState)(nil),       // 3: ethereum.eth.v1alpha1.ShardState
	(*ShardTransition)(nil),  // 4: ethereum.eth.v1alpha1.ShardTransition
	(*CompactCommittee)(nil), // 5: ethereum.eth.v1alpha1.CompactCommittee
}
var file_eth_v1alpha1_shard_proto_depIdxs = []int32{
	0, // 0: ethereum.eth.v1alpha1.SignedShardBlock.message:type_name -> ethereum.eth.v1alpha1.ShardBlock
	3, // 1: ethereum.eth.v1alpha1.ShardTransition.shard_states:type_name -> ethereum.eth.v1alpha1.ShardState
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_eth_v1alpha1_shard_proto_init() }
func file_eth_v1alpha1_shard_proto_init() {
	if File_eth_v1alpha1_shard_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eth_v1alpha1_shard_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShardBlock); i {
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
		file_eth_v1alpha1_shard_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignedShardBlock); i {
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
		file_eth_v1alpha1_shard_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShardBlockHeader); i {
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
		file_eth_v1alpha1_shard_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShardState); i {
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
		file_eth_v1alpha1_shard_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShardTransition); i {
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
		file_eth_v1alpha1_shard_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompactCommittee); i {
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
			RawDescriptor: file_eth_v1alpha1_shard_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eth_v1alpha1_shard_proto_goTypes,
		DependencyIndexes: file_eth_v1alpha1_shard_proto_depIdxs,
		MessageInfos:      file_eth_v1alpha1_shard_proto_msgTypes,
	}.Build()
	File_eth_v1alpha1_shard_proto = out.File
	file_eth_v1alpha1_shard_proto_rawDesc = nil
	file_eth_v1alpha1_shard_proto_goTypes = nil
	file_eth_v1alpha1_shard_proto_depIdxs = nil
}
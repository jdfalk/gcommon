// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: pkg/cache/proto/messages/eviction_result.proto

//go:build !protoopaque

package cachepb

import (
	proto "github.com/jdfalk/gcommon/pkg/common/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/gofeaturespb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// *
// Result of cache eviction operations.
// Provides details about items removed from cache.
type EvictionResult struct {
	state protoimpl.MessageState `protogen:"hybrid.v1"`
	// Number of items evicted
	EvictedCount *int64 `protobuf:"varint,1,opt,name=evicted_count,json=evictedCount" json:"evicted_count,omitempty"`
	// List of evicted keys
	EvictedKeys []string `protobuf:"bytes,2,rep,name=evicted_keys,json=evictedKeys" json:"evicted_keys,omitempty"`
	// Eviction policy used
	PolicyUsed *proto.EvictionPolicy `protobuf:"varint,3,opt,name=policy_used,json=policyUsed,enum=gcommon.v1.common.EvictionPolicy" json:"policy_used,omitempty"`
	// Reason for eviction
	EvictionReason *string `protobuf:"bytes,4,opt,name=eviction_reason,json=evictionReason" json:"eviction_reason,omitempty"`
	// Timestamp of eviction
	EvictedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=evicted_at,json=evictedAt" json:"evicted_at,omitempty"`
	// Memory freed by eviction (bytes)
	MemoryFreed *int64 `protobuf:"varint,6,opt,name=memory_freed,json=memoryFreed" json:"memory_freed,omitempty"`
	// Whether eviction was successful
	Success       *bool `protobuf:"varint,7,opt,name=success" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EvictionResult) Reset() {
	*x = EvictionResult{}
	mi := &file_pkg_cache_proto_messages_eviction_result_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EvictionResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvictionResult) ProtoMessage() {}

func (x *EvictionResult) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_cache_proto_messages_eviction_result_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *EvictionResult) GetEvictedCount() int64 {
	if x != nil && x.EvictedCount != nil {
		return *x.EvictedCount
	}
	return 0
}

func (x *EvictionResult) GetEvictedKeys() []string {
	if x != nil {
		return x.EvictedKeys
	}
	return nil
}

func (x *EvictionResult) GetPolicyUsed() proto.EvictionPolicy {
	if x != nil && x.PolicyUsed != nil {
		return *x.PolicyUsed
	}
	return proto.EvictionPolicy(0)
}

func (x *EvictionResult) GetEvictionReason() string {
	if x != nil && x.EvictionReason != nil {
		return *x.EvictionReason
	}
	return ""
}

func (x *EvictionResult) GetEvictedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.EvictedAt
	}
	return nil
}

func (x *EvictionResult) GetMemoryFreed() int64 {
	if x != nil && x.MemoryFreed != nil {
		return *x.MemoryFreed
	}
	return 0
}

func (x *EvictionResult) GetSuccess() bool {
	if x != nil && x.Success != nil {
		return *x.Success
	}
	return false
}

func (x *EvictionResult) SetEvictedCount(v int64) {
	x.EvictedCount = &v
}

func (x *EvictionResult) SetEvictedKeys(v []string) {
	x.EvictedKeys = v
}

func (x *EvictionResult) SetPolicyUsed(v proto.EvictionPolicy) {
	x.PolicyUsed = &v
}

func (x *EvictionResult) SetEvictionReason(v string) {
	x.EvictionReason = &v
}

func (x *EvictionResult) SetEvictedAt(v *timestamppb.Timestamp) {
	x.EvictedAt = v
}

func (x *EvictionResult) SetMemoryFreed(v int64) {
	x.MemoryFreed = &v
}

func (x *EvictionResult) SetSuccess(v bool) {
	x.Success = &v
}

func (x *EvictionResult) HasEvictedCount() bool {
	if x == nil {
		return false
	}
	return x.EvictedCount != nil
}

func (x *EvictionResult) HasPolicyUsed() bool {
	if x == nil {
		return false
	}
	return x.PolicyUsed != nil
}

func (x *EvictionResult) HasEvictionReason() bool {
	if x == nil {
		return false
	}
	return x.EvictionReason != nil
}

func (x *EvictionResult) HasEvictedAt() bool {
	if x == nil {
		return false
	}
	return x.EvictedAt != nil
}

func (x *EvictionResult) HasMemoryFreed() bool {
	if x == nil {
		return false
	}
	return x.MemoryFreed != nil
}

func (x *EvictionResult) HasSuccess() bool {
	if x == nil {
		return false
	}
	return x.Success != nil
}

func (x *EvictionResult) ClearEvictedCount() {
	x.EvictedCount = nil
}

func (x *EvictionResult) ClearPolicyUsed() {
	x.PolicyUsed = nil
}

func (x *EvictionResult) ClearEvictionReason() {
	x.EvictionReason = nil
}

func (x *EvictionResult) ClearEvictedAt() {
	x.EvictedAt = nil
}

func (x *EvictionResult) ClearMemoryFreed() {
	x.MemoryFreed = nil
}

func (x *EvictionResult) ClearSuccess() {
	x.Success = nil
}

type EvictionResult_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// Number of items evicted
	EvictedCount *int64
	// List of evicted keys
	EvictedKeys []string
	// Eviction policy used
	PolicyUsed *proto.EvictionPolicy
	// Reason for eviction
	EvictionReason *string
	// Timestamp of eviction
	EvictedAt *timestamppb.Timestamp
	// Memory freed by eviction (bytes)
	MemoryFreed *int64
	// Whether eviction was successful
	Success *bool
}

func (b0 EvictionResult_builder) Build() *EvictionResult {
	m0 := &EvictionResult{}
	b, x := &b0, m0
	_, _ = b, x
	x.EvictedCount = b.EvictedCount
	x.EvictedKeys = b.EvictedKeys
	x.PolicyUsed = b.PolicyUsed
	x.EvictionReason = b.EvictionReason
	x.EvictedAt = b.EvictedAt
	x.MemoryFreed = b.MemoryFreed
	x.Success = b.Success
	return m0
}

var File_pkg_cache_proto_messages_eviction_result_proto protoreflect.FileDescriptor

const file_pkg_cache_proto_messages_eviction_result_proto_rawDesc = "" +
	"\n" +
	".pkg/cache/proto/messages/eviction_result.proto\x12\x10gcommon.v1.cache\x1a!google/protobuf/go_features.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a,pkg/common/proto/enums/eviction_policy.proto\"\xbd\x02\n" +
	"\x0eEvictionResult\x12#\n" +
	"\revicted_count\x18\x01 \x01(\x03R\fevictedCount\x12!\n" +
	"\fevicted_keys\x18\x02 \x03(\tR\vevictedKeys\x12B\n" +
	"\vpolicy_used\x18\x03 \x01(\x0e2!.gcommon.v1.common.EvictionPolicyR\n" +
	"policyUsed\x12'\n" +
	"\x0feviction_reason\x18\x04 \x01(\tR\x0eevictionReason\x129\n" +
	"\n" +
	"evicted_at\x18\x05 \x01(\v2\x1a.google.protobuf.TimestampR\tevictedAt\x12!\n" +
	"\fmemory_freed\x18\x06 \x01(\x03R\vmemoryFreed\x12\x18\n" +
	"\asuccess\x18\a \x01(\bR\asuccessB\xc8\x01\n" +
	"\x14com.gcommon.v1.cacheB\x13EvictionResultProtoP\x01Z1github.com/jdfalk/gcommon/pkg/cache/proto;cachepb\xa2\x02\x03GVC\xaa\x02\x10Gcommon.V1.Cache\xca\x02\x10Gcommon\\V1\\Cache\xe2\x02\x1cGcommon\\V1\\Cache\\GPBMetadata\xea\x02\x12Gcommon::V1::Cache\x92\x03\x05\xd2>\x02\x10\x02b\beditionsp\xe8\a"

var file_pkg_cache_proto_messages_eviction_result_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pkg_cache_proto_messages_eviction_result_proto_goTypes = []any{
	(*EvictionResult)(nil),        // 0: gcommon.v1.cache.EvictionResult
	(proto.EvictionPolicy)(0),     // 1: gcommon.v1.common.EvictionPolicy
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_pkg_cache_proto_messages_eviction_result_proto_depIdxs = []int32{
	1, // 0: gcommon.v1.cache.EvictionResult.policy_used:type_name -> gcommon.v1.common.EvictionPolicy
	2, // 1: gcommon.v1.cache.EvictionResult.evicted_at:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_cache_proto_messages_eviction_result_proto_init() }
func file_pkg_cache_proto_messages_eviction_result_proto_init() {
	if File_pkg_cache_proto_messages_eviction_result_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pkg_cache_proto_messages_eviction_result_proto_rawDesc), len(file_pkg_cache_proto_messages_eviction_result_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_cache_proto_messages_eviction_result_proto_goTypes,
		DependencyIndexes: file_pkg_cache_proto_messages_eviction_result_proto_depIdxs,
		MessageInfos:      file_pkg_cache_proto_messages_eviction_result_proto_msgTypes,
	}.Build()
	File_pkg_cache_proto_messages_eviction_result_proto = out.File
	file_pkg_cache_proto_messages_eviction_result_proto_goTypes = nil
	file_pkg_cache_proto_messages_eviction_result_proto_depIdxs = nil
}

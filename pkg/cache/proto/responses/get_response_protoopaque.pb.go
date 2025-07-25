// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: pkg/cache/proto/responses/get_response.proto

//go:build protoopaque

package cachepb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/gofeaturespb"
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
// Response containing a cached value and metadata.
// Includes cache hit/miss information and entry details.
type GetResponse struct {
	state               protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Entry    *CacheEntry            `protobuf:"bytes,1,opt,name=entry"`
	xxx_hidden_Found    bool                   `protobuf:"varint,2,opt,name=found"`
	xxx_hidden_CacheHit bool                   `protobuf:"varint,3,opt,name=cache_hit,json=cacheHit"`
	// Deprecated: Do not use. This will be deleted in the near future.
	XXX_lazyUnmarshalInfo  protoimpl.LazyUnmarshalInfo
	XXX_raceDetectHookData protoimpl.RaceDetectHookData
	XXX_presence           [1]uint32
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	mi := &file_pkg_cache_proto_responses_get_response_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_cache_proto_responses_get_response_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *GetResponse) GetEntry() *CacheEntry {
	if x != nil {
		if protoimpl.X.Present(&(x.XXX_presence[0]), 0) {
			if protoimpl.X.AtomicCheckPointerIsNil(&x.xxx_hidden_Entry) {
				protoimpl.X.UnmarshalField(x, 1)
			}
			var rv *CacheEntry
			protoimpl.X.AtomicLoadPointer(protoimpl.Pointer(&x.xxx_hidden_Entry), protoimpl.Pointer(&rv))
			return rv
		}
	}
	return nil
}

func (x *GetResponse) GetFound() bool {
	if x != nil {
		return x.xxx_hidden_Found
	}
	return false
}

func (x *GetResponse) GetCacheHit() bool {
	if x != nil {
		return x.xxx_hidden_CacheHit
	}
	return false
}

func (x *GetResponse) SetEntry(v *CacheEntry) {
	protoimpl.X.AtomicSetPointer(&x.xxx_hidden_Entry, v)
	if v == nil {
		protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	} else {
		protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 3)
	}
}

func (x *GetResponse) SetFound(v bool) {
	x.xxx_hidden_Found = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 1, 3)
}

func (x *GetResponse) SetCacheHit(v bool) {
	x.xxx_hidden_CacheHit = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 2, 3)
}

func (x *GetResponse) HasEntry() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *GetResponse) HasFound() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 1)
}

func (x *GetResponse) HasCacheHit() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 2)
}

func (x *GetResponse) ClearEntry() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	protoimpl.X.AtomicSetPointer(&x.xxx_hidden_Entry, (*CacheEntry)(nil))
}

func (x *GetResponse) ClearFound() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 1)
	x.xxx_hidden_Found = false
}

func (x *GetResponse) ClearCacheHit() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 2)
	x.xxx_hidden_CacheHit = false
}

type GetResponse_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// The cached entry (only present if found)
	Entry *CacheEntry
	// Whether the key was found in the cache
	Found *bool
	// Cache hit/miss information for metrics
	CacheHit *bool
}

func (b0 GetResponse_builder) Build() *GetResponse {
	m0 := &GetResponse{}
	b, x := &b0, m0
	_, _ = b, x
	if b.Entry != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 3)
		x.xxx_hidden_Entry = b.Entry
	}
	if b.Found != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 1, 3)
		x.xxx_hidden_Found = *b.Found
	}
	if b.CacheHit != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 2, 3)
		x.xxx_hidden_CacheHit = *b.CacheHit
	}
	return m0
}

var File_pkg_cache_proto_responses_get_response_proto protoreflect.FileDescriptor

const file_pkg_cache_proto_responses_get_response_proto_rawDesc = "" +
	"\n" +
	",pkg/cache/proto/responses/get_response.proto\x12\x10gcommon.v1.cache\x1a!google/protobuf/go_features.proto\x1a*pkg/cache/proto/messages/cache_entry.proto\"x\n" +
	"\vGetResponse\x126\n" +
	"\x05entry\x18\x01 \x01(\v2\x1c.gcommon.v1.cache.CacheEntryB\x02(\x01R\x05entry\x12\x14\n" +
	"\x05found\x18\x02 \x01(\bR\x05found\x12\x1b\n" +
	"\tcache_hit\x18\x03 \x01(\bR\bcacheHitB\xc5\x01\n" +
	"\x14com.gcommon.v1.cacheB\x10GetResponseProtoP\x01Z1github.com/jdfalk/gcommon/pkg/cache/proto;cachepb\xa2\x02\x03GVC\xaa\x02\x10Gcommon.V1.Cache\xca\x02\x10Gcommon\\V1\\Cache\xe2\x02\x1cGcommon\\V1\\Cache\\GPBMetadata\xea\x02\x12Gcommon::V1::Cache\x92\x03\x05\xd2>\x02\x10\x02b\beditionsp\xe8\a"

var file_pkg_cache_proto_responses_get_response_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pkg_cache_proto_responses_get_response_proto_goTypes = []any{
	(*GetResponse)(nil), // 0: gcommon.v1.cache.GetResponse
	(*CacheEntry)(nil),  // 1: gcommon.v1.cache.CacheEntry
}
var file_pkg_cache_proto_responses_get_response_proto_depIdxs = []int32{
	1, // 0: gcommon.v1.cache.GetResponse.entry:type_name -> gcommon.v1.cache.CacheEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_cache_proto_responses_get_response_proto_init() }
func file_pkg_cache_proto_responses_get_response_proto_init() {
	if File_pkg_cache_proto_responses_get_response_proto != nil {
		return
	}
	file_pkg_cache_proto_messages_cache_entry_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pkg_cache_proto_responses_get_response_proto_rawDesc), len(file_pkg_cache_proto_responses_get_response_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_cache_proto_responses_get_response_proto_goTypes,
		DependencyIndexes: file_pkg_cache_proto_responses_get_response_proto_depIdxs,
		MessageInfos:      file_pkg_cache_proto_responses_get_response_proto_msgTypes,
	}.Build()
	File_pkg_cache_proto_responses_get_response_proto = out.File
	file_pkg_cache_proto_responses_get_response_proto_goTypes = nil
	file_pkg_cache_proto_responses_get_response_proto_depIdxs = nil
}

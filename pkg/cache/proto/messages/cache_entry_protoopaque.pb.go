// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: pkg/cache/proto/messages/cache_entry.proto

//go:build protoopaque

package cachepb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/gofeaturespb"
	anypb "google.golang.org/protobuf/types/known/anypb"
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
// Cache entry containing the cached value and metadata.
// Supports multiple value types with comprehensive expiration
// and access tracking for cache policies.
type CacheEntry struct {
	state                     protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Key            *string                `protobuf:"bytes,1,opt,name=key"`
	xxx_hidden_Value          *anypb.Any             `protobuf:"bytes,2,opt,name=value"`
	xxx_hidden_CreatedAt      *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt"`
	xxx_hidden_LastAccessedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=last_accessed_at,json=lastAccessedAt"`
	xxx_hidden_ExpiresAt      *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=expires_at,json=expiresAt"`
	xxx_hidden_AccessCount    int64                  `protobuf:"varint,6,opt,name=access_count,json=accessCount"`
	xxx_hidden_SizeBytes      int64                  `protobuf:"varint,7,opt,name=size_bytes,json=sizeBytes"`
	xxx_hidden_Metadata       map[string]string      `protobuf:"bytes,8,rep,name=metadata" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	xxx_hidden_Namespace      *string                `protobuf:"bytes,9,opt,name=namespace"`
	// Deprecated: Do not use. This will be deleted in the near future.
	XXX_lazyUnmarshalInfo  protoimpl.LazyUnmarshalInfo
	XXX_raceDetectHookData protoimpl.RaceDetectHookData
	XXX_presence           [1]uint32
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *CacheEntry) Reset() {
	*x = CacheEntry{}
	mi := &file_pkg_cache_proto_messages_cache_entry_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CacheEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CacheEntry) ProtoMessage() {}

func (x *CacheEntry) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_cache_proto_messages_cache_entry_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *CacheEntry) GetKey() string {
	if x != nil {
		if x.xxx_hidden_Key != nil {
			return *x.xxx_hidden_Key
		}
		return ""
	}
	return ""
}

func (x *CacheEntry) GetValue() *anypb.Any {
	if x != nil {
		if protoimpl.X.Present(&(x.XXX_presence[0]), 1) {
			if protoimpl.X.AtomicCheckPointerIsNil(&x.xxx_hidden_Value) {
				protoimpl.X.UnmarshalField(x, 2)
			}
			var rv *anypb.Any
			protoimpl.X.AtomicLoadPointer(protoimpl.Pointer(&x.xxx_hidden_Value), protoimpl.Pointer(&rv))
			return rv
		}
	}
	return nil
}

func (x *CacheEntry) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		if protoimpl.X.Present(&(x.XXX_presence[0]), 2) {
			if protoimpl.X.AtomicCheckPointerIsNil(&x.xxx_hidden_CreatedAt) {
				protoimpl.X.UnmarshalField(x, 3)
			}
			var rv *timestamppb.Timestamp
			protoimpl.X.AtomicLoadPointer(protoimpl.Pointer(&x.xxx_hidden_CreatedAt), protoimpl.Pointer(&rv))
			return rv
		}
	}
	return nil
}

func (x *CacheEntry) GetLastAccessedAt() *timestamppb.Timestamp {
	if x != nil {
		if protoimpl.X.Present(&(x.XXX_presence[0]), 3) {
			if protoimpl.X.AtomicCheckPointerIsNil(&x.xxx_hidden_LastAccessedAt) {
				protoimpl.X.UnmarshalField(x, 4)
			}
			var rv *timestamppb.Timestamp
			protoimpl.X.AtomicLoadPointer(protoimpl.Pointer(&x.xxx_hidden_LastAccessedAt), protoimpl.Pointer(&rv))
			return rv
		}
	}
	return nil
}

func (x *CacheEntry) GetExpiresAt() *timestamppb.Timestamp {
	if x != nil {
		if protoimpl.X.Present(&(x.XXX_presence[0]), 4) {
			if protoimpl.X.AtomicCheckPointerIsNil(&x.xxx_hidden_ExpiresAt) {
				protoimpl.X.UnmarshalField(x, 5)
			}
			var rv *timestamppb.Timestamp
			protoimpl.X.AtomicLoadPointer(protoimpl.Pointer(&x.xxx_hidden_ExpiresAt), protoimpl.Pointer(&rv))
			return rv
		}
	}
	return nil
}

func (x *CacheEntry) GetAccessCount() int64 {
	if x != nil {
		return x.xxx_hidden_AccessCount
	}
	return 0
}

func (x *CacheEntry) GetSizeBytes() int64 {
	if x != nil {
		return x.xxx_hidden_SizeBytes
	}
	return 0
}

func (x *CacheEntry) GetMetadata() map[string]string {
	if x != nil {
		if protoimpl.X.Present(&(x.XXX_presence[0]), 7) {
			if protoimpl.X.AtomicCheckPointerIsNil(&x.xxx_hidden_Metadata) {
				protoimpl.X.UnmarshalField(x, 8)
			}
			return x.xxx_hidden_Metadata
		}
	}
	return nil
}

func (x *CacheEntry) GetNamespace() string {
	if x != nil {
		if x.xxx_hidden_Namespace != nil {
			return *x.xxx_hidden_Namespace
		}
		return ""
	}
	return ""
}

func (x *CacheEntry) SetKey(v string) {
	x.xxx_hidden_Key = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 9)
}

func (x *CacheEntry) SetValue(v *anypb.Any) {
	protoimpl.X.AtomicSetPointer(&x.xxx_hidden_Value, v)
	if v == nil {
		protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 1)
	} else {
		protoimpl.X.SetPresent(&(x.XXX_presence[0]), 1, 9)
	}
}

func (x *CacheEntry) SetCreatedAt(v *timestamppb.Timestamp) {
	protoimpl.X.AtomicSetPointer(&x.xxx_hidden_CreatedAt, v)
	if v == nil {
		protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 2)
	} else {
		protoimpl.X.SetPresent(&(x.XXX_presence[0]), 2, 9)
	}
}

func (x *CacheEntry) SetLastAccessedAt(v *timestamppb.Timestamp) {
	protoimpl.X.AtomicSetPointer(&x.xxx_hidden_LastAccessedAt, v)
	if v == nil {
		protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 3)
	} else {
		protoimpl.X.SetPresent(&(x.XXX_presence[0]), 3, 9)
	}
}

func (x *CacheEntry) SetExpiresAt(v *timestamppb.Timestamp) {
	protoimpl.X.AtomicSetPointer(&x.xxx_hidden_ExpiresAt, v)
	if v == nil {
		protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 4)
	} else {
		protoimpl.X.SetPresent(&(x.XXX_presence[0]), 4, 9)
	}
}

func (x *CacheEntry) SetAccessCount(v int64) {
	x.xxx_hidden_AccessCount = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 5, 9)
}

func (x *CacheEntry) SetSizeBytes(v int64) {
	x.xxx_hidden_SizeBytes = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 6, 9)
}

func (x *CacheEntry) SetMetadata(v map[string]string) {
	x.xxx_hidden_Metadata = v
	if v == nil {
		protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 7)
	} else {
		protoimpl.X.SetPresent(&(x.XXX_presence[0]), 7, 9)
	}
}

func (x *CacheEntry) SetNamespace(v string) {
	x.xxx_hidden_Namespace = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 8, 9)
}

func (x *CacheEntry) HasKey() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *CacheEntry) HasValue() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 1)
}

func (x *CacheEntry) HasCreatedAt() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 2)
}

func (x *CacheEntry) HasLastAccessedAt() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 3)
}

func (x *CacheEntry) HasExpiresAt() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 4)
}

func (x *CacheEntry) HasAccessCount() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 5)
}

func (x *CacheEntry) HasSizeBytes() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 6)
}

func (x *CacheEntry) HasNamespace() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 8)
}

func (x *CacheEntry) ClearKey() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_Key = nil
}

func (x *CacheEntry) ClearValue() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 1)
	protoimpl.X.AtomicSetPointer(&x.xxx_hidden_Value, (*anypb.Any)(nil))
}

func (x *CacheEntry) ClearCreatedAt() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 2)
	protoimpl.X.AtomicSetPointer(&x.xxx_hidden_CreatedAt, (*timestamppb.Timestamp)(nil))
}

func (x *CacheEntry) ClearLastAccessedAt() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 3)
	protoimpl.X.AtomicSetPointer(&x.xxx_hidden_LastAccessedAt, (*timestamppb.Timestamp)(nil))
}

func (x *CacheEntry) ClearExpiresAt() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 4)
	protoimpl.X.AtomicSetPointer(&x.xxx_hidden_ExpiresAt, (*timestamppb.Timestamp)(nil))
}

func (x *CacheEntry) ClearAccessCount() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 5)
	x.xxx_hidden_AccessCount = 0
}

func (x *CacheEntry) ClearSizeBytes() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 6)
	x.xxx_hidden_SizeBytes = 0
}

func (x *CacheEntry) ClearNamespace() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 8)
	x.xxx_hidden_Namespace = nil
}

type CacheEntry_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// Cache key (immutable)
	Key *string
	// The cached value (flexible type support)
	Value *anypb.Any
	// When the entry was created
	CreatedAt *timestamppb.Timestamp
	// When the entry was last accessed
	LastAccessedAt *timestamppb.Timestamp
	// When the entry expires (optional)
	ExpiresAt *timestamppb.Timestamp
	// Number of times the entry has been accessed
	AccessCount *int64
	// Size of the entry in bytes
	SizeBytes *int64
	// Entry metadata for extensibility
	Metadata map[string]string
	// Cache namespace this entry belongs to
	Namespace *string
}

func (b0 CacheEntry_builder) Build() *CacheEntry {
	m0 := &CacheEntry{}
	b, x := &b0, m0
	_, _ = b, x
	if b.Key != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 9)
		x.xxx_hidden_Key = b.Key
	}
	if b.Value != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 1, 9)
		x.xxx_hidden_Value = b.Value
	}
	if b.CreatedAt != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 2, 9)
		x.xxx_hidden_CreatedAt = b.CreatedAt
	}
	if b.LastAccessedAt != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 3, 9)
		x.xxx_hidden_LastAccessedAt = b.LastAccessedAt
	}
	if b.ExpiresAt != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 4, 9)
		x.xxx_hidden_ExpiresAt = b.ExpiresAt
	}
	if b.AccessCount != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 5, 9)
		x.xxx_hidden_AccessCount = *b.AccessCount
	}
	if b.SizeBytes != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 6, 9)
		x.xxx_hidden_SizeBytes = *b.SizeBytes
	}
	if b.Metadata != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 7, 9)
		x.xxx_hidden_Metadata = b.Metadata
	}
	if b.Namespace != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 8, 9)
		x.xxx_hidden_Namespace = b.Namespace
	}
	return m0
}

var File_pkg_cache_proto_messages_cache_entry_proto protoreflect.FileDescriptor

const file_pkg_cache_proto_messages_cache_entry_proto_rawDesc = "" +
	"\n" +
	"*pkg/cache/proto/messages/cache_entry.proto\x12\x10gcommon.v1.cache\x1a!google/protobuf/go_features.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x19google/protobuf/any.proto\"\xff\x03\n" +
	"\n" +
	"CacheEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12.\n" +
	"\x05value\x18\x02 \x01(\v2\x14.google.protobuf.AnyB\x02(\x01R\x05value\x12=\n" +
	"\n" +
	"created_at\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampB\x02(\x01R\tcreatedAt\x12H\n" +
	"\x10last_accessed_at\x18\x04 \x01(\v2\x1a.google.protobuf.TimestampB\x02(\x01R\x0elastAccessedAt\x12=\n" +
	"\n" +
	"expires_at\x18\x05 \x01(\v2\x1a.google.protobuf.TimestampB\x02(\x01R\texpiresAt\x12!\n" +
	"\faccess_count\x18\x06 \x01(\x03R\vaccessCount\x12\x1d\n" +
	"\n" +
	"size_bytes\x18\a \x01(\x03R\tsizeBytes\x12J\n" +
	"\bmetadata\x18\b \x03(\v2*.gcommon.v1.cache.CacheEntry.MetadataEntryB\x02(\x01R\bmetadata\x12\x1c\n" +
	"\tnamespace\x18\t \x01(\tR\tnamespace\x1a;\n" +
	"\rMetadataEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01B\xc4\x01\n" +
	"\x14com.gcommon.v1.cacheB\x0fCacheEntryProtoP\x01Z1github.com/jdfalk/gcommon/pkg/cache/proto;cachepb\xa2\x02\x03GVC\xaa\x02\x10Gcommon.V1.Cache\xca\x02\x10Gcommon\\V1\\Cache\xe2\x02\x1cGcommon\\V1\\Cache\\GPBMetadata\xea\x02\x12Gcommon::V1::Cache\x92\x03\x05\xd2>\x02\x10\x02b\beditionsp\xe8\a"

var file_pkg_cache_proto_messages_cache_entry_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_cache_proto_messages_cache_entry_proto_goTypes = []any{
	(*CacheEntry)(nil),            // 0: gcommon.v1.cache.CacheEntry
	nil,                           // 1: gcommon.v1.cache.CacheEntry.MetadataEntry
	(*anypb.Any)(nil),             // 2: google.protobuf.Any
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_pkg_cache_proto_messages_cache_entry_proto_depIdxs = []int32{
	2, // 0: gcommon.v1.cache.CacheEntry.value:type_name -> google.protobuf.Any
	3, // 1: gcommon.v1.cache.CacheEntry.created_at:type_name -> google.protobuf.Timestamp
	3, // 2: gcommon.v1.cache.CacheEntry.last_accessed_at:type_name -> google.protobuf.Timestamp
	3, // 3: gcommon.v1.cache.CacheEntry.expires_at:type_name -> google.protobuf.Timestamp
	1, // 4: gcommon.v1.cache.CacheEntry.metadata:type_name -> gcommon.v1.cache.CacheEntry.MetadataEntry
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_pkg_cache_proto_messages_cache_entry_proto_init() }
func file_pkg_cache_proto_messages_cache_entry_proto_init() {
	if File_pkg_cache_proto_messages_cache_entry_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pkg_cache_proto_messages_cache_entry_proto_rawDesc), len(file_pkg_cache_proto_messages_cache_entry_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_cache_proto_messages_cache_entry_proto_goTypes,
		DependencyIndexes: file_pkg_cache_proto_messages_cache_entry_proto_depIdxs,
		MessageInfos:      file_pkg_cache_proto_messages_cache_entry_proto_msgTypes,
	}.Build()
	File_pkg_cache_proto_messages_cache_entry_proto = out.File
	file_pkg_cache_proto_messages_cache_entry_proto_goTypes = nil
	file_pkg_cache_proto_messages_cache_entry_proto_depIdxs = nil
}

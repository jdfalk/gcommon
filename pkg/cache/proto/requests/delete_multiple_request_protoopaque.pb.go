// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: pkg/cache/proto/requests/delete_multiple_request.proto

//go:build protoopaque

package cachepb

import (
	proto "github.com/jdfalk/gcommon/pkg/common/proto"
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
// Request to delete multiple cache keys.
type DeleteMultipleRequest struct {
	state                protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Keys      []string               `protobuf:"bytes,1,rep,name=keys"`
	xxx_hidden_Namespace *string                `protobuf:"bytes,2,opt,name=namespace"`
	xxx_hidden_Metadata  *proto.RequestMetadata `protobuf:"bytes,3,opt,name=metadata"`
	// Deprecated: Do not use. This will be deleted in the near future.
	XXX_lazyUnmarshalInfo  protoimpl.LazyUnmarshalInfo
	XXX_raceDetectHookData protoimpl.RaceDetectHookData
	XXX_presence           [1]uint32
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *DeleteMultipleRequest) Reset() {
	*x = DeleteMultipleRequest{}
	mi := &file_pkg_cache_proto_requests_delete_multiple_request_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteMultipleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMultipleRequest) ProtoMessage() {}

func (x *DeleteMultipleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_cache_proto_requests_delete_multiple_request_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *DeleteMultipleRequest) GetKeys() []string {
	if x != nil {
		return x.xxx_hidden_Keys
	}
	return nil
}

func (x *DeleteMultipleRequest) GetNamespace() string {
	if x != nil {
		if x.xxx_hidden_Namespace != nil {
			return *x.xxx_hidden_Namespace
		}
		return ""
	}
	return ""
}

func (x *DeleteMultipleRequest) GetMetadata() *proto.RequestMetadata {
	if x != nil {
		if protoimpl.X.Present(&(x.XXX_presence[0]), 2) {
			if protoimpl.X.AtomicCheckPointerIsNil(&x.xxx_hidden_Metadata) {
				protoimpl.X.UnmarshalField(x, 3)
			}
			var rv *proto.RequestMetadata
			protoimpl.X.AtomicLoadPointer(protoimpl.Pointer(&x.xxx_hidden_Metadata), protoimpl.Pointer(&rv))
			return rv
		}
	}
	return nil
}

func (x *DeleteMultipleRequest) SetKeys(v []string) {
	x.xxx_hidden_Keys = v
}

func (x *DeleteMultipleRequest) SetNamespace(v string) {
	x.xxx_hidden_Namespace = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 1, 3)
}

func (x *DeleteMultipleRequest) SetMetadata(v *proto.RequestMetadata) {
	protoimpl.X.AtomicSetPointer(&x.xxx_hidden_Metadata, v)
	if v == nil {
		protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 2)
	} else {
		protoimpl.X.SetPresent(&(x.XXX_presence[0]), 2, 3)
	}
}

func (x *DeleteMultipleRequest) HasNamespace() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 1)
}

func (x *DeleteMultipleRequest) HasMetadata() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 2)
}

func (x *DeleteMultipleRequest) ClearNamespace() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 1)
	x.xxx_hidden_Namespace = nil
}

func (x *DeleteMultipleRequest) ClearMetadata() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 2)
	protoimpl.X.AtomicSetPointer(&x.xxx_hidden_Metadata, (*proto.RequestMetadata)(nil))
}

type DeleteMultipleRequest_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// Keys to delete
	Keys []string
	// Optional namespace
	Namespace *string
	// Request metadata
	Metadata *proto.RequestMetadata
}

func (b0 DeleteMultipleRequest_builder) Build() *DeleteMultipleRequest {
	m0 := &DeleteMultipleRequest{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Keys = b.Keys
	if b.Namespace != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 1, 3)
		x.xxx_hidden_Namespace = b.Namespace
	}
	if b.Metadata != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 2, 3)
		x.xxx_hidden_Metadata = b.Metadata
	}
	return m0
}

var File_pkg_cache_proto_requests_delete_multiple_request_proto protoreflect.FileDescriptor

const file_pkg_cache_proto_requests_delete_multiple_request_proto_rawDesc = "" +
	"\n" +
	"6pkg/cache/proto/requests/delete_multiple_request.proto\x12\x10gcommon.v1.cache\x1a!google/protobuf/go_features.proto\x1a0pkg/common/proto/messages/request_metadata.proto\"\x8d\x01\n" +
	"\x15DeleteMultipleRequest\x12\x12\n" +
	"\x04keys\x18\x01 \x03(\tR\x04keys\x12\x1c\n" +
	"\tnamespace\x18\x02 \x01(\tR\tnamespace\x12B\n" +
	"\bmetadata\x18\x03 \x01(\v2\".gcommon.v1.common.RequestMetadataB\x02(\x01R\bmetadataB\xcf\x01\n" +
	"\x14com.gcommon.v1.cacheB\x1aDeleteMultipleRequestProtoP\x01Z1github.com/jdfalk/gcommon/pkg/cache/proto;cachepb\xa2\x02\x03GVC\xaa\x02\x10Gcommon.V1.Cache\xca\x02\x10Gcommon\\V1\\Cache\xe2\x02\x1cGcommon\\V1\\Cache\\GPBMetadata\xea\x02\x12Gcommon::V1::Cache\x92\x03\x05\xd2>\x02\x10\x02b\beditionsp\xe8\a"

var file_pkg_cache_proto_requests_delete_multiple_request_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pkg_cache_proto_requests_delete_multiple_request_proto_goTypes = []any{
	(*DeleteMultipleRequest)(nil), // 0: gcommon.v1.cache.DeleteMultipleRequest
	(*proto.RequestMetadata)(nil), // 1: gcommon.v1.common.RequestMetadata
}
var file_pkg_cache_proto_requests_delete_multiple_request_proto_depIdxs = []int32{
	1, // 0: gcommon.v1.cache.DeleteMultipleRequest.metadata:type_name -> gcommon.v1.common.RequestMetadata
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_cache_proto_requests_delete_multiple_request_proto_init() }
func file_pkg_cache_proto_requests_delete_multiple_request_proto_init() {
	if File_pkg_cache_proto_requests_delete_multiple_request_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pkg_cache_proto_requests_delete_multiple_request_proto_rawDesc), len(file_pkg_cache_proto_requests_delete_multiple_request_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_cache_proto_requests_delete_multiple_request_proto_goTypes,
		DependencyIndexes: file_pkg_cache_proto_requests_delete_multiple_request_proto_depIdxs,
		MessageInfos:      file_pkg_cache_proto_requests_delete_multiple_request_proto_msgTypes,
	}.Build()
	File_pkg_cache_proto_requests_delete_multiple_request_proto = out.File
	file_pkg_cache_proto_requests_delete_multiple_request_proto_goTypes = nil
	file_pkg_cache_proto_requests_delete_multiple_request_proto_depIdxs = nil
}

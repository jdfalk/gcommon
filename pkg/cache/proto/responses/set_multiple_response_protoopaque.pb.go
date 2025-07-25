// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: pkg/cache/proto/responses/set_multiple_response.proto

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
// Response for multiple cache value set operations.
// Indicates success/failure of batch set operation.
type SetMultipleResponse struct {
	state                  protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Success     bool                   `protobuf:"varint,1,opt,name=success"`
	xxx_hidden_FailedKeys  []string               `protobuf:"bytes,2,rep,name=failed_keys,json=failedKeys"`
	xxx_hidden_Error       *proto.Error           `protobuf:"bytes,3,opt,name=error"`
	xxx_hidden_SetCount    int32                  `protobuf:"varint,4,opt,name=set_count,json=setCount"`
	XXX_raceDetectHookData protoimpl.RaceDetectHookData
	XXX_presence           [1]uint32
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *SetMultipleResponse) Reset() {
	*x = SetMultipleResponse{}
	mi := &file_pkg_cache_proto_responses_set_multiple_response_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetMultipleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetMultipleResponse) ProtoMessage() {}

func (x *SetMultipleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_cache_proto_responses_set_multiple_response_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *SetMultipleResponse) GetSuccess() bool {
	if x != nil {
		return x.xxx_hidden_Success
	}
	return false
}

func (x *SetMultipleResponse) GetFailedKeys() []string {
	if x != nil {
		return x.xxx_hidden_FailedKeys
	}
	return nil
}

func (x *SetMultipleResponse) GetError() *proto.Error {
	if x != nil {
		return x.xxx_hidden_Error
	}
	return nil
}

func (x *SetMultipleResponse) GetSetCount() int32 {
	if x != nil {
		return x.xxx_hidden_SetCount
	}
	return 0
}

func (x *SetMultipleResponse) SetSuccess(v bool) {
	x.xxx_hidden_Success = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 4)
}

func (x *SetMultipleResponse) SetFailedKeys(v []string) {
	x.xxx_hidden_FailedKeys = v
}

func (x *SetMultipleResponse) SetError(v *proto.Error) {
	x.xxx_hidden_Error = v
}

func (x *SetMultipleResponse) SetSetCount(v int32) {
	x.xxx_hidden_SetCount = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 3, 4)
}

func (x *SetMultipleResponse) HasSuccess() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *SetMultipleResponse) HasError() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Error != nil
}

func (x *SetMultipleResponse) HasSetCount() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 3)
}

func (x *SetMultipleResponse) ClearSuccess() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_Success = false
}

func (x *SetMultipleResponse) ClearError() {
	x.xxx_hidden_Error = nil
}

func (x *SetMultipleResponse) ClearSetCount() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 3)
	x.xxx_hidden_SetCount = 0
}

type SetMultipleResponse_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// Whether all values were successfully set
	Success *bool
	// List of keys that failed to be set
	FailedKeys []string
	// Error details if operation failed
	Error *proto.Error
	// Number of keys that were successfully set
	SetCount *int32
}

func (b0 SetMultipleResponse_builder) Build() *SetMultipleResponse {
	m0 := &SetMultipleResponse{}
	b, x := &b0, m0
	_, _ = b, x
	if b.Success != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 4)
		x.xxx_hidden_Success = *b.Success
	}
	x.xxx_hidden_FailedKeys = b.FailedKeys
	x.xxx_hidden_Error = b.Error
	if b.SetCount != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 3, 4)
		x.xxx_hidden_SetCount = *b.SetCount
	}
	return m0
}

var File_pkg_cache_proto_responses_set_multiple_response_proto protoreflect.FileDescriptor

const file_pkg_cache_proto_responses_set_multiple_response_proto_rawDesc = "" +
	"\n" +
	"5pkg/cache/proto/responses/set_multiple_response.proto\x12\x10gcommon.v1.cache\x1a!google/protobuf/go_features.proto\x1a%pkg/common/proto/messages/error.proto\"\x9d\x01\n" +
	"\x13SetMultipleResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\x12\x1f\n" +
	"\vfailed_keys\x18\x02 \x03(\tR\n" +
	"failedKeys\x12.\n" +
	"\x05error\x18\x03 \x01(\v2\x18.gcommon.v1.common.ErrorR\x05error\x12\x1b\n" +
	"\tset_count\x18\x04 \x01(\x05R\bsetCountB\xcd\x01\n" +
	"\x14com.gcommon.v1.cacheB\x18SetMultipleResponseProtoP\x01Z1github.com/jdfalk/gcommon/pkg/cache/proto;cachepb\xa2\x02\x03GVC\xaa\x02\x10Gcommon.V1.Cache\xca\x02\x10Gcommon\\V1\\Cache\xe2\x02\x1cGcommon\\V1\\Cache\\GPBMetadata\xea\x02\x12Gcommon::V1::Cache\x92\x03\x05\xd2>\x02\x10\x02b\beditionsp\xe8\a"

var file_pkg_cache_proto_responses_set_multiple_response_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pkg_cache_proto_responses_set_multiple_response_proto_goTypes = []any{
	(*SetMultipleResponse)(nil), // 0: gcommon.v1.cache.SetMultipleResponse
	(*proto.Error)(nil),         // 1: gcommon.v1.common.Error
}
var file_pkg_cache_proto_responses_set_multiple_response_proto_depIdxs = []int32{
	1, // 0: gcommon.v1.cache.SetMultipleResponse.error:type_name -> gcommon.v1.common.Error
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_cache_proto_responses_set_multiple_response_proto_init() }
func file_pkg_cache_proto_responses_set_multiple_response_proto_init() {
	if File_pkg_cache_proto_responses_set_multiple_response_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pkg_cache_proto_responses_set_multiple_response_proto_rawDesc), len(file_pkg_cache_proto_responses_set_multiple_response_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_cache_proto_responses_set_multiple_response_proto_goTypes,
		DependencyIndexes: file_pkg_cache_proto_responses_set_multiple_response_proto_depIdxs,
		MessageInfos:      file_pkg_cache_proto_responses_set_multiple_response_proto_msgTypes,
	}.Build()
	File_pkg_cache_proto_responses_set_multiple_response_proto = out.File
	file_pkg_cache_proto_responses_set_multiple_response_proto_goTypes = nil
	file_pkg_cache_proto_responses_set_multiple_response_proto_depIdxs = nil
}

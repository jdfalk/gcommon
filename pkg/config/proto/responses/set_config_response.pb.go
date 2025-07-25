// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: pkg/config/proto/responses/set_config_response.proto

//go:build !protoopaque

package configpb

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
// SetConfigResponse indicates the result of a configuration set operation.
type SetConfigResponse struct {
	state protoimpl.MessageState `protogen:"hybrid.v1"`
	// Whether the operation succeeded
	Success *bool `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	// Previous value if it existed
	PreviousEntry *ConfigEntry `protobuf:"bytes,2,opt,name=previous_entry,json=previousEntry" json:"previous_entry,omitempty"`
	// Error information
	Error         *proto.Error `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetConfigResponse) Reset() {
	*x = SetConfigResponse{}
	mi := &file_pkg_config_proto_responses_set_config_response_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetConfigResponse) ProtoMessage() {}

func (x *SetConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_config_proto_responses_set_config_response_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *SetConfigResponse) GetSuccess() bool {
	if x != nil && x.Success != nil {
		return *x.Success
	}
	return false
}

func (x *SetConfigResponse) GetPreviousEntry() *ConfigEntry {
	if x != nil {
		return x.PreviousEntry
	}
	return nil
}

func (x *SetConfigResponse) GetError() *proto.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *SetConfigResponse) SetSuccess(v bool) {
	x.Success = &v
}

func (x *SetConfigResponse) SetPreviousEntry(v *ConfigEntry) {
	x.PreviousEntry = v
}

func (x *SetConfigResponse) SetError(v *proto.Error) {
	x.Error = v
}

func (x *SetConfigResponse) HasSuccess() bool {
	if x == nil {
		return false
	}
	return x.Success != nil
}

func (x *SetConfigResponse) HasPreviousEntry() bool {
	if x == nil {
		return false
	}
	return x.PreviousEntry != nil
}

func (x *SetConfigResponse) HasError() bool {
	if x == nil {
		return false
	}
	return x.Error != nil
}

func (x *SetConfigResponse) ClearSuccess() {
	x.Success = nil
}

func (x *SetConfigResponse) ClearPreviousEntry() {
	x.PreviousEntry = nil
}

func (x *SetConfigResponse) ClearError() {
	x.Error = nil
}

type SetConfigResponse_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// Whether the operation succeeded
	Success *bool
	// Previous value if it existed
	PreviousEntry *ConfigEntry
	// Error information
	Error *proto.Error
}

func (b0 SetConfigResponse_builder) Build() *SetConfigResponse {
	m0 := &SetConfigResponse{}
	b, x := &b0, m0
	_, _ = b, x
	x.Success = b.Success
	x.PreviousEntry = b.PreviousEntry
	x.Error = b.Error
	return m0
}

var File_pkg_config_proto_responses_set_config_response_proto protoreflect.FileDescriptor

const file_pkg_config_proto_responses_set_config_response_proto_rawDesc = "" +
	"\n" +
	"4pkg/config/proto/responses/set_config_response.proto\x12\x11gcommon.v1.config\x1a!google/protobuf/go_features.proto\x1a%pkg/common/proto/messages/error.proto\x1a,pkg/config/proto/messages/config_entry.proto\"\xa4\x01\n" +
	"\x11SetConfigResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\x12E\n" +
	"\x0eprevious_entry\x18\x02 \x01(\v2\x1e.gcommon.v1.config.ConfigEntryR\rpreviousEntry\x12.\n" +
	"\x05error\x18\x03 \x01(\v2\x18.gcommon.v1.common.ErrorR\x05errorB\xd2\x01\n" +
	"\x15com.gcommon.v1.configB\x16SetConfigResponseProtoP\x01Z3github.com/jdfalk/gcommon/pkg/config/proto;configpb\xa2\x02\x03GVC\xaa\x02\x11Gcommon.V1.Config\xca\x02\x11Gcommon\\V1\\Config\xe2\x02\x1dGcommon\\V1\\Config\\GPBMetadata\xea\x02\x13Gcommon::V1::Config\x92\x03\x05\xd2>\x02\x10\x02b\beditionsp\xe8\a"

var file_pkg_config_proto_responses_set_config_response_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pkg_config_proto_responses_set_config_response_proto_goTypes = []any{
	(*SetConfigResponse)(nil), // 0: gcommon.v1.config.SetConfigResponse
	(*ConfigEntry)(nil),       // 1: gcommon.v1.config.ConfigEntry
	(*proto.Error)(nil),       // 2: gcommon.v1.common.Error
}
var file_pkg_config_proto_responses_set_config_response_proto_depIdxs = []int32{
	1, // 0: gcommon.v1.config.SetConfigResponse.previous_entry:type_name -> gcommon.v1.config.ConfigEntry
	2, // 1: gcommon.v1.config.SetConfigResponse.error:type_name -> gcommon.v1.common.Error
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_config_proto_responses_set_config_response_proto_init() }
func file_pkg_config_proto_responses_set_config_response_proto_init() {
	if File_pkg_config_proto_responses_set_config_response_proto != nil {
		return
	}
	file_pkg_config_proto_messages_config_entry_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pkg_config_proto_responses_set_config_response_proto_rawDesc), len(file_pkg_config_proto_responses_set_config_response_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_config_proto_responses_set_config_response_proto_goTypes,
		DependencyIndexes: file_pkg_config_proto_responses_set_config_response_proto_depIdxs,
		MessageInfos:      file_pkg_config_proto_responses_set_config_response_proto_msgTypes,
	}.Build()
	File_pkg_config_proto_responses_set_config_response_proto = out.File
	file_pkg_config_proto_responses_set_config_response_proto_goTypes = nil
	file_pkg_config_proto_responses_set_config_response_proto_depIdxs = nil
}

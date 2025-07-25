// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: pkg/config/proto/requests/get_config_request.proto

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
// GetConfigRequest retrieves a single configuration value.
type GetConfigRequest struct {
	state protoimpl.MessageState `protogen:"hybrid.v1"`
	// Configuration key
	Key *string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	// Optional namespace/environment
	Namespace *string `protobuf:"bytes,2,opt,name=namespace" json:"namespace,omitempty"`
	// Request metadata
	Metadata *proto.RequestMetadata `protobuf:"bytes,3,opt,name=metadata" json:"metadata,omitempty"`
	// Whether to decrypt encrypted values
	Decrypt       *bool `protobuf:"varint,4,opt,name=decrypt" json:"decrypt,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetConfigRequest) Reset() {
	*x = GetConfigRequest{}
	mi := &file_pkg_config_proto_requests_get_config_request_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfigRequest) ProtoMessage() {}

func (x *GetConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_config_proto_requests_get_config_request_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *GetConfigRequest) GetKey() string {
	if x != nil && x.Key != nil {
		return *x.Key
	}
	return ""
}

func (x *GetConfigRequest) GetNamespace() string {
	if x != nil && x.Namespace != nil {
		return *x.Namespace
	}
	return ""
}

func (x *GetConfigRequest) GetMetadata() *proto.RequestMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *GetConfigRequest) GetDecrypt() bool {
	if x != nil && x.Decrypt != nil {
		return *x.Decrypt
	}
	return false
}

func (x *GetConfigRequest) SetKey(v string) {
	x.Key = &v
}

func (x *GetConfigRequest) SetNamespace(v string) {
	x.Namespace = &v
}

func (x *GetConfigRequest) SetMetadata(v *proto.RequestMetadata) {
	x.Metadata = v
}

func (x *GetConfigRequest) SetDecrypt(v bool) {
	x.Decrypt = &v
}

func (x *GetConfigRequest) HasKey() bool {
	if x == nil {
		return false
	}
	return x.Key != nil
}

func (x *GetConfigRequest) HasNamespace() bool {
	if x == nil {
		return false
	}
	return x.Namespace != nil
}

func (x *GetConfigRequest) HasMetadata() bool {
	if x == nil {
		return false
	}
	return x.Metadata != nil
}

func (x *GetConfigRequest) HasDecrypt() bool {
	if x == nil {
		return false
	}
	return x.Decrypt != nil
}

func (x *GetConfigRequest) ClearKey() {
	x.Key = nil
}

func (x *GetConfigRequest) ClearNamespace() {
	x.Namespace = nil
}

func (x *GetConfigRequest) ClearMetadata() {
	x.Metadata = nil
}

func (x *GetConfigRequest) ClearDecrypt() {
	x.Decrypt = nil
}

type GetConfigRequest_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// Configuration key
	Key *string
	// Optional namespace/environment
	Namespace *string
	// Request metadata
	Metadata *proto.RequestMetadata
	// Whether to decrypt encrypted values
	Decrypt *bool
}

func (b0 GetConfigRequest_builder) Build() *GetConfigRequest {
	m0 := &GetConfigRequest{}
	b, x := &b0, m0
	_, _ = b, x
	x.Key = b.Key
	x.Namespace = b.Namespace
	x.Metadata = b.Metadata
	x.Decrypt = b.Decrypt
	return m0
}

var File_pkg_config_proto_requests_get_config_request_proto protoreflect.FileDescriptor

const file_pkg_config_proto_requests_get_config_request_proto_rawDesc = "" +
	"\n" +
	"2pkg/config/proto/requests/get_config_request.proto\x12\x11gcommon.v1.config\x1a!google/protobuf/go_features.proto\x1a0pkg/common/proto/messages/request_metadata.proto\"\x9c\x01\n" +
	"\x10GetConfigRequest\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x1c\n" +
	"\tnamespace\x18\x02 \x01(\tR\tnamespace\x12>\n" +
	"\bmetadata\x18\x03 \x01(\v2\".gcommon.v1.common.RequestMetadataR\bmetadata\x12\x18\n" +
	"\adecrypt\x18\x04 \x01(\bR\adecryptB\xd1\x01\n" +
	"\x15com.gcommon.v1.configB\x15GetConfigRequestProtoP\x01Z3github.com/jdfalk/gcommon/pkg/config/proto;configpb\xa2\x02\x03GVC\xaa\x02\x11Gcommon.V1.Config\xca\x02\x11Gcommon\\V1\\Config\xe2\x02\x1dGcommon\\V1\\Config\\GPBMetadata\xea\x02\x13Gcommon::V1::Config\x92\x03\x05\xd2>\x02\x10\x02b\beditionsp\xe8\a"

var file_pkg_config_proto_requests_get_config_request_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pkg_config_proto_requests_get_config_request_proto_goTypes = []any{
	(*GetConfigRequest)(nil),      // 0: gcommon.v1.config.GetConfigRequest
	(*proto.RequestMetadata)(nil), // 1: gcommon.v1.common.RequestMetadata
}
var file_pkg_config_proto_requests_get_config_request_proto_depIdxs = []int32{
	1, // 0: gcommon.v1.config.GetConfigRequest.metadata:type_name -> gcommon.v1.common.RequestMetadata
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_config_proto_requests_get_config_request_proto_init() }
func file_pkg_config_proto_requests_get_config_request_proto_init() {
	if File_pkg_config_proto_requests_get_config_request_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pkg_config_proto_requests_get_config_request_proto_rawDesc), len(file_pkg_config_proto_requests_get_config_request_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_config_proto_requests_get_config_request_proto_goTypes,
		DependencyIndexes: file_pkg_config_proto_requests_get_config_request_proto_depIdxs,
		MessageInfos:      file_pkg_config_proto_requests_get_config_request_proto_msgTypes,
	}.Build()
	File_pkg_config_proto_requests_get_config_request_proto = out.File
	file_pkg_config_proto_requests_get_config_request_proto_goTypes = nil
	file_pkg_config_proto_requests_get_config_request_proto_depIdxs = nil
}

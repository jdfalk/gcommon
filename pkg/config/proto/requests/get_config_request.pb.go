// file: pkg/config/proto/requests/get_config_request.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
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
// Request to retrieve a configuration value.
// Supports hierarchical key lookup and environment-specific overrides.
type GetConfigRequest struct {
	state protoimpl.MessageState `protogen:"hybrid.v1"`
	// Configuration key (supports dot notation for hierarchy)
	Key *string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	// Environment context (e.g., "dev", "staging", "prod")
	Environment *string `protobuf:"bytes,2,opt,name=environment" json:"environment,omitempty"`
	// Service/module context for scoped configuration
	Service *string `protobuf:"bytes,3,opt,name=service" json:"service,omitempty"`
	// Request metadata for tracing and correlation
	Metadata *proto.RequestMetadata `protobuf:"bytes,4,opt,name=metadata" json:"metadata,omitempty"`
	// Whether to include metadata in response
	IncludeMetadata *bool `protobuf:"varint,5,opt,name=include_metadata,json=includeMetadata" json:"include_metadata,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
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

func (x *GetConfigRequest) GetEnvironment() string {
	if x != nil && x.Environment != nil {
		return *x.Environment
	}
	return ""
}

func (x *GetConfigRequest) GetService() string {
	if x != nil && x.Service != nil {
		return *x.Service
	}
	return ""
}

func (x *GetConfigRequest) GetMetadata() *proto.RequestMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *GetConfigRequest) GetIncludeMetadata() bool {
	if x != nil && x.IncludeMetadata != nil {
		return *x.IncludeMetadata
	}
	return false
}

func (x *GetConfigRequest) SetKey(v string) {
	x.Key = &v
}

func (x *GetConfigRequest) SetEnvironment(v string) {
	x.Environment = &v
}

func (x *GetConfigRequest) SetService(v string) {
	x.Service = &v
}

func (x *GetConfigRequest) SetMetadata(v *proto.RequestMetadata) {
	x.Metadata = v
}

func (x *GetConfigRequest) SetIncludeMetadata(v bool) {
	x.IncludeMetadata = &v
}

func (x *GetConfigRequest) HasKey() bool {
	if x == nil {
		return false
	}
	return x.Key != nil
}

func (x *GetConfigRequest) HasEnvironment() bool {
	if x == nil {
		return false
	}
	return x.Environment != nil
}

func (x *GetConfigRequest) HasService() bool {
	if x == nil {
		return false
	}
	return x.Service != nil
}

func (x *GetConfigRequest) HasMetadata() bool {
	if x == nil {
		return false
	}
	return x.Metadata != nil
}

func (x *GetConfigRequest) HasIncludeMetadata() bool {
	if x == nil {
		return false
	}
	return x.IncludeMetadata != nil
}

func (x *GetConfigRequest) ClearKey() {
	x.Key = nil
}

func (x *GetConfigRequest) ClearEnvironment() {
	x.Environment = nil
}

func (x *GetConfigRequest) ClearService() {
	x.Service = nil
}

func (x *GetConfigRequest) ClearMetadata() {
	x.Metadata = nil
}

func (x *GetConfigRequest) ClearIncludeMetadata() {
	x.IncludeMetadata = nil
}

type GetConfigRequest_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// Configuration key (supports dot notation for hierarchy)
	Key *string
	// Environment context (e.g., "dev", "staging", "prod")
	Environment *string
	// Service/module context for scoped configuration
	Service *string
	// Request metadata for tracing and correlation
	Metadata *proto.RequestMetadata
	// Whether to include metadata in response
	IncludeMetadata *bool
}

func (b0 GetConfigRequest_builder) Build() *GetConfigRequest {
	m0 := &GetConfigRequest{}
	b, x := &b0, m0
	_, _ = b, x
	x.Key = b.Key
	x.Environment = b.Environment
	x.Service = b.Service
	x.Metadata = b.Metadata
	x.IncludeMetadata = b.IncludeMetadata
	return m0
}

var File_pkg_config_proto_requests_get_config_request_proto protoreflect.FileDescriptor

const file_pkg_config_proto_requests_get_config_request_proto_rawDesc = "" +
	"\n" +
	"2pkg/config/proto/requests/get_config_request.proto\x12\x11gcommon.v1.config\x1a!google/protobuf/go_features.proto\x1a0pkg/common/proto/messages/request_metadata.proto\"\xcf\x01\n" +
	"\x10GetConfigRequest\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12 \n" +
	"\venvironment\x18\x02 \x01(\tR\venvironment\x12\x18\n" +
	"\aservice\x18\x03 \x01(\tR\aservice\x12B\n" +
	"\bmetadata\x18\x04 \x01(\v2\".gcommon.v1.common.RequestMetadataB\x02(\x01R\bmetadata\x12)\n" +
	"\x10include_metadata\x18\x05 \x01(\bR\x0fincludeMetadataB=Z3github.com/jdfalk/gcommon/pkg/config/proto;configpb\x92\x03\x05\xd2>\x02\x10\x02b\beditionsp\xe8\a"

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

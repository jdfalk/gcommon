// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: pkg/common/proto/types/resource_reference.proto

//go:build protoopaque

package commonpb

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
// Resource reference for cross-module operations and relationships.
// Provides a standardized way to reference resources across different
// GCommon modules with consistent identification and ownership tracking.
type ResourceReference struct {
	state                  protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Type        *string                `protobuf:"bytes,1,opt,name=type"`
	xxx_hidden_Id          *string                `protobuf:"bytes,2,opt,name=id"`
	xxx_hidden_Name        *string                `protobuf:"bytes,3,opt,name=name"`
	xxx_hidden_Module      *string                `protobuf:"bytes,4,opt,name=module"`
	XXX_raceDetectHookData protoimpl.RaceDetectHookData
	XXX_presence           [1]uint32
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *ResourceReference) Reset() {
	*x = ResourceReference{}
	mi := &file_pkg_common_proto_types_resource_reference_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResourceReference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceReference) ProtoMessage() {}

func (x *ResourceReference) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_common_proto_types_resource_reference_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *ResourceReference) GetType() string {
	if x != nil {
		if x.xxx_hidden_Type != nil {
			return *x.xxx_hidden_Type
		}
		return ""
	}
	return ""
}

func (x *ResourceReference) GetId() string {
	if x != nil {
		if x.xxx_hidden_Id != nil {
			return *x.xxx_hidden_Id
		}
		return ""
	}
	return ""
}

func (x *ResourceReference) GetName() string {
	if x != nil {
		if x.xxx_hidden_Name != nil {
			return *x.xxx_hidden_Name
		}
		return ""
	}
	return ""
}

func (x *ResourceReference) GetModule() string {
	if x != nil {
		if x.xxx_hidden_Module != nil {
			return *x.xxx_hidden_Module
		}
		return ""
	}
	return ""
}

func (x *ResourceReference) SetType(v string) {
	x.xxx_hidden_Type = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 4)
}

func (x *ResourceReference) SetId(v string) {
	x.xxx_hidden_Id = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 1, 4)
}

func (x *ResourceReference) SetName(v string) {
	x.xxx_hidden_Name = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 2, 4)
}

func (x *ResourceReference) SetModule(v string) {
	x.xxx_hidden_Module = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 3, 4)
}

func (x *ResourceReference) HasType() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *ResourceReference) HasId() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 1)
}

func (x *ResourceReference) HasName() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 2)
}

func (x *ResourceReference) HasModule() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 3)
}

func (x *ResourceReference) ClearType() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_Type = nil
}

func (x *ResourceReference) ClearId() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 1)
	x.xxx_hidden_Id = nil
}

func (x *ResourceReference) ClearName() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 2)
	x.xxx_hidden_Name = nil
}

func (x *ResourceReference) ClearModule() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 3)
	x.xxx_hidden_Module = nil
}

type ResourceReference_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// Resource type identifier (e.g., "user", "config", "queue", "metric")
	Type *string
	// Unique resource identifier within the module
	Id *string
	// Human-readable resource name for display purposes
	Name *string
	// Module that owns and manages this resource
	Module *string
}

func (b0 ResourceReference_builder) Build() *ResourceReference {
	m0 := &ResourceReference{}
	b, x := &b0, m0
	_, _ = b, x
	if b.Type != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 4)
		x.xxx_hidden_Type = b.Type
	}
	if b.Id != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 1, 4)
		x.xxx_hidden_Id = b.Id
	}
	if b.Name != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 2, 4)
		x.xxx_hidden_Name = b.Name
	}
	if b.Module != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 3, 4)
		x.xxx_hidden_Module = b.Module
	}
	return m0
}

var File_pkg_common_proto_types_resource_reference_proto protoreflect.FileDescriptor

const file_pkg_common_proto_types_resource_reference_proto_rawDesc = "" +
	"\n" +
	"/pkg/common/proto/types/resource_reference.proto\x12\x11gcommon.v1.common\x1a!google/protobuf/go_features.proto\"c\n" +
	"\x11ResourceReference\x12\x12\n" +
	"\x04type\x18\x01 \x01(\tR\x04type\x12\x0e\n" +
	"\x02id\x18\x02 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x03 \x01(\tR\x04name\x12\x16\n" +
	"\x06module\x18\x04 \x01(\tR\x06moduleB\xd2\x01\n" +
	"\x15com.gcommon.v1.commonB\x16ResourceReferenceProtoP\x01Z3github.com/jdfalk/gcommon/pkg/common/proto;commonpb\xa2\x02\x03GVC\xaa\x02\x11Gcommon.V1.Common\xca\x02\x11Gcommon\\V1\\Common\xe2\x02\x1dGcommon\\V1\\Common\\GPBMetadata\xea\x02\x13Gcommon::V1::Common\x92\x03\x05\xd2>\x02\x10\x02b\beditionsp\xe8\a"

var file_pkg_common_proto_types_resource_reference_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pkg_common_proto_types_resource_reference_proto_goTypes = []any{
	(*ResourceReference)(nil), // 0: gcommon.v1.common.ResourceReference
}
var file_pkg_common_proto_types_resource_reference_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_common_proto_types_resource_reference_proto_init() }
func file_pkg_common_proto_types_resource_reference_proto_init() {
	if File_pkg_common_proto_types_resource_reference_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pkg_common_proto_types_resource_reference_proto_rawDesc), len(file_pkg_common_proto_types_resource_reference_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_common_proto_types_resource_reference_proto_goTypes,
		DependencyIndexes: file_pkg_common_proto_types_resource_reference_proto_depIdxs,
		MessageInfos:      file_pkg_common_proto_types_resource_reference_proto_msgTypes,
	}.Build()
	File_pkg_common_proto_types_resource_reference_proto = out.File
	file_pkg_common_proto_types_resource_reference_proto_goTypes = nil
	file_pkg_common_proto_types_resource_reference_proto_depIdxs = nil
}

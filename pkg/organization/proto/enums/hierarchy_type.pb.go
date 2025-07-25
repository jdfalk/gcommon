// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: pkg/organization/proto/enums/hierarchy_type.proto

//go:build !protoopaque

package organizationpb

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
// Hierarchy type enumeration defining the type of organizational structure.
// Used to categorize different organizational units and their relationships.
type HierarchyType int32

const (
	// Default value indicating no hierarchy type was specified
	HierarchyType_HIERARCHY_TYPE_UNSPECIFIED HierarchyType = 0
	// Department-based hierarchical structure
	HierarchyType_HIERARCHY_TYPE_DEPARTMENT HierarchyType = 1
	// Team-based organizational structure
	HierarchyType_HIERARCHY_TYPE_TEAM HierarchyType = 2
	// Project-based organizational structure
	HierarchyType_HIERARCHY_TYPE_PROJECT HierarchyType = 3
	// Geographic/location-based structure
	HierarchyType_HIERARCHY_TYPE_GEOGRAPHIC HierarchyType = 4
	// Functional role-based structure
	HierarchyType_HIERARCHY_TYPE_FUNCTIONAL HierarchyType = 5
	// Matrix organizational structure (multi-dimensional)
	HierarchyType_HIERARCHY_TYPE_MATRIX HierarchyType = 6
	// Flat organizational structure (minimal hierarchy)
	HierarchyType_HIERARCHY_TYPE_FLAT HierarchyType = 7
)

// Enum value maps for HierarchyType.
var (
	HierarchyType_name = map[int32]string{
		0: "HIERARCHY_TYPE_UNSPECIFIED",
		1: "HIERARCHY_TYPE_DEPARTMENT",
		2: "HIERARCHY_TYPE_TEAM",
		3: "HIERARCHY_TYPE_PROJECT",
		4: "HIERARCHY_TYPE_GEOGRAPHIC",
		5: "HIERARCHY_TYPE_FUNCTIONAL",
		6: "HIERARCHY_TYPE_MATRIX",
		7: "HIERARCHY_TYPE_FLAT",
	}
	HierarchyType_value = map[string]int32{
		"HIERARCHY_TYPE_UNSPECIFIED": 0,
		"HIERARCHY_TYPE_DEPARTMENT":  1,
		"HIERARCHY_TYPE_TEAM":        2,
		"HIERARCHY_TYPE_PROJECT":     3,
		"HIERARCHY_TYPE_GEOGRAPHIC":  4,
		"HIERARCHY_TYPE_FUNCTIONAL":  5,
		"HIERARCHY_TYPE_MATRIX":      6,
		"HIERARCHY_TYPE_FLAT":        7,
	}
)

func (x HierarchyType) Enum() *HierarchyType {
	p := new(HierarchyType)
	*p = x
	return p
}

func (x HierarchyType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HierarchyType) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_organization_proto_enums_hierarchy_type_proto_enumTypes[0].Descriptor()
}

func (HierarchyType) Type() protoreflect.EnumType {
	return &file_pkg_organization_proto_enums_hierarchy_type_proto_enumTypes[0]
}

func (x HierarchyType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

var File_pkg_organization_proto_enums_hierarchy_type_proto protoreflect.FileDescriptor

const file_pkg_organization_proto_enums_hierarchy_type_proto_rawDesc = "" +
	"\n" +
	"1pkg/organization/proto/enums/hierarchy_type.proto\x12\x17gcommon.v1.organization\x1a!google/protobuf/go_features.proto*\xf5\x01\n" +
	"\rHierarchyType\x12\x1e\n" +
	"\x1aHIERARCHY_TYPE_UNSPECIFIED\x10\x00\x12\x1d\n" +
	"\x19HIERARCHY_TYPE_DEPARTMENT\x10\x01\x12\x17\n" +
	"\x13HIERARCHY_TYPE_TEAM\x10\x02\x12\x1a\n" +
	"\x16HIERARCHY_TYPE_PROJECT\x10\x03\x12\x1d\n" +
	"\x19HIERARCHY_TYPE_GEOGRAPHIC\x10\x04\x12\x1d\n" +
	"\x19HIERARCHY_TYPE_FUNCTIONAL\x10\x05\x12\x19\n" +
	"\x15HIERARCHY_TYPE_MATRIX\x10\x06\x12\x17\n" +
	"\x13HIERARCHY_TYPE_FLAT\x10\aB\xf8\x01\n" +
	"\x1bcom.gcommon.v1.organizationB\x12HierarchyTypeProtoP\x01Z?github.com/jdfalk/gcommon/pkg/organization/proto;organizationpb\xa2\x02\x03GVO\xaa\x02\x17Gcommon.V1.Organization\xca\x02\x17Gcommon\\V1\\Organization\xe2\x02#Gcommon\\V1\\Organization\\GPBMetadata\xea\x02\x19Gcommon::V1::Organization\x92\x03\x05\xd2>\x02\x10\x02b\beditionsp\xe8\a"

var file_pkg_organization_proto_enums_hierarchy_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkg_organization_proto_enums_hierarchy_type_proto_goTypes = []any{
	(HierarchyType)(0), // 0: gcommon.v1.organization.HierarchyType
}
var file_pkg_organization_proto_enums_hierarchy_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_organization_proto_enums_hierarchy_type_proto_init() }
func file_pkg_organization_proto_enums_hierarchy_type_proto_init() {
	if File_pkg_organization_proto_enums_hierarchy_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pkg_organization_proto_enums_hierarchy_type_proto_rawDesc), len(file_pkg_organization_proto_enums_hierarchy_type_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_organization_proto_enums_hierarchy_type_proto_goTypes,
		DependencyIndexes: file_pkg_organization_proto_enums_hierarchy_type_proto_depIdxs,
		EnumInfos:         file_pkg_organization_proto_enums_hierarchy_type_proto_enumTypes,
	}.Build()
	File_pkg_organization_proto_enums_hierarchy_type_proto = out.File
	file_pkg_organization_proto_enums_hierarchy_type_proto_goTypes = nil
	file_pkg_organization_proto_enums_hierarchy_type_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: pkg/organization/proto/enums/organization_status.proto

//go:build protoopaque

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
// Organization status enumeration defining the state of an organization.
// Used to track organization lifecycle, operational status, and access permissions.
type OrganizationStatus int32

const (
	// Default value indicating no status was specified
	OrganizationStatus_ORGANIZATION_STATUS_UNSPECIFIED OrganizationStatus = 0
	// Organization is active and operational
	OrganizationStatus_ORGANIZATION_STATUS_ACTIVE OrganizationStatus = 1
	// Organization is inactive (temporarily suspended operations)
	OrganizationStatus_ORGANIZATION_STATUS_INACTIVE OrganizationStatus = 2
	// Organization is suspended due to policy violations or billing issues
	OrganizationStatus_ORGANIZATION_STATUS_SUSPENDED OrganizationStatus = 3
	// Organization is pending verification or onboarding completion
	OrganizationStatus_ORGANIZATION_STATUS_PENDING OrganizationStatus = 4
	// Organization is archived (read-only access, no new operations)
	OrganizationStatus_ORGANIZATION_STATUS_ARCHIVED OrganizationStatus = 5
	// Organization is marked for deletion and undergoing cleanup
	OrganizationStatus_ORGANIZATION_STATUS_DELETED OrganizationStatus = 6
)

// Enum value maps for OrganizationStatus.
var (
	OrganizationStatus_name = map[int32]string{
		0: "ORGANIZATION_STATUS_UNSPECIFIED",
		1: "ORGANIZATION_STATUS_ACTIVE",
		2: "ORGANIZATION_STATUS_INACTIVE",
		3: "ORGANIZATION_STATUS_SUSPENDED",
		4: "ORGANIZATION_STATUS_PENDING",
		5: "ORGANIZATION_STATUS_ARCHIVED",
		6: "ORGANIZATION_STATUS_DELETED",
	}
	OrganizationStatus_value = map[string]int32{
		"ORGANIZATION_STATUS_UNSPECIFIED": 0,
		"ORGANIZATION_STATUS_ACTIVE":      1,
		"ORGANIZATION_STATUS_INACTIVE":    2,
		"ORGANIZATION_STATUS_SUSPENDED":   3,
		"ORGANIZATION_STATUS_PENDING":     4,
		"ORGANIZATION_STATUS_ARCHIVED":    5,
		"ORGANIZATION_STATUS_DELETED":     6,
	}
)

func (x OrganizationStatus) Enum() *OrganizationStatus {
	p := new(OrganizationStatus)
	*p = x
	return p
}

func (x OrganizationStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrganizationStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_organization_proto_enums_organization_status_proto_enumTypes[0].Descriptor()
}

func (OrganizationStatus) Type() protoreflect.EnumType {
	return &file_pkg_organization_proto_enums_organization_status_proto_enumTypes[0]
}

func (x OrganizationStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

var File_pkg_organization_proto_enums_organization_status_proto protoreflect.FileDescriptor

const file_pkg_organization_proto_enums_organization_status_proto_rawDesc = "" +
	"\n" +
	"6pkg/organization/proto/enums/organization_status.proto\x12\x17gcommon.v1.organization\x1a!google/protobuf/go_features.proto*\x82\x02\n" +
	"\x12OrganizationStatus\x12#\n" +
	"\x1fORGANIZATION_STATUS_UNSPECIFIED\x10\x00\x12\x1e\n" +
	"\x1aORGANIZATION_STATUS_ACTIVE\x10\x01\x12 \n" +
	"\x1cORGANIZATION_STATUS_INACTIVE\x10\x02\x12!\n" +
	"\x1dORGANIZATION_STATUS_SUSPENDED\x10\x03\x12\x1f\n" +
	"\x1bORGANIZATION_STATUS_PENDING\x10\x04\x12 \n" +
	"\x1cORGANIZATION_STATUS_ARCHIVED\x10\x05\x12\x1f\n" +
	"\x1bORGANIZATION_STATUS_DELETED\x10\x06B\xfd\x01\n" +
	"\x1bcom.gcommon.v1.organizationB\x17OrganizationStatusProtoP\x01Z?github.com/jdfalk/gcommon/pkg/organization/proto;organizationpb\xa2\x02\x03GVO\xaa\x02\x17Gcommon.V1.Organization\xca\x02\x17Gcommon\\V1\\Organization\xe2\x02#Gcommon\\V1\\Organization\\GPBMetadata\xea\x02\x19Gcommon::V1::Organization\x92\x03\x05\xd2>\x02\x10\x02b\beditionsp\xe8\a"

var file_pkg_organization_proto_enums_organization_status_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkg_organization_proto_enums_organization_status_proto_goTypes = []any{
	(OrganizationStatus)(0), // 0: gcommon.v1.organization.OrganizationStatus
}
var file_pkg_organization_proto_enums_organization_status_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_organization_proto_enums_organization_status_proto_init() }
func file_pkg_organization_proto_enums_organization_status_proto_init() {
	if File_pkg_organization_proto_enums_organization_status_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pkg_organization_proto_enums_organization_status_proto_rawDesc), len(file_pkg_organization_proto_enums_organization_status_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_organization_proto_enums_organization_status_proto_goTypes,
		DependencyIndexes: file_pkg_organization_proto_enums_organization_status_proto_depIdxs,
		EnumInfos:         file_pkg_organization_proto_enums_organization_status_proto_enumTypes,
	}.Build()
	File_pkg_organization_proto_enums_organization_status_proto = out.File
	file_pkg_organization_proto_enums_organization_status_proto_goTypes = nil
	file_pkg_organization_proto_enums_organization_status_proto_depIdxs = nil
}

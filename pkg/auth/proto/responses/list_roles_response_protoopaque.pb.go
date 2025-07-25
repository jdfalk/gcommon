// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: pkg/auth/proto/responses/list_roles_response.proto

//go:build protoopaque

package authpb

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
// Response containing a list of roles.
// Includes pagination information for large result sets.
// Used for role management interfaces and administration.
type ListRolesResponse struct {
	state                 protoimpl.MessageState   `protogen:"opaque.v1"`
	xxx_hidden_Roles      *[]*Role                 `protobuf:"bytes,1,rep,name=roles"`
	xxx_hidden_Pagination *proto.PaginatedResponse `protobuf:"bytes,2,opt,name=pagination"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *ListRolesResponse) Reset() {
	*x = ListRolesResponse{}
	mi := &file_pkg_auth_proto_responses_list_roles_response_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRolesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRolesResponse) ProtoMessage() {}

func (x *ListRolesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_auth_proto_responses_list_roles_response_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *ListRolesResponse) GetRoles() []*Role {
	if x != nil {
		if x.xxx_hidden_Roles != nil {
			return *x.xxx_hidden_Roles
		}
	}
	return nil
}

func (x *ListRolesResponse) GetPagination() *proto.PaginatedResponse {
	if x != nil {
		return x.xxx_hidden_Pagination
	}
	return nil
}

func (x *ListRolesResponse) SetRoles(v []*Role) {
	x.xxx_hidden_Roles = &v
}

func (x *ListRolesResponse) SetPagination(v *proto.PaginatedResponse) {
	x.xxx_hidden_Pagination = v
}

func (x *ListRolesResponse) HasPagination() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Pagination != nil
}

func (x *ListRolesResponse) ClearPagination() {
	x.xxx_hidden_Pagination = nil
}

type ListRolesResponse_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// List of roles matching the request criteria
	Roles []*Role
	// Pagination information for the response
	Pagination *proto.PaginatedResponse
}

func (b0 ListRolesResponse_builder) Build() *ListRolesResponse {
	m0 := &ListRolesResponse{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Roles = &b.Roles
	x.xxx_hidden_Pagination = b.Pagination
	return m0
}

var File_pkg_auth_proto_responses_list_roles_response_proto protoreflect.FileDescriptor

const file_pkg_auth_proto_responses_list_roles_response_proto_rawDesc = "" +
	"\n" +
	"2pkg/auth/proto/responses/list_roles_response.proto\x12\x0fgcommon.v1.auth\x1a!google/protobuf/go_features.proto\x1a2pkg/common/proto/messages/paginated_response.proto\x1a\x1fpkg/auth/proto/types/role.proto\"\x86\x01\n" +
	"\x11ListRolesResponse\x12+\n" +
	"\x05roles\x18\x01 \x03(\v2\x15.gcommon.v1.auth.RoleR\x05roles\x12D\n" +
	"\n" +
	"pagination\x18\x02 \x01(\v2$.gcommon.v1.common.PaginatedResponseR\n" +
	"paginationB\xc4\x01\n" +
	"\x13com.gcommon.v1.authB\x16ListRolesResponseProtoP\x01Z/github.com/jdfalk/gcommon/pkg/auth/proto;authpb\xa2\x02\x03GVA\xaa\x02\x0fGcommon.V1.Auth\xca\x02\x0fGcommon\\V1\\Auth\xe2\x02\x1bGcommon\\V1\\Auth\\GPBMetadata\xea\x02\x11Gcommon::V1::Auth\x92\x03\x05\xd2>\x02\x10\x02b\beditionsp\xe8\a"

var file_pkg_auth_proto_responses_list_roles_response_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pkg_auth_proto_responses_list_roles_response_proto_goTypes = []any{
	(*ListRolesResponse)(nil),       // 0: gcommon.v1.auth.ListRolesResponse
	(*Role)(nil),                    // 1: gcommon.v1.auth.Role
	(*proto.PaginatedResponse)(nil), // 2: gcommon.v1.common.PaginatedResponse
}
var file_pkg_auth_proto_responses_list_roles_response_proto_depIdxs = []int32{
	1, // 0: gcommon.v1.auth.ListRolesResponse.roles:type_name -> gcommon.v1.auth.Role
	2, // 1: gcommon.v1.auth.ListRolesResponse.pagination:type_name -> gcommon.v1.common.PaginatedResponse
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_auth_proto_responses_list_roles_response_proto_init() }
func file_pkg_auth_proto_responses_list_roles_response_proto_init() {
	if File_pkg_auth_proto_responses_list_roles_response_proto != nil {
		return
	}
	file_pkg_auth_proto_types_role_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pkg_auth_proto_responses_list_roles_response_proto_rawDesc), len(file_pkg_auth_proto_responses_list_roles_response_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_auth_proto_responses_list_roles_response_proto_goTypes,
		DependencyIndexes: file_pkg_auth_proto_responses_list_roles_response_proto_depIdxs,
		MessageInfos:      file_pkg_auth_proto_responses_list_roles_response_proto_msgTypes,
	}.Build()
	File_pkg_auth_proto_responses_list_roles_response_proto = out.File
	file_pkg_auth_proto_responses_list_roles_response_proto_goTypes = nil
	file_pkg_auth_proto_responses_list_roles_response_proto_depIdxs = nil
}

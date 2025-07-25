// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: pkg/auth/proto/messages/role_assignment.proto

//go:build !protoopaque

package authpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/gofeaturespb"
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
// Role assignment entity representing role grants to users.
// Used for tracking role-based access control assignments.
// Supports scoped roles and expiration.
type RoleAssignment struct {
	state protoimpl.MessageState `protogen:"hybrid.v1"`
	// Unique assignment identifier
	Id *string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// User ID receiving the role
	UserId *string `protobuf:"bytes,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	// Role ID being assigned
	RoleId *string `protobuf:"bytes,3,opt,name=role_id,json=roleId" json:"role_id,omitempty"`
	// Resource the role applies to (optional, for scoped roles)
	Resource *string `protobuf:"bytes,4,opt,name=resource" json:"resource,omitempty"`
	// Role scope
	Scope *RoleScope `protobuf:"varint,5,opt,name=scope,enum=gcommon.v1.auth.RoleScope" json:"scope,omitempty"`
	// Assignment creation timestamp
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	// Assignment expiration timestamp (optional)
	ExpiresAt *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=expires_at,json=expiresAt" json:"expires_at,omitempty"`
	// User who assigned the role
	AssignedByUserId *string `protobuf:"bytes,8,opt,name=assigned_by_user_id,json=assignedByUserId" json:"assigned_by_user_id,omitempty"`
	// Assignment metadata
	Metadata map[string]string `protobuf:"bytes,9,rep,name=metadata" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Assignment active flag
	Active        *bool `protobuf:"varint,10,opt,name=active" json:"active,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RoleAssignment) Reset() {
	*x = RoleAssignment{}
	mi := &file_pkg_auth_proto_messages_role_assignment_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RoleAssignment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleAssignment) ProtoMessage() {}

func (x *RoleAssignment) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_auth_proto_messages_role_assignment_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *RoleAssignment) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *RoleAssignment) GetUserId() string {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return ""
}

func (x *RoleAssignment) GetRoleId() string {
	if x != nil && x.RoleId != nil {
		return *x.RoleId
	}
	return ""
}

func (x *RoleAssignment) GetResource() string {
	if x != nil && x.Resource != nil {
		return *x.Resource
	}
	return ""
}

func (x *RoleAssignment) GetScope() RoleScope {
	if x != nil && x.Scope != nil {
		return *x.Scope
	}
	return RoleScope_ROLE_SCOPE_UNSPECIFIED
}

func (x *RoleAssignment) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *RoleAssignment) GetExpiresAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiresAt
	}
	return nil
}

func (x *RoleAssignment) GetAssignedByUserId() string {
	if x != nil && x.AssignedByUserId != nil {
		return *x.AssignedByUserId
	}
	return ""
}

func (x *RoleAssignment) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *RoleAssignment) GetActive() bool {
	if x != nil && x.Active != nil {
		return *x.Active
	}
	return false
}

func (x *RoleAssignment) SetId(v string) {
	x.Id = &v
}

func (x *RoleAssignment) SetUserId(v string) {
	x.UserId = &v
}

func (x *RoleAssignment) SetRoleId(v string) {
	x.RoleId = &v
}

func (x *RoleAssignment) SetResource(v string) {
	x.Resource = &v
}

func (x *RoleAssignment) SetScope(v RoleScope) {
	x.Scope = &v
}

func (x *RoleAssignment) SetCreatedAt(v *timestamppb.Timestamp) {
	x.CreatedAt = v
}

func (x *RoleAssignment) SetExpiresAt(v *timestamppb.Timestamp) {
	x.ExpiresAt = v
}

func (x *RoleAssignment) SetAssignedByUserId(v string) {
	x.AssignedByUserId = &v
}

func (x *RoleAssignment) SetMetadata(v map[string]string) {
	x.Metadata = v
}

func (x *RoleAssignment) SetActive(v bool) {
	x.Active = &v
}

func (x *RoleAssignment) HasId() bool {
	if x == nil {
		return false
	}
	return x.Id != nil
}

func (x *RoleAssignment) HasUserId() bool {
	if x == nil {
		return false
	}
	return x.UserId != nil
}

func (x *RoleAssignment) HasRoleId() bool {
	if x == nil {
		return false
	}
	return x.RoleId != nil
}

func (x *RoleAssignment) HasResource() bool {
	if x == nil {
		return false
	}
	return x.Resource != nil
}

func (x *RoleAssignment) HasScope() bool {
	if x == nil {
		return false
	}
	return x.Scope != nil
}

func (x *RoleAssignment) HasCreatedAt() bool {
	if x == nil {
		return false
	}
	return x.CreatedAt != nil
}

func (x *RoleAssignment) HasExpiresAt() bool {
	if x == nil {
		return false
	}
	return x.ExpiresAt != nil
}

func (x *RoleAssignment) HasAssignedByUserId() bool {
	if x == nil {
		return false
	}
	return x.AssignedByUserId != nil
}

func (x *RoleAssignment) HasActive() bool {
	if x == nil {
		return false
	}
	return x.Active != nil
}

func (x *RoleAssignment) ClearId() {
	x.Id = nil
}

func (x *RoleAssignment) ClearUserId() {
	x.UserId = nil
}

func (x *RoleAssignment) ClearRoleId() {
	x.RoleId = nil
}

func (x *RoleAssignment) ClearResource() {
	x.Resource = nil
}

func (x *RoleAssignment) ClearScope() {
	x.Scope = nil
}

func (x *RoleAssignment) ClearCreatedAt() {
	x.CreatedAt = nil
}

func (x *RoleAssignment) ClearExpiresAt() {
	x.ExpiresAt = nil
}

func (x *RoleAssignment) ClearAssignedByUserId() {
	x.AssignedByUserId = nil
}

func (x *RoleAssignment) ClearActive() {
	x.Active = nil
}

type RoleAssignment_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// Unique assignment identifier
	Id *string
	// User ID receiving the role
	UserId *string
	// Role ID being assigned
	RoleId *string
	// Resource the role applies to (optional, for scoped roles)
	Resource *string
	// Role scope
	Scope *RoleScope
	// Assignment creation timestamp
	CreatedAt *timestamppb.Timestamp
	// Assignment expiration timestamp (optional)
	ExpiresAt *timestamppb.Timestamp
	// User who assigned the role
	AssignedByUserId *string
	// Assignment metadata
	Metadata map[string]string
	// Assignment active flag
	Active *bool
}

func (b0 RoleAssignment_builder) Build() *RoleAssignment {
	m0 := &RoleAssignment{}
	b, x := &b0, m0
	_, _ = b, x
	x.Id = b.Id
	x.UserId = b.UserId
	x.RoleId = b.RoleId
	x.Resource = b.Resource
	x.Scope = b.Scope
	x.CreatedAt = b.CreatedAt
	x.ExpiresAt = b.ExpiresAt
	x.AssignedByUserId = b.AssignedByUserId
	x.Metadata = b.Metadata
	x.Active = b.Active
	return m0
}

var File_pkg_auth_proto_messages_role_assignment_proto protoreflect.FileDescriptor

const file_pkg_auth_proto_messages_role_assignment_proto_rawDesc = "" +
	"\n" +
	"-pkg/auth/proto/messages/role_assignment.proto\x12\x0fgcommon.v1.auth\x1a!google/protobuf/go_features.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a%pkg/auth/proto/enums/role_scope.proto\"\xf1\x03\n" +
	"\x0eRoleAssignment\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x17\n" +
	"\auser_id\x18\x02 \x01(\tR\x06userId\x12\x17\n" +
	"\arole_id\x18\x03 \x01(\tR\x06roleId\x12\x1a\n" +
	"\bresource\x18\x04 \x01(\tR\bresource\x120\n" +
	"\x05scope\x18\x05 \x01(\x0e2\x1a.gcommon.v1.auth.RoleScopeR\x05scope\x12=\n" +
	"\n" +
	"created_at\x18\x06 \x01(\v2\x1a.google.protobuf.TimestampB\x02(\x01R\tcreatedAt\x12=\n" +
	"\n" +
	"expires_at\x18\a \x01(\v2\x1a.google.protobuf.TimestampB\x02(\x01R\texpiresAt\x12-\n" +
	"\x13assigned_by_user_id\x18\b \x01(\tR\x10assignedByUserId\x12M\n" +
	"\bmetadata\x18\t \x03(\v2-.gcommon.v1.auth.RoleAssignment.MetadataEntryB\x02(\x01R\bmetadata\x12\x16\n" +
	"\x06active\x18\n" +
	" \x01(\bR\x06active\x1a;\n" +
	"\rMetadataEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01B\xc1\x01\n" +
	"\x13com.gcommon.v1.authB\x13RoleAssignmentProtoP\x01Z/github.com/jdfalk/gcommon/pkg/auth/proto;authpb\xa2\x02\x03GVA\xaa\x02\x0fGcommon.V1.Auth\xca\x02\x0fGcommon\\V1\\Auth\xe2\x02\x1bGcommon\\V1\\Auth\\GPBMetadata\xea\x02\x11Gcommon::V1::Auth\x92\x03\x05\xd2>\x02\x10\x02b\beditionsp\xe8\a"

var file_pkg_auth_proto_messages_role_assignment_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_auth_proto_messages_role_assignment_proto_goTypes = []any{
	(*RoleAssignment)(nil),        // 0: gcommon.v1.auth.RoleAssignment
	nil,                           // 1: gcommon.v1.auth.RoleAssignment.MetadataEntry
	(RoleScope)(0),                // 2: gcommon.v1.auth.RoleScope
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_pkg_auth_proto_messages_role_assignment_proto_depIdxs = []int32{
	2, // 0: gcommon.v1.auth.RoleAssignment.scope:type_name -> gcommon.v1.auth.RoleScope
	3, // 1: gcommon.v1.auth.RoleAssignment.created_at:type_name -> google.protobuf.Timestamp
	3, // 2: gcommon.v1.auth.RoleAssignment.expires_at:type_name -> google.protobuf.Timestamp
	1, // 3: gcommon.v1.auth.RoleAssignment.metadata:type_name -> gcommon.v1.auth.RoleAssignment.MetadataEntry
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_pkg_auth_proto_messages_role_assignment_proto_init() }
func file_pkg_auth_proto_messages_role_assignment_proto_init() {
	if File_pkg_auth_proto_messages_role_assignment_proto != nil {
		return
	}
	file_pkg_auth_proto_enums_role_scope_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pkg_auth_proto_messages_role_assignment_proto_rawDesc), len(file_pkg_auth_proto_messages_role_assignment_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_auth_proto_messages_role_assignment_proto_goTypes,
		DependencyIndexes: file_pkg_auth_proto_messages_role_assignment_proto_depIdxs,
		MessageInfos:      file_pkg_auth_proto_messages_role_assignment_proto_msgTypes,
	}.Build()
	File_pkg_auth_proto_messages_role_assignment_proto = out.File
	file_pkg_auth_proto_messages_role_assignment_proto_goTypes = nil
	file_pkg_auth_proto_messages_role_assignment_proto_depIdxs = nil
}

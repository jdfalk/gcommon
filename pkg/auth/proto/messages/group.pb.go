// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: pkg/auth/proto/messages/group.proto

//go:build !protoopaque

package authpb

import (
	proto "github.com/jdfalk/gcommon/pkg/common/proto"
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
// Group entity for organizing users into collections.
// Used for bulk permission management and organizational structure.
// Supports hierarchical group relationships.
type Group struct {
	state protoimpl.MessageState `protogen:"hybrid.v1"`
	// Unique group identifier
	Id *string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// Group name
	Name *string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// Group description
	Description *string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	// Parent group ID (for hierarchical groups)
	ParentGroupId *string `protobuf:"bytes,4,opt,name=parent_group_id,json=parentGroupId" json:"parent_group_id,omitempty"`
	// Group permissions
	Permissions []string `protobuf:"bytes,5,rep,name=permissions" json:"permissions,omitempty"`
	// Group metadata for extensibility
	Metadata map[string]string `protobuf:"bytes,6,rep,name=metadata" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Group creation timestamp
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	// Group status
	Status *proto.ResourceStatus `protobuf:"varint,8,opt,name=status,enum=gcommon.v1.common.ResourceStatus" json:"status,omitempty"`
	// Group member count
	MemberCount *int32 `protobuf:"varint,9,opt,name=member_count,json=memberCount" json:"member_count,omitempty"`
	// Group administrator user IDs
	AdminUserIds  []string `protobuf:"bytes,10,rep,name=admin_user_ids,json=adminUserIds" json:"admin_user_ids,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Group) Reset() {
	*x = Group{}
	mi := &file_pkg_auth_proto_messages_group_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group) ProtoMessage() {}

func (x *Group) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_auth_proto_messages_group_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *Group) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *Group) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Group) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *Group) GetParentGroupId() string {
	if x != nil && x.ParentGroupId != nil {
		return *x.ParentGroupId
	}
	return ""
}

func (x *Group) GetPermissions() []string {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *Group) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Group) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Group) GetStatus() proto.ResourceStatus {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return proto.ResourceStatus(0)
}

func (x *Group) GetMemberCount() int32 {
	if x != nil && x.MemberCount != nil {
		return *x.MemberCount
	}
	return 0
}

func (x *Group) GetAdminUserIds() []string {
	if x != nil {
		return x.AdminUserIds
	}
	return nil
}

func (x *Group) SetId(v string) {
	x.Id = &v
}

func (x *Group) SetName(v string) {
	x.Name = &v
}

func (x *Group) SetDescription(v string) {
	x.Description = &v
}

func (x *Group) SetParentGroupId(v string) {
	x.ParentGroupId = &v
}

func (x *Group) SetPermissions(v []string) {
	x.Permissions = v
}

func (x *Group) SetMetadata(v map[string]string) {
	x.Metadata = v
}

func (x *Group) SetCreatedAt(v *timestamppb.Timestamp) {
	x.CreatedAt = v
}

func (x *Group) SetStatus(v proto.ResourceStatus) {
	x.Status = &v
}

func (x *Group) SetMemberCount(v int32) {
	x.MemberCount = &v
}

func (x *Group) SetAdminUserIds(v []string) {
	x.AdminUserIds = v
}

func (x *Group) HasId() bool {
	if x == nil {
		return false
	}
	return x.Id != nil
}

func (x *Group) HasName() bool {
	if x == nil {
		return false
	}
	return x.Name != nil
}

func (x *Group) HasDescription() bool {
	if x == nil {
		return false
	}
	return x.Description != nil
}

func (x *Group) HasParentGroupId() bool {
	if x == nil {
		return false
	}
	return x.ParentGroupId != nil
}

func (x *Group) HasCreatedAt() bool {
	if x == nil {
		return false
	}
	return x.CreatedAt != nil
}

func (x *Group) HasStatus() bool {
	if x == nil {
		return false
	}
	return x.Status != nil
}

func (x *Group) HasMemberCount() bool {
	if x == nil {
		return false
	}
	return x.MemberCount != nil
}

func (x *Group) ClearId() {
	x.Id = nil
}

func (x *Group) ClearName() {
	x.Name = nil
}

func (x *Group) ClearDescription() {
	x.Description = nil
}

func (x *Group) ClearParentGroupId() {
	x.ParentGroupId = nil
}

func (x *Group) ClearCreatedAt() {
	x.CreatedAt = nil
}

func (x *Group) ClearStatus() {
	x.Status = nil
}

func (x *Group) ClearMemberCount() {
	x.MemberCount = nil
}

type Group_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// Unique group identifier
	Id *string
	// Group name
	Name *string
	// Group description
	Description *string
	// Parent group ID (for hierarchical groups)
	ParentGroupId *string
	// Group permissions
	Permissions []string
	// Group metadata for extensibility
	Metadata map[string]string
	// Group creation timestamp
	CreatedAt *timestamppb.Timestamp
	// Group status
	Status *proto.ResourceStatus
	// Group member count
	MemberCount *int32
	// Group administrator user IDs
	AdminUserIds []string
}

func (b0 Group_builder) Build() *Group {
	m0 := &Group{}
	b, x := &b0, m0
	_, _ = b, x
	x.Id = b.Id
	x.Name = b.Name
	x.Description = b.Description
	x.ParentGroupId = b.ParentGroupId
	x.Permissions = b.Permissions
	x.Metadata = b.Metadata
	x.CreatedAt = b.CreatedAt
	x.Status = b.Status
	x.MemberCount = b.MemberCount
	x.AdminUserIds = b.AdminUserIds
	return m0
}

var File_pkg_auth_proto_messages_group_proto protoreflect.FileDescriptor

const file_pkg_auth_proto_messages_group_proto_rawDesc = "" +
	"\n" +
	"#pkg/auth/proto/messages/group.proto\x12\x0fgcommon.v1.auth\x1a!google/protobuf/go_features.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a,pkg/common/proto/enums/resource_status.proto\"\xdd\x03\n" +
	"\x05Group\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12&\n" +
	"\x0fparent_group_id\x18\x04 \x01(\tR\rparentGroupId\x12 \n" +
	"\vpermissions\x18\x05 \x03(\tR\vpermissions\x12D\n" +
	"\bmetadata\x18\x06 \x03(\v2$.gcommon.v1.auth.Group.MetadataEntryB\x02(\x01R\bmetadata\x12=\n" +
	"\n" +
	"created_at\x18\a \x01(\v2\x1a.google.protobuf.TimestampB\x02(\x01R\tcreatedAt\x129\n" +
	"\x06status\x18\b \x01(\x0e2!.gcommon.v1.common.ResourceStatusR\x06status\x12!\n" +
	"\fmember_count\x18\t \x01(\x05R\vmemberCount\x12$\n" +
	"\x0eadmin_user_ids\x18\n" +
	" \x03(\tR\fadminUserIds\x1a;\n" +
	"\rMetadataEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01B\xb8\x01\n" +
	"\x13com.gcommon.v1.authB\n" +
	"GroupProtoP\x01Z/github.com/jdfalk/gcommon/pkg/auth/proto;authpb\xa2\x02\x03GVA\xaa\x02\x0fGcommon.V1.Auth\xca\x02\x0fGcommon\\V1\\Auth\xe2\x02\x1bGcommon\\V1\\Auth\\GPBMetadata\xea\x02\x11Gcommon::V1::Auth\x92\x03\x05\xd2>\x02\x10\x02b\beditionsp\xe8\a"

var file_pkg_auth_proto_messages_group_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_auth_proto_messages_group_proto_goTypes = []any{
	(*Group)(nil),                 // 0: gcommon.v1.auth.Group
	nil,                           // 1: gcommon.v1.auth.Group.MetadataEntry
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
	(proto.ResourceStatus)(0),     // 3: gcommon.v1.common.ResourceStatus
}
var file_pkg_auth_proto_messages_group_proto_depIdxs = []int32{
	1, // 0: gcommon.v1.auth.Group.metadata:type_name -> gcommon.v1.auth.Group.MetadataEntry
	2, // 1: gcommon.v1.auth.Group.created_at:type_name -> google.protobuf.Timestamp
	3, // 2: gcommon.v1.auth.Group.status:type_name -> gcommon.v1.common.ResourceStatus
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pkg_auth_proto_messages_group_proto_init() }
func file_pkg_auth_proto_messages_group_proto_init() {
	if File_pkg_auth_proto_messages_group_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pkg_auth_proto_messages_group_proto_rawDesc), len(file_pkg_auth_proto_messages_group_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_auth_proto_messages_group_proto_goTypes,
		DependencyIndexes: file_pkg_auth_proto_messages_group_proto_depIdxs,
		MessageInfos:      file_pkg_auth_proto_messages_group_proto_msgTypes,
	}.Build()
	File_pkg_auth_proto_messages_group_proto = out.File
	file_pkg_auth_proto_messages_group_proto_goTypes = nil
	file_pkg_auth_proto_messages_group_proto_depIdxs = nil
}

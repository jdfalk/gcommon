// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: pkg/auth/proto/requests/update_user_request.proto

package authpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// *
// Request to update an existing user account.
type UpdateUserRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Unique identifier of the user to update
	UserId *string `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	// New email address (optional)
	Email *string `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
	// New password (should be hashed, optional)
	Password *string `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
	// New full name (optional)
	FullName *string `protobuf:"bytes,4,opt,name=full_name,json=fullName" json:"full_name,omitempty"`
	// Whether to enable/disable the account (optional)
	Enabled *bool `protobuf:"varint,5,opt,name=enabled" json:"enabled,omitempty"`
	// New roles to assign (replaces existing roles)
	Roles []string `protobuf:"bytes,6,rep,name=roles" json:"roles,omitempty"`
	// Additional metadata updates
	Metadata map[string]string `protobuf:"bytes,7,rep,name=metadata" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// New account expiration time (optional)
	ExpiresAt *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=expires_at,json=expiresAt" json:"expires_at,omitempty"`
	// Fields to update (if empty, all provided fields are updated)
	UpdateMask    []string `protobuf:"bytes,9,rep,name=update_mask,json=updateMask" json:"update_mask,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateUserRequest) Reset() {
	*x = UpdateUserRequest{}
	mi := &file_pkg_auth_proto_requests_update_user_request_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserRequest) ProtoMessage() {}

func (x *UpdateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_auth_proto_requests_update_user_request_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return file_pkg_auth_proto_requests_update_user_request_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateUserRequest) GetUserId() string {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return ""
}

func (x *UpdateUserRequest) GetEmail() string {
	if x != nil && x.Email != nil {
		return *x.Email
	}
	return ""
}

func (x *UpdateUserRequest) GetPassword() string {
	if x != nil && x.Password != nil {
		return *x.Password
	}
	return ""
}

func (x *UpdateUserRequest) GetFullName() string {
	if x != nil && x.FullName != nil {
		return *x.FullName
	}
	return ""
}

func (x *UpdateUserRequest) GetEnabled() bool {
	if x != nil && x.Enabled != nil {
		return *x.Enabled
	}
	return false
}

func (x *UpdateUserRequest) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *UpdateUserRequest) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *UpdateUserRequest) GetExpiresAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiresAt
	}
	return nil
}

func (x *UpdateUserRequest) GetUpdateMask() []string {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

var File_pkg_auth_proto_requests_update_user_request_proto protoreflect.FileDescriptor

const file_pkg_auth_proto_requests_update_user_request_proto_rawDesc = "" +
	"\n" +
	"1pkg/auth/proto/requests/update_user_request.proto\x12\x0fgcommon.v1.auth\x1a\x1fgoogle/protobuf/timestamp.proto\"\x92\x03\n" +
	"\x11UpdateUserRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x12\x14\n" +
	"\x05email\x18\x02 \x01(\tR\x05email\x12\x1a\n" +
	"\bpassword\x18\x03 \x01(\tR\bpassword\x12\x1b\n" +
	"\tfull_name\x18\x04 \x01(\tR\bfullName\x12\x18\n" +
	"\aenabled\x18\x05 \x01(\bR\aenabled\x12\x14\n" +
	"\x05roles\x18\x06 \x03(\tR\x05roles\x12L\n" +
	"\bmetadata\x18\a \x03(\v20.gcommon.v1.auth.UpdateUserRequest.MetadataEntryR\bmetadata\x129\n" +
	"\n" +
	"expires_at\x18\b \x01(\v2\x1a.google.protobuf.TimestampR\texpiresAt\x12\x1f\n" +
	"\vupdate_mask\x18\t \x03(\tR\n" +
	"updateMask\x1a;\n" +
	"\rMetadataEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01B\xbc\x01\n" +
	"\x13com.gcommon.v1.authB\x16UpdateUserRequestProtoP\x01Z/github.com/jdfalk/gcommon/pkg/auth/proto;authpb\xa2\x02\x03GVA\xaa\x02\x0fGcommon.V1.Auth\xca\x02\x0fGcommon\\V1\\Auth\xe2\x02\x1bGcommon\\V1\\Auth\\GPBMetadata\xea\x02\x11Gcommon::V1::Authb\beditionsp\xe8\a"

var (
	file_pkg_auth_proto_requests_update_user_request_proto_rawDescOnce sync.Once
	file_pkg_auth_proto_requests_update_user_request_proto_rawDescData []byte
)

func file_pkg_auth_proto_requests_update_user_request_proto_rawDescGZIP() []byte {
	file_pkg_auth_proto_requests_update_user_request_proto_rawDescOnce.Do(func() {
		file_pkg_auth_proto_requests_update_user_request_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_pkg_auth_proto_requests_update_user_request_proto_rawDesc), len(file_pkg_auth_proto_requests_update_user_request_proto_rawDesc)))
	})
	return file_pkg_auth_proto_requests_update_user_request_proto_rawDescData
}

var file_pkg_auth_proto_requests_update_user_request_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_auth_proto_requests_update_user_request_proto_goTypes = []any{
	(*UpdateUserRequest)(nil),     // 0: gcommon.v1.auth.UpdateUserRequest
	nil,                           // 1: gcommon.v1.auth.UpdateUserRequest.MetadataEntry
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_pkg_auth_proto_requests_update_user_request_proto_depIdxs = []int32{
	1, // 0: gcommon.v1.auth.UpdateUserRequest.metadata:type_name -> gcommon.v1.auth.UpdateUserRequest.MetadataEntry
	2, // 1: gcommon.v1.auth.UpdateUserRequest.expires_at:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_auth_proto_requests_update_user_request_proto_init() }
func file_pkg_auth_proto_requests_update_user_request_proto_init() {
	if File_pkg_auth_proto_requests_update_user_request_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pkg_auth_proto_requests_update_user_request_proto_rawDesc), len(file_pkg_auth_proto_requests_update_user_request_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_auth_proto_requests_update_user_request_proto_goTypes,
		DependencyIndexes: file_pkg_auth_proto_requests_update_user_request_proto_depIdxs,
		MessageInfos:      file_pkg_auth_proto_requests_update_user_request_proto_msgTypes,
	}.Build()
	File_pkg_auth_proto_requests_update_user_request_proto = out.File
	file_pkg_auth_proto_requests_update_user_request_proto_goTypes = nil
	file_pkg_auth_proto_requests_update_user_request_proto_depIdxs = nil
}

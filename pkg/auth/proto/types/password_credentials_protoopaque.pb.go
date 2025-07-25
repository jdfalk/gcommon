// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: pkg/auth/proto/types/password_credentials.proto

//go:build protoopaque

package authpb

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
// Username/password credentials for traditional authentication.
// Supports both username and email-based authentication with optional
// remember-me functionality for extended session duration.
type PasswordCredentials struct {
	state                  protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Username    *string                `protobuf:"bytes,1,opt,name=username"`
	xxx_hidden_Password    *string                `protobuf:"bytes,2,opt,name=password"`
	xxx_hidden_RememberMe  bool                   `protobuf:"varint,3,opt,name=remember_me,json=rememberMe"`
	XXX_raceDetectHookData protoimpl.RaceDetectHookData
	XXX_presence           [1]uint32
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *PasswordCredentials) Reset() {
	*x = PasswordCredentials{}
	mi := &file_pkg_auth_proto_types_password_credentials_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PasswordCredentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PasswordCredentials) ProtoMessage() {}

func (x *PasswordCredentials) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_auth_proto_types_password_credentials_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *PasswordCredentials) GetUsername() string {
	if x != nil {
		if x.xxx_hidden_Username != nil {
			return *x.xxx_hidden_Username
		}
		return ""
	}
	return ""
}

func (x *PasswordCredentials) GetPassword() string {
	if x != nil {
		if x.xxx_hidden_Password != nil {
			return *x.xxx_hidden_Password
		}
		return ""
	}
	return ""
}

func (x *PasswordCredentials) GetRememberMe() bool {
	if x != nil {
		return x.xxx_hidden_RememberMe
	}
	return false
}

func (x *PasswordCredentials) SetUsername(v string) {
	x.xxx_hidden_Username = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 3)
}

func (x *PasswordCredentials) SetPassword(v string) {
	x.xxx_hidden_Password = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 1, 3)
}

func (x *PasswordCredentials) SetRememberMe(v bool) {
	x.xxx_hidden_RememberMe = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 2, 3)
}

func (x *PasswordCredentials) HasUsername() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *PasswordCredentials) HasPassword() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 1)
}

func (x *PasswordCredentials) HasRememberMe() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 2)
}

func (x *PasswordCredentials) ClearUsername() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_Username = nil
}

func (x *PasswordCredentials) ClearPassword() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 1)
	x.xxx_hidden_Password = nil
}

func (x *PasswordCredentials) ClearRememberMe() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 2)
	x.xxx_hidden_RememberMe = false
}

type PasswordCredentials_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// Username or email address for authentication
	Username *string
	// User's password (should be transmitted over secure channels only)
	Password *string
	// Remember me option for extended session duration
	// When true, session may have longer expiration time
	RememberMe *bool
}

func (b0 PasswordCredentials_builder) Build() *PasswordCredentials {
	m0 := &PasswordCredentials{}
	b, x := &b0, m0
	_, _ = b, x
	if b.Username != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 3)
		x.xxx_hidden_Username = b.Username
	}
	if b.Password != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 1, 3)
		x.xxx_hidden_Password = b.Password
	}
	if b.RememberMe != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 2, 3)
		x.xxx_hidden_RememberMe = *b.RememberMe
	}
	return m0
}

var File_pkg_auth_proto_types_password_credentials_proto protoreflect.FileDescriptor

const file_pkg_auth_proto_types_password_credentials_proto_rawDesc = "" +
	"\n" +
	"/pkg/auth/proto/types/password_credentials.proto\x12\x0fgcommon.v1.auth\x1a!google/protobuf/go_features.proto\"n\n" +
	"\x13PasswordCredentials\x12\x1a\n" +
	"\busername\x18\x01 \x01(\tR\busername\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword\x12\x1f\n" +
	"\vremember_me\x18\x03 \x01(\bR\n" +
	"rememberMeB\xc6\x01\n" +
	"\x13com.gcommon.v1.authB\x18PasswordCredentialsProtoP\x01Z/github.com/jdfalk/gcommon/pkg/auth/proto;authpb\xa2\x02\x03GVA\xaa\x02\x0fGcommon.V1.Auth\xca\x02\x0fGcommon\\V1\\Auth\xe2\x02\x1bGcommon\\V1\\Auth\\GPBMetadata\xea\x02\x11Gcommon::V1::Auth\x92\x03\x05\xd2>\x02\x10\x02b\beditionsp\xe8\a"

var file_pkg_auth_proto_types_password_credentials_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pkg_auth_proto_types_password_credentials_proto_goTypes = []any{
	(*PasswordCredentials)(nil), // 0: gcommon.v1.auth.PasswordCredentials
}
var file_pkg_auth_proto_types_password_credentials_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_auth_proto_types_password_credentials_proto_init() }
func file_pkg_auth_proto_types_password_credentials_proto_init() {
	if File_pkg_auth_proto_types_password_credentials_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pkg_auth_proto_types_password_credentials_proto_rawDesc), len(file_pkg_auth_proto_types_password_credentials_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_auth_proto_types_password_credentials_proto_goTypes,
		DependencyIndexes: file_pkg_auth_proto_types_password_credentials_proto_depIdxs,
		MessageInfos:      file_pkg_auth_proto_types_password_credentials_proto_msgTypes,
	}.Build()
	File_pkg_auth_proto_types_password_credentials_proto = out.File
	file_pkg_auth_proto_types_password_credentials_proto_goTypes = nil
	file_pkg_auth_proto_types_password_credentials_proto_depIdxs = nil
}

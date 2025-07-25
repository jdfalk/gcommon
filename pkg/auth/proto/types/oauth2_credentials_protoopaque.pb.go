// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: pkg/auth/proto/types/oauth2_credentials.proto

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
// OAuth2 credentials for authorization code flow authentication.
// Implements standard OAuth2 authorization code exchange for tokens.
type OAuth2Credentials struct {
	state                   protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Code         *string                `protobuf:"bytes,1,opt,name=code"`
	xxx_hidden_RedirectUri  *string                `protobuf:"bytes,2,opt,name=redirect_uri,json=redirectUri"`
	xxx_hidden_ClientId     *string                `protobuf:"bytes,3,opt,name=client_id,json=clientId"`
	xxx_hidden_ClientSecret *string                `protobuf:"bytes,4,opt,name=client_secret,json=clientSecret"`
	XXX_raceDetectHookData  protoimpl.RaceDetectHookData
	XXX_presence            [1]uint32
	unknownFields           protoimpl.UnknownFields
	sizeCache               protoimpl.SizeCache
}

func (x *OAuth2Credentials) Reset() {
	*x = OAuth2Credentials{}
	mi := &file_pkg_auth_proto_types_oauth2_credentials_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OAuth2Credentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OAuth2Credentials) ProtoMessage() {}

func (x *OAuth2Credentials) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_auth_proto_types_oauth2_credentials_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *OAuth2Credentials) GetCode() string {
	if x != nil {
		if x.xxx_hidden_Code != nil {
			return *x.xxx_hidden_Code
		}
		return ""
	}
	return ""
}

func (x *OAuth2Credentials) GetRedirectUri() string {
	if x != nil {
		if x.xxx_hidden_RedirectUri != nil {
			return *x.xxx_hidden_RedirectUri
		}
		return ""
	}
	return ""
}

func (x *OAuth2Credentials) GetClientId() string {
	if x != nil {
		if x.xxx_hidden_ClientId != nil {
			return *x.xxx_hidden_ClientId
		}
		return ""
	}
	return ""
}

func (x *OAuth2Credentials) GetClientSecret() string {
	if x != nil {
		if x.xxx_hidden_ClientSecret != nil {
			return *x.xxx_hidden_ClientSecret
		}
		return ""
	}
	return ""
}

func (x *OAuth2Credentials) SetCode(v string) {
	x.xxx_hidden_Code = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 4)
}

func (x *OAuth2Credentials) SetRedirectUri(v string) {
	x.xxx_hidden_RedirectUri = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 1, 4)
}

func (x *OAuth2Credentials) SetClientId(v string) {
	x.xxx_hidden_ClientId = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 2, 4)
}

func (x *OAuth2Credentials) SetClientSecret(v string) {
	x.xxx_hidden_ClientSecret = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 3, 4)
}

func (x *OAuth2Credentials) HasCode() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *OAuth2Credentials) HasRedirectUri() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 1)
}

func (x *OAuth2Credentials) HasClientId() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 2)
}

func (x *OAuth2Credentials) HasClientSecret() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 3)
}

func (x *OAuth2Credentials) ClearCode() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_Code = nil
}

func (x *OAuth2Credentials) ClearRedirectUri() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 1)
	x.xxx_hidden_RedirectUri = nil
}

func (x *OAuth2Credentials) ClearClientId() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 2)
	x.xxx_hidden_ClientId = nil
}

func (x *OAuth2Credentials) ClearClientSecret() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 3)
	x.xxx_hidden_ClientSecret = nil
}

type OAuth2Credentials_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// Authorization code received from OAuth2 provider
	Code *string
	// Redirect URI that was used in the authorization request
	// Must match the URI registered with the OAuth2 provider
	RedirectUri *string
	// OAuth2 client identifier
	ClientId *string
	// OAuth2 client secret (for confidential clients only)
	// Should be omitted for public clients (e.g., mobile apps, SPAs)
	ClientSecret *string
}

func (b0 OAuth2Credentials_builder) Build() *OAuth2Credentials {
	m0 := &OAuth2Credentials{}
	b, x := &b0, m0
	_, _ = b, x
	if b.Code != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 4)
		x.xxx_hidden_Code = b.Code
	}
	if b.RedirectUri != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 1, 4)
		x.xxx_hidden_RedirectUri = b.RedirectUri
	}
	if b.ClientId != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 2, 4)
		x.xxx_hidden_ClientId = b.ClientId
	}
	if b.ClientSecret != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 3, 4)
		x.xxx_hidden_ClientSecret = b.ClientSecret
	}
	return m0
}

var File_pkg_auth_proto_types_oauth2_credentials_proto protoreflect.FileDescriptor

const file_pkg_auth_proto_types_oauth2_credentials_proto_rawDesc = "" +
	"\n" +
	"-pkg/auth/proto/types/oauth2_credentials.proto\x12\x0fgcommon.v1.auth\x1a!google/protobuf/go_features.proto\"\x8c\x01\n" +
	"\x11OAuth2Credentials\x12\x12\n" +
	"\x04code\x18\x01 \x01(\tR\x04code\x12!\n" +
	"\fredirect_uri\x18\x02 \x01(\tR\vredirectUri\x12\x1b\n" +
	"\tclient_id\x18\x03 \x01(\tR\bclientId\x12#\n" +
	"\rclient_secret\x18\x04 \x01(\tR\fclientSecretB\xc4\x01\n" +
	"\x13com.gcommon.v1.authB\x16Oauth2CredentialsProtoP\x01Z/github.com/jdfalk/gcommon/pkg/auth/proto;authpb\xa2\x02\x03GVA\xaa\x02\x0fGcommon.V1.Auth\xca\x02\x0fGcommon\\V1\\Auth\xe2\x02\x1bGcommon\\V1\\Auth\\GPBMetadata\xea\x02\x11Gcommon::V1::Auth\x92\x03\x05\xd2>\x02\x10\x02b\beditionsp\xe8\a"

var file_pkg_auth_proto_types_oauth2_credentials_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pkg_auth_proto_types_oauth2_credentials_proto_goTypes = []any{
	(*OAuth2Credentials)(nil), // 0: gcommon.v1.auth.OAuth2Credentials
}
var file_pkg_auth_proto_types_oauth2_credentials_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_auth_proto_types_oauth2_credentials_proto_init() }
func file_pkg_auth_proto_types_oauth2_credentials_proto_init() {
	if File_pkg_auth_proto_types_oauth2_credentials_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pkg_auth_proto_types_oauth2_credentials_proto_rawDesc), len(file_pkg_auth_proto_types_oauth2_credentials_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_auth_proto_types_oauth2_credentials_proto_goTypes,
		DependencyIndexes: file_pkg_auth_proto_types_oauth2_credentials_proto_depIdxs,
		MessageInfos:      file_pkg_auth_proto_types_oauth2_credentials_proto_msgTypes,
	}.Build()
	File_pkg_auth_proto_types_oauth2_credentials_proto = out.File
	file_pkg_auth_proto_types_oauth2_credentials_proto_goTypes = nil
	file_pkg_auth_proto_types_oauth2_credentials_proto_depIdxs = nil
}

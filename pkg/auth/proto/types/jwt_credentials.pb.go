// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: pkg/auth/proto/types/jwt_credentials.proto

//go:build !protoopaque

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
// JWT (JSON Web Token) credentials for token-based authentication.
// Supports validation of externally issued JWTs with optional issuer verification.
type JWTCredentials struct {
	state protoimpl.MessageState `protogen:"hybrid.v1"`
	// JWT token string (header.payload.signature format)
	Token *string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	// Expected issuer for JWT validation (optional)
	// When provided, the JWT's 'iss' claim must match this value
	Issuer        *string `protobuf:"bytes,2,opt,name=issuer" json:"issuer,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *JWTCredentials) Reset() {
	*x = JWTCredentials{}
	mi := &file_pkg_auth_proto_types_jwt_credentials_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JWTCredentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JWTCredentials) ProtoMessage() {}

func (x *JWTCredentials) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_auth_proto_types_jwt_credentials_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *JWTCredentials) GetToken() string {
	if x != nil && x.Token != nil {
		return *x.Token
	}
	return ""
}

func (x *JWTCredentials) GetIssuer() string {
	if x != nil && x.Issuer != nil {
		return *x.Issuer
	}
	return ""
}

func (x *JWTCredentials) SetToken(v string) {
	x.Token = &v
}

func (x *JWTCredentials) SetIssuer(v string) {
	x.Issuer = &v
}

func (x *JWTCredentials) HasToken() bool {
	if x == nil {
		return false
	}
	return x.Token != nil
}

func (x *JWTCredentials) HasIssuer() bool {
	if x == nil {
		return false
	}
	return x.Issuer != nil
}

func (x *JWTCredentials) ClearToken() {
	x.Token = nil
}

func (x *JWTCredentials) ClearIssuer() {
	x.Issuer = nil
}

type JWTCredentials_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// JWT token string (header.payload.signature format)
	Token *string
	// Expected issuer for JWT validation (optional)
	// When provided, the JWT's 'iss' claim must match this value
	Issuer *string
}

func (b0 JWTCredentials_builder) Build() *JWTCredentials {
	m0 := &JWTCredentials{}
	b, x := &b0, m0
	_, _ = b, x
	x.Token = b.Token
	x.Issuer = b.Issuer
	return m0
}

var File_pkg_auth_proto_types_jwt_credentials_proto protoreflect.FileDescriptor

const file_pkg_auth_proto_types_jwt_credentials_proto_rawDesc = "" +
	"\n" +
	"*pkg/auth/proto/types/jwt_credentials.proto\x12\x0fgcommon.v1.auth\x1a!google/protobuf/go_features.proto\">\n" +
	"\x0eJWTCredentials\x12\x14\n" +
	"\x05token\x18\x01 \x01(\tR\x05token\x12\x16\n" +
	"\x06issuer\x18\x02 \x01(\tR\x06issuerB\xc1\x01\n" +
	"\x13com.gcommon.v1.authB\x13JwtCredentialsProtoP\x01Z/github.com/jdfalk/gcommon/pkg/auth/proto;authpb\xa2\x02\x03GVA\xaa\x02\x0fGcommon.V1.Auth\xca\x02\x0fGcommon\\V1\\Auth\xe2\x02\x1bGcommon\\V1\\Auth\\GPBMetadata\xea\x02\x11Gcommon::V1::Auth\x92\x03\x05\xd2>\x02\x10\x02b\beditionsp\xe8\a"

var file_pkg_auth_proto_types_jwt_credentials_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pkg_auth_proto_types_jwt_credentials_proto_goTypes = []any{
	(*JWTCredentials)(nil), // 0: gcommon.v1.auth.JWTCredentials
}
var file_pkg_auth_proto_types_jwt_credentials_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_auth_proto_types_jwt_credentials_proto_init() }
func file_pkg_auth_proto_types_jwt_credentials_proto_init() {
	if File_pkg_auth_proto_types_jwt_credentials_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pkg_auth_proto_types_jwt_credentials_proto_rawDesc), len(file_pkg_auth_proto_types_jwt_credentials_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_auth_proto_types_jwt_credentials_proto_goTypes,
		DependencyIndexes: file_pkg_auth_proto_types_jwt_credentials_proto_depIdxs,
		MessageInfos:      file_pkg_auth_proto_types_jwt_credentials_proto_msgTypes,
	}.Build()
	File_pkg_auth_proto_types_jwt_credentials_proto = out.File
	file_pkg_auth_proto_types_jwt_credentials_proto_goTypes = nil
	file_pkg_auth_proto_types_jwt_credentials_proto_depIdxs = nil
}

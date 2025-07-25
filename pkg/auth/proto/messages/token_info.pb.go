// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: pkg/auth/proto/messages/token_info.proto

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
// Token information for lightweight token tracking.
// Contains essential token data without sensitive information.
// Used in responses where full token data is not needed.
type TokenInfo struct {
	state protoimpl.MessageState `protogen:"hybrid.v1"`
	// Token identifier
	TokenId *string `protobuf:"bytes,1,opt,name=token_id,json=tokenId" json:"token_id,omitempty"`
	// Token type
	Type *TokenType `protobuf:"varint,2,opt,name=type,enum=gcommon.v1.auth.TokenType" json:"type,omitempty"`
	// User ID associated with token
	UserId *string `protobuf:"bytes,3,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	// Client ID that owns the token
	ClientId *string `protobuf:"bytes,4,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
	// Token scopes
	Scopes []string `protobuf:"bytes,5,rep,name=scopes" json:"scopes,omitempty"`
	// Token creation time
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	// Token expiration time
	ExpiresAt *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=expires_at,json=expiresAt" json:"expires_at,omitempty"`
	// Token active flag
	Active *bool `protobuf:"varint,8,opt,name=active" json:"active,omitempty"`
	// Time until expiration (seconds)
	ExpiresIn     *int64 `protobuf:"varint,9,opt,name=expires_in,json=expiresIn" json:"expires_in,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TokenInfo) Reset() {
	*x = TokenInfo{}
	mi := &file_pkg_auth_proto_messages_token_info_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TokenInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenInfo) ProtoMessage() {}

func (x *TokenInfo) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_auth_proto_messages_token_info_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *TokenInfo) GetTokenId() string {
	if x != nil && x.TokenId != nil {
		return *x.TokenId
	}
	return ""
}

func (x *TokenInfo) GetType() TokenType {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return TokenType_TOKEN_TYPE_UNSPECIFIED
}

func (x *TokenInfo) GetUserId() string {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return ""
}

func (x *TokenInfo) GetClientId() string {
	if x != nil && x.ClientId != nil {
		return *x.ClientId
	}
	return ""
}

func (x *TokenInfo) GetScopes() []string {
	if x != nil {
		return x.Scopes
	}
	return nil
}

func (x *TokenInfo) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *TokenInfo) GetExpiresAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiresAt
	}
	return nil
}

func (x *TokenInfo) GetActive() bool {
	if x != nil && x.Active != nil {
		return *x.Active
	}
	return false
}

func (x *TokenInfo) GetExpiresIn() int64 {
	if x != nil && x.ExpiresIn != nil {
		return *x.ExpiresIn
	}
	return 0
}

func (x *TokenInfo) SetTokenId(v string) {
	x.TokenId = &v
}

func (x *TokenInfo) SetType(v TokenType) {
	x.Type = &v
}

func (x *TokenInfo) SetUserId(v string) {
	x.UserId = &v
}

func (x *TokenInfo) SetClientId(v string) {
	x.ClientId = &v
}

func (x *TokenInfo) SetScopes(v []string) {
	x.Scopes = v
}

func (x *TokenInfo) SetCreatedAt(v *timestamppb.Timestamp) {
	x.CreatedAt = v
}

func (x *TokenInfo) SetExpiresAt(v *timestamppb.Timestamp) {
	x.ExpiresAt = v
}

func (x *TokenInfo) SetActive(v bool) {
	x.Active = &v
}

func (x *TokenInfo) SetExpiresIn(v int64) {
	x.ExpiresIn = &v
}

func (x *TokenInfo) HasTokenId() bool {
	if x == nil {
		return false
	}
	return x.TokenId != nil
}

func (x *TokenInfo) HasType() bool {
	if x == nil {
		return false
	}
	return x.Type != nil
}

func (x *TokenInfo) HasUserId() bool {
	if x == nil {
		return false
	}
	return x.UserId != nil
}

func (x *TokenInfo) HasClientId() bool {
	if x == nil {
		return false
	}
	return x.ClientId != nil
}

func (x *TokenInfo) HasCreatedAt() bool {
	if x == nil {
		return false
	}
	return x.CreatedAt != nil
}

func (x *TokenInfo) HasExpiresAt() bool {
	if x == nil {
		return false
	}
	return x.ExpiresAt != nil
}

func (x *TokenInfo) HasActive() bool {
	if x == nil {
		return false
	}
	return x.Active != nil
}

func (x *TokenInfo) HasExpiresIn() bool {
	if x == nil {
		return false
	}
	return x.ExpiresIn != nil
}

func (x *TokenInfo) ClearTokenId() {
	x.TokenId = nil
}

func (x *TokenInfo) ClearType() {
	x.Type = nil
}

func (x *TokenInfo) ClearUserId() {
	x.UserId = nil
}

func (x *TokenInfo) ClearClientId() {
	x.ClientId = nil
}

func (x *TokenInfo) ClearCreatedAt() {
	x.CreatedAt = nil
}

func (x *TokenInfo) ClearExpiresAt() {
	x.ExpiresAt = nil
}

func (x *TokenInfo) ClearActive() {
	x.Active = nil
}

func (x *TokenInfo) ClearExpiresIn() {
	x.ExpiresIn = nil
}

type TokenInfo_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// Token identifier
	TokenId *string
	// Token type
	Type *TokenType
	// User ID associated with token
	UserId *string
	// Client ID that owns the token
	ClientId *string
	// Token scopes
	Scopes []string
	// Token creation time
	CreatedAt *timestamppb.Timestamp
	// Token expiration time
	ExpiresAt *timestamppb.Timestamp
	// Token active flag
	Active *bool
	// Time until expiration (seconds)
	ExpiresIn *int64
}

func (b0 TokenInfo_builder) Build() *TokenInfo {
	m0 := &TokenInfo{}
	b, x := &b0, m0
	_, _ = b, x
	x.TokenId = b.TokenId
	x.Type = b.Type
	x.UserId = b.UserId
	x.ClientId = b.ClientId
	x.Scopes = b.Scopes
	x.CreatedAt = b.CreatedAt
	x.ExpiresAt = b.ExpiresAt
	x.Active = b.Active
	x.ExpiresIn = b.ExpiresIn
	return m0
}

var File_pkg_auth_proto_messages_token_info_proto protoreflect.FileDescriptor

const file_pkg_auth_proto_messages_token_info_proto_rawDesc = "" +
	"\n" +
	"(pkg/auth/proto/messages/token_info.proto\x12\x0fgcommon.v1.auth\x1a!google/protobuf/go_features.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a%pkg/auth/proto/enums/token_type.proto\"\xd1\x02\n" +
	"\tTokenInfo\x12\x19\n" +
	"\btoken_id\x18\x01 \x01(\tR\atokenId\x12.\n" +
	"\x04type\x18\x02 \x01(\x0e2\x1a.gcommon.v1.auth.TokenTypeR\x04type\x12\x17\n" +
	"\auser_id\x18\x03 \x01(\tR\x06userId\x12\x1b\n" +
	"\tclient_id\x18\x04 \x01(\tR\bclientId\x12\x16\n" +
	"\x06scopes\x18\x05 \x03(\tR\x06scopes\x129\n" +
	"\n" +
	"created_at\x18\x06 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x129\n" +
	"\n" +
	"expires_at\x18\a \x01(\v2\x1a.google.protobuf.TimestampR\texpiresAt\x12\x16\n" +
	"\x06active\x18\b \x01(\bR\x06active\x12\x1d\n" +
	"\n" +
	"expires_in\x18\t \x01(\x03R\texpiresInB\xbc\x01\n" +
	"\x13com.gcommon.v1.authB\x0eTokenInfoProtoP\x01Z/github.com/jdfalk/gcommon/pkg/auth/proto;authpb\xa2\x02\x03GVA\xaa\x02\x0fGcommon.V1.Auth\xca\x02\x0fGcommon\\V1\\Auth\xe2\x02\x1bGcommon\\V1\\Auth\\GPBMetadata\xea\x02\x11Gcommon::V1::Auth\x92\x03\x05\xd2>\x02\x10\x02b\beditionsp\xe8\a"

var file_pkg_auth_proto_messages_token_info_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pkg_auth_proto_messages_token_info_proto_goTypes = []any{
	(*TokenInfo)(nil),             // 0: gcommon.v1.auth.TokenInfo
	(TokenType)(0),                // 1: gcommon.v1.auth.TokenType
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_pkg_auth_proto_messages_token_info_proto_depIdxs = []int32{
	1, // 0: gcommon.v1.auth.TokenInfo.type:type_name -> gcommon.v1.auth.TokenType
	2, // 1: gcommon.v1.auth.TokenInfo.created_at:type_name -> google.protobuf.Timestamp
	2, // 2: gcommon.v1.auth.TokenInfo.expires_at:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pkg_auth_proto_messages_token_info_proto_init() }
func file_pkg_auth_proto_messages_token_info_proto_init() {
	if File_pkg_auth_proto_messages_token_info_proto != nil {
		return
	}
	file_pkg_auth_proto_enums_token_type_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pkg_auth_proto_messages_token_info_proto_rawDesc), len(file_pkg_auth_proto_messages_token_info_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_auth_proto_messages_token_info_proto_goTypes,
		DependencyIndexes: file_pkg_auth_proto_messages_token_info_proto_depIdxs,
		MessageInfos:      file_pkg_auth_proto_messages_token_info_proto_msgTypes,
	}.Build()
	File_pkg_auth_proto_messages_token_info_proto = out.File
	file_pkg_auth_proto_messages_token_info_proto_goTypes = nil
	file_pkg_auth_proto_messages_token_info_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: pkg/auth/proto/messages/oauth_client.proto

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
// OAuth2 client configuration for third-party integrations.
// Used for OAuth2 authorization code and implicit flows.
// Contains client credentials and configuration.
type OAuthClient struct {
	state protoimpl.MessageState `protogen:"hybrid.v1"`
	// Unique client identifier
	ClientId *string `protobuf:"bytes,1,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
	// Client secret (for confidential clients)
	ClientSecret *string `protobuf:"bytes,2,opt,name=client_secret,json=clientSecret" json:"client_secret,omitempty"`
	// Client name
	Name *string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	// Client description
	Description *string `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
	// Client type ("public" or "confidential")
	ClientType *string `protobuf:"bytes,5,opt,name=client_type,json=clientType" json:"client_type,omitempty"`
	// Allowed redirect URIs
	RedirectUris []string `protobuf:"bytes,6,rep,name=redirect_uris,json=redirectUris" json:"redirect_uris,omitempty"`
	// Allowed grant types
	GrantTypes []string `protobuf:"bytes,7,rep,name=grant_types,json=grantTypes" json:"grant_types,omitempty"`
	// Allowed response types
	ResponseTypes []string `protobuf:"bytes,8,rep,name=response_types,json=responseTypes" json:"response_types,omitempty"`
	// Allowed scopes
	Scopes []string `protobuf:"bytes,9,rep,name=scopes" json:"scopes,omitempty"`
	// Client creation timestamp
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	// Client status
	Status *proto.ResourceStatus `protobuf:"varint,11,opt,name=status,enum=gcommon.v1.common.ResourceStatus" json:"status,omitempty"`
	// Client metadata
	Metadata map[string]string `protobuf:"bytes,12,rep,name=metadata" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Client logo URL
	LogoUrl *string `protobuf:"bytes,13,opt,name=logo_url,json=logoUrl" json:"logo_url,omitempty"`
	// Client website URL
	WebsiteUrl *string `protobuf:"bytes,14,opt,name=website_url,json=websiteUrl" json:"website_url,omitempty"`
	// Client owner user ID
	OwnerUserId   *string `protobuf:"bytes,15,opt,name=owner_user_id,json=ownerUserId" json:"owner_user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OAuthClient) Reset() {
	*x = OAuthClient{}
	mi := &file_pkg_auth_proto_messages_oauth_client_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OAuthClient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OAuthClient) ProtoMessage() {}

func (x *OAuthClient) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_auth_proto_messages_oauth_client_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *OAuthClient) GetClientId() string {
	if x != nil && x.ClientId != nil {
		return *x.ClientId
	}
	return ""
}

func (x *OAuthClient) GetClientSecret() string {
	if x != nil && x.ClientSecret != nil {
		return *x.ClientSecret
	}
	return ""
}

func (x *OAuthClient) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *OAuthClient) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *OAuthClient) GetClientType() string {
	if x != nil && x.ClientType != nil {
		return *x.ClientType
	}
	return ""
}

func (x *OAuthClient) GetRedirectUris() []string {
	if x != nil {
		return x.RedirectUris
	}
	return nil
}

func (x *OAuthClient) GetGrantTypes() []string {
	if x != nil {
		return x.GrantTypes
	}
	return nil
}

func (x *OAuthClient) GetResponseTypes() []string {
	if x != nil {
		return x.ResponseTypes
	}
	return nil
}

func (x *OAuthClient) GetScopes() []string {
	if x != nil {
		return x.Scopes
	}
	return nil
}

func (x *OAuthClient) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *OAuthClient) GetStatus() proto.ResourceStatus {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return proto.ResourceStatus(0)
}

func (x *OAuthClient) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *OAuthClient) GetLogoUrl() string {
	if x != nil && x.LogoUrl != nil {
		return *x.LogoUrl
	}
	return ""
}

func (x *OAuthClient) GetWebsiteUrl() string {
	if x != nil && x.WebsiteUrl != nil {
		return *x.WebsiteUrl
	}
	return ""
}

func (x *OAuthClient) GetOwnerUserId() string {
	if x != nil && x.OwnerUserId != nil {
		return *x.OwnerUserId
	}
	return ""
}

func (x *OAuthClient) SetClientId(v string) {
	x.ClientId = &v
}

func (x *OAuthClient) SetClientSecret(v string) {
	x.ClientSecret = &v
}

func (x *OAuthClient) SetName(v string) {
	x.Name = &v
}

func (x *OAuthClient) SetDescription(v string) {
	x.Description = &v
}

func (x *OAuthClient) SetClientType(v string) {
	x.ClientType = &v
}

func (x *OAuthClient) SetRedirectUris(v []string) {
	x.RedirectUris = v
}

func (x *OAuthClient) SetGrantTypes(v []string) {
	x.GrantTypes = v
}

func (x *OAuthClient) SetResponseTypes(v []string) {
	x.ResponseTypes = v
}

func (x *OAuthClient) SetScopes(v []string) {
	x.Scopes = v
}

func (x *OAuthClient) SetCreatedAt(v *timestamppb.Timestamp) {
	x.CreatedAt = v
}

func (x *OAuthClient) SetStatus(v proto.ResourceStatus) {
	x.Status = &v
}

func (x *OAuthClient) SetMetadata(v map[string]string) {
	x.Metadata = v
}

func (x *OAuthClient) SetLogoUrl(v string) {
	x.LogoUrl = &v
}

func (x *OAuthClient) SetWebsiteUrl(v string) {
	x.WebsiteUrl = &v
}

func (x *OAuthClient) SetOwnerUserId(v string) {
	x.OwnerUserId = &v
}

func (x *OAuthClient) HasClientId() bool {
	if x == nil {
		return false
	}
	return x.ClientId != nil
}

func (x *OAuthClient) HasClientSecret() bool {
	if x == nil {
		return false
	}
	return x.ClientSecret != nil
}

func (x *OAuthClient) HasName() bool {
	if x == nil {
		return false
	}
	return x.Name != nil
}

func (x *OAuthClient) HasDescription() bool {
	if x == nil {
		return false
	}
	return x.Description != nil
}

func (x *OAuthClient) HasClientType() bool {
	if x == nil {
		return false
	}
	return x.ClientType != nil
}

func (x *OAuthClient) HasCreatedAt() bool {
	if x == nil {
		return false
	}
	return x.CreatedAt != nil
}

func (x *OAuthClient) HasStatus() bool {
	if x == nil {
		return false
	}
	return x.Status != nil
}

func (x *OAuthClient) HasLogoUrl() bool {
	if x == nil {
		return false
	}
	return x.LogoUrl != nil
}

func (x *OAuthClient) HasWebsiteUrl() bool {
	if x == nil {
		return false
	}
	return x.WebsiteUrl != nil
}

func (x *OAuthClient) HasOwnerUserId() bool {
	if x == nil {
		return false
	}
	return x.OwnerUserId != nil
}

func (x *OAuthClient) ClearClientId() {
	x.ClientId = nil
}

func (x *OAuthClient) ClearClientSecret() {
	x.ClientSecret = nil
}

func (x *OAuthClient) ClearName() {
	x.Name = nil
}

func (x *OAuthClient) ClearDescription() {
	x.Description = nil
}

func (x *OAuthClient) ClearClientType() {
	x.ClientType = nil
}

func (x *OAuthClient) ClearCreatedAt() {
	x.CreatedAt = nil
}

func (x *OAuthClient) ClearStatus() {
	x.Status = nil
}

func (x *OAuthClient) ClearLogoUrl() {
	x.LogoUrl = nil
}

func (x *OAuthClient) ClearWebsiteUrl() {
	x.WebsiteUrl = nil
}

func (x *OAuthClient) ClearOwnerUserId() {
	x.OwnerUserId = nil
}

type OAuthClient_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// Unique client identifier
	ClientId *string
	// Client secret (for confidential clients)
	ClientSecret *string
	// Client name
	Name *string
	// Client description
	Description *string
	// Client type ("public" or "confidential")
	ClientType *string
	// Allowed redirect URIs
	RedirectUris []string
	// Allowed grant types
	GrantTypes []string
	// Allowed response types
	ResponseTypes []string
	// Allowed scopes
	Scopes []string
	// Client creation timestamp
	CreatedAt *timestamppb.Timestamp
	// Client status
	Status *proto.ResourceStatus
	// Client metadata
	Metadata map[string]string
	// Client logo URL
	LogoUrl *string
	// Client website URL
	WebsiteUrl *string
	// Client owner user ID
	OwnerUserId *string
}

func (b0 OAuthClient_builder) Build() *OAuthClient {
	m0 := &OAuthClient{}
	b, x := &b0, m0
	_, _ = b, x
	x.ClientId = b.ClientId
	x.ClientSecret = b.ClientSecret
	x.Name = b.Name
	x.Description = b.Description
	x.ClientType = b.ClientType
	x.RedirectUris = b.RedirectUris
	x.GrantTypes = b.GrantTypes
	x.ResponseTypes = b.ResponseTypes
	x.Scopes = b.Scopes
	x.CreatedAt = b.CreatedAt
	x.Status = b.Status
	x.Metadata = b.Metadata
	x.LogoUrl = b.LogoUrl
	x.WebsiteUrl = b.WebsiteUrl
	x.OwnerUserId = b.OwnerUserId
	return m0
}

var File_pkg_auth_proto_messages_oauth_client_proto protoreflect.FileDescriptor

const file_pkg_auth_proto_messages_oauth_client_proto_rawDesc = "" +
	"\n" +
	"*pkg/auth/proto/messages/oauth_client.proto\x12\x0fgcommon.v1.auth\x1a!google/protobuf/go_features.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a,pkg/common/proto/enums/resource_status.proto\"\x8e\x05\n" +
	"\vOAuthClient\x12\x1b\n" +
	"\tclient_id\x18\x01 \x01(\tR\bclientId\x12#\n" +
	"\rclient_secret\x18\x02 \x01(\tR\fclientSecret\x12\x12\n" +
	"\x04name\x18\x03 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x04 \x01(\tR\vdescription\x12\x1f\n" +
	"\vclient_type\x18\x05 \x01(\tR\n" +
	"clientType\x12#\n" +
	"\rredirect_uris\x18\x06 \x03(\tR\fredirectUris\x12\x1f\n" +
	"\vgrant_types\x18\a \x03(\tR\n" +
	"grantTypes\x12%\n" +
	"\x0eresponse_types\x18\b \x03(\tR\rresponseTypes\x12\x16\n" +
	"\x06scopes\x18\t \x03(\tR\x06scopes\x12=\n" +
	"\n" +
	"created_at\x18\n" +
	" \x01(\v2\x1a.google.protobuf.TimestampB\x02(\x01R\tcreatedAt\x129\n" +
	"\x06status\x18\v \x01(\x0e2!.gcommon.v1.common.ResourceStatusR\x06status\x12J\n" +
	"\bmetadata\x18\f \x03(\v2*.gcommon.v1.auth.OAuthClient.MetadataEntryB\x02(\x01R\bmetadata\x12\x19\n" +
	"\blogo_url\x18\r \x01(\tR\alogoUrl\x12\x1f\n" +
	"\vwebsite_url\x18\x0e \x01(\tR\n" +
	"websiteUrl\x12\"\n" +
	"\rowner_user_id\x18\x0f \x01(\tR\vownerUserId\x1a;\n" +
	"\rMetadataEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01B\xbe\x01\n" +
	"\x13com.gcommon.v1.authB\x10OauthClientProtoP\x01Z/github.com/jdfalk/gcommon/pkg/auth/proto;authpb\xa2\x02\x03GVA\xaa\x02\x0fGcommon.V1.Auth\xca\x02\x0fGcommon\\V1\\Auth\xe2\x02\x1bGcommon\\V1\\Auth\\GPBMetadata\xea\x02\x11Gcommon::V1::Auth\x92\x03\x05\xd2>\x02\x10\x02b\beditionsp\xe8\a"

var file_pkg_auth_proto_messages_oauth_client_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_auth_proto_messages_oauth_client_proto_goTypes = []any{
	(*OAuthClient)(nil),           // 0: gcommon.v1.auth.OAuthClient
	nil,                           // 1: gcommon.v1.auth.OAuthClient.MetadataEntry
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
	(proto.ResourceStatus)(0),     // 3: gcommon.v1.common.ResourceStatus
}
var file_pkg_auth_proto_messages_oauth_client_proto_depIdxs = []int32{
	2, // 0: gcommon.v1.auth.OAuthClient.created_at:type_name -> google.protobuf.Timestamp
	3, // 1: gcommon.v1.auth.OAuthClient.status:type_name -> gcommon.v1.common.ResourceStatus
	1, // 2: gcommon.v1.auth.OAuthClient.metadata:type_name -> gcommon.v1.auth.OAuthClient.MetadataEntry
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pkg_auth_proto_messages_oauth_client_proto_init() }
func file_pkg_auth_proto_messages_oauth_client_proto_init() {
	if File_pkg_auth_proto_messages_oauth_client_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pkg_auth_proto_messages_oauth_client_proto_rawDesc), len(file_pkg_auth_proto_messages_oauth_client_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_auth_proto_messages_oauth_client_proto_goTypes,
		DependencyIndexes: file_pkg_auth_proto_messages_oauth_client_proto_depIdxs,
		MessageInfos:      file_pkg_auth_proto_messages_oauth_client_proto_msgTypes,
	}.Build()
	File_pkg_auth_proto_messages_oauth_client_proto = out.File
	file_pkg_auth_proto_messages_oauth_client_proto_goTypes = nil
	file_pkg_auth_proto_messages_oauth_client_proto_depIdxs = nil
}

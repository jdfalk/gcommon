// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: pkg/auth/proto/requests/list_user_sessions_request.proto

package authpb

import (
	proto "github.com/jdfalk/gcommon/pkg/common/proto"
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
// ListUserSessionsRequest requests a list of active sessions for a user.
// Used by administrators to monitor user sessions and by users to view
// their own active sessions across devices.
//
// Follows 1-1-1 pattern: one message per file.
type ListUserSessionsRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// *
	// The ID of the user whose sessions should be listed.
	// This can be the requesting user's own ID or another user's ID
	// if the requester has appropriate permissions.
	UserId *string `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	// *
	// Standard request metadata including authentication context,
	// tracing information, and client details.
	Metadata *proto.RequestMetadata `protobuf:"bytes,11,opt,name=metadata" json:"metadata,omitempty"`
	// *
	// Pagination options for large result sets.
	// Defaults to first 50 sessions if not specified.
	Pagination *proto.Pagination `protobuf:"bytes,12,opt,name=pagination" json:"pagination,omitempty"`
	// *
	// Filter sessions by status (active, expired, terminated).
	// If not specified, returns all sessions.
	StatusFilter *string `protobuf:"bytes,13,opt,name=status_filter,json=statusFilter" json:"status_filter,omitempty"`
	// *
	// Filter sessions by device type (web, mobile, api, etc.).
	// If not specified, returns sessions from all device types.
	DeviceTypeFilter *string `protobuf:"bytes,14,opt,name=device_type_filter,json=deviceTypeFilter" json:"device_type_filter,omitempty"`
	// *
	// Only return sessions created after this timestamp.
	// Useful for finding recent sessions.
	CreatedAfter *timestamppb.Timestamp `protobuf:"bytes,15,opt,name=created_after,json=createdAfter" json:"created_after,omitempty"`
	// *
	// Only return sessions created before this timestamp.
	// Useful for historical analysis.
	CreatedBefore *timestamppb.Timestamp `protobuf:"bytes,16,opt,name=created_before,json=createdBefore" json:"created_before,omitempty"`
	// *
	// Include detailed session information (IP addresses, user agents, etc.).
	// May require additional permissions. Defaults to false.
	IncludeDetails *bool `protobuf:"varint,17,opt,name=include_details,json=includeDetails" json:"include_details,omitempty"`
	// *
	// Sort order for results. Valid values: "created_asc", "created_desc",
	// "last_activity_asc", "last_activity_desc". Defaults to "created_desc".
	SortOrder     *string `protobuf:"bytes,18,opt,name=sort_order,json=sortOrder" json:"sort_order,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListUserSessionsRequest) Reset() {
	*x = ListUserSessionsRequest{}
	mi := &file_pkg_auth_proto_requests_list_user_sessions_request_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListUserSessionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserSessionsRequest) ProtoMessage() {}

func (x *ListUserSessionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_auth_proto_requests_list_user_sessions_request_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserSessionsRequest.ProtoReflect.Descriptor instead.
func (*ListUserSessionsRequest) Descriptor() ([]byte, []int) {
	return file_pkg_auth_proto_requests_list_user_sessions_request_proto_rawDescGZIP(), []int{0}
}

func (x *ListUserSessionsRequest) GetUserId() string {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return ""
}

func (x *ListUserSessionsRequest) GetMetadata() *proto.RequestMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *ListUserSessionsRequest) GetPagination() *proto.Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *ListUserSessionsRequest) GetStatusFilter() string {
	if x != nil && x.StatusFilter != nil {
		return *x.StatusFilter
	}
	return ""
}

func (x *ListUserSessionsRequest) GetDeviceTypeFilter() string {
	if x != nil && x.DeviceTypeFilter != nil {
		return *x.DeviceTypeFilter
	}
	return ""
}

func (x *ListUserSessionsRequest) GetCreatedAfter() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAfter
	}
	return nil
}

func (x *ListUserSessionsRequest) GetCreatedBefore() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedBefore
	}
	return nil
}

func (x *ListUserSessionsRequest) GetIncludeDetails() bool {
	if x != nil && x.IncludeDetails != nil {
		return *x.IncludeDetails
	}
	return false
}

func (x *ListUserSessionsRequest) GetSortOrder() string {
	if x != nil && x.SortOrder != nil {
		return *x.SortOrder
	}
	return ""
}

var File_pkg_auth_proto_requests_list_user_sessions_request_proto protoreflect.FileDescriptor

const file_pkg_auth_proto_requests_list_user_sessions_request_proto_rawDesc = "" +
	"\n" +
	"8pkg/auth/proto/requests/list_user_sessions_request.proto\x12\x0fgcommon.v1.auth\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1dpkg/common/proto/common.proto\"\xd0\x03\n" +
	"\x17ListUserSessionsRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x12>\n" +
	"\bmetadata\x18\v \x01(\v2\".gcommon.v1.common.RequestMetadataR\bmetadata\x12=\n" +
	"\n" +
	"pagination\x18\f \x01(\v2\x1d.gcommon.v1.common.PaginationR\n" +
	"pagination\x12#\n" +
	"\rstatus_filter\x18\r \x01(\tR\fstatusFilter\x12,\n" +
	"\x12device_type_filter\x18\x0e \x01(\tR\x10deviceTypeFilter\x12?\n" +
	"\rcreated_after\x18\x0f \x01(\v2\x1a.google.protobuf.TimestampR\fcreatedAfter\x12A\n" +
	"\x0ecreated_before\x18\x10 \x01(\v2\x1a.google.protobuf.TimestampR\rcreatedBefore\x12'\n" +
	"\x0finclude_details\x18\x11 \x01(\bR\x0eincludeDetails\x12\x1d\n" +
	"\n" +
	"sort_order\x18\x12 \x01(\tR\tsortOrderB\xc2\x01\n" +
	"\x13com.gcommon.v1.authB\x1cListUserSessionsRequestProtoP\x01Z/github.com/jdfalk/gcommon/pkg/auth/proto;authpb\xa2\x02\x03GVA\xaa\x02\x0fGcommon.V1.Auth\xca\x02\x0fGcommon\\V1\\Auth\xe2\x02\x1bGcommon\\V1\\Auth\\GPBMetadata\xea\x02\x11Gcommon::V1::Authb\beditionsp\xe8\a"

var (
	file_pkg_auth_proto_requests_list_user_sessions_request_proto_rawDescOnce sync.Once
	file_pkg_auth_proto_requests_list_user_sessions_request_proto_rawDescData []byte
)

func file_pkg_auth_proto_requests_list_user_sessions_request_proto_rawDescGZIP() []byte {
	file_pkg_auth_proto_requests_list_user_sessions_request_proto_rawDescOnce.Do(func() {
		file_pkg_auth_proto_requests_list_user_sessions_request_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_pkg_auth_proto_requests_list_user_sessions_request_proto_rawDesc), len(file_pkg_auth_proto_requests_list_user_sessions_request_proto_rawDesc)))
	})
	return file_pkg_auth_proto_requests_list_user_sessions_request_proto_rawDescData
}

var file_pkg_auth_proto_requests_list_user_sessions_request_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pkg_auth_proto_requests_list_user_sessions_request_proto_goTypes = []any{
	(*ListUserSessionsRequest)(nil), // 0: gcommon.v1.auth.ListUserSessionsRequest
	(*proto.RequestMetadata)(nil),   // 1: gcommon.v1.common.RequestMetadata
	(*proto.Pagination)(nil),        // 2: gcommon.v1.common.Pagination
	(*timestamppb.Timestamp)(nil),   // 3: google.protobuf.Timestamp
}
var file_pkg_auth_proto_requests_list_user_sessions_request_proto_depIdxs = []int32{
	1, // 0: gcommon.v1.auth.ListUserSessionsRequest.metadata:type_name -> gcommon.v1.common.RequestMetadata
	2, // 1: gcommon.v1.auth.ListUserSessionsRequest.pagination:type_name -> gcommon.v1.common.Pagination
	3, // 2: gcommon.v1.auth.ListUserSessionsRequest.created_after:type_name -> google.protobuf.Timestamp
	3, // 3: gcommon.v1.auth.ListUserSessionsRequest.created_before:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_pkg_auth_proto_requests_list_user_sessions_request_proto_init() }
func file_pkg_auth_proto_requests_list_user_sessions_request_proto_init() {
	if File_pkg_auth_proto_requests_list_user_sessions_request_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pkg_auth_proto_requests_list_user_sessions_request_proto_rawDesc), len(file_pkg_auth_proto_requests_list_user_sessions_request_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_auth_proto_requests_list_user_sessions_request_proto_goTypes,
		DependencyIndexes: file_pkg_auth_proto_requests_list_user_sessions_request_proto_depIdxs,
		MessageInfos:      file_pkg_auth_proto_requests_list_user_sessions_request_proto_msgTypes,
	}.Build()
	File_pkg_auth_proto_requests_list_user_sessions_request_proto = out.File
	file_pkg_auth_proto_requests_list_user_sessions_request_proto_goTypes = nil
	file_pkg_auth_proto_requests_list_user_sessions_request_proto_depIdxs = nil
}

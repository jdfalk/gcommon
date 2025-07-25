// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: pkg/config/proto/requests/list_config_request.proto

//go:build protoopaque

package configpb

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
// ListConfigRequest lists configuration entries with optional filtering.
type ListConfigRequest struct {
	state                  protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Prefix      *string                `protobuf:"bytes,1,opt,name=prefix"`
	xxx_hidden_Namespace   *string                `protobuf:"bytes,2,opt,name=namespace"`
	xxx_hidden_Pagination  *proto.Pagination      `protobuf:"bytes,3,opt,name=pagination"`
	xxx_hidden_Filter      *proto.FilterOptions   `protobuf:"bytes,4,opt,name=filter"`
	xxx_hidden_Sort        *proto.SortOptions     `protobuf:"bytes,5,opt,name=sort"`
	xxx_hidden_Metadata    *proto.RequestMetadata `protobuf:"bytes,6,opt,name=metadata"`
	XXX_raceDetectHookData protoimpl.RaceDetectHookData
	XXX_presence           [1]uint32
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *ListConfigRequest) Reset() {
	*x = ListConfigRequest{}
	mi := &file_pkg_config_proto_requests_list_config_request_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListConfigRequest) ProtoMessage() {}

func (x *ListConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_config_proto_requests_list_config_request_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *ListConfigRequest) GetPrefix() string {
	if x != nil {
		if x.xxx_hidden_Prefix != nil {
			return *x.xxx_hidden_Prefix
		}
		return ""
	}
	return ""
}

func (x *ListConfigRequest) GetNamespace() string {
	if x != nil {
		if x.xxx_hidden_Namespace != nil {
			return *x.xxx_hidden_Namespace
		}
		return ""
	}
	return ""
}

func (x *ListConfigRequest) GetPagination() *proto.Pagination {
	if x != nil {
		return x.xxx_hidden_Pagination
	}
	return nil
}

func (x *ListConfigRequest) GetFilter() *proto.FilterOptions {
	if x != nil {
		return x.xxx_hidden_Filter
	}
	return nil
}

func (x *ListConfigRequest) GetSort() *proto.SortOptions {
	if x != nil {
		return x.xxx_hidden_Sort
	}
	return nil
}

func (x *ListConfigRequest) GetMetadata() *proto.RequestMetadata {
	if x != nil {
		return x.xxx_hidden_Metadata
	}
	return nil
}

func (x *ListConfigRequest) SetPrefix(v string) {
	x.xxx_hidden_Prefix = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 6)
}

func (x *ListConfigRequest) SetNamespace(v string) {
	x.xxx_hidden_Namespace = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 1, 6)
}

func (x *ListConfigRequest) SetPagination(v *proto.Pagination) {
	x.xxx_hidden_Pagination = v
}

func (x *ListConfigRequest) SetFilter(v *proto.FilterOptions) {
	x.xxx_hidden_Filter = v
}

func (x *ListConfigRequest) SetSort(v *proto.SortOptions) {
	x.xxx_hidden_Sort = v
}

func (x *ListConfigRequest) SetMetadata(v *proto.RequestMetadata) {
	x.xxx_hidden_Metadata = v
}

func (x *ListConfigRequest) HasPrefix() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *ListConfigRequest) HasNamespace() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 1)
}

func (x *ListConfigRequest) HasPagination() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Pagination != nil
}

func (x *ListConfigRequest) HasFilter() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Filter != nil
}

func (x *ListConfigRequest) HasSort() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Sort != nil
}

func (x *ListConfigRequest) HasMetadata() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Metadata != nil
}

func (x *ListConfigRequest) ClearPrefix() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_Prefix = nil
}

func (x *ListConfigRequest) ClearNamespace() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 1)
	x.xxx_hidden_Namespace = nil
}

func (x *ListConfigRequest) ClearPagination() {
	x.xxx_hidden_Pagination = nil
}

func (x *ListConfigRequest) ClearFilter() {
	x.xxx_hidden_Filter = nil
}

func (x *ListConfigRequest) ClearSort() {
	x.xxx_hidden_Sort = nil
}

func (x *ListConfigRequest) ClearMetadata() {
	x.xxx_hidden_Metadata = nil
}

type ListConfigRequest_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// Key prefix filter
	Prefix *string
	// Optional namespace/environment
	Namespace *string
	// Pagination options
	Pagination *proto.Pagination
	// Filter options
	Filter *proto.FilterOptions
	// Sort options
	Sort *proto.SortOptions
	// Request metadata
	Metadata *proto.RequestMetadata
}

func (b0 ListConfigRequest_builder) Build() *ListConfigRequest {
	m0 := &ListConfigRequest{}
	b, x := &b0, m0
	_, _ = b, x
	if b.Prefix != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 6)
		x.xxx_hidden_Prefix = b.Prefix
	}
	if b.Namespace != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 1, 6)
		x.xxx_hidden_Namespace = b.Namespace
	}
	x.xxx_hidden_Pagination = b.Pagination
	x.xxx_hidden_Filter = b.Filter
	x.xxx_hidden_Sort = b.Sort
	x.xxx_hidden_Metadata = b.Metadata
	return m0
}

var File_pkg_config_proto_requests_list_config_request_proto protoreflect.FileDescriptor

const file_pkg_config_proto_requests_list_config_request_proto_rawDesc = "" +
	"\n" +
	"3pkg/config/proto/requests/list_config_request.proto\x12\x11gcommon.v1.config\x1a!google/protobuf/go_features.proto\x1a*pkg/common/proto/messages/pagination.proto\x1a.pkg/common/proto/messages/filter_options.proto\x1a!pkg/common/proto/types/sort.proto\x1a0pkg/common/proto/messages/request_metadata.proto\"\xb6\x02\n" +
	"\x11ListConfigRequest\x12\x16\n" +
	"\x06prefix\x18\x01 \x01(\tR\x06prefix\x12\x1c\n" +
	"\tnamespace\x18\x02 \x01(\tR\tnamespace\x12=\n" +
	"\n" +
	"pagination\x18\x03 \x01(\v2\x1d.gcommon.v1.common.PaginationR\n" +
	"pagination\x128\n" +
	"\x06filter\x18\x04 \x01(\v2 .gcommon.v1.common.FilterOptionsR\x06filter\x122\n" +
	"\x04sort\x18\x05 \x01(\v2\x1e.gcommon.v1.common.SortOptionsR\x04sort\x12>\n" +
	"\bmetadata\x18\x06 \x01(\v2\".gcommon.v1.common.RequestMetadataR\bmetadataB\xd2\x01\n" +
	"\x15com.gcommon.v1.configB\x16ListConfigRequestProtoP\x01Z3github.com/jdfalk/gcommon/pkg/config/proto;configpb\xa2\x02\x03GVC\xaa\x02\x11Gcommon.V1.Config\xca\x02\x11Gcommon\\V1\\Config\xe2\x02\x1dGcommon\\V1\\Config\\GPBMetadata\xea\x02\x13Gcommon::V1::Config\x92\x03\x05\xd2>\x02\x10\x02b\beditionsp\xe8\a"

var file_pkg_config_proto_requests_list_config_request_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pkg_config_proto_requests_list_config_request_proto_goTypes = []any{
	(*ListConfigRequest)(nil),     // 0: gcommon.v1.config.ListConfigRequest
	(*proto.Pagination)(nil),      // 1: gcommon.v1.common.Pagination
	(*proto.FilterOptions)(nil),   // 2: gcommon.v1.common.FilterOptions
	(*proto.SortOptions)(nil),     // 3: gcommon.v1.common.SortOptions
	(*proto.RequestMetadata)(nil), // 4: gcommon.v1.common.RequestMetadata
}
var file_pkg_config_proto_requests_list_config_request_proto_depIdxs = []int32{
	1, // 0: gcommon.v1.config.ListConfigRequest.pagination:type_name -> gcommon.v1.common.Pagination
	2, // 1: gcommon.v1.config.ListConfigRequest.filter:type_name -> gcommon.v1.common.FilterOptions
	3, // 2: gcommon.v1.config.ListConfigRequest.sort:type_name -> gcommon.v1.common.SortOptions
	4, // 3: gcommon.v1.config.ListConfigRequest.metadata:type_name -> gcommon.v1.common.RequestMetadata
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_pkg_config_proto_requests_list_config_request_proto_init() }
func file_pkg_config_proto_requests_list_config_request_proto_init() {
	if File_pkg_config_proto_requests_list_config_request_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pkg_config_proto_requests_list_config_request_proto_rawDesc), len(file_pkg_config_proto_requests_list_config_request_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_config_proto_requests_list_config_request_proto_goTypes,
		DependencyIndexes: file_pkg_config_proto_requests_list_config_request_proto_depIdxs,
		MessageInfos:      file_pkg_config_proto_requests_list_config_request_proto_msgTypes,
	}.Build()
	File_pkg_config_proto_requests_list_config_request_proto = out.File
	file_pkg_config_proto_requests_list_config_request_proto_goTypes = nil
	file_pkg_config_proto_requests_list_config_request_proto_depIdxs = nil
}

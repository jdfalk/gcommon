// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: pkg/health/proto/responses/enable_check_response.proto

//go:build protoopaque

package healthpb

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
// Response message for enabling a health check.
// Contains the result of enabling a previously disabled check.
type EnableCheckResponse struct {
	state                  protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Success     bool                   `protobuf:"varint,1,opt,name=success"`
	xxx_hidden_CheckId     *string                `protobuf:"bytes,2,opt,name=check_id,json=checkId"`
	xxx_hidden_Error       *proto.Error           `protobuf:"bytes,3,opt,name=error"`
	xxx_hidden_Status      *string                `protobuf:"bytes,4,opt,name=status"`
	XXX_raceDetectHookData protoimpl.RaceDetectHookData
	XXX_presence           [1]uint32
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *EnableCheckResponse) Reset() {
	*x = EnableCheckResponse{}
	mi := &file_pkg_health_proto_responses_enable_check_response_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EnableCheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnableCheckResponse) ProtoMessage() {}

func (x *EnableCheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_health_proto_responses_enable_check_response_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *EnableCheckResponse) GetSuccess() bool {
	if x != nil {
		return x.xxx_hidden_Success
	}
	return false
}

func (x *EnableCheckResponse) GetCheckId() string {
	if x != nil {
		if x.xxx_hidden_CheckId != nil {
			return *x.xxx_hidden_CheckId
		}
		return ""
	}
	return ""
}

func (x *EnableCheckResponse) GetError() *proto.Error {
	if x != nil {
		return x.xxx_hidden_Error
	}
	return nil
}

func (x *EnableCheckResponse) GetStatus() string {
	if x != nil {
		if x.xxx_hidden_Status != nil {
			return *x.xxx_hidden_Status
		}
		return ""
	}
	return ""
}

func (x *EnableCheckResponse) SetSuccess(v bool) {
	x.xxx_hidden_Success = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 4)
}

func (x *EnableCheckResponse) SetCheckId(v string) {
	x.xxx_hidden_CheckId = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 1, 4)
}

func (x *EnableCheckResponse) SetError(v *proto.Error) {
	x.xxx_hidden_Error = v
}

func (x *EnableCheckResponse) SetStatus(v string) {
	x.xxx_hidden_Status = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 3, 4)
}

func (x *EnableCheckResponse) HasSuccess() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *EnableCheckResponse) HasCheckId() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 1)
}

func (x *EnableCheckResponse) HasError() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Error != nil
}

func (x *EnableCheckResponse) HasStatus() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 3)
}

func (x *EnableCheckResponse) ClearSuccess() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_Success = false
}

func (x *EnableCheckResponse) ClearCheckId() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 1)
	x.xxx_hidden_CheckId = nil
}

func (x *EnableCheckResponse) ClearError() {
	x.xxx_hidden_Error = nil
}

func (x *EnableCheckResponse) ClearStatus() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 3)
	x.xxx_hidden_Status = nil
}

type EnableCheckResponse_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// Success status
	Success *bool
	// Check ID that was enabled
	CheckId *string
	// Error information if enabling failed
	Error *proto.Error
	// Check status after enabling
	Status *string
}

func (b0 EnableCheckResponse_builder) Build() *EnableCheckResponse {
	m0 := &EnableCheckResponse{}
	b, x := &b0, m0
	_, _ = b, x
	if b.Success != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 4)
		x.xxx_hidden_Success = *b.Success
	}
	if b.CheckId != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 1, 4)
		x.xxx_hidden_CheckId = b.CheckId
	}
	x.xxx_hidden_Error = b.Error
	if b.Status != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 3, 4)
		x.xxx_hidden_Status = b.Status
	}
	return m0
}

var File_pkg_health_proto_responses_enable_check_response_proto protoreflect.FileDescriptor

const file_pkg_health_proto_responses_enable_check_response_proto_rawDesc = "" +
	"\n" +
	"6pkg/health/proto/responses/enable_check_response.proto\x12\x11gcommon.v1.health\x1a!google/protobuf/go_features.proto\x1a\x1dpkg/common/proto/common.proto\"\x92\x01\n" +
	"\x13EnableCheckResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\x12\x19\n" +
	"\bcheck_id\x18\x02 \x01(\tR\acheckId\x12.\n" +
	"\x05error\x18\x03 \x01(\v2\x18.gcommon.v1.common.ErrorR\x05error\x12\x16\n" +
	"\x06status\x18\x04 \x01(\tR\x06statusB\xd4\x01\n" +
	"\x15com.gcommon.v1.healthB\x18EnableCheckResponseProtoP\x01Z3github.com/jdfalk/gcommon/pkg/health/proto;healthpb\xa2\x02\x03GVH\xaa\x02\x11Gcommon.V1.Health\xca\x02\x11Gcommon\\V1\\Health\xe2\x02\x1dGcommon\\V1\\Health\\GPBMetadata\xea\x02\x13Gcommon::V1::Health\x92\x03\x05\xd2>\x02\x10\x02b\beditionsp\xe8\a"

var file_pkg_health_proto_responses_enable_check_response_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pkg_health_proto_responses_enable_check_response_proto_goTypes = []any{
	(*EnableCheckResponse)(nil), // 0: gcommon.v1.health.EnableCheckResponse
	(*proto.Error)(nil),         // 1: gcommon.v1.common.Error
}
var file_pkg_health_proto_responses_enable_check_response_proto_depIdxs = []int32{
	1, // 0: gcommon.v1.health.EnableCheckResponse.error:type_name -> gcommon.v1.common.Error
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_health_proto_responses_enable_check_response_proto_init() }
func file_pkg_health_proto_responses_enable_check_response_proto_init() {
	if File_pkg_health_proto_responses_enable_check_response_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pkg_health_proto_responses_enable_check_response_proto_rawDesc), len(file_pkg_health_proto_responses_enable_check_response_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_health_proto_responses_enable_check_response_proto_goTypes,
		DependencyIndexes: file_pkg_health_proto_responses_enable_check_response_proto_depIdxs,
		MessageInfos:      file_pkg_health_proto_responses_enable_check_response_proto_msgTypes,
	}.Build()
	File_pkg_health_proto_responses_enable_check_response_proto = out.File
	file_pkg_health_proto_responses_enable_check_response_proto_goTypes = nil
	file_pkg_health_proto_responses_enable_check_response_proto_depIdxs = nil
}

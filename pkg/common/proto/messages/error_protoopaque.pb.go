// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: pkg/common/proto/messages/error.proto

//go:build protoopaque

package commonpb

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
// Common error message structure for standardized error handling.
// Provides comprehensive error information including code, message,
// debugging details, and traceability across all GCommon modules.
type Error struct {
	state                  protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Code        ErrorCode              `protobuf:"varint,1,opt,name=code,enum=gcommon.v1.common.ErrorCode"`
	xxx_hidden_Message     *string                `protobuf:"bytes,2,opt,name=message"`
	xxx_hidden_Details     map[string]string      `protobuf:"bytes,3,rep,name=details" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	xxx_hidden_TraceId     *string                `protobuf:"bytes,4,opt,name=trace_id,json=traceId"`
	xxx_hidden_Timestamp   *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=timestamp"`
	xxx_hidden_Source      *string                `protobuf:"bytes,6,opt,name=source"`
	XXX_raceDetectHookData protoimpl.RaceDetectHookData
	XXX_presence           [1]uint32
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *Error) Reset() {
	*x = Error{}
	mi := &file_pkg_common_proto_messages_error_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_common_proto_messages_error_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *Error) GetCode() ErrorCode {
	if x != nil {
		if protoimpl.X.Present(&(x.XXX_presence[0]), 0) {
			return x.xxx_hidden_Code
		}
	}
	return ErrorCode_ERROR_CODE_UNSPECIFIED
}

func (x *Error) GetMessage() string {
	if x != nil {
		if x.xxx_hidden_Message != nil {
			return *x.xxx_hidden_Message
		}
		return ""
	}
	return ""
}

func (x *Error) GetDetails() map[string]string {
	if x != nil {
		return x.xxx_hidden_Details
	}
	return nil
}

func (x *Error) GetTraceId() string {
	if x != nil {
		if x.xxx_hidden_TraceId != nil {
			return *x.xxx_hidden_TraceId
		}
		return ""
	}
	return ""
}

func (x *Error) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.xxx_hidden_Timestamp
	}
	return nil
}

func (x *Error) GetSource() string {
	if x != nil {
		if x.xxx_hidden_Source != nil {
			return *x.xxx_hidden_Source
		}
		return ""
	}
	return ""
}

func (x *Error) SetCode(v ErrorCode) {
	x.xxx_hidden_Code = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 6)
}

func (x *Error) SetMessage(v string) {
	x.xxx_hidden_Message = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 1, 6)
}

func (x *Error) SetDetails(v map[string]string) {
	x.xxx_hidden_Details = v
}

func (x *Error) SetTraceId(v string) {
	x.xxx_hidden_TraceId = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 3, 6)
}

func (x *Error) SetTimestamp(v *timestamppb.Timestamp) {
	x.xxx_hidden_Timestamp = v
}

func (x *Error) SetSource(v string) {
	x.xxx_hidden_Source = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 5, 6)
}

func (x *Error) HasCode() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *Error) HasMessage() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 1)
}

func (x *Error) HasTraceId() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 3)
}

func (x *Error) HasTimestamp() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Timestamp != nil
}

func (x *Error) HasSource() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 5)
}

func (x *Error) ClearCode() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_Code = ErrorCode_ERROR_CODE_UNSPECIFIED
}

func (x *Error) ClearMessage() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 1)
	x.xxx_hidden_Message = nil
}

func (x *Error) ClearTraceId() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 3)
	x.xxx_hidden_TraceId = nil
}

func (x *Error) ClearTimestamp() {
	x.xxx_hidden_Timestamp = nil
}

func (x *Error) ClearSource() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 5)
	x.xxx_hidden_Source = nil
}

type Error_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// Standardized error code for programmatic handling
	Code *ErrorCode
	// Human-readable error message describing what went wrong
	Message *string
	// Additional error details for debugging and troubleshooting
	Details map[string]string
	// Distributed trace ID for request correlation across services
	TraceId *string
	// Timestamp when the error occurred
	Timestamp *timestamppb.Timestamp
	// Source module or component that generated the error
	Source *string
}

func (b0 Error_builder) Build() *Error {
	m0 := &Error{}
	b, x := &b0, m0
	_, _ = b, x
	if b.Code != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 6)
		x.xxx_hidden_Code = *b.Code
	}
	if b.Message != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 1, 6)
		x.xxx_hidden_Message = b.Message
	}
	x.xxx_hidden_Details = b.Details
	if b.TraceId != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 3, 6)
		x.xxx_hidden_TraceId = b.TraceId
	}
	x.xxx_hidden_Timestamp = b.Timestamp
	if b.Source != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 5, 6)
		x.xxx_hidden_Source = b.Source
	}
	return m0
}

var File_pkg_common_proto_messages_error_proto protoreflect.FileDescriptor

const file_pkg_common_proto_messages_error_proto_rawDesc = "" +
	"\n" +
	"%pkg/common/proto/messages/error.proto\x12\x11gcommon.v1.common\x1a'pkg/common/proto/enums/error_code.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a!google/protobuf/go_features.proto\"\xbd\x02\n" +
	"\x05Error\x120\n" +
	"\x04code\x18\x01 \x01(\x0e2\x1c.gcommon.v1.common.ErrorCodeR\x04code\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\x12?\n" +
	"\adetails\x18\x03 \x03(\v2%.gcommon.v1.common.Error.DetailsEntryR\adetails\x12\x19\n" +
	"\btrace_id\x18\x04 \x01(\tR\atraceId\x128\n" +
	"\ttimestamp\x18\x05 \x01(\v2\x1a.google.protobuf.TimestampR\ttimestamp\x12\x16\n" +
	"\x06source\x18\x06 \x01(\tR\x06source\x1a:\n" +
	"\fDetailsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01B\xc6\x01\n" +
	"\x15com.gcommon.v1.commonB\n" +
	"ErrorProtoP\x01Z3github.com/jdfalk/gcommon/pkg/common/proto;commonpb\xa2\x02\x03GVC\xaa\x02\x11Gcommon.V1.Common\xca\x02\x11Gcommon\\V1\\Common\xe2\x02\x1dGcommon\\V1\\Common\\GPBMetadata\xea\x02\x13Gcommon::V1::Common\x92\x03\x05\xd2>\x02\x10\x02b\beditionsp\xe8\a"

var file_pkg_common_proto_messages_error_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_common_proto_messages_error_proto_goTypes = []any{
	(*Error)(nil),                 // 0: gcommon.v1.common.Error
	nil,                           // 1: gcommon.v1.common.Error.DetailsEntry
	(ErrorCode)(0),                // 2: gcommon.v1.common.ErrorCode
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_pkg_common_proto_messages_error_proto_depIdxs = []int32{
	2, // 0: gcommon.v1.common.Error.code:type_name -> gcommon.v1.common.ErrorCode
	1, // 1: gcommon.v1.common.Error.details:type_name -> gcommon.v1.common.Error.DetailsEntry
	3, // 2: gcommon.v1.common.Error.timestamp:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pkg_common_proto_messages_error_proto_init() }
func file_pkg_common_proto_messages_error_proto_init() {
	if File_pkg_common_proto_messages_error_proto != nil {
		return
	}
	file_pkg_common_proto_enums_error_code_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pkg_common_proto_messages_error_proto_rawDesc), len(file_pkg_common_proto_messages_error_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_common_proto_messages_error_proto_goTypes,
		DependencyIndexes: file_pkg_common_proto_messages_error_proto_depIdxs,
		MessageInfos:      file_pkg_common_proto_messages_error_proto_msgTypes,
	}.Build()
	File_pkg_common_proto_messages_error_proto = out.File
	file_pkg_common_proto_messages_error_proto_goTypes = nil
	file_pkg_common_proto_messages_error_proto_depIdxs = nil
}

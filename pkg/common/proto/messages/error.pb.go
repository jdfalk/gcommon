// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: pkg/common/proto/messages/error.proto

//go:build !protoopaque

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
	state protoimpl.MessageState `protogen:"hybrid.v1"`
	// Standardized error code for programmatic handling
	Code *ErrorCode `protobuf:"varint,1,opt,name=code,enum=gcommon.v1.common.ErrorCode" json:"code,omitempty"`
	// Human-readable error message describing what went wrong
	Message *string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	// Additional error details for debugging and troubleshooting
	Details map[string]string `protobuf:"bytes,3,rep,name=details" json:"details,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Distributed trace ID for request correlation across services
	TraceId *string `protobuf:"bytes,4,opt,name=trace_id,json=traceId" json:"trace_id,omitempty"`
	// Timestamp when the error occurred
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=timestamp" json:"timestamp,omitempty"`
	// Source module or component that generated the error
	Source        *string `protobuf:"bytes,6,opt,name=source" json:"source,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
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
	if x != nil && x.Code != nil {
		return *x.Code
	}
	return ErrorCode_ERROR_CODE_UNSPECIFIED
}

func (x *Error) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

func (x *Error) GetDetails() map[string]string {
	if x != nil {
		return x.Details
	}
	return nil
}

func (x *Error) GetTraceId() string {
	if x != nil && x.TraceId != nil {
		return *x.TraceId
	}
	return ""
}

func (x *Error) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *Error) GetSource() string {
	if x != nil && x.Source != nil {
		return *x.Source
	}
	return ""
}

func (x *Error) SetCode(v ErrorCode) {
	x.Code = &v
}

func (x *Error) SetMessage(v string) {
	x.Message = &v
}

func (x *Error) SetDetails(v map[string]string) {
	x.Details = v
}

func (x *Error) SetTraceId(v string) {
	x.TraceId = &v
}

func (x *Error) SetTimestamp(v *timestamppb.Timestamp) {
	x.Timestamp = v
}

func (x *Error) SetSource(v string) {
	x.Source = &v
}

func (x *Error) HasCode() bool {
	if x == nil {
		return false
	}
	return x.Code != nil
}

func (x *Error) HasMessage() bool {
	if x == nil {
		return false
	}
	return x.Message != nil
}

func (x *Error) HasTraceId() bool {
	if x == nil {
		return false
	}
	return x.TraceId != nil
}

func (x *Error) HasTimestamp() bool {
	if x == nil {
		return false
	}
	return x.Timestamp != nil
}

func (x *Error) HasSource() bool {
	if x == nil {
		return false
	}
	return x.Source != nil
}

func (x *Error) ClearCode() {
	x.Code = nil
}

func (x *Error) ClearMessage() {
	x.Message = nil
}

func (x *Error) ClearTraceId() {
	x.TraceId = nil
}

func (x *Error) ClearTimestamp() {
	x.Timestamp = nil
}

func (x *Error) ClearSource() {
	x.Source = nil
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
	x.Code = b.Code
	x.Message = b.Message
	x.Details = b.Details
	x.TraceId = b.TraceId
	x.Timestamp = b.Timestamp
	x.Source = b.Source
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

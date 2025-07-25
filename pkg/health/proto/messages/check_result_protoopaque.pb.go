// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: pkg/health/proto/messages/check_result.proto

//go:build protoopaque

package healthpb

import (
	proto "github.com/jdfalk/gcommon/pkg/common/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/gofeaturespb"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
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
// Individual health check result for a specific component or subsystem.
// Provides detailed information about the health status of a single check.
type CheckResult struct {
	state                    protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Name          *string                `protobuf:"bytes,1,opt,name=name"`
	xxx_hidden_Status        proto.HealthStatus     `protobuf:"varint,2,opt,name=status,enum=gcommon.v1.common.HealthStatus"`
	xxx_hidden_Timestamp     *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=timestamp"`
	xxx_hidden_ExecutionTime *durationpb.Duration   `protobuf:"bytes,4,opt,name=execution_time,json=executionTime"`
	xxx_hidden_Message       *string                `protobuf:"bytes,5,opt,name=message"`
	xxx_hidden_Error         *proto.Error           `protobuf:"bytes,6,opt,name=error"`
	xxx_hidden_Metadata      map[string]string      `protobuf:"bytes,7,rep,name=metadata" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_raceDetectHookData   protoimpl.RaceDetectHookData
	XXX_presence             [1]uint32
	unknownFields            protoimpl.UnknownFields
	sizeCache                protoimpl.SizeCache
}

func (x *CheckResult) Reset() {
	*x = CheckResult{}
	mi := &file_pkg_health_proto_messages_check_result_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckResult) ProtoMessage() {}

func (x *CheckResult) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_health_proto_messages_check_result_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *CheckResult) GetName() string {
	if x != nil {
		if x.xxx_hidden_Name != nil {
			return *x.xxx_hidden_Name
		}
		return ""
	}
	return ""
}

func (x *CheckResult) GetStatus() proto.HealthStatus {
	if x != nil {
		if protoimpl.X.Present(&(x.XXX_presence[0]), 1) {
			return x.xxx_hidden_Status
		}
	}
	return proto.HealthStatus(0)
}

func (x *CheckResult) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.xxx_hidden_Timestamp
	}
	return nil
}

func (x *CheckResult) GetExecutionTime() *durationpb.Duration {
	if x != nil {
		return x.xxx_hidden_ExecutionTime
	}
	return nil
}

func (x *CheckResult) GetMessage() string {
	if x != nil {
		if x.xxx_hidden_Message != nil {
			return *x.xxx_hidden_Message
		}
		return ""
	}
	return ""
}

func (x *CheckResult) GetError() *proto.Error {
	if x != nil {
		return x.xxx_hidden_Error
	}
	return nil
}

func (x *CheckResult) GetMetadata() map[string]string {
	if x != nil {
		return x.xxx_hidden_Metadata
	}
	return nil
}

func (x *CheckResult) SetName(v string) {
	x.xxx_hidden_Name = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 7)
}

func (x *CheckResult) SetStatus(v proto.HealthStatus) {
	x.xxx_hidden_Status = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 1, 7)
}

func (x *CheckResult) SetTimestamp(v *timestamppb.Timestamp) {
	x.xxx_hidden_Timestamp = v
}

func (x *CheckResult) SetExecutionTime(v *durationpb.Duration) {
	x.xxx_hidden_ExecutionTime = v
}

func (x *CheckResult) SetMessage(v string) {
	x.xxx_hidden_Message = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 4, 7)
}

func (x *CheckResult) SetError(v *proto.Error) {
	x.xxx_hidden_Error = v
}

func (x *CheckResult) SetMetadata(v map[string]string) {
	x.xxx_hidden_Metadata = v
}

func (x *CheckResult) HasName() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *CheckResult) HasStatus() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 1)
}

func (x *CheckResult) HasTimestamp() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Timestamp != nil
}

func (x *CheckResult) HasExecutionTime() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_ExecutionTime != nil
}

func (x *CheckResult) HasMessage() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 4)
}

func (x *CheckResult) HasError() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Error != nil
}

func (x *CheckResult) ClearName() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_Name = nil
}

func (x *CheckResult) ClearStatus() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 1)
	x.xxx_hidden_Status = proto.HealthStatus_HEALTH_STATUS_UNSPECIFIED
}

func (x *CheckResult) ClearTimestamp() {
	x.xxx_hidden_Timestamp = nil
}

func (x *CheckResult) ClearExecutionTime() {
	x.xxx_hidden_ExecutionTime = nil
}

func (x *CheckResult) ClearMessage() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 4)
	x.xxx_hidden_Message = nil
}

func (x *CheckResult) ClearError() {
	x.xxx_hidden_Error = nil
}

type CheckResult_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// Check name or identifier
	Name *string
	// Health status of this specific check
	Status *proto.HealthStatus
	// Check execution timestamp
	Timestamp *timestamppb.Timestamp
	// Time taken to execute this check
	ExecutionTime *durationpb.Duration
	// Human-readable message about the check result
	Message *string
	// Error details if the check failed
	Error *proto.Error
	// Additional metadata for this check
	Metadata map[string]string
}

func (b0 CheckResult_builder) Build() *CheckResult {
	m0 := &CheckResult{}
	b, x := &b0, m0
	_, _ = b, x
	if b.Name != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 7)
		x.xxx_hidden_Name = b.Name
	}
	if b.Status != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 1, 7)
		x.xxx_hidden_Status = *b.Status
	}
	x.xxx_hidden_Timestamp = b.Timestamp
	x.xxx_hidden_ExecutionTime = b.ExecutionTime
	if b.Message != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 4, 7)
		x.xxx_hidden_Message = b.Message
	}
	x.xxx_hidden_Error = b.Error
	x.xxx_hidden_Metadata = b.Metadata
	return m0
}

var File_pkg_health_proto_messages_check_result_proto protoreflect.FileDescriptor

const file_pkg_health_proto_messages_check_result_proto_rawDesc = "" +
	"\n" +
	",pkg/health/proto/messages/check_result.proto\x12\x11gcommon.v1.health\x1a!google/protobuf/go_features.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1egoogle/protobuf/duration.proto\x1a\x1dpkg/common/proto/common.proto\"\xa7\x03\n" +
	"\vCheckResult\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x127\n" +
	"\x06status\x18\x02 \x01(\x0e2\x1f.gcommon.v1.common.HealthStatusR\x06status\x128\n" +
	"\ttimestamp\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampR\ttimestamp\x12@\n" +
	"\x0eexecution_time\x18\x04 \x01(\v2\x19.google.protobuf.DurationR\rexecutionTime\x12\x18\n" +
	"\amessage\x18\x05 \x01(\tR\amessage\x12.\n" +
	"\x05error\x18\x06 \x01(\v2\x18.gcommon.v1.common.ErrorR\x05error\x12H\n" +
	"\bmetadata\x18\a \x03(\v2,.gcommon.v1.health.CheckResult.MetadataEntryR\bmetadata\x1a;\n" +
	"\rMetadataEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01B\xcc\x01\n" +
	"\x15com.gcommon.v1.healthB\x10CheckResultProtoP\x01Z3github.com/jdfalk/gcommon/pkg/health/proto;healthpb\xa2\x02\x03GVH\xaa\x02\x11Gcommon.V1.Health\xca\x02\x11Gcommon\\V1\\Health\xe2\x02\x1dGcommon\\V1\\Health\\GPBMetadata\xea\x02\x13Gcommon::V1::Health\x92\x03\x05\xd2>\x02\x10\x02b\beditionsp\xe8\a"

var file_pkg_health_proto_messages_check_result_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_health_proto_messages_check_result_proto_goTypes = []any{
	(*CheckResult)(nil),           // 0: gcommon.v1.health.CheckResult
	nil,                           // 1: gcommon.v1.health.CheckResult.MetadataEntry
	(proto.HealthStatus)(0),       // 2: gcommon.v1.common.HealthStatus
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),   // 4: google.protobuf.Duration
	(*proto.Error)(nil),           // 5: gcommon.v1.common.Error
}
var file_pkg_health_proto_messages_check_result_proto_depIdxs = []int32{
	2, // 0: gcommon.v1.health.CheckResult.status:type_name -> gcommon.v1.common.HealthStatus
	3, // 1: gcommon.v1.health.CheckResult.timestamp:type_name -> google.protobuf.Timestamp
	4, // 2: gcommon.v1.health.CheckResult.execution_time:type_name -> google.protobuf.Duration
	5, // 3: gcommon.v1.health.CheckResult.error:type_name -> gcommon.v1.common.Error
	1, // 4: gcommon.v1.health.CheckResult.metadata:type_name -> gcommon.v1.health.CheckResult.MetadataEntry
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_pkg_health_proto_messages_check_result_proto_init() }
func file_pkg_health_proto_messages_check_result_proto_init() {
	if File_pkg_health_proto_messages_check_result_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pkg_health_proto_messages_check_result_proto_rawDesc), len(file_pkg_health_proto_messages_check_result_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_health_proto_messages_check_result_proto_goTypes,
		DependencyIndexes: file_pkg_health_proto_messages_check_result_proto_depIdxs,
		MessageInfos:      file_pkg_health_proto_messages_check_result_proto_msgTypes,
	}.Build()
	File_pkg_health_proto_messages_check_result_proto = out.File
	file_pkg_health_proto_messages_check_result_proto_goTypes = nil
	file_pkg_health_proto_messages_check_result_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: pkg/metrics/proto/responses/record_metric_response.proto

//go:build !protoopaque

package metricspb

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
// RecordMetricResponse contains the result of recording a metric data point.
type RecordMetricResponse struct {
	state protoimpl.MessageState `protogen:"hybrid.v1"`
	// Success status of the operation
	Success *bool `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	// Error information if the operation failed
	Error *proto.Error `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
	// Unique ID assigned to the recorded metric (if applicable)
	MetricId *string `protobuf:"bytes,3,opt,name=metric_id,json=metricId" json:"metric_id,omitempty"`
	// Timestamp when the metric was actually recorded
	RecordedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=recorded_at,json=recordedAt" json:"recorded_at,omitempty"`
	// Provider that handled the metric
	ProviderId *string `protobuf:"bytes,5,opt,name=provider_id,json=providerId" json:"provider_id,omitempty"`
	// Validation results if validation was requested
	Validation *ValidationResult `protobuf:"bytes,6,opt,name=validation" json:"validation,omitempty"`
	// Performance metrics about the recording operation
	Stats *RecordingStats `protobuf:"bytes,7,opt,name=stats" json:"stats,omitempty"`
	// Warnings or informational messages
	Warnings      []string `protobuf:"bytes,8,rep,name=warnings" json:"warnings,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RecordMetricResponse) Reset() {
	*x = RecordMetricResponse{}
	mi := &file_pkg_metrics_proto_responses_record_metric_response_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RecordMetricResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordMetricResponse) ProtoMessage() {}

func (x *RecordMetricResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_metrics_proto_responses_record_metric_response_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *RecordMetricResponse) GetSuccess() bool {
	if x != nil && x.Success != nil {
		return *x.Success
	}
	return false
}

func (x *RecordMetricResponse) GetError() *proto.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *RecordMetricResponse) GetMetricId() string {
	if x != nil && x.MetricId != nil {
		return *x.MetricId
	}
	return ""
}

func (x *RecordMetricResponse) GetRecordedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.RecordedAt
	}
	return nil
}

func (x *RecordMetricResponse) GetProviderId() string {
	if x != nil && x.ProviderId != nil {
		return *x.ProviderId
	}
	return ""
}

func (x *RecordMetricResponse) GetValidation() *ValidationResult {
	if x != nil {
		return x.Validation
	}
	return nil
}

func (x *RecordMetricResponse) GetStats() *RecordingStats {
	if x != nil {
		return x.Stats
	}
	return nil
}

func (x *RecordMetricResponse) GetWarnings() []string {
	if x != nil {
		return x.Warnings
	}
	return nil
}

func (x *RecordMetricResponse) SetSuccess(v bool) {
	x.Success = &v
}

func (x *RecordMetricResponse) SetError(v *proto.Error) {
	x.Error = v
}

func (x *RecordMetricResponse) SetMetricId(v string) {
	x.MetricId = &v
}

func (x *RecordMetricResponse) SetRecordedAt(v *timestamppb.Timestamp) {
	x.RecordedAt = v
}

func (x *RecordMetricResponse) SetProviderId(v string) {
	x.ProviderId = &v
}

func (x *RecordMetricResponse) SetValidation(v *ValidationResult) {
	x.Validation = v
}

func (x *RecordMetricResponse) SetStats(v *RecordingStats) {
	x.Stats = v
}

func (x *RecordMetricResponse) SetWarnings(v []string) {
	x.Warnings = v
}

func (x *RecordMetricResponse) HasSuccess() bool {
	if x == nil {
		return false
	}
	return x.Success != nil
}

func (x *RecordMetricResponse) HasError() bool {
	if x == nil {
		return false
	}
	return x.Error != nil
}

func (x *RecordMetricResponse) HasMetricId() bool {
	if x == nil {
		return false
	}
	return x.MetricId != nil
}

func (x *RecordMetricResponse) HasRecordedAt() bool {
	if x == nil {
		return false
	}
	return x.RecordedAt != nil
}

func (x *RecordMetricResponse) HasProviderId() bool {
	if x == nil {
		return false
	}
	return x.ProviderId != nil
}

func (x *RecordMetricResponse) HasValidation() bool {
	if x == nil {
		return false
	}
	return x.Validation != nil
}

func (x *RecordMetricResponse) HasStats() bool {
	if x == nil {
		return false
	}
	return x.Stats != nil
}

func (x *RecordMetricResponse) ClearSuccess() {
	x.Success = nil
}

func (x *RecordMetricResponse) ClearError() {
	x.Error = nil
}

func (x *RecordMetricResponse) ClearMetricId() {
	x.MetricId = nil
}

func (x *RecordMetricResponse) ClearRecordedAt() {
	x.RecordedAt = nil
}

func (x *RecordMetricResponse) ClearProviderId() {
	x.ProviderId = nil
}

func (x *RecordMetricResponse) ClearValidation() {
	x.Validation = nil
}

func (x *RecordMetricResponse) ClearStats() {
	x.Stats = nil
}

type RecordMetricResponse_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// Success status of the operation
	Success *bool
	// Error information if the operation failed
	Error *proto.Error
	// Unique ID assigned to the recorded metric (if applicable)
	MetricId *string
	// Timestamp when the metric was actually recorded
	RecordedAt *timestamppb.Timestamp
	// Provider that handled the metric
	ProviderId *string
	// Validation results if validation was requested
	Validation *ValidationResult
	// Performance metrics about the recording operation
	Stats *RecordingStats
	// Warnings or informational messages
	Warnings []string
}

func (b0 RecordMetricResponse_builder) Build() *RecordMetricResponse {
	m0 := &RecordMetricResponse{}
	b, x := &b0, m0
	_, _ = b, x
	x.Success = b.Success
	x.Error = b.Error
	x.MetricId = b.MetricId
	x.RecordedAt = b.RecordedAt
	x.ProviderId = b.ProviderId
	x.Validation = b.Validation
	x.Stats = b.Stats
	x.Warnings = b.Warnings
	return m0
}

var File_pkg_metrics_proto_responses_record_metric_response_proto protoreflect.FileDescriptor

const file_pkg_metrics_proto_responses_record_metric_response_proto_rawDesc = "" +
	"\n" +
	"8pkg/metrics/proto/responses/record_metric_response.proto\x12\x12gcommon.v1.metrics\x1a!google/protobuf/go_features.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a%pkg/common/proto/messages/error.proto\x1a/pkg/metrics/proto/types/validation_result.proto\x1a-pkg/metrics/proto/types/recording_stats.proto\"\xf7\x02\n" +
	"\x14RecordMetricResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\x12.\n" +
	"\x05error\x18\x02 \x01(\v2\x18.gcommon.v1.common.ErrorR\x05error\x12\x1b\n" +
	"\tmetric_id\x18\x03 \x01(\tR\bmetricId\x12;\n" +
	"\vrecorded_at\x18\x04 \x01(\v2\x1a.google.protobuf.TimestampR\n" +
	"recordedAt\x12\x1f\n" +
	"\vprovider_id\x18\x05 \x01(\tR\n" +
	"providerId\x12D\n" +
	"\n" +
	"validation\x18\x06 \x01(\v2$.gcommon.v1.metrics.ValidationResultR\n" +
	"validation\x128\n" +
	"\x05stats\x18\a \x01(\v2\".gcommon.v1.metrics.RecordingStatsR\x05stats\x12\x1a\n" +
	"\bwarnings\x18\b \x03(\tR\bwarningsB\xdc\x01\n" +
	"\x16com.gcommon.v1.metricsB\x19RecordMetricResponseProtoP\x01Z5github.com/jdfalk/gcommon/pkg/metrics/proto;metricspb\xa2\x02\x03GVM\xaa\x02\x12Gcommon.V1.Metrics\xca\x02\x12Gcommon\\V1\\Metrics\xe2\x02\x1eGcommon\\V1\\Metrics\\GPBMetadata\xea\x02\x14Gcommon::V1::Metrics\x92\x03\x05\xd2>\x02\x10\x02b\beditionsp\xe8\a"

var file_pkg_metrics_proto_responses_record_metric_response_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pkg_metrics_proto_responses_record_metric_response_proto_goTypes = []any{
	(*RecordMetricResponse)(nil),  // 0: gcommon.v1.metrics.RecordMetricResponse
	(*proto.Error)(nil),           // 1: gcommon.v1.common.Error
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
	(*ValidationResult)(nil),      // 3: gcommon.v1.metrics.ValidationResult
	(*RecordingStats)(nil),        // 4: gcommon.v1.metrics.RecordingStats
}
var file_pkg_metrics_proto_responses_record_metric_response_proto_depIdxs = []int32{
	1, // 0: gcommon.v1.metrics.RecordMetricResponse.error:type_name -> gcommon.v1.common.Error
	2, // 1: gcommon.v1.metrics.RecordMetricResponse.recorded_at:type_name -> google.protobuf.Timestamp
	3, // 2: gcommon.v1.metrics.RecordMetricResponse.validation:type_name -> gcommon.v1.metrics.ValidationResult
	4, // 3: gcommon.v1.metrics.RecordMetricResponse.stats:type_name -> gcommon.v1.metrics.RecordingStats
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_pkg_metrics_proto_responses_record_metric_response_proto_init() }
func file_pkg_metrics_proto_responses_record_metric_response_proto_init() {
	if File_pkg_metrics_proto_responses_record_metric_response_proto != nil {
		return
	}
	file_pkg_metrics_proto_types_validation_result_proto_init()
	file_pkg_metrics_proto_types_recording_stats_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pkg_metrics_proto_responses_record_metric_response_proto_rawDesc), len(file_pkg_metrics_proto_responses_record_metric_response_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_metrics_proto_responses_record_metric_response_proto_goTypes,
		DependencyIndexes: file_pkg_metrics_proto_responses_record_metric_response_proto_depIdxs,
		MessageInfos:      file_pkg_metrics_proto_responses_record_metric_response_proto_msgTypes,
	}.Build()
	File_pkg_metrics_proto_responses_record_metric_response_proto = out.File
	file_pkg_metrics_proto_responses_record_metric_response_proto_goTypes = nil
	file_pkg_metrics_proto_responses_record_metric_response_proto_depIdxs = nil
}

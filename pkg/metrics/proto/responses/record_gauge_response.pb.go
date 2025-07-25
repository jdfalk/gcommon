// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: pkg/metrics/proto/responses/record_gauge_response.proto

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
// RecordGaugeResponse is returned after recording a gauge metric.
type RecordGaugeResponse struct {
	state protoimpl.MessageState `protogen:"hybrid.v1"`
	// Whether the operation was successful
	Success *bool `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	// Error information if the operation failed
	Error *proto.Error `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
	// The recorded gauge metric with updated value
	Metric *GaugeMetric `protobuf:"bytes,3,opt,name=metric" json:"metric,omitempty"`
	// Previous value of the gauge (if applicable)
	PreviousValue *float64 `protobuf:"fixed64,4,opt,name=previous_value,json=previousValue" json:"previous_value,omitempty"`
	// Timestamp when the gauge was recorded
	RecordedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=recorded_at,json=recordedAt" json:"recorded_at,omitempty"`
	// Whether this was a new gauge or an update to existing
	IsNewMetric *bool `protobuf:"varint,6,opt,name=is_new_metric,json=isNewMetric" json:"is_new_metric,omitempty"`
	// Processing statistics
	Stats         *RecordingStats `protobuf:"bytes,7,opt,name=stats" json:"stats,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RecordGaugeResponse) Reset() {
	*x = RecordGaugeResponse{}
	mi := &file_pkg_metrics_proto_responses_record_gauge_response_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RecordGaugeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordGaugeResponse) ProtoMessage() {}

func (x *RecordGaugeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_metrics_proto_responses_record_gauge_response_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *RecordGaugeResponse) GetSuccess() bool {
	if x != nil && x.Success != nil {
		return *x.Success
	}
	return false
}

func (x *RecordGaugeResponse) GetError() *proto.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *RecordGaugeResponse) GetMetric() *GaugeMetric {
	if x != nil {
		return x.Metric
	}
	return nil
}

func (x *RecordGaugeResponse) GetPreviousValue() float64 {
	if x != nil && x.PreviousValue != nil {
		return *x.PreviousValue
	}
	return 0
}

func (x *RecordGaugeResponse) GetRecordedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.RecordedAt
	}
	return nil
}

func (x *RecordGaugeResponse) GetIsNewMetric() bool {
	if x != nil && x.IsNewMetric != nil {
		return *x.IsNewMetric
	}
	return false
}

func (x *RecordGaugeResponse) GetStats() *RecordingStats {
	if x != nil {
		return x.Stats
	}
	return nil
}

func (x *RecordGaugeResponse) SetSuccess(v bool) {
	x.Success = &v
}

func (x *RecordGaugeResponse) SetError(v *proto.Error) {
	x.Error = v
}

func (x *RecordGaugeResponse) SetMetric(v *GaugeMetric) {
	x.Metric = v
}

func (x *RecordGaugeResponse) SetPreviousValue(v float64) {
	x.PreviousValue = &v
}

func (x *RecordGaugeResponse) SetRecordedAt(v *timestamppb.Timestamp) {
	x.RecordedAt = v
}

func (x *RecordGaugeResponse) SetIsNewMetric(v bool) {
	x.IsNewMetric = &v
}

func (x *RecordGaugeResponse) SetStats(v *RecordingStats) {
	x.Stats = v
}

func (x *RecordGaugeResponse) HasSuccess() bool {
	if x == nil {
		return false
	}
	return x.Success != nil
}

func (x *RecordGaugeResponse) HasError() bool {
	if x == nil {
		return false
	}
	return x.Error != nil
}

func (x *RecordGaugeResponse) HasMetric() bool {
	if x == nil {
		return false
	}
	return x.Metric != nil
}

func (x *RecordGaugeResponse) HasPreviousValue() bool {
	if x == nil {
		return false
	}
	return x.PreviousValue != nil
}

func (x *RecordGaugeResponse) HasRecordedAt() bool {
	if x == nil {
		return false
	}
	return x.RecordedAt != nil
}

func (x *RecordGaugeResponse) HasIsNewMetric() bool {
	if x == nil {
		return false
	}
	return x.IsNewMetric != nil
}

func (x *RecordGaugeResponse) HasStats() bool {
	if x == nil {
		return false
	}
	return x.Stats != nil
}

func (x *RecordGaugeResponse) ClearSuccess() {
	x.Success = nil
}

func (x *RecordGaugeResponse) ClearError() {
	x.Error = nil
}

func (x *RecordGaugeResponse) ClearMetric() {
	x.Metric = nil
}

func (x *RecordGaugeResponse) ClearPreviousValue() {
	x.PreviousValue = nil
}

func (x *RecordGaugeResponse) ClearRecordedAt() {
	x.RecordedAt = nil
}

func (x *RecordGaugeResponse) ClearIsNewMetric() {
	x.IsNewMetric = nil
}

func (x *RecordGaugeResponse) ClearStats() {
	x.Stats = nil
}

type RecordGaugeResponse_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// Whether the operation was successful
	Success *bool
	// Error information if the operation failed
	Error *proto.Error
	// The recorded gauge metric with updated value
	Metric *GaugeMetric
	// Previous value of the gauge (if applicable)
	PreviousValue *float64
	// Timestamp when the gauge was recorded
	RecordedAt *timestamppb.Timestamp
	// Whether this was a new gauge or an update to existing
	IsNewMetric *bool
	// Processing statistics
	Stats *RecordingStats
}

func (b0 RecordGaugeResponse_builder) Build() *RecordGaugeResponse {
	m0 := &RecordGaugeResponse{}
	b, x := &b0, m0
	_, _ = b, x
	x.Success = b.Success
	x.Error = b.Error
	x.Metric = b.Metric
	x.PreviousValue = b.PreviousValue
	x.RecordedAt = b.RecordedAt
	x.IsNewMetric = b.IsNewMetric
	x.Stats = b.Stats
	return m0
}

var File_pkg_metrics_proto_responses_record_gauge_response_proto protoreflect.FileDescriptor

const file_pkg_metrics_proto_responses_record_gauge_response_proto_rawDesc = "" +
	"\n" +
	"7pkg/metrics/proto/responses/record_gauge_response.proto\x12\x12gcommon.v1.metrics\x1a!google/protobuf/go_features.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a%pkg/common/proto/messages/error.proto\x1a-pkg/metrics/proto/types/recording_stats.proto\x1a-pkg/metrics/proto/messages/gauge_metric.proto\"\xda\x02\n" +
	"\x13RecordGaugeResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\x12.\n" +
	"\x05error\x18\x02 \x01(\v2\x18.gcommon.v1.common.ErrorR\x05error\x127\n" +
	"\x06metric\x18\x03 \x01(\v2\x1f.gcommon.v1.metrics.GaugeMetricR\x06metric\x12%\n" +
	"\x0eprevious_value\x18\x04 \x01(\x01R\rpreviousValue\x12;\n" +
	"\vrecorded_at\x18\x05 \x01(\v2\x1a.google.protobuf.TimestampR\n" +
	"recordedAt\x12\"\n" +
	"\ris_new_metric\x18\x06 \x01(\bR\visNewMetric\x128\n" +
	"\x05stats\x18\a \x01(\v2\".gcommon.v1.metrics.RecordingStatsR\x05statsB\xdb\x01\n" +
	"\x16com.gcommon.v1.metricsB\x18RecordGaugeResponseProtoP\x01Z5github.com/jdfalk/gcommon/pkg/metrics/proto;metricspb\xa2\x02\x03GVM\xaa\x02\x12Gcommon.V1.Metrics\xca\x02\x12Gcommon\\V1\\Metrics\xe2\x02\x1eGcommon\\V1\\Metrics\\GPBMetadata\xea\x02\x14Gcommon::V1::Metrics\x92\x03\x05\xd2>\x02\x10\x02b\beditionsp\xe8\a"

var file_pkg_metrics_proto_responses_record_gauge_response_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pkg_metrics_proto_responses_record_gauge_response_proto_goTypes = []any{
	(*RecordGaugeResponse)(nil),   // 0: gcommon.v1.metrics.RecordGaugeResponse
	(*proto.Error)(nil),           // 1: gcommon.v1.common.Error
	(*GaugeMetric)(nil),           // 2: gcommon.v1.metrics.GaugeMetric
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(*RecordingStats)(nil),        // 4: gcommon.v1.metrics.RecordingStats
}
var file_pkg_metrics_proto_responses_record_gauge_response_proto_depIdxs = []int32{
	1, // 0: gcommon.v1.metrics.RecordGaugeResponse.error:type_name -> gcommon.v1.common.Error
	2, // 1: gcommon.v1.metrics.RecordGaugeResponse.metric:type_name -> gcommon.v1.metrics.GaugeMetric
	3, // 2: gcommon.v1.metrics.RecordGaugeResponse.recorded_at:type_name -> google.protobuf.Timestamp
	4, // 3: gcommon.v1.metrics.RecordGaugeResponse.stats:type_name -> gcommon.v1.metrics.RecordingStats
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_pkg_metrics_proto_responses_record_gauge_response_proto_init() }
func file_pkg_metrics_proto_responses_record_gauge_response_proto_init() {
	if File_pkg_metrics_proto_responses_record_gauge_response_proto != nil {
		return
	}
	file_pkg_metrics_proto_types_recording_stats_proto_init()
	file_pkg_metrics_proto_messages_gauge_metric_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pkg_metrics_proto_responses_record_gauge_response_proto_rawDesc), len(file_pkg_metrics_proto_responses_record_gauge_response_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_metrics_proto_responses_record_gauge_response_proto_goTypes,
		DependencyIndexes: file_pkg_metrics_proto_responses_record_gauge_response_proto_depIdxs,
		MessageInfos:      file_pkg_metrics_proto_responses_record_gauge_response_proto_msgTypes,
	}.Build()
	File_pkg_metrics_proto_responses_record_gauge_response_proto = out.File
	file_pkg_metrics_proto_responses_record_gauge_response_proto_goTypes = nil
	file_pkg_metrics_proto_responses_record_gauge_response_proto_depIdxs = nil
}

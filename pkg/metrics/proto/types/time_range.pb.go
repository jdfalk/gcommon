// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: pkg/metrics/proto/types/time_range.proto

package metricspb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
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
// Time range specification for metrics queries.
type TimeRange struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Start time for the range
	StartTime *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
	// End time for the range
	EndTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=end_time,json=endTime" json:"end_time,omitempty"`
	// Duration of the time range (alternative to end_time)
	Duration      *durationpb.Duration `protobuf:"bytes,3,opt,name=duration" json:"duration,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TimeRange) Reset() {
	*x = TimeRange{}
	mi := &file_pkg_metrics_proto_types_time_range_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TimeRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeRange) ProtoMessage() {}

func (x *TimeRange) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_metrics_proto_types_time_range_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeRange.ProtoReflect.Descriptor instead.
func (*TimeRange) Descriptor() ([]byte, []int) {
	return file_pkg_metrics_proto_types_time_range_proto_rawDescGZIP(), []int{0}
}

func (x *TimeRange) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *TimeRange) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *TimeRange) GetDuration() *durationpb.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

var File_pkg_metrics_proto_types_time_range_proto protoreflect.FileDescriptor

const file_pkg_metrics_proto_types_time_range_proto_rawDesc = "" +
	"\n" +
	"(pkg/metrics/proto/types/time_range.proto\x12\x12gcommon.v1.metrics\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1egoogle/protobuf/duration.proto\"\xb4\x01\n" +
	"\tTimeRange\x129\n" +
	"\n" +
	"start_time\x18\x01 \x01(\v2\x1a.google.protobuf.TimestampR\tstartTime\x125\n" +
	"\bend_time\x18\x02 \x01(\v2\x1a.google.protobuf.TimestampR\aendTime\x125\n" +
	"\bduration\x18\x03 \x01(\v2\x19.google.protobuf.DurationR\bdurationB\xc9\x01\n" +
	"\x16com.gcommon.v1.metricsB\x0eTimeRangeProtoP\x01Z5github.com/jdfalk/gcommon/pkg/metrics/proto;metricspb\xa2\x02\x03GVM\xaa\x02\x12Gcommon.V1.Metrics\xca\x02\x12Gcommon\\V1\\Metrics\xe2\x02\x1eGcommon\\V1\\Metrics\\GPBMetadata\xea\x02\x14Gcommon::V1::Metricsb\beditionsp\xe8\a"

var (
	file_pkg_metrics_proto_types_time_range_proto_rawDescOnce sync.Once
	file_pkg_metrics_proto_types_time_range_proto_rawDescData []byte
)

func file_pkg_metrics_proto_types_time_range_proto_rawDescGZIP() []byte {
	file_pkg_metrics_proto_types_time_range_proto_rawDescOnce.Do(func() {
		file_pkg_metrics_proto_types_time_range_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_pkg_metrics_proto_types_time_range_proto_rawDesc), len(file_pkg_metrics_proto_types_time_range_proto_rawDesc)))
	})
	return file_pkg_metrics_proto_types_time_range_proto_rawDescData
}

var file_pkg_metrics_proto_types_time_range_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pkg_metrics_proto_types_time_range_proto_goTypes = []any{
	(*TimeRange)(nil),             // 0: gcommon.v1.metrics.TimeRange
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),   // 2: google.protobuf.Duration
}
var file_pkg_metrics_proto_types_time_range_proto_depIdxs = []int32{
	1, // 0: gcommon.v1.metrics.TimeRange.start_time:type_name -> google.protobuf.Timestamp
	1, // 1: gcommon.v1.metrics.TimeRange.end_time:type_name -> google.protobuf.Timestamp
	2, // 2: gcommon.v1.metrics.TimeRange.duration:type_name -> google.protobuf.Duration
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pkg_metrics_proto_types_time_range_proto_init() }
func file_pkg_metrics_proto_types_time_range_proto_init() {
	if File_pkg_metrics_proto_types_time_range_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pkg_metrics_proto_types_time_range_proto_rawDesc), len(file_pkg_metrics_proto_types_time_range_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_metrics_proto_types_time_range_proto_goTypes,
		DependencyIndexes: file_pkg_metrics_proto_types_time_range_proto_depIdxs,
		MessageInfos:      file_pkg_metrics_proto_types_time_range_proto_msgTypes,
	}.Build()
	File_pkg_metrics_proto_types_time_range_proto = out.File
	file_pkg_metrics_proto_types_time_range_proto_goTypes = nil
	file_pkg_metrics_proto_types_time_range_proto_depIdxs = nil
}

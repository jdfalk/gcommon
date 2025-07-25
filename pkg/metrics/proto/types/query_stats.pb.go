// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: pkg/metrics/proto/types/query_stats.proto

//go:build !protoopaque

package metricspb

import (
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
// QueryStats contains performance information about the query.
type QueryStats struct {
	state protoimpl.MessageState `protogen:"hybrid.v1"`
	// Time taken to execute the query (milliseconds)
	ExecutionTimeMs *int64 `protobuf:"varint,1,opt,name=execution_time_ms,json=executionTimeMs" json:"execution_time_ms,omitempty"`
	// Number of data points examined
	DataPointsExamined *int64 `protobuf:"varint,2,opt,name=data_points_examined,json=dataPointsExamined" json:"data_points_examined,omitempty"`
	// Number of metrics returned
	MetricsReturned *int64 `protobuf:"varint,3,opt,name=metrics_returned,json=metricsReturned" json:"metrics_returned,omitempty"`
	// Number of data points returned
	DataPointsReturned *int64 `protobuf:"varint,4,opt,name=data_points_returned,json=dataPointsReturned" json:"data_points_returned,omitempty"`
	// Cache hit rate for the query (0.0 to 1.0)
	CacheHitRate *float64 `protobuf:"fixed64,5,opt,name=cache_hit_rate,json=cacheHitRate" json:"cache_hit_rate,omitempty"`
	// Memory used during query execution (bytes)
	MemoryUsedBytes *int64 `protobuf:"varint,6,opt,name=memory_used_bytes,json=memoryUsedBytes" json:"memory_used_bytes,omitempty"`
	// Storage backends queried
	StorageBackends []string `protobuf:"bytes,7,rep,name=storage_backends,json=storageBackends" json:"storage_backends,omitempty"`
	// Whether the query was optimized
	QueryOptimized *bool `protobuf:"varint,8,opt,name=query_optimized,json=queryOptimized" json:"query_optimized,omitempty"`
	// Query optimization details
	OptimizationDetails *string `protobuf:"bytes,9,opt,name=optimization_details,json=optimizationDetails" json:"optimization_details,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *QueryStats) Reset() {
	*x = QueryStats{}
	mi := &file_pkg_metrics_proto_types_query_stats_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *QueryStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryStats) ProtoMessage() {}

func (x *QueryStats) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_metrics_proto_types_query_stats_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *QueryStats) GetExecutionTimeMs() int64 {
	if x != nil && x.ExecutionTimeMs != nil {
		return *x.ExecutionTimeMs
	}
	return 0
}

func (x *QueryStats) GetDataPointsExamined() int64 {
	if x != nil && x.DataPointsExamined != nil {
		return *x.DataPointsExamined
	}
	return 0
}

func (x *QueryStats) GetMetricsReturned() int64 {
	if x != nil && x.MetricsReturned != nil {
		return *x.MetricsReturned
	}
	return 0
}

func (x *QueryStats) GetDataPointsReturned() int64 {
	if x != nil && x.DataPointsReturned != nil {
		return *x.DataPointsReturned
	}
	return 0
}

func (x *QueryStats) GetCacheHitRate() float64 {
	if x != nil && x.CacheHitRate != nil {
		return *x.CacheHitRate
	}
	return 0
}

func (x *QueryStats) GetMemoryUsedBytes() int64 {
	if x != nil && x.MemoryUsedBytes != nil {
		return *x.MemoryUsedBytes
	}
	return 0
}

func (x *QueryStats) GetStorageBackends() []string {
	if x != nil {
		return x.StorageBackends
	}
	return nil
}

func (x *QueryStats) GetQueryOptimized() bool {
	if x != nil && x.QueryOptimized != nil {
		return *x.QueryOptimized
	}
	return false
}

func (x *QueryStats) GetOptimizationDetails() string {
	if x != nil && x.OptimizationDetails != nil {
		return *x.OptimizationDetails
	}
	return ""
}

func (x *QueryStats) SetExecutionTimeMs(v int64) {
	x.ExecutionTimeMs = &v
}

func (x *QueryStats) SetDataPointsExamined(v int64) {
	x.DataPointsExamined = &v
}

func (x *QueryStats) SetMetricsReturned(v int64) {
	x.MetricsReturned = &v
}

func (x *QueryStats) SetDataPointsReturned(v int64) {
	x.DataPointsReturned = &v
}

func (x *QueryStats) SetCacheHitRate(v float64) {
	x.CacheHitRate = &v
}

func (x *QueryStats) SetMemoryUsedBytes(v int64) {
	x.MemoryUsedBytes = &v
}

func (x *QueryStats) SetStorageBackends(v []string) {
	x.StorageBackends = v
}

func (x *QueryStats) SetQueryOptimized(v bool) {
	x.QueryOptimized = &v
}

func (x *QueryStats) SetOptimizationDetails(v string) {
	x.OptimizationDetails = &v
}

func (x *QueryStats) HasExecutionTimeMs() bool {
	if x == nil {
		return false
	}
	return x.ExecutionTimeMs != nil
}

func (x *QueryStats) HasDataPointsExamined() bool {
	if x == nil {
		return false
	}
	return x.DataPointsExamined != nil
}

func (x *QueryStats) HasMetricsReturned() bool {
	if x == nil {
		return false
	}
	return x.MetricsReturned != nil
}

func (x *QueryStats) HasDataPointsReturned() bool {
	if x == nil {
		return false
	}
	return x.DataPointsReturned != nil
}

func (x *QueryStats) HasCacheHitRate() bool {
	if x == nil {
		return false
	}
	return x.CacheHitRate != nil
}

func (x *QueryStats) HasMemoryUsedBytes() bool {
	if x == nil {
		return false
	}
	return x.MemoryUsedBytes != nil
}

func (x *QueryStats) HasQueryOptimized() bool {
	if x == nil {
		return false
	}
	return x.QueryOptimized != nil
}

func (x *QueryStats) HasOptimizationDetails() bool {
	if x == nil {
		return false
	}
	return x.OptimizationDetails != nil
}

func (x *QueryStats) ClearExecutionTimeMs() {
	x.ExecutionTimeMs = nil
}

func (x *QueryStats) ClearDataPointsExamined() {
	x.DataPointsExamined = nil
}

func (x *QueryStats) ClearMetricsReturned() {
	x.MetricsReturned = nil
}

func (x *QueryStats) ClearDataPointsReturned() {
	x.DataPointsReturned = nil
}

func (x *QueryStats) ClearCacheHitRate() {
	x.CacheHitRate = nil
}

func (x *QueryStats) ClearMemoryUsedBytes() {
	x.MemoryUsedBytes = nil
}

func (x *QueryStats) ClearQueryOptimized() {
	x.QueryOptimized = nil
}

func (x *QueryStats) ClearOptimizationDetails() {
	x.OptimizationDetails = nil
}

type QueryStats_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// Time taken to execute the query (milliseconds)
	ExecutionTimeMs *int64
	// Number of data points examined
	DataPointsExamined *int64
	// Number of metrics returned
	MetricsReturned *int64
	// Number of data points returned
	DataPointsReturned *int64
	// Cache hit rate for the query (0.0 to 1.0)
	CacheHitRate *float64
	// Memory used during query execution (bytes)
	MemoryUsedBytes *int64
	// Storage backends queried
	StorageBackends []string
	// Whether the query was optimized
	QueryOptimized *bool
	// Query optimization details
	OptimizationDetails *string
}

func (b0 QueryStats_builder) Build() *QueryStats {
	m0 := &QueryStats{}
	b, x := &b0, m0
	_, _ = b, x
	x.ExecutionTimeMs = b.ExecutionTimeMs
	x.DataPointsExamined = b.DataPointsExamined
	x.MetricsReturned = b.MetricsReturned
	x.DataPointsReturned = b.DataPointsReturned
	x.CacheHitRate = b.CacheHitRate
	x.MemoryUsedBytes = b.MemoryUsedBytes
	x.StorageBackends = b.StorageBackends
	x.QueryOptimized = b.QueryOptimized
	x.OptimizationDetails = b.OptimizationDetails
	return m0
}

var File_pkg_metrics_proto_types_query_stats_proto protoreflect.FileDescriptor

const file_pkg_metrics_proto_types_query_stats_proto_rawDesc = "" +
	"\n" +
	")pkg/metrics/proto/types/query_stats.proto\x12\x12gcommon.v1.metrics\x1a!google/protobuf/go_features.proto\"\xa0\x03\n" +
	"\n" +
	"QueryStats\x12*\n" +
	"\x11execution_time_ms\x18\x01 \x01(\x03R\x0fexecutionTimeMs\x120\n" +
	"\x14data_points_examined\x18\x02 \x01(\x03R\x12dataPointsExamined\x12)\n" +
	"\x10metrics_returned\x18\x03 \x01(\x03R\x0fmetricsReturned\x120\n" +
	"\x14data_points_returned\x18\x04 \x01(\x03R\x12dataPointsReturned\x12$\n" +
	"\x0ecache_hit_rate\x18\x05 \x01(\x01R\fcacheHitRate\x12*\n" +
	"\x11memory_used_bytes\x18\x06 \x01(\x03R\x0fmemoryUsedBytes\x12)\n" +
	"\x10storage_backends\x18\a \x03(\tR\x0fstorageBackends\x12'\n" +
	"\x0fquery_optimized\x18\b \x01(\bR\x0equeryOptimized\x121\n" +
	"\x14optimization_details\x18\t \x01(\tR\x13optimizationDetailsB\xd2\x01\n" +
	"\x16com.gcommon.v1.metricsB\x0fQueryStatsProtoP\x01Z5github.com/jdfalk/gcommon/pkg/metrics/proto;metricspb\xa2\x02\x03GVM\xaa\x02\x12Gcommon.V1.Metrics\xca\x02\x12Gcommon\\V1\\Metrics\xe2\x02\x1eGcommon\\V1\\Metrics\\GPBMetadata\xea\x02\x14Gcommon::V1::Metrics\x92\x03\x05\xd2>\x02\x10\x02b\beditionsp\xe8\a"

var file_pkg_metrics_proto_types_query_stats_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pkg_metrics_proto_types_query_stats_proto_goTypes = []any{
	(*QueryStats)(nil), // 0: gcommon.v1.metrics.QueryStats
}
var file_pkg_metrics_proto_types_query_stats_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_metrics_proto_types_query_stats_proto_init() }
func file_pkg_metrics_proto_types_query_stats_proto_init() {
	if File_pkg_metrics_proto_types_query_stats_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pkg_metrics_proto_types_query_stats_proto_rawDesc), len(file_pkg_metrics_proto_types_query_stats_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_metrics_proto_types_query_stats_proto_goTypes,
		DependencyIndexes: file_pkg_metrics_proto_types_query_stats_proto_depIdxs,
		MessageInfos:      file_pkg_metrics_proto_types_query_stats_proto_msgTypes,
	}.Build()
	File_pkg_metrics_proto_types_query_stats_proto = out.File
	file_pkg_metrics_proto_types_query_stats_proto_goTypes = nil
	file_pkg_metrics_proto_types_query_stats_proto_depIdxs = nil
}

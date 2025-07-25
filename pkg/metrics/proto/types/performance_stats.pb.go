// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: pkg/metrics/proto/types/performance_stats.proto

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
// PerformanceStats contains performance metrics for provider operations.
type PerformanceStats struct {
	state protoimpl.MessageState `protogen:"hybrid.v1"`
	// Average response time in milliseconds
	AvgResponseTimeMs *float64 `protobuf:"fixed64,1,opt,name=avg_response_time_ms,json=avgResponseTimeMs" json:"avg_response_time_ms,omitempty"`
	// Maximum response time in milliseconds
	MaxResponseTimeMs *float64 `protobuf:"fixed64,2,opt,name=max_response_time_ms,json=maxResponseTimeMs" json:"max_response_time_ms,omitempty"`
	// Minimum response time in milliseconds
	MinResponseTimeMs *float64 `protobuf:"fixed64,3,opt,name=min_response_time_ms,json=minResponseTimeMs" json:"min_response_time_ms,omitempty"`
	// 95th percentile response time in milliseconds
	P95ResponseTimeMs *float64 `protobuf:"fixed64,4,opt,name=p95_response_time_ms,json=p95ResponseTimeMs" json:"p95_response_time_ms,omitempty"`
	// 99th percentile response time in milliseconds
	P99ResponseTimeMs *float64 `protobuf:"fixed64,5,opt,name=p99_response_time_ms,json=p99ResponseTimeMs" json:"p99_response_time_ms,omitempty"`
	// Requests per second
	RequestsPerSecond *float64 `protobuf:"fixed64,6,opt,name=requests_per_second,json=requestsPerSecond" json:"requests_per_second,omitempty"`
	// Total number of requests processed
	TotalRequests *uint64 `protobuf:"varint,7,opt,name=total_requests,json=totalRequests" json:"total_requests,omitempty"`
	// Number of successful requests
	SuccessfulRequests *uint64 `protobuf:"varint,8,opt,name=successful_requests,json=successfulRequests" json:"successful_requests,omitempty"`
	// Number of failed requests
	FailedRequests *uint64 `protobuf:"varint,9,opt,name=failed_requests,json=failedRequests" json:"failed_requests,omitempty"`
	// Success rate (0.0 to 1.0)
	SuccessRate *float64 `protobuf:"fixed64,10,opt,name=success_rate,json=successRate" json:"success_rate,omitempty"`
	// CPU utilization percentage (0.0 to 100.0)
	CpuUtilization *float64 `protobuf:"fixed64,11,opt,name=cpu_utilization,json=cpuUtilization" json:"cpu_utilization,omitempty"`
	// Memory utilization percentage (0.0 to 100.0)
	MemoryUtilization *float64 `protobuf:"fixed64,12,opt,name=memory_utilization,json=memoryUtilization" json:"memory_utilization,omitempty"`
	// Network I/O in bytes per second
	NetworkIoBytesPerSec *float64 `protobuf:"fixed64,13,opt,name=network_io_bytes_per_sec,json=networkIoBytesPerSec" json:"network_io_bytes_per_sec,omitempty"`
	// Disk I/O in bytes per second
	DiskIoBytesPerSec *float64 `protobuf:"fixed64,14,opt,name=disk_io_bytes_per_sec,json=diskIoBytesPerSec" json:"disk_io_bytes_per_sec,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *PerformanceStats) Reset() {
	*x = PerformanceStats{}
	mi := &file_pkg_metrics_proto_types_performance_stats_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PerformanceStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PerformanceStats) ProtoMessage() {}

func (x *PerformanceStats) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_metrics_proto_types_performance_stats_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *PerformanceStats) GetAvgResponseTimeMs() float64 {
	if x != nil && x.AvgResponseTimeMs != nil {
		return *x.AvgResponseTimeMs
	}
	return 0
}

func (x *PerformanceStats) GetMaxResponseTimeMs() float64 {
	if x != nil && x.MaxResponseTimeMs != nil {
		return *x.MaxResponseTimeMs
	}
	return 0
}

func (x *PerformanceStats) GetMinResponseTimeMs() float64 {
	if x != nil && x.MinResponseTimeMs != nil {
		return *x.MinResponseTimeMs
	}
	return 0
}

func (x *PerformanceStats) GetP95ResponseTimeMs() float64 {
	if x != nil && x.P95ResponseTimeMs != nil {
		return *x.P95ResponseTimeMs
	}
	return 0
}

func (x *PerformanceStats) GetP99ResponseTimeMs() float64 {
	if x != nil && x.P99ResponseTimeMs != nil {
		return *x.P99ResponseTimeMs
	}
	return 0
}

func (x *PerformanceStats) GetRequestsPerSecond() float64 {
	if x != nil && x.RequestsPerSecond != nil {
		return *x.RequestsPerSecond
	}
	return 0
}

func (x *PerformanceStats) GetTotalRequests() uint64 {
	if x != nil && x.TotalRequests != nil {
		return *x.TotalRequests
	}
	return 0
}

func (x *PerformanceStats) GetSuccessfulRequests() uint64 {
	if x != nil && x.SuccessfulRequests != nil {
		return *x.SuccessfulRequests
	}
	return 0
}

func (x *PerformanceStats) GetFailedRequests() uint64 {
	if x != nil && x.FailedRequests != nil {
		return *x.FailedRequests
	}
	return 0
}

func (x *PerformanceStats) GetSuccessRate() float64 {
	if x != nil && x.SuccessRate != nil {
		return *x.SuccessRate
	}
	return 0
}

func (x *PerformanceStats) GetCpuUtilization() float64 {
	if x != nil && x.CpuUtilization != nil {
		return *x.CpuUtilization
	}
	return 0
}

func (x *PerformanceStats) GetMemoryUtilization() float64 {
	if x != nil && x.MemoryUtilization != nil {
		return *x.MemoryUtilization
	}
	return 0
}

func (x *PerformanceStats) GetNetworkIoBytesPerSec() float64 {
	if x != nil && x.NetworkIoBytesPerSec != nil {
		return *x.NetworkIoBytesPerSec
	}
	return 0
}

func (x *PerformanceStats) GetDiskIoBytesPerSec() float64 {
	if x != nil && x.DiskIoBytesPerSec != nil {
		return *x.DiskIoBytesPerSec
	}
	return 0
}

func (x *PerformanceStats) SetAvgResponseTimeMs(v float64) {
	x.AvgResponseTimeMs = &v
}

func (x *PerformanceStats) SetMaxResponseTimeMs(v float64) {
	x.MaxResponseTimeMs = &v
}

func (x *PerformanceStats) SetMinResponseTimeMs(v float64) {
	x.MinResponseTimeMs = &v
}

func (x *PerformanceStats) SetP95ResponseTimeMs(v float64) {
	x.P95ResponseTimeMs = &v
}

func (x *PerformanceStats) SetP99ResponseTimeMs(v float64) {
	x.P99ResponseTimeMs = &v
}

func (x *PerformanceStats) SetRequestsPerSecond(v float64) {
	x.RequestsPerSecond = &v
}

func (x *PerformanceStats) SetTotalRequests(v uint64) {
	x.TotalRequests = &v
}

func (x *PerformanceStats) SetSuccessfulRequests(v uint64) {
	x.SuccessfulRequests = &v
}

func (x *PerformanceStats) SetFailedRequests(v uint64) {
	x.FailedRequests = &v
}

func (x *PerformanceStats) SetSuccessRate(v float64) {
	x.SuccessRate = &v
}

func (x *PerformanceStats) SetCpuUtilization(v float64) {
	x.CpuUtilization = &v
}

func (x *PerformanceStats) SetMemoryUtilization(v float64) {
	x.MemoryUtilization = &v
}

func (x *PerformanceStats) SetNetworkIoBytesPerSec(v float64) {
	x.NetworkIoBytesPerSec = &v
}

func (x *PerformanceStats) SetDiskIoBytesPerSec(v float64) {
	x.DiskIoBytesPerSec = &v
}

func (x *PerformanceStats) HasAvgResponseTimeMs() bool {
	if x == nil {
		return false
	}
	return x.AvgResponseTimeMs != nil
}

func (x *PerformanceStats) HasMaxResponseTimeMs() bool {
	if x == nil {
		return false
	}
	return x.MaxResponseTimeMs != nil
}

func (x *PerformanceStats) HasMinResponseTimeMs() bool {
	if x == nil {
		return false
	}
	return x.MinResponseTimeMs != nil
}

func (x *PerformanceStats) HasP95ResponseTimeMs() bool {
	if x == nil {
		return false
	}
	return x.P95ResponseTimeMs != nil
}

func (x *PerformanceStats) HasP99ResponseTimeMs() bool {
	if x == nil {
		return false
	}
	return x.P99ResponseTimeMs != nil
}

func (x *PerformanceStats) HasRequestsPerSecond() bool {
	if x == nil {
		return false
	}
	return x.RequestsPerSecond != nil
}

func (x *PerformanceStats) HasTotalRequests() bool {
	if x == nil {
		return false
	}
	return x.TotalRequests != nil
}

func (x *PerformanceStats) HasSuccessfulRequests() bool {
	if x == nil {
		return false
	}
	return x.SuccessfulRequests != nil
}

func (x *PerformanceStats) HasFailedRequests() bool {
	if x == nil {
		return false
	}
	return x.FailedRequests != nil
}

func (x *PerformanceStats) HasSuccessRate() bool {
	if x == nil {
		return false
	}
	return x.SuccessRate != nil
}

func (x *PerformanceStats) HasCpuUtilization() bool {
	if x == nil {
		return false
	}
	return x.CpuUtilization != nil
}

func (x *PerformanceStats) HasMemoryUtilization() bool {
	if x == nil {
		return false
	}
	return x.MemoryUtilization != nil
}

func (x *PerformanceStats) HasNetworkIoBytesPerSec() bool {
	if x == nil {
		return false
	}
	return x.NetworkIoBytesPerSec != nil
}

func (x *PerformanceStats) HasDiskIoBytesPerSec() bool {
	if x == nil {
		return false
	}
	return x.DiskIoBytesPerSec != nil
}

func (x *PerformanceStats) ClearAvgResponseTimeMs() {
	x.AvgResponseTimeMs = nil
}

func (x *PerformanceStats) ClearMaxResponseTimeMs() {
	x.MaxResponseTimeMs = nil
}

func (x *PerformanceStats) ClearMinResponseTimeMs() {
	x.MinResponseTimeMs = nil
}

func (x *PerformanceStats) ClearP95ResponseTimeMs() {
	x.P95ResponseTimeMs = nil
}

func (x *PerformanceStats) ClearP99ResponseTimeMs() {
	x.P99ResponseTimeMs = nil
}

func (x *PerformanceStats) ClearRequestsPerSecond() {
	x.RequestsPerSecond = nil
}

func (x *PerformanceStats) ClearTotalRequests() {
	x.TotalRequests = nil
}

func (x *PerformanceStats) ClearSuccessfulRequests() {
	x.SuccessfulRequests = nil
}

func (x *PerformanceStats) ClearFailedRequests() {
	x.FailedRequests = nil
}

func (x *PerformanceStats) ClearSuccessRate() {
	x.SuccessRate = nil
}

func (x *PerformanceStats) ClearCpuUtilization() {
	x.CpuUtilization = nil
}

func (x *PerformanceStats) ClearMemoryUtilization() {
	x.MemoryUtilization = nil
}

func (x *PerformanceStats) ClearNetworkIoBytesPerSec() {
	x.NetworkIoBytesPerSec = nil
}

func (x *PerformanceStats) ClearDiskIoBytesPerSec() {
	x.DiskIoBytesPerSec = nil
}

type PerformanceStats_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// Average response time in milliseconds
	AvgResponseTimeMs *float64
	// Maximum response time in milliseconds
	MaxResponseTimeMs *float64
	// Minimum response time in milliseconds
	MinResponseTimeMs *float64
	// 95th percentile response time in milliseconds
	P95ResponseTimeMs *float64
	// 99th percentile response time in milliseconds
	P99ResponseTimeMs *float64
	// Requests per second
	RequestsPerSecond *float64
	// Total number of requests processed
	TotalRequests *uint64
	// Number of successful requests
	SuccessfulRequests *uint64
	// Number of failed requests
	FailedRequests *uint64
	// Success rate (0.0 to 1.0)
	SuccessRate *float64
	// CPU utilization percentage (0.0 to 100.0)
	CpuUtilization *float64
	// Memory utilization percentage (0.0 to 100.0)
	MemoryUtilization *float64
	// Network I/O in bytes per second
	NetworkIoBytesPerSec *float64
	// Disk I/O in bytes per second
	DiskIoBytesPerSec *float64
}

func (b0 PerformanceStats_builder) Build() *PerformanceStats {
	m0 := &PerformanceStats{}
	b, x := &b0, m0
	_, _ = b, x
	x.AvgResponseTimeMs = b.AvgResponseTimeMs
	x.MaxResponseTimeMs = b.MaxResponseTimeMs
	x.MinResponseTimeMs = b.MinResponseTimeMs
	x.P95ResponseTimeMs = b.P95ResponseTimeMs
	x.P99ResponseTimeMs = b.P99ResponseTimeMs
	x.RequestsPerSecond = b.RequestsPerSecond
	x.TotalRequests = b.TotalRequests
	x.SuccessfulRequests = b.SuccessfulRequests
	x.FailedRequests = b.FailedRequests
	x.SuccessRate = b.SuccessRate
	x.CpuUtilization = b.CpuUtilization
	x.MemoryUtilization = b.MemoryUtilization
	x.NetworkIoBytesPerSec = b.NetworkIoBytesPerSec
	x.DiskIoBytesPerSec = b.DiskIoBytesPerSec
	return m0
}

var File_pkg_metrics_proto_types_performance_stats_proto protoreflect.FileDescriptor

const file_pkg_metrics_proto_types_performance_stats_proto_rawDesc = "" +
	"\n" +
	"/pkg/metrics/proto/types/performance_stats.proto\x12\x12gcommon.v1.metrics\x1a!google/protobuf/go_features.proto\"\x9d\x05\n" +
	"\x10PerformanceStats\x12/\n" +
	"\x14avg_response_time_ms\x18\x01 \x01(\x01R\x11avgResponseTimeMs\x12/\n" +
	"\x14max_response_time_ms\x18\x02 \x01(\x01R\x11maxResponseTimeMs\x12/\n" +
	"\x14min_response_time_ms\x18\x03 \x01(\x01R\x11minResponseTimeMs\x12/\n" +
	"\x14p95_response_time_ms\x18\x04 \x01(\x01R\x11p95ResponseTimeMs\x12/\n" +
	"\x14p99_response_time_ms\x18\x05 \x01(\x01R\x11p99ResponseTimeMs\x12.\n" +
	"\x13requests_per_second\x18\x06 \x01(\x01R\x11requestsPerSecond\x12%\n" +
	"\x0etotal_requests\x18\a \x01(\x04R\rtotalRequests\x12/\n" +
	"\x13successful_requests\x18\b \x01(\x04R\x12successfulRequests\x12'\n" +
	"\x0ffailed_requests\x18\t \x01(\x04R\x0efailedRequests\x12!\n" +
	"\fsuccess_rate\x18\n" +
	" \x01(\x01R\vsuccessRate\x12'\n" +
	"\x0fcpu_utilization\x18\v \x01(\x01R\x0ecpuUtilization\x12-\n" +
	"\x12memory_utilization\x18\f \x01(\x01R\x11memoryUtilization\x126\n" +
	"\x18network_io_bytes_per_sec\x18\r \x01(\x01R\x14networkIoBytesPerSec\x120\n" +
	"\x15disk_io_bytes_per_sec\x18\x0e \x01(\x01R\x11diskIoBytesPerSecB\xd8\x01\n" +
	"\x16com.gcommon.v1.metricsB\x15PerformanceStatsProtoP\x01Z5github.com/jdfalk/gcommon/pkg/metrics/proto;metricspb\xa2\x02\x03GVM\xaa\x02\x12Gcommon.V1.Metrics\xca\x02\x12Gcommon\\V1\\Metrics\xe2\x02\x1eGcommon\\V1\\Metrics\\GPBMetadata\xea\x02\x14Gcommon::V1::Metrics\x92\x03\x05\xd2>\x02\x10\x02b\beditionsp\xe8\a"

var file_pkg_metrics_proto_types_performance_stats_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pkg_metrics_proto_types_performance_stats_proto_goTypes = []any{
	(*PerformanceStats)(nil), // 0: gcommon.v1.metrics.PerformanceStats
}
var file_pkg_metrics_proto_types_performance_stats_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_metrics_proto_types_performance_stats_proto_init() }
func file_pkg_metrics_proto_types_performance_stats_proto_init() {
	if File_pkg_metrics_proto_types_performance_stats_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pkg_metrics_proto_types_performance_stats_proto_rawDesc), len(file_pkg_metrics_proto_types_performance_stats_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_metrics_proto_types_performance_stats_proto_goTypes,
		DependencyIndexes: file_pkg_metrics_proto_types_performance_stats_proto_depIdxs,
		MessageInfos:      file_pkg_metrics_proto_types_performance_stats_proto_msgTypes,
	}.Build()
	File_pkg_metrics_proto_types_performance_stats_proto = out.File
	file_pkg_metrics_proto_types_performance_stats_proto_goTypes = nil
	file_pkg_metrics_proto_types_performance_stats_proto_depIdxs = nil
}

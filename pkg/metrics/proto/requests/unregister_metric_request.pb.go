// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: pkg/metrics/proto/requests/unregister_metric_request.proto

//go:build !protoopaque

package metricspb

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
// UnregisterMetricRequest represents a request to unregister an existing metric.
// This removes the metric definition and optionally cleans up associated data.
type UnregisterMetricRequest struct {
	state protoimpl.MessageState `protogen:"hybrid.v1"`
	// Standard request metadata (tracing, auth, etc.)
	Metadata *proto.RequestMetadata `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	// Metric identifier (either name or ID)
	//
	// Types that are valid to be assigned to MetricIdentifier:
	//
	//	*UnregisterMetricRequest_MetricName
	//	*UnregisterMetricRequest_MetricId
	MetricIdentifier isUnregisterMetricRequest_MetricIdentifier `protobuf_oneof:"metric_identifier"`
	// Optional provider ID to unregister from
	ProviderId *string `protobuf:"bytes,4,opt,name=provider_id,json=providerId" json:"provider_id,omitempty"`
	// Options for the unregistration process
	Options       *UnregistrationOptions `protobuf:"bytes,5,opt,name=options" json:"options,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UnregisterMetricRequest) Reset() {
	*x = UnregisterMetricRequest{}
	mi := &file_pkg_metrics_proto_requests_unregister_metric_request_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UnregisterMetricRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnregisterMetricRequest) ProtoMessage() {}

func (x *UnregisterMetricRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_metrics_proto_requests_unregister_metric_request_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *UnregisterMetricRequest) GetMetadata() *proto.RequestMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *UnregisterMetricRequest) GetMetricIdentifier() isUnregisterMetricRequest_MetricIdentifier {
	if x != nil {
		return x.MetricIdentifier
	}
	return nil
}

func (x *UnregisterMetricRequest) GetMetricName() string {
	if x != nil {
		if x, ok := x.MetricIdentifier.(*UnregisterMetricRequest_MetricName); ok {
			return x.MetricName
		}
	}
	return ""
}

func (x *UnregisterMetricRequest) GetMetricId() string {
	if x != nil {
		if x, ok := x.MetricIdentifier.(*UnregisterMetricRequest_MetricId); ok {
			return x.MetricId
		}
	}
	return ""
}

func (x *UnregisterMetricRequest) GetProviderId() string {
	if x != nil && x.ProviderId != nil {
		return *x.ProviderId
	}
	return ""
}

func (x *UnregisterMetricRequest) GetOptions() *UnregistrationOptions {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *UnregisterMetricRequest) SetMetadata(v *proto.RequestMetadata) {
	x.Metadata = v
}

func (x *UnregisterMetricRequest) SetMetricName(v string) {
	x.MetricIdentifier = &UnregisterMetricRequest_MetricName{v}
}

func (x *UnregisterMetricRequest) SetMetricId(v string) {
	x.MetricIdentifier = &UnregisterMetricRequest_MetricId{v}
}

func (x *UnregisterMetricRequest) SetProviderId(v string) {
	x.ProviderId = &v
}

func (x *UnregisterMetricRequest) SetOptions(v *UnregistrationOptions) {
	x.Options = v
}

func (x *UnregisterMetricRequest) HasMetadata() bool {
	if x == nil {
		return false
	}
	return x.Metadata != nil
}

func (x *UnregisterMetricRequest) HasMetricIdentifier() bool {
	if x == nil {
		return false
	}
	return x.MetricIdentifier != nil
}

func (x *UnregisterMetricRequest) HasMetricName() bool {
	if x == nil {
		return false
	}
	_, ok := x.MetricIdentifier.(*UnregisterMetricRequest_MetricName)
	return ok
}

func (x *UnregisterMetricRequest) HasMetricId() bool {
	if x == nil {
		return false
	}
	_, ok := x.MetricIdentifier.(*UnregisterMetricRequest_MetricId)
	return ok
}

func (x *UnregisterMetricRequest) HasProviderId() bool {
	if x == nil {
		return false
	}
	return x.ProviderId != nil
}

func (x *UnregisterMetricRequest) HasOptions() bool {
	if x == nil {
		return false
	}
	return x.Options != nil
}

func (x *UnregisterMetricRequest) ClearMetadata() {
	x.Metadata = nil
}

func (x *UnregisterMetricRequest) ClearMetricIdentifier() {
	x.MetricIdentifier = nil
}

func (x *UnregisterMetricRequest) ClearMetricName() {
	if _, ok := x.MetricIdentifier.(*UnregisterMetricRequest_MetricName); ok {
		x.MetricIdentifier = nil
	}
}

func (x *UnregisterMetricRequest) ClearMetricId() {
	if _, ok := x.MetricIdentifier.(*UnregisterMetricRequest_MetricId); ok {
		x.MetricIdentifier = nil
	}
}

func (x *UnregisterMetricRequest) ClearProviderId() {
	x.ProviderId = nil
}

func (x *UnregisterMetricRequest) ClearOptions() {
	x.Options = nil
}

const UnregisterMetricRequest_MetricIdentifier_not_set_case case_UnregisterMetricRequest_MetricIdentifier = 0
const UnregisterMetricRequest_MetricName_case case_UnregisterMetricRequest_MetricIdentifier = 2
const UnregisterMetricRequest_MetricId_case case_UnregisterMetricRequest_MetricIdentifier = 3

func (x *UnregisterMetricRequest) WhichMetricIdentifier() case_UnregisterMetricRequest_MetricIdentifier {
	if x == nil {
		return UnregisterMetricRequest_MetricIdentifier_not_set_case
	}
	switch x.MetricIdentifier.(type) {
	case *UnregisterMetricRequest_MetricName:
		return UnregisterMetricRequest_MetricName_case
	case *UnregisterMetricRequest_MetricId:
		return UnregisterMetricRequest_MetricId_case
	default:
		return UnregisterMetricRequest_MetricIdentifier_not_set_case
	}
}

type UnregisterMetricRequest_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// Standard request metadata (tracing, auth, etc.)
	Metadata *proto.RequestMetadata
	// Metric identifier (either name or ID)

	// Fields of oneof MetricIdentifier:
	// Metric name to unregister
	MetricName *string
	// Metric ID to unregister
	MetricId *string
	// -- end of MetricIdentifier
	// Optional provider ID to unregister from
	ProviderId *string
	// Options for the unregistration process
	Options *UnregistrationOptions
}

func (b0 UnregisterMetricRequest_builder) Build() *UnregisterMetricRequest {
	m0 := &UnregisterMetricRequest{}
	b, x := &b0, m0
	_, _ = b, x
	x.Metadata = b.Metadata
	if b.MetricName != nil {
		x.MetricIdentifier = &UnregisterMetricRequest_MetricName{*b.MetricName}
	}
	if b.MetricId != nil {
		x.MetricIdentifier = &UnregisterMetricRequest_MetricId{*b.MetricId}
	}
	x.ProviderId = b.ProviderId
	x.Options = b.Options
	return m0
}

type case_UnregisterMetricRequest_MetricIdentifier protoreflect.FieldNumber

func (x case_UnregisterMetricRequest_MetricIdentifier) String() string {
	md := file_pkg_metrics_proto_requests_unregister_metric_request_proto_msgTypes[0].Descriptor()
	if x == 0 {
		return "not set"
	}
	return protoimpl.X.MessageFieldStringOf(md, protoreflect.FieldNumber(x))
}

type isUnregisterMetricRequest_MetricIdentifier interface {
	isUnregisterMetricRequest_MetricIdentifier()
}

type UnregisterMetricRequest_MetricName struct {
	// Metric name to unregister
	MetricName string `protobuf:"bytes,2,opt,name=metric_name,json=metricName,oneof"`
}

type UnregisterMetricRequest_MetricId struct {
	// Metric ID to unregister
	MetricId string `protobuf:"bytes,3,opt,name=metric_id,json=metricId,oneof"`
}

func (*UnregisterMetricRequest_MetricName) isUnregisterMetricRequest_MetricIdentifier() {}

func (*UnregisterMetricRequest_MetricId) isUnregisterMetricRequest_MetricIdentifier() {}

// *
// UnregistrationOptions configure the unregistration process.
type UnregistrationOptions struct {
	state protoimpl.MessageState `protogen:"hybrid.v1"`
	// Whether to delete all associated data
	DeleteData *bool `protobuf:"varint,1,opt,name=delete_data,json=deleteData" json:"delete_data,omitempty"`
	// Whether to delete associated indices
	DeleteIndices *bool `protobuf:"varint,2,opt,name=delete_indices,json=deleteIndices" json:"delete_indices,omitempty"`
	// Whether to remove alert rules
	RemoveAlerts *bool `protobuf:"varint,3,opt,name=remove_alerts,json=removeAlerts" json:"remove_alerts,omitempty"`
	// Whether to stop export configurations
	StopExports *bool `protobuf:"varint,4,opt,name=stop_exports,json=stopExports" json:"stop_exports,omitempty"`
	// Grace period before actual deletion (duration string like "24h")
	GracePeriod *string `protobuf:"bytes,5,opt,name=grace_period,json=gracePeriod" json:"grace_period,omitempty"`
	// Whether to perform a dry run (show what would be deleted)
	DryRun *bool `protobuf:"varint,6,opt,name=dry_run,json=dryRun" json:"dry_run,omitempty"`
	// Whether to force deletion even if metric is in use
	Force *bool `protobuf:"varint,7,opt,name=force" json:"force,omitempty"`
	// Whether to create a backup before deletion
	CreateBackup  *bool `protobuf:"varint,8,opt,name=create_backup,json=createBackup" json:"create_backup,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UnregistrationOptions) Reset() {
	*x = UnregistrationOptions{}
	mi := &file_pkg_metrics_proto_requests_unregister_metric_request_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UnregistrationOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnregistrationOptions) ProtoMessage() {}

func (x *UnregistrationOptions) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_metrics_proto_requests_unregister_metric_request_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *UnregistrationOptions) GetDeleteData() bool {
	if x != nil && x.DeleteData != nil {
		return *x.DeleteData
	}
	return false
}

func (x *UnregistrationOptions) GetDeleteIndices() bool {
	if x != nil && x.DeleteIndices != nil {
		return *x.DeleteIndices
	}
	return false
}

func (x *UnregistrationOptions) GetRemoveAlerts() bool {
	if x != nil && x.RemoveAlerts != nil {
		return *x.RemoveAlerts
	}
	return false
}

func (x *UnregistrationOptions) GetStopExports() bool {
	if x != nil && x.StopExports != nil {
		return *x.StopExports
	}
	return false
}

func (x *UnregistrationOptions) GetGracePeriod() string {
	if x != nil && x.GracePeriod != nil {
		return *x.GracePeriod
	}
	return ""
}

func (x *UnregistrationOptions) GetDryRun() bool {
	if x != nil && x.DryRun != nil {
		return *x.DryRun
	}
	return false
}

func (x *UnregistrationOptions) GetForce() bool {
	if x != nil && x.Force != nil {
		return *x.Force
	}
	return false
}

func (x *UnregistrationOptions) GetCreateBackup() bool {
	if x != nil && x.CreateBackup != nil {
		return *x.CreateBackup
	}
	return false
}

func (x *UnregistrationOptions) SetDeleteData(v bool) {
	x.DeleteData = &v
}

func (x *UnregistrationOptions) SetDeleteIndices(v bool) {
	x.DeleteIndices = &v
}

func (x *UnregistrationOptions) SetRemoveAlerts(v bool) {
	x.RemoveAlerts = &v
}

func (x *UnregistrationOptions) SetStopExports(v bool) {
	x.StopExports = &v
}

func (x *UnregistrationOptions) SetGracePeriod(v string) {
	x.GracePeriod = &v
}

func (x *UnregistrationOptions) SetDryRun(v bool) {
	x.DryRun = &v
}

func (x *UnregistrationOptions) SetForce(v bool) {
	x.Force = &v
}

func (x *UnregistrationOptions) SetCreateBackup(v bool) {
	x.CreateBackup = &v
}

func (x *UnregistrationOptions) HasDeleteData() bool {
	if x == nil {
		return false
	}
	return x.DeleteData != nil
}

func (x *UnregistrationOptions) HasDeleteIndices() bool {
	if x == nil {
		return false
	}
	return x.DeleteIndices != nil
}

func (x *UnregistrationOptions) HasRemoveAlerts() bool {
	if x == nil {
		return false
	}
	return x.RemoveAlerts != nil
}

func (x *UnregistrationOptions) HasStopExports() bool {
	if x == nil {
		return false
	}
	return x.StopExports != nil
}

func (x *UnregistrationOptions) HasGracePeriod() bool {
	if x == nil {
		return false
	}
	return x.GracePeriod != nil
}

func (x *UnregistrationOptions) HasDryRun() bool {
	if x == nil {
		return false
	}
	return x.DryRun != nil
}

func (x *UnregistrationOptions) HasForce() bool {
	if x == nil {
		return false
	}
	return x.Force != nil
}

func (x *UnregistrationOptions) HasCreateBackup() bool {
	if x == nil {
		return false
	}
	return x.CreateBackup != nil
}

func (x *UnregistrationOptions) ClearDeleteData() {
	x.DeleteData = nil
}

func (x *UnregistrationOptions) ClearDeleteIndices() {
	x.DeleteIndices = nil
}

func (x *UnregistrationOptions) ClearRemoveAlerts() {
	x.RemoveAlerts = nil
}

func (x *UnregistrationOptions) ClearStopExports() {
	x.StopExports = nil
}

func (x *UnregistrationOptions) ClearGracePeriod() {
	x.GracePeriod = nil
}

func (x *UnregistrationOptions) ClearDryRun() {
	x.DryRun = nil
}

func (x *UnregistrationOptions) ClearForce() {
	x.Force = nil
}

func (x *UnregistrationOptions) ClearCreateBackup() {
	x.CreateBackup = nil
}

type UnregistrationOptions_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// Whether to delete all associated data
	DeleteData *bool
	// Whether to delete associated indices
	DeleteIndices *bool
	// Whether to remove alert rules
	RemoveAlerts *bool
	// Whether to stop export configurations
	StopExports *bool
	// Grace period before actual deletion (duration string like "24h")
	GracePeriod *string
	// Whether to perform a dry run (show what would be deleted)
	DryRun *bool
	// Whether to force deletion even if metric is in use
	Force *bool
	// Whether to create a backup before deletion
	CreateBackup *bool
}

func (b0 UnregistrationOptions_builder) Build() *UnregistrationOptions {
	m0 := &UnregistrationOptions{}
	b, x := &b0, m0
	_, _ = b, x
	x.DeleteData = b.DeleteData
	x.DeleteIndices = b.DeleteIndices
	x.RemoveAlerts = b.RemoveAlerts
	x.StopExports = b.StopExports
	x.GracePeriod = b.GracePeriod
	x.DryRun = b.DryRun
	x.Force = b.Force
	x.CreateBackup = b.CreateBackup
	return m0
}

var File_pkg_metrics_proto_requests_unregister_metric_request_proto protoreflect.FileDescriptor

const file_pkg_metrics_proto_requests_unregister_metric_request_proto_rawDesc = "" +
	"\n" +
	":pkg/metrics/proto/requests/unregister_metric_request.proto\x12\x12gcommon.v1.metrics\x1a!google/protobuf/go_features.proto\x1a0pkg/common/proto/messages/request_metadata.proto\"\x96\x02\n" +
	"\x17UnregisterMetricRequest\x12>\n" +
	"\bmetadata\x18\x01 \x01(\v2\".gcommon.v1.common.RequestMetadataR\bmetadata\x12!\n" +
	"\vmetric_name\x18\x02 \x01(\tH\x00R\n" +
	"metricName\x12\x1d\n" +
	"\tmetric_id\x18\x03 \x01(\tH\x00R\bmetricId\x12\x1f\n" +
	"\vprovider_id\x18\x04 \x01(\tR\n" +
	"providerId\x12C\n" +
	"\aoptions\x18\x05 \x01(\v2).gcommon.v1.metrics.UnregistrationOptionsR\aoptionsB\x13\n" +
	"\x11metric_identifier\"\x9e\x02\n" +
	"\x15UnregistrationOptions\x12\x1f\n" +
	"\vdelete_data\x18\x01 \x01(\bR\n" +
	"deleteData\x12%\n" +
	"\x0edelete_indices\x18\x02 \x01(\bR\rdeleteIndices\x12#\n" +
	"\rremove_alerts\x18\x03 \x01(\bR\fremoveAlerts\x12!\n" +
	"\fstop_exports\x18\x04 \x01(\bR\vstopExports\x12!\n" +
	"\fgrace_period\x18\x05 \x01(\tR\vgracePeriod\x12\x17\n" +
	"\adry_run\x18\x06 \x01(\bR\x06dryRun\x12\x14\n" +
	"\x05force\x18\a \x01(\bR\x05force\x12#\n" +
	"\rcreate_backup\x18\b \x01(\bR\fcreateBackupB\xdf\x01\n" +
	"\x16com.gcommon.v1.metricsB\x1cUnregisterMetricRequestProtoP\x01Z5github.com/jdfalk/gcommon/pkg/metrics/proto;metricspb\xa2\x02\x03GVM\xaa\x02\x12Gcommon.V1.Metrics\xca\x02\x12Gcommon\\V1\\Metrics\xe2\x02\x1eGcommon\\V1\\Metrics\\GPBMetadata\xea\x02\x14Gcommon::V1::Metrics\x92\x03\x05\xd2>\x02\x10\x02b\beditionsp\xe8\a"

var file_pkg_metrics_proto_requests_unregister_metric_request_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_metrics_proto_requests_unregister_metric_request_proto_goTypes = []any{
	(*UnregisterMetricRequest)(nil), // 0: gcommon.v1.metrics.UnregisterMetricRequest
	(*UnregistrationOptions)(nil),   // 1: gcommon.v1.metrics.UnregistrationOptions
	(*proto.RequestMetadata)(nil),   // 2: gcommon.v1.common.RequestMetadata
}
var file_pkg_metrics_proto_requests_unregister_metric_request_proto_depIdxs = []int32{
	2, // 0: gcommon.v1.metrics.UnregisterMetricRequest.metadata:type_name -> gcommon.v1.common.RequestMetadata
	1, // 1: gcommon.v1.metrics.UnregisterMetricRequest.options:type_name -> gcommon.v1.metrics.UnregistrationOptions
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_metrics_proto_requests_unregister_metric_request_proto_init() }
func file_pkg_metrics_proto_requests_unregister_metric_request_proto_init() {
	if File_pkg_metrics_proto_requests_unregister_metric_request_proto != nil {
		return
	}
	file_pkg_metrics_proto_requests_unregister_metric_request_proto_msgTypes[0].OneofWrappers = []any{
		(*UnregisterMetricRequest_MetricName)(nil),
		(*UnregisterMetricRequest_MetricId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pkg_metrics_proto_requests_unregister_metric_request_proto_rawDesc), len(file_pkg_metrics_proto_requests_unregister_metric_request_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_metrics_proto_requests_unregister_metric_request_proto_goTypes,
		DependencyIndexes: file_pkg_metrics_proto_requests_unregister_metric_request_proto_depIdxs,
		MessageInfos:      file_pkg_metrics_proto_requests_unregister_metric_request_proto_msgTypes,
	}.Build()
	File_pkg_metrics_proto_requests_unregister_metric_request_proto = out.File
	file_pkg_metrics_proto_requests_unregister_metric_request_proto_goTypes = nil
	file_pkg_metrics_proto_requests_unregister_metric_request_proto_depIdxs = nil
}

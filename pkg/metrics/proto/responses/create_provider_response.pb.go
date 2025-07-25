// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: pkg/metrics/proto/responses/create_provider_response.proto

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
// CreateProviderResponse contains the result of creating a metrics provider.
type CreateProviderResponse struct {
	state protoimpl.MessageState `protogen:"hybrid.v1"`
	// Success status of the creation
	Success *bool `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	// Error information if creation failed
	Error *proto.Error `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
	// ID of the created provider
	ProviderId *string `protobuf:"bytes,3,opt,name=provider_id,json=providerId" json:"provider_id,omitempty"`
	// When the provider was created
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	// Current status of the provider
	Status *ProviderStatus `protobuf:"bytes,5,opt,name=status" json:"status,omitempty"`
	// Validation results
	Validation *ValidationResult `protobuf:"bytes,6,opt,name=validation" json:"validation,omitempty"`
	// Configuration details that were applied
	AppliedConfig *AppliedConfig `protobuf:"bytes,7,opt,name=applied_config,json=appliedConfig" json:"applied_config,omitempty"`
	// Warnings or informational messages
	Warnings []string `protobuf:"bytes,8,rep,name=warnings" json:"warnings,omitempty"`
	// Endpoint information for accessing the provider
	Endpoints     *ProviderEndpoints `protobuf:"bytes,9,opt,name=endpoints" json:"endpoints,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateProviderResponse) Reset() {
	*x = CreateProviderResponse{}
	mi := &file_pkg_metrics_proto_responses_create_provider_response_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateProviderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProviderResponse) ProtoMessage() {}

func (x *CreateProviderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_metrics_proto_responses_create_provider_response_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *CreateProviderResponse) GetSuccess() bool {
	if x != nil && x.Success != nil {
		return *x.Success
	}
	return false
}

func (x *CreateProviderResponse) GetError() *proto.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *CreateProviderResponse) GetProviderId() string {
	if x != nil && x.ProviderId != nil {
		return *x.ProviderId
	}
	return ""
}

func (x *CreateProviderResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *CreateProviderResponse) GetStatus() *ProviderStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *CreateProviderResponse) GetValidation() *ValidationResult {
	if x != nil {
		return x.Validation
	}
	return nil
}

func (x *CreateProviderResponse) GetAppliedConfig() *AppliedConfig {
	if x != nil {
		return x.AppliedConfig
	}
	return nil
}

func (x *CreateProviderResponse) GetWarnings() []string {
	if x != nil {
		return x.Warnings
	}
	return nil
}

func (x *CreateProviderResponse) GetEndpoints() *ProviderEndpoints {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

func (x *CreateProviderResponse) SetSuccess(v bool) {
	x.Success = &v
}

func (x *CreateProviderResponse) SetError(v *proto.Error) {
	x.Error = v
}

func (x *CreateProviderResponse) SetProviderId(v string) {
	x.ProviderId = &v
}

func (x *CreateProviderResponse) SetCreatedAt(v *timestamppb.Timestamp) {
	x.CreatedAt = v
}

func (x *CreateProviderResponse) SetStatus(v *ProviderStatus) {
	x.Status = v
}

func (x *CreateProviderResponse) SetValidation(v *ValidationResult) {
	x.Validation = v
}

func (x *CreateProviderResponse) SetAppliedConfig(v *AppliedConfig) {
	x.AppliedConfig = v
}

func (x *CreateProviderResponse) SetWarnings(v []string) {
	x.Warnings = v
}

func (x *CreateProviderResponse) SetEndpoints(v *ProviderEndpoints) {
	x.Endpoints = v
}

func (x *CreateProviderResponse) HasSuccess() bool {
	if x == nil {
		return false
	}
	return x.Success != nil
}

func (x *CreateProviderResponse) HasError() bool {
	if x == nil {
		return false
	}
	return x.Error != nil
}

func (x *CreateProviderResponse) HasProviderId() bool {
	if x == nil {
		return false
	}
	return x.ProviderId != nil
}

func (x *CreateProviderResponse) HasCreatedAt() bool {
	if x == nil {
		return false
	}
	return x.CreatedAt != nil
}

func (x *CreateProviderResponse) HasStatus() bool {
	if x == nil {
		return false
	}
	return x.Status != nil
}

func (x *CreateProviderResponse) HasValidation() bool {
	if x == nil {
		return false
	}
	return x.Validation != nil
}

func (x *CreateProviderResponse) HasAppliedConfig() bool {
	if x == nil {
		return false
	}
	return x.AppliedConfig != nil
}

func (x *CreateProviderResponse) HasEndpoints() bool {
	if x == nil {
		return false
	}
	return x.Endpoints != nil
}

func (x *CreateProviderResponse) ClearSuccess() {
	x.Success = nil
}

func (x *CreateProviderResponse) ClearError() {
	x.Error = nil
}

func (x *CreateProviderResponse) ClearProviderId() {
	x.ProviderId = nil
}

func (x *CreateProviderResponse) ClearCreatedAt() {
	x.CreatedAt = nil
}

func (x *CreateProviderResponse) ClearStatus() {
	x.Status = nil
}

func (x *CreateProviderResponse) ClearValidation() {
	x.Validation = nil
}

func (x *CreateProviderResponse) ClearAppliedConfig() {
	x.AppliedConfig = nil
}

func (x *CreateProviderResponse) ClearEndpoints() {
	x.Endpoints = nil
}

type CreateProviderResponse_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// Success status of the creation
	Success *bool
	// Error information if creation failed
	Error *proto.Error
	// ID of the created provider
	ProviderId *string
	// When the provider was created
	CreatedAt *timestamppb.Timestamp
	// Current status of the provider
	Status *ProviderStatus
	// Validation results
	Validation *ValidationResult
	// Configuration details that were applied
	AppliedConfig *AppliedConfig
	// Warnings or informational messages
	Warnings []string
	// Endpoint information for accessing the provider
	Endpoints *ProviderEndpoints
}

func (b0 CreateProviderResponse_builder) Build() *CreateProviderResponse {
	m0 := &CreateProviderResponse{}
	b, x := &b0, m0
	_, _ = b, x
	x.Success = b.Success
	x.Error = b.Error
	x.ProviderId = b.ProviderId
	x.CreatedAt = b.CreatedAt
	x.Status = b.Status
	x.Validation = b.Validation
	x.AppliedConfig = b.AppliedConfig
	x.Warnings = b.Warnings
	x.Endpoints = b.Endpoints
	return m0
}

// *
// AppliedConfig contains information about the configuration that was applied.
type AppliedConfig struct {
	state protoimpl.MessageState `protogen:"hybrid.v1"`
	// Configuration that was actually applied (may differ from requested)
	ConfigSummary *string `protobuf:"bytes,1,opt,name=config_summary,json=configSummary" json:"config_summary,omitempty"`
	// Default values that were applied
	AppliedDefaults map[string]string `protobuf:"bytes,2,rep,name=applied_defaults,json=appliedDefaults" json:"applied_defaults,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Configuration overrides that were applied
	AppliedOverrides map[string]string `protobuf:"bytes,3,rep,name=applied_overrides,json=appliedOverrides" json:"applied_overrides,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Resource allocations that were made
	ResourceAllocations *ResourceAllocations `protobuf:"bytes,4,opt,name=resource_allocations,json=resourceAllocations" json:"resource_allocations,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *AppliedConfig) Reset() {
	*x = AppliedConfig{}
	mi := &file_pkg_metrics_proto_responses_create_provider_response_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AppliedConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppliedConfig) ProtoMessage() {}

func (x *AppliedConfig) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_metrics_proto_responses_create_provider_response_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *AppliedConfig) GetConfigSummary() string {
	if x != nil && x.ConfigSummary != nil {
		return *x.ConfigSummary
	}
	return ""
}

func (x *AppliedConfig) GetAppliedDefaults() map[string]string {
	if x != nil {
		return x.AppliedDefaults
	}
	return nil
}

func (x *AppliedConfig) GetAppliedOverrides() map[string]string {
	if x != nil {
		return x.AppliedOverrides
	}
	return nil
}

func (x *AppliedConfig) GetResourceAllocations() *ResourceAllocations {
	if x != nil {
		return x.ResourceAllocations
	}
	return nil
}

func (x *AppliedConfig) SetConfigSummary(v string) {
	x.ConfigSummary = &v
}

func (x *AppliedConfig) SetAppliedDefaults(v map[string]string) {
	x.AppliedDefaults = v
}

func (x *AppliedConfig) SetAppliedOverrides(v map[string]string) {
	x.AppliedOverrides = v
}

func (x *AppliedConfig) SetResourceAllocations(v *ResourceAllocations) {
	x.ResourceAllocations = v
}

func (x *AppliedConfig) HasConfigSummary() bool {
	if x == nil {
		return false
	}
	return x.ConfigSummary != nil
}

func (x *AppliedConfig) HasResourceAllocations() bool {
	if x == nil {
		return false
	}
	return x.ResourceAllocations != nil
}

func (x *AppliedConfig) ClearConfigSummary() {
	x.ConfigSummary = nil
}

func (x *AppliedConfig) ClearResourceAllocations() {
	x.ResourceAllocations = nil
}

type AppliedConfig_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// Configuration that was actually applied (may differ from requested)
	ConfigSummary *string
	// Default values that were applied
	AppliedDefaults map[string]string
	// Configuration overrides that were applied
	AppliedOverrides map[string]string
	// Resource allocations that were made
	ResourceAllocations *ResourceAllocations
}

func (b0 AppliedConfig_builder) Build() *AppliedConfig {
	m0 := &AppliedConfig{}
	b, x := &b0, m0
	_, _ = b, x
	x.ConfigSummary = b.ConfigSummary
	x.AppliedDefaults = b.AppliedDefaults
	x.AppliedOverrides = b.AppliedOverrides
	x.ResourceAllocations = b.ResourceAllocations
	return m0
}

// *
// ResourceAllocations contains information about allocated resources.
type ResourceAllocations struct {
	state protoimpl.MessageState `protogen:"hybrid.v1"`
	// Allocated memory (bytes)
	AllocatedMemoryBytes *int64 `protobuf:"varint,1,opt,name=allocated_memory_bytes,json=allocatedMemoryBytes" json:"allocated_memory_bytes,omitempty"`
	// Allocated CPU (percentage)
	AllocatedCpuPercent *float64 `protobuf:"fixed64,2,opt,name=allocated_cpu_percent,json=allocatedCpuPercent" json:"allocated_cpu_percent,omitempty"`
	// Allocated disk space (bytes)
	AllocatedDiskBytes *int64 `protobuf:"varint,3,opt,name=allocated_disk_bytes,json=allocatedDiskBytes" json:"allocated_disk_bytes,omitempty"`
	// Network ports allocated
	AllocatedPorts []int32 `protobuf:"varint,4,rep,packed,name=allocated_ports,json=allocatedPorts" json:"allocated_ports,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *ResourceAllocations) Reset() {
	*x = ResourceAllocations{}
	mi := &file_pkg_metrics_proto_responses_create_provider_response_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResourceAllocations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceAllocations) ProtoMessage() {}

func (x *ResourceAllocations) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_metrics_proto_responses_create_provider_response_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *ResourceAllocations) GetAllocatedMemoryBytes() int64 {
	if x != nil && x.AllocatedMemoryBytes != nil {
		return *x.AllocatedMemoryBytes
	}
	return 0
}

func (x *ResourceAllocations) GetAllocatedCpuPercent() float64 {
	if x != nil && x.AllocatedCpuPercent != nil {
		return *x.AllocatedCpuPercent
	}
	return 0
}

func (x *ResourceAllocations) GetAllocatedDiskBytes() int64 {
	if x != nil && x.AllocatedDiskBytes != nil {
		return *x.AllocatedDiskBytes
	}
	return 0
}

func (x *ResourceAllocations) GetAllocatedPorts() []int32 {
	if x != nil {
		return x.AllocatedPorts
	}
	return nil
}

func (x *ResourceAllocations) SetAllocatedMemoryBytes(v int64) {
	x.AllocatedMemoryBytes = &v
}

func (x *ResourceAllocations) SetAllocatedCpuPercent(v float64) {
	x.AllocatedCpuPercent = &v
}

func (x *ResourceAllocations) SetAllocatedDiskBytes(v int64) {
	x.AllocatedDiskBytes = &v
}

func (x *ResourceAllocations) SetAllocatedPorts(v []int32) {
	x.AllocatedPorts = v
}

func (x *ResourceAllocations) HasAllocatedMemoryBytes() bool {
	if x == nil {
		return false
	}
	return x.AllocatedMemoryBytes != nil
}

func (x *ResourceAllocations) HasAllocatedCpuPercent() bool {
	if x == nil {
		return false
	}
	return x.AllocatedCpuPercent != nil
}

func (x *ResourceAllocations) HasAllocatedDiskBytes() bool {
	if x == nil {
		return false
	}
	return x.AllocatedDiskBytes != nil
}

func (x *ResourceAllocations) ClearAllocatedMemoryBytes() {
	x.AllocatedMemoryBytes = nil
}

func (x *ResourceAllocations) ClearAllocatedCpuPercent() {
	x.AllocatedCpuPercent = nil
}

func (x *ResourceAllocations) ClearAllocatedDiskBytes() {
	x.AllocatedDiskBytes = nil
}

type ResourceAllocations_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// Allocated memory (bytes)
	AllocatedMemoryBytes *int64
	// Allocated CPU (percentage)
	AllocatedCpuPercent *float64
	// Allocated disk space (bytes)
	AllocatedDiskBytes *int64
	// Network ports allocated
	AllocatedPorts []int32
}

func (b0 ResourceAllocations_builder) Build() *ResourceAllocations {
	m0 := &ResourceAllocations{}
	b, x := &b0, m0
	_, _ = b, x
	x.AllocatedMemoryBytes = b.AllocatedMemoryBytes
	x.AllocatedCpuPercent = b.AllocatedCpuPercent
	x.AllocatedDiskBytes = b.AllocatedDiskBytes
	x.AllocatedPorts = b.AllocatedPorts
	return m0
}

// *
// ProviderEndpoints contains endpoint information for the provider.
type ProviderEndpoints struct {
	state protoimpl.MessageState `protogen:"hybrid.v1"`
	// Main service endpoint
	ServiceEndpoint *string `protobuf:"bytes,1,opt,name=service_endpoint,json=serviceEndpoint" json:"service_endpoint,omitempty"`
	// Metrics endpoint (for scraping)
	MetricsEndpoint *string `protobuf:"bytes,2,opt,name=metrics_endpoint,json=metricsEndpoint" json:"metrics_endpoint,omitempty"`
	// Health check endpoint
	HealthEndpoint *string `protobuf:"bytes,3,opt,name=health_endpoint,json=healthEndpoint" json:"health_endpoint,omitempty"`
	// Admin/management endpoint
	AdminEndpoint *string `protobuf:"bytes,4,opt,name=admin_endpoint,json=adminEndpoint" json:"admin_endpoint,omitempty"`
	// Additional endpoints
	AdditionalEndpoints map[string]string `protobuf:"bytes,5,rep,name=additional_endpoints,json=additionalEndpoints" json:"additional_endpoints,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *ProviderEndpoints) Reset() {
	*x = ProviderEndpoints{}
	mi := &file_pkg_metrics_proto_responses_create_provider_response_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProviderEndpoints) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProviderEndpoints) ProtoMessage() {}

func (x *ProviderEndpoints) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_metrics_proto_responses_create_provider_response_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *ProviderEndpoints) GetServiceEndpoint() string {
	if x != nil && x.ServiceEndpoint != nil {
		return *x.ServiceEndpoint
	}
	return ""
}

func (x *ProviderEndpoints) GetMetricsEndpoint() string {
	if x != nil && x.MetricsEndpoint != nil {
		return *x.MetricsEndpoint
	}
	return ""
}

func (x *ProviderEndpoints) GetHealthEndpoint() string {
	if x != nil && x.HealthEndpoint != nil {
		return *x.HealthEndpoint
	}
	return ""
}

func (x *ProviderEndpoints) GetAdminEndpoint() string {
	if x != nil && x.AdminEndpoint != nil {
		return *x.AdminEndpoint
	}
	return ""
}

func (x *ProviderEndpoints) GetAdditionalEndpoints() map[string]string {
	if x != nil {
		return x.AdditionalEndpoints
	}
	return nil
}

func (x *ProviderEndpoints) SetServiceEndpoint(v string) {
	x.ServiceEndpoint = &v
}

func (x *ProviderEndpoints) SetMetricsEndpoint(v string) {
	x.MetricsEndpoint = &v
}

func (x *ProviderEndpoints) SetHealthEndpoint(v string) {
	x.HealthEndpoint = &v
}

func (x *ProviderEndpoints) SetAdminEndpoint(v string) {
	x.AdminEndpoint = &v
}

func (x *ProviderEndpoints) SetAdditionalEndpoints(v map[string]string) {
	x.AdditionalEndpoints = v
}

func (x *ProviderEndpoints) HasServiceEndpoint() bool {
	if x == nil {
		return false
	}
	return x.ServiceEndpoint != nil
}

func (x *ProviderEndpoints) HasMetricsEndpoint() bool {
	if x == nil {
		return false
	}
	return x.MetricsEndpoint != nil
}

func (x *ProviderEndpoints) HasHealthEndpoint() bool {
	if x == nil {
		return false
	}
	return x.HealthEndpoint != nil
}

func (x *ProviderEndpoints) HasAdminEndpoint() bool {
	if x == nil {
		return false
	}
	return x.AdminEndpoint != nil
}

func (x *ProviderEndpoints) ClearServiceEndpoint() {
	x.ServiceEndpoint = nil
}

func (x *ProviderEndpoints) ClearMetricsEndpoint() {
	x.MetricsEndpoint = nil
}

func (x *ProviderEndpoints) ClearHealthEndpoint() {
	x.HealthEndpoint = nil
}

func (x *ProviderEndpoints) ClearAdminEndpoint() {
	x.AdminEndpoint = nil
}

type ProviderEndpoints_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// Main service endpoint
	ServiceEndpoint *string
	// Metrics endpoint (for scraping)
	MetricsEndpoint *string
	// Health check endpoint
	HealthEndpoint *string
	// Admin/management endpoint
	AdminEndpoint *string
	// Additional endpoints
	AdditionalEndpoints map[string]string
}

func (b0 ProviderEndpoints_builder) Build() *ProviderEndpoints {
	m0 := &ProviderEndpoints{}
	b, x := &b0, m0
	_, _ = b, x
	x.ServiceEndpoint = b.ServiceEndpoint
	x.MetricsEndpoint = b.MetricsEndpoint
	x.HealthEndpoint = b.HealthEndpoint
	x.AdminEndpoint = b.AdminEndpoint
	x.AdditionalEndpoints = b.AdditionalEndpoints
	return m0
}

var File_pkg_metrics_proto_responses_create_provider_response_proto protoreflect.FileDescriptor

const file_pkg_metrics_proto_responses_create_provider_response_proto_rawDesc = "" +
	"\n" +
	":pkg/metrics/proto/responses/create_provider_response.proto\x12\x12gcommon.v1.metrics\x1a!google/protobuf/go_features.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a%pkg/common/proto/messages/error.proto\x1a-pkg/metrics/proto/types/provider_status.proto\x1a,pkg/metrics/proto/enums/provider_state.proto\x1a/pkg/metrics/proto/types/validation_result.proto\"\xeb\x03\n" +
	"\x16CreateProviderResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\x12.\n" +
	"\x05error\x18\x02 \x01(\v2\x18.gcommon.v1.common.ErrorR\x05error\x12\x1f\n" +
	"\vprovider_id\x18\x03 \x01(\tR\n" +
	"providerId\x129\n" +
	"\n" +
	"created_at\x18\x04 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x12:\n" +
	"\x06status\x18\x05 \x01(\v2\".gcommon.v1.metrics.ProviderStatusR\x06status\x12D\n" +
	"\n" +
	"validation\x18\x06 \x01(\v2$.gcommon.v1.metrics.ValidationResultR\n" +
	"validation\x12H\n" +
	"\x0eapplied_config\x18\a \x01(\v2!.gcommon.v1.metrics.AppliedConfigR\rappliedConfig\x12\x1a\n" +
	"\bwarnings\x18\b \x03(\tR\bwarnings\x12C\n" +
	"\tendpoints\x18\t \x01(\v2%.gcommon.v1.metrics.ProviderEndpointsR\tendpoints\"\xe4\x03\n" +
	"\rAppliedConfig\x12%\n" +
	"\x0econfig_summary\x18\x01 \x01(\tR\rconfigSummary\x12a\n" +
	"\x10applied_defaults\x18\x02 \x03(\v26.gcommon.v1.metrics.AppliedConfig.AppliedDefaultsEntryR\x0fappliedDefaults\x12d\n" +
	"\x11applied_overrides\x18\x03 \x03(\v27.gcommon.v1.metrics.AppliedConfig.AppliedOverridesEntryR\x10appliedOverrides\x12Z\n" +
	"\x14resource_allocations\x18\x04 \x01(\v2'.gcommon.v1.metrics.ResourceAllocationsR\x13resourceAllocations\x1aB\n" +
	"\x14AppliedDefaultsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\x1aC\n" +
	"\x15AppliedOverridesEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\"\xda\x01\n" +
	"\x13ResourceAllocations\x124\n" +
	"\x16allocated_memory_bytes\x18\x01 \x01(\x03R\x14allocatedMemoryBytes\x122\n" +
	"\x15allocated_cpu_percent\x18\x02 \x01(\x01R\x13allocatedCpuPercent\x120\n" +
	"\x14allocated_disk_bytes\x18\x03 \x01(\x03R\x12allocatedDiskBytes\x12'\n" +
	"\x0fallocated_ports\x18\x04 \x03(\x05R\x0eallocatedPorts\"\xf4\x02\n" +
	"\x11ProviderEndpoints\x12)\n" +
	"\x10service_endpoint\x18\x01 \x01(\tR\x0fserviceEndpoint\x12)\n" +
	"\x10metrics_endpoint\x18\x02 \x01(\tR\x0fmetricsEndpoint\x12'\n" +
	"\x0fhealth_endpoint\x18\x03 \x01(\tR\x0ehealthEndpoint\x12%\n" +
	"\x0eadmin_endpoint\x18\x04 \x01(\tR\radminEndpoint\x12q\n" +
	"\x14additional_endpoints\x18\x05 \x03(\v2>.gcommon.v1.metrics.ProviderEndpoints.AdditionalEndpointsEntryR\x13additionalEndpoints\x1aF\n" +
	"\x18AdditionalEndpointsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01B\xde\x01\n" +
	"\x16com.gcommon.v1.metricsB\x1bCreateProviderResponseProtoP\x01Z5github.com/jdfalk/gcommon/pkg/metrics/proto;metricspb\xa2\x02\x03GVM\xaa\x02\x12Gcommon.V1.Metrics\xca\x02\x12Gcommon\\V1\\Metrics\xe2\x02\x1eGcommon\\V1\\Metrics\\GPBMetadata\xea\x02\x14Gcommon::V1::Metrics\x92\x03\x05\xd2>\x02\x10\x02b\beditionsp\xe8\a"

var file_pkg_metrics_proto_responses_create_provider_response_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pkg_metrics_proto_responses_create_provider_response_proto_goTypes = []any{
	(*CreateProviderResponse)(nil), // 0: gcommon.v1.metrics.CreateProviderResponse
	(*AppliedConfig)(nil),          // 1: gcommon.v1.metrics.AppliedConfig
	(*ResourceAllocations)(nil),    // 2: gcommon.v1.metrics.ResourceAllocations
	(*ProviderEndpoints)(nil),      // 3: gcommon.v1.metrics.ProviderEndpoints
	nil,                            // 4: gcommon.v1.metrics.AppliedConfig.AppliedDefaultsEntry
	nil,                            // 5: gcommon.v1.metrics.AppliedConfig.AppliedOverridesEntry
	nil,                            // 6: gcommon.v1.metrics.ProviderEndpoints.AdditionalEndpointsEntry
	(*proto.Error)(nil),            // 7: gcommon.v1.common.Error
	(*timestamppb.Timestamp)(nil),  // 8: google.protobuf.Timestamp
	(*ProviderStatus)(nil),         // 9: gcommon.v1.metrics.ProviderStatus
	(*ValidationResult)(nil),       // 10: gcommon.v1.metrics.ValidationResult
}
var file_pkg_metrics_proto_responses_create_provider_response_proto_depIdxs = []int32{
	7,  // 0: gcommon.v1.metrics.CreateProviderResponse.error:type_name -> gcommon.v1.common.Error
	8,  // 1: gcommon.v1.metrics.CreateProviderResponse.created_at:type_name -> google.protobuf.Timestamp
	9,  // 2: gcommon.v1.metrics.CreateProviderResponse.status:type_name -> gcommon.v1.metrics.ProviderStatus
	10, // 3: gcommon.v1.metrics.CreateProviderResponse.validation:type_name -> gcommon.v1.metrics.ValidationResult
	1,  // 4: gcommon.v1.metrics.CreateProviderResponse.applied_config:type_name -> gcommon.v1.metrics.AppliedConfig
	3,  // 5: gcommon.v1.metrics.CreateProviderResponse.endpoints:type_name -> gcommon.v1.metrics.ProviderEndpoints
	4,  // 6: gcommon.v1.metrics.AppliedConfig.applied_defaults:type_name -> gcommon.v1.metrics.AppliedConfig.AppliedDefaultsEntry
	5,  // 7: gcommon.v1.metrics.AppliedConfig.applied_overrides:type_name -> gcommon.v1.metrics.AppliedConfig.AppliedOverridesEntry
	2,  // 8: gcommon.v1.metrics.AppliedConfig.resource_allocations:type_name -> gcommon.v1.metrics.ResourceAllocations
	6,  // 9: gcommon.v1.metrics.ProviderEndpoints.additional_endpoints:type_name -> gcommon.v1.metrics.ProviderEndpoints.AdditionalEndpointsEntry
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_pkg_metrics_proto_responses_create_provider_response_proto_init() }
func file_pkg_metrics_proto_responses_create_provider_response_proto_init() {
	if File_pkg_metrics_proto_responses_create_provider_response_proto != nil {
		return
	}
	file_pkg_metrics_proto_types_provider_status_proto_init()
	file_pkg_metrics_proto_enums_provider_state_proto_init()
	file_pkg_metrics_proto_types_validation_result_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pkg_metrics_proto_responses_create_provider_response_proto_rawDesc), len(file_pkg_metrics_proto_responses_create_provider_response_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_metrics_proto_responses_create_provider_response_proto_goTypes,
		DependencyIndexes: file_pkg_metrics_proto_responses_create_provider_response_proto_depIdxs,
		MessageInfos:      file_pkg_metrics_proto_responses_create_provider_response_proto_msgTypes,
	}.Build()
	File_pkg_metrics_proto_responses_create_provider_response_proto = out.File
	file_pkg_metrics_proto_responses_create_provider_response_proto_goTypes = nil
	file_pkg_metrics_proto_responses_create_provider_response_proto_depIdxs = nil
}

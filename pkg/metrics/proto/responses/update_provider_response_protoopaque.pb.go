// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: pkg/metrics/proto/responses/update_provider_response.proto

//go:build protoopaque

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
// UpdateAction indicates what action was taken during the update.
type UpdateAction int32

const (
	UpdateAction_UPDATE_ACTION_UNSPECIFIED UpdateAction = 0
	UpdateAction_UPDATE_ACTION_UPDATED     UpdateAction = 1
	UpdateAction_UPDATE_ACTION_NO_CHANGE   UpdateAction = 2
	UpdateAction_UPDATE_ACTION_RESTARTED   UpdateAction = 3
	UpdateAction_UPDATE_ACTION_RECREATED   UpdateAction = 4
)

// Enum value maps for UpdateAction.
var (
	UpdateAction_name = map[int32]string{
		0: "UPDATE_ACTION_UNSPECIFIED",
		1: "UPDATE_ACTION_UPDATED",
		2: "UPDATE_ACTION_NO_CHANGE",
		3: "UPDATE_ACTION_RESTARTED",
		4: "UPDATE_ACTION_RECREATED",
	}
	UpdateAction_value = map[string]int32{
		"UPDATE_ACTION_UNSPECIFIED": 0,
		"UPDATE_ACTION_UPDATED":     1,
		"UPDATE_ACTION_NO_CHANGE":   2,
		"UPDATE_ACTION_RESTARTED":   3,
		"UPDATE_ACTION_RECREATED":   4,
	}
)

func (x UpdateAction) Enum() *UpdateAction {
	p := new(UpdateAction)
	*p = x
	return p
}

func (x UpdateAction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UpdateAction) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_metrics_proto_responses_update_provider_response_proto_enumTypes[0].Descriptor()
}

func (UpdateAction) Type() protoreflect.EnumType {
	return &file_pkg_metrics_proto_responses_update_provider_response_proto_enumTypes[0]
}

func (x UpdateAction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// *
// ChangeType defines the type of configuration change.
type ChangeType int32

const (
	ChangeType_CHANGE_TYPE_UNSPECIFIED ChangeType = 0
	ChangeType_CHANGE_TYPE_ADDED       ChangeType = 1
	ChangeType_CHANGE_TYPE_UPDATED     ChangeType = 2
	ChangeType_CHANGE_TYPE_REMOVED     ChangeType = 3
	ChangeType_CHANGE_TYPE_REPLACED    ChangeType = 4
)

// Enum value maps for ChangeType.
var (
	ChangeType_name = map[int32]string{
		0: "CHANGE_TYPE_UNSPECIFIED",
		1: "CHANGE_TYPE_ADDED",
		2: "CHANGE_TYPE_UPDATED",
		3: "CHANGE_TYPE_REMOVED",
		4: "CHANGE_TYPE_REPLACED",
	}
	ChangeType_value = map[string]int32{
		"CHANGE_TYPE_UNSPECIFIED": 0,
		"CHANGE_TYPE_ADDED":       1,
		"CHANGE_TYPE_UPDATED":     2,
		"CHANGE_TYPE_REMOVED":     3,
		"CHANGE_TYPE_REPLACED":    4,
	}
)

func (x ChangeType) Enum() *ChangeType {
	p := new(ChangeType)
	*p = x
	return p
}

func (x ChangeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChangeType) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_metrics_proto_responses_update_provider_response_proto_enumTypes[1].Descriptor()
}

func (ChangeType) Type() protoreflect.EnumType {
	return &file_pkg_metrics_proto_responses_update_provider_response_proto_enumTypes[1]
}

func (x ChangeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// *
// UpdateProviderResponse contains the result of updating a metrics provider.
type UpdateProviderResponse struct {
	state                   protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Success      bool                   `protobuf:"varint,1,opt,name=success"`
	xxx_hidden_Error        *proto.Error           `protobuf:"bytes,2,opt,name=error"`
	xxx_hidden_ProviderId   *string                `protobuf:"bytes,3,opt,name=provider_id,json=providerId"`
	xxx_hidden_UpdatedAt    *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt"`
	xxx_hidden_Status       *ProviderStatus        `protobuf:"bytes,5,opt,name=status"`
	xxx_hidden_UpdateResult *UpdateResult          `protobuf:"bytes,6,opt,name=update_result,json=updateResult"`
	xxx_hidden_Validation   *ValidationResult      `protobuf:"bytes,7,opt,name=validation"`
	xxx_hidden_Warnings     []string               `protobuf:"bytes,8,rep,name=warnings"`
	xxx_hidden_BackupInfo   *BackupInfo            `protobuf:"bytes,9,opt,name=backup_info,json=backupInfo"`
	XXX_raceDetectHookData  protoimpl.RaceDetectHookData
	XXX_presence            [1]uint32
	unknownFields           protoimpl.UnknownFields
	sizeCache               protoimpl.SizeCache
}

func (x *UpdateProviderResponse) Reset() {
	*x = UpdateProviderResponse{}
	mi := &file_pkg_metrics_proto_responses_update_provider_response_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateProviderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProviderResponse) ProtoMessage() {}

func (x *UpdateProviderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_metrics_proto_responses_update_provider_response_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *UpdateProviderResponse) GetSuccess() bool {
	if x != nil {
		return x.xxx_hidden_Success
	}
	return false
}

func (x *UpdateProviderResponse) GetError() *proto.Error {
	if x != nil {
		return x.xxx_hidden_Error
	}
	return nil
}

func (x *UpdateProviderResponse) GetProviderId() string {
	if x != nil {
		if x.xxx_hidden_ProviderId != nil {
			return *x.xxx_hidden_ProviderId
		}
		return ""
	}
	return ""
}

func (x *UpdateProviderResponse) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.xxx_hidden_UpdatedAt
	}
	return nil
}

func (x *UpdateProviderResponse) GetStatus() *ProviderStatus {
	if x != nil {
		return x.xxx_hidden_Status
	}
	return nil
}

func (x *UpdateProviderResponse) GetUpdateResult() *UpdateResult {
	if x != nil {
		return x.xxx_hidden_UpdateResult
	}
	return nil
}

func (x *UpdateProviderResponse) GetValidation() *ValidationResult {
	if x != nil {
		return x.xxx_hidden_Validation
	}
	return nil
}

func (x *UpdateProviderResponse) GetWarnings() []string {
	if x != nil {
		return x.xxx_hidden_Warnings
	}
	return nil
}

func (x *UpdateProviderResponse) GetBackupInfo() *BackupInfo {
	if x != nil {
		return x.xxx_hidden_BackupInfo
	}
	return nil
}

func (x *UpdateProviderResponse) SetSuccess(v bool) {
	x.xxx_hidden_Success = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 9)
}

func (x *UpdateProviderResponse) SetError(v *proto.Error) {
	x.xxx_hidden_Error = v
}

func (x *UpdateProviderResponse) SetProviderId(v string) {
	x.xxx_hidden_ProviderId = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 2, 9)
}

func (x *UpdateProviderResponse) SetUpdatedAt(v *timestamppb.Timestamp) {
	x.xxx_hidden_UpdatedAt = v
}

func (x *UpdateProviderResponse) SetStatus(v *ProviderStatus) {
	x.xxx_hidden_Status = v
}

func (x *UpdateProviderResponse) SetUpdateResult(v *UpdateResult) {
	x.xxx_hidden_UpdateResult = v
}

func (x *UpdateProviderResponse) SetValidation(v *ValidationResult) {
	x.xxx_hidden_Validation = v
}

func (x *UpdateProviderResponse) SetWarnings(v []string) {
	x.xxx_hidden_Warnings = v
}

func (x *UpdateProviderResponse) SetBackupInfo(v *BackupInfo) {
	x.xxx_hidden_BackupInfo = v
}

func (x *UpdateProviderResponse) HasSuccess() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *UpdateProviderResponse) HasError() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Error != nil
}

func (x *UpdateProviderResponse) HasProviderId() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 2)
}

func (x *UpdateProviderResponse) HasUpdatedAt() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_UpdatedAt != nil
}

func (x *UpdateProviderResponse) HasStatus() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Status != nil
}

func (x *UpdateProviderResponse) HasUpdateResult() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_UpdateResult != nil
}

func (x *UpdateProviderResponse) HasValidation() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Validation != nil
}

func (x *UpdateProviderResponse) HasBackupInfo() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_BackupInfo != nil
}

func (x *UpdateProviderResponse) ClearSuccess() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_Success = false
}

func (x *UpdateProviderResponse) ClearError() {
	x.xxx_hidden_Error = nil
}

func (x *UpdateProviderResponse) ClearProviderId() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 2)
	x.xxx_hidden_ProviderId = nil
}

func (x *UpdateProviderResponse) ClearUpdatedAt() {
	x.xxx_hidden_UpdatedAt = nil
}

func (x *UpdateProviderResponse) ClearStatus() {
	x.xxx_hidden_Status = nil
}

func (x *UpdateProviderResponse) ClearUpdateResult() {
	x.xxx_hidden_UpdateResult = nil
}

func (x *UpdateProviderResponse) ClearValidation() {
	x.xxx_hidden_Validation = nil
}

func (x *UpdateProviderResponse) ClearBackupInfo() {
	x.xxx_hidden_BackupInfo = nil
}

type UpdateProviderResponse_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// Success status of the update
	Success *bool
	// Error information if update failed
	Error *proto.Error
	// Provider ID that was updated
	ProviderId *string
	// When the update was completed
	UpdatedAt *timestamppb.Timestamp
	// New status of the provider after update
	Status *ProviderStatus
	// Update results and changes applied
	UpdateResult *UpdateResult
	// Validation results
	Validation *ValidationResult
	// Warnings or informational messages
	Warnings []string
	// Backup information (if backup was created)
	BackupInfo *BackupInfo
}

func (b0 UpdateProviderResponse_builder) Build() *UpdateProviderResponse {
	m0 := &UpdateProviderResponse{}
	b, x := &b0, m0
	_, _ = b, x
	if b.Success != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 9)
		x.xxx_hidden_Success = *b.Success
	}
	x.xxx_hidden_Error = b.Error
	if b.ProviderId != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 2, 9)
		x.xxx_hidden_ProviderId = b.ProviderId
	}
	x.xxx_hidden_UpdatedAt = b.UpdatedAt
	x.xxx_hidden_Status = b.Status
	x.xxx_hidden_UpdateResult = b.UpdateResult
	x.xxx_hidden_Validation = b.Validation
	x.xxx_hidden_Warnings = b.Warnings
	x.xxx_hidden_BackupInfo = b.BackupInfo
	return m0
}

// *
// UpdateResult contains information about what was changed.
type UpdateResult struct {
	state                      protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Action          UpdateAction           `protobuf:"varint,1,opt,name=action,enum=gcommon.v1.metrics.UpdateAction"`
	xxx_hidden_ConfigChanges   *[]*ConfigChange       `protobuf:"bytes,2,rep,name=config_changes,json=configChanges"`
	xxx_hidden_UpdatedSettings []string               `protobuf:"bytes,3,rep,name=updated_settings,json=updatedSettings"`
	xxx_hidden_RemovedSettings []string               `protobuf:"bytes,4,rep,name=removed_settings,json=removedSettings"`
	xxx_hidden_Restarted       bool                   `protobuf:"varint,5,opt,name=restarted"`
	xxx_hidden_StrategyUsed    *string                `protobuf:"bytes,6,opt,name=strategy_used,json=strategyUsed"`
	xxx_hidden_UpdateDuration  *string                `protobuf:"bytes,7,opt,name=update_duration,json=updateDuration"`
	XXX_raceDetectHookData     protoimpl.RaceDetectHookData
	XXX_presence               [1]uint32
	unknownFields              protoimpl.UnknownFields
	sizeCache                  protoimpl.SizeCache
}

func (x *UpdateResult) Reset() {
	*x = UpdateResult{}
	mi := &file_pkg_metrics_proto_responses_update_provider_response_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResult) ProtoMessage() {}

func (x *UpdateResult) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_metrics_proto_responses_update_provider_response_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *UpdateResult) GetAction() UpdateAction {
	if x != nil {
		if protoimpl.X.Present(&(x.XXX_presence[0]), 0) {
			return x.xxx_hidden_Action
		}
	}
	return UpdateAction_UPDATE_ACTION_UNSPECIFIED
}

func (x *UpdateResult) GetConfigChanges() []*ConfigChange {
	if x != nil {
		if x.xxx_hidden_ConfigChanges != nil {
			return *x.xxx_hidden_ConfigChanges
		}
	}
	return nil
}

func (x *UpdateResult) GetUpdatedSettings() []string {
	if x != nil {
		return x.xxx_hidden_UpdatedSettings
	}
	return nil
}

func (x *UpdateResult) GetRemovedSettings() []string {
	if x != nil {
		return x.xxx_hidden_RemovedSettings
	}
	return nil
}

func (x *UpdateResult) GetRestarted() bool {
	if x != nil {
		return x.xxx_hidden_Restarted
	}
	return false
}

func (x *UpdateResult) GetStrategyUsed() string {
	if x != nil {
		if x.xxx_hidden_StrategyUsed != nil {
			return *x.xxx_hidden_StrategyUsed
		}
		return ""
	}
	return ""
}

func (x *UpdateResult) GetUpdateDuration() string {
	if x != nil {
		if x.xxx_hidden_UpdateDuration != nil {
			return *x.xxx_hidden_UpdateDuration
		}
		return ""
	}
	return ""
}

func (x *UpdateResult) SetAction(v UpdateAction) {
	x.xxx_hidden_Action = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 7)
}

func (x *UpdateResult) SetConfigChanges(v []*ConfigChange) {
	x.xxx_hidden_ConfigChanges = &v
}

func (x *UpdateResult) SetUpdatedSettings(v []string) {
	x.xxx_hidden_UpdatedSettings = v
}

func (x *UpdateResult) SetRemovedSettings(v []string) {
	x.xxx_hidden_RemovedSettings = v
}

func (x *UpdateResult) SetRestarted(v bool) {
	x.xxx_hidden_Restarted = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 4, 7)
}

func (x *UpdateResult) SetStrategyUsed(v string) {
	x.xxx_hidden_StrategyUsed = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 5, 7)
}

func (x *UpdateResult) SetUpdateDuration(v string) {
	x.xxx_hidden_UpdateDuration = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 6, 7)
}

func (x *UpdateResult) HasAction() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *UpdateResult) HasRestarted() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 4)
}

func (x *UpdateResult) HasStrategyUsed() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 5)
}

func (x *UpdateResult) HasUpdateDuration() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 6)
}

func (x *UpdateResult) ClearAction() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_Action = UpdateAction_UPDATE_ACTION_UNSPECIFIED
}

func (x *UpdateResult) ClearRestarted() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 4)
	x.xxx_hidden_Restarted = false
}

func (x *UpdateResult) ClearStrategyUsed() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 5)
	x.xxx_hidden_StrategyUsed = nil
}

func (x *UpdateResult) ClearUpdateDuration() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 6)
	x.xxx_hidden_UpdateDuration = nil
}

type UpdateResult_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// What update action was taken
	Action *UpdateAction
	// Configuration changes that were applied
	ConfigChanges []*ConfigChange
	// Settings that were updated
	UpdatedSettings []string
	// Settings that were removed
	RemovedSettings []string
	// Whether a restart occurred
	Restarted *bool
	// Update strategy that was used
	StrategyUsed *string
	// Time taken for the update
	UpdateDuration *string
}

func (b0 UpdateResult_builder) Build() *UpdateResult {
	m0 := &UpdateResult{}
	b, x := &b0, m0
	_, _ = b, x
	if b.Action != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 7)
		x.xxx_hidden_Action = *b.Action
	}
	x.xxx_hidden_ConfigChanges = &b.ConfigChanges
	x.xxx_hidden_UpdatedSettings = b.UpdatedSettings
	x.xxx_hidden_RemovedSettings = b.RemovedSettings
	if b.Restarted != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 4, 7)
		x.xxx_hidden_Restarted = *b.Restarted
	}
	if b.StrategyUsed != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 5, 7)
		x.xxx_hidden_StrategyUsed = b.StrategyUsed
	}
	if b.UpdateDuration != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 6, 7)
		x.xxx_hidden_UpdateDuration = b.UpdateDuration
	}
	return m0
}

// *
// ConfigChange describes a configuration change that was made.
type ConfigChange struct {
	state                  protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_ChangeType  ChangeType             `protobuf:"varint,1,opt,name=change_type,json=changeType,enum=gcommon.v1.metrics.ChangeType"`
	xxx_hidden_SettingPath *string                `protobuf:"bytes,2,opt,name=setting_path,json=settingPath"`
	xxx_hidden_OldValue    *string                `protobuf:"bytes,3,opt,name=old_value,json=oldValue"`
	xxx_hidden_NewValue    *string                `protobuf:"bytes,4,opt,name=new_value,json=newValue"`
	xxx_hidden_Description *string                `protobuf:"bytes,5,opt,name=description"`
	XXX_raceDetectHookData protoimpl.RaceDetectHookData
	XXX_presence           [1]uint32
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *ConfigChange) Reset() {
	*x = ConfigChange{}
	mi := &file_pkg_metrics_proto_responses_update_provider_response_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConfigChange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigChange) ProtoMessage() {}

func (x *ConfigChange) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_metrics_proto_responses_update_provider_response_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *ConfigChange) GetChangeType() ChangeType {
	if x != nil {
		if protoimpl.X.Present(&(x.XXX_presence[0]), 0) {
			return x.xxx_hidden_ChangeType
		}
	}
	return ChangeType_CHANGE_TYPE_UNSPECIFIED
}

func (x *ConfigChange) GetSettingPath() string {
	if x != nil {
		if x.xxx_hidden_SettingPath != nil {
			return *x.xxx_hidden_SettingPath
		}
		return ""
	}
	return ""
}

func (x *ConfigChange) GetOldValue() string {
	if x != nil {
		if x.xxx_hidden_OldValue != nil {
			return *x.xxx_hidden_OldValue
		}
		return ""
	}
	return ""
}

func (x *ConfigChange) GetNewValue() string {
	if x != nil {
		if x.xxx_hidden_NewValue != nil {
			return *x.xxx_hidden_NewValue
		}
		return ""
	}
	return ""
}

func (x *ConfigChange) GetDescription() string {
	if x != nil {
		if x.xxx_hidden_Description != nil {
			return *x.xxx_hidden_Description
		}
		return ""
	}
	return ""
}

func (x *ConfigChange) SetChangeType(v ChangeType) {
	x.xxx_hidden_ChangeType = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 5)
}

func (x *ConfigChange) SetSettingPath(v string) {
	x.xxx_hidden_SettingPath = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 1, 5)
}

func (x *ConfigChange) SetOldValue(v string) {
	x.xxx_hidden_OldValue = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 2, 5)
}

func (x *ConfigChange) SetNewValue(v string) {
	x.xxx_hidden_NewValue = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 3, 5)
}

func (x *ConfigChange) SetDescription(v string) {
	x.xxx_hidden_Description = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 4, 5)
}

func (x *ConfigChange) HasChangeType() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *ConfigChange) HasSettingPath() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 1)
}

func (x *ConfigChange) HasOldValue() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 2)
}

func (x *ConfigChange) HasNewValue() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 3)
}

func (x *ConfigChange) HasDescription() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 4)
}

func (x *ConfigChange) ClearChangeType() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_ChangeType = ChangeType_CHANGE_TYPE_UNSPECIFIED
}

func (x *ConfigChange) ClearSettingPath() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 1)
	x.xxx_hidden_SettingPath = nil
}

func (x *ConfigChange) ClearOldValue() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 2)
	x.xxx_hidden_OldValue = nil
}

func (x *ConfigChange) ClearNewValue() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 3)
	x.xxx_hidden_NewValue = nil
}

func (x *ConfigChange) ClearDescription() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 4)
	x.xxx_hidden_Description = nil
}

type ConfigChange_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// Type of change
	ChangeType *ChangeType
	// Setting that was changed
	SettingPath *string
	// Old value (if applicable)
	OldValue *string
	// New value
	NewValue *string
	// Description of the change
	Description *string
}

func (b0 ConfigChange_builder) Build() *ConfigChange {
	m0 := &ConfigChange{}
	b, x := &b0, m0
	_, _ = b, x
	if b.ChangeType != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 5)
		x.xxx_hidden_ChangeType = *b.ChangeType
	}
	if b.SettingPath != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 1, 5)
		x.xxx_hidden_SettingPath = b.SettingPath
	}
	if b.OldValue != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 2, 5)
		x.xxx_hidden_OldValue = b.OldValue
	}
	if b.NewValue != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 3, 5)
		x.xxx_hidden_NewValue = b.NewValue
	}
	if b.Description != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 4, 5)
		x.xxx_hidden_Description = b.Description
	}
	return m0
}

var File_pkg_metrics_proto_responses_update_provider_response_proto protoreflect.FileDescriptor

const file_pkg_metrics_proto_responses_update_provider_response_proto_rawDesc = "" +
	"\n" +
	":pkg/metrics/proto/responses/update_provider_response.proto\x12\x12gcommon.v1.metrics\x1a!google/protobuf/go_features.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a%pkg/common/proto/messages/error.proto\x1a-pkg/metrics/proto/types/provider_status.proto\x1a/pkg/metrics/proto/types/validation_result.proto\x1a)pkg/metrics/proto/types/backup_info.proto\"\xe4\x03\n" +
	"\x16UpdateProviderResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\x12.\n" +
	"\x05error\x18\x02 \x01(\v2\x18.gcommon.v1.common.ErrorR\x05error\x12\x1f\n" +
	"\vprovider_id\x18\x03 \x01(\tR\n" +
	"providerId\x129\n" +
	"\n" +
	"updated_at\x18\x04 \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAt\x12:\n" +
	"\x06status\x18\x05 \x01(\v2\".gcommon.v1.metrics.ProviderStatusR\x06status\x12E\n" +
	"\rupdate_result\x18\x06 \x01(\v2 .gcommon.v1.metrics.UpdateResultR\fupdateResult\x12D\n" +
	"\n" +
	"validation\x18\a \x01(\v2$.gcommon.v1.metrics.ValidationResultR\n" +
	"validation\x12\x1a\n" +
	"\bwarnings\x18\b \x03(\tR\bwarnings\x12?\n" +
	"\vbackup_info\x18\t \x01(\v2\x1e.gcommon.v1.metrics.BackupInfoR\n" +
	"backupInfo\"\xd3\x02\n" +
	"\fUpdateResult\x128\n" +
	"\x06action\x18\x01 \x01(\x0e2 .gcommon.v1.metrics.UpdateActionR\x06action\x12G\n" +
	"\x0econfig_changes\x18\x02 \x03(\v2 .gcommon.v1.metrics.ConfigChangeR\rconfigChanges\x12)\n" +
	"\x10updated_settings\x18\x03 \x03(\tR\x0fupdatedSettings\x12)\n" +
	"\x10removed_settings\x18\x04 \x03(\tR\x0fremovedSettings\x12\x1c\n" +
	"\trestarted\x18\x05 \x01(\bR\trestarted\x12#\n" +
	"\rstrategy_used\x18\x06 \x01(\tR\fstrategyUsed\x12'\n" +
	"\x0fupdate_duration\x18\a \x01(\tR\x0eupdateDuration\"\xce\x01\n" +
	"\fConfigChange\x12?\n" +
	"\vchange_type\x18\x01 \x01(\x0e2\x1e.gcommon.v1.metrics.ChangeTypeR\n" +
	"changeType\x12!\n" +
	"\fsetting_path\x18\x02 \x01(\tR\vsettingPath\x12\x1b\n" +
	"\told_value\x18\x03 \x01(\tR\boldValue\x12\x1b\n" +
	"\tnew_value\x18\x04 \x01(\tR\bnewValue\x12 \n" +
	"\vdescription\x18\x05 \x01(\tR\vdescription*\x9f\x01\n" +
	"\fUpdateAction\x12\x1d\n" +
	"\x19UPDATE_ACTION_UNSPECIFIED\x10\x00\x12\x19\n" +
	"\x15UPDATE_ACTION_UPDATED\x10\x01\x12\x1b\n" +
	"\x17UPDATE_ACTION_NO_CHANGE\x10\x02\x12\x1b\n" +
	"\x17UPDATE_ACTION_RESTARTED\x10\x03\x12\x1b\n" +
	"\x17UPDATE_ACTION_RECREATED\x10\x04*\x8c\x01\n" +
	"\n" +
	"ChangeType\x12\x1b\n" +
	"\x17CHANGE_TYPE_UNSPECIFIED\x10\x00\x12\x15\n" +
	"\x11CHANGE_TYPE_ADDED\x10\x01\x12\x17\n" +
	"\x13CHANGE_TYPE_UPDATED\x10\x02\x12\x17\n" +
	"\x13CHANGE_TYPE_REMOVED\x10\x03\x12\x18\n" +
	"\x14CHANGE_TYPE_REPLACED\x10\x04B\xde\x01\n" +
	"\x16com.gcommon.v1.metricsB\x1bUpdateProviderResponseProtoP\x01Z5github.com/jdfalk/gcommon/pkg/metrics/proto;metricspb\xa2\x02\x03GVM\xaa\x02\x12Gcommon.V1.Metrics\xca\x02\x12Gcommon\\V1\\Metrics\xe2\x02\x1eGcommon\\V1\\Metrics\\GPBMetadata\xea\x02\x14Gcommon::V1::Metrics\x92\x03\x05\xd2>\x02\x10\x02b\beditionsp\xe8\a"

var file_pkg_metrics_proto_responses_update_provider_response_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_pkg_metrics_proto_responses_update_provider_response_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pkg_metrics_proto_responses_update_provider_response_proto_goTypes = []any{
	(UpdateAction)(0),              // 0: gcommon.v1.metrics.UpdateAction
	(ChangeType)(0),                // 1: gcommon.v1.metrics.ChangeType
	(*UpdateProviderResponse)(nil), // 2: gcommon.v1.metrics.UpdateProviderResponse
	(*UpdateResult)(nil),           // 3: gcommon.v1.metrics.UpdateResult
	(*ConfigChange)(nil),           // 4: gcommon.v1.metrics.ConfigChange
	(*proto.Error)(nil),            // 5: gcommon.v1.common.Error
	(*timestamppb.Timestamp)(nil),  // 6: google.protobuf.Timestamp
	(*ProviderStatus)(nil),         // 7: gcommon.v1.metrics.ProviderStatus
	(*ValidationResult)(nil),       // 8: gcommon.v1.metrics.ValidationResult
	(*BackupInfo)(nil),             // 9: gcommon.v1.metrics.BackupInfo
}
var file_pkg_metrics_proto_responses_update_provider_response_proto_depIdxs = []int32{
	5, // 0: gcommon.v1.metrics.UpdateProviderResponse.error:type_name -> gcommon.v1.common.Error
	6, // 1: gcommon.v1.metrics.UpdateProviderResponse.updated_at:type_name -> google.protobuf.Timestamp
	7, // 2: gcommon.v1.metrics.UpdateProviderResponse.status:type_name -> gcommon.v1.metrics.ProviderStatus
	3, // 3: gcommon.v1.metrics.UpdateProviderResponse.update_result:type_name -> gcommon.v1.metrics.UpdateResult
	8, // 4: gcommon.v1.metrics.UpdateProviderResponse.validation:type_name -> gcommon.v1.metrics.ValidationResult
	9, // 5: gcommon.v1.metrics.UpdateProviderResponse.backup_info:type_name -> gcommon.v1.metrics.BackupInfo
	0, // 6: gcommon.v1.metrics.UpdateResult.action:type_name -> gcommon.v1.metrics.UpdateAction
	4, // 7: gcommon.v1.metrics.UpdateResult.config_changes:type_name -> gcommon.v1.metrics.ConfigChange
	1, // 8: gcommon.v1.metrics.ConfigChange.change_type:type_name -> gcommon.v1.metrics.ChangeType
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_pkg_metrics_proto_responses_update_provider_response_proto_init() }
func file_pkg_metrics_proto_responses_update_provider_response_proto_init() {
	if File_pkg_metrics_proto_responses_update_provider_response_proto != nil {
		return
	}
	file_pkg_metrics_proto_types_provider_status_proto_init()
	file_pkg_metrics_proto_types_validation_result_proto_init()
	file_pkg_metrics_proto_types_backup_info_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pkg_metrics_proto_responses_update_provider_response_proto_rawDesc), len(file_pkg_metrics_proto_responses_update_provider_response_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_metrics_proto_responses_update_provider_response_proto_goTypes,
		DependencyIndexes: file_pkg_metrics_proto_responses_update_provider_response_proto_depIdxs,
		EnumInfos:         file_pkg_metrics_proto_responses_update_provider_response_proto_enumTypes,
		MessageInfos:      file_pkg_metrics_proto_responses_update_provider_response_proto_msgTypes,
	}.Build()
	File_pkg_metrics_proto_responses_update_provider_response_proto = out.File
	file_pkg_metrics_proto_responses_update_provider_response_proto_goTypes = nil
	file_pkg_metrics_proto_responses_update_provider_response_proto_depIdxs = nil
}

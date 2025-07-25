// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: pkg/config/proto/messages/config_backup.proto

//go:build !protoopaque

package configpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/gofeaturespb"
	anypb "google.golang.org/protobuf/types/known/anypb"
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
// Represents a configuration backup.
// Stores complete configuration state for recovery purposes.
type ConfigBackup struct {
	state protoimpl.MessageState `protogen:"hybrid.v1"`
	// Unique identifier for this backup
	BackupId *string `protobuf:"bytes,1,opt,name=backup_id,json=backupId" json:"backup_id,omitempty"`
	// Timestamp when backup was created
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	// Configuration values at backup time
	ConfigValues map[string]*anypb.Any `protobuf:"bytes,3,rep,name=config_values,json=configValues" json:"config_values,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Version of the configuration
	Version *string `protobuf:"bytes,4,opt,name=version" json:"version,omitempty"`
	// Environment or context (e.g., "production", "staging")
	Environment *string `protobuf:"bytes,5,opt,name=environment" json:"environment,omitempty"`
	// User or service that created the backup
	CreatedBy *string `protobuf:"bytes,6,opt,name=created_by,json=createdBy" json:"created_by,omitempty"`
	// Description or reason for the backup
	Description *string `protobuf:"bytes,7,opt,name=description" json:"description,omitempty"`
	// Backup type (MANUAL, SCHEDULED, AUTOMATIC)
	BackupType *string `protobuf:"bytes,8,opt,name=backup_type,json=backupType" json:"backup_type,omitempty"`
	// Checksum or hash of the configuration
	Checksum *string `protobuf:"bytes,9,opt,name=checksum" json:"checksum,omitempty"`
	// Size of the backup in bytes
	SizeBytes *int64 `protobuf:"varint,10,opt,name=size_bytes,json=sizeBytes" json:"size_bytes,omitempty"`
	// Compression used (if any)
	Compression *string `protobuf:"bytes,11,opt,name=compression" json:"compression,omitempty"`
	// Storage location or path
	StoragePath *string `protobuf:"bytes,12,opt,name=storage_path,json=storagePath" json:"storage_path,omitempty"`
	// Retention policy for this backup
	RetentionPolicy *string `protobuf:"bytes,13,opt,name=retention_policy,json=retentionPolicy" json:"retention_policy,omitempty"`
	// Additional metadata
	Metadata      map[string]string `protobuf:"bytes,14,rep,name=metadata" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConfigBackup) Reset() {
	*x = ConfigBackup{}
	mi := &file_pkg_config_proto_messages_config_backup_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConfigBackup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigBackup) ProtoMessage() {}

func (x *ConfigBackup) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_config_proto_messages_config_backup_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *ConfigBackup) GetBackupId() string {
	if x != nil && x.BackupId != nil {
		return *x.BackupId
	}
	return ""
}

func (x *ConfigBackup) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *ConfigBackup) GetConfigValues() map[string]*anypb.Any {
	if x != nil {
		return x.ConfigValues
	}
	return nil
}

func (x *ConfigBackup) GetVersion() string {
	if x != nil && x.Version != nil {
		return *x.Version
	}
	return ""
}

func (x *ConfigBackup) GetEnvironment() string {
	if x != nil && x.Environment != nil {
		return *x.Environment
	}
	return ""
}

func (x *ConfigBackup) GetCreatedBy() string {
	if x != nil && x.CreatedBy != nil {
		return *x.CreatedBy
	}
	return ""
}

func (x *ConfigBackup) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *ConfigBackup) GetBackupType() string {
	if x != nil && x.BackupType != nil {
		return *x.BackupType
	}
	return ""
}

func (x *ConfigBackup) GetChecksum() string {
	if x != nil && x.Checksum != nil {
		return *x.Checksum
	}
	return ""
}

func (x *ConfigBackup) GetSizeBytes() int64 {
	if x != nil && x.SizeBytes != nil {
		return *x.SizeBytes
	}
	return 0
}

func (x *ConfigBackup) GetCompression() string {
	if x != nil && x.Compression != nil {
		return *x.Compression
	}
	return ""
}

func (x *ConfigBackup) GetStoragePath() string {
	if x != nil && x.StoragePath != nil {
		return *x.StoragePath
	}
	return ""
}

func (x *ConfigBackup) GetRetentionPolicy() string {
	if x != nil && x.RetentionPolicy != nil {
		return *x.RetentionPolicy
	}
	return ""
}

func (x *ConfigBackup) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *ConfigBackup) SetBackupId(v string) {
	x.BackupId = &v
}

func (x *ConfigBackup) SetCreatedAt(v *timestamppb.Timestamp) {
	x.CreatedAt = v
}

func (x *ConfigBackup) SetConfigValues(v map[string]*anypb.Any) {
	x.ConfigValues = v
}

func (x *ConfigBackup) SetVersion(v string) {
	x.Version = &v
}

func (x *ConfigBackup) SetEnvironment(v string) {
	x.Environment = &v
}

func (x *ConfigBackup) SetCreatedBy(v string) {
	x.CreatedBy = &v
}

func (x *ConfigBackup) SetDescription(v string) {
	x.Description = &v
}

func (x *ConfigBackup) SetBackupType(v string) {
	x.BackupType = &v
}

func (x *ConfigBackup) SetChecksum(v string) {
	x.Checksum = &v
}

func (x *ConfigBackup) SetSizeBytes(v int64) {
	x.SizeBytes = &v
}

func (x *ConfigBackup) SetCompression(v string) {
	x.Compression = &v
}

func (x *ConfigBackup) SetStoragePath(v string) {
	x.StoragePath = &v
}

func (x *ConfigBackup) SetRetentionPolicy(v string) {
	x.RetentionPolicy = &v
}

func (x *ConfigBackup) SetMetadata(v map[string]string) {
	x.Metadata = v
}

func (x *ConfigBackup) HasBackupId() bool {
	if x == nil {
		return false
	}
	return x.BackupId != nil
}

func (x *ConfigBackup) HasCreatedAt() bool {
	if x == nil {
		return false
	}
	return x.CreatedAt != nil
}

func (x *ConfigBackup) HasVersion() bool {
	if x == nil {
		return false
	}
	return x.Version != nil
}

func (x *ConfigBackup) HasEnvironment() bool {
	if x == nil {
		return false
	}
	return x.Environment != nil
}

func (x *ConfigBackup) HasCreatedBy() bool {
	if x == nil {
		return false
	}
	return x.CreatedBy != nil
}

func (x *ConfigBackup) HasDescription() bool {
	if x == nil {
		return false
	}
	return x.Description != nil
}

func (x *ConfigBackup) HasBackupType() bool {
	if x == nil {
		return false
	}
	return x.BackupType != nil
}

func (x *ConfigBackup) HasChecksum() bool {
	if x == nil {
		return false
	}
	return x.Checksum != nil
}

func (x *ConfigBackup) HasSizeBytes() bool {
	if x == nil {
		return false
	}
	return x.SizeBytes != nil
}

func (x *ConfigBackup) HasCompression() bool {
	if x == nil {
		return false
	}
	return x.Compression != nil
}

func (x *ConfigBackup) HasStoragePath() bool {
	if x == nil {
		return false
	}
	return x.StoragePath != nil
}

func (x *ConfigBackup) HasRetentionPolicy() bool {
	if x == nil {
		return false
	}
	return x.RetentionPolicy != nil
}

func (x *ConfigBackup) ClearBackupId() {
	x.BackupId = nil
}

func (x *ConfigBackup) ClearCreatedAt() {
	x.CreatedAt = nil
}

func (x *ConfigBackup) ClearVersion() {
	x.Version = nil
}

func (x *ConfigBackup) ClearEnvironment() {
	x.Environment = nil
}

func (x *ConfigBackup) ClearCreatedBy() {
	x.CreatedBy = nil
}

func (x *ConfigBackup) ClearDescription() {
	x.Description = nil
}

func (x *ConfigBackup) ClearBackupType() {
	x.BackupType = nil
}

func (x *ConfigBackup) ClearChecksum() {
	x.Checksum = nil
}

func (x *ConfigBackup) ClearSizeBytes() {
	x.SizeBytes = nil
}

func (x *ConfigBackup) ClearCompression() {
	x.Compression = nil
}

func (x *ConfigBackup) ClearStoragePath() {
	x.StoragePath = nil
}

func (x *ConfigBackup) ClearRetentionPolicy() {
	x.RetentionPolicy = nil
}

type ConfigBackup_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// Unique identifier for this backup
	BackupId *string
	// Timestamp when backup was created
	CreatedAt *timestamppb.Timestamp
	// Configuration values at backup time
	ConfigValues map[string]*anypb.Any
	// Version of the configuration
	Version *string
	// Environment or context (e.g., "production", "staging")
	Environment *string
	// User or service that created the backup
	CreatedBy *string
	// Description or reason for the backup
	Description *string
	// Backup type (MANUAL, SCHEDULED, AUTOMATIC)
	BackupType *string
	// Checksum or hash of the configuration
	Checksum *string
	// Size of the backup in bytes
	SizeBytes *int64
	// Compression used (if any)
	Compression *string
	// Storage location or path
	StoragePath *string
	// Retention policy for this backup
	RetentionPolicy *string
	// Additional metadata
	Metadata map[string]string
}

func (b0 ConfigBackup_builder) Build() *ConfigBackup {
	m0 := &ConfigBackup{}
	b, x := &b0, m0
	_, _ = b, x
	x.BackupId = b.BackupId
	x.CreatedAt = b.CreatedAt
	x.ConfigValues = b.ConfigValues
	x.Version = b.Version
	x.Environment = b.Environment
	x.CreatedBy = b.CreatedBy
	x.Description = b.Description
	x.BackupType = b.BackupType
	x.Checksum = b.Checksum
	x.SizeBytes = b.SizeBytes
	x.Compression = b.Compression
	x.StoragePath = b.StoragePath
	x.RetentionPolicy = b.RetentionPolicy
	x.Metadata = b.Metadata
	return m0
}

var File_pkg_config_proto_messages_config_backup_proto protoreflect.FileDescriptor

const file_pkg_config_proto_messages_config_backup_proto_rawDesc = "" +
	"\n" +
	"-pkg/config/proto/messages/config_backup.proto\x12\x11gcommon.v1.config\x1a!google/protobuf/go_features.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x19google/protobuf/any.proto\"\xe6\x05\n" +
	"\fConfigBackup\x12\x1b\n" +
	"\tbackup_id\x18\x01 \x01(\tR\bbackupId\x129\n" +
	"\n" +
	"created_at\x18\x02 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x12V\n" +
	"\rconfig_values\x18\x03 \x03(\v21.gcommon.v1.config.ConfigBackup.ConfigValuesEntryR\fconfigValues\x12\x18\n" +
	"\aversion\x18\x04 \x01(\tR\aversion\x12 \n" +
	"\venvironment\x18\x05 \x01(\tR\venvironment\x12\x1d\n" +
	"\n" +
	"created_by\x18\x06 \x01(\tR\tcreatedBy\x12 \n" +
	"\vdescription\x18\a \x01(\tR\vdescription\x12\x1f\n" +
	"\vbackup_type\x18\b \x01(\tR\n" +
	"backupType\x12\x1a\n" +
	"\bchecksum\x18\t \x01(\tR\bchecksum\x12\x1d\n" +
	"\n" +
	"size_bytes\x18\n" +
	" \x01(\x03R\tsizeBytes\x12 \n" +
	"\vcompression\x18\v \x01(\tR\vcompression\x12!\n" +
	"\fstorage_path\x18\f \x01(\tR\vstoragePath\x12)\n" +
	"\x10retention_policy\x18\r \x01(\tR\x0fretentionPolicy\x12I\n" +
	"\bmetadata\x18\x0e \x03(\v2-.gcommon.v1.config.ConfigBackup.MetadataEntryR\bmetadata\x1aU\n" +
	"\x11ConfigValuesEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12*\n" +
	"\x05value\x18\x02 \x01(\v2\x14.google.protobuf.AnyR\x05value:\x028\x01\x1a;\n" +
	"\rMetadataEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01B\xcd\x01\n" +
	"\x15com.gcommon.v1.configB\x11ConfigBackupProtoP\x01Z3github.com/jdfalk/gcommon/pkg/config/proto;configpb\xa2\x02\x03GVC\xaa\x02\x11Gcommon.V1.Config\xca\x02\x11Gcommon\\V1\\Config\xe2\x02\x1dGcommon\\V1\\Config\\GPBMetadata\xea\x02\x13Gcommon::V1::Config\x92\x03\x05\xd2>\x02\x10\x02b\beditionsp\xe8\a"

var file_pkg_config_proto_messages_config_backup_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pkg_config_proto_messages_config_backup_proto_goTypes = []any{
	(*ConfigBackup)(nil),          // 0: gcommon.v1.config.ConfigBackup
	nil,                           // 1: gcommon.v1.config.ConfigBackup.ConfigValuesEntry
	nil,                           // 2: gcommon.v1.config.ConfigBackup.MetadataEntry
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(*anypb.Any)(nil),             // 4: google.protobuf.Any
}
var file_pkg_config_proto_messages_config_backup_proto_depIdxs = []int32{
	3, // 0: gcommon.v1.config.ConfigBackup.created_at:type_name -> google.protobuf.Timestamp
	1, // 1: gcommon.v1.config.ConfigBackup.config_values:type_name -> gcommon.v1.config.ConfigBackup.ConfigValuesEntry
	2, // 2: gcommon.v1.config.ConfigBackup.metadata:type_name -> gcommon.v1.config.ConfigBackup.MetadataEntry
	4, // 3: gcommon.v1.config.ConfigBackup.ConfigValuesEntry.value:type_name -> google.protobuf.Any
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_pkg_config_proto_messages_config_backup_proto_init() }
func file_pkg_config_proto_messages_config_backup_proto_init() {
	if File_pkg_config_proto_messages_config_backup_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pkg_config_proto_messages_config_backup_proto_rawDesc), len(file_pkg_config_proto_messages_config_backup_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_config_proto_messages_config_backup_proto_goTypes,
		DependencyIndexes: file_pkg_config_proto_messages_config_backup_proto_depIdxs,
		MessageInfos:      file_pkg_config_proto_messages_config_backup_proto_msgTypes,
	}.Build()
	File_pkg_config_proto_messages_config_backup_proto = out.File
	file_pkg_config_proto_messages_config_backup_proto_goTypes = nil
	file_pkg_config_proto_messages_config_backup_proto_depIdxs = nil
}

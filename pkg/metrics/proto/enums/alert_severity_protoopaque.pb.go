// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: pkg/metrics/proto/enums/alert_severity.proto

//go:build protoopaque

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
// AlertSeverity defines the severity levels for metric alerts.
// Used to classify the urgency and impact of metric-based alerts.
type AlertSeverity int32

const (
	// Unspecified severity (default)
	AlertSeverity_ALERT_SEVERITY_UNSPECIFIED AlertSeverity = 0
	// Low severity - informational alerts
	AlertSeverity_ALERT_SEVERITY_LOW AlertSeverity = 1
	// Medium severity - warnings that need attention
	AlertSeverity_ALERT_SEVERITY_MEDIUM AlertSeverity = 2
	// High severity - serious issues requiring immediate attention
	AlertSeverity_ALERT_SEVERITY_HIGH AlertSeverity = 3
	// Critical severity - service-affecting issues requiring urgent response
	AlertSeverity_ALERT_SEVERITY_CRITICAL AlertSeverity = 4
	// Emergency severity - complete system failure or security incident
	AlertSeverity_ALERT_SEVERITY_EMERGENCY AlertSeverity = 5
)

// Enum value maps for AlertSeverity.
var (
	AlertSeverity_name = map[int32]string{
		0: "ALERT_SEVERITY_UNSPECIFIED",
		1: "ALERT_SEVERITY_LOW",
		2: "ALERT_SEVERITY_MEDIUM",
		3: "ALERT_SEVERITY_HIGH",
		4: "ALERT_SEVERITY_CRITICAL",
		5: "ALERT_SEVERITY_EMERGENCY",
	}
	AlertSeverity_value = map[string]int32{
		"ALERT_SEVERITY_UNSPECIFIED": 0,
		"ALERT_SEVERITY_LOW":         1,
		"ALERT_SEVERITY_MEDIUM":      2,
		"ALERT_SEVERITY_HIGH":        3,
		"ALERT_SEVERITY_CRITICAL":    4,
		"ALERT_SEVERITY_EMERGENCY":   5,
	}
)

func (x AlertSeverity) Enum() *AlertSeverity {
	p := new(AlertSeverity)
	*p = x
	return p
}

func (x AlertSeverity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AlertSeverity) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_metrics_proto_enums_alert_severity_proto_enumTypes[0].Descriptor()
}

func (AlertSeverity) Type() protoreflect.EnumType {
	return &file_pkg_metrics_proto_enums_alert_severity_proto_enumTypes[0]
}

func (x AlertSeverity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

var File_pkg_metrics_proto_enums_alert_severity_proto protoreflect.FileDescriptor

const file_pkg_metrics_proto_enums_alert_severity_proto_rawDesc = "" +
	"\n" +
	",pkg/metrics/proto/enums/alert_severity.proto\x12\x12gcommon.v1.metrics\x1a!google/protobuf/go_features.proto*\xb6\x01\n" +
	"\rAlertSeverity\x12\x1e\n" +
	"\x1aALERT_SEVERITY_UNSPECIFIED\x10\x00\x12\x16\n" +
	"\x12ALERT_SEVERITY_LOW\x10\x01\x12\x19\n" +
	"\x15ALERT_SEVERITY_MEDIUM\x10\x02\x12\x17\n" +
	"\x13ALERT_SEVERITY_HIGH\x10\x03\x12\x1b\n" +
	"\x17ALERT_SEVERITY_CRITICAL\x10\x04\x12\x1c\n" +
	"\x18ALERT_SEVERITY_EMERGENCY\x10\x05B\xd5\x01\n" +
	"\x16com.gcommon.v1.metricsB\x12AlertSeverityProtoP\x01Z5github.com/jdfalk/gcommon/pkg/metrics/proto;metricspb\xa2\x02\x03GVM\xaa\x02\x12Gcommon.V1.Metrics\xca\x02\x12Gcommon\\V1\\Metrics\xe2\x02\x1eGcommon\\V1\\Metrics\\GPBMetadata\xea\x02\x14Gcommon::V1::Metrics\x92\x03\x05\xd2>\x02\x10\x02b\beditionsp\xe8\a"

var file_pkg_metrics_proto_enums_alert_severity_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkg_metrics_proto_enums_alert_severity_proto_goTypes = []any{
	(AlertSeverity)(0), // 0: gcommon.v1.metrics.AlertSeverity
}
var file_pkg_metrics_proto_enums_alert_severity_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_metrics_proto_enums_alert_severity_proto_init() }
func file_pkg_metrics_proto_enums_alert_severity_proto_init() {
	if File_pkg_metrics_proto_enums_alert_severity_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pkg_metrics_proto_enums_alert_severity_proto_rawDesc), len(file_pkg_metrics_proto_enums_alert_severity_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_metrics_proto_enums_alert_severity_proto_goTypes,
		DependencyIndexes: file_pkg_metrics_proto_enums_alert_severity_proto_depIdxs,
		EnumInfos:         file_pkg_metrics_proto_enums_alert_severity_proto_enumTypes,
	}.Build()
	File_pkg_metrics_proto_enums_alert_severity_proto = out.File
	file_pkg_metrics_proto_enums_alert_severity_proto_goTypes = nil
	file_pkg_metrics_proto_enums_alert_severity_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: pkg/metrics/proto/enums/notification_type.proto

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
// NotificationType defines the types of notifications for metric alerts.
// Used to specify how alerts should be delivered to users.
type NotificationType int32

const (
	// Unspecified notification type (default)
	NotificationType_NOTIFICATION_TYPE_UNSPECIFIED NotificationType = 0
	// Email notification
	NotificationType_NOTIFICATION_TYPE_EMAIL NotificationType = 1
	// SMS text message
	NotificationType_NOTIFICATION_TYPE_SMS NotificationType = 2
	// Push notification (mobile app)
	NotificationType_NOTIFICATION_TYPE_PUSH NotificationType = 3
	// Slack message
	NotificationType_NOTIFICATION_TYPE_SLACK NotificationType = 4
	// Microsoft Teams message
	NotificationType_NOTIFICATION_TYPE_TEAMS NotificationType = 5
	// Discord message
	NotificationType_NOTIFICATION_TYPE_DISCORD NotificationType = 6
	// PagerDuty incident
	NotificationType_NOTIFICATION_TYPE_PAGERDUTY NotificationType = 7
	// Webhook/HTTP POST
	NotificationType_NOTIFICATION_TYPE_WEBHOOK NotificationType = 8
	// In-app notification
	NotificationType_NOTIFICATION_TYPE_IN_APP NotificationType = 9
	// SNMP trap
	NotificationType_NOTIFICATION_TYPE_SNMP NotificationType = 10
	// Telegram message
	NotificationType_NOTIFICATION_TYPE_TELEGRAM NotificationType = 11
	// Matrix message
	NotificationType_NOTIFICATION_TYPE_MATRIX NotificationType = 12
	// Voice call
	NotificationType_NOTIFICATION_TYPE_VOICE NotificationType = 13
)

// Enum value maps for NotificationType.
var (
	NotificationType_name = map[int32]string{
		0:  "NOTIFICATION_TYPE_UNSPECIFIED",
		1:  "NOTIFICATION_TYPE_EMAIL",
		2:  "NOTIFICATION_TYPE_SMS",
		3:  "NOTIFICATION_TYPE_PUSH",
		4:  "NOTIFICATION_TYPE_SLACK",
		5:  "NOTIFICATION_TYPE_TEAMS",
		6:  "NOTIFICATION_TYPE_DISCORD",
		7:  "NOTIFICATION_TYPE_PAGERDUTY",
		8:  "NOTIFICATION_TYPE_WEBHOOK",
		9:  "NOTIFICATION_TYPE_IN_APP",
		10: "NOTIFICATION_TYPE_SNMP",
		11: "NOTIFICATION_TYPE_TELEGRAM",
		12: "NOTIFICATION_TYPE_MATRIX",
		13: "NOTIFICATION_TYPE_VOICE",
	}
	NotificationType_value = map[string]int32{
		"NOTIFICATION_TYPE_UNSPECIFIED": 0,
		"NOTIFICATION_TYPE_EMAIL":       1,
		"NOTIFICATION_TYPE_SMS":         2,
		"NOTIFICATION_TYPE_PUSH":        3,
		"NOTIFICATION_TYPE_SLACK":       4,
		"NOTIFICATION_TYPE_TEAMS":       5,
		"NOTIFICATION_TYPE_DISCORD":     6,
		"NOTIFICATION_TYPE_PAGERDUTY":   7,
		"NOTIFICATION_TYPE_WEBHOOK":     8,
		"NOTIFICATION_TYPE_IN_APP":      9,
		"NOTIFICATION_TYPE_SNMP":        10,
		"NOTIFICATION_TYPE_TELEGRAM":    11,
		"NOTIFICATION_TYPE_MATRIX":      12,
		"NOTIFICATION_TYPE_VOICE":       13,
	}
)

func (x NotificationType) Enum() *NotificationType {
	p := new(NotificationType)
	*p = x
	return p
}

func (x NotificationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NotificationType) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_metrics_proto_enums_notification_type_proto_enumTypes[0].Descriptor()
}

func (NotificationType) Type() protoreflect.EnumType {
	return &file_pkg_metrics_proto_enums_notification_type_proto_enumTypes[0]
}

func (x NotificationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

var File_pkg_metrics_proto_enums_notification_type_proto protoreflect.FileDescriptor

const file_pkg_metrics_proto_enums_notification_type_proto_rawDesc = "" +
	"\n" +
	"/pkg/metrics/proto/enums/notification_type.proto\x12\x12gcommon.v1.metrics\x1a!google/protobuf/go_features.proto*\xb7\x03\n" +
	"\x10NotificationType\x12!\n" +
	"\x1dNOTIFICATION_TYPE_UNSPECIFIED\x10\x00\x12\x1b\n" +
	"\x17NOTIFICATION_TYPE_EMAIL\x10\x01\x12\x19\n" +
	"\x15NOTIFICATION_TYPE_SMS\x10\x02\x12\x1a\n" +
	"\x16NOTIFICATION_TYPE_PUSH\x10\x03\x12\x1b\n" +
	"\x17NOTIFICATION_TYPE_SLACK\x10\x04\x12\x1b\n" +
	"\x17NOTIFICATION_TYPE_TEAMS\x10\x05\x12\x1d\n" +
	"\x19NOTIFICATION_TYPE_DISCORD\x10\x06\x12\x1f\n" +
	"\x1bNOTIFICATION_TYPE_PAGERDUTY\x10\a\x12\x1d\n" +
	"\x19NOTIFICATION_TYPE_WEBHOOK\x10\b\x12\x1c\n" +
	"\x18NOTIFICATION_TYPE_IN_APP\x10\t\x12\x1a\n" +
	"\x16NOTIFICATION_TYPE_SNMP\x10\n" +
	"\x12\x1e\n" +
	"\x1aNOTIFICATION_TYPE_TELEGRAM\x10\v\x12\x1c\n" +
	"\x18NOTIFICATION_TYPE_MATRIX\x10\f\x12\x1b\n" +
	"\x17NOTIFICATION_TYPE_VOICE\x10\rB\xd8\x01\n" +
	"\x16com.gcommon.v1.metricsB\x15NotificationTypeProtoP\x01Z5github.com/jdfalk/gcommon/pkg/metrics/proto;metricspb\xa2\x02\x03GVM\xaa\x02\x12Gcommon.V1.Metrics\xca\x02\x12Gcommon\\V1\\Metrics\xe2\x02\x1eGcommon\\V1\\Metrics\\GPBMetadata\xea\x02\x14Gcommon::V1::Metrics\x92\x03\x05\xd2>\x02\x10\x02b\beditionsp\xe8\a"

var file_pkg_metrics_proto_enums_notification_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkg_metrics_proto_enums_notification_type_proto_goTypes = []any{
	(NotificationType)(0), // 0: gcommon.v1.metrics.NotificationType
}
var file_pkg_metrics_proto_enums_notification_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_metrics_proto_enums_notification_type_proto_init() }
func file_pkg_metrics_proto_enums_notification_type_proto_init() {
	if File_pkg_metrics_proto_enums_notification_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pkg_metrics_proto_enums_notification_type_proto_rawDesc), len(file_pkg_metrics_proto_enums_notification_type_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_metrics_proto_enums_notification_type_proto_goTypes,
		DependencyIndexes: file_pkg_metrics_proto_enums_notification_type_proto_depIdxs,
		EnumInfos:         file_pkg_metrics_proto_enums_notification_type_proto_enumTypes,
	}.Build()
	File_pkg_metrics_proto_enums_notification_type_proto = out.File
	file_pkg_metrics_proto_enums_notification_type_proto_goTypes = nil
	file_pkg_metrics_proto_enums_notification_type_proto_depIdxs = nil
}

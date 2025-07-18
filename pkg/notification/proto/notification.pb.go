// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: pkg/notification/proto/notification.proto

//go:build !protoopaque

package proto

import (
	enums "github.com/jdfalk/gcommon/pkg/notification/proto/enums"
	messages "github.com/jdfalk/gcommon/pkg/notification/proto/messages"
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

// Symbols defined in public import of pkg/notification/proto/messages/notification_message.proto.

type NotificationMessage = messages.NotificationMessage
type NotificationMessage_builder = messages.NotificationMessage_builder

// Symbols defined in public import of pkg/notification/proto/messages/delivery_channel.proto.

type DeliveryChannelType = messages.DeliveryChannelType

const DeliveryChannelType_DELIVERY_CHANNEL_TYPE_UNSPECIFIED = messages.DeliveryChannelType_DELIVERY_CHANNEL_TYPE_UNSPECIFIED
const DeliveryChannelType_DELIVERY_CHANNEL_TYPE_EMAIL = messages.DeliveryChannelType_DELIVERY_CHANNEL_TYPE_EMAIL
const DeliveryChannelType_DELIVERY_CHANNEL_TYPE_SMS = messages.DeliveryChannelType_DELIVERY_CHANNEL_TYPE_SMS
const DeliveryChannelType_DELIVERY_CHANNEL_TYPE_SLACK = messages.DeliveryChannelType_DELIVERY_CHANNEL_TYPE_SLACK
const DeliveryChannelType_DELIVERY_CHANNEL_TYPE_WEBHOOK = messages.DeliveryChannelType_DELIVERY_CHANNEL_TYPE_WEBHOOK

var DeliveryChannelType_name = messages.DeliveryChannelType_name
var DeliveryChannelType_value = messages.DeliveryChannelType_value

type DeliveryChannel = messages.DeliveryChannel
type DeliveryChannel_builder = messages.DeliveryChannel_builder

// Symbols defined in public import of pkg/notification/proto/messages/template.proto.

type Template = messages.Template
type Template_builder = messages.Template_builder

// Symbols defined in public import of pkg/notification/proto/messages/subscription_preferences.proto.

type SubscriptionPreferences = messages.SubscriptionPreferences
type SubscriptionPreferences_builder = messages.SubscriptionPreferences_builder

// Symbols defined in public import of pkg/notification/proto/messages/event_notification.proto.

type EventNotification = messages.EventNotification
type EventNotification_builder = messages.EventNotification_builder

// Symbols defined in public import of pkg/notification/proto/enums/delivery_status.proto.

type DeliveryStatus = enums.DeliveryStatus

const DeliveryStatus_DELIVERY_STATUS_UNSPECIFIED = enums.DeliveryStatus_DELIVERY_STATUS_UNSPECIFIED
const DeliveryStatus_DELIVERY_STATUS_PENDING = enums.DeliveryStatus_DELIVERY_STATUS_PENDING
const DeliveryStatus_DELIVERY_STATUS_SENT = enums.DeliveryStatus_DELIVERY_STATUS_SENT
const DeliveryStatus_DELIVERY_STATUS_FAILED = enums.DeliveryStatus_DELIVERY_STATUS_FAILED
const DeliveryStatus_DELIVERY_STATUS_ACKNOWLEDGED = enums.DeliveryStatus_DELIVERY_STATUS_ACKNOWLEDGED

var DeliveryStatus_name = enums.DeliveryStatus_name
var DeliveryStatus_value = enums.DeliveryStatus_value

var File_pkg_notification_proto_notification_proto protoreflect.FileDescriptor

const file_pkg_notification_proto_notification_proto_rawDesc = "" +
	"\n" +
	")pkg/notification/proto/notification.proto\x12\x17gcommon.v1.notification\x1a:pkg/notification/proto/messages/notification_message.proto\x1a6pkg/notification/proto/messages/delivery_channel.proto\x1a.pkg/notification/proto/messages/template.proto\x1a>pkg/notification/proto/messages/subscription_preferences.proto\x1a8pkg/notification/proto/messages/event_notification.proto\x1a2pkg/notification/proto/enums/delivery_status.proto\x1a!google/protobuf/go_features.protoB\xe8\x01\n" +
	"\x1bcom.gcommon.v1.notificationB\x11NotificationProtoP\x01Z0github.com/jdfalk/gcommon/pkg/notification/proto\xa2\x02\x03GVN\xaa\x02\x17Gcommon.V1.Notification\xca\x02\x17Gcommon\\V1\\Notification\xe2\x02#Gcommon\\V1\\Notification\\GPBMetadata\xea\x02\x19Gcommon::V1::Notification\x92\x03\x05\xd2>\x02\x10\x02P\x00P\x01P\x02P\x03P\x04P\x05b\beditionsp\xe8\a"

var file_pkg_notification_proto_notification_proto_goTypes = []any{}
var file_pkg_notification_proto_notification_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_notification_proto_notification_proto_init() }
func file_pkg_notification_proto_notification_proto_init() {
	if File_pkg_notification_proto_notification_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pkg_notification_proto_notification_proto_rawDesc), len(file_pkg_notification_proto_notification_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_notification_proto_notification_proto_goTypes,
		DependencyIndexes: file_pkg_notification_proto_notification_proto_depIdxs,
	}.Build()
	File_pkg_notification_proto_notification_proto = out.File
	file_pkg_notification_proto_notification_proto_goTypes = nil
	file_pkg_notification_proto_notification_proto_depIdxs = nil
}

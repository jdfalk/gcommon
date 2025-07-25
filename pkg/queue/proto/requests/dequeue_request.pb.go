// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: pkg/queue/proto/requests/dequeue_request.proto

package queuepb

import (
	proto "github.com/jdfalk/gcommon/pkg/common/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
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
// DequeueRequest retrieves one or more messages from a queue.
// Supports various consumption patterns including polling,
// long polling, and batch operations.
//
// Follows 1-1-1 pattern: one message per file.
type DequeueRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// *
	// The name/identifier of the queue to receive messages from.
	// Must be a valid existing queue.
	QueueName *string `protobuf:"bytes,1,opt,name=queue_name,json=queueName" json:"queue_name,omitempty"`
	// *
	// Standard request metadata including authentication context,
	// tracing information, and client details.
	Metadata *proto.RequestMetadata `protobuf:"bytes,11,opt,name=metadata" json:"metadata,omitempty"`
	// *
	// Maximum number of messages to receive in this request.
	// Range: 1-100. Default: 1.
	MaxMessages *int32 `protobuf:"varint,12,opt,name=max_messages,json=maxMessages" json:"max_messages,omitempty"`
	// *
	// Visibility timeout - how long the message is hidden from
	// other consumers after being received. Must be acknowledged
	// or rejected within this time. Default: queue configuration.
	VisibilityTimeout *durationpb.Duration `protobuf:"bytes,13,opt,name=visibility_timeout,json=visibilityTimeout" json:"visibility_timeout,omitempty"`
	// *
	// Wait time for long polling. If no messages are available,
	// the request will wait up to this duration for messages
	// to arrive. Set to 0 for immediate return.
	WaitTime *durationpb.Duration `protobuf:"bytes,14,opt,name=wait_time,json=waitTime" json:"wait_time,omitempty"`
	// *
	// Message group ID filter. If specified, only messages
	// from this group will be returned. Useful for ordered processing.
	GroupIdFilter *string `protobuf:"bytes,15,opt,name=group_id_filter,json=groupIdFilter" json:"group_id_filter,omitempty"`
	// *
	// Attribute filters for selective message consumption.
	// Only messages matching all specified attributes will be returned.
	AttributeFilters map[string]string `protobuf:"bytes,16,rep,name=attribute_filters,json=attributeFilters" json:"attribute_filters,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// *
	// Message type filter. If specified, only messages of
	// this type will be returned.
	MessageTypeFilter *string `protobuf:"bytes,17,opt,name=message_type_filter,json=messageTypeFilter" json:"message_type_filter,omitempty"`
	// *
	// Consumer ID for tracking and load balancing.
	// Helps with consumer group coordination and metrics.
	ConsumerId *string `protobuf:"bytes,18,opt,name=consumer_id,json=consumerId" json:"consumer_id,omitempty"`
	// *
	// Include message attributes in the response.
	// Default: true. Set to false to reduce response size.
	IncludeAttributes *bool `protobuf:"varint,19,opt,name=include_attributes,json=includeAttributes" json:"include_attributes,omitempty"`
	// *
	// Include message metadata (timestamps, delivery count, etc.)
	// in the response. Default: true.
	IncludeMetadata *bool `protobuf:"varint,20,opt,name=include_metadata,json=includeMetadata" json:"include_metadata,omitempty"`
	// *
	// Peek mode - return messages without removing them from
	// the queue. Useful for inspection. Default: false.
	PeekOnly *bool `protobuf:"varint,21,opt,name=peek_only,json=peekOnly" json:"peek_only,omitempty"`
	// *
	// Priority threshold - only return messages with priority
	// greater than or equal to this value. Range: 0-255.
	MinPriority   *int32 `protobuf:"varint,22,opt,name=min_priority,json=minPriority" json:"min_priority,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DequeueRequest) Reset() {
	*x = DequeueRequest{}
	mi := &file_pkg_queue_proto_requests_dequeue_request_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DequeueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DequeueRequest) ProtoMessage() {}

func (x *DequeueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_queue_proto_requests_dequeue_request_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DequeueRequest.ProtoReflect.Descriptor instead.
func (*DequeueRequest) Descriptor() ([]byte, []int) {
	return file_pkg_queue_proto_requests_dequeue_request_proto_rawDescGZIP(), []int{0}
}

func (x *DequeueRequest) GetQueueName() string {
	if x != nil && x.QueueName != nil {
		return *x.QueueName
	}
	return ""
}

func (x *DequeueRequest) GetMetadata() *proto.RequestMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *DequeueRequest) GetMaxMessages() int32 {
	if x != nil && x.MaxMessages != nil {
		return *x.MaxMessages
	}
	return 0
}

func (x *DequeueRequest) GetVisibilityTimeout() *durationpb.Duration {
	if x != nil {
		return x.VisibilityTimeout
	}
	return nil
}

func (x *DequeueRequest) GetWaitTime() *durationpb.Duration {
	if x != nil {
		return x.WaitTime
	}
	return nil
}

func (x *DequeueRequest) GetGroupIdFilter() string {
	if x != nil && x.GroupIdFilter != nil {
		return *x.GroupIdFilter
	}
	return ""
}

func (x *DequeueRequest) GetAttributeFilters() map[string]string {
	if x != nil {
		return x.AttributeFilters
	}
	return nil
}

func (x *DequeueRequest) GetMessageTypeFilter() string {
	if x != nil && x.MessageTypeFilter != nil {
		return *x.MessageTypeFilter
	}
	return ""
}

func (x *DequeueRequest) GetConsumerId() string {
	if x != nil && x.ConsumerId != nil {
		return *x.ConsumerId
	}
	return ""
}

func (x *DequeueRequest) GetIncludeAttributes() bool {
	if x != nil && x.IncludeAttributes != nil {
		return *x.IncludeAttributes
	}
	return false
}

func (x *DequeueRequest) GetIncludeMetadata() bool {
	if x != nil && x.IncludeMetadata != nil {
		return *x.IncludeMetadata
	}
	return false
}

func (x *DequeueRequest) GetPeekOnly() bool {
	if x != nil && x.PeekOnly != nil {
		return *x.PeekOnly
	}
	return false
}

func (x *DequeueRequest) GetMinPriority() int32 {
	if x != nil && x.MinPriority != nil {
		return *x.MinPriority
	}
	return 0
}

var File_pkg_queue_proto_requests_dequeue_request_proto protoreflect.FileDescriptor

const file_pkg_queue_proto_requests_dequeue_request_proto_rawDesc = "" +
	"\n" +
	".pkg/queue/proto/requests/dequeue_request.proto\x12\x10gcommon.v1.queue\x1a\x1egoogle/protobuf/duration.proto\x1a\x1dpkg/common/proto/common.proto\"\xd1\x05\n" +
	"\x0eDequeueRequest\x12\x1d\n" +
	"\n" +
	"queue_name\x18\x01 \x01(\tR\tqueueName\x12>\n" +
	"\bmetadata\x18\v \x01(\v2\".gcommon.v1.common.RequestMetadataR\bmetadata\x12!\n" +
	"\fmax_messages\x18\f \x01(\x05R\vmaxMessages\x12H\n" +
	"\x12visibility_timeout\x18\r \x01(\v2\x19.google.protobuf.DurationR\x11visibilityTimeout\x126\n" +
	"\twait_time\x18\x0e \x01(\v2\x19.google.protobuf.DurationR\bwaitTime\x12&\n" +
	"\x0fgroup_id_filter\x18\x0f \x01(\tR\rgroupIdFilter\x12c\n" +
	"\x11attribute_filters\x18\x10 \x03(\v26.gcommon.v1.queue.DequeueRequest.AttributeFiltersEntryR\x10attributeFilters\x12.\n" +
	"\x13message_type_filter\x18\x11 \x01(\tR\x11messageTypeFilter\x12\x1f\n" +
	"\vconsumer_id\x18\x12 \x01(\tR\n" +
	"consumerId\x12-\n" +
	"\x12include_attributes\x18\x13 \x01(\bR\x11includeAttributes\x12)\n" +
	"\x10include_metadata\x18\x14 \x01(\bR\x0fincludeMetadata\x12\x1b\n" +
	"\tpeek_only\x18\x15 \x01(\bR\bpeekOnly\x12!\n" +
	"\fmin_priority\x18\x16 \x01(\x05R\vminPriority\x1aC\n" +
	"\x15AttributeFiltersEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01B\xc0\x01\n" +
	"\x14com.gcommon.v1.queueB\x13DequeueRequestProtoP\x01Z1github.com/jdfalk/gcommon/pkg/queue/proto;queuepb\xa2\x02\x03GVQ\xaa\x02\x10Gcommon.V1.Queue\xca\x02\x10Gcommon\\V1\\Queue\xe2\x02\x1cGcommon\\V1\\Queue\\GPBMetadata\xea\x02\x12Gcommon::V1::Queueb\beditionsp\xe8\a"

var (
	file_pkg_queue_proto_requests_dequeue_request_proto_rawDescOnce sync.Once
	file_pkg_queue_proto_requests_dequeue_request_proto_rawDescData []byte
)

func file_pkg_queue_proto_requests_dequeue_request_proto_rawDescGZIP() []byte {
	file_pkg_queue_proto_requests_dequeue_request_proto_rawDescOnce.Do(func() {
		file_pkg_queue_proto_requests_dequeue_request_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_pkg_queue_proto_requests_dequeue_request_proto_rawDesc), len(file_pkg_queue_proto_requests_dequeue_request_proto_rawDesc)))
	})
	return file_pkg_queue_proto_requests_dequeue_request_proto_rawDescData
}

var file_pkg_queue_proto_requests_dequeue_request_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_queue_proto_requests_dequeue_request_proto_goTypes = []any{
	(*DequeueRequest)(nil),        // 0: gcommon.v1.queue.DequeueRequest
	nil,                           // 1: gcommon.v1.queue.DequeueRequest.AttributeFiltersEntry
	(*proto.RequestMetadata)(nil), // 2: gcommon.v1.common.RequestMetadata
	(*durationpb.Duration)(nil),   // 3: google.protobuf.Duration
}
var file_pkg_queue_proto_requests_dequeue_request_proto_depIdxs = []int32{
	2, // 0: gcommon.v1.queue.DequeueRequest.metadata:type_name -> gcommon.v1.common.RequestMetadata
	3, // 1: gcommon.v1.queue.DequeueRequest.visibility_timeout:type_name -> google.protobuf.Duration
	3, // 2: gcommon.v1.queue.DequeueRequest.wait_time:type_name -> google.protobuf.Duration
	1, // 3: gcommon.v1.queue.DequeueRequest.attribute_filters:type_name -> gcommon.v1.queue.DequeueRequest.AttributeFiltersEntry
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_pkg_queue_proto_requests_dequeue_request_proto_init() }
func file_pkg_queue_proto_requests_dequeue_request_proto_init() {
	if File_pkg_queue_proto_requests_dequeue_request_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pkg_queue_proto_requests_dequeue_request_proto_rawDesc), len(file_pkg_queue_proto_requests_dequeue_request_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_queue_proto_requests_dequeue_request_proto_goTypes,
		DependencyIndexes: file_pkg_queue_proto_requests_dequeue_request_proto_depIdxs,
		MessageInfos:      file_pkg_queue_proto_requests_dequeue_request_proto_msgTypes,
	}.Build()
	File_pkg_queue_proto_requests_dequeue_request_proto = out.File
	file_pkg_queue_proto_requests_dequeue_request_proto_goTypes = nil
	file_pkg_queue_proto_requests_dequeue_request_proto_depIdxs = nil
}

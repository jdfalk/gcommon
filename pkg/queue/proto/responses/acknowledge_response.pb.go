// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: pkg/queue/proto/responses/acknowledge_response.proto

package queuepb

import (
	proto "github.com/jdfalk/gcommon/pkg/common/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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
// MessageAckResult represents the acknowledgment result for a single message.
// Contains success status and any error information for individual messages
// within a batch acknowledgment request.
type MessageAckResult struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// *
	// Receipt handle of the message this result applies to.
	ReceiptHandle *string `protobuf:"bytes,1,opt,name=receipt_handle,json=receiptHandle" json:"receipt_handle,omitempty"`
	// *
	// Whether this specific message was successfully acknowledged.
	Success *bool `protobuf:"varint,2,opt,name=success" json:"success,omitempty"`
	// *
	// Error information if acknowledgment failed for this message.
	Error *proto.Error `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
	// *
	// Message ID for correlation (if available).
	MessageId *string `protobuf:"bytes,4,opt,name=message_id,json=messageId" json:"message_id,omitempty"`
	// *
	// Processing result that was recorded (echoed from request).
	ProcessingResult *string `protobuf:"bytes,5,opt,name=processing_result,json=processingResult" json:"processing_result,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *MessageAckResult) Reset() {
	*x = MessageAckResult{}
	mi := &file_pkg_queue_proto_responses_acknowledge_response_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageAckResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageAckResult) ProtoMessage() {}

func (x *MessageAckResult) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_queue_proto_responses_acknowledge_response_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageAckResult.ProtoReflect.Descriptor instead.
func (*MessageAckResult) Descriptor() ([]byte, []int) {
	return file_pkg_queue_proto_responses_acknowledge_response_proto_rawDescGZIP(), []int{0}
}

func (x *MessageAckResult) GetReceiptHandle() string {
	if x != nil && x.ReceiptHandle != nil {
		return *x.ReceiptHandle
	}
	return ""
}

func (x *MessageAckResult) GetSuccess() bool {
	if x != nil && x.Success != nil {
		return *x.Success
	}
	return false
}

func (x *MessageAckResult) GetError() *proto.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *MessageAckResult) GetMessageId() string {
	if x != nil && x.MessageId != nil {
		return *x.MessageId
	}
	return ""
}

func (x *MessageAckResult) GetProcessingResult() string {
	if x != nil && x.ProcessingResult != nil {
		return *x.ProcessingResult
	}
	return ""
}

// *
// AcknowledgeResponse confirms the acknowledgment of processed messages.
// Returns success status for each message and overall operation metrics
// for monitoring and debugging.
//
// Follows 1-1-1 pattern: one message per file.
type AcknowledgeResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// *
	// Overall success status of the acknowledgment operation.
	// True if all messages were successfully acknowledged.
	Success *bool `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	// *
	// Number of messages that were successfully acknowledged.
	AcknowledgedCount *int32 `protobuf:"varint,2,opt,name=acknowledged_count,json=acknowledgedCount" json:"acknowledged_count,omitempty"`
	// *
	// Number of messages that failed to be acknowledged.
	FailedCount *int32 `protobuf:"varint,3,opt,name=failed_count,json=failedCount" json:"failed_count,omitempty"`
	// *
	// Request processing metadata including timing, request ID,
	// and other observability information.
	RequestMetadata *proto.RequestMetadata `protobuf:"bytes,11,opt,name=request_metadata,json=requestMetadata" json:"request_metadata,omitempty"`
	// *
	// Name of the queue where messages were acknowledged.
	// Echoed from the request for verification.
	QueueName *string `protobuf:"bytes,12,opt,name=queue_name,json=queueName" json:"queue_name,omitempty"`
	// *
	// Detailed results for each message acknowledgment.
	// Only populated if there were failures or if detailed results were requested.
	MessageResults []*MessageAckResult `protobuf:"bytes,13,rep,name=message_results,json=messageResults" json:"message_results,omitempty"`
	// *
	// Consumer ID that was used for acknowledgment.
	// Echoed from the request for verification.
	ConsumerId *string `protobuf:"bytes,14,opt,name=consumer_id,json=consumerId" json:"consumer_id,omitempty"`
	// *
	// Total processing time for the acknowledgment operation in milliseconds.
	OperationTimeMs *int64 `protobuf:"varint,15,opt,name=operation_time_ms,json=operationTimeMs" json:"operation_time_ms,omitempty"`
	// *
	// Whether the operation was completed in batch mode.
	// Echoed from the request for verification.
	BatchMode *bool `protobuf:"varint,16,opt,name=batch_mode,json=batchMode" json:"batch_mode,omitempty"`
	// *
	// Number of messages that were already acknowledged (duplicates).
	// These don't count as failures but indicate potential issues.
	AlreadyAcknowledgedCount *int32 `protobuf:"varint,17,opt,name=already_acknowledged_count,json=alreadyAcknowledgedCount" json:"already_acknowledged_count,omitempty"`
	// *
	// Number of messages with expired visibility timeouts.
	// These may have been redelivered to other consumers.
	ExpiredTimeoutCount *int32 `protobuf:"varint,18,opt,name=expired_timeout_count,json=expiredTimeoutCount" json:"expired_timeout_count,omitempty"`
	// *
	// Error information if the overall acknowledgment operation failed
	// or completed with warnings.
	Error *proto.Error `protobuf:"bytes,61,opt,name=error" json:"error,omitempty"`
	// *
	// Timestamp when the acknowledgment operation was processed.
	AcknowledgedAt *timestamppb.Timestamp `protobuf:"bytes,51,opt,name=acknowledged_at,json=acknowledgedAt" json:"acknowledged_at,omitempty"`
	// *
	// Timestamp when this response was generated.
	ResponseGeneratedAt *timestamppb.Timestamp `protobuf:"bytes,52,opt,name=response_generated_at,json=responseGeneratedAt" json:"response_generated_at,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *AcknowledgeResponse) Reset() {
	*x = AcknowledgeResponse{}
	mi := &file_pkg_queue_proto_responses_acknowledge_response_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AcknowledgeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcknowledgeResponse) ProtoMessage() {}

func (x *AcknowledgeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_queue_proto_responses_acknowledge_response_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcknowledgeResponse.ProtoReflect.Descriptor instead.
func (*AcknowledgeResponse) Descriptor() ([]byte, []int) {
	return file_pkg_queue_proto_responses_acknowledge_response_proto_rawDescGZIP(), []int{1}
}

func (x *AcknowledgeResponse) GetSuccess() bool {
	if x != nil && x.Success != nil {
		return *x.Success
	}
	return false
}

func (x *AcknowledgeResponse) GetAcknowledgedCount() int32 {
	if x != nil && x.AcknowledgedCount != nil {
		return *x.AcknowledgedCount
	}
	return 0
}

func (x *AcknowledgeResponse) GetFailedCount() int32 {
	if x != nil && x.FailedCount != nil {
		return *x.FailedCount
	}
	return 0
}

func (x *AcknowledgeResponse) GetRequestMetadata() *proto.RequestMetadata {
	if x != nil {
		return x.RequestMetadata
	}
	return nil
}

func (x *AcknowledgeResponse) GetQueueName() string {
	if x != nil && x.QueueName != nil {
		return *x.QueueName
	}
	return ""
}

func (x *AcknowledgeResponse) GetMessageResults() []*MessageAckResult {
	if x != nil {
		return x.MessageResults
	}
	return nil
}

func (x *AcknowledgeResponse) GetConsumerId() string {
	if x != nil && x.ConsumerId != nil {
		return *x.ConsumerId
	}
	return ""
}

func (x *AcknowledgeResponse) GetOperationTimeMs() int64 {
	if x != nil && x.OperationTimeMs != nil {
		return *x.OperationTimeMs
	}
	return 0
}

func (x *AcknowledgeResponse) GetBatchMode() bool {
	if x != nil && x.BatchMode != nil {
		return *x.BatchMode
	}
	return false
}

func (x *AcknowledgeResponse) GetAlreadyAcknowledgedCount() int32 {
	if x != nil && x.AlreadyAcknowledgedCount != nil {
		return *x.AlreadyAcknowledgedCount
	}
	return 0
}

func (x *AcknowledgeResponse) GetExpiredTimeoutCount() int32 {
	if x != nil && x.ExpiredTimeoutCount != nil {
		return *x.ExpiredTimeoutCount
	}
	return 0
}

func (x *AcknowledgeResponse) GetError() *proto.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *AcknowledgeResponse) GetAcknowledgedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.AcknowledgedAt
	}
	return nil
}

func (x *AcknowledgeResponse) GetResponseGeneratedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ResponseGeneratedAt
	}
	return nil
}

var File_pkg_queue_proto_responses_acknowledge_response_proto protoreflect.FileDescriptor

const file_pkg_queue_proto_responses_acknowledge_response_proto_rawDesc = "" +
	"\n" +
	"4pkg/queue/proto/responses/acknowledge_response.proto\x12\x10gcommon.v1.queue\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1dpkg/common/proto/common.proto\"\xcf\x01\n" +
	"\x10MessageAckResult\x12%\n" +
	"\x0ereceipt_handle\x18\x01 \x01(\tR\rreceiptHandle\x12\x18\n" +
	"\asuccess\x18\x02 \x01(\bR\asuccess\x12.\n" +
	"\x05error\x18\x03 \x01(\v2\x18.gcommon.v1.common.ErrorR\x05error\x12\x1d\n" +
	"\n" +
	"message_id\x18\x04 \x01(\tR\tmessageId\x12+\n" +
	"\x11processing_result\x18\x05 \x01(\tR\x10processingResult\"\xdf\x05\n" +
	"\x13AcknowledgeResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\x12-\n" +
	"\x12acknowledged_count\x18\x02 \x01(\x05R\x11acknowledgedCount\x12!\n" +
	"\ffailed_count\x18\x03 \x01(\x05R\vfailedCount\x12M\n" +
	"\x10request_metadata\x18\v \x01(\v2\".gcommon.v1.common.RequestMetadataR\x0frequestMetadata\x12\x1d\n" +
	"\n" +
	"queue_name\x18\f \x01(\tR\tqueueName\x12K\n" +
	"\x0fmessage_results\x18\r \x03(\v2\".gcommon.v1.queue.MessageAckResultR\x0emessageResults\x12\x1f\n" +
	"\vconsumer_id\x18\x0e \x01(\tR\n" +
	"consumerId\x12*\n" +
	"\x11operation_time_ms\x18\x0f \x01(\x03R\x0foperationTimeMs\x12\x1d\n" +
	"\n" +
	"batch_mode\x18\x10 \x01(\bR\tbatchMode\x12<\n" +
	"\x1aalready_acknowledged_count\x18\x11 \x01(\x05R\x18alreadyAcknowledgedCount\x122\n" +
	"\x15expired_timeout_count\x18\x12 \x01(\x05R\x13expiredTimeoutCount\x12.\n" +
	"\x05error\x18= \x01(\v2\x18.gcommon.v1.common.ErrorR\x05error\x12C\n" +
	"\x0facknowledged_at\x183 \x01(\v2\x1a.google.protobuf.TimestampR\x0eacknowledgedAt\x12N\n" +
	"\x15response_generated_at\x184 \x01(\v2\x1a.google.protobuf.TimestampR\x13responseGeneratedAtB\xc5\x01\n" +
	"\x14com.gcommon.v1.queueB\x18AcknowledgeResponseProtoP\x01Z1github.com/jdfalk/gcommon/pkg/queue/proto;queuepb\xa2\x02\x03GVQ\xaa\x02\x10Gcommon.V1.Queue\xca\x02\x10Gcommon\\V1\\Queue\xe2\x02\x1cGcommon\\V1\\Queue\\GPBMetadata\xea\x02\x12Gcommon::V1::Queueb\beditionsp\xe8\a"

var (
	file_pkg_queue_proto_responses_acknowledge_response_proto_rawDescOnce sync.Once
	file_pkg_queue_proto_responses_acknowledge_response_proto_rawDescData []byte
)

func file_pkg_queue_proto_responses_acknowledge_response_proto_rawDescGZIP() []byte {
	file_pkg_queue_proto_responses_acknowledge_response_proto_rawDescOnce.Do(func() {
		file_pkg_queue_proto_responses_acknowledge_response_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_pkg_queue_proto_responses_acknowledge_response_proto_rawDesc), len(file_pkg_queue_proto_responses_acknowledge_response_proto_rawDesc)))
	})
	return file_pkg_queue_proto_responses_acknowledge_response_proto_rawDescData
}

var file_pkg_queue_proto_responses_acknowledge_response_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_queue_proto_responses_acknowledge_response_proto_goTypes = []any{
	(*MessageAckResult)(nil),      // 0: gcommon.v1.queue.MessageAckResult
	(*AcknowledgeResponse)(nil),   // 1: gcommon.v1.queue.AcknowledgeResponse
	(*proto.Error)(nil),           // 2: gcommon.v1.common.Error
	(*proto.RequestMetadata)(nil), // 3: gcommon.v1.common.RequestMetadata
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_pkg_queue_proto_responses_acknowledge_response_proto_depIdxs = []int32{
	2, // 0: gcommon.v1.queue.MessageAckResult.error:type_name -> gcommon.v1.common.Error
	3, // 1: gcommon.v1.queue.AcknowledgeResponse.request_metadata:type_name -> gcommon.v1.common.RequestMetadata
	0, // 2: gcommon.v1.queue.AcknowledgeResponse.message_results:type_name -> gcommon.v1.queue.MessageAckResult
	2, // 3: gcommon.v1.queue.AcknowledgeResponse.error:type_name -> gcommon.v1.common.Error
	4, // 4: gcommon.v1.queue.AcknowledgeResponse.acknowledged_at:type_name -> google.protobuf.Timestamp
	4, // 5: gcommon.v1.queue.AcknowledgeResponse.response_generated_at:type_name -> google.protobuf.Timestamp
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_pkg_queue_proto_responses_acknowledge_response_proto_init() }
func file_pkg_queue_proto_responses_acknowledge_response_proto_init() {
	if File_pkg_queue_proto_responses_acknowledge_response_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pkg_queue_proto_responses_acknowledge_response_proto_rawDesc), len(file_pkg_queue_proto_responses_acknowledge_response_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_queue_proto_responses_acknowledge_response_proto_goTypes,
		DependencyIndexes: file_pkg_queue_proto_responses_acknowledge_response_proto_depIdxs,
		MessageInfos:      file_pkg_queue_proto_responses_acknowledge_response_proto_msgTypes,
	}.Build()
	File_pkg_queue_proto_responses_acknowledge_response_proto = out.File
	file_pkg_queue_proto_responses_acknowledge_response_proto_goTypes = nil
	file_pkg_queue_proto_responses_acknowledge_response_proto_depIdxs = nil
}

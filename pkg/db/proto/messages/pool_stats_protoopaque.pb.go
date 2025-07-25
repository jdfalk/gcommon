// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: pkg/db/proto/messages/pool_stats.proto

//go:build protoopaque

package dbpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/gofeaturespb"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
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
// PoolStats provides detailed statistics about connection pool usage.
// Used for monitoring pool efficiency and connection management.
type PoolStats struct {
	state                          protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_TotalCreated        int64                  `protobuf:"varint,1,opt,name=total_created,json=totalCreated"`
	xxx_hidden_TotalClosed         int64                  `protobuf:"varint,2,opt,name=total_closed,json=totalClosed"`
	xxx_hidden_AcquisitionFailures int64                  `protobuf:"varint,3,opt,name=acquisition_failures,json=acquisitionFailures"`
	xxx_hidden_AvgAcquisitionTime  *durationpb.Duration   `protobuf:"bytes,4,opt,name=avg_acquisition_time,json=avgAcquisitionTime"`
	// Deprecated: Do not use. This will be deleted in the near future.
	XXX_lazyUnmarshalInfo  protoimpl.LazyUnmarshalInfo
	XXX_raceDetectHookData protoimpl.RaceDetectHookData
	XXX_presence           [1]uint32
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *PoolStats) Reset() {
	*x = PoolStats{}
	mi := &file_pkg_db_proto_messages_pool_stats_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PoolStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PoolStats) ProtoMessage() {}

func (x *PoolStats) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_db_proto_messages_pool_stats_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *PoolStats) GetTotalCreated() int64 {
	if x != nil {
		return x.xxx_hidden_TotalCreated
	}
	return 0
}

func (x *PoolStats) GetTotalClosed() int64 {
	if x != nil {
		return x.xxx_hidden_TotalClosed
	}
	return 0
}

func (x *PoolStats) GetAcquisitionFailures() int64 {
	if x != nil {
		return x.xxx_hidden_AcquisitionFailures
	}
	return 0
}

func (x *PoolStats) GetAvgAcquisitionTime() *durationpb.Duration {
	if x != nil {
		if protoimpl.X.Present(&(x.XXX_presence[0]), 3) {
			if protoimpl.X.AtomicCheckPointerIsNil(&x.xxx_hidden_AvgAcquisitionTime) {
				protoimpl.X.UnmarshalField(x, 4)
			}
			var rv *durationpb.Duration
			protoimpl.X.AtomicLoadPointer(protoimpl.Pointer(&x.xxx_hidden_AvgAcquisitionTime), protoimpl.Pointer(&rv))
			return rv
		}
	}
	return nil
}

func (x *PoolStats) SetTotalCreated(v int64) {
	x.xxx_hidden_TotalCreated = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 4)
}

func (x *PoolStats) SetTotalClosed(v int64) {
	x.xxx_hidden_TotalClosed = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 1, 4)
}

func (x *PoolStats) SetAcquisitionFailures(v int64) {
	x.xxx_hidden_AcquisitionFailures = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 2, 4)
}

func (x *PoolStats) SetAvgAcquisitionTime(v *durationpb.Duration) {
	protoimpl.X.AtomicSetPointer(&x.xxx_hidden_AvgAcquisitionTime, v)
	if v == nil {
		protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 3)
	} else {
		protoimpl.X.SetPresent(&(x.XXX_presence[0]), 3, 4)
	}
}

func (x *PoolStats) HasTotalCreated() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *PoolStats) HasTotalClosed() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 1)
}

func (x *PoolStats) HasAcquisitionFailures() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 2)
}

func (x *PoolStats) HasAvgAcquisitionTime() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 3)
}

func (x *PoolStats) ClearTotalCreated() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_TotalCreated = 0
}

func (x *PoolStats) ClearTotalClosed() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 1)
	x.xxx_hidden_TotalClosed = 0
}

func (x *PoolStats) ClearAcquisitionFailures() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 2)
	x.xxx_hidden_AcquisitionFailures = 0
}

func (x *PoolStats) ClearAvgAcquisitionTime() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 3)
	protoimpl.X.AtomicSetPointer(&x.xxx_hidden_AvgAcquisitionTime, (*durationpb.Duration)(nil))
}

type PoolStats_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// Total number of connections created since pool initialization
	TotalCreated *int64
	// Total number of connections closed since pool initialization
	TotalClosed *int64
	// Number of failed attempts to acquire connections
	AcquisitionFailures *int64
	// Average time to acquire a connection from the pool
	AvgAcquisitionTime *durationpb.Duration
}

func (b0 PoolStats_builder) Build() *PoolStats {
	m0 := &PoolStats{}
	b, x := &b0, m0
	_, _ = b, x
	if b.TotalCreated != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 4)
		x.xxx_hidden_TotalCreated = *b.TotalCreated
	}
	if b.TotalClosed != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 1, 4)
		x.xxx_hidden_TotalClosed = *b.TotalClosed
	}
	if b.AcquisitionFailures != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 2, 4)
		x.xxx_hidden_AcquisitionFailures = *b.AcquisitionFailures
	}
	if b.AvgAcquisitionTime != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 3, 4)
		x.xxx_hidden_AvgAcquisitionTime = b.AvgAcquisitionTime
	}
	return m0
}

var File_pkg_db_proto_messages_pool_stats_proto protoreflect.FileDescriptor

const file_pkg_db_proto_messages_pool_stats_proto_rawDesc = "" +
	"\n" +
	"&pkg/db/proto/messages/pool_stats.proto\x12\x13gcommon.v1.database\x1a\x1egoogle/protobuf/duration.proto\x1a!google/protobuf/go_features.proto\"\xd7\x01\n" +
	"\tPoolStats\x12#\n" +
	"\rtotal_created\x18\x01 \x01(\x03R\ftotalCreated\x12!\n" +
	"\ftotal_closed\x18\x02 \x01(\x03R\vtotalClosed\x121\n" +
	"\x14acquisition_failures\x18\x03 \x01(\x03R\x13acquisitionFailures\x12O\n" +
	"\x14avg_acquisition_time\x18\x04 \x01(\v2\x19.google.protobuf.DurationB\x02(\x01R\x12avgAcquisitionTimeB\xcc\x01\n" +
	"\x17com.gcommon.v1.databaseB\x0ePoolStatsProtoP\x01Z+github.com/jdfalk/gcommon/pkg/db/proto;dbpb\xa2\x02\x03GVD\xaa\x02\x13Gcommon.V1.Database\xca\x02\x13Gcommon\\V1\\Database\xe2\x02\x1fGcommon\\V1\\Database\\GPBMetadata\xea\x02\x15Gcommon::V1::Database\x92\x03\x05\xd2>\x02\x10\x02b\beditionsp\xe8\a"

var file_pkg_db_proto_messages_pool_stats_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pkg_db_proto_messages_pool_stats_proto_goTypes = []any{
	(*PoolStats)(nil),           // 0: gcommon.v1.database.PoolStats
	(*durationpb.Duration)(nil), // 1: google.protobuf.Duration
}
var file_pkg_db_proto_messages_pool_stats_proto_depIdxs = []int32{
	1, // 0: gcommon.v1.database.PoolStats.avg_acquisition_time:type_name -> google.protobuf.Duration
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_db_proto_messages_pool_stats_proto_init() }
func file_pkg_db_proto_messages_pool_stats_proto_init() {
	if File_pkg_db_proto_messages_pool_stats_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pkg_db_proto_messages_pool_stats_proto_rawDesc), len(file_pkg_db_proto_messages_pool_stats_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_db_proto_messages_pool_stats_proto_goTypes,
		DependencyIndexes: file_pkg_db_proto_messages_pool_stats_proto_depIdxs,
		MessageInfos:      file_pkg_db_proto_messages_pool_stats_proto_msgTypes,
	}.Build()
	File_pkg_db_proto_messages_pool_stats_proto = out.File
	file_pkg_db_proto_messages_pool_stats_proto_goTypes = nil
	file_pkg_db_proto_messages_pool_stats_proto_depIdxs = nil
}

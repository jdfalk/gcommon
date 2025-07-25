// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: pkg/db/proto/types/column_metadata.proto

//go:build protoopaque

package dbpb

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
// ColumnMetadata describes the structure and properties of a database column.
// Used in result sets to provide type information for proper data handling.
type ColumnMetadata struct {
	state               protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Name     *string                `protobuf:"bytes,1,opt,name=name"`
	xxx_hidden_Type     *string                `protobuf:"bytes,2,opt,name=type"`
	xxx_hidden_Nullable bool                   `protobuf:"varint,3,opt,name=nullable"`
	xxx_hidden_Size     int32                  `protobuf:"varint,4,opt,name=size"`
	xxx_hidden_Scale    int32                  `protobuf:"varint,5,opt,name=scale"`
	xxx_hidden_Metadata map[string]string      `protobuf:"bytes,6,rep,name=metadata" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Deprecated: Do not use. This will be deleted in the near future.
	XXX_lazyUnmarshalInfo  protoimpl.LazyUnmarshalInfo
	XXX_raceDetectHookData protoimpl.RaceDetectHookData
	XXX_presence           [1]uint32
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *ColumnMetadata) Reset() {
	*x = ColumnMetadata{}
	mi := &file_pkg_db_proto_types_column_metadata_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ColumnMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ColumnMetadata) ProtoMessage() {}

func (x *ColumnMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_db_proto_types_column_metadata_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *ColumnMetadata) GetName() string {
	if x != nil {
		if x.xxx_hidden_Name != nil {
			return *x.xxx_hidden_Name
		}
		return ""
	}
	return ""
}

func (x *ColumnMetadata) GetType() string {
	if x != nil {
		if x.xxx_hidden_Type != nil {
			return *x.xxx_hidden_Type
		}
		return ""
	}
	return ""
}

func (x *ColumnMetadata) GetNullable() bool {
	if x != nil {
		return x.xxx_hidden_Nullable
	}
	return false
}

func (x *ColumnMetadata) GetSize() int32 {
	if x != nil {
		return x.xxx_hidden_Size
	}
	return 0
}

func (x *ColumnMetadata) GetScale() int32 {
	if x != nil {
		return x.xxx_hidden_Scale
	}
	return 0
}

func (x *ColumnMetadata) GetMetadata() map[string]string {
	if x != nil {
		if protoimpl.X.Present(&(x.XXX_presence[0]), 5) {
			if protoimpl.X.AtomicCheckPointerIsNil(&x.xxx_hidden_Metadata) {
				protoimpl.X.UnmarshalField(x, 6)
			}
			return x.xxx_hidden_Metadata
		}
	}
	return nil
}

func (x *ColumnMetadata) SetName(v string) {
	x.xxx_hidden_Name = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 6)
}

func (x *ColumnMetadata) SetType(v string) {
	x.xxx_hidden_Type = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 1, 6)
}

func (x *ColumnMetadata) SetNullable(v bool) {
	x.xxx_hidden_Nullable = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 2, 6)
}

func (x *ColumnMetadata) SetSize(v int32) {
	x.xxx_hidden_Size = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 3, 6)
}

func (x *ColumnMetadata) SetScale(v int32) {
	x.xxx_hidden_Scale = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 4, 6)
}

func (x *ColumnMetadata) SetMetadata(v map[string]string) {
	x.xxx_hidden_Metadata = v
	if v == nil {
		protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 5)
	} else {
		protoimpl.X.SetPresent(&(x.XXX_presence[0]), 5, 6)
	}
}

func (x *ColumnMetadata) HasName() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *ColumnMetadata) HasType() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 1)
}

func (x *ColumnMetadata) HasNullable() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 2)
}

func (x *ColumnMetadata) HasSize() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 3)
}

func (x *ColumnMetadata) HasScale() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 4)
}

func (x *ColumnMetadata) ClearName() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_Name = nil
}

func (x *ColumnMetadata) ClearType() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 1)
	x.xxx_hidden_Type = nil
}

func (x *ColumnMetadata) ClearNullable() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 2)
	x.xxx_hidden_Nullable = false
}

func (x *ColumnMetadata) ClearSize() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 3)
	x.xxx_hidden_Size = 0
}

func (x *ColumnMetadata) ClearScale() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 4)
	x.xxx_hidden_Scale = 0
}

type ColumnMetadata_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// Column name as defined in the database
	Name *string
	// Column data type (database-specific)
	Type *string
	// Whether the column allows NULL values
	Nullable *bool
	// Column size/precision for numeric and string types
	Size *int32
	// Column scale for decimal/numeric types
	Scale *int32
	// Additional column-specific metadata
	Metadata map[string]string
}

func (b0 ColumnMetadata_builder) Build() *ColumnMetadata {
	m0 := &ColumnMetadata{}
	b, x := &b0, m0
	_, _ = b, x
	if b.Name != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 6)
		x.xxx_hidden_Name = b.Name
	}
	if b.Type != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 1, 6)
		x.xxx_hidden_Type = b.Type
	}
	if b.Nullable != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 2, 6)
		x.xxx_hidden_Nullable = *b.Nullable
	}
	if b.Size != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 3, 6)
		x.xxx_hidden_Size = *b.Size
	}
	if b.Scale != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 4, 6)
		x.xxx_hidden_Scale = *b.Scale
	}
	if b.Metadata != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 5, 6)
		x.xxx_hidden_Metadata = b.Metadata
	}
	return m0
}

var File_pkg_db_proto_types_column_metadata_proto protoreflect.FileDescriptor

const file_pkg_db_proto_types_column_metadata_proto_rawDesc = "" +
	"\n" +
	"(pkg/db/proto/types/column_metadata.proto\x12\x13gcommon.v1.database\x1a!google/protobuf/go_features.proto\"\x8e\x02\n" +
	"\x0eColumnMetadata\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x12\n" +
	"\x04type\x18\x02 \x01(\tR\x04type\x12\x1a\n" +
	"\bnullable\x18\x03 \x01(\bR\bnullable\x12\x12\n" +
	"\x04size\x18\x04 \x01(\x05R\x04size\x12\x14\n" +
	"\x05scale\x18\x05 \x01(\x05R\x05scale\x12Q\n" +
	"\bmetadata\x18\x06 \x03(\v21.gcommon.v1.database.ColumnMetadata.MetadataEntryB\x02(\x01R\bmetadata\x1a;\n" +
	"\rMetadataEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01B\xd1\x01\n" +
	"\x17com.gcommon.v1.databaseB\x13ColumnMetadataProtoP\x01Z+github.com/jdfalk/gcommon/pkg/db/proto;dbpb\xa2\x02\x03GVD\xaa\x02\x13Gcommon.V1.Database\xca\x02\x13Gcommon\\V1\\Database\xe2\x02\x1fGcommon\\V1\\Database\\GPBMetadata\xea\x02\x15Gcommon::V1::Database\x92\x03\x05\xd2>\x02\x10\x02b\beditionsp\xe8\a"

var file_pkg_db_proto_types_column_metadata_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_db_proto_types_column_metadata_proto_goTypes = []any{
	(*ColumnMetadata)(nil), // 0: gcommon.v1.database.ColumnMetadata
	nil,                    // 1: gcommon.v1.database.ColumnMetadata.MetadataEntry
}
var file_pkg_db_proto_types_column_metadata_proto_depIdxs = []int32{
	1, // 0: gcommon.v1.database.ColumnMetadata.metadata:type_name -> gcommon.v1.database.ColumnMetadata.MetadataEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_db_proto_types_column_metadata_proto_init() }
func file_pkg_db_proto_types_column_metadata_proto_init() {
	if File_pkg_db_proto_types_column_metadata_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pkg_db_proto_types_column_metadata_proto_rawDesc), len(file_pkg_db_proto_types_column_metadata_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_db_proto_types_column_metadata_proto_goTypes,
		DependencyIndexes: file_pkg_db_proto_types_column_metadata_proto_depIdxs,
		MessageInfos:      file_pkg_db_proto_types_column_metadata_proto_msgTypes,
	}.Build()
	File_pkg_db_proto_types_column_metadata_proto = out.File
	file_pkg_db_proto_types_column_metadata_proto_goTypes = nil
	file_pkg_db_proto_types_column_metadata_proto_depIdxs = nil
}

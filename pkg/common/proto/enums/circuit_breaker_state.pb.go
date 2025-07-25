// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: pkg/common/proto/enums/circuit_breaker_state.proto

//go:build !protoopaque

package commonpb

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
// Circuit breaker state enumeration for fault tolerance patterns.
// Defines the current state of circuit breaker components used
// for resilience and stability across all GCommon modules.
type CircuitBreakerState int32

const (
	// Default value indicating no circuit breaker state was specified
	CircuitBreakerState_CIRCUIT_BREAKER_STATE_UNSPECIFIED CircuitBreakerState = 0
	// Circuit is closed - requests are flowing normally
	CircuitBreakerState_CIRCUIT_BREAKER_STATE_CLOSED CircuitBreakerState = 1
	// Circuit is open - requests are blocked due to failures
	CircuitBreakerState_CIRCUIT_BREAKER_STATE_OPEN CircuitBreakerState = 2
	// Circuit is half-open - testing if service has recovered
	CircuitBreakerState_CIRCUIT_BREAKER_STATE_HALF_OPEN CircuitBreakerState = 3
)

// Enum value maps for CircuitBreakerState.
var (
	CircuitBreakerState_name = map[int32]string{
		0: "CIRCUIT_BREAKER_STATE_UNSPECIFIED",
		1: "CIRCUIT_BREAKER_STATE_CLOSED",
		2: "CIRCUIT_BREAKER_STATE_OPEN",
		3: "CIRCUIT_BREAKER_STATE_HALF_OPEN",
	}
	CircuitBreakerState_value = map[string]int32{
		"CIRCUIT_BREAKER_STATE_UNSPECIFIED": 0,
		"CIRCUIT_BREAKER_STATE_CLOSED":      1,
		"CIRCUIT_BREAKER_STATE_OPEN":        2,
		"CIRCUIT_BREAKER_STATE_HALF_OPEN":   3,
	}
)

func (x CircuitBreakerState) Enum() *CircuitBreakerState {
	p := new(CircuitBreakerState)
	*p = x
	return p
}

func (x CircuitBreakerState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CircuitBreakerState) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_common_proto_enums_circuit_breaker_state_proto_enumTypes[0].Descriptor()
}

func (CircuitBreakerState) Type() protoreflect.EnumType {
	return &file_pkg_common_proto_enums_circuit_breaker_state_proto_enumTypes[0]
}

func (x CircuitBreakerState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

var File_pkg_common_proto_enums_circuit_breaker_state_proto protoreflect.FileDescriptor

const file_pkg_common_proto_enums_circuit_breaker_state_proto_rawDesc = "" +
	"\n" +
	"2pkg/common/proto/enums/circuit_breaker_state.proto\x12\x11gcommon.v1.common\x1a!google/protobuf/go_features.proto*\xa3\x01\n" +
	"\x13CircuitBreakerState\x12%\n" +
	"!CIRCUIT_BREAKER_STATE_UNSPECIFIED\x10\x00\x12 \n" +
	"\x1cCIRCUIT_BREAKER_STATE_CLOSED\x10\x01\x12\x1e\n" +
	"\x1aCIRCUIT_BREAKER_STATE_OPEN\x10\x02\x12#\n" +
	"\x1fCIRCUIT_BREAKER_STATE_HALF_OPEN\x10\x03B\xd4\x01\n" +
	"\x15com.gcommon.v1.commonB\x18CircuitBreakerStateProtoP\x01Z3github.com/jdfalk/gcommon/pkg/common/proto;commonpb\xa2\x02\x03GVC\xaa\x02\x11Gcommon.V1.Common\xca\x02\x11Gcommon\\V1\\Common\xe2\x02\x1dGcommon\\V1\\Common\\GPBMetadata\xea\x02\x13Gcommon::V1::Common\x92\x03\x05\xd2>\x02\x10\x02b\beditionsp\xe8\a"

var file_pkg_common_proto_enums_circuit_breaker_state_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkg_common_proto_enums_circuit_breaker_state_proto_goTypes = []any{
	(CircuitBreakerState)(0), // 0: gcommon.v1.common.CircuitBreakerState
}
var file_pkg_common_proto_enums_circuit_breaker_state_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_common_proto_enums_circuit_breaker_state_proto_init() }
func file_pkg_common_proto_enums_circuit_breaker_state_proto_init() {
	if File_pkg_common_proto_enums_circuit_breaker_state_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pkg_common_proto_enums_circuit_breaker_state_proto_rawDesc), len(file_pkg_common_proto_enums_circuit_breaker_state_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_common_proto_enums_circuit_breaker_state_proto_goTypes,
		DependencyIndexes: file_pkg_common_proto_enums_circuit_breaker_state_proto_depIdxs,
		EnumInfos:         file_pkg_common_proto_enums_circuit_breaker_state_proto_enumTypes,
	}.Build()
	File_pkg_common_proto_enums_circuit_breaker_state_proto = out.File
	file_pkg_common_proto_enums_circuit_breaker_state_proto_goTypes = nil
	file_pkg_common_proto_enums_circuit_breaker_state_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: pkg/metrics/proto/messages/metric_filter.proto

//go:build !protoopaque

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

// Field to sort by
type SortCriteria_SortField int32

const (
	SortCriteria_SORT_FIELD_UNSPECIFIED SortCriteria_SortField = 0
	SortCriteria_SORT_FIELD_NAME        SortCriteria_SortField = 1
	SortCriteria_SORT_FIELD_TIMESTAMP   SortCriteria_SortField = 2
	SortCriteria_SORT_FIELD_VALUE       SortCriteria_SortField = 3
	SortCriteria_SORT_FIELD_CREATED_AT  SortCriteria_SortField = 4
)

// Enum value maps for SortCriteria_SortField.
var (
	SortCriteria_SortField_name = map[int32]string{
		0: "SORT_FIELD_UNSPECIFIED",
		1: "SORT_FIELD_NAME",
		2: "SORT_FIELD_TIMESTAMP",
		3: "SORT_FIELD_VALUE",
		4: "SORT_FIELD_CREATED_AT",
	}
	SortCriteria_SortField_value = map[string]int32{
		"SORT_FIELD_UNSPECIFIED": 0,
		"SORT_FIELD_NAME":        1,
		"SORT_FIELD_TIMESTAMP":   2,
		"SORT_FIELD_VALUE":       3,
		"SORT_FIELD_CREATED_AT":  4,
	}
)

func (x SortCriteria_SortField) Enum() *SortCriteria_SortField {
	p := new(SortCriteria_SortField)
	*p = x
	return p
}

func (x SortCriteria_SortField) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SortCriteria_SortField) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_metrics_proto_messages_metric_filter_proto_enumTypes[0].Descriptor()
}

func (SortCriteria_SortField) Type() protoreflect.EnumType {
	return &file_pkg_metrics_proto_messages_metric_filter_proto_enumTypes[0]
}

func (x SortCriteria_SortField) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Sort direction
type SortCriteria_SortOrder int32

const (
	SortCriteria_SORT_ORDER_UNSPECIFIED SortCriteria_SortOrder = 0
	SortCriteria_SORT_ORDER_ASCENDING   SortCriteria_SortOrder = 1
	SortCriteria_SORT_ORDER_DESCENDING  SortCriteria_SortOrder = 2
)

// Enum value maps for SortCriteria_SortOrder.
var (
	SortCriteria_SortOrder_name = map[int32]string{
		0: "SORT_ORDER_UNSPECIFIED",
		1: "SORT_ORDER_ASCENDING",
		2: "SORT_ORDER_DESCENDING",
	}
	SortCriteria_SortOrder_value = map[string]int32{
		"SORT_ORDER_UNSPECIFIED": 0,
		"SORT_ORDER_ASCENDING":   1,
		"SORT_ORDER_DESCENDING":  2,
	}
)

func (x SortCriteria_SortOrder) Enum() *SortCriteria_SortOrder {
	p := new(SortCriteria_SortOrder)
	*p = x
	return p
}

func (x SortCriteria_SortOrder) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SortCriteria_SortOrder) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_metrics_proto_messages_metric_filter_proto_enumTypes[1].Descriptor()
}

func (SortCriteria_SortOrder) Type() protoreflect.EnumType {
	return &file_pkg_metrics_proto_messages_metric_filter_proto_enumTypes[1]
}

func (x SortCriteria_SortOrder) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// *
// MetricFilter defines criteria for filtering metrics in queries.
// Used to specify which metrics to include/exclude from operations.
type MetricFilter struct {
	state protoimpl.MessageState `protogen:"hybrid.v1"`
	// Filter by metric name (supports wildcards and regex)
	MetricNames []string `protobuf:"bytes,1,rep,name=metric_names,json=metricNames" json:"metric_names,omitempty"`
	// Filter by metric type
	MetricTypes []MetricType `protobuf:"varint,2,rep,packed,name=metric_types,json=metricTypes,enum=gcommon.v1.metrics.MetricType" json:"metric_types,omitempty"`
	// Filter by labels/tags
	LabelFilters []*LabelFilter `protobuf:"bytes,3,rep,name=label_filters,json=labelFilters" json:"label_filters,omitempty"`
	// Filter by namespace
	Namespaces []string `protobuf:"bytes,4,rep,name=namespaces" json:"namespaces,omitempty"`
	// Filter by source system
	Sources []string `protobuf:"bytes,5,rep,name=sources" json:"sources,omitempty"`
	// Time range filter
	TimeFilter *TimeFilter `protobuf:"bytes,6,opt,name=time_filter,json=timeFilter" json:"time_filter,omitempty"`
	// Filter by value range
	ValueFilters []*ValueFilter `protobuf:"bytes,7,rep,name=value_filters,json=valueFilters" json:"value_filters,omitempty"`
	// Limit number of results
	Limit *int32 `protobuf:"varint,8,opt,name=limit" json:"limit,omitempty"`
	// Skip/offset for pagination
	Offset        *int32 `protobuf:"varint,9,opt,name=offset" json:"offset,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MetricFilter) Reset() {
	*x = MetricFilter{}
	mi := &file_pkg_metrics_proto_messages_metric_filter_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MetricFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricFilter) ProtoMessage() {}

func (x *MetricFilter) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_metrics_proto_messages_metric_filter_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *MetricFilter) GetMetricNames() []string {
	if x != nil {
		return x.MetricNames
	}
	return nil
}

func (x *MetricFilter) GetMetricTypes() []MetricType {
	if x != nil {
		return x.MetricTypes
	}
	return nil
}

func (x *MetricFilter) GetLabelFilters() []*LabelFilter {
	if x != nil {
		return x.LabelFilters
	}
	return nil
}

func (x *MetricFilter) GetNamespaces() []string {
	if x != nil {
		return x.Namespaces
	}
	return nil
}

func (x *MetricFilter) GetSources() []string {
	if x != nil {
		return x.Sources
	}
	return nil
}

func (x *MetricFilter) GetTimeFilter() *TimeFilter {
	if x != nil {
		return x.TimeFilter
	}
	return nil
}

func (x *MetricFilter) GetValueFilters() []*ValueFilter {
	if x != nil {
		return x.ValueFilters
	}
	return nil
}

func (x *MetricFilter) GetLimit() int32 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

func (x *MetricFilter) GetOffset() int32 {
	if x != nil && x.Offset != nil {
		return *x.Offset
	}
	return 0
}

func (x *MetricFilter) SetMetricNames(v []string) {
	x.MetricNames = v
}

func (x *MetricFilter) SetMetricTypes(v []MetricType) {
	x.MetricTypes = v
}

func (x *MetricFilter) SetLabelFilters(v []*LabelFilter) {
	x.LabelFilters = v
}

func (x *MetricFilter) SetNamespaces(v []string) {
	x.Namespaces = v
}

func (x *MetricFilter) SetSources(v []string) {
	x.Sources = v
}

func (x *MetricFilter) SetTimeFilter(v *TimeFilter) {
	x.TimeFilter = v
}

func (x *MetricFilter) SetValueFilters(v []*ValueFilter) {
	x.ValueFilters = v
}

func (x *MetricFilter) SetLimit(v int32) {
	x.Limit = &v
}

func (x *MetricFilter) SetOffset(v int32) {
	x.Offset = &v
}

func (x *MetricFilter) HasTimeFilter() bool {
	if x == nil {
		return false
	}
	return x.TimeFilter != nil
}

func (x *MetricFilter) HasLimit() bool {
	if x == nil {
		return false
	}
	return x.Limit != nil
}

func (x *MetricFilter) HasOffset() bool {
	if x == nil {
		return false
	}
	return x.Offset != nil
}

func (x *MetricFilter) ClearTimeFilter() {
	x.TimeFilter = nil
}

func (x *MetricFilter) ClearLimit() {
	x.Limit = nil
}

func (x *MetricFilter) ClearOffset() {
	x.Offset = nil
}

type MetricFilter_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// Filter by metric name (supports wildcards and regex)
	MetricNames []string
	// Filter by metric type
	MetricTypes []MetricType
	// Filter by labels/tags
	LabelFilters []*LabelFilter
	// Filter by namespace
	Namespaces []string
	// Filter by source system
	Sources []string
	// Time range filter
	TimeFilter *TimeFilter
	// Filter by value range
	ValueFilters []*ValueFilter
	// Limit number of results
	Limit *int32
	// Skip/offset for pagination
	Offset *int32
}

func (b0 MetricFilter_builder) Build() *MetricFilter {
	m0 := &MetricFilter{}
	b, x := &b0, m0
	_, _ = b, x
	x.MetricNames = b.MetricNames
	x.MetricTypes = b.MetricTypes
	x.LabelFilters = b.LabelFilters
	x.Namespaces = b.Namespaces
	x.Sources = b.Sources
	x.TimeFilter = b.TimeFilter
	x.ValueFilters = b.ValueFilters
	x.Limit = b.Limit
	x.Offset = b.Offset
	return m0
}

// *
// LabelFilter filters metrics based on label values.
type LabelFilter struct {
	state protoimpl.MessageState `protogen:"hybrid.v1"`
	// Label key to filter on
	Key *string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	// Comparison operator for the filter
	Operator *ComparisonOperator `protobuf:"varint,2,opt,name=operator,enum=gcommon.v1.metrics.ComparisonOperator" json:"operator,omitempty"`
	// Value(s) to compare against
	Values []string `protobuf:"bytes,3,rep,name=values" json:"values,omitempty"`
	// Whether this filter should be negated (NOT condition)
	Negated       *bool `protobuf:"varint,4,opt,name=negated" json:"negated,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LabelFilter) Reset() {
	*x = LabelFilter{}
	mi := &file_pkg_metrics_proto_messages_metric_filter_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LabelFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabelFilter) ProtoMessage() {}

func (x *LabelFilter) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_metrics_proto_messages_metric_filter_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *LabelFilter) GetKey() string {
	if x != nil && x.Key != nil {
		return *x.Key
	}
	return ""
}

func (x *LabelFilter) GetOperator() ComparisonOperator {
	if x != nil && x.Operator != nil {
		return *x.Operator
	}
	return ComparisonOperator_COMPARISON_OPERATOR_UNSPECIFIED
}

func (x *LabelFilter) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

func (x *LabelFilter) GetNegated() bool {
	if x != nil && x.Negated != nil {
		return *x.Negated
	}
	return false
}

func (x *LabelFilter) SetKey(v string) {
	x.Key = &v
}

func (x *LabelFilter) SetOperator(v ComparisonOperator) {
	x.Operator = &v
}

func (x *LabelFilter) SetValues(v []string) {
	x.Values = v
}

func (x *LabelFilter) SetNegated(v bool) {
	x.Negated = &v
}

func (x *LabelFilter) HasKey() bool {
	if x == nil {
		return false
	}
	return x.Key != nil
}

func (x *LabelFilter) HasOperator() bool {
	if x == nil {
		return false
	}
	return x.Operator != nil
}

func (x *LabelFilter) HasNegated() bool {
	if x == nil {
		return false
	}
	return x.Negated != nil
}

func (x *LabelFilter) ClearKey() {
	x.Key = nil
}

func (x *LabelFilter) ClearOperator() {
	x.Operator = nil
}

func (x *LabelFilter) ClearNegated() {
	x.Negated = nil
}

type LabelFilter_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// Label key to filter on
	Key *string
	// Comparison operator for the filter
	Operator *ComparisonOperator
	// Value(s) to compare against
	Values []string
	// Whether this filter should be negated (NOT condition)
	Negated *bool
}

func (b0 LabelFilter_builder) Build() *LabelFilter {
	m0 := &LabelFilter{}
	b, x := &b0, m0
	_, _ = b, x
	x.Key = b.Key
	x.Operator = b.Operator
	x.Values = b.Values
	x.Negated = b.Negated
	return m0
}

// *
// ValueFilter filters metrics based on their actual values.
type ValueFilter struct {
	state protoimpl.MessageState `protogen:"hybrid.v1"`
	// Comparison operator
	Operator *ComparisonOperator `protobuf:"varint,1,opt,name=operator,enum=gcommon.v1.metrics.ComparisonOperator" json:"operator,omitempty"`
	// Value to compare against
	//
	// Types that are valid to be assigned to Threshold:
	//
	//	*ValueFilter_DoubleValue
	//	*ValueFilter_IntValue
	//	*ValueFilter_StringValue
	//	*ValueFilter_BoolValue
	Threshold isValueFilter_Threshold `protobuf_oneof:"threshold"`
	// Whether this filter should be negated
	Negated       *bool `protobuf:"varint,6,opt,name=negated" json:"negated,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ValueFilter) Reset() {
	*x = ValueFilter{}
	mi := &file_pkg_metrics_proto_messages_metric_filter_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ValueFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValueFilter) ProtoMessage() {}

func (x *ValueFilter) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_metrics_proto_messages_metric_filter_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *ValueFilter) GetOperator() ComparisonOperator {
	if x != nil && x.Operator != nil {
		return *x.Operator
	}
	return ComparisonOperator_COMPARISON_OPERATOR_UNSPECIFIED
}

func (x *ValueFilter) GetThreshold() isValueFilter_Threshold {
	if x != nil {
		return x.Threshold
	}
	return nil
}

func (x *ValueFilter) GetDoubleValue() float64 {
	if x != nil {
		if x, ok := x.Threshold.(*ValueFilter_DoubleValue); ok {
			return x.DoubleValue
		}
	}
	return 0
}

func (x *ValueFilter) GetIntValue() int64 {
	if x != nil {
		if x, ok := x.Threshold.(*ValueFilter_IntValue); ok {
			return x.IntValue
		}
	}
	return 0
}

func (x *ValueFilter) GetStringValue() string {
	if x != nil {
		if x, ok := x.Threshold.(*ValueFilter_StringValue); ok {
			return x.StringValue
		}
	}
	return ""
}

func (x *ValueFilter) GetBoolValue() bool {
	if x != nil {
		if x, ok := x.Threshold.(*ValueFilter_BoolValue); ok {
			return x.BoolValue
		}
	}
	return false
}

func (x *ValueFilter) GetNegated() bool {
	if x != nil && x.Negated != nil {
		return *x.Negated
	}
	return false
}

func (x *ValueFilter) SetOperator(v ComparisonOperator) {
	x.Operator = &v
}

func (x *ValueFilter) SetDoubleValue(v float64) {
	x.Threshold = &ValueFilter_DoubleValue{v}
}

func (x *ValueFilter) SetIntValue(v int64) {
	x.Threshold = &ValueFilter_IntValue{v}
}

func (x *ValueFilter) SetStringValue(v string) {
	x.Threshold = &ValueFilter_StringValue{v}
}

func (x *ValueFilter) SetBoolValue(v bool) {
	x.Threshold = &ValueFilter_BoolValue{v}
}

func (x *ValueFilter) SetNegated(v bool) {
	x.Negated = &v
}

func (x *ValueFilter) HasOperator() bool {
	if x == nil {
		return false
	}
	return x.Operator != nil
}

func (x *ValueFilter) HasThreshold() bool {
	if x == nil {
		return false
	}
	return x.Threshold != nil
}

func (x *ValueFilter) HasDoubleValue() bool {
	if x == nil {
		return false
	}
	_, ok := x.Threshold.(*ValueFilter_DoubleValue)
	return ok
}

func (x *ValueFilter) HasIntValue() bool {
	if x == nil {
		return false
	}
	_, ok := x.Threshold.(*ValueFilter_IntValue)
	return ok
}

func (x *ValueFilter) HasStringValue() bool {
	if x == nil {
		return false
	}
	_, ok := x.Threshold.(*ValueFilter_StringValue)
	return ok
}

func (x *ValueFilter) HasBoolValue() bool {
	if x == nil {
		return false
	}
	_, ok := x.Threshold.(*ValueFilter_BoolValue)
	return ok
}

func (x *ValueFilter) HasNegated() bool {
	if x == nil {
		return false
	}
	return x.Negated != nil
}

func (x *ValueFilter) ClearOperator() {
	x.Operator = nil
}

func (x *ValueFilter) ClearThreshold() {
	x.Threshold = nil
}

func (x *ValueFilter) ClearDoubleValue() {
	if _, ok := x.Threshold.(*ValueFilter_DoubleValue); ok {
		x.Threshold = nil
	}
}

func (x *ValueFilter) ClearIntValue() {
	if _, ok := x.Threshold.(*ValueFilter_IntValue); ok {
		x.Threshold = nil
	}
}

func (x *ValueFilter) ClearStringValue() {
	if _, ok := x.Threshold.(*ValueFilter_StringValue); ok {
		x.Threshold = nil
	}
}

func (x *ValueFilter) ClearBoolValue() {
	if _, ok := x.Threshold.(*ValueFilter_BoolValue); ok {
		x.Threshold = nil
	}
}

func (x *ValueFilter) ClearNegated() {
	x.Negated = nil
}

const ValueFilter_Threshold_not_set_case case_ValueFilter_Threshold = 0
const ValueFilter_DoubleValue_case case_ValueFilter_Threshold = 2
const ValueFilter_IntValue_case case_ValueFilter_Threshold = 3
const ValueFilter_StringValue_case case_ValueFilter_Threshold = 4
const ValueFilter_BoolValue_case case_ValueFilter_Threshold = 5

func (x *ValueFilter) WhichThreshold() case_ValueFilter_Threshold {
	if x == nil {
		return ValueFilter_Threshold_not_set_case
	}
	switch x.Threshold.(type) {
	case *ValueFilter_DoubleValue:
		return ValueFilter_DoubleValue_case
	case *ValueFilter_IntValue:
		return ValueFilter_IntValue_case
	case *ValueFilter_StringValue:
		return ValueFilter_StringValue_case
	case *ValueFilter_BoolValue:
		return ValueFilter_BoolValue_case
	default:
		return ValueFilter_Threshold_not_set_case
	}
}

type ValueFilter_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// Comparison operator
	Operator *ComparisonOperator
	// Value to compare against

	// Fields of oneof Threshold:
	DoubleValue *float64
	IntValue    *int64
	StringValue *string
	BoolValue   *bool
	// -- end of Threshold
	// Whether this filter should be negated
	Negated *bool
}

func (b0 ValueFilter_builder) Build() *ValueFilter {
	m0 := &ValueFilter{}
	b, x := &b0, m0
	_, _ = b, x
	x.Operator = b.Operator
	if b.DoubleValue != nil {
		x.Threshold = &ValueFilter_DoubleValue{*b.DoubleValue}
	}
	if b.IntValue != nil {
		x.Threshold = &ValueFilter_IntValue{*b.IntValue}
	}
	if b.StringValue != nil {
		x.Threshold = &ValueFilter_StringValue{*b.StringValue}
	}
	if b.BoolValue != nil {
		x.Threshold = &ValueFilter_BoolValue{*b.BoolValue}
	}
	x.Negated = b.Negated
	return m0
}

type case_ValueFilter_Threshold protoreflect.FieldNumber

func (x case_ValueFilter_Threshold) String() string {
	md := file_pkg_metrics_proto_messages_metric_filter_proto_msgTypes[2].Descriptor()
	if x == 0 {
		return "not set"
	}
	return protoimpl.X.MessageFieldStringOf(md, protoreflect.FieldNumber(x))
}

type isValueFilter_Threshold interface {
	isValueFilter_Threshold()
}

type ValueFilter_DoubleValue struct {
	DoubleValue float64 `protobuf:"fixed64,2,opt,name=double_value,json=doubleValue,oneof"`
}

type ValueFilter_IntValue struct {
	IntValue int64 `protobuf:"varint,3,opt,name=int_value,json=intValue,oneof"`
}

type ValueFilter_StringValue struct {
	StringValue string `protobuf:"bytes,4,opt,name=string_value,json=stringValue,oneof"`
}

type ValueFilter_BoolValue struct {
	BoolValue bool `protobuf:"varint,5,opt,name=bool_value,json=boolValue,oneof"`
}

func (*ValueFilter_DoubleValue) isValueFilter_Threshold() {}

func (*ValueFilter_IntValue) isValueFilter_Threshold() {}

func (*ValueFilter_StringValue) isValueFilter_Threshold() {}

func (*ValueFilter_BoolValue) isValueFilter_Threshold() {}

// *
// SortCriteria defines how to sort metric results.
type SortCriteria struct {
	state         protoimpl.MessageState  `protogen:"hybrid.v1"`
	Field         *SortCriteria_SortField `protobuf:"varint,1,opt,name=field,enum=gcommon.v1.metrics.SortCriteria_SortField" json:"field,omitempty"`
	Order         *SortCriteria_SortOrder `protobuf:"varint,2,opt,name=order,enum=gcommon.v1.metrics.SortCriteria_SortOrder" json:"order,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SortCriteria) Reset() {
	*x = SortCriteria{}
	mi := &file_pkg_metrics_proto_messages_metric_filter_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SortCriteria) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SortCriteria) ProtoMessage() {}

func (x *SortCriteria) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_metrics_proto_messages_metric_filter_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *SortCriteria) GetField() SortCriteria_SortField {
	if x != nil && x.Field != nil {
		return *x.Field
	}
	return SortCriteria_SORT_FIELD_UNSPECIFIED
}

func (x *SortCriteria) GetOrder() SortCriteria_SortOrder {
	if x != nil && x.Order != nil {
		return *x.Order
	}
	return SortCriteria_SORT_ORDER_UNSPECIFIED
}

func (x *SortCriteria) SetField(v SortCriteria_SortField) {
	x.Field = &v
}

func (x *SortCriteria) SetOrder(v SortCriteria_SortOrder) {
	x.Order = &v
}

func (x *SortCriteria) HasField() bool {
	if x == nil {
		return false
	}
	return x.Field != nil
}

func (x *SortCriteria) HasOrder() bool {
	if x == nil {
		return false
	}
	return x.Order != nil
}

func (x *SortCriteria) ClearField() {
	x.Field = nil
}

func (x *SortCriteria) ClearOrder() {
	x.Order = nil
}

type SortCriteria_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Field *SortCriteria_SortField
	Order *SortCriteria_SortOrder
}

func (b0 SortCriteria_builder) Build() *SortCriteria {
	m0 := &SortCriteria{}
	b, x := &b0, m0
	_, _ = b, x
	x.Field = b.Field
	x.Order = b.Order
	return m0
}

var File_pkg_metrics_proto_messages_metric_filter_proto protoreflect.FileDescriptor

const file_pkg_metrics_proto_messages_metric_filter_proto_rawDesc = "" +
	"\n" +
	".pkg/metrics/proto/messages/metric_filter.proto\x12\x12gcommon.v1.metrics\x1a!google/protobuf/go_features.proto\x1a)pkg/metrics/proto/enums/metric_type.proto\x1a1pkg/metrics/proto/enums/comparison_operator.proto\x1a-pkg/metrics/proto/types/timestamp_range.proto\"\xa9\x03\n" +
	"\fMetricFilter\x12!\n" +
	"\fmetric_names\x18\x01 \x03(\tR\vmetricNames\x12A\n" +
	"\fmetric_types\x18\x02 \x03(\x0e2\x1e.gcommon.v1.metrics.MetricTypeR\vmetricTypes\x12D\n" +
	"\rlabel_filters\x18\x03 \x03(\v2\x1f.gcommon.v1.metrics.LabelFilterR\flabelFilters\x12\x1e\n" +
	"\n" +
	"namespaces\x18\x04 \x03(\tR\n" +
	"namespaces\x12\x18\n" +
	"\asources\x18\x05 \x03(\tR\asources\x12?\n" +
	"\vtime_filter\x18\x06 \x01(\v2\x1e.gcommon.v1.metrics.TimeFilterR\n" +
	"timeFilter\x12D\n" +
	"\rvalue_filters\x18\a \x03(\v2\x1f.gcommon.v1.metrics.ValueFilterR\fvalueFilters\x12\x14\n" +
	"\x05limit\x18\b \x01(\x05R\x05limit\x12\x16\n" +
	"\x06offset\x18\t \x01(\x05R\x06offset\"\x95\x01\n" +
	"\vLabelFilter\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12B\n" +
	"\boperator\x18\x02 \x01(\x0e2&.gcommon.v1.metrics.ComparisonOperatorR\boperator\x12\x16\n" +
	"\x06values\x18\x03 \x03(\tR\x06values\x12\x18\n" +
	"\anegated\x18\x04 \x01(\bR\anegated\"\x82\x02\n" +
	"\vValueFilter\x12B\n" +
	"\boperator\x18\x01 \x01(\x0e2&.gcommon.v1.metrics.ComparisonOperatorR\boperator\x12#\n" +
	"\fdouble_value\x18\x02 \x01(\x01H\x00R\vdoubleValue\x12\x1d\n" +
	"\tint_value\x18\x03 \x01(\x03H\x00R\bintValue\x12#\n" +
	"\fstring_value\x18\x04 \x01(\tH\x00R\vstringValue\x12\x1f\n" +
	"\n" +
	"bool_value\x18\x05 \x01(\bH\x00R\tboolValue\x12\x18\n" +
	"\anegated\x18\x06 \x01(\bR\anegatedB\v\n" +
	"\tthreshold\"\xfa\x02\n" +
	"\fSortCriteria\x12@\n" +
	"\x05field\x18\x01 \x01(\x0e2*.gcommon.v1.metrics.SortCriteria.SortFieldR\x05field\x12@\n" +
	"\x05order\x18\x02 \x01(\x0e2*.gcommon.v1.metrics.SortCriteria.SortOrderR\x05order\"\x87\x01\n" +
	"\tSortField\x12\x1a\n" +
	"\x16SORT_FIELD_UNSPECIFIED\x10\x00\x12\x13\n" +
	"\x0fSORT_FIELD_NAME\x10\x01\x12\x18\n" +
	"\x14SORT_FIELD_TIMESTAMP\x10\x02\x12\x14\n" +
	"\x10SORT_FIELD_VALUE\x10\x03\x12\x19\n" +
	"\x15SORT_FIELD_CREATED_AT\x10\x04\"\\\n" +
	"\tSortOrder\x12\x1a\n" +
	"\x16SORT_ORDER_UNSPECIFIED\x10\x00\x12\x18\n" +
	"\x14SORT_ORDER_ASCENDING\x10\x01\x12\x19\n" +
	"\x15SORT_ORDER_DESCENDING\x10\x02B\xd4\x01\n" +
	"\x16com.gcommon.v1.metricsB\x11MetricFilterProtoP\x01Z5github.com/jdfalk/gcommon/pkg/metrics/proto;metricspb\xa2\x02\x03GVM\xaa\x02\x12Gcommon.V1.Metrics\xca\x02\x12Gcommon\\V1\\Metrics\xe2\x02\x1eGcommon\\V1\\Metrics\\GPBMetadata\xea\x02\x14Gcommon::V1::Metrics\x92\x03\x05\xd2>\x02\x10\x02b\beditionsp\xe8\a"

var file_pkg_metrics_proto_messages_metric_filter_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_pkg_metrics_proto_messages_metric_filter_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pkg_metrics_proto_messages_metric_filter_proto_goTypes = []any{
	(SortCriteria_SortField)(0), // 0: gcommon.v1.metrics.SortCriteria.SortField
	(SortCriteria_SortOrder)(0), // 1: gcommon.v1.metrics.SortCriteria.SortOrder
	(*MetricFilter)(nil),        // 2: gcommon.v1.metrics.MetricFilter
	(*LabelFilter)(nil),         // 3: gcommon.v1.metrics.LabelFilter
	(*ValueFilter)(nil),         // 4: gcommon.v1.metrics.ValueFilter
	(*SortCriteria)(nil),        // 5: gcommon.v1.metrics.SortCriteria
	(MetricType)(0),             // 6: gcommon.v1.metrics.MetricType
	(*TimeFilter)(nil),          // 7: gcommon.v1.metrics.TimeFilter
	(ComparisonOperator)(0),     // 8: gcommon.v1.metrics.ComparisonOperator
}
var file_pkg_metrics_proto_messages_metric_filter_proto_depIdxs = []int32{
	6, // 0: gcommon.v1.metrics.MetricFilter.metric_types:type_name -> gcommon.v1.metrics.MetricType
	3, // 1: gcommon.v1.metrics.MetricFilter.label_filters:type_name -> gcommon.v1.metrics.LabelFilter
	7, // 2: gcommon.v1.metrics.MetricFilter.time_filter:type_name -> gcommon.v1.metrics.TimeFilter
	4, // 3: gcommon.v1.metrics.MetricFilter.value_filters:type_name -> gcommon.v1.metrics.ValueFilter
	8, // 4: gcommon.v1.metrics.LabelFilter.operator:type_name -> gcommon.v1.metrics.ComparisonOperator
	8, // 5: gcommon.v1.metrics.ValueFilter.operator:type_name -> gcommon.v1.metrics.ComparisonOperator
	0, // 6: gcommon.v1.metrics.SortCriteria.field:type_name -> gcommon.v1.metrics.SortCriteria.SortField
	1, // 7: gcommon.v1.metrics.SortCriteria.order:type_name -> gcommon.v1.metrics.SortCriteria.SortOrder
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_pkg_metrics_proto_messages_metric_filter_proto_init() }
func file_pkg_metrics_proto_messages_metric_filter_proto_init() {
	if File_pkg_metrics_proto_messages_metric_filter_proto != nil {
		return
	}
	file_pkg_metrics_proto_enums_metric_type_proto_init()
	file_pkg_metrics_proto_enums_comparison_operator_proto_init()
	file_pkg_metrics_proto_types_timestamp_range_proto_init()
	file_pkg_metrics_proto_messages_metric_filter_proto_msgTypes[2].OneofWrappers = []any{
		(*ValueFilter_DoubleValue)(nil),
		(*ValueFilter_IntValue)(nil),
		(*ValueFilter_StringValue)(nil),
		(*ValueFilter_BoolValue)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pkg_metrics_proto_messages_metric_filter_proto_rawDesc), len(file_pkg_metrics_proto_messages_metric_filter_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_metrics_proto_messages_metric_filter_proto_goTypes,
		DependencyIndexes: file_pkg_metrics_proto_messages_metric_filter_proto_depIdxs,
		EnumInfos:         file_pkg_metrics_proto_messages_metric_filter_proto_enumTypes,
		MessageInfos:      file_pkg_metrics_proto_messages_metric_filter_proto_msgTypes,
	}.Build()
	File_pkg_metrics_proto_messages_metric_filter_proto = out.File
	file_pkg_metrics_proto_messages_metric_filter_proto_goTypes = nil
	file_pkg_metrics_proto_messages_metric_filter_proto_depIdxs = nil
}

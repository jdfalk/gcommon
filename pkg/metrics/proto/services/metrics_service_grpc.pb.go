// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: pkg/metrics/proto/services/metrics_service.proto

package metricspb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	MetricsService_RecordMetric_FullMethodName       = "/gcommon.v1.metrics.MetricsService/RecordMetric"
	MetricsService_RecordBatchMetrics_FullMethodName = "/gcommon.v1.metrics.MetricsService/RecordBatchMetrics"
	MetricsService_GetMetrics_FullMethodName         = "/gcommon.v1.metrics.MetricsService/GetMetrics"
	MetricsService_StreamMetrics_FullMethodName      = "/gcommon.v1.metrics.MetricsService/StreamMetrics"
	MetricsService_RegisterMetric_FullMethodName     = "/gcommon.v1.metrics.MetricsService/RegisterMetric"
	MetricsService_UnregisterMetric_FullMethodName   = "/gcommon.v1.metrics.MetricsService/UnregisterMetric"
	MetricsService_GetMetricMetadata_FullMethodName  = "/gcommon.v1.metrics.MetricsService/GetMetricMetadata"
	MetricsService_QueryMetrics_FullMethodName       = "/gcommon.v1.metrics.MetricsService/QueryMetrics"
	MetricsService_GetMetricsSummary_FullMethodName  = "/gcommon.v1.metrics.MetricsService/GetMetricsSummary"
)

// MetricsServiceClient is the client API for MetricsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// *
// MetricsService provides comprehensive metrics collection and querying capabilities.
// Supports counters, gauges, histograms, streaming, and custom metrics aggregation.
type MetricsServiceClient interface {
	// Record a single metric data point
	RecordMetric(ctx context.Context, in *RecordMetricRequest, opts ...grpc.CallOption) (*RecordMetricResponse, error)
	// Record multiple metrics in batch for efficiency
	RecordBatchMetrics(ctx context.Context, in *RecordMetricsRequest, opts ...grpc.CallOption) (*RecordMetricsResponse, error)
	// Retrieve metrics data with filtering and aggregation
	GetMetrics(ctx context.Context, in *GetMetricsRequest, opts ...grpc.CallOption) (*GetMetricsResponse, error)
	// Stream metrics data in real-time
	StreamMetrics(ctx context.Context, in *StreamMetricsRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[MetricData], error)
	// Register a new metric definition
	RegisterMetric(ctx context.Context, in *RegisterMetricRequest, opts ...grpc.CallOption) (*RegisterMetricResponse, error)
	// Unregister an existing metric
	UnregisterMetric(ctx context.Context, in *UnregisterMetricRequest, opts ...grpc.CallOption) (*UnregisterMetricResponse, error)
	// Get metadata for a specific metric
	GetMetricMetadata(ctx context.Context, in *GetMetricMetadataRequest, opts ...grpc.CallOption) (*GetMetricMetadataResponse, error)
	// Query metrics data using complex query expressions
	QueryMetrics(ctx context.Context, in *QueryMetricsRequest, opts ...grpc.CallOption) (*QueryMetricsResponse, error)
	// Get summary statistics about metrics
	GetMetricsSummary(ctx context.Context, in *GetMetricsSummaryRequest, opts ...grpc.CallOption) (*GetMetricsSummaryResponse, error)
}

type metricsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMetricsServiceClient(cc grpc.ClientConnInterface) MetricsServiceClient {
	return &metricsServiceClient{cc}
}

func (c *metricsServiceClient) RecordMetric(ctx context.Context, in *RecordMetricRequest, opts ...grpc.CallOption) (*RecordMetricResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RecordMetricResponse)
	err := c.cc.Invoke(ctx, MetricsService_RecordMetric_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsServiceClient) RecordBatchMetrics(ctx context.Context, in *RecordMetricsRequest, opts ...grpc.CallOption) (*RecordMetricsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RecordMetricsResponse)
	err := c.cc.Invoke(ctx, MetricsService_RecordBatchMetrics_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsServiceClient) GetMetrics(ctx context.Context, in *GetMetricsRequest, opts ...grpc.CallOption) (*GetMetricsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMetricsResponse)
	err := c.cc.Invoke(ctx, MetricsService_GetMetrics_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsServiceClient) StreamMetrics(ctx context.Context, in *StreamMetricsRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[MetricData], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &MetricsService_ServiceDesc.Streams[0], MetricsService_StreamMetrics_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[StreamMetricsRequest, MetricData]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type MetricsService_StreamMetricsClient = grpc.ServerStreamingClient[MetricData]

func (c *metricsServiceClient) RegisterMetric(ctx context.Context, in *RegisterMetricRequest, opts ...grpc.CallOption) (*RegisterMetricResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterMetricResponse)
	err := c.cc.Invoke(ctx, MetricsService_RegisterMetric_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsServiceClient) UnregisterMetric(ctx context.Context, in *UnregisterMetricRequest, opts ...grpc.CallOption) (*UnregisterMetricResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UnregisterMetricResponse)
	err := c.cc.Invoke(ctx, MetricsService_UnregisterMetric_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsServiceClient) GetMetricMetadata(ctx context.Context, in *GetMetricMetadataRequest, opts ...grpc.CallOption) (*GetMetricMetadataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMetricMetadataResponse)
	err := c.cc.Invoke(ctx, MetricsService_GetMetricMetadata_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsServiceClient) QueryMetrics(ctx context.Context, in *QueryMetricsRequest, opts ...grpc.CallOption) (*QueryMetricsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueryMetricsResponse)
	err := c.cc.Invoke(ctx, MetricsService_QueryMetrics_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsServiceClient) GetMetricsSummary(ctx context.Context, in *GetMetricsSummaryRequest, opts ...grpc.CallOption) (*GetMetricsSummaryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMetricsSummaryResponse)
	err := c.cc.Invoke(ctx, MetricsService_GetMetricsSummary_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetricsServiceServer is the server API for MetricsService service.
// All implementations should embed UnimplementedMetricsServiceServer
// for forward compatibility.
//
// *
// MetricsService provides comprehensive metrics collection and querying capabilities.
// Supports counters, gauges, histograms, streaming, and custom metrics aggregation.
type MetricsServiceServer interface {
	// Record a single metric data point
	RecordMetric(context.Context, *RecordMetricRequest) (*RecordMetricResponse, error)
	// Record multiple metrics in batch for efficiency
	RecordBatchMetrics(context.Context, *RecordMetricsRequest) (*RecordMetricsResponse, error)
	// Retrieve metrics data with filtering and aggregation
	GetMetrics(context.Context, *GetMetricsRequest) (*GetMetricsResponse, error)
	// Stream metrics data in real-time
	StreamMetrics(*StreamMetricsRequest, grpc.ServerStreamingServer[MetricData]) error
	// Register a new metric definition
	RegisterMetric(context.Context, *RegisterMetricRequest) (*RegisterMetricResponse, error)
	// Unregister an existing metric
	UnregisterMetric(context.Context, *UnregisterMetricRequest) (*UnregisterMetricResponse, error)
	// Get metadata for a specific metric
	GetMetricMetadata(context.Context, *GetMetricMetadataRequest) (*GetMetricMetadataResponse, error)
	// Query metrics data using complex query expressions
	QueryMetrics(context.Context, *QueryMetricsRequest) (*QueryMetricsResponse, error)
	// Get summary statistics about metrics
	GetMetricsSummary(context.Context, *GetMetricsSummaryRequest) (*GetMetricsSummaryResponse, error)
}

// UnimplementedMetricsServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMetricsServiceServer struct{}

func (UnimplementedMetricsServiceServer) RecordMetric(context.Context, *RecordMetricRequest) (*RecordMetricResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecordMetric not implemented")
}
func (UnimplementedMetricsServiceServer) RecordBatchMetrics(context.Context, *RecordMetricsRequest) (*RecordMetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecordBatchMetrics not implemented")
}
func (UnimplementedMetricsServiceServer) GetMetrics(context.Context, *GetMetricsRequest) (*GetMetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetrics not implemented")
}
func (UnimplementedMetricsServiceServer) StreamMetrics(*StreamMetricsRequest, grpc.ServerStreamingServer[MetricData]) error {
	return status.Errorf(codes.Unimplemented, "method StreamMetrics not implemented")
}
func (UnimplementedMetricsServiceServer) RegisterMetric(context.Context, *RegisterMetricRequest) (*RegisterMetricResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterMetric not implemented")
}
func (UnimplementedMetricsServiceServer) UnregisterMetric(context.Context, *UnregisterMetricRequest) (*UnregisterMetricResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnregisterMetric not implemented")
}
func (UnimplementedMetricsServiceServer) GetMetricMetadata(context.Context, *GetMetricMetadataRequest) (*GetMetricMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetricMetadata not implemented")
}
func (UnimplementedMetricsServiceServer) QueryMetrics(context.Context, *QueryMetricsRequest) (*QueryMetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryMetrics not implemented")
}
func (UnimplementedMetricsServiceServer) GetMetricsSummary(context.Context, *GetMetricsSummaryRequest) (*GetMetricsSummaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetricsSummary not implemented")
}
func (UnimplementedMetricsServiceServer) testEmbeddedByValue() {}

// UnsafeMetricsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MetricsServiceServer will
// result in compilation errors.
type UnsafeMetricsServiceServer interface {
	mustEmbedUnimplementedMetricsServiceServer()
}

func RegisterMetricsServiceServer(s grpc.ServiceRegistrar, srv MetricsServiceServer) {
	// If the following call pancis, it indicates UnimplementedMetricsServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MetricsService_ServiceDesc, srv)
}

func _MetricsService_RecordMetric_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecordMetricRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServiceServer).RecordMetric(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MetricsService_RecordMetric_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServiceServer).RecordMetric(ctx, req.(*RecordMetricRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsService_RecordBatchMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecordMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServiceServer).RecordBatchMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MetricsService_RecordBatchMetrics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServiceServer).RecordBatchMetrics(ctx, req.(*RecordMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsService_GetMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServiceServer).GetMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MetricsService_GetMetrics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServiceServer).GetMetrics(ctx, req.(*GetMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsService_StreamMetrics_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamMetricsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MetricsServiceServer).StreamMetrics(m, &grpc.GenericServerStream[StreamMetricsRequest, MetricData]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type MetricsService_StreamMetricsServer = grpc.ServerStreamingServer[MetricData]

func _MetricsService_RegisterMetric_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterMetricRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServiceServer).RegisterMetric(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MetricsService_RegisterMetric_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServiceServer).RegisterMetric(ctx, req.(*RegisterMetricRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsService_UnregisterMetric_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnregisterMetricRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServiceServer).UnregisterMetric(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MetricsService_UnregisterMetric_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServiceServer).UnregisterMetric(ctx, req.(*UnregisterMetricRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsService_GetMetricMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMetricMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServiceServer).GetMetricMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MetricsService_GetMetricMetadata_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServiceServer).GetMetricMetadata(ctx, req.(*GetMetricMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsService_QueryMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServiceServer).QueryMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MetricsService_QueryMetrics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServiceServer).QueryMetrics(ctx, req.(*QueryMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsService_GetMetricsSummary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMetricsSummaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServiceServer).GetMetricsSummary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MetricsService_GetMetricsSummary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServiceServer).GetMetricsSummary(ctx, req.(*GetMetricsSummaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MetricsService_ServiceDesc is the grpc.ServiceDesc for MetricsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MetricsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gcommon.v1.metrics.MetricsService",
	HandlerType: (*MetricsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RecordMetric",
			Handler:    _MetricsService_RecordMetric_Handler,
		},
		{
			MethodName: "RecordBatchMetrics",
			Handler:    _MetricsService_RecordBatchMetrics_Handler,
		},
		{
			MethodName: "GetMetrics",
			Handler:    _MetricsService_GetMetrics_Handler,
		},
		{
			MethodName: "RegisterMetric",
			Handler:    _MetricsService_RegisterMetric_Handler,
		},
		{
			MethodName: "UnregisterMetric",
			Handler:    _MetricsService_UnregisterMetric_Handler,
		},
		{
			MethodName: "GetMetricMetadata",
			Handler:    _MetricsService_GetMetricMetadata_Handler,
		},
		{
			MethodName: "QueryMetrics",
			Handler:    _MetricsService_QueryMetrics_Handler,
		},
		{
			MethodName: "GetMetricsSummary",
			Handler:    _MetricsService_GetMetricsSummary_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamMetrics",
			Handler:       _MetricsService_StreamMetrics_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pkg/metrics/proto/services/metrics_service.proto",
}

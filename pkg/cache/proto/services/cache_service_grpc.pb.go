// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: pkg/cache/proto/services/cache_service.proto

package services

import (
	context "context"
	requests "github.com/jdfalk/gcommon/pkg/cache/proto/requests"
	responses "github.com/jdfalk/gcommon/pkg/cache/proto/responses"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	CacheService_Get_FullMethodName             = "/gcommon.v1.cache.CacheService/Get"
	CacheService_Set_FullMethodName             = "/gcommon.v1.cache.CacheService/Set"
	CacheService_Delete_FullMethodName          = "/gcommon.v1.cache.CacheService/Delete"
	CacheService_Exists_FullMethodName          = "/gcommon.v1.cache.CacheService/Exists"
	CacheService_GetMultiple_FullMethodName     = "/gcommon.v1.cache.CacheService/GetMultiple"
	CacheService_SetMultiple_FullMethodName     = "/gcommon.v1.cache.CacheService/SetMultiple"
	CacheService_DeleteMultiple_FullMethodName  = "/gcommon.v1.cache.CacheService/DeleteMultiple"
	CacheService_Increment_FullMethodName       = "/gcommon.v1.cache.CacheService/Increment"
	CacheService_Decrement_FullMethodName       = "/gcommon.v1.cache.CacheService/Decrement"
	CacheService_Clear_FullMethodName           = "/gcommon.v1.cache.CacheService/Clear"
	CacheService_Keys_FullMethodName            = "/gcommon.v1.cache.CacheService/Keys"
	CacheService_GetStats_FullMethodName        = "/gcommon.v1.cache.CacheService/GetStats"
	CacheService_Flush_FullMethodName           = "/gcommon.v1.cache.CacheService/Flush"
	CacheService_TouchExpiration_FullMethodName = "/gcommon.v1.cache.CacheService/TouchExpiration"
)

// CacheServiceClient is the client API for CacheService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// *
// CacheService provides comprehensive caching capabilities.
// Supports CRUD operations, batch operations, atomic operations,
// and cache management with flexible expiration policies.
type CacheServiceClient interface {
	// Get retrieves a value from the cache by key
	Get(ctx context.Context, in *requests.GetRequest, opts ...grpc.CallOption) (*responses.GetResponse, error)
	// Set stores a value in the cache with optional expiration
	Set(ctx context.Context, in *requests.SetRequest, opts ...grpc.CallOption) (*responses.SetResponse, error)
	// Delete removes a value from the cache
	Delete(ctx context.Context, in *requests.DeleteRequest, opts ...grpc.CallOption) (*responses.DeleteResponse, error)
	// Exists checks if a key exists in the cache
	Exists(ctx context.Context, in *requests.ExistsRequest, opts ...grpc.CallOption) (*responses.ExistsResponse, error)
	// GetMultiple retrieves multiple values from the cache in a single operation
	GetMultiple(ctx context.Context, in *requests.GetMultipleRequest, opts ...grpc.CallOption) (*responses.GetMultipleResponse, error)
	// SetMultiple stores multiple values in the cache atomically
	SetMultiple(ctx context.Context, in *requests.SetMultipleRequest, opts ...grpc.CallOption) (*responses.SetMultipleResponse, error)
	// DeleteMultiple removes multiple values from the cache atomically
	DeleteMultiple(ctx context.Context, in *requests.DeleteMultipleRequest, opts ...grpc.CallOption) (*responses.DeleteMultipleResponse, error)
	// Increment atomically increments a numeric value
	Increment(ctx context.Context, in *requests.IncrementRequest, opts ...grpc.CallOption) (*responses.IncrementResponse, error)
	// Decrement atomically decrements a numeric value
	Decrement(ctx context.Context, in *requests.DecrementRequest, opts ...grpc.CallOption) (*responses.DecrementResponse, error)
	// Clear removes all entries from the cache or by pattern
	Clear(ctx context.Context, in *requests.ClearRequest, opts ...grpc.CallOption) (*responses.ClearResponse, error)
	// Keys returns all keys matching a pattern
	Keys(ctx context.Context, in *requests.KeysRequest, opts ...grpc.CallOption) (*responses.KeysResponse, error)
	// GetStats returns cache statistics and performance metrics
	GetStats(ctx context.Context, in *requests.GetStatsRequest, opts ...grpc.CallOption) (*responses.GetStatsResponse, error)
	// Flush forces cache persistence if supported by the backend
	Flush(ctx context.Context, in *requests.FlushRequest, opts ...grpc.CallOption) (*responses.FlushResponse, error)
	// TouchExpiration updates the expiration time of an existing key
	TouchExpiration(ctx context.Context, in *requests.TouchExpirationRequest, opts ...grpc.CallOption) (*responses.TouchExpirationResponse, error)
}

type cacheServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCacheServiceClient(cc grpc.ClientConnInterface) CacheServiceClient {
	return &cacheServiceClient{cc}
}

func (c *cacheServiceClient) Get(ctx context.Context, in *requests.GetRequest, opts ...grpc.CallOption) (*responses.GetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(responses.GetResponse)
	err := c.cc.Invoke(ctx, CacheService_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) Set(ctx context.Context, in *requests.SetRequest, opts ...grpc.CallOption) (*responses.SetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(responses.SetResponse)
	err := c.cc.Invoke(ctx, CacheService_Set_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) Delete(ctx context.Context, in *requests.DeleteRequest, opts ...grpc.CallOption) (*responses.DeleteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(responses.DeleteResponse)
	err := c.cc.Invoke(ctx, CacheService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) Exists(ctx context.Context, in *requests.ExistsRequest, opts ...grpc.CallOption) (*responses.ExistsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(responses.ExistsResponse)
	err := c.cc.Invoke(ctx, CacheService_Exists_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) GetMultiple(ctx context.Context, in *requests.GetMultipleRequest, opts ...grpc.CallOption) (*responses.GetMultipleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(responses.GetMultipleResponse)
	err := c.cc.Invoke(ctx, CacheService_GetMultiple_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) SetMultiple(ctx context.Context, in *requests.SetMultipleRequest, opts ...grpc.CallOption) (*responses.SetMultipleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(responses.SetMultipleResponse)
	err := c.cc.Invoke(ctx, CacheService_SetMultiple_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) DeleteMultiple(ctx context.Context, in *requests.DeleteMultipleRequest, opts ...grpc.CallOption) (*responses.DeleteMultipleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(responses.DeleteMultipleResponse)
	err := c.cc.Invoke(ctx, CacheService_DeleteMultiple_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) Increment(ctx context.Context, in *requests.IncrementRequest, opts ...grpc.CallOption) (*responses.IncrementResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(responses.IncrementResponse)
	err := c.cc.Invoke(ctx, CacheService_Increment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) Decrement(ctx context.Context, in *requests.DecrementRequest, opts ...grpc.CallOption) (*responses.DecrementResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(responses.DecrementResponse)
	err := c.cc.Invoke(ctx, CacheService_Decrement_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) Clear(ctx context.Context, in *requests.ClearRequest, opts ...grpc.CallOption) (*responses.ClearResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(responses.ClearResponse)
	err := c.cc.Invoke(ctx, CacheService_Clear_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) Keys(ctx context.Context, in *requests.KeysRequest, opts ...grpc.CallOption) (*responses.KeysResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(responses.KeysResponse)
	err := c.cc.Invoke(ctx, CacheService_Keys_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) GetStats(ctx context.Context, in *requests.GetStatsRequest, opts ...grpc.CallOption) (*responses.GetStatsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(responses.GetStatsResponse)
	err := c.cc.Invoke(ctx, CacheService_GetStats_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) Flush(ctx context.Context, in *requests.FlushRequest, opts ...grpc.CallOption) (*responses.FlushResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(responses.FlushResponse)
	err := c.cc.Invoke(ctx, CacheService_Flush_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) TouchExpiration(ctx context.Context, in *requests.TouchExpirationRequest, opts ...grpc.CallOption) (*responses.TouchExpirationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(responses.TouchExpirationResponse)
	err := c.cc.Invoke(ctx, CacheService_TouchExpiration_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CacheServiceServer is the server API for CacheService service.
// All implementations should embed UnimplementedCacheServiceServer
// for forward compatibility.
//
// *
// CacheService provides comprehensive caching capabilities.
// Supports CRUD operations, batch operations, atomic operations,
// and cache management with flexible expiration policies.
type CacheServiceServer interface {
	// Get retrieves a value from the cache by key
	Get(context.Context, *requests.GetRequest) (*responses.GetResponse, error)
	// Set stores a value in the cache with optional expiration
	Set(context.Context, *requests.SetRequest) (*responses.SetResponse, error)
	// Delete removes a value from the cache
	Delete(context.Context, *requests.DeleteRequest) (*responses.DeleteResponse, error)
	// Exists checks if a key exists in the cache
	Exists(context.Context, *requests.ExistsRequest) (*responses.ExistsResponse, error)
	// GetMultiple retrieves multiple values from the cache in a single operation
	GetMultiple(context.Context, *requests.GetMultipleRequest) (*responses.GetMultipleResponse, error)
	// SetMultiple stores multiple values in the cache atomically
	SetMultiple(context.Context, *requests.SetMultipleRequest) (*responses.SetMultipleResponse, error)
	// DeleteMultiple removes multiple values from the cache atomically
	DeleteMultiple(context.Context, *requests.DeleteMultipleRequest) (*responses.DeleteMultipleResponse, error)
	// Increment atomically increments a numeric value
	Increment(context.Context, *requests.IncrementRequest) (*responses.IncrementResponse, error)
	// Decrement atomically decrements a numeric value
	Decrement(context.Context, *requests.DecrementRequest) (*responses.DecrementResponse, error)
	// Clear removes all entries from the cache or by pattern
	Clear(context.Context, *requests.ClearRequest) (*responses.ClearResponse, error)
	// Keys returns all keys matching a pattern
	Keys(context.Context, *requests.KeysRequest) (*responses.KeysResponse, error)
	// GetStats returns cache statistics and performance metrics
	GetStats(context.Context, *requests.GetStatsRequest) (*responses.GetStatsResponse, error)
	// Flush forces cache persistence if supported by the backend
	Flush(context.Context, *requests.FlushRequest) (*responses.FlushResponse, error)
	// TouchExpiration updates the expiration time of an existing key
	TouchExpiration(context.Context, *requests.TouchExpirationRequest) (*responses.TouchExpirationResponse, error)
}

// UnimplementedCacheServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCacheServiceServer struct{}

func (UnimplementedCacheServiceServer) Get(context.Context, *requests.GetRequest) (*responses.GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedCacheServiceServer) Set(context.Context, *requests.SetRequest) (*responses.SetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (UnimplementedCacheServiceServer) Delete(context.Context, *requests.DeleteRequest) (*responses.DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedCacheServiceServer) Exists(context.Context, *requests.ExistsRequest) (*responses.ExistsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Exists not implemented")
}
func (UnimplementedCacheServiceServer) GetMultiple(context.Context, *requests.GetMultipleRequest) (*responses.GetMultipleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMultiple not implemented")
}
func (UnimplementedCacheServiceServer) SetMultiple(context.Context, *requests.SetMultipleRequest) (*responses.SetMultipleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetMultiple not implemented")
}
func (UnimplementedCacheServiceServer) DeleteMultiple(context.Context, *requests.DeleteMultipleRequest) (*responses.DeleteMultipleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMultiple not implemented")
}
func (UnimplementedCacheServiceServer) Increment(context.Context, *requests.IncrementRequest) (*responses.IncrementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Increment not implemented")
}
func (UnimplementedCacheServiceServer) Decrement(context.Context, *requests.DecrementRequest) (*responses.DecrementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Decrement not implemented")
}
func (UnimplementedCacheServiceServer) Clear(context.Context, *requests.ClearRequest) (*responses.ClearResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Clear not implemented")
}
func (UnimplementedCacheServiceServer) Keys(context.Context, *requests.KeysRequest) (*responses.KeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Keys not implemented")
}
func (UnimplementedCacheServiceServer) GetStats(context.Context, *requests.GetStatsRequest) (*responses.GetStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStats not implemented")
}
func (UnimplementedCacheServiceServer) Flush(context.Context, *requests.FlushRequest) (*responses.FlushResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Flush not implemented")
}
func (UnimplementedCacheServiceServer) TouchExpiration(context.Context, *requests.TouchExpirationRequest) (*responses.TouchExpirationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TouchExpiration not implemented")
}
func (UnimplementedCacheServiceServer) testEmbeddedByValue() {}

// UnsafeCacheServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CacheServiceServer will
// result in compilation errors.
type UnsafeCacheServiceServer interface {
	mustEmbedUnimplementedCacheServiceServer()
}

func RegisterCacheServiceServer(s grpc.ServiceRegistrar, srv CacheServiceServer) {
	// If the following call pancis, it indicates UnimplementedCacheServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CacheService_ServiceDesc, srv)
}

func _CacheService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(requests.GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CacheService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).Get(ctx, req.(*requests.GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(requests.SetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CacheService_Set_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).Set(ctx, req.(*requests.SetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(requests.DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CacheService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).Delete(ctx, req.(*requests.DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_Exists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(requests.ExistsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).Exists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CacheService_Exists_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).Exists(ctx, req.(*requests.ExistsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_GetMultiple_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(requests.GetMultipleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).GetMultiple(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CacheService_GetMultiple_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).GetMultiple(ctx, req.(*requests.GetMultipleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_SetMultiple_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(requests.SetMultipleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).SetMultiple(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CacheService_SetMultiple_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).SetMultiple(ctx, req.(*requests.SetMultipleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_DeleteMultiple_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(requests.DeleteMultipleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).DeleteMultiple(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CacheService_DeleteMultiple_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).DeleteMultiple(ctx, req.(*requests.DeleteMultipleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_Increment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(requests.IncrementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).Increment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CacheService_Increment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).Increment(ctx, req.(*requests.IncrementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_Decrement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(requests.DecrementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).Decrement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CacheService_Decrement_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).Decrement(ctx, req.(*requests.DecrementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_Clear_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(requests.ClearRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).Clear(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CacheService_Clear_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).Clear(ctx, req.(*requests.ClearRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_Keys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(requests.KeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).Keys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CacheService_Keys_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).Keys(ctx, req.(*requests.KeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_GetStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(requests.GetStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).GetStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CacheService_GetStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).GetStats(ctx, req.(*requests.GetStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_Flush_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(requests.FlushRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).Flush(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CacheService_Flush_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).Flush(ctx, req.(*requests.FlushRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_TouchExpiration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(requests.TouchExpirationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).TouchExpiration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CacheService_TouchExpiration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).TouchExpiration(ctx, req.(*requests.TouchExpirationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CacheService_ServiceDesc is the grpc.ServiceDesc for CacheService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CacheService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gcommon.v1.cache.CacheService",
	HandlerType: (*CacheServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _CacheService_Get_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _CacheService_Set_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CacheService_Delete_Handler,
		},
		{
			MethodName: "Exists",
			Handler:    _CacheService_Exists_Handler,
		},
		{
			MethodName: "GetMultiple",
			Handler:    _CacheService_GetMultiple_Handler,
		},
		{
			MethodName: "SetMultiple",
			Handler:    _CacheService_SetMultiple_Handler,
		},
		{
			MethodName: "DeleteMultiple",
			Handler:    _CacheService_DeleteMultiple_Handler,
		},
		{
			MethodName: "Increment",
			Handler:    _CacheService_Increment_Handler,
		},
		{
			MethodName: "Decrement",
			Handler:    _CacheService_Decrement_Handler,
		},
		{
			MethodName: "Clear",
			Handler:    _CacheService_Clear_Handler,
		},
		{
			MethodName: "Keys",
			Handler:    _CacheService_Keys_Handler,
		},
		{
			MethodName: "GetStats",
			Handler:    _CacheService_GetStats_Handler,
		},
		{
			MethodName: "Flush",
			Handler:    _CacheService_Flush_Handler,
		},
		{
			MethodName: "TouchExpiration",
			Handler:    _CacheService_TouchExpiration_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/cache/proto/services/cache_service.proto",
}

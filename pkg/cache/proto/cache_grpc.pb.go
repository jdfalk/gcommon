// file: pkg/cache/proto/cache.proto

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: pkg/cache/proto/cache.proto

package cachepb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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
// CacheService provides comprehensive caching capabilities
type CacheServiceClient interface {
	// Get retrieves a value from the cache
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	// Set stores a value in the cache
	Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetResponse, error)
	// Delete removes a value from the cache
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	// Exists checks if a key exists in the cache
	Exists(ctx context.Context, in *ExistsRequest, opts ...grpc.CallOption) (*ExistsResponse, error)
	// GetMultiple retrieves multiple values from the cache
	GetMultiple(ctx context.Context, in *GetMultipleRequest, opts ...grpc.CallOption) (*GetMultipleResponse, error)
	// SetMultiple stores multiple values in the cache
	SetMultiple(ctx context.Context, in *SetMultipleRequest, opts ...grpc.CallOption) (*SetMultipleResponse, error)
	// DeleteMultiple removes multiple values from the cache
	DeleteMultiple(ctx context.Context, in *DeleteMultipleRequest, opts ...grpc.CallOption) (*DeleteMultipleResponse, error)
	// Increment atomically increments a numeric value
	Increment(ctx context.Context, in *IncrementRequest, opts ...grpc.CallOption) (*IncrementResponse, error)
	// Decrement atomically decrements a numeric value
	Decrement(ctx context.Context, in *DecrementRequest, opts ...grpc.CallOption) (*DecrementResponse, error)
	// Clear removes all entries from the cache or by pattern
	Clear(ctx context.Context, in *ClearRequest, opts ...grpc.CallOption) (*ClearResponse, error)
	// Keys returns all keys matching a pattern
	Keys(ctx context.Context, in *KeysRequest, opts ...grpc.CallOption) (*KeysResponse, error)
	// GetStats returns cache statistics and metrics
	GetStats(ctx context.Context, in *GetStatsRequest, opts ...grpc.CallOption) (*GetStatsResponse, error)
	// Flush forces cache persistence if supported
	Flush(ctx context.Context, in *FlushRequest, opts ...grpc.CallOption) (*FlushResponse, error)
	// TouchExpiration updates the expiration time of a key
	TouchExpiration(ctx context.Context, in *TouchExpirationRequest, opts ...grpc.CallOption) (*TouchExpirationResponse, error)
}

type cacheServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCacheServiceClient(cc grpc.ClientConnInterface) CacheServiceClient {
	return &cacheServiceClient{cc}
}

func (c *cacheServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, CacheService_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetResponse)
	err := c.cc.Invoke(ctx, CacheService_Set_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, CacheService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) Exists(ctx context.Context, in *ExistsRequest, opts ...grpc.CallOption) (*ExistsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExistsResponse)
	err := c.cc.Invoke(ctx, CacheService_Exists_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) GetMultiple(ctx context.Context, in *GetMultipleRequest, opts ...grpc.CallOption) (*GetMultipleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMultipleResponse)
	err := c.cc.Invoke(ctx, CacheService_GetMultiple_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) SetMultiple(ctx context.Context, in *SetMultipleRequest, opts ...grpc.CallOption) (*SetMultipleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetMultipleResponse)
	err := c.cc.Invoke(ctx, CacheService_SetMultiple_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) DeleteMultiple(ctx context.Context, in *DeleteMultipleRequest, opts ...grpc.CallOption) (*DeleteMultipleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteMultipleResponse)
	err := c.cc.Invoke(ctx, CacheService_DeleteMultiple_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) Increment(ctx context.Context, in *IncrementRequest, opts ...grpc.CallOption) (*IncrementResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(IncrementResponse)
	err := c.cc.Invoke(ctx, CacheService_Increment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) Decrement(ctx context.Context, in *DecrementRequest, opts ...grpc.CallOption) (*DecrementResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DecrementResponse)
	err := c.cc.Invoke(ctx, CacheService_Decrement_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) Clear(ctx context.Context, in *ClearRequest, opts ...grpc.CallOption) (*ClearResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ClearResponse)
	err := c.cc.Invoke(ctx, CacheService_Clear_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) Keys(ctx context.Context, in *KeysRequest, opts ...grpc.CallOption) (*KeysResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(KeysResponse)
	err := c.cc.Invoke(ctx, CacheService_Keys_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) GetStats(ctx context.Context, in *GetStatsRequest, opts ...grpc.CallOption) (*GetStatsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetStatsResponse)
	err := c.cc.Invoke(ctx, CacheService_GetStats_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) Flush(ctx context.Context, in *FlushRequest, opts ...grpc.CallOption) (*FlushResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FlushResponse)
	err := c.cc.Invoke(ctx, CacheService_Flush_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) TouchExpiration(ctx context.Context, in *TouchExpirationRequest, opts ...grpc.CallOption) (*TouchExpirationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TouchExpirationResponse)
	err := c.cc.Invoke(ctx, CacheService_TouchExpiration_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CacheServiceServer is the server API for CacheService service.
// All implementations must embed UnimplementedCacheServiceServer
// for forward compatibility.
//
// CacheService provides comprehensive caching capabilities
type CacheServiceServer interface {
	// Get retrieves a value from the cache
	Get(context.Context, *GetRequest) (*GetResponse, error)
	// Set stores a value in the cache
	Set(context.Context, *SetRequest) (*SetResponse, error)
	// Delete removes a value from the cache
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	// Exists checks if a key exists in the cache
	Exists(context.Context, *ExistsRequest) (*ExistsResponse, error)
	// GetMultiple retrieves multiple values from the cache
	GetMultiple(context.Context, *GetMultipleRequest) (*GetMultipleResponse, error)
	// SetMultiple stores multiple values in the cache
	SetMultiple(context.Context, *SetMultipleRequest) (*SetMultipleResponse, error)
	// DeleteMultiple removes multiple values from the cache
	DeleteMultiple(context.Context, *DeleteMultipleRequest) (*DeleteMultipleResponse, error)
	// Increment atomically increments a numeric value
	Increment(context.Context, *IncrementRequest) (*IncrementResponse, error)
	// Decrement atomically decrements a numeric value
	Decrement(context.Context, *DecrementRequest) (*DecrementResponse, error)
	// Clear removes all entries from the cache or by pattern
	Clear(context.Context, *ClearRequest) (*ClearResponse, error)
	// Keys returns all keys matching a pattern
	Keys(context.Context, *KeysRequest) (*KeysResponse, error)
	// GetStats returns cache statistics and metrics
	GetStats(context.Context, *GetStatsRequest) (*GetStatsResponse, error)
	// Flush forces cache persistence if supported
	Flush(context.Context, *FlushRequest) (*FlushResponse, error)
	// TouchExpiration updates the expiration time of a key
	TouchExpiration(context.Context, *TouchExpirationRequest) (*TouchExpirationResponse, error)
	mustEmbedUnimplementedCacheServiceServer()
}

// UnimplementedCacheServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCacheServiceServer struct{}

func (UnimplementedCacheServiceServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedCacheServiceServer) Set(context.Context, *SetRequest) (*SetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (UnimplementedCacheServiceServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedCacheServiceServer) Exists(context.Context, *ExistsRequest) (*ExistsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Exists not implemented")
}
func (UnimplementedCacheServiceServer) GetMultiple(context.Context, *GetMultipleRequest) (*GetMultipleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMultiple not implemented")
}
func (UnimplementedCacheServiceServer) SetMultiple(context.Context, *SetMultipleRequest) (*SetMultipleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetMultiple not implemented")
}
func (UnimplementedCacheServiceServer) DeleteMultiple(context.Context, *DeleteMultipleRequest) (*DeleteMultipleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMultiple not implemented")
}
func (UnimplementedCacheServiceServer) Increment(context.Context, *IncrementRequest) (*IncrementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Increment not implemented")
}
func (UnimplementedCacheServiceServer) Decrement(context.Context, *DecrementRequest) (*DecrementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Decrement not implemented")
}
func (UnimplementedCacheServiceServer) Clear(context.Context, *ClearRequest) (*ClearResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Clear not implemented")
}
func (UnimplementedCacheServiceServer) Keys(context.Context, *KeysRequest) (*KeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Keys not implemented")
}
func (UnimplementedCacheServiceServer) GetStats(context.Context, *GetStatsRequest) (*GetStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStats not implemented")
}
func (UnimplementedCacheServiceServer) Flush(context.Context, *FlushRequest) (*FlushResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Flush not implemented")
}
func (UnimplementedCacheServiceServer) TouchExpiration(context.Context, *TouchExpirationRequest) (*TouchExpirationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TouchExpiration not implemented")
}
func (UnimplementedCacheServiceServer) mustEmbedUnimplementedCacheServiceServer() {}
func (UnimplementedCacheServiceServer) testEmbeddedByValue()                      {}

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
	in := new(GetRequest)
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
		return srv.(CacheServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRequest)
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
		return srv.(CacheServiceServer).Set(ctx, req.(*SetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
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
		return srv.(CacheServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_Exists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistsRequest)
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
		return srv.(CacheServiceServer).Exists(ctx, req.(*ExistsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_GetMultiple_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMultipleRequest)
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
		return srv.(CacheServiceServer).GetMultiple(ctx, req.(*GetMultipleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_SetMultiple_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetMultipleRequest)
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
		return srv.(CacheServiceServer).SetMultiple(ctx, req.(*SetMultipleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_DeleteMultiple_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMultipleRequest)
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
		return srv.(CacheServiceServer).DeleteMultiple(ctx, req.(*DeleteMultipleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_Increment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncrementRequest)
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
		return srv.(CacheServiceServer).Increment(ctx, req.(*IncrementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_Decrement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecrementRequest)
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
		return srv.(CacheServiceServer).Decrement(ctx, req.(*DecrementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_Clear_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClearRequest)
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
		return srv.(CacheServiceServer).Clear(ctx, req.(*ClearRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_Keys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeysRequest)
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
		return srv.(CacheServiceServer).Keys(ctx, req.(*KeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_GetStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatsRequest)
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
		return srv.(CacheServiceServer).GetStats(ctx, req.(*GetStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_Flush_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlushRequest)
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
		return srv.(CacheServiceServer).Flush(ctx, req.(*FlushRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_TouchExpiration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TouchExpirationRequest)
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
		return srv.(CacheServiceServer).TouchExpiration(ctx, req.(*TouchExpirationRequest))
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
	Metadata: "pkg/cache/proto/cache.proto",
}

const (
	CacheAdminService_CreateNamespace_FullMethodName   = "/gcommon.v1.cache.CacheAdminService/CreateNamespace"
	CacheAdminService_DeleteNamespace_FullMethodName   = "/gcommon.v1.cache.CacheAdminService/DeleteNamespace"
	CacheAdminService_ListNamespaces_FullMethodName    = "/gcommon.v1.cache.CacheAdminService/ListNamespaces"
	CacheAdminService_GetNamespaceStats_FullMethodName = "/gcommon.v1.cache.CacheAdminService/GetNamespaceStats"
	CacheAdminService_ConfigurePolicy_FullMethodName   = "/gcommon.v1.cache.CacheAdminService/ConfigurePolicy"
)

// CacheAdminServiceClient is the client API for CacheAdminService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Cache management service for administrative operations
type CacheAdminServiceClient interface {
	// CreateNamespace creates a new cache namespace
	CreateNamespace(ctx context.Context, in *CreateNamespaceRequest, opts ...grpc.CallOption) (*CreateNamespaceResponse, error)
	// DeleteNamespace removes a cache namespace
	DeleteNamespace(ctx context.Context, in *DeleteNamespaceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// ListNamespaces returns all available namespaces
	ListNamespaces(ctx context.Context, in *ListNamespacesRequest, opts ...grpc.CallOption) (*ListNamespacesResponse, error)
	// GetNamespaceStats returns statistics for a namespace
	GetNamespaceStats(ctx context.Context, in *GetNamespaceStatsRequest, opts ...grpc.CallOption) (*GetNamespaceStatsResponse, error)
	// ConfigurePolicy sets cache policies for a namespace
	ConfigurePolicy(ctx context.Context, in *ConfigurePolicyRequest, opts ...grpc.CallOption) (*ConfigurePolicyResponse, error)
}

type cacheAdminServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCacheAdminServiceClient(cc grpc.ClientConnInterface) CacheAdminServiceClient {
	return &cacheAdminServiceClient{cc}
}

func (c *cacheAdminServiceClient) CreateNamespace(ctx context.Context, in *CreateNamespaceRequest, opts ...grpc.CallOption) (*CreateNamespaceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateNamespaceResponse)
	err := c.cc.Invoke(ctx, CacheAdminService_CreateNamespace_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheAdminServiceClient) DeleteNamespace(ctx context.Context, in *DeleteNamespaceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, CacheAdminService_DeleteNamespace_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheAdminServiceClient) ListNamespaces(ctx context.Context, in *ListNamespacesRequest, opts ...grpc.CallOption) (*ListNamespacesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListNamespacesResponse)
	err := c.cc.Invoke(ctx, CacheAdminService_ListNamespaces_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheAdminServiceClient) GetNamespaceStats(ctx context.Context, in *GetNamespaceStatsRequest, opts ...grpc.CallOption) (*GetNamespaceStatsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetNamespaceStatsResponse)
	err := c.cc.Invoke(ctx, CacheAdminService_GetNamespaceStats_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheAdminServiceClient) ConfigurePolicy(ctx context.Context, in *ConfigurePolicyRequest, opts ...grpc.CallOption) (*ConfigurePolicyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ConfigurePolicyResponse)
	err := c.cc.Invoke(ctx, CacheAdminService_ConfigurePolicy_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CacheAdminServiceServer is the server API for CacheAdminService service.
// All implementations must embed UnimplementedCacheAdminServiceServer
// for forward compatibility.
//
// Cache management service for administrative operations
type CacheAdminServiceServer interface {
	// CreateNamespace creates a new cache namespace
	CreateNamespace(context.Context, *CreateNamespaceRequest) (*CreateNamespaceResponse, error)
	// DeleteNamespace removes a cache namespace
	DeleteNamespace(context.Context, *DeleteNamespaceRequest) (*emptypb.Empty, error)
	// ListNamespaces returns all available namespaces
	ListNamespaces(context.Context, *ListNamespacesRequest) (*ListNamespacesResponse, error)
	// GetNamespaceStats returns statistics for a namespace
	GetNamespaceStats(context.Context, *GetNamespaceStatsRequest) (*GetNamespaceStatsResponse, error)
	// ConfigurePolicy sets cache policies for a namespace
	ConfigurePolicy(context.Context, *ConfigurePolicyRequest) (*ConfigurePolicyResponse, error)
	mustEmbedUnimplementedCacheAdminServiceServer()
}

// UnimplementedCacheAdminServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCacheAdminServiceServer struct{}

func (UnimplementedCacheAdminServiceServer) CreateNamespace(context.Context, *CreateNamespaceRequest) (*CreateNamespaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNamespace not implemented")
}
func (UnimplementedCacheAdminServiceServer) DeleteNamespace(context.Context, *DeleteNamespaceRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNamespace not implemented")
}
func (UnimplementedCacheAdminServiceServer) ListNamespaces(context.Context, *ListNamespacesRequest) (*ListNamespacesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNamespaces not implemented")
}
func (UnimplementedCacheAdminServiceServer) GetNamespaceStats(context.Context, *GetNamespaceStatsRequest) (*GetNamespaceStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNamespaceStats not implemented")
}
func (UnimplementedCacheAdminServiceServer) ConfigurePolicy(context.Context, *ConfigurePolicyRequest) (*ConfigurePolicyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigurePolicy not implemented")
}
func (UnimplementedCacheAdminServiceServer) mustEmbedUnimplementedCacheAdminServiceServer() {}
func (UnimplementedCacheAdminServiceServer) testEmbeddedByValue()                           {}

// UnsafeCacheAdminServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CacheAdminServiceServer will
// result in compilation errors.
type UnsafeCacheAdminServiceServer interface {
	mustEmbedUnimplementedCacheAdminServiceServer()
}

func RegisterCacheAdminServiceServer(s grpc.ServiceRegistrar, srv CacheAdminServiceServer) {
	// If the following call pancis, it indicates UnimplementedCacheAdminServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CacheAdminService_ServiceDesc, srv)
}

func _CacheAdminService_CreateNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheAdminServiceServer).CreateNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CacheAdminService_CreateNamespace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheAdminServiceServer).CreateNamespace(ctx, req.(*CreateNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheAdminService_DeleteNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheAdminServiceServer).DeleteNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CacheAdminService_DeleteNamespace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheAdminServiceServer).DeleteNamespace(ctx, req.(*DeleteNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheAdminService_ListNamespaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNamespacesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheAdminServiceServer).ListNamespaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CacheAdminService_ListNamespaces_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheAdminServiceServer).ListNamespaces(ctx, req.(*ListNamespacesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheAdminService_GetNamespaceStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNamespaceStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheAdminServiceServer).GetNamespaceStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CacheAdminService_GetNamespaceStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheAdminServiceServer).GetNamespaceStats(ctx, req.(*GetNamespaceStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheAdminService_ConfigurePolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigurePolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheAdminServiceServer).ConfigurePolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CacheAdminService_ConfigurePolicy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheAdminServiceServer).ConfigurePolicy(ctx, req.(*ConfigurePolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CacheAdminService_ServiceDesc is the grpc.ServiceDesc for CacheAdminService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CacheAdminService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gcommon.v1.cache.CacheAdminService",
	HandlerType: (*CacheAdminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNamespace",
			Handler:    _CacheAdminService_CreateNamespace_Handler,
		},
		{
			MethodName: "DeleteNamespace",
			Handler:    _CacheAdminService_DeleteNamespace_Handler,
		},
		{
			MethodName: "ListNamespaces",
			Handler:    _CacheAdminService_ListNamespaces_Handler,
		},
		{
			MethodName: "GetNamespaceStats",
			Handler:    _CacheAdminService_GetNamespaceStats_Handler,
		},
		{
			MethodName: "ConfigurePolicy",
			Handler:    _CacheAdminService_ConfigurePolicy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/cache/proto/cache.proto",
}

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: pkg/organization/proto/services/tenant_service.proto

package organizationpb

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
	TenantService_CreateTenant_FullMethodName             = "/gcommon.v1.organization.TenantService/CreateTenant"
	TenantService_GetTenant_FullMethodName                = "/gcommon.v1.organization.TenantService/GetTenant"
	TenantService_UpdateTenant_FullMethodName             = "/gcommon.v1.organization.TenantService/UpdateTenant"
	TenantService_DeleteTenant_FullMethodName             = "/gcommon.v1.organization.TenantService/DeleteTenant"
	TenantService_ListTenants_FullMethodName              = "/gcommon.v1.organization.TenantService/ListTenants"
	TenantService_ConfigureTenantIsolation_FullMethodName = "/gcommon.v1.organization.TenantService/ConfigureTenantIsolation"
	TenantService_GetTenantIsolation_FullMethodName       = "/gcommon.v1.organization.TenantService/GetTenantIsolation"
	TenantService_UpdateTenantQuota_FullMethodName        = "/gcommon.v1.organization.TenantService/UpdateTenantQuota"
	TenantService_GetTenantUsage_FullMethodName           = "/gcommon.v1.organization.TenantService/GetTenantUsage"
)

// TenantServiceClient is the client API for TenantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// *
// TenantService provides comprehensive tenant management capabilities.
// Handles tenant CRUD operations, isolation configuration, quota management,
// and multi-tenant resource administration.
type TenantServiceClient interface {
	// Create a new tenant within an organization
	CreateTenant(ctx context.Context, in *CreateTenantRequest, opts ...grpc.CallOption) (*CreateTenantResponse, error)
	// Get a tenant by ID
	GetTenant(ctx context.Context, in *GetTenantRequest, opts ...grpc.CallOption) (*GetTenantResponse, error)
	// Update an existing tenant
	UpdateTenant(ctx context.Context, in *UpdateTenantRequest, opts ...grpc.CallOption) (*UpdateTenantResponse, error)
	// Delete a tenant (soft delete)
	DeleteTenant(ctx context.Context, in *DeleteTenantRequest, opts ...grpc.CallOption) (*DeleteTenantResponse, error)
	// List tenants within an organization (with pagination and filtering)
	ListTenants(ctx context.Context, in *ListTenantsRequest, opts ...grpc.CallOption) (*ListTenantsResponse, error)
	// Configure tenant isolation settings
	ConfigureTenantIsolation(ctx context.Context, in *ConfigureTenantIsolationRequest, opts ...grpc.CallOption) (*ConfigureTenantIsolationResponse, error)
	// Get tenant isolation configuration
	GetTenantIsolation(ctx context.Context, in *GetTenantIsolationRequest, opts ...grpc.CallOption) (*GetTenantIsolationResponse, error)
	// Update tenant resource quotas
	UpdateTenantQuota(ctx context.Context, in *UpdateTenantQuotaRequest, opts ...grpc.CallOption) (*UpdateTenantQuotaResponse, error)
	// Get tenant resource usage statistics
	GetTenantUsage(ctx context.Context, in *GetTenantUsageRequest, opts ...grpc.CallOption) (*GetTenantUsageResponse, error)
}

type tenantServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTenantServiceClient(cc grpc.ClientConnInterface) TenantServiceClient {
	return &tenantServiceClient{cc}
}

func (c *tenantServiceClient) CreateTenant(ctx context.Context, in *CreateTenantRequest, opts ...grpc.CallOption) (*CreateTenantResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateTenantResponse)
	err := c.cc.Invoke(ctx, TenantService_CreateTenant_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantServiceClient) GetTenant(ctx context.Context, in *GetTenantRequest, opts ...grpc.CallOption) (*GetTenantResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTenantResponse)
	err := c.cc.Invoke(ctx, TenantService_GetTenant_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantServiceClient) UpdateTenant(ctx context.Context, in *UpdateTenantRequest, opts ...grpc.CallOption) (*UpdateTenantResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateTenantResponse)
	err := c.cc.Invoke(ctx, TenantService_UpdateTenant_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantServiceClient) DeleteTenant(ctx context.Context, in *DeleteTenantRequest, opts ...grpc.CallOption) (*DeleteTenantResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteTenantResponse)
	err := c.cc.Invoke(ctx, TenantService_DeleteTenant_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantServiceClient) ListTenants(ctx context.Context, in *ListTenantsRequest, opts ...grpc.CallOption) (*ListTenantsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListTenantsResponse)
	err := c.cc.Invoke(ctx, TenantService_ListTenants_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantServiceClient) ConfigureTenantIsolation(ctx context.Context, in *ConfigureTenantIsolationRequest, opts ...grpc.CallOption) (*ConfigureTenantIsolationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ConfigureTenantIsolationResponse)
	err := c.cc.Invoke(ctx, TenantService_ConfigureTenantIsolation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantServiceClient) GetTenantIsolation(ctx context.Context, in *GetTenantIsolationRequest, opts ...grpc.CallOption) (*GetTenantIsolationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTenantIsolationResponse)
	err := c.cc.Invoke(ctx, TenantService_GetTenantIsolation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantServiceClient) UpdateTenantQuota(ctx context.Context, in *UpdateTenantQuotaRequest, opts ...grpc.CallOption) (*UpdateTenantQuotaResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateTenantQuotaResponse)
	err := c.cc.Invoke(ctx, TenantService_UpdateTenantQuota_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantServiceClient) GetTenantUsage(ctx context.Context, in *GetTenantUsageRequest, opts ...grpc.CallOption) (*GetTenantUsageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTenantUsageResponse)
	err := c.cc.Invoke(ctx, TenantService_GetTenantUsage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TenantServiceServer is the server API for TenantService service.
// All implementations should embed UnimplementedTenantServiceServer
// for forward compatibility.
//
// *
// TenantService provides comprehensive tenant management capabilities.
// Handles tenant CRUD operations, isolation configuration, quota management,
// and multi-tenant resource administration.
type TenantServiceServer interface {
	// Create a new tenant within an organization
	CreateTenant(context.Context, *CreateTenantRequest) (*CreateTenantResponse, error)
	// Get a tenant by ID
	GetTenant(context.Context, *GetTenantRequest) (*GetTenantResponse, error)
	// Update an existing tenant
	UpdateTenant(context.Context, *UpdateTenantRequest) (*UpdateTenantResponse, error)
	// Delete a tenant (soft delete)
	DeleteTenant(context.Context, *DeleteTenantRequest) (*DeleteTenantResponse, error)
	// List tenants within an organization (with pagination and filtering)
	ListTenants(context.Context, *ListTenantsRequest) (*ListTenantsResponse, error)
	// Configure tenant isolation settings
	ConfigureTenantIsolation(context.Context, *ConfigureTenantIsolationRequest) (*ConfigureTenantIsolationResponse, error)
	// Get tenant isolation configuration
	GetTenantIsolation(context.Context, *GetTenantIsolationRequest) (*GetTenantIsolationResponse, error)
	// Update tenant resource quotas
	UpdateTenantQuota(context.Context, *UpdateTenantQuotaRequest) (*UpdateTenantQuotaResponse, error)
	// Get tenant resource usage statistics
	GetTenantUsage(context.Context, *GetTenantUsageRequest) (*GetTenantUsageResponse, error)
}

// UnimplementedTenantServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTenantServiceServer struct{}

func (UnimplementedTenantServiceServer) CreateTenant(context.Context, *CreateTenantRequest) (*CreateTenantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTenant not implemented")
}
func (UnimplementedTenantServiceServer) GetTenant(context.Context, *GetTenantRequest) (*GetTenantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTenant not implemented")
}
func (UnimplementedTenantServiceServer) UpdateTenant(context.Context, *UpdateTenantRequest) (*UpdateTenantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTenant not implemented")
}
func (UnimplementedTenantServiceServer) DeleteTenant(context.Context, *DeleteTenantRequest) (*DeleteTenantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTenant not implemented")
}
func (UnimplementedTenantServiceServer) ListTenants(context.Context, *ListTenantsRequest) (*ListTenantsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTenants not implemented")
}
func (UnimplementedTenantServiceServer) ConfigureTenantIsolation(context.Context, *ConfigureTenantIsolationRequest) (*ConfigureTenantIsolationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigureTenantIsolation not implemented")
}
func (UnimplementedTenantServiceServer) GetTenantIsolation(context.Context, *GetTenantIsolationRequest) (*GetTenantIsolationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTenantIsolation not implemented")
}
func (UnimplementedTenantServiceServer) UpdateTenantQuota(context.Context, *UpdateTenantQuotaRequest) (*UpdateTenantQuotaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTenantQuota not implemented")
}
func (UnimplementedTenantServiceServer) GetTenantUsage(context.Context, *GetTenantUsageRequest) (*GetTenantUsageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTenantUsage not implemented")
}
func (UnimplementedTenantServiceServer) testEmbeddedByValue() {}

// UnsafeTenantServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TenantServiceServer will
// result in compilation errors.
type UnsafeTenantServiceServer interface {
	mustEmbedUnimplementedTenantServiceServer()
}

func RegisterTenantServiceServer(s grpc.ServiceRegistrar, srv TenantServiceServer) {
	// If the following call pancis, it indicates UnimplementedTenantServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TenantService_ServiceDesc, srv)
}

func _TenantService_CreateTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServiceServer).CreateTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TenantService_CreateTenant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServiceServer).CreateTenant(ctx, req.(*CreateTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantService_GetTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServiceServer).GetTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TenantService_GetTenant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServiceServer).GetTenant(ctx, req.(*GetTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantService_UpdateTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServiceServer).UpdateTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TenantService_UpdateTenant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServiceServer).UpdateTenant(ctx, req.(*UpdateTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantService_DeleteTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServiceServer).DeleteTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TenantService_DeleteTenant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServiceServer).DeleteTenant(ctx, req.(*DeleteTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantService_ListTenants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTenantsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServiceServer).ListTenants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TenantService_ListTenants_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServiceServer).ListTenants(ctx, req.(*ListTenantsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantService_ConfigureTenantIsolation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigureTenantIsolationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServiceServer).ConfigureTenantIsolation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TenantService_ConfigureTenantIsolation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServiceServer).ConfigureTenantIsolation(ctx, req.(*ConfigureTenantIsolationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantService_GetTenantIsolation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTenantIsolationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServiceServer).GetTenantIsolation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TenantService_GetTenantIsolation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServiceServer).GetTenantIsolation(ctx, req.(*GetTenantIsolationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantService_UpdateTenantQuota_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTenantQuotaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServiceServer).UpdateTenantQuota(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TenantService_UpdateTenantQuota_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServiceServer).UpdateTenantQuota(ctx, req.(*UpdateTenantQuotaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantService_GetTenantUsage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTenantUsageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServiceServer).GetTenantUsage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TenantService_GetTenantUsage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServiceServer).GetTenantUsage(ctx, req.(*GetTenantUsageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TenantService_ServiceDesc is the grpc.ServiceDesc for TenantService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TenantService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gcommon.v1.organization.TenantService",
	HandlerType: (*TenantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTenant",
			Handler:    _TenantService_CreateTenant_Handler,
		},
		{
			MethodName: "GetTenant",
			Handler:    _TenantService_GetTenant_Handler,
		},
		{
			MethodName: "UpdateTenant",
			Handler:    _TenantService_UpdateTenant_Handler,
		},
		{
			MethodName: "DeleteTenant",
			Handler:    _TenantService_DeleteTenant_Handler,
		},
		{
			MethodName: "ListTenants",
			Handler:    _TenantService_ListTenants_Handler,
		},
		{
			MethodName: "ConfigureTenantIsolation",
			Handler:    _TenantService_ConfigureTenantIsolation_Handler,
		},
		{
			MethodName: "GetTenantIsolation",
			Handler:    _TenantService_GetTenantIsolation_Handler,
		},
		{
			MethodName: "UpdateTenantQuota",
			Handler:    _TenantService_UpdateTenantQuota_Handler,
		},
		{
			MethodName: "GetTenantUsage",
			Handler:    _TenantService_GetTenantUsage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/organization/proto/services/tenant_service.proto",
}

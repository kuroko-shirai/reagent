// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: service_a.proto

package api

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
	ServiceA_GetHealth_FullMethodName = "/proto.ServiceA/GetHealth"
)

// ServiceAClient is the client API for ServiceA service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Example of some microservice for which we will generate proto.
type ServiceAClient interface {
	// Health check of the microservice for internal use.
	GetHealth(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error)
}

type serviceAClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceAClient(cc grpc.ClientConnInterface) ServiceAClient {
	return &serviceAClient{cc}
}

func (c *serviceAClient) GetHealth(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HealthResponse)
	err := c.cc.Invoke(ctx, ServiceA_GetHealth_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceAServer is the server API for ServiceA service.
// All implementations must embed UnimplementedServiceAServer
// for forward compatibility.
//
// Example of some microservice for which we will generate proto.
type ServiceAServer interface {
	// Health check of the microservice for internal use.
	GetHealth(context.Context, *HealthRequest) (*HealthResponse, error)
	mustEmbedUnimplementedServiceAServer()
}

// UnimplementedServiceAServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedServiceAServer struct{}

func (UnimplementedServiceAServer) GetHealth(context.Context, *HealthRequest) (*HealthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHealth not implemented")
}
func (UnimplementedServiceAServer) mustEmbedUnimplementedServiceAServer() {}
func (UnimplementedServiceAServer) testEmbeddedByValue()                  {}

// UnsafeServiceAServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceAServer will
// result in compilation errors.
type UnsafeServiceAServer interface {
	mustEmbedUnimplementedServiceAServer()
}

func RegisterServiceAServer(s grpc.ServiceRegistrar, srv ServiceAServer) {
	// If the following call pancis, it indicates UnimplementedServiceAServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ServiceA_ServiceDesc, srv)
}

func _ServiceA_GetHealth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceAServer).GetHealth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceA_GetHealth_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceAServer).GetHealth(ctx, req.(*HealthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ServiceA_ServiceDesc is the grpc.ServiceDesc for ServiceA service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServiceA_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ServiceA",
	HandlerType: (*ServiceAServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHealth",
			Handler:    _ServiceA_GetHealth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_a.proto",
}
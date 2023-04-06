// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.2
// source: authentication.proto

package go_pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	AuthenticationService_GetTokenDemo_FullMethodName = "/authenticate.AuthenticationService/GetTokenDemo"
)

// AuthenticationServiceClient is the client API for AuthenticationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthenticationServiceClient interface {
	GetTokenDemo(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
}

type authenticationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthenticationServiceClient(cc grpc.ClientConnInterface) AuthenticationServiceClient {
	return &authenticationServiceClient{cc}
}

func (c *authenticationServiceClient) GetTokenDemo(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, AuthenticationService_GetTokenDemo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticationServiceServer is the server API for AuthenticationService service.
// All implementations must embed UnimplementedAuthenticationServiceServer
// for forward compatibility
type AuthenticationServiceServer interface {
	GetTokenDemo(context.Context, *LoginRequest) (*LoginResponse, error)
	mustEmbedUnimplementedAuthenticationServiceServer()
}

// UnimplementedAuthenticationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthenticationServiceServer struct {
}

func (UnimplementedAuthenticationServiceServer) GetTokenDemo(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTokenDemo not implemented")
}
func (UnimplementedAuthenticationServiceServer) mustEmbedUnimplementedAuthenticationServiceServer() {}

// UnsafeAuthenticationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthenticationServiceServer will
// result in compilation errors.
type UnsafeAuthenticationServiceServer interface {
	mustEmbedUnimplementedAuthenticationServiceServer()
}

func RegisterAuthenticationServiceServer(s grpc.ServiceRegistrar, srv AuthenticationServiceServer) {
	s.RegisterService(&AuthenticationService_ServiceDesc, srv)
}

func _AuthenticationService_GetTokenDemo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).GetTokenDemo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthenticationService_GetTokenDemo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).GetTokenDemo(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthenticationService_ServiceDesc is the grpc.ServiceDesc for AuthenticationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthenticationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "authenticate.AuthenticationService",
	HandlerType: (*AuthenticationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTokenDemo",
			Handler:    _AuthenticationService_GetTokenDemo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authentication.proto",
}

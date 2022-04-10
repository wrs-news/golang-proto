// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/security/security-microservice.proto

package security

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SecurityServiceClient is the client API for SecurityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SecurityServiceClient interface {
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*TokensPair, error)
	AuthCheck(ctx context.Context, in *AuthCheckReq, opts ...grpc.CallOption) (*empty.Empty, error)
	RefreshToken(ctx context.Context, in *RefreshTokenReq, opts ...grpc.CallOption) (*TokensPair, error)
	Logout(ctx context.Context, in *LogoutReq, opts ...grpc.CallOption) (*empty.Empty, error)
	HeartbeatCheck(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
}

type securityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSecurityServiceClient(cc grpc.ClientConnInterface) SecurityServiceClient {
	return &securityServiceClient{cc}
}

func (c *securityServiceClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*TokensPair, error) {
	out := new(TokensPair)
	err := c.cc.Invoke(ctx, "/security.SecurityService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityServiceClient) AuthCheck(ctx context.Context, in *AuthCheckReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/security.SecurityService/AuthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityServiceClient) RefreshToken(ctx context.Context, in *RefreshTokenReq, opts ...grpc.CallOption) (*TokensPair, error) {
	out := new(TokensPair)
	err := c.cc.Invoke(ctx, "/security.SecurityService/RefreshToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityServiceClient) Logout(ctx context.Context, in *LogoutReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/security.SecurityService/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityServiceClient) HeartbeatCheck(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/security.SecurityService/HeartbeatCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecurityServiceServer is the server API for SecurityService service.
// All implementations must embed UnimplementedSecurityServiceServer
// for forward compatibility
type SecurityServiceServer interface {
	Login(context.Context, *LoginReq) (*TokensPair, error)
	AuthCheck(context.Context, *AuthCheckReq) (*empty.Empty, error)
	RefreshToken(context.Context, *RefreshTokenReq) (*TokensPair, error)
	Logout(context.Context, *LogoutReq) (*empty.Empty, error)
	HeartbeatCheck(context.Context, *empty.Empty) (*empty.Empty, error)
	mustEmbedUnimplementedSecurityServiceServer()
}

// UnimplementedSecurityServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSecurityServiceServer struct {
}

func (UnimplementedSecurityServiceServer) Login(context.Context, *LoginReq) (*TokensPair, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedSecurityServiceServer) AuthCheck(context.Context, *AuthCheckReq) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthCheck not implemented")
}
func (UnimplementedSecurityServiceServer) RefreshToken(context.Context, *RefreshTokenReq) (*TokensPair, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshToken not implemented")
}
func (UnimplementedSecurityServiceServer) Logout(context.Context, *LogoutReq) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedSecurityServiceServer) HeartbeatCheck(context.Context, *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HeartbeatCheck not implemented")
}
func (UnimplementedSecurityServiceServer) mustEmbedUnimplementedSecurityServiceServer() {}

// UnsafeSecurityServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SecurityServiceServer will
// result in compilation errors.
type UnsafeSecurityServiceServer interface {
	mustEmbedUnimplementedSecurityServiceServer()
}

func RegisterSecurityServiceServer(s grpc.ServiceRegistrar, srv SecurityServiceServer) {
	s.RegisterService(&SecurityService_ServiceDesc, srv)
}

func _SecurityService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/security.SecurityService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityServiceServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityService_AuthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthCheckReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityServiceServer).AuthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/security.SecurityService/AuthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityServiceServer).AuthCheck(ctx, req.(*AuthCheckReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityService_RefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityServiceServer).RefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/security.SecurityService/RefreshToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityServiceServer).RefreshToken(ctx, req.(*RefreshTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityService_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityServiceServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/security.SecurityService/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityServiceServer).Logout(ctx, req.(*LogoutReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityService_HeartbeatCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityServiceServer).HeartbeatCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/security.SecurityService/HeartbeatCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityServiceServer).HeartbeatCheck(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// SecurityService_ServiceDesc is the grpc.ServiceDesc for SecurityService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SecurityService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "security.SecurityService",
	HandlerType: (*SecurityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _SecurityService_Login_Handler,
		},
		{
			MethodName: "AuthCheck",
			Handler:    _SecurityService_AuthCheck_Handler,
		},
		{
			MethodName: "RefreshToken",
			Handler:    _SecurityService_RefreshToken_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _SecurityService_Logout_Handler,
		},
		{
			MethodName: "HeartbeatCheck",
			Handler:    _SecurityService_HeartbeatCheck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/security/security-microservice.proto",
}

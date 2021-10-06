// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package secret

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SecretServiceClient is the client API for SecretService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SecretServiceClient interface {
	CreateSecret(ctx context.Context, in *CreateSecretReq, opts ...grpc.CallOption) (*CreateSecretRsp, error)
	GetSecretList(ctx context.Context, in *GetSecretListReq, opts ...grpc.CallOption) (*GetSecretListRsp, error)
	UpdateSecret(ctx context.Context, in *UpdateSecretReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteSecret(ctx context.Context, in *DeleteSecretReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type secretServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSecretServiceClient(cc grpc.ClientConnInterface) SecretServiceClient {
	return &secretServiceClient{cc}
}

func (c *secretServiceClient) CreateSecret(ctx context.Context, in *CreateSecretReq, opts ...grpc.CallOption) (*CreateSecretRsp, error) {
	out := new(CreateSecretRsp)
	err := c.cc.Invoke(ctx, "/SecretService/CreateSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretServiceClient) GetSecretList(ctx context.Context, in *GetSecretListReq, opts ...grpc.CallOption) (*GetSecretListRsp, error) {
	out := new(GetSecretListRsp)
	err := c.cc.Invoke(ctx, "/SecretService/GetSecretList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretServiceClient) UpdateSecret(ctx context.Context, in *UpdateSecretReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/SecretService/UpdateSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretServiceClient) DeleteSecret(ctx context.Context, in *DeleteSecretReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/SecretService/DeleteSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecretServiceServer is the server API for SecretService service.
// All implementations must embed UnimplementedSecretServiceServer
// for forward compatibility
type SecretServiceServer interface {
	CreateSecret(context.Context, *CreateSecretReq) (*CreateSecretRsp, error)
	GetSecretList(context.Context, *GetSecretListReq) (*GetSecretListRsp, error)
	UpdateSecret(context.Context, *UpdateSecretReq) (*emptypb.Empty, error)
	DeleteSecret(context.Context, *DeleteSecretReq) (*emptypb.Empty, error)
	mustEmbedUnimplementedSecretServiceServer()
}

// UnimplementedSecretServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSecretServiceServer struct {
}

func (UnimplementedSecretServiceServer) CreateSecret(context.Context, *CreateSecretReq) (*CreateSecretRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSecret not implemented")
}
func (UnimplementedSecretServiceServer) GetSecretList(context.Context, *GetSecretListReq) (*GetSecretListRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSecretList not implemented")
}
func (UnimplementedSecretServiceServer) UpdateSecret(context.Context, *UpdateSecretReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSecret not implemented")
}
func (UnimplementedSecretServiceServer) DeleteSecret(context.Context, *DeleteSecretReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSecret not implemented")
}
func (UnimplementedSecretServiceServer) mustEmbedUnimplementedSecretServiceServer() {}

// UnsafeSecretServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SecretServiceServer will
// result in compilation errors.
type UnsafeSecretServiceServer interface {
	mustEmbedUnimplementedSecretServiceServer()
}

func RegisterSecretServiceServer(s grpc.ServiceRegistrar, srv SecretServiceServer) {
	s.RegisterService(&SecretService_ServiceDesc, srv)
}

func _SecretService_CreateSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSecretReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretServiceServer).CreateSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SecretService/CreateSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretServiceServer).CreateSecret(ctx, req.(*CreateSecretReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretService_GetSecretList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSecretListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretServiceServer).GetSecretList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SecretService/GetSecretList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretServiceServer).GetSecretList(ctx, req.(*GetSecretListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretService_UpdateSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSecretReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretServiceServer).UpdateSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SecretService/UpdateSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretServiceServer).UpdateSecret(ctx, req.(*UpdateSecretReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretService_DeleteSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSecretReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretServiceServer).DeleteSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SecretService/DeleteSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretServiceServer).DeleteSecret(ctx, req.(*DeleteSecretReq))
	}
	return interceptor(ctx, in, info, handler)
}

// SecretService_ServiceDesc is the grpc.ServiceDesc for SecretService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SecretService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "SecretService",
	HandlerType: (*SecretServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSecret",
			Handler:    _SecretService_CreateSecret_Handler,
		},
		{
			MethodName: "GetSecretList",
			Handler:    _SecretService_GetSecretList_Handler,
		},
		{
			MethodName: "UpdateSecret",
			Handler:    _SecretService_UpdateSecret_Handler,
		},
		{
			MethodName: "DeleteSecret",
			Handler:    _SecretService_DeleteSecret_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "secret/secret.proto",
}

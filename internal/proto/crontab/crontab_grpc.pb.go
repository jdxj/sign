// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package crontab

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

// CrontabServiceClient is the client API for CrontabService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CrontabServiceClient interface {
	CreateTask(ctx context.Context, in *CreateTaskReq, opts ...grpc.CallOption) (*CreateTaskRsp, error)
	GetTasks(ctx context.Context, in *GetTasksReq, opts ...grpc.CallOption) (*GetTasksRsp, error)
	DeleteTask(ctx context.Context, in *DeleteTaskReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type crontabServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCrontabServiceClient(cc grpc.ClientConnInterface) CrontabServiceClient {
	return &crontabServiceClient{cc}
}

func (c *crontabServiceClient) CreateTask(ctx context.Context, in *CreateTaskReq, opts ...grpc.CallOption) (*CreateTaskRsp, error) {
	out := new(CreateTaskRsp)
	err := c.cc.Invoke(ctx, "/CrontabService/CreateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crontabServiceClient) GetTasks(ctx context.Context, in *GetTasksReq, opts ...grpc.CallOption) (*GetTasksRsp, error) {
	out := new(GetTasksRsp)
	err := c.cc.Invoke(ctx, "/CrontabService/GetTasks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crontabServiceClient) DeleteTask(ctx context.Context, in *DeleteTaskReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/CrontabService/DeleteTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CrontabServiceServer is the server API for CrontabService service.
// All implementations must embed UnimplementedCrontabServiceServer
// for forward compatibility
type CrontabServiceServer interface {
	CreateTask(context.Context, *CreateTaskReq) (*CreateTaskRsp, error)
	GetTasks(context.Context, *GetTasksReq) (*GetTasksRsp, error)
	DeleteTask(context.Context, *DeleteTaskReq) (*emptypb.Empty, error)
	mustEmbedUnimplementedCrontabServiceServer()
}

// UnimplementedCrontabServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCrontabServiceServer struct {
}

func (UnimplementedCrontabServiceServer) CreateTask(context.Context, *CreateTaskReq) (*CreateTaskRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTask not implemented")
}
func (UnimplementedCrontabServiceServer) GetTasks(context.Context, *GetTasksReq) (*GetTasksRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTasks not implemented")
}
func (UnimplementedCrontabServiceServer) DeleteTask(context.Context, *DeleteTaskReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTask not implemented")
}
func (UnimplementedCrontabServiceServer) mustEmbedUnimplementedCrontabServiceServer() {}

// UnsafeCrontabServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CrontabServiceServer will
// result in compilation errors.
type UnsafeCrontabServiceServer interface {
	mustEmbedUnimplementedCrontabServiceServer()
}

func RegisterCrontabServiceServer(s grpc.ServiceRegistrar, srv CrontabServiceServer) {
	s.RegisterService(&CrontabService_ServiceDesc, srv)
}

func _CrontabService_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTaskReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrontabServiceServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CrontabService/CreateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrontabServiceServer).CreateTask(ctx, req.(*CreateTaskReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CrontabService_GetTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTasksReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrontabServiceServer).GetTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CrontabService/GetTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrontabServiceServer).GetTasks(ctx, req.(*GetTasksReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CrontabService_DeleteTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTaskReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrontabServiceServer).DeleteTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CrontabService/DeleteTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrontabServiceServer).DeleteTask(ctx, req.(*DeleteTaskReq))
	}
	return interceptor(ctx, in, info, handler)
}

// CrontabService_ServiceDesc is the grpc.ServiceDesc for CrontabService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CrontabService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "CrontabService",
	HandlerType: (*CrontabServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTask",
			Handler:    _CrontabService_CreateTask_Handler,
		},
		{
			MethodName: "GetTasks",
			Handler:    _CrontabService_GetTasks_Handler,
		},
		{
			MethodName: "DeleteTask",
			Handler:    _CrontabService_DeleteTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "crontab/crontab.proto",
}
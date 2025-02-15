// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.6
// source: mpaas/apps/job/pb/rpc.proto

package job

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
	RPC_CreateJob_FullMethodName       = "/infraboard.mpaas.job.RPC/CreateJob"
	RPC_QueryJob_FullMethodName        = "/infraboard.mpaas.job.RPC/QueryJob"
	RPC_DescribeJob_FullMethodName     = "/infraboard.mpaas.job.RPC/DescribeJob"
	RPC_UpdateJob_FullMethodName       = "/infraboard.mpaas.job.RPC/UpdateJob"
	RPC_UpdateJobStatus_FullMethodName = "/infraboard.mpaas.job.RPC/UpdateJobStatus"
)

// RPCClient is the client API for RPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RPCClient interface {
	CreateJob(ctx context.Context, in *CreateJobRequest, opts ...grpc.CallOption) (*Job, error)
	QueryJob(ctx context.Context, in *QueryJobRequest, opts ...grpc.CallOption) (*JobSet, error)
	DescribeJob(ctx context.Context, in *DescribeJobRequest, opts ...grpc.CallOption) (*Job, error)
	UpdateJob(ctx context.Context, in *UpdateJobRequest, opts ...grpc.CallOption) (*Job, error)
	// 编辑job状态, 比如发布
	UpdateJobStatus(ctx context.Context, in *UpdateJobStatusRequest, opts ...grpc.CallOption) (*Job, error)
}

type rPCClient struct {
	cc grpc.ClientConnInterface
}

func NewRPCClient(cc grpc.ClientConnInterface) RPCClient {
	return &rPCClient{cc}
}

func (c *rPCClient) CreateJob(ctx context.Context, in *CreateJobRequest, opts ...grpc.CallOption) (*Job, error) {
	out := new(Job)
	err := c.cc.Invoke(ctx, RPC_CreateJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) QueryJob(ctx context.Context, in *QueryJobRequest, opts ...grpc.CallOption) (*JobSet, error) {
	out := new(JobSet)
	err := c.cc.Invoke(ctx, RPC_QueryJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) DescribeJob(ctx context.Context, in *DescribeJobRequest, opts ...grpc.CallOption) (*Job, error) {
	out := new(Job)
	err := c.cc.Invoke(ctx, RPC_DescribeJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) UpdateJob(ctx context.Context, in *UpdateJobRequest, opts ...grpc.CallOption) (*Job, error) {
	out := new(Job)
	err := c.cc.Invoke(ctx, RPC_UpdateJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) UpdateJobStatus(ctx context.Context, in *UpdateJobStatusRequest, opts ...grpc.CallOption) (*Job, error) {
	out := new(Job)
	err := c.cc.Invoke(ctx, RPC_UpdateJobStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RPCServer is the server API for RPC service.
// All implementations must embed UnimplementedRPCServer
// for forward compatibility
type RPCServer interface {
	CreateJob(context.Context, *CreateJobRequest) (*Job, error)
	QueryJob(context.Context, *QueryJobRequest) (*JobSet, error)
	DescribeJob(context.Context, *DescribeJobRequest) (*Job, error)
	UpdateJob(context.Context, *UpdateJobRequest) (*Job, error)
	// 编辑job状态, 比如发布
	UpdateJobStatus(context.Context, *UpdateJobStatusRequest) (*Job, error)
	mustEmbedUnimplementedRPCServer()
}

// UnimplementedRPCServer must be embedded to have forward compatible implementations.
type UnimplementedRPCServer struct {
}

func (UnimplementedRPCServer) CreateJob(context.Context, *CreateJobRequest) (*Job, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateJob not implemented")
}
func (UnimplementedRPCServer) QueryJob(context.Context, *QueryJobRequest) (*JobSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryJob not implemented")
}
func (UnimplementedRPCServer) DescribeJob(context.Context, *DescribeJobRequest) (*Job, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeJob not implemented")
}
func (UnimplementedRPCServer) UpdateJob(context.Context, *UpdateJobRequest) (*Job, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateJob not implemented")
}
func (UnimplementedRPCServer) UpdateJobStatus(context.Context, *UpdateJobStatusRequest) (*Job, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateJobStatus not implemented")
}
func (UnimplementedRPCServer) mustEmbedUnimplementedRPCServer() {}

// UnsafeRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RPCServer will
// result in compilation errors.
type UnsafeRPCServer interface {
	mustEmbedUnimplementedRPCServer()
}

func RegisterRPCServer(s grpc.ServiceRegistrar, srv RPCServer) {
	s.RegisterService(&RPC_ServiceDesc, srv)
}

func _RPC_CreateJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).CreateJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_CreateJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).CreateJob(ctx, req.(*CreateJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_QueryJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).QueryJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_QueryJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).QueryJob(ctx, req.(*QueryJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_DescribeJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).DescribeJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_DescribeJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).DescribeJob(ctx, req.(*DescribeJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_UpdateJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).UpdateJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_UpdateJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).UpdateJob(ctx, req.(*UpdateJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_UpdateJobStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateJobStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).UpdateJobStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPC_UpdateJobStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).UpdateJobStatus(ctx, req.(*UpdateJobStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RPC_ServiceDesc is the grpc.ServiceDesc for RPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "infraboard.mpaas.job.RPC",
	HandlerType: (*RPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateJob",
			Handler:    _RPC_CreateJob_Handler,
		},
		{
			MethodName: "QueryJob",
			Handler:    _RPC_QueryJob_Handler,
		},
		{
			MethodName: "DescribeJob",
			Handler:    _RPC_DescribeJob_Handler,
		},
		{
			MethodName: "UpdateJob",
			Handler:    _RPC_UpdateJob_Handler,
		},
		{
			MethodName: "UpdateJobStatus",
			Handler:    _RPC_UpdateJobStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mpaas/apps/job/pb/rpc.proto",
}

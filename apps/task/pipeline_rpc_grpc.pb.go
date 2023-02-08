// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: apps/task/pb/pipeline_rpc.proto

package task

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

// PipelineRPCClient is the client API for PipelineRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PipelineRPCClient interface {
	// 执行Pipeline
	RunPipeline(ctx context.Context, in *RunPipelineRequest, opts ...grpc.CallOption) (*PipelineTask, error)
	// 查询Pipeline任务列表
	QueryPipelineTask(ctx context.Context, in *QueryPipelineTaskRequest, opts ...grpc.CallOption) (*PipelineTaskSet, error)
	// 查询Pipeline任务详情
	DescribePipelineTask(ctx context.Context, in *DescribePipelineTaskRequest, opts ...grpc.CallOption) (*PipelineTask, error)
}

type pipelineRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewPipelineRPCClient(cc grpc.ClientConnInterface) PipelineRPCClient {
	return &pipelineRPCClient{cc}
}

func (c *pipelineRPCClient) RunPipeline(ctx context.Context, in *RunPipelineRequest, opts ...grpc.CallOption) (*PipelineTask, error) {
	out := new(PipelineTask)
	err := c.cc.Invoke(ctx, "/infraboard.mpaas.task.PipelineRPC/RunPipeline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelineRPCClient) QueryPipelineTask(ctx context.Context, in *QueryPipelineTaskRequest, opts ...grpc.CallOption) (*PipelineTaskSet, error) {
	out := new(PipelineTaskSet)
	err := c.cc.Invoke(ctx, "/infraboard.mpaas.task.PipelineRPC/QueryPipelineTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelineRPCClient) DescribePipelineTask(ctx context.Context, in *DescribePipelineTaskRequest, opts ...grpc.CallOption) (*PipelineTask, error) {
	out := new(PipelineTask)
	err := c.cc.Invoke(ctx, "/infraboard.mpaas.task.PipelineRPC/DescribePipelineTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PipelineRPCServer is the server API for PipelineRPC service.
// All implementations must embed UnimplementedPipelineRPCServer
// for forward compatibility
type PipelineRPCServer interface {
	// 执行Pipeline
	RunPipeline(context.Context, *RunPipelineRequest) (*PipelineTask, error)
	// 查询Pipeline任务列表
	QueryPipelineTask(context.Context, *QueryPipelineTaskRequest) (*PipelineTaskSet, error)
	// 查询Pipeline任务详情
	DescribePipelineTask(context.Context, *DescribePipelineTaskRequest) (*PipelineTask, error)
	mustEmbedUnimplementedPipelineRPCServer()
}

// UnimplementedPipelineRPCServer must be embedded to have forward compatible implementations.
type UnimplementedPipelineRPCServer struct {
}

func (UnimplementedPipelineRPCServer) RunPipeline(context.Context, *RunPipelineRequest) (*PipelineTask, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunPipeline not implemented")
}
func (UnimplementedPipelineRPCServer) QueryPipelineTask(context.Context, *QueryPipelineTaskRequest) (*PipelineTaskSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryPipelineTask not implemented")
}
func (UnimplementedPipelineRPCServer) DescribePipelineTask(context.Context, *DescribePipelineTaskRequest) (*PipelineTask, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribePipelineTask not implemented")
}
func (UnimplementedPipelineRPCServer) mustEmbedUnimplementedPipelineRPCServer() {}

// UnsafePipelineRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PipelineRPCServer will
// result in compilation errors.
type UnsafePipelineRPCServer interface {
	mustEmbedUnimplementedPipelineRPCServer()
}

func RegisterPipelineRPCServer(s grpc.ServiceRegistrar, srv PipelineRPCServer) {
	s.RegisterService(&PipelineRPC_ServiceDesc, srv)
}

func _PipelineRPC_RunPipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunPipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelineRPCServer).RunPipeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.mpaas.task.PipelineRPC/RunPipeline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelineRPCServer).RunPipeline(ctx, req.(*RunPipelineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PipelineRPC_QueryPipelineTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPipelineTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelineRPCServer).QueryPipelineTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.mpaas.task.PipelineRPC/QueryPipelineTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelineRPCServer).QueryPipelineTask(ctx, req.(*QueryPipelineTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PipelineRPC_DescribePipelineTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribePipelineTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelineRPCServer).DescribePipelineTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.mpaas.task.PipelineRPC/DescribePipelineTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelineRPCServer).DescribePipelineTask(ctx, req.(*DescribePipelineTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PipelineRPC_ServiceDesc is the grpc.ServiceDesc for PipelineRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PipelineRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "infraboard.mpaas.task.PipelineRPC",
	HandlerType: (*PipelineRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RunPipeline",
			Handler:    _PipelineRPC_RunPipeline_Handler,
		},
		{
			MethodName: "QueryPipelineTask",
			Handler:    _PipelineRPC_QueryPipelineTask_Handler,
		},
		{
			MethodName: "DescribePipelineTask",
			Handler:    _PipelineRPC_DescribePipelineTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apps/task/pb/pipeline_rpc.proto",
}

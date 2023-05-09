// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.6
// source: mpaas/apps/pipeline/pb/rpc.proto

package pipeline

import (
	request "github.com/infraboard/mcube/http/request"
	request1 "github.com/infraboard/mcube/pb/request"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type QueryPipelineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 分页请求
	// @gotags: json:"page"
	Page *request.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page"`
	// pipeline 所属域
	// @gotags: json:"domain"
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain"`
	// pipeline 所属空间
	// @gotags: json:"namespace"
	Namespace string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace"`
	// 是否是模版, 用于快速继承模版参数进行修改, 模版不能用于执行
	// @gotags: json:"is_template"
	IsTemplate *bool `protobuf:"varint,4,opt,name=is_template,json=isTemplate,proto3,oneof" json:"is_template"`
	// pipeline Id列表
	// @gotags: json:"ids"
	Ids []string `protobuf:"bytes,15,rep,name=ids,proto3" json:"ids"`
}

func (x *QueryPipelineRequest) Reset() {
	*x = QueryPipelineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mpaas_apps_pipeline_pb_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryPipelineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryPipelineRequest) ProtoMessage() {}

func (x *QueryPipelineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mpaas_apps_pipeline_pb_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryPipelineRequest.ProtoReflect.Descriptor instead.
func (*QueryPipelineRequest) Descriptor() ([]byte, []int) {
	return file_mpaas_apps_pipeline_pb_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *QueryPipelineRequest) GetPage() *request.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *QueryPipelineRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *QueryPipelineRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *QueryPipelineRequest) GetIsTemplate() bool {
	if x != nil && x.IsTemplate != nil {
		return *x.IsTemplate
	}
	return false
}

func (x *QueryPipelineRequest) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type DescribePipelineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// pipeline id
	// @gotags: json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
}

func (x *DescribePipelineRequest) Reset() {
	*x = DescribePipelineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mpaas_apps_pipeline_pb_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribePipelineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribePipelineRequest) ProtoMessage() {}

func (x *DescribePipelineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mpaas_apps_pipeline_pb_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribePipelineRequest.ProtoReflect.Descriptor instead.
func (*DescribePipelineRequest) Descriptor() ([]byte, []int) {
	return file_mpaas_apps_pipeline_pb_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *DescribePipelineRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdatePipelineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// pipeline id
	// @gotags: json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	// 更新模式
	// @gotags: json:"update_mode"
	UpdateMode request1.UpdateMode `protobuf:"varint,2,opt,name=update_mode,json=updateMode,proto3,enum=infraboard.mcube.request.UpdateMode" json:"update_mode"`
	// 更新人
	// @gotags: json:"update_by"
	UpdateBy string `protobuf:"bytes,3,opt,name=update_by,json=updateBy,proto3" json:"update_by"`
	// 更新时间
	// @gotags: json:"update_at"
	UpdateAt int64 `protobuf:"varint,4,opt,name=update_at,json=updateAt,proto3" json:"update_at"`
	// 定义
	// @gotags: json:"spec"
	Spec *CreatePipelineRequest `protobuf:"bytes,5,opt,name=spec,proto3" json:"spec"`
}

func (x *UpdatePipelineRequest) Reset() {
	*x = UpdatePipelineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mpaas_apps_pipeline_pb_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePipelineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePipelineRequest) ProtoMessage() {}

func (x *UpdatePipelineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mpaas_apps_pipeline_pb_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePipelineRequest.ProtoReflect.Descriptor instead.
func (*UpdatePipelineRequest) Descriptor() ([]byte, []int) {
	return file_mpaas_apps_pipeline_pb_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *UpdatePipelineRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdatePipelineRequest) GetUpdateMode() request1.UpdateMode {
	if x != nil {
		return x.UpdateMode
	}
	return request1.UpdateMode(0)
}

func (x *UpdatePipelineRequest) GetUpdateBy() string {
	if x != nil {
		return x.UpdateBy
	}
	return ""
}

func (x *UpdatePipelineRequest) GetUpdateAt() int64 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

func (x *UpdatePipelineRequest) GetSpec() *CreatePipelineRequest {
	if x != nil {
		return x.Spec
	}
	return nil
}

type DeletePipelineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// pipeline id
	// @gotags: json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
}

func (x *DeletePipelineRequest) Reset() {
	*x = DeletePipelineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mpaas_apps_pipeline_pb_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePipelineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePipelineRequest) ProtoMessage() {}

func (x *DeletePipelineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mpaas_apps_pipeline_pb_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePipelineRequest.ProtoReflect.Descriptor instead.
func (*DeletePipelineRequest) Descriptor() ([]byte, []int) {
	return file_mpaas_apps_pipeline_pb_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *DeletePipelineRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_mpaas_apps_pipeline_pb_rpc_proto protoreflect.FileDescriptor

var file_mpaas_apps_pipeline_pb_rpc_proto_rawDesc = []byte{
	0x0a, 0x20, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x70, 0x69, 0x70,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x19, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d,
	0x70, 0x61, 0x61, 0x73, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x1a, 0x25, 0x6d,
	0x70, 0x61, 0x61, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x70,
	0x61, 0x67, 0x65, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcc,
	0x01, 0x0a, 0x14, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x0a, 0x69, 0x73,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x12, 0x10, 0x0a, 0x03, 0x69,
	0x64, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73, 0x42, 0x0e, 0x0a,
	0x0c, 0x5f, 0x69, 0x73, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x22, 0x29, 0x0a,
	0x17, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xee, 0x01, 0x0a, 0x15, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x45, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x62, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x41, 0x74, 0x12, 0x44, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x30, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d,
	0x70, 0x61, 0x61, 0x73, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x22, 0x27, 0x0a, 0x15, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x32, 0x97, 0x04, 0x0a, 0x03, 0x52, 0x50, 0x43, 0x12, 0x68, 0x0a, 0x0d, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x2f, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x70,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x69, 0x70,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e,
	0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x53, 0x65, 0x74, 0x12, 0x6b, 0x0a, 0x10, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x32, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x70, 0x69, 0x70, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x69, 0x70,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e,
	0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x12, 0x67, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x12, 0x30, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x67, 0x0a, 0x0e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x30, 0x2e, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e,
	0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23,
	0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61,
	0x73, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x12, 0x67, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x69, 0x70,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x30, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x42, 0x2b, 0x5a, 0x29,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x73,
	0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_mpaas_apps_pipeline_pb_rpc_proto_rawDescOnce sync.Once
	file_mpaas_apps_pipeline_pb_rpc_proto_rawDescData = file_mpaas_apps_pipeline_pb_rpc_proto_rawDesc
)

func file_mpaas_apps_pipeline_pb_rpc_proto_rawDescGZIP() []byte {
	file_mpaas_apps_pipeline_pb_rpc_proto_rawDescOnce.Do(func() {
		file_mpaas_apps_pipeline_pb_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_mpaas_apps_pipeline_pb_rpc_proto_rawDescData)
	})
	return file_mpaas_apps_pipeline_pb_rpc_proto_rawDescData
}

var file_mpaas_apps_pipeline_pb_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_mpaas_apps_pipeline_pb_rpc_proto_goTypes = []interface{}{
	(*QueryPipelineRequest)(nil),    // 0: infraboard.mpaas.pipeline.QueryPipelineRequest
	(*DescribePipelineRequest)(nil), // 1: infraboard.mpaas.pipeline.DescribePipelineRequest
	(*UpdatePipelineRequest)(nil),   // 2: infraboard.mpaas.pipeline.UpdatePipelineRequest
	(*DeletePipelineRequest)(nil),   // 3: infraboard.mpaas.pipeline.DeletePipelineRequest
	(*request.PageRequest)(nil),     // 4: infraboard.mcube.page.PageRequest
	(request1.UpdateMode)(0),        // 5: infraboard.mcube.request.UpdateMode
	(*CreatePipelineRequest)(nil),   // 6: infraboard.mpaas.pipeline.CreatePipelineRequest
	(*PipelineSet)(nil),             // 7: infraboard.mpaas.pipeline.PipelineSet
	(*Pipeline)(nil),                // 8: infraboard.mpaas.pipeline.Pipeline
}
var file_mpaas_apps_pipeline_pb_rpc_proto_depIdxs = []int32{
	4, // 0: infraboard.mpaas.pipeline.QueryPipelineRequest.page:type_name -> infraboard.mcube.page.PageRequest
	5, // 1: infraboard.mpaas.pipeline.UpdatePipelineRequest.update_mode:type_name -> infraboard.mcube.request.UpdateMode
	6, // 2: infraboard.mpaas.pipeline.UpdatePipelineRequest.spec:type_name -> infraboard.mpaas.pipeline.CreatePipelineRequest
	0, // 3: infraboard.mpaas.pipeline.RPC.QueryPipeline:input_type -> infraboard.mpaas.pipeline.QueryPipelineRequest
	1, // 4: infraboard.mpaas.pipeline.RPC.DescribePipeline:input_type -> infraboard.mpaas.pipeline.DescribePipelineRequest
	6, // 5: infraboard.mpaas.pipeline.RPC.CreatePipeline:input_type -> infraboard.mpaas.pipeline.CreatePipelineRequest
	2, // 6: infraboard.mpaas.pipeline.RPC.UpdatePipeline:input_type -> infraboard.mpaas.pipeline.UpdatePipelineRequest
	3, // 7: infraboard.mpaas.pipeline.RPC.DeletePipeline:input_type -> infraboard.mpaas.pipeline.DeletePipelineRequest
	7, // 8: infraboard.mpaas.pipeline.RPC.QueryPipeline:output_type -> infraboard.mpaas.pipeline.PipelineSet
	8, // 9: infraboard.mpaas.pipeline.RPC.DescribePipeline:output_type -> infraboard.mpaas.pipeline.Pipeline
	8, // 10: infraboard.mpaas.pipeline.RPC.CreatePipeline:output_type -> infraboard.mpaas.pipeline.Pipeline
	8, // 11: infraboard.mpaas.pipeline.RPC.UpdatePipeline:output_type -> infraboard.mpaas.pipeline.Pipeline
	8, // 12: infraboard.mpaas.pipeline.RPC.DeletePipeline:output_type -> infraboard.mpaas.pipeline.Pipeline
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_mpaas_apps_pipeline_pb_rpc_proto_init() }
func file_mpaas_apps_pipeline_pb_rpc_proto_init() {
	if File_mpaas_apps_pipeline_pb_rpc_proto != nil {
		return
	}
	file_mpaas_apps_pipeline_pb_pipeline_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_mpaas_apps_pipeline_pb_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryPipelineRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mpaas_apps_pipeline_pb_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribePipelineRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mpaas_apps_pipeline_pb_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePipelineRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mpaas_apps_pipeline_pb_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePipelineRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_mpaas_apps_pipeline_pb_rpc_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mpaas_apps_pipeline_pb_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mpaas_apps_pipeline_pb_rpc_proto_goTypes,
		DependencyIndexes: file_mpaas_apps_pipeline_pb_rpc_proto_depIdxs,
		MessageInfos:      file_mpaas_apps_pipeline_pb_rpc_proto_msgTypes,
	}.Build()
	File_mpaas_apps_pipeline_pb_rpc_proto = out.File
	file_mpaas_apps_pipeline_pb_rpc_proto_rawDesc = nil
	file_mpaas_apps_pipeline_pb_rpc_proto_goTypes = nil
	file_mpaas_apps_pipeline_pb_rpc_proto_depIdxs = nil
}

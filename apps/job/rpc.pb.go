// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.6
// source: mpaas/apps/job/pb/rpc.proto

package job

import (
	request "github.com/infraboard/mcube/http/request"
	request1 "github.com/infraboard/mcube/pb/request"
	resource "github.com/infraboard/mcube/pb/resource"
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

type DESCRIBE_BY int32

const (
	// Job的ID
	DESCRIBE_BY_JOB_ID DESCRIBE_BY = 0
	// Job的唯一名称, <job_name>.<namespace>.<domain>
	DESCRIBE_BY_JOB_UNIQ_NAME DESCRIBE_BY = 1
)

// Enum value maps for DESCRIBE_BY.
var (
	DESCRIBE_BY_name = map[int32]string{
		0: "JOB_ID",
		1: "JOB_UNIQ_NAME",
	}
	DESCRIBE_BY_value = map[string]int32{
		"JOB_ID":        0,
		"JOB_UNIQ_NAME": 1,
	}
)

func (x DESCRIBE_BY) Enum() *DESCRIBE_BY {
	p := new(DESCRIBE_BY)
	*p = x
	return p
}

func (x DESCRIBE_BY) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DESCRIBE_BY) Descriptor() protoreflect.EnumDescriptor {
	return file_mpaas_apps_job_pb_rpc_proto_enumTypes[0].Descriptor()
}

func (DESCRIBE_BY) Type() protoreflect.EnumType {
	return &file_mpaas_apps_job_pb_rpc_proto_enumTypes[0]
}

func (x DESCRIBE_BY) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DESCRIBE_BY.Descriptor instead.
func (DESCRIBE_BY) EnumDescriptor() ([]byte, []int) {
	return file_mpaas_apps_job_pb_rpc_proto_rawDescGZIP(), []int{0}
}

type QueryJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 资源范围
	// @gotags: json:"mscopeeta"
	Scope *resource.Scope `protobuf:"bytes,1,opt,name=scope,proto3" json:"mscopeeta"`
	// 资源标签过滤
	// @gotags: json:"filters"
	Filters []*resource.LabelRequirement `protobuf:"bytes,2,rep,name=filters,proto3" json:"filters"`
	// 分页请求
	// @gotags: json:"page"
	Page *request.PageRequest `protobuf:"bytes,3,opt,name=page,proto3" json:"page"`
	// 是否是公开Job, 默认只能本空间内访问
	// @gotags: bson:"visiable_mode" json:"visiable_mode"
	VisiableMode *resource.VISIABLE `protobuf:"varint,4,opt,name=visiable_mode,json=visiableMode,proto3,enum=infraboard.mcube.resource.VISIABLE,oneof" json:"visiable_mode" bson:"visiable_mode"`
	// job Id列表
	// @gotags: json:"ids"
	Ids []string `protobuf:"bytes,5,rep,name=ids,proto3" json:"ids"`
	// job 名称列表
	// @gotags: json:"names"
	Names []string `protobuf:"bytes,6,rep,name=names,proto3" json:"names"`
	// job 标签
	// @gotags: json:"label"
	Label map[string]string `protobuf:"bytes,7,rep,name=label,proto3" json:"label" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *QueryJobRequest) Reset() {
	*x = QueryJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mpaas_apps_job_pb_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryJobRequest) ProtoMessage() {}

func (x *QueryJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mpaas_apps_job_pb_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryJobRequest.ProtoReflect.Descriptor instead.
func (*QueryJobRequest) Descriptor() ([]byte, []int) {
	return file_mpaas_apps_job_pb_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *QueryJobRequest) GetScope() *resource.Scope {
	if x != nil {
		return x.Scope
	}
	return nil
}

func (x *QueryJobRequest) GetFilters() []*resource.LabelRequirement {
	if x != nil {
		return x.Filters
	}
	return nil
}

func (x *QueryJobRequest) GetPage() *request.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *QueryJobRequest) GetVisiableMode() resource.VISIABLE {
	if x != nil && x.VisiableMode != nil {
		return *x.VisiableMode
	}
	return resource.VISIABLE(0)
}

func (x *QueryJobRequest) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *QueryJobRequest) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

func (x *QueryJobRequest) GetLabel() map[string]string {
	if x != nil {
		return x.Label
	}
	return nil
}

type DescribeJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 查询方式
	// @gotags: json:"describe_by"
	DescribeBy DESCRIBE_BY `protobuf:"varint,1,opt,name=describe_by,json=describeBy,proto3,enum=infraboard.mpaas.job.DESCRIBE_BY" json:"describe_by"`
	// 查询值
	// @gotags: json:"describe_value"  validate:"required"
	DescribeValue string `protobuf:"bytes,2,opt,name=describe_value,json=describeValue,proto3" json:"describe_value" validate:"required"`
}

func (x *DescribeJobRequest) Reset() {
	*x = DescribeJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mpaas_apps_job_pb_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeJobRequest) ProtoMessage() {}

func (x *DescribeJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mpaas_apps_job_pb_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeJobRequest.ProtoReflect.Descriptor instead.
func (*DescribeJobRequest) Descriptor() ([]byte, []int) {
	return file_mpaas_apps_job_pb_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *DescribeJobRequest) GetDescribeBy() DESCRIBE_BY {
	if x != nil {
		return x.DescribeBy
	}
	return DESCRIBE_BY_JOB_ID
}

func (x *DescribeJobRequest) GetDescribeValue() string {
	if x != nil {
		return x.DescribeValue
	}
	return ""
}

type UpdateJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Cluster id
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
	// 更新的信息
	// @gotags: json:"spec"
	Spec *CreateJobRequest `protobuf:"bytes,5,opt,name=spec,proto3" json:"spec"`
}

func (x *UpdateJobRequest) Reset() {
	*x = UpdateJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mpaas_apps_job_pb_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateJobRequest) ProtoMessage() {}

func (x *UpdateJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mpaas_apps_job_pb_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateJobRequest.ProtoReflect.Descriptor instead.
func (*UpdateJobRequest) Descriptor() ([]byte, []int) {
	return file_mpaas_apps_job_pb_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateJobRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateJobRequest) GetUpdateMode() request1.UpdateMode {
	if x != nil {
		return x.UpdateMode
	}
	return request1.UpdateMode(0)
}

func (x *UpdateJobRequest) GetUpdateBy() string {
	if x != nil {
		return x.UpdateBy
	}
	return ""
}

func (x *UpdateJobRequest) GetUpdateAt() int64 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

func (x *UpdateJobRequest) GetSpec() *CreateJobRequest {
	if x != nil {
		return x.Spec
	}
	return nil
}

type UpdateJobStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// job id
	// @gotags: json:"id" validate:"required"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" validate:"required"`
	// job状态, 注意 Job是带有版本管理的
	// @gotags: bson:"status" json:"status"
	Status *JobStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status" bson:"status"`
}

func (x *UpdateJobStatusRequest) Reset() {
	*x = UpdateJobStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mpaas_apps_job_pb_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateJobStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateJobStatusRequest) ProtoMessage() {}

func (x *UpdateJobStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mpaas_apps_job_pb_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateJobStatusRequest.ProtoReflect.Descriptor instead.
func (*UpdateJobStatusRequest) Descriptor() ([]byte, []int) {
	return file_mpaas_apps_job_pb_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateJobStatusRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateJobStatusRequest) GetStatus() *JobStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

type DeleteJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// job id
	// @gotags: json:"id" validate:"required"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" validate:"required"`
}

func (x *DeleteJobRequest) Reset() {
	*x = DeleteJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mpaas_apps_job_pb_rpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteJobRequest) ProtoMessage() {}

func (x *DeleteJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mpaas_apps_job_pb_rpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteJobRequest.ProtoReflect.Descriptor instead.
func (*DeleteJobRequest) Descriptor() ([]byte, []int) {
	return file_mpaas_apps_job_pb_rpc_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteJobRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_mpaas_apps_job_pb_rpc_proto protoreflect.FileDescriptor

var file_mpaas_apps_job_pb_rpc_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x6a, 0x6f, 0x62,
	0x2f, 0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e,
	0x6a, 0x6f, 0x62, 0x1a, 0x1b, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f,
	0x6a, 0x6f, 0x62, 0x2f, 0x70, 0x62, 0x2f, 0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x18, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f,
	0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x6d, 0x63, 0x75, 0x62,
	0x65, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x6d, 0x63, 0x75, 0x62, 0x65,
	0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2f, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x6d, 0x63, 0x75, 0x62, 0x65,
	0x2f, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x6d, 0x65, 0x74,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x70,
	0x62, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x76, 0x69, 0x73, 0x69, 0x61,
	0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd3, 0x03, 0x0a, 0x0f, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a,
	0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52, 0x05,
	0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x45, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x36, 0x0a, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x70, 0x61,
	0x67, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x12, 0x4d, 0x0a, 0x0d, 0x76, 0x69, 0x73, 0x69, 0x61, 0x62, 0x6c, 0x65,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x56, 0x49, 0x53, 0x49, 0x41, 0x42, 0x4c, 0x45,
	0x48, 0x00, 0x52, 0x0c, 0x76, 0x69, 0x73, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x4d, 0x6f, 0x64, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x46, 0x0a, 0x05, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x6a, 0x6f,
	0x62, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x1a, 0x38, 0x0a, 0x0a, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x10, 0x0a,
	0x0e, 0x5f, 0x76, 0x69, 0x73, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x22,
	0x7f, 0x0a, 0x12, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x42, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x5f, 0x62, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x6a, 0x6f,
	0x62, 0x2e, 0x44, 0x45, 0x53, 0x43, 0x52, 0x49, 0x42, 0x45, 0x5f, 0x42, 0x59, 0x52, 0x0a, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x42, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0xdf, 0x01, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x45, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x6d, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65,
	0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x62, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x3a, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x73, 0x70,
	0x65, 0x63, 0x22, 0x61, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x37, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e,
	0x6a, 0x6f, 0x62, 0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x22, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4a,
	0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x2a, 0x2c, 0x0a, 0x0b, 0x44, 0x45, 0x53,
	0x43, 0x52, 0x49, 0x42, 0x45, 0x5f, 0x42, 0x59, 0x12, 0x0a, 0x0a, 0x06, 0x4a, 0x4f, 0x42, 0x5f,
	0x49, 0x44, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x4a, 0x4f, 0x42, 0x5f, 0x55, 0x4e, 0x49, 0x51,
	0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x01, 0x32, 0xa6, 0x03, 0x0a, 0x03, 0x52, 0x50, 0x43, 0x12,
	0x4e, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x12, 0x26, 0x2e, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e,
	0x6a, 0x6f, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x4a, 0x6f, 0x62, 0x12,
	0x4f, 0x0a, 0x08, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4a, 0x6f, 0x62, 0x12, 0x25, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x6a,
	0x6f, 0x62, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x65, 0x74,
	0x12, 0x52, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4a, 0x6f, 0x62, 0x12,
	0x28, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61,
	0x61, 0x73, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4a,
	0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x6a, 0x6f, 0x62,
	0x2e, 0x4a, 0x6f, 0x62, 0x12, 0x4e, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f,
	0x62, 0x12, 0x26, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d,
	0x70, 0x61, 0x61, 0x73, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a,
	0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x6a, 0x6f, 0x62,
	0x2e, 0x4a, 0x6f, 0x62, 0x12, 0x5a, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f,
	0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2c, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x4a, 0x6f, 0x62,
	0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2f,
	0x61, 0x70, 0x70, 0x73, 0x2f, 0x6a, 0x6f, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mpaas_apps_job_pb_rpc_proto_rawDescOnce sync.Once
	file_mpaas_apps_job_pb_rpc_proto_rawDescData = file_mpaas_apps_job_pb_rpc_proto_rawDesc
)

func file_mpaas_apps_job_pb_rpc_proto_rawDescGZIP() []byte {
	file_mpaas_apps_job_pb_rpc_proto_rawDescOnce.Do(func() {
		file_mpaas_apps_job_pb_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_mpaas_apps_job_pb_rpc_proto_rawDescData)
	})
	return file_mpaas_apps_job_pb_rpc_proto_rawDescData
}

var file_mpaas_apps_job_pb_rpc_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mpaas_apps_job_pb_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_mpaas_apps_job_pb_rpc_proto_goTypes = []interface{}{
	(DESCRIBE_BY)(0),                  // 0: infraboard.mpaas.job.DESCRIBE_BY
	(*QueryJobRequest)(nil),           // 1: infraboard.mpaas.job.QueryJobRequest
	(*DescribeJobRequest)(nil),        // 2: infraboard.mpaas.job.DescribeJobRequest
	(*UpdateJobRequest)(nil),          // 3: infraboard.mpaas.job.UpdateJobRequest
	(*UpdateJobStatusRequest)(nil),    // 4: infraboard.mpaas.job.UpdateJobStatusRequest
	(*DeleteJobRequest)(nil),          // 5: infraboard.mpaas.job.DeleteJobRequest
	nil,                               // 6: infraboard.mpaas.job.QueryJobRequest.LabelEntry
	(*resource.Scope)(nil),            // 7: infraboard.mcube.resource.Scope
	(*resource.LabelRequirement)(nil), // 8: infraboard.mcube.resource.LabelRequirement
	(*request.PageRequest)(nil),       // 9: infraboard.mcube.page.PageRequest
	(resource.VISIABLE)(0),            // 10: infraboard.mcube.resource.VISIABLE
	(request1.UpdateMode)(0),          // 11: infraboard.mcube.request.UpdateMode
	(*CreateJobRequest)(nil),          // 12: infraboard.mpaas.job.CreateJobRequest
	(*JobStatus)(nil),                 // 13: infraboard.mpaas.job.JobStatus
	(*Job)(nil),                       // 14: infraboard.mpaas.job.Job
	(*JobSet)(nil),                    // 15: infraboard.mpaas.job.JobSet
}
var file_mpaas_apps_job_pb_rpc_proto_depIdxs = []int32{
	7,  // 0: infraboard.mpaas.job.QueryJobRequest.scope:type_name -> infraboard.mcube.resource.Scope
	8,  // 1: infraboard.mpaas.job.QueryJobRequest.filters:type_name -> infraboard.mcube.resource.LabelRequirement
	9,  // 2: infraboard.mpaas.job.QueryJobRequest.page:type_name -> infraboard.mcube.page.PageRequest
	10, // 3: infraboard.mpaas.job.QueryJobRequest.visiable_mode:type_name -> infraboard.mcube.resource.VISIABLE
	6,  // 4: infraboard.mpaas.job.QueryJobRequest.label:type_name -> infraboard.mpaas.job.QueryJobRequest.LabelEntry
	0,  // 5: infraboard.mpaas.job.DescribeJobRequest.describe_by:type_name -> infraboard.mpaas.job.DESCRIBE_BY
	11, // 6: infraboard.mpaas.job.UpdateJobRequest.update_mode:type_name -> infraboard.mcube.request.UpdateMode
	12, // 7: infraboard.mpaas.job.UpdateJobRequest.spec:type_name -> infraboard.mpaas.job.CreateJobRequest
	13, // 8: infraboard.mpaas.job.UpdateJobStatusRequest.status:type_name -> infraboard.mpaas.job.JobStatus
	12, // 9: infraboard.mpaas.job.RPC.CreateJob:input_type -> infraboard.mpaas.job.CreateJobRequest
	1,  // 10: infraboard.mpaas.job.RPC.QueryJob:input_type -> infraboard.mpaas.job.QueryJobRequest
	2,  // 11: infraboard.mpaas.job.RPC.DescribeJob:input_type -> infraboard.mpaas.job.DescribeJobRequest
	3,  // 12: infraboard.mpaas.job.RPC.UpdateJob:input_type -> infraboard.mpaas.job.UpdateJobRequest
	4,  // 13: infraboard.mpaas.job.RPC.UpdateJobStatus:input_type -> infraboard.mpaas.job.UpdateJobStatusRequest
	14, // 14: infraboard.mpaas.job.RPC.CreateJob:output_type -> infraboard.mpaas.job.Job
	15, // 15: infraboard.mpaas.job.RPC.QueryJob:output_type -> infraboard.mpaas.job.JobSet
	14, // 16: infraboard.mpaas.job.RPC.DescribeJob:output_type -> infraboard.mpaas.job.Job
	14, // 17: infraboard.mpaas.job.RPC.UpdateJob:output_type -> infraboard.mpaas.job.Job
	14, // 18: infraboard.mpaas.job.RPC.UpdateJobStatus:output_type -> infraboard.mpaas.job.Job
	14, // [14:19] is the sub-list for method output_type
	9,  // [9:14] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_mpaas_apps_job_pb_rpc_proto_init() }
func file_mpaas_apps_job_pb_rpc_proto_init() {
	if File_mpaas_apps_job_pb_rpc_proto != nil {
		return
	}
	file_mpaas_apps_job_pb_job_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_mpaas_apps_job_pb_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryJobRequest); i {
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
		file_mpaas_apps_job_pb_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeJobRequest); i {
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
		file_mpaas_apps_job_pb_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateJobRequest); i {
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
		file_mpaas_apps_job_pb_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateJobStatusRequest); i {
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
		file_mpaas_apps_job_pb_rpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteJobRequest); i {
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
	file_mpaas_apps_job_pb_rpc_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mpaas_apps_job_pb_rpc_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mpaas_apps_job_pb_rpc_proto_goTypes,
		DependencyIndexes: file_mpaas_apps_job_pb_rpc_proto_depIdxs,
		EnumInfos:         file_mpaas_apps_job_pb_rpc_proto_enumTypes,
		MessageInfos:      file_mpaas_apps_job_pb_rpc_proto_msgTypes,
	}.Build()
	File_mpaas_apps_job_pb_rpc_proto = out.File
	file_mpaas_apps_job_pb_rpc_proto_rawDesc = nil
	file_mpaas_apps_job_pb_rpc_proto_goTypes = nil
	file_mpaas_apps_job_pb_rpc_proto_depIdxs = nil
}

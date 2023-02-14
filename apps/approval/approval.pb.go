// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: apps/approval/pb/approval.proto

package approval

import (
	pipeline "github.com/infraboard/mpaas/apps/pipeline"
	meta "github.com/infraboard/mpaas/common/meta"
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

type STAGE int32

const (
	// 草稿
	STAGE_DRAFT STAGE = 0
	// 待审核
	STAGE_PENDDING STAGE = 1
	// 已过期, 如果一直没审核, 多久后会过期
	STAGE_EXPIRED STAGE = 2
	// 审核拒绝
	STAGE_DENY STAGE = 3
	// 审核通过
	STAGE_PASSED STAGE = 4
	// 发布成功
	STAGE_SUCCEEDED STAGE = 11
	// 发布失败
	STAGE_FAILED STAGE = 12
	// 发布关闭, 发布成功后,验证没问题, 发布结束
	STAGE_CLOSED STAGE = 13
)

// Enum value maps for STAGE.
var (
	STAGE_name = map[int32]string{
		0:  "DRAFT",
		1:  "PENDDING",
		2:  "EXPIRED",
		3:  "DENY",
		4:  "PASSED",
		11: "SUCCEEDED",
		12: "FAILED",
		13: "CLOSED",
	}
	STAGE_value = map[string]int32{
		"DRAFT":     0,
		"PENDDING":  1,
		"EXPIRED":   2,
		"DENY":      3,
		"PASSED":    4,
		"SUCCEEDED": 11,
		"FAILED":    12,
		"CLOSED":    13,
	}
)

func (x STAGE) Enum() *STAGE {
	p := new(STAGE)
	*p = x
	return p
}

func (x STAGE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (STAGE) Descriptor() protoreflect.EnumDescriptor {
	return file_apps_approval_pb_approval_proto_enumTypes[0].Descriptor()
}

func (STAGE) Type() protoreflect.EnumType {
	return &file_apps_approval_pb_approval_proto_enumTypes[0]
}

func (x STAGE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use STAGE.Descriptor instead.
func (STAGE) EnumDescriptor() ([]byte, []int) {
	return file_apps_approval_pb_approval_proto_rawDescGZIP(), []int{0}
}

type ApprovalSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 总数
	// @gotags: bson:"total" json:"total"
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total" bson:"total"`
	// 列表
	// @gotags: bson:"items" json:"items"
	Items []*Approval `protobuf:"bytes,2,rep,name=items,proto3" json:"items" bson:"items"`
}

func (x *ApprovalSet) Reset() {
	*x = ApprovalSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_approval_pb_approval_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApprovalSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApprovalSet) ProtoMessage() {}

func (x *ApprovalSet) ProtoReflect() protoreflect.Message {
	mi := &file_apps_approval_pb_approval_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApprovalSet.ProtoReflect.Descriptor instead.
func (*ApprovalSet) Descriptor() ([]byte, []int) {
	return file_apps_approval_pb_approval_proto_rawDescGZIP(), []int{0}
}

func (x *ApprovalSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ApprovalSet) GetItems() []*Approval {
	if x != nil {
		return x.Items
	}
	return nil
}

// 发布申请单
type Approval struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 元信息
	// @gotags: bson:",inline" json:"meta"
	Meta *meta.Meta `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta" bson:",inline"`
	// 创建信息
	// @gotags: bson:",inline" json:"spec"
	Spec *CreateApprovalRequest `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec" bson:",inline"`
	// 部署流水线配置
	// @gotags: bson:"-" json:"deploy_pipeline"
	DeployPipeline *pipeline.Pipeline `protobuf:"bytes,7,opt,name=deploy_pipeline,json=deployPipeline,proto3" json:"deploy_pipeline" bson:"-"`
	// 发布单当前状态
	// @gotags: bson:"status" json:"status"
	Status *Status `protobuf:"bytes,3,opt,name=status,proto3" json:"status" bson:"status"`
}

func (x *Approval) Reset() {
	*x = Approval{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_approval_pb_approval_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Approval) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Approval) ProtoMessage() {}

func (x *Approval) ProtoReflect() protoreflect.Message {
	mi := &file_apps_approval_pb_approval_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Approval.ProtoReflect.Descriptor instead.
func (*Approval) Descriptor() ([]byte, []int) {
	return file_apps_approval_pb_approval_proto_rawDescGZIP(), []int{1}
}

func (x *Approval) GetMeta() *meta.Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Approval) GetSpec() *CreateApprovalRequest {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *Approval) GetDeployPipeline() *pipeline.Pipeline {
	if x != nil {
		return x.DeployPipeline
	}
	return nil
}

func (x *Approval) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

// 创建发布申请
type CreateApprovalRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 是否是模版
	// @gotags: bson:",inline" json:"meta"
	IsTemplate bool `protobuf:"varint,1,opt,name=is_template,json=isTemplate,proto3" json:"meta" bson:",inline"`
	// 申请人列表, 申请通过后, 由申请人执行发布
	// @gotags: bson:"proposers" json:"proposers"
	Proposers []string `protobuf:"bytes,2,rep,name=proposers,proto3" json:"proposers" bson:"proposers"`
	// 审核人列表
	// @gotags: bson:"auditors" json:"auditors"
	Auditors []string `protobuf:"bytes,3,rep,name=auditors,proto3" json:"auditors" bson:"auditors"`
	// 发布号
	// @gotags: bson:"version" json:"version" validate:"required"
	Version string `protobuf:"bytes,4,opt,name=version,proto3" json:"version" bson:"version" validate:"required"`
	// 发布说明
	// @gotags: bson:"describe" json:"describe" validate:"required"
	Describe string `protobuf:"bytes,5,opt,name=describe,proto3" json:"describe" bson:"describe" validate:"required"`
	// 审核通过后, 是否自动执行发布
	// @gotags: bson:"auto_publish" json:"auto_publish"
	AutoPublish bool `protobuf:"varint,6,opt,name=auto_publish,json=autoPublish,proto3" json:"auto_publish" bson:"auto_publish"`
	// 部署流水线配置
	// @gotags: bson:"-" json:"deploy_pipeline_spec"
	DeployPipelineSpec *pipeline.CreatePipelineRequest `protobuf:"bytes,7,opt,name=deploy_pipeline_spec,json=deployPipelineSpec,proto3" json:"deploy_pipeline_spec" bson:"-"`
	// 部署流水线配置Id
	// @gotags: bson:"deploy_pipeline_id" json:"deploy_pipeline_id"
	DeployPipelineId string `protobuf:"bytes,8,opt,name=deploy_pipeline_id,json=deployPipelineId,proto3" json:"deploy_pipeline_id" bson:"deploy_pipeline_id"`
}

func (x *CreateApprovalRequest) Reset() {
	*x = CreateApprovalRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_approval_pb_approval_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateApprovalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateApprovalRequest) ProtoMessage() {}

func (x *CreateApprovalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apps_approval_pb_approval_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateApprovalRequest.ProtoReflect.Descriptor instead.
func (*CreateApprovalRequest) Descriptor() ([]byte, []int) {
	return file_apps_approval_pb_approval_proto_rawDescGZIP(), []int{2}
}

func (x *CreateApprovalRequest) GetIsTemplate() bool {
	if x != nil {
		return x.IsTemplate
	}
	return false
}

func (x *CreateApprovalRequest) GetProposers() []string {
	if x != nil {
		return x.Proposers
	}
	return nil
}

func (x *CreateApprovalRequest) GetAuditors() []string {
	if x != nil {
		return x.Auditors
	}
	return nil
}

func (x *CreateApprovalRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *CreateApprovalRequest) GetDescribe() string {
	if x != nil {
		return x.Describe
	}
	return ""
}

func (x *CreateApprovalRequest) GetAutoPublish() bool {
	if x != nil {
		return x.AutoPublish
	}
	return false
}

func (x *CreateApprovalRequest) GetDeployPipelineSpec() *pipeline.CreatePipelineRequest {
	if x != nil {
		return x.DeployPipelineSpec
	}
	return nil
}

func (x *CreateApprovalRequest) GetDeployPipelineId() string {
	if x != nil {
		return x.DeployPipelineId
	}
	return ""
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 当前状态
	// @gotags: bson:"stage" json:"stage"
	Stage STAGE `protobuf:"varint,1,opt,name=stage,proto3,enum=infraboard.mpaas.approval.STAGE" json:"stage" bson:"stage"`
	// 审核通过的时间
	// @gotags: bson:"pass_at" json:"pass_at"
	PassAt int64 `protobuf:"varint,2,opt,name=pass_at,json=passAt,proto3" json:"pass_at" bson:"pass_at"`
	// 审核通过的意见
	// @gotags: bson:"comment" json:"comment"
	Comment string `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment" bson:"comment"`
	// 开始发布的时间
	// @gotags: bson:"publish_at" json:"publish_at"
	PublishAt int64 `protobuf:"varint,4,opt,name=publish_at,json=publishAt,proto3" json:"publish_at" bson:"publish_at"`
	// 发布结束的时间
	// @gotags: bson:"end_at" json:"end_at"
	EndAt int64 `protobuf:"varint,5,opt,name=end_at,json=endAt,proto3" json:"end_at" bson:"end_at"`
	// 发布关闭的时间
	// @gotags: bson:"close_at" json:"close_at"
	CloseAt int64 `protobuf:"varint,6,opt,name=close_at,json=closeAt,proto3" json:"close_at" bson:"close_at"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_approval_pb_approval_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_apps_approval_pb_approval_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_apps_approval_pb_approval_proto_rawDescGZIP(), []int{3}
}

func (x *Status) GetStage() STAGE {
	if x != nil {
		return x.Stage
	}
	return STAGE_DRAFT
}

func (x *Status) GetPassAt() int64 {
	if x != nil {
		return x.PassAt
	}
	return 0
}

func (x *Status) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *Status) GetPublishAt() int64 {
	if x != nil {
		return x.PublishAt
	}
	return 0
}

func (x *Status) GetEndAt() int64 {
	if x != nil {
		return x.EndAt
	}
	return 0
}

func (x *Status) GetCloseAt() int64 {
	if x != nil {
		return x.CloseAt
	}
	return 0
}

var File_apps_approval_pb_approval_proto protoreflect.FileDescriptor

var file_apps_approval_pb_approval_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x2f,
	0x70, 0x62, 0x2f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x19, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70,
	0x61, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x1a, 0x16, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5e, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61,
	0x6c, 0x53, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x39, 0x0a, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x70,
	0x72, 0x6f, 0x76, 0x61, 0x6c, 0x2e, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x52, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x91, 0x02, 0x0a, 0x08, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76,
	0x61, 0x6c, 0x12, 0x36, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70,
	0x61, 0x61, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e,
	0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x44, 0x0a, 0x04, 0x73, 0x70,
	0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x70, 0x72,
	0x6f, 0x76, 0x61, 0x6c, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x72, 0x6f,
	0x76, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63,
	0x12, 0x4c, 0x0a, 0x0f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x5f, 0x70, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x70, 0x69, 0x70,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x0e,
	0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x39,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61,
	0x73, 0x2e, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xdd, 0x02, 0x0a, 0x15, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65,
	0x72, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x75, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x61, 0x75, 0x74, 0x6f,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x12, 0x62, 0x0a, 0x14, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x5f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x12, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x50,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x2c, 0x0a, 0x12, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x5f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x50,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x49, 0x64, 0x22, 0xc4, 0x01, 0x0a, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x36, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x2e,
	0x53, 0x54, 0x41, 0x47, 0x45, 0x52, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x12, 0x17, 0x0a, 0x07,
	0x70, 0x61, 0x73, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70,
	0x61, 0x73, 0x73, 0x41, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x41, 0x74, 0x12, 0x15,
	0x0a, 0x06, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x65, 0x6e, 0x64, 0x41, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x5f, 0x61,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x41, 0x74,
	0x2a, 0x6a, 0x0a, 0x05, 0x53, 0x54, 0x41, 0x47, 0x45, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x52, 0x41,
	0x46, 0x54, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x45, 0x4e, 0x44, 0x44, 0x49, 0x4e, 0x47,
	0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x58, 0x50, 0x49, 0x52, 0x45, 0x44, 0x10, 0x02, 0x12,
	0x08, 0x0a, 0x04, 0x44, 0x45, 0x4e, 0x59, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x41, 0x53,
	0x53, 0x45, 0x44, 0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x55, 0x43, 0x43, 0x45, 0x45, 0x44,
	0x45, 0x44, 0x10, 0x0b, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x0c,
	0x12, 0x0a, 0x0a, 0x06, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x44, 0x10, 0x0d, 0x42, 0x2b, 0x5a, 0x29,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x73,
	0x2f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_apps_approval_pb_approval_proto_rawDescOnce sync.Once
	file_apps_approval_pb_approval_proto_rawDescData = file_apps_approval_pb_approval_proto_rawDesc
)

func file_apps_approval_pb_approval_proto_rawDescGZIP() []byte {
	file_apps_approval_pb_approval_proto_rawDescOnce.Do(func() {
		file_apps_approval_pb_approval_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_approval_pb_approval_proto_rawDescData)
	})
	return file_apps_approval_pb_approval_proto_rawDescData
}

var file_apps_approval_pb_approval_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_apps_approval_pb_approval_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_apps_approval_pb_approval_proto_goTypes = []interface{}{
	(STAGE)(0),                             // 0: infraboard.mpaas.approval.STAGE
	(*ApprovalSet)(nil),                    // 1: infraboard.mpaas.approval.ApprovalSet
	(*Approval)(nil),                       // 2: infraboard.mpaas.approval.Approval
	(*CreateApprovalRequest)(nil),          // 3: infraboard.mpaas.approval.CreateApprovalRequest
	(*Status)(nil),                         // 4: infraboard.mpaas.approval.Status
	(*meta.Meta)(nil),                      // 5: infraboard.mpaas.common.meta.Meta
	(*pipeline.Pipeline)(nil),              // 6: infraboard.mpaas.pipeline.Pipeline
	(*pipeline.CreatePipelineRequest)(nil), // 7: infraboard.mpaas.pipeline.CreatePipelineRequest
}
var file_apps_approval_pb_approval_proto_depIdxs = []int32{
	2, // 0: infraboard.mpaas.approval.ApprovalSet.items:type_name -> infraboard.mpaas.approval.Approval
	5, // 1: infraboard.mpaas.approval.Approval.meta:type_name -> infraboard.mpaas.common.meta.Meta
	3, // 2: infraboard.mpaas.approval.Approval.spec:type_name -> infraboard.mpaas.approval.CreateApprovalRequest
	6, // 3: infraboard.mpaas.approval.Approval.deploy_pipeline:type_name -> infraboard.mpaas.pipeline.Pipeline
	4, // 4: infraboard.mpaas.approval.Approval.status:type_name -> infraboard.mpaas.approval.Status
	7, // 5: infraboard.mpaas.approval.CreateApprovalRequest.deploy_pipeline_spec:type_name -> infraboard.mpaas.pipeline.CreatePipelineRequest
	0, // 6: infraboard.mpaas.approval.Status.stage:type_name -> infraboard.mpaas.approval.STAGE
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_apps_approval_pb_approval_proto_init() }
func file_apps_approval_pb_approval_proto_init() {
	if File_apps_approval_pb_approval_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_approval_pb_approval_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApprovalSet); i {
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
		file_apps_approval_pb_approval_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Approval); i {
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
		file_apps_approval_pb_approval_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateApprovalRequest); i {
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
		file_apps_approval_pb_approval_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_apps_approval_pb_approval_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apps_approval_pb_approval_proto_goTypes,
		DependencyIndexes: file_apps_approval_pb_approval_proto_depIdxs,
		EnumInfos:         file_apps_approval_pb_approval_proto_enumTypes,
		MessageInfos:      file_apps_approval_pb_approval_proto_msgTypes,
	}.Build()
	File_apps_approval_pb_approval_proto = out.File
	file_apps_approval_pb_approval_proto_rawDesc = nil
	file_apps_approval_pb_approval_proto_goTypes = nil
	file_apps_approval_pb_approval_proto_depIdxs = nil
}

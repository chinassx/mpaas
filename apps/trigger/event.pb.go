// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: apps/trigger/pb/event.proto

package trigger

import (
	build "github.com/infraboard/mpaas/apps/build"
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

type EVENT_PROVIDER int32

const (
	// 来自gitlab的事件
	EVENT_PROVIDER_GITLAB EVENT_PROVIDER = 0
)

// Enum value maps for EVENT_PROVIDER.
var (
	EVENT_PROVIDER_name = map[int32]string{
		0: "GITLAB",
	}
	EVENT_PROVIDER_value = map[string]int32{
		"GITLAB": 0,
	}
)

func (x EVENT_PROVIDER) Enum() *EVENT_PROVIDER {
	p := new(EVENT_PROVIDER)
	*p = x
	return p
}

func (x EVENT_PROVIDER) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EVENT_PROVIDER) Descriptor() protoreflect.EnumDescriptor {
	return file_apps_trigger_pb_event_proto_enumTypes[0].Descriptor()
}

func (EVENT_PROVIDER) Type() protoreflect.EnumType {
	return &file_apps_trigger_pb_event_proto_enumTypes[0]
}

func (x EVENT_PROVIDER) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EVENT_PROVIDER.Descriptor instead.
func (EVENT_PROVIDER) EnumDescriptor() ([]byte, []int) {
	return file_apps_trigger_pb_event_proto_rawDescGZIP(), []int{0}
}

type RecordSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 总数
	// @gotags: bson:"total" json:"total"
	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total" bson:"total"`
	// 列表
	// @gotags: bson:"items" json:"items"
	Items []*Record `protobuf:"bytes,2,rep,name=items,proto3" json:"items" bson:"items"`
}

func (x *RecordSet) Reset() {
	*x = RecordSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_trigger_pb_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordSet) ProtoMessage() {}

func (x *RecordSet) ProtoReflect() protoreflect.Message {
	mi := &file_apps_trigger_pb_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordSet.ProtoReflect.Descriptor instead.
func (*RecordSet) Descriptor() ([]byte, []int) {
	return file_apps_trigger_pb_event_proto_rawDescGZIP(), []int{0}
}

func (x *RecordSet) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *RecordSet) GetItems() []*Record {
	if x != nil {
		return x.Items
	}
	return nil
}

// 事件记录
type Record struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 元信息
	// @gotags: bson:",inline" json:"meta"
	Meta *meta.Meta `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta" bson:",inline"`
	// event相关定义
	// @gotags: bson:",inline" json:"event"
	Event *Event `protobuf:"bytes,2,opt,name=event,proto3" json:"event" bson:",inline"`
	// 构建状态
	// @gotags: bson:"build_status" json:"build_status"
	BuildStatus []*BuildStatus `protobuf:"bytes,3,rep,name=build_status,json=buildStatus,proto3" json:"build_status" bson:"build_status"`
}

func (x *Record) Reset() {
	*x = Record{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_trigger_pb_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Record) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Record) ProtoMessage() {}

func (x *Record) ProtoReflect() protoreflect.Message {
	mi := &file_apps_trigger_pb_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Record.ProtoReflect.Descriptor instead.
func (*Record) Descriptor() ([]byte, []int) {
	return file_apps_trigger_pb_event_proto_rawDescGZIP(), []int{1}
}

func (x *Record) GetMeta() *meta.Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Record) GetEvent() *Event {
	if x != nil {
		return x.Event
	}
	return nil
}

func (x *Record) GetBuildStatus() []*BuildStatus {
	if x != nil {
		return x.BuildStatus
	}
	return nil
}

type BuildStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 构建信息, 由于可能修改, 因此需要保存整个对象
	// @gotags: bson:"build_config" json:"build_config"
	BuildConfig *build.BuildConfig `protobuf:"bytes,1,opt,name=build_config,json=buildConfig,proto3" json:"build_config" bson:"build_config"`
	// 构建信息
	// @gotags: bson:"pipline_task_id" json:"pipline_task_id"
	PiplineTaskId string `protobuf:"bytes,2,opt,name=pipline_task_id,json=piplineTaskId,proto3" json:"pipline_task_id" bson:"pipline_task_id"`
	// 如果流水线运行报错的报错信息
	// @gotags: bson:"error_message" json:"error_message"
	ErrorMessage string `protobuf:"bytes,3,opt,name=error_message,json=errorMessage,proto3" json:"error_message" bson:"error_message"`
}

func (x *BuildStatus) Reset() {
	*x = BuildStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_trigger_pb_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildStatus) ProtoMessage() {}

func (x *BuildStatus) ProtoReflect() protoreflect.Message {
	mi := &file_apps_trigger_pb_event_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildStatus.ProtoReflect.Descriptor instead.
func (*BuildStatus) Descriptor() ([]byte, []int) {
	return file_apps_trigger_pb_event_proto_rawDescGZIP(), []int{2}
}

func (x *BuildStatus) GetBuildConfig() *build.BuildConfig {
	if x != nil {
		return x.BuildConfig
	}
	return nil
}

func (x *BuildStatus) GetPiplineTaskId() string {
	if x != nil {
		return x.PiplineTaskId
	}
	return ""
}

func (x *BuildStatus) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 服务Id
	// @gotags: json:"service_id" validate:"required"
	ServiceId string `protobuf:"bytes,1,opt,name=service_id,json=serviceId,proto3" json:"service_id" validate:"required"`
	// 应用ID
	// @gotags: json:"provider"
	Provider EVENT_PROVIDER `protobuf:"varint,2,opt,name=provider,proto3,enum=infraboard.mpaas.trigger.EVENT_PROVIDER" json:"provider"`
	// gitlab webhook事件
	// @gotags: json:"gitlab_event" validate:"-"
	GitlabEvent *GitlabWebHookEvent `protobuf:"bytes,3,opt,name=gitlab_event,json=gitlabEvent,proto3" json:"gitlab_event" validate:"-"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_trigger_pb_event_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_apps_trigger_pb_event_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_apps_trigger_pb_event_proto_rawDescGZIP(), []int{3}
}

func (x *Event) GetServiceId() string {
	if x != nil {
		return x.ServiceId
	}
	return ""
}

func (x *Event) GetProvider() EVENT_PROVIDER {
	if x != nil {
		return x.Provider
	}
	return EVENT_PROVIDER_GITLAB
}

func (x *Event) GetGitlabEvent() *GitlabWebHookEvent {
	if x != nil {
		return x.GitlabEvent
	}
	return nil
}

var File_apps_trigger_pb_event_proto protoreflect.FileDescriptor

var file_apps_trigger_pb_event_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x70,
	0x62, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e,
	0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x1a, 0x1c, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x74, 0x72,
	0x69, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x2f, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x2f, 0x70, 0x62, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x16, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x6d, 0x65,
	0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x59, 0x0a, 0x09, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x53, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x36, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x74, 0x72,
	0x69, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x22, 0xc1, 0x01, 0x0a, 0x06, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x36,
	0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x4d, 0x65, 0x74, 0x61,
	0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x35, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72,
	0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x48, 0x0a,
	0x0c, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x42,
	0x75, 0x69, 0x6c, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0b, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xa2, 0x01, 0x0a, 0x0b, 0x42, 0x75, 0x69, 0x6c,
	0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x46, 0x0a, 0x0c, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73,
	0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x0b, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x26, 0x0a, 0x0f, 0x70, 0x69, 0x70, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x69, 0x70, 0x6c, 0x69, 0x6e,
	0x65, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xbd, 0x01, 0x0a,
	0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x44, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x74, 0x72, 0x69, 0x67, 0x67,
	0x65, 0x72, 0x2e, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45,
	0x52, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x4f, 0x0a, 0x0c, 0x67,
	0x69, 0x74, 0x6c, 0x61, 0x62, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2c, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d,
	0x70, 0x61, 0x61, 0x73, 0x2e, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x47, 0x69, 0x74,
	0x6c, 0x61, 0x62, 0x57, 0x65, 0x62, 0x48, 0x6f, 0x6f, 0x6b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52,
	0x0b, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2a, 0x1c, 0x0a, 0x0e,
	0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x12, 0x0a,
	0x0a, 0x06, 0x47, 0x49, 0x54, 0x4c, 0x41, 0x42, 0x10, 0x00, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2f, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x74,
	0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_trigger_pb_event_proto_rawDescOnce sync.Once
	file_apps_trigger_pb_event_proto_rawDescData = file_apps_trigger_pb_event_proto_rawDesc
)

func file_apps_trigger_pb_event_proto_rawDescGZIP() []byte {
	file_apps_trigger_pb_event_proto_rawDescOnce.Do(func() {
		file_apps_trigger_pb_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_trigger_pb_event_proto_rawDescData)
	})
	return file_apps_trigger_pb_event_proto_rawDescData
}

var file_apps_trigger_pb_event_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_apps_trigger_pb_event_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_apps_trigger_pb_event_proto_goTypes = []interface{}{
	(EVENT_PROVIDER)(0),        // 0: infraboard.mpaas.trigger.EVENT_PROVIDER
	(*RecordSet)(nil),          // 1: infraboard.mpaas.trigger.RecordSet
	(*Record)(nil),             // 2: infraboard.mpaas.trigger.Record
	(*BuildStatus)(nil),        // 3: infraboard.mpaas.trigger.BuildStatus
	(*Event)(nil),              // 4: infraboard.mpaas.trigger.Event
	(*meta.Meta)(nil),          // 5: infraboard.mpaas.common.meta.Meta
	(*build.BuildConfig)(nil),  // 6: infraboard.mpaas.build.BuildConfig
	(*GitlabWebHookEvent)(nil), // 7: infraboard.mpaas.trigger.GitlabWebHookEvent
}
var file_apps_trigger_pb_event_proto_depIdxs = []int32{
	2, // 0: infraboard.mpaas.trigger.RecordSet.items:type_name -> infraboard.mpaas.trigger.Record
	5, // 1: infraboard.mpaas.trigger.Record.meta:type_name -> infraboard.mpaas.common.meta.Meta
	4, // 2: infraboard.mpaas.trigger.Record.event:type_name -> infraboard.mpaas.trigger.Event
	3, // 3: infraboard.mpaas.trigger.Record.build_status:type_name -> infraboard.mpaas.trigger.BuildStatus
	6, // 4: infraboard.mpaas.trigger.BuildStatus.build_config:type_name -> infraboard.mpaas.build.BuildConfig
	0, // 5: infraboard.mpaas.trigger.Event.provider:type_name -> infraboard.mpaas.trigger.EVENT_PROVIDER
	7, // 6: infraboard.mpaas.trigger.Event.gitlab_event:type_name -> infraboard.mpaas.trigger.GitlabWebHookEvent
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_apps_trigger_pb_event_proto_init() }
func file_apps_trigger_pb_event_proto_init() {
	if File_apps_trigger_pb_event_proto != nil {
		return
	}
	file_apps_trigger_pb_gitlab_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_apps_trigger_pb_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordSet); i {
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
		file_apps_trigger_pb_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Record); i {
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
		file_apps_trigger_pb_event_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildStatus); i {
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
		file_apps_trigger_pb_event_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
			RawDescriptor: file_apps_trigger_pb_event_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apps_trigger_pb_event_proto_goTypes,
		DependencyIndexes: file_apps_trigger_pb_event_proto_depIdxs,
		EnumInfos:         file_apps_trigger_pb_event_proto_enumTypes,
		MessageInfos:      file_apps_trigger_pb_event_proto_msgTypes,
	}.Build()
	File_apps_trigger_pb_event_proto = out.File
	file_apps_trigger_pb_event_proto_rawDesc = nil
	file_apps_trigger_pb_event_proto_goTypes = nil
	file_apps_trigger_pb_event_proto_depIdxs = nil
}

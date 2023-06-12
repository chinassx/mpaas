// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.6
// source: mpaas/apps/event/pb/event.proto

package event

import (
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

// 事件的级别
type LEVEL int32

const (
	LEVEL_TRACE LEVEL = 0
	LEVEL_DEBUG LEVEL = 1
	LEVEL_INFO  LEVEL = 2
	LEVEL_WARNN LEVEL = 3
	LEVEL_ERROR LEVEL = 4
	LEVEL_PANIC LEVEL = 5
)

// Enum value maps for LEVEL.
var (
	LEVEL_name = map[int32]string{
		0: "TRACE",
		1: "DEBUG",
		2: "INFO",
		3: "WARNN",
		4: "ERROR",
		5: "PANIC",
	}
	LEVEL_value = map[string]int32{
		"TRACE": 0,
		"DEBUG": 1,
		"INFO":  2,
		"WARNN": 3,
		"ERROR": 4,
		"PANIC": 5,
	}
)

func (x LEVEL) Enum() *LEVEL {
	p := new(LEVEL)
	*p = x
	return p
}

func (x LEVEL) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LEVEL) Descriptor() protoreflect.EnumDescriptor {
	return file_mpaas_apps_event_pb_event_proto_enumTypes[0].Descriptor()
}

func (LEVEL) Type() protoreflect.EnumType {
	return &file_mpaas_apps_event_pb_event_proto_enumTypes[0]
}

func (x LEVEL) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LEVEL.Descriptor instead.
func (LEVEL) EnumDescriptor() ([]byte, []int) {
	return file_mpaas_apps_event_pb_event_proto_rawDescGZIP(), []int{0}
}

// 事件状态
type STAGE int32

const (
	// 告警中
	STAGE_FIRING STAGE = 0
	// 告警恢复
	STAGE_RESOLVED STAGE = 1
	// 过期
	STAGE_EXPIRED STAGE = 2
)

// Enum value maps for STAGE.
var (
	STAGE_name = map[int32]string{
		0: "FIRING",
		1: "RESOLVED",
		2: "EXPIRED",
	}
	STAGE_value = map[string]int32{
		"FIRING":   0,
		"RESOLVED": 1,
		"EXPIRED":  2,
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
	return file_mpaas_apps_event_pb_event_proto_enumTypes[1].Descriptor()
}

func (STAGE) Type() protoreflect.EnumType {
	return &file_mpaas_apps_event_pb_event_proto_enumTypes[1]
}

func (x STAGE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use STAGE.Descriptor instead.
func (STAGE) EnumDescriptor() ([]byte, []int) {
	return file_mpaas_apps_event_pb_event_proto_rawDescGZIP(), []int{1}
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 创建信息
	// @gotags: bson:",inline" json:"spec"
	Spec *CreateEventRequest `protobuf:"bytes,1,opt,name=spec,proto3" json:"spec" bson:",inline"`
	// 事件状态
	// @gotags: bson:"status" json:"status"
	Status *Status `protobuf:"bytes,2,opt,name=status,proto3" json:"status" bson:"status"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mpaas_apps_event_pb_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_mpaas_apps_event_pb_event_proto_msgTypes[0]
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
	return file_mpaas_apps_event_pb_event_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetSpec() *CreateEventRequest {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *Event) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

// 创建事件
type CreateEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 事件Id, 如果不传会自动生成
	// @gotags: bson:"_id" json:"id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"_id"`
	// 事件触发的事件
	// @gotags: bson:"time" json:"time"
	Time int64 `protobuf:"varint,2,opt,name=time,proto3" json:"time" bson:"time"`
	// 事件的级别
	// @gotags: bson:"level" json:"level"
	Level LEVEL `protobuf:"varint,3,opt,name=level,proto3,enum=infraboard.mpaas.event.LEVEL" json:"level" bson:"level"`
	// 是否是临时事件
	// @gotags: bson:"is_temporary" json:"is_temporary"
	IsTemporary bool `protobuf:"varint,4,opt,name=is_temporary,json=isTemporary,proto3" json:"is_temporary" bson:"is_temporary"`
}

func (x *CreateEventRequest) Reset() {
	*x = CreateEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mpaas_apps_event_pb_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEventRequest) ProtoMessage() {}

func (x *CreateEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mpaas_apps_event_pb_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEventRequest.ProtoReflect.Descriptor instead.
func (*CreateEventRequest) Descriptor() ([]byte, []int) {
	return file_mpaas_apps_event_pb_event_proto_rawDescGZIP(), []int{1}
}

func (x *CreateEventRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateEventRequest) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *CreateEventRequest) GetLevel() LEVEL {
	if x != nil {
		return x.Level
	}
	return LEVEL_TRACE
}

func (x *CreateEventRequest) GetIsTemporary() bool {
	if x != nil {
		return x.IsTemporary
	}
	return false
}

// 事件状态
type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 当前状态
	// @gotags: bson:"stage" json:"stage"
	Stage STAGE `protobuf:"varint,1,opt,name=stage,proto3,enum=infraboard.mpaas.event.STAGE" json:"stage" bson:"stage"`
	// 更新时间
	// @gotags: bson:"update_at" json:"update_at"
	UpdateAt int64 `protobuf:"varint,2,opt,name=update_at,json=updateAt,proto3" json:"update_at" bson:"update_at"`
	// 恢复时间
	// @gotags: bson:"resolved_at" json:"resolved_at"
	ResolvedAt int64 `protobuf:"varint,3,opt,name=resolved_at,json=resolvedAt,proto3" json:"resolved_at" bson:"resolved_at"`
	// 过期时间
	// @gotags: bson:"expired_at" json:"expired_at"
	ExpiredAt int64 `protobuf:"varint,4,opt,name=expired_at,json=expiredAt,proto3" json:"expired_at" bson:"expired_at"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mpaas_apps_event_pb_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_mpaas_apps_event_pb_event_proto_msgTypes[2]
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
	return file_mpaas_apps_event_pb_event_proto_rawDescGZIP(), []int{2}
}

func (x *Status) GetStage() STAGE {
	if x != nil {
		return x.Stage
	}
	return STAGE_FIRING
}

func (x *Status) GetUpdateAt() int64 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

func (x *Status) GetResolvedAt() int64 {
	if x != nil {
		return x.ResolvedAt
	}
	return 0
}

func (x *Status) GetExpiredAt() int64 {
	if x != nil {
		return x.ExpiredAt
	}
	return 0
}

// 事件跟踪信息
type TraceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 跟进时间
	// @gotags: bson:"trace_at" json:"trace_at"
	TraceAt int64 `protobuf:"varint,2,opt,name=trace_at,json=traceAt,proto3" json:"trace_at" bson:"trace_at"`
	// 根据人
	// @gotags: bson:"trace_by" json:"trace_by"
	TraceBy string `protobuf:"bytes,3,opt,name=trace_by,json=traceBy,proto3" json:"trace_by" bson:"trace_by"`
}

func (x *TraceInfo) Reset() {
	*x = TraceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mpaas_apps_event_pb_event_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TraceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TraceInfo) ProtoMessage() {}

func (x *TraceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_mpaas_apps_event_pb_event_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TraceInfo.ProtoReflect.Descriptor instead.
func (*TraceInfo) Descriptor() ([]byte, []int) {
	return file_mpaas_apps_event_pb_event_proto_rawDescGZIP(), []int{3}
}

func (x *TraceInfo) GetTraceAt() int64 {
	if x != nil {
		return x.TraceAt
	}
	return 0
}

func (x *TraceInfo) GetTraceBy() string {
	if x != nil {
		return x.TraceBy
	}
	return ""
}

var File_mpaas_apps_event_pb_event_proto protoreflect.FileDescriptor

var file_mpaas_apps_event_pb_event_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x2f, 0x70, 0x62, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x16, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70,
	0x61, 0x61, 0x73, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x7f, 0x0a, 0x05, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x3e, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2a, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70,
	0x61, 0x61, 0x73, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x73, 0x70,
	0x65, 0x63, 0x12, 0x36, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x90, 0x01, 0x0a, 0x12, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x4c, 0x45,
	0x56, 0x45, 0x4c, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x73,
	0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0b, 0x69, 0x73, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x72, 0x79, 0x22, 0x9a, 0x01,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x33, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x2e, 0x53, 0x54, 0x41, 0x47, 0x45, 0x52, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65,
	0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x41, 0x74, 0x22, 0x41, 0x0a, 0x09, 0x54, 0x72,
	0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x63, 0x65,
	0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x74, 0x72, 0x61, 0x63, 0x65,
	0x41, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x62, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x61, 0x63, 0x65, 0x42, 0x79, 0x2a, 0x48, 0x0a,
	0x05, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x12, 0x09, 0x0a, 0x05, 0x54, 0x52, 0x41, 0x43, 0x45, 0x10,
	0x00, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x45, 0x42, 0x55, 0x47, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04,
	0x49, 0x4e, 0x46, 0x4f, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x57, 0x41, 0x52, 0x4e, 0x4e, 0x10,
	0x03, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05,
	0x50, 0x41, 0x4e, 0x49, 0x43, 0x10, 0x05, 0x2a, 0x2e, 0x0a, 0x05, 0x53, 0x54, 0x41, 0x47, 0x45,
	0x12, 0x0a, 0x0a, 0x06, 0x46, 0x49, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08,
	0x52, 0x45, 0x53, 0x4f, 0x4c, 0x56, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x58,
	0x50, 0x49, 0x52, 0x45, 0x44, 0x10, 0x02, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2f, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mpaas_apps_event_pb_event_proto_rawDescOnce sync.Once
	file_mpaas_apps_event_pb_event_proto_rawDescData = file_mpaas_apps_event_pb_event_proto_rawDesc
)

func file_mpaas_apps_event_pb_event_proto_rawDescGZIP() []byte {
	file_mpaas_apps_event_pb_event_proto_rawDescOnce.Do(func() {
		file_mpaas_apps_event_pb_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_mpaas_apps_event_pb_event_proto_rawDescData)
	})
	return file_mpaas_apps_event_pb_event_proto_rawDescData
}

var file_mpaas_apps_event_pb_event_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_mpaas_apps_event_pb_event_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_mpaas_apps_event_pb_event_proto_goTypes = []interface{}{
	(LEVEL)(0),                 // 0: infraboard.mpaas.event.LEVEL
	(STAGE)(0),                 // 1: infraboard.mpaas.event.STAGE
	(*Event)(nil),              // 2: infraboard.mpaas.event.Event
	(*CreateEventRequest)(nil), // 3: infraboard.mpaas.event.CreateEventRequest
	(*Status)(nil),             // 4: infraboard.mpaas.event.Status
	(*TraceInfo)(nil),          // 5: infraboard.mpaas.event.TraceInfo
}
var file_mpaas_apps_event_pb_event_proto_depIdxs = []int32{
	3, // 0: infraboard.mpaas.event.Event.spec:type_name -> infraboard.mpaas.event.CreateEventRequest
	4, // 1: infraboard.mpaas.event.Event.status:type_name -> infraboard.mpaas.event.Status
	0, // 2: infraboard.mpaas.event.CreateEventRequest.level:type_name -> infraboard.mpaas.event.LEVEL
	1, // 3: infraboard.mpaas.event.Status.stage:type_name -> infraboard.mpaas.event.STAGE
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_mpaas_apps_event_pb_event_proto_init() }
func file_mpaas_apps_event_pb_event_proto_init() {
	if File_mpaas_apps_event_pb_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mpaas_apps_event_pb_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_mpaas_apps_event_pb_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEventRequest); i {
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
		file_mpaas_apps_event_pb_event_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_mpaas_apps_event_pb_event_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TraceInfo); i {
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
			RawDescriptor: file_mpaas_apps_event_pb_event_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mpaas_apps_event_pb_event_proto_goTypes,
		DependencyIndexes: file_mpaas_apps_event_pb_event_proto_depIdxs,
		EnumInfos:         file_mpaas_apps_event_pb_event_proto_enumTypes,
		MessageInfos:      file_mpaas_apps_event_pb_event_proto_msgTypes,
	}.Build()
	File_mpaas_apps_event_pb_event_proto = out.File
	file_mpaas_apps_event_pb_event_proto_rawDesc = nil
	file_mpaas_apps_event_pb_event_proto_goTypes = nil
	file_mpaas_apps_event_pb_event_proto_depIdxs = nil
}

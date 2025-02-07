// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.6
// source: mpaas/apps/trigger/pb/rpc.proto

package trigger

import (
	request "github.com/infraboard/mcube/http/request"
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

type QueryRecordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 分页请求
	// @gotags: json:"page"
	Page *request.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page"`
	// 服务Id, 查询某个服务的事件
	// @gotags: json:"service_id"
	ServiceId string `protobuf:"bytes,2,opt,name=service_id,json=serviceId,proto3" json:"service_id"`
	// 查询PipelineTask关联的事件
	// @gotags: json:"pipeline_task_id"
	PipelineTaskId string `protobuf:"bytes,3,opt,name=pipeline_task_id,json=pipelineTaskId,proto3" json:"pipeline_task_id"`
}

func (x *QueryRecordRequest) Reset() {
	*x = QueryRecordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mpaas_apps_trigger_pb_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRecordRequest) ProtoMessage() {}

func (x *QueryRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mpaas_apps_trigger_pb_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRecordRequest.ProtoReflect.Descriptor instead.
func (*QueryRecordRequest) Descriptor() ([]byte, []int) {
	return file_mpaas_apps_trigger_pb_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *QueryRecordRequest) GetPage() *request.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *QueryRecordRequest) GetServiceId() string {
	if x != nil {
		return x.ServiceId
	}
	return ""
}

func (x *QueryRecordRequest) GetPipelineTaskId() string {
	if x != nil {
		return x.PipelineTaskId
	}
	return ""
}

type EventQueueTaskCompleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 执行完成的PipelineTask任务Id
	// @gotags: json:"pipeline_task_id"
	PipelineTaskId string `protobuf:"bytes,1,opt,name=pipeline_task_id,json=pipelineTaskId,proto3" json:"pipeline_task_id"`
}

func (x *EventQueueTaskCompleteRequest) Reset() {
	*x = EventQueueTaskCompleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mpaas_apps_trigger_pb_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventQueueTaskCompleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventQueueTaskCompleteRequest) ProtoMessage() {}

func (x *EventQueueTaskCompleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mpaas_apps_trigger_pb_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventQueueTaskCompleteRequest.ProtoReflect.Descriptor instead.
func (*EventQueueTaskCompleteRequest) Descriptor() ([]byte, []int) {
	return file_mpaas_apps_trigger_pb_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *EventQueueTaskCompleteRequest) GetPipelineTaskId() string {
	if x != nil {
		return x.PipelineTaskId
	}
	return ""
}

var File_mpaas_apps_trigger_pb_rpc_proto protoreflect.FileDescriptor

var file_mpaas_apps_trigger_pb_rpc_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x74, 0x72, 0x69,
	0x67, 0x67, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x18, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x70,
	0x61, 0x61, 0x73, 0x2e, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x1a, 0x21, 0x6d, 0x70, 0x61,
	0x61, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x2f,
	0x70, 0x62, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18,
	0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x61,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x95, 0x01, 0x0a, 0x12, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x36, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65,
	0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x64,
	0x22, 0x49, 0x0a, 0x1d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x51, 0x75, 0x65, 0x75, 0x65, 0x54, 0x61,
	0x73, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x28, 0x0a, 0x10, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x74, 0x61,
	0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x69, 0x70,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x32, 0xb9, 0x01, 0x0a, 0x03,
	0x52, 0x50, 0x43, 0x12, 0x50, 0x0a, 0x0b, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x1f, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x1a, 0x20, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x60, 0x0a, 0x0b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x12, 0x2c, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2e, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x2e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x23, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x6d, 0x70, 0x61, 0x61, 0x73, 0x2e, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x53, 0x65, 0x74, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2f, 0x6d, 0x70, 0x61, 0x61, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x74, 0x72, 0x69, 0x67,
	0x67, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mpaas_apps_trigger_pb_rpc_proto_rawDescOnce sync.Once
	file_mpaas_apps_trigger_pb_rpc_proto_rawDescData = file_mpaas_apps_trigger_pb_rpc_proto_rawDesc
)

func file_mpaas_apps_trigger_pb_rpc_proto_rawDescGZIP() []byte {
	file_mpaas_apps_trigger_pb_rpc_proto_rawDescOnce.Do(func() {
		file_mpaas_apps_trigger_pb_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_mpaas_apps_trigger_pb_rpc_proto_rawDescData)
	})
	return file_mpaas_apps_trigger_pb_rpc_proto_rawDescData
}

var file_mpaas_apps_trigger_pb_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_mpaas_apps_trigger_pb_rpc_proto_goTypes = []interface{}{
	(*QueryRecordRequest)(nil),            // 0: infraboard.mpaas.trigger.QueryRecordRequest
	(*EventQueueTaskCompleteRequest)(nil), // 1: infraboard.mpaas.trigger.EventQueueTaskCompleteRequest
	(*request.PageRequest)(nil),           // 2: infraboard.mcube.page.PageRequest
	(*Event)(nil),                         // 3: infraboard.mpaas.trigger.Event
	(*Record)(nil),                        // 4: infraboard.mpaas.trigger.Record
	(*RecordSet)(nil),                     // 5: infraboard.mpaas.trigger.RecordSet
}
var file_mpaas_apps_trigger_pb_rpc_proto_depIdxs = []int32{
	2, // 0: infraboard.mpaas.trigger.QueryRecordRequest.page:type_name -> infraboard.mcube.page.PageRequest
	3, // 1: infraboard.mpaas.trigger.RPC.HandleEvent:input_type -> infraboard.mpaas.trigger.Event
	0, // 2: infraboard.mpaas.trigger.RPC.QueryRecord:input_type -> infraboard.mpaas.trigger.QueryRecordRequest
	4, // 3: infraboard.mpaas.trigger.RPC.HandleEvent:output_type -> infraboard.mpaas.trigger.Record
	5, // 4: infraboard.mpaas.trigger.RPC.QueryRecord:output_type -> infraboard.mpaas.trigger.RecordSet
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_mpaas_apps_trigger_pb_rpc_proto_init() }
func file_mpaas_apps_trigger_pb_rpc_proto_init() {
	if File_mpaas_apps_trigger_pb_rpc_proto != nil {
		return
	}
	file_mpaas_apps_trigger_pb_event_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_mpaas_apps_trigger_pb_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryRecordRequest); i {
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
		file_mpaas_apps_trigger_pb_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventQueueTaskCompleteRequest); i {
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
			RawDescriptor: file_mpaas_apps_trigger_pb_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mpaas_apps_trigger_pb_rpc_proto_goTypes,
		DependencyIndexes: file_mpaas_apps_trigger_pb_rpc_proto_depIdxs,
		MessageInfos:      file_mpaas_apps_trigger_pb_rpc_proto_msgTypes,
	}.Build()
	File_mpaas_apps_trigger_pb_rpc_proto = out.File
	file_mpaas_apps_trigger_pb_rpc_proto_rawDesc = nil
	file_mpaas_apps_trigger_pb_rpc_proto_goTypes = nil
	file_mpaas_apps_trigger_pb_rpc_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: mq.proto

package mq

import (
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type PublishRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Queue string `protobuf:"bytes,1,opt,name=queue,proto3" json:"queue,omitempty"`
	Data  string `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *PublishRequest) Reset() {
	*x = PublishRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishRequest) ProtoMessage() {}

func (x *PublishRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishRequest.ProtoReflect.Descriptor instead.
func (*PublishRequest) Descriptor() ([]byte, []int) {
	return file_mq_proto_rawDescGZIP(), []int{0}
}

func (x *PublishRequest) GetQueue() string {
	if x != nil {
		return x.Queue
	}
	return ""
}

func (x *PublishRequest) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type DelayPublishRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Queue        string `protobuf:"bytes,1,opt,name=queue,proto3" json:"queue,omitempty"`
	Data         string `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	DelaySeconds uint64 `protobuf:"varint,3,opt,name=delaySeconds,proto3" json:"delaySeconds,omitempty"`
}

func (x *DelayPublishRequest) Reset() {
	*x = DelayPublishRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mq_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelayPublishRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelayPublishRequest) ProtoMessage() {}

func (x *DelayPublishRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mq_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelayPublishRequest.ProtoReflect.Descriptor instead.
func (*DelayPublishRequest) Descriptor() ([]byte, []int) {
	return file_mq_proto_rawDescGZIP(), []int{1}
}

func (x *DelayPublishRequest) GetQueue() string {
	if x != nil {
		return x.Queue
	}
	return ""
}

func (x *DelayPublishRequest) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *DelayPublishRequest) GetDelaySeconds() uint64 {
	if x != nil {
		return x.DelaySeconds
	}
	return 0
}

type SubscribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Queue string `protobuf:"bytes,1,opt,name=queue,proto3" json:"queue,omitempty"`
}

func (x *SubscribeRequest) Reset() {
	*x = SubscribeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mq_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeRequest) ProtoMessage() {}

func (x *SubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mq_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeRequest.ProtoReflect.Descriptor instead.
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return file_mq_proto_rawDescGZIP(), []int{2}
}

func (x *SubscribeRequest) GetQueue() string {
	if x != nil {
		return x.Queue
	}
	return ""
}

type SubscribeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Data  string `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Queue string `protobuf:"bytes,3,opt,name=queue,proto3" json:"queue,omitempty"`
}

func (x *SubscribeResponse) Reset() {
	*x = SubscribeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mq_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeResponse) ProtoMessage() {}

func (x *SubscribeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mq_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeResponse.ProtoReflect.Descriptor instead.
func (*SubscribeResponse) Descriptor() ([]byte, []int) {
	return file_mq_proto_rawDescGZIP(), []int{3}
}

func (x *SubscribeResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SubscribeResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *SubscribeResponse) GetQueue() string {
	if x != nil {
		return x.Queue
	}
	return ""
}

type QueueId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *QueueId) Reset() {
	*x = QueueId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mq_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueueId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueueId) ProtoMessage() {}

func (x *QueueId) ProtoReflect() protoreflect.Message {
	mi := &file_mq_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueueId.ProtoReflect.Descriptor instead.
func (*QueueId) Descriptor() ([]byte, []int) {
	return file_mq_proto_rawDescGZIP(), []int{4}
}

func (x *QueueId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type TopicPublishRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Data  string `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *TopicPublishRequest) Reset() {
	*x = TopicPublishRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mq_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopicPublishRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopicPublishRequest) ProtoMessage() {}

func (x *TopicPublishRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mq_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopicPublishRequest.ProtoReflect.Descriptor instead.
func (*TopicPublishRequest) Descriptor() ([]byte, []int) {
	return file_mq_proto_rawDescGZIP(), []int{5}
}

func (x *TopicPublishRequest) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *TopicPublishRequest) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type DelayTopicPublishRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic        string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Data         string `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	DelaySeconds uint64 `protobuf:"varint,3,opt,name=delaySeconds,proto3" json:"delaySeconds,omitempty"`
}

func (x *DelayTopicPublishRequest) Reset() {
	*x = DelayTopicPublishRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mq_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelayTopicPublishRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelayTopicPublishRequest) ProtoMessage() {}

func (x *DelayTopicPublishRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mq_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelayTopicPublishRequest.ProtoReflect.Descriptor instead.
func (*DelayTopicPublishRequest) Descriptor() ([]byte, []int) {
	return file_mq_proto_rawDescGZIP(), []int{6}
}

func (x *DelayTopicPublishRequest) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *DelayTopicPublishRequest) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *DelayTopicPublishRequest) GetDelaySeconds() uint64 {
	if x != nil {
		return x.DelaySeconds
	}
	return 0
}

type TopicSubscribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic     string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	QueueName string `protobuf:"bytes,2,opt,name=queueName,proto3" json:"queueName,omitempty"`
}

func (x *TopicSubscribeRequest) Reset() {
	*x = TopicSubscribeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mq_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopicSubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopicSubscribeRequest) ProtoMessage() {}

func (x *TopicSubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mq_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopicSubscribeRequest.ProtoReflect.Descriptor instead.
func (*TopicSubscribeRequest) Descriptor() ([]byte, []int) {
	return file_mq_proto_rawDescGZIP(), []int{7}
}

func (x *TopicSubscribeRequest) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *TopicSubscribeRequest) GetQueueName() string {
	if x != nil {
		return x.QueueName
	}
	return ""
}

var File_mq_proto protoreflect.FileDescriptor

var file_mq_proto_rawDesc = []byte{
	0x0a, 0x08, 0x6d, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x6d, 0x71, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3a, 0x0a, 0x0e, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x71, 0x75, 0x65, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75,
	0x65, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x63, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x61, 0x79,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x71, 0x75, 0x65, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71,
	0x75, 0x65, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x65, 0x6c, 0x61,
	0x79, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c,
	0x64, 0x65, 0x6c, 0x61, 0x79, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x22, 0x28, 0x0a, 0x10,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x71, 0x75, 0x65, 0x75, 0x65, 0x22, 0x4d, 0x0a, 0x11, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x71, 0x75, 0x65, 0x75, 0x65, 0x22, 0x19, 0x0a, 0x07, 0x51, 0x75, 0x65, 0x75, 0x65, 0x49, 0x64,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x3f, 0x0a, 0x13, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x68, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x70, 0x69, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x65, 0x6c, 0x61, 0x79,
	0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x64,
	0x65, 0x6c, 0x61, 0x79, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x22, 0x4b, 0x0a, 0x15, 0x54,
	0x6f, 0x70, 0x69, 0x63, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x1c, 0x0a, 0x09, 0x71, 0x75,
	0x65, 0x75, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x71,
	0x75, 0x65, 0x75, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0x8f, 0x02, 0x0a, 0x02, 0x4d, 0x71, 0x12,
	0x35, 0x0a, 0x07, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x12, 0x12, 0x2e, 0x6d, 0x71, 0x2e,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3f, 0x0a, 0x0c, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x12, 0x17, 0x2e, 0x6d, 0x71, 0x2e, 0x44, 0x65, 0x6c, 0x61,
	0x79, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x38, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x12, 0x14, 0x2e, 0x6d, 0x71, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6d, 0x71, 0x2e,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2a, 0x0a, 0x03, 0x61, 0x63, 0x6b, 0x12, 0x0b, 0x2e, 0x6d, 0x71, 0x2e, 0x51, 0x75,
	0x65, 0x75, 0x65, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x2b, 0x0a,
	0x04, 0x6e, 0x61, 0x63, 0x6b, 0x12, 0x0b, 0x2e, 0x6d, 0x71, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65,
	0x49, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0xa3, 0x02, 0x0a, 0x07, 0x4d,
	0x71, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x3a, 0x0a, 0x07, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x12, 0x17, 0x2e, 0x6d, 0x71, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x12, 0x44, 0x0a, 0x0c, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x12, 0x1c, 0x2e, 0x6d, 0x71, 0x2e, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x54, 0x6f, 0x70,
	0x69, 0x63, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3d, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x19, 0x2e, 0x6d, 0x71, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x15, 0x2e, 0x6d, 0x71, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x03, 0x61, 0x63, 0x6b, 0x12, 0x0b,
	0x2e, 0x6d, 0x71, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x2b, 0x0a, 0x04, 0x6e, 0x61, 0x63, 0x6b, 0x12, 0x0b, 0x2e, 0x6d, 0x71,
	0x2e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x42, 0x32, 0x5a, 0x04, 0x2e, 0x3b, 0x6d, 0x71, 0xca, 0x02, 0x13, 0x44, 0x75, 0x63, 0x43, 0x6e,
	0x7a, 0x6a, 0x5c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x75, 0x73, 0x5c, 0x4d, 0x71, 0xe2, 0x02,
	0x13, 0x44, 0x75, 0x63, 0x43, 0x6e, 0x7a, 0x6a, 0x5c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x75,
	0x73, 0x5c, 0x4d, 0x71, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mq_proto_rawDescOnce sync.Once
	file_mq_proto_rawDescData = file_mq_proto_rawDesc
)

func file_mq_proto_rawDescGZIP() []byte {
	file_mq_proto_rawDescOnce.Do(func() {
		file_mq_proto_rawDescData = protoimpl.X.CompressGZIP(file_mq_proto_rawDescData)
	})
	return file_mq_proto_rawDescData
}

var file_mq_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_mq_proto_goTypes = []interface{}{
	(*PublishRequest)(nil),           // 0: mq.PublishRequest
	(*DelayPublishRequest)(nil),      // 1: mq.DelayPublishRequest
	(*SubscribeRequest)(nil),         // 2: mq.SubscribeRequest
	(*SubscribeResponse)(nil),        // 3: mq.SubscribeResponse
	(*QueueId)(nil),                  // 4: mq.QueueId
	(*TopicPublishRequest)(nil),      // 5: mq.TopicPublishRequest
	(*DelayTopicPublishRequest)(nil), // 6: mq.DelayTopicPublishRequest
	(*TopicSubscribeRequest)(nil),    // 7: mq.TopicSubscribeRequest
	(*empty.Empty)(nil),              // 8: google.protobuf.Empty
}
var file_mq_proto_depIdxs = []int32{
	0,  // 0: mq.Mq.publish:input_type -> mq.PublishRequest
	1,  // 1: mq.Mq.delayPublish:input_type -> mq.DelayPublishRequest
	2,  // 2: mq.Mq.subscribe:input_type -> mq.SubscribeRequest
	4,  // 3: mq.Mq.ack:input_type -> mq.QueueId
	4,  // 4: mq.Mq.nack:input_type -> mq.QueueId
	5,  // 5: mq.MqTopic.publish:input_type -> mq.TopicPublishRequest
	6,  // 6: mq.MqTopic.delayPublish:input_type -> mq.DelayTopicPublishRequest
	7,  // 7: mq.MqTopic.subscribe:input_type -> mq.TopicSubscribeRequest
	4,  // 8: mq.MqTopic.ack:input_type -> mq.QueueId
	4,  // 9: mq.MqTopic.nack:input_type -> mq.QueueId
	8,  // 10: mq.Mq.publish:output_type -> google.protobuf.Empty
	8,  // 11: mq.Mq.delayPublish:output_type -> google.protobuf.Empty
	3,  // 12: mq.Mq.subscribe:output_type -> mq.SubscribeResponse
	8,  // 13: mq.Mq.ack:output_type -> google.protobuf.Empty
	8,  // 14: mq.Mq.nack:output_type -> google.protobuf.Empty
	8,  // 15: mq.MqTopic.publish:output_type -> google.protobuf.Empty
	8,  // 16: mq.MqTopic.delayPublish:output_type -> google.protobuf.Empty
	3,  // 17: mq.MqTopic.subscribe:output_type -> mq.SubscribeResponse
	8,  // 18: mq.MqTopic.ack:output_type -> google.protobuf.Empty
	8,  // 19: mq.MqTopic.nack:output_type -> google.protobuf.Empty
	10, // [10:20] is the sub-list for method output_type
	0,  // [0:10] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_mq_proto_init() }
func file_mq_proto_init() {
	if File_mq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishRequest); i {
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
		file_mq_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelayPublishRequest); i {
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
		file_mq_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeRequest); i {
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
		file_mq_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeResponse); i {
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
		file_mq_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueueId); i {
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
		file_mq_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopicPublishRequest); i {
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
		file_mq_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelayTopicPublishRequest); i {
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
		file_mq_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopicSubscribeRequest); i {
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
			RawDescriptor: file_mq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_mq_proto_goTypes,
		DependencyIndexes: file_mq_proto_depIdxs,
		MessageInfos:      file_mq_proto_msgTypes,
	}.Build()
	File_mq_proto = out.File
	file_mq_proto_rawDesc = nil
	file_mq_proto_goTypes = nil
	file_mq_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: api.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
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

type Status int32

const (
	Status_OK    Status = 0
	Status_ERROR Status = 1
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "OK",
		1: "ERROR",
	}
	Status_value = map[string]int32{
		"OK":    0,
		"ERROR": 1,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_api_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version     string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	PeerId      string `protobuf:"bytes,3,opt,name=peerId,proto3" json:"peerId,omitempty"`
	Os          string `protobuf:"bytes,4,opt,name=os,proto3" json:"os,omitempty"`
	NetworkName string `protobuf:"bytes,5,opt,name=networkName,proto3" json:"networkName,omitempty"`
	Type        string `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RegisterRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *RegisterRequest) GetPeerId() string {
	if x != nil {
		return x.PeerId
	}
	return ""
}

func (x *RegisterRequest) GetOs() string {
	if x != nil {
		return x.Os
	}
	return ""
}

func (x *RegisterRequest) GetNetworkName() string {
	if x != nil {
		return x.NetworkName
	}
	return ""
}

func (x *RegisterRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type RegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token   string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Status  Status `protobuf:"varint,2,opt,name=status,proto3,enum=proto.Status" json:"status,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *RegisterResponse) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_OK
}

func (x *RegisterResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type HeartbeatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HeartbeatRequest) Reset() {
	*x = HeartbeatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartbeatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatRequest) ProtoMessage() {}

func (x *HeartbeatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatRequest.ProtoReflect.Descriptor instead.
func (*HeartbeatRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

type HeartbeatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  Status `protobuf:"varint,1,opt,name=status,proto3,enum=proto.Status" json:"status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HeartbeatResponse) Reset() {
	*x = HeartbeatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartbeatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatResponse) ProtoMessage() {}

func (x *HeartbeatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatResponse.ProtoReflect.Descriptor instead.
func (*HeartbeatResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{3}
}

func (x *HeartbeatResponse) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_OK
}

func (x *HeartbeatResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DefaultResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  Status `protobuf:"varint,1,opt,name=status,proto3,enum=proto.Status" json:"status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DefaultResponse) Reset() {
	*x = DefaultResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DefaultResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DefaultResponse) ProtoMessage() {}

func (x *DefaultResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DefaultResponse.ProtoReflect.Descriptor instead.
func (*DefaultResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{4}
}

func (x *DefaultResponse) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_OK
}

func (x *DefaultResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ChainHeadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TipsetHeight int64             `protobuf:"varint,1,opt,name=tipsetHeight,proto3" json:"tipsetHeight,omitempty"`
	Blocks       []*ChainHeadBlock `protobuf:"bytes,2,rep,name=blocks,proto3" json:"blocks,omitempty"`
	ReceivedAt   string            `protobuf:"bytes,3,opt,name=receivedAt,proto3" json:"receivedAt,omitempty"`
}

func (x *ChainHeadRequest) Reset() {
	*x = ChainHeadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChainHeadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChainHeadRequest) ProtoMessage() {}

func (x *ChainHeadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChainHeadRequest.ProtoReflect.Descriptor instead.
func (*ChainHeadRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{5}
}

func (x *ChainHeadRequest) GetTipsetHeight() int64 {
	if x != nil {
		return x.TipsetHeight
	}
	return 0
}

func (x *ChainHeadRequest) GetBlocks() []*ChainHeadBlock {
	if x != nil {
		return x.Blocks
	}
	return nil
}

func (x *ChainHeadRequest) GetReceivedAt() string {
	if x != nil {
		return x.ReceivedAt
	}
	return ""
}

type ChainHeadBlock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cid              string `protobuf:"bytes,1,opt,name=cid,proto3" json:"cid,omitempty"`
	ParentWeight     string `protobuf:"bytes,2,opt,name=parentWeight,proto3" json:"parentWeight,omitempty"`
	Miner            string `protobuf:"bytes,3,opt,name=miner,proto3" json:"miner,omitempty"`
	NumberOfMessages int64  `protobuf:"varint,4,opt,name=numberOfMessages,proto3" json:"numberOfMessages,omitempty"`
	Timestamp        uint64 `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	CurrentWeight    string `protobuf:"bytes,6,opt,name=currentWeight,proto3" json:"currentWeight,omitempty"`
}

func (x *ChainHeadBlock) Reset() {
	*x = ChainHeadBlock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChainHeadBlock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChainHeadBlock) ProtoMessage() {}

func (x *ChainHeadBlock) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChainHeadBlock.ProtoReflect.Descriptor instead.
func (*ChainHeadBlock) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{6}
}

func (x *ChainHeadBlock) GetCid() string {
	if x != nil {
		return x.Cid
	}
	return ""
}

func (x *ChainHeadBlock) GetParentWeight() string {
	if x != nil {
		return x.ParentWeight
	}
	return ""
}

func (x *ChainHeadBlock) GetMiner() string {
	if x != nil {
		return x.Miner
	}
	return ""
}

func (x *ChainHeadBlock) GetNumberOfMessages() int64 {
	if x != nil {
		return x.NumberOfMessages
	}
	return 0
}

func (x *ChainHeadBlock) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *ChainHeadBlock) GetCurrentWeight() string {
	if x != nil {
		return x.CurrentWeight
	}
	return ""
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x9d, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x65, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x65, 0x65, 0x72, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02,
	0x6f, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6f, 0x73, 0x12, 0x20, 0x0a, 0x0b,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x22, 0x69, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x25, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x12, 0x0a,
	0x10, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x54, 0x0a, 0x11, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x52, 0x0a, 0x0f, 0x44, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x85, 0x01, 0x0a, 0x10,
	0x43, 0x68, 0x61, 0x69, 0x6e, 0x48, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x22, 0x0a, 0x0c, 0x74, 0x69, 0x70, 0x73, 0x65, 0x74, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x74, 0x69, 0x70, 0x73, 0x65, 0x74, 0x48, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x12, 0x2d, 0x0a, 0x06, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x61,
	0x69, 0x6e, 0x48, 0x65, 0x61, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x06, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x64, 0x41, 0x74, 0x22, 0xcc, 0x01, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x48, 0x65, 0x61,
	0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x6d, 0x69, 0x6e, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x69, 0x6e,
	0x65, 0x72, 0x12, 0x2a, 0x0a, 0x10, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x1c,
	0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x24, 0x0a, 0x0d,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x57, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x2a, 0x1b, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x06, 0x0a, 0x02,
	0x4f, 0x4b, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x01, 0x32,
	0xcb, 0x01, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x73, 0x74, 0x61, 0x74, 0x73, 0x12, 0x3d, 0x0a, 0x08,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x09, 0x48,
	0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62,
	0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a,
	0x09, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x48, 0x65, 0x61, 0x64, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x48, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x36, 0x5a,
	0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69,
	0x74, 0x61, 0x6c, 0x2d, 0x6d, 0x6f, 0x62, 0x2d, 0x66, 0x69, 0x6c, 0x65, 0x63, 0x6f, 0x69, 0x6e,
	0x2f, 0x66, 0x69, 0x6c, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_proto_goTypes = []interface{}{
	(Status)(0),               // 0: proto.Status
	(*RegisterRequest)(nil),   // 1: proto.RegisterRequest
	(*RegisterResponse)(nil),  // 2: proto.RegisterResponse
	(*HeartbeatRequest)(nil),  // 3: proto.HeartbeatRequest
	(*HeartbeatResponse)(nil), // 4: proto.HeartbeatResponse
	(*DefaultResponse)(nil),   // 5: proto.DefaultResponse
	(*ChainHeadRequest)(nil),  // 6: proto.ChainHeadRequest
	(*ChainHeadBlock)(nil),    // 7: proto.ChainHeadBlock
}
var file_api_proto_depIdxs = []int32{
	0, // 0: proto.RegisterResponse.status:type_name -> proto.Status
	0, // 1: proto.HeartbeatResponse.status:type_name -> proto.Status
	0, // 2: proto.DefaultResponse.status:type_name -> proto.Status
	7, // 3: proto.ChainHeadRequest.blocks:type_name -> proto.ChainHeadBlock
	1, // 4: proto.Filstats.Register:input_type -> proto.RegisterRequest
	3, // 5: proto.Filstats.Heartbeat:input_type -> proto.HeartbeatRequest
	6, // 6: proto.Filstats.ChainHead:input_type -> proto.ChainHeadRequest
	2, // 7: proto.Filstats.Register:output_type -> proto.RegisterResponse
	4, // 8: proto.Filstats.Heartbeat:output_type -> proto.HeartbeatResponse
	5, // 9: proto.Filstats.ChainHead:output_type -> proto.DefaultResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
		file_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResponse); i {
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
		file_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartbeatRequest); i {
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
		file_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartbeatResponse); i {
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
		file_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DefaultResponse); i {
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
		file_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChainHeadRequest); i {
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
		file_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChainHeadBlock); i {
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
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		EnumInfos:         file_api_proto_enumTypes,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}

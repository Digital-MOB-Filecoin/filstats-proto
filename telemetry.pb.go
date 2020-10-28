// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: telemetry.proto

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

type PeersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Peers int64 `protobuf:"varint,1,opt,name=peers,proto3" json:"peers,omitempty"`
}

func (x *PeersRequest) Reset() {
	*x = PeersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_telemetry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeersRequest) ProtoMessage() {}

func (x *PeersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_telemetry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeersRequest.ProtoReflect.Descriptor instead.
func (*PeersRequest) Descriptor() ([]byte, []int) {
	return file_telemetry_proto_rawDescGZIP(), []int{0}
}

func (x *PeersRequest) GetPeers() int64 {
	if x != nil {
		return x.Peers
	}
	return 0
}

type MpoolSizeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size int64 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *MpoolSizeRequest) Reset() {
	*x = MpoolSizeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_telemetry_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MpoolSizeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MpoolSizeRequest) ProtoMessage() {}

func (x *MpoolSizeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_telemetry_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MpoolSizeRequest.ProtoReflect.Descriptor instead.
func (*MpoolSizeRequest) Descriptor() ([]byte, []int) {
	return file_telemetry_proto_rawDescGZIP(), []int{1}
}

func (x *MpoolSizeRequest) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type SyncingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsSyncing bool `protobuf:"varint,1,opt,name=isSyncing,proto3" json:"isSyncing,omitempty"`
}

func (x *SyncingRequest) Reset() {
	*x = SyncingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_telemetry_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncingRequest) ProtoMessage() {}

func (x *SyncingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_telemetry_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncingRequest.ProtoReflect.Descriptor instead.
func (*SyncingRequest) Descriptor() ([]byte, []int) {
	return file_telemetry_proto_rawDescGZIP(), []int{2}
}

func (x *SyncingRequest) GetIsSyncing() bool {
	if x != nil {
		return x.IsSyncing
	}
	return false
}

type NSPRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Power string `protobuf:"bytes,1,opt,name=power,proto3" json:"power,omitempty"`
}

func (x *NSPRequest) Reset() {
	*x = NSPRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_telemetry_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NSPRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NSPRequest) ProtoMessage() {}

func (x *NSPRequest) ProtoReflect() protoreflect.Message {
	mi := &file_telemetry_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NSPRequest.ProtoReflect.Descriptor instead.
func (*NSPRequest) Descriptor() ([]byte, []int) {
	return file_telemetry_proto_rawDescGZIP(), []int{3}
}

func (x *NSPRequest) GetPower() string {
	if x != nil {
		return x.Power
	}
	return ""
}

var File_telemetry_proto protoreflect.FileDescriptor

var file_telemetry_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x24, 0x0a, 0x0c, 0x50, 0x65, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x22, 0x26, 0x0a, 0x10, 0x4d, 0x70, 0x6f,
	0x6f, 0x6c, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x22, 0x2e, 0x0a, 0x0e, 0x53, 0x79, 0x6e, 0x63, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x53, 0x79, 0x6e, 0x63, 0x69, 0x6e, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x53, 0x79, 0x6e, 0x63, 0x69, 0x6e,
	0x67, 0x22, 0x22, 0x0a, 0x0a, 0x4e, 0x53, 0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x70, 0x6f, 0x77, 0x65, 0x72, 0x32, 0x83, 0x02, 0x0a, 0x09, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x65,
	0x74, 0x72, 0x79, 0x12, 0x36, 0x0a, 0x05, 0x50, 0x65, 0x65, 0x72, 0x73, 0x12, 0x13, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x09, 0x4d,
	0x70, 0x6f, 0x6f, 0x6c, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4d, 0x70, 0x6f, 0x6f, 0x6c, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x07, 0x53,
	0x79, 0x6e, 0x63, 0x69, 0x6e, 0x67, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53,
	0x79, 0x6e, 0x63, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x13, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x11,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x53, 0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x36, 0x5a, 0x34, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61,
	0x6c, 0x2d, 0x6d, 0x6f, 0x62, 0x2d, 0x66, 0x69, 0x6c, 0x65, 0x63, 0x6f, 0x69, 0x6e, 0x2f, 0x66,
	0x69, 0x6c, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_telemetry_proto_rawDescOnce sync.Once
	file_telemetry_proto_rawDescData = file_telemetry_proto_rawDesc
)

func file_telemetry_proto_rawDescGZIP() []byte {
	file_telemetry_proto_rawDescOnce.Do(func() {
		file_telemetry_proto_rawDescData = protoimpl.X.CompressGZIP(file_telemetry_proto_rawDescData)
	})
	return file_telemetry_proto_rawDescData
}

var file_telemetry_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_telemetry_proto_goTypes = []interface{}{
	(*PeersRequest)(nil),     // 0: proto.PeersRequest
	(*MpoolSizeRequest)(nil), // 1: proto.MpoolSizeRequest
	(*SyncingRequest)(nil),   // 2: proto.SyncingRequest
	(*NSPRequest)(nil),       // 3: proto.NSPRequest
	(*DefaultResponse)(nil),  // 4: proto.DefaultResponse
}
var file_telemetry_proto_depIdxs = []int32{
	0, // 0: proto.Telemetry.Peers:input_type -> proto.PeersRequest
	1, // 1: proto.Telemetry.MpoolSize:input_type -> proto.MpoolSizeRequest
	2, // 2: proto.Telemetry.Syncing:input_type -> proto.SyncingRequest
	3, // 3: proto.Telemetry.NetworkStoragePower:input_type -> proto.NSPRequest
	4, // 4: proto.Telemetry.Peers:output_type -> proto.DefaultResponse
	4, // 5: proto.Telemetry.MpoolSize:output_type -> proto.DefaultResponse
	4, // 6: proto.Telemetry.Syncing:output_type -> proto.DefaultResponse
	4, // 7: proto.Telemetry.NetworkStoragePower:output_type -> proto.DefaultResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_telemetry_proto_init() }
func file_telemetry_proto_init() {
	if File_telemetry_proto != nil {
		return
	}
	file_api_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_telemetry_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeersRequest); i {
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
		file_telemetry_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MpoolSizeRequest); i {
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
		file_telemetry_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncingRequest); i {
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
		file_telemetry_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NSPRequest); i {
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
			RawDescriptor: file_telemetry_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_telemetry_proto_goTypes,
		DependencyIndexes: file_telemetry_proto_depIdxs,
		MessageInfos:      file_telemetry_proto_msgTypes,
	}.Build()
	File_telemetry_proto = out.File
	file_telemetry_proto_rawDesc = nil
	file_telemetry_proto_goTypes = nil
	file_telemetry_proto_depIdxs = nil
}

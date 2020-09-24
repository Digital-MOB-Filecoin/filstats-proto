// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: telemetry.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

var File_telemetry_proto protoreflect.FileDescriptor

var file_telemetry_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x24, 0x0a, 0x0c, 0x50, 0x65, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x32, 0x43, 0x0a, 0x09, 0x54, 0x65, 0x6c,
	0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x12, 0x36, 0x0a, 0x05, 0x50, 0x65, 0x65, 0x72, 0x73, 0x12,
	0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x36,
	0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67,
	0x69, 0x74, 0x61, 0x6c, 0x2d, 0x6d, 0x6f, 0x62, 0x2d, 0x66, 0x69, 0x6c, 0x65, 0x63, 0x6f, 0x69,
	0x6e, 0x2f, 0x66, 0x69, 0x6c, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_telemetry_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_telemetry_proto_goTypes = []interface{}{
	(*PeersRequest)(nil),    // 0: proto.PeersRequest
	(*DefaultResponse)(nil), // 1: proto.DefaultResponse
}
var file_telemetry_proto_depIdxs = []int32{
	0, // 0: proto.Telemetry.Peers:input_type -> proto.PeersRequest
	1, // 1: proto.Telemetry.Peers:output_type -> proto.DefaultResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_telemetry_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TelemetryClient is the client API for Telemetry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TelemetryClient interface {
	Peers(ctx context.Context, in *PeersRequest, opts ...grpc.CallOption) (*DefaultResponse, error)
}

type telemetryClient struct {
	cc grpc.ClientConnInterface
}

func NewTelemetryClient(cc grpc.ClientConnInterface) TelemetryClient {
	return &telemetryClient{cc}
}

func (c *telemetryClient) Peers(ctx context.Context, in *PeersRequest, opts ...grpc.CallOption) (*DefaultResponse, error) {
	out := new(DefaultResponse)
	err := c.cc.Invoke(ctx, "/proto.Telemetry/Peers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TelemetryServer is the server API for Telemetry service.
type TelemetryServer interface {
	Peers(context.Context, *PeersRequest) (*DefaultResponse, error)
}

// UnimplementedTelemetryServer can be embedded to have forward compatible implementations.
type UnimplementedTelemetryServer struct {
}

func (*UnimplementedTelemetryServer) Peers(context.Context, *PeersRequest) (*DefaultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Peers not implemented")
}

func RegisterTelemetryServer(s *grpc.Server, srv TelemetryServer) {
	s.RegisterService(&_Telemetry_serviceDesc, srv)
}

func _Telemetry_Peers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TelemetryServer).Peers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Telemetry/Peers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TelemetryServer).Peers(ctx, req.(*PeersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Telemetry_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Telemetry",
	HandlerType: (*TelemetryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Peers",
			Handler:    _Telemetry_Peers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "telemetry.proto",
}

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// FilstatsClient is the client API for Filstats service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FilstatsClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	Heartbeat(ctx context.Context, in *HeartbeatRequest, opts ...grpc.CallOption) (*HeartbeatResponse, error)
	ChainHead(ctx context.Context, in *ChainHeadRequest, opts ...grpc.CallOption) (*DefaultResponse, error)
}

type filstatsClient struct {
	cc grpc.ClientConnInterface
}

func NewFilstatsClient(cc grpc.ClientConnInterface) FilstatsClient {
	return &filstatsClient{cc}
}

func (c *filstatsClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/proto.Filstats/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filstatsClient) Heartbeat(ctx context.Context, in *HeartbeatRequest, opts ...grpc.CallOption) (*HeartbeatResponse, error) {
	out := new(HeartbeatResponse)
	err := c.cc.Invoke(ctx, "/proto.Filstats/Heartbeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filstatsClient) ChainHead(ctx context.Context, in *ChainHeadRequest, opts ...grpc.CallOption) (*DefaultResponse, error) {
	out := new(DefaultResponse)
	err := c.cc.Invoke(ctx, "/proto.Filstats/ChainHead", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FilstatsServer is the server API for Filstats service.
// All implementations must embed UnimplementedFilstatsServer
// for forward compatibility
type FilstatsServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	Heartbeat(context.Context, *HeartbeatRequest) (*HeartbeatResponse, error)
	ChainHead(context.Context, *ChainHeadRequest) (*DefaultResponse, error)
	mustEmbedUnimplementedFilstatsServer()
}

// UnimplementedFilstatsServer must be embedded to have forward compatible implementations.
type UnimplementedFilstatsServer struct {
}

func (UnimplementedFilstatsServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedFilstatsServer) Heartbeat(context.Context, *HeartbeatRequest) (*HeartbeatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Heartbeat not implemented")
}
func (UnimplementedFilstatsServer) ChainHead(context.Context, *ChainHeadRequest) (*DefaultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChainHead not implemented")
}
func (UnimplementedFilstatsServer) mustEmbedUnimplementedFilstatsServer() {}

// UnsafeFilstatsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FilstatsServer will
// result in compilation errors.
type UnsafeFilstatsServer interface {
	mustEmbedUnimplementedFilstatsServer()
}

func RegisterFilstatsServer(s grpc.ServiceRegistrar, srv FilstatsServer) {
	s.RegisterService(&_Filstats_serviceDesc, srv)
}

func _Filstats_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilstatsServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Filstats/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilstatsServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Filstats_Heartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartbeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilstatsServer).Heartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Filstats/Heartbeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilstatsServer).Heartbeat(ctx, req.(*HeartbeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Filstats_ChainHead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChainHeadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilstatsServer).ChainHead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Filstats/ChainHead",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilstatsServer).ChainHead(ctx, req.(*ChainHeadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Filstats_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Filstats",
	HandlerType: (*FilstatsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Filstats_Register_Handler,
		},
		{
			MethodName: "Heartbeat",
			Handler:    _Filstats_Heartbeat_Handler,
		},
		{
			MethodName: "ChainHead",
			Handler:    _Filstats_ChainHead_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

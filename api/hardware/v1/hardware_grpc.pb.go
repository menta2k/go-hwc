// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: hardware/v1/hardware.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// HardwareClient is the client API for Hardware service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HardwareClient interface {
	// Sends a greeting
	Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendReply, error)
	GetHardware(ctx context.Context, in *GetHardwareRequest, opts ...grpc.CallOption) (*GetHardwareReply, error)
}

type hardwareClient struct {
	cc grpc.ClientConnInterface
}

func NewHardwareClient(cc grpc.ClientConnInterface) HardwareClient {
	return &hardwareClient{cc}
}

func (c *hardwareClient) Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendReply, error) {
	out := new(SendReply)
	err := c.cc.Invoke(ctx, "/hardware.v1.Hardware/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hardwareClient) GetHardware(ctx context.Context, in *GetHardwareRequest, opts ...grpc.CallOption) (*GetHardwareReply, error) {
	out := new(GetHardwareReply)
	err := c.cc.Invoke(ctx, "/hardware.v1.Hardware/GetHardware", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HardwareServer is the server API for Hardware service.
// All implementations must embed UnimplementedHardwareServer
// for forward compatibility
type HardwareServer interface {
	// Sends a greeting
	Send(context.Context, *SendRequest) (*SendReply, error)
	GetHardware(context.Context, *GetHardwareRequest) (*GetHardwareReply, error)
	mustEmbedUnimplementedHardwareServer()
}

// UnimplementedHardwareServer must be embedded to have forward compatible implementations.
type UnimplementedHardwareServer struct {
}

func (UnimplementedHardwareServer) Send(context.Context, *SendRequest) (*SendReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedHardwareServer) GetHardware(context.Context, *GetHardwareRequest) (*GetHardwareReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHardware not implemented")
}
func (UnimplementedHardwareServer) mustEmbedUnimplementedHardwareServer() {}

// UnsafeHardwareServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HardwareServer will
// result in compilation errors.
type UnsafeHardwareServer interface {
	mustEmbedUnimplementedHardwareServer()
}

func RegisterHardwareServer(s grpc.ServiceRegistrar, srv HardwareServer) {
	s.RegisterService(&Hardware_ServiceDesc, srv)
}

func _Hardware_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HardwareServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hardware.v1.Hardware/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HardwareServer).Send(ctx, req.(*SendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hardware_GetHardware_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHardwareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HardwareServer).GetHardware(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hardware.v1.Hardware/GetHardware",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HardwareServer).GetHardware(ctx, req.(*GetHardwareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Hardware_ServiceDesc is the grpc.ServiceDesc for Hardware service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Hardware_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hardware.v1.Hardware",
	HandlerType: (*HardwareServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _Hardware_Send_Handler,
		},
		{
			MethodName: "GetHardware",
			Handler:    _Hardware_GetHardware_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hardware/v1/hardware.proto",
}
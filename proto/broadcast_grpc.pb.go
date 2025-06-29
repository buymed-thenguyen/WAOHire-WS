// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.5
// source: proto/broadcast.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Broadcaster_BroadcastUserAnswered_FullMethodName = "/ws.Broadcaster/BroadcastUserAnswered"
	Broadcaster_BroadcastUserJoined_FullMethodName   = "/ws.Broadcaster/BroadcastUserJoined"
	Broadcaster_BroadcastUserLeaved_FullMethodName   = "/ws.Broadcaster/BroadcastUserLeaved"
	Broadcaster_BroadcastStartSession_FullMethodName = "/ws.Broadcaster/BroadcastStartSession"
)

// BroadcasterClient is the client API for Broadcaster service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BroadcasterClient interface {
	BroadcastUserAnswered(ctx context.Context, in *UserAnsweredRequest, opts ...grpc.CallOption) (*BroadcastResponse, error)
	BroadcastUserJoined(ctx context.Context, in *UserJoinedRequest, opts ...grpc.CallOption) (*BroadcastResponse, error)
	BroadcastUserLeaved(ctx context.Context, in *UserLeavedRequest, opts ...grpc.CallOption) (*BroadcastResponse, error)
	BroadcastStartSession(ctx context.Context, in *StartSessionRequest, opts ...grpc.CallOption) (*BroadcastResponse, error)
}

type broadcasterClient struct {
	cc grpc.ClientConnInterface
}

func NewBroadcasterClient(cc grpc.ClientConnInterface) BroadcasterClient {
	return &broadcasterClient{cc}
}

func (c *broadcasterClient) BroadcastUserAnswered(ctx context.Context, in *UserAnsweredRequest, opts ...grpc.CallOption) (*BroadcastResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BroadcastResponse)
	err := c.cc.Invoke(ctx, Broadcaster_BroadcastUserAnswered_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *broadcasterClient) BroadcastUserJoined(ctx context.Context, in *UserJoinedRequest, opts ...grpc.CallOption) (*BroadcastResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BroadcastResponse)
	err := c.cc.Invoke(ctx, Broadcaster_BroadcastUserJoined_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *broadcasterClient) BroadcastUserLeaved(ctx context.Context, in *UserLeavedRequest, opts ...grpc.CallOption) (*BroadcastResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BroadcastResponse)
	err := c.cc.Invoke(ctx, Broadcaster_BroadcastUserLeaved_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *broadcasterClient) BroadcastStartSession(ctx context.Context, in *StartSessionRequest, opts ...grpc.CallOption) (*BroadcastResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BroadcastResponse)
	err := c.cc.Invoke(ctx, Broadcaster_BroadcastStartSession_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BroadcasterServer is the server API for Broadcaster service.
// All implementations must embed UnimplementedBroadcasterServer
// for forward compatibility.
type BroadcasterServer interface {
	BroadcastUserAnswered(context.Context, *UserAnsweredRequest) (*BroadcastResponse, error)
	BroadcastUserJoined(context.Context, *UserJoinedRequest) (*BroadcastResponse, error)
	BroadcastUserLeaved(context.Context, *UserLeavedRequest) (*BroadcastResponse, error)
	BroadcastStartSession(context.Context, *StartSessionRequest) (*BroadcastResponse, error)
	mustEmbedUnimplementedBroadcasterServer()
}

// UnimplementedBroadcasterServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBroadcasterServer struct{}

func (UnimplementedBroadcasterServer) BroadcastUserAnswered(context.Context, *UserAnsweredRequest) (*BroadcastResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BroadcastUserAnswered not implemented")
}
func (UnimplementedBroadcasterServer) BroadcastUserJoined(context.Context, *UserJoinedRequest) (*BroadcastResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BroadcastUserJoined not implemented")
}
func (UnimplementedBroadcasterServer) BroadcastUserLeaved(context.Context, *UserLeavedRequest) (*BroadcastResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BroadcastUserLeaved not implemented")
}
func (UnimplementedBroadcasterServer) BroadcastStartSession(context.Context, *StartSessionRequest) (*BroadcastResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BroadcastStartSession not implemented")
}
func (UnimplementedBroadcasterServer) mustEmbedUnimplementedBroadcasterServer() {}
func (UnimplementedBroadcasterServer) testEmbeddedByValue()                     {}

// UnsafeBroadcasterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BroadcasterServer will
// result in compilation errors.
type UnsafeBroadcasterServer interface {
	mustEmbedUnimplementedBroadcasterServer()
}

func RegisterBroadcasterServer(s grpc.ServiceRegistrar, srv BroadcasterServer) {
	// If the following call pancis, it indicates UnimplementedBroadcasterServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Broadcaster_ServiceDesc, srv)
}

func _Broadcaster_BroadcastUserAnswered_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAnsweredRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BroadcasterServer).BroadcastUserAnswered(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Broadcaster_BroadcastUserAnswered_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BroadcasterServer).BroadcastUserAnswered(ctx, req.(*UserAnsweredRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Broadcaster_BroadcastUserJoined_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserJoinedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BroadcasterServer).BroadcastUserJoined(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Broadcaster_BroadcastUserJoined_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BroadcasterServer).BroadcastUserJoined(ctx, req.(*UserJoinedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Broadcaster_BroadcastUserLeaved_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLeavedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BroadcasterServer).BroadcastUserLeaved(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Broadcaster_BroadcastUserLeaved_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BroadcasterServer).BroadcastUserLeaved(ctx, req.(*UserLeavedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Broadcaster_BroadcastStartSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BroadcasterServer).BroadcastStartSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Broadcaster_BroadcastStartSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BroadcasterServer).BroadcastStartSession(ctx, req.(*StartSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Broadcaster_ServiceDesc is the grpc.ServiceDesc for Broadcaster service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Broadcaster_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ws.Broadcaster",
	HandlerType: (*BroadcasterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BroadcastUserAnswered",
			Handler:    _Broadcaster_BroadcastUserAnswered_Handler,
		},
		{
			MethodName: "BroadcastUserJoined",
			Handler:    _Broadcaster_BroadcastUserJoined_Handler,
		},
		{
			MethodName: "BroadcastUserLeaved",
			Handler:    _Broadcaster_BroadcastUserLeaved_Handler,
		},
		{
			MethodName: "BroadcastStartSession",
			Handler:    _Broadcaster_BroadcastStartSession_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/broadcast.proto",
}

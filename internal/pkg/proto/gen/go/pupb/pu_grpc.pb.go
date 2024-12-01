// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.0
// source: pu/pu.proto

package pupb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ParcelsUser_GetParcel_FullMethodName          = "/pupb.ParcelsUser/GetParcel"
	ParcelsUser_GetCheckpoints_FullMethodName     = "/pupb.ParcelsUser/GetCheckpoints"
	ParcelsUser_AddSubscription_FullMethodName    = "/pupb.ParcelsUser/AddSubscription"
	ParcelsUser_DeleteSubscription_FullMethodName = "/pupb.ParcelsUser/DeleteSubscription"
)

// ParcelsUserClient is the client API for ParcelsUser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ParcelsUserClient interface {
	GetParcel(ctx context.Context, in *GetParcelRequest, opts ...grpc.CallOption) (*GetParcelResponse, error)
	GetCheckpoints(ctx context.Context, in *GetCheckpointsRequest, opts ...grpc.CallOption) (*GetCheckpointsResponse, error)
	AddSubscription(ctx context.Context, in *AddSubscriptionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteSubscription(ctx context.Context, in *DeleteSubscriptionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type parcelsUserClient struct {
	cc grpc.ClientConnInterface
}

func NewParcelsUserClient(cc grpc.ClientConnInterface) ParcelsUserClient {
	return &parcelsUserClient{cc}
}

func (c *parcelsUserClient) GetParcel(ctx context.Context, in *GetParcelRequest, opts ...grpc.CallOption) (*GetParcelResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetParcelResponse)
	err := c.cc.Invoke(ctx, ParcelsUser_GetParcel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *parcelsUserClient) GetCheckpoints(ctx context.Context, in *GetCheckpointsRequest, opts ...grpc.CallOption) (*GetCheckpointsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCheckpointsResponse)
	err := c.cc.Invoke(ctx, ParcelsUser_GetCheckpoints_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *parcelsUserClient) AddSubscription(ctx context.Context, in *AddSubscriptionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ParcelsUser_AddSubscription_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *parcelsUserClient) DeleteSubscription(ctx context.Context, in *DeleteSubscriptionRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ParcelsUser_DeleteSubscription_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ParcelsUserServer is the server API for ParcelsUser service.
// All implementations must embed UnimplementedParcelsUserServer
// for forward compatibility.
type ParcelsUserServer interface {
	GetParcel(context.Context, *GetParcelRequest) (*GetParcelResponse, error)
	GetCheckpoints(context.Context, *GetCheckpointsRequest) (*GetCheckpointsResponse, error)
	AddSubscription(context.Context, *AddSubscriptionRequest) (*emptypb.Empty, error)
	DeleteSubscription(context.Context, *DeleteSubscriptionRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedParcelsUserServer()
}

// UnimplementedParcelsUserServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedParcelsUserServer struct{}

func (UnimplementedParcelsUserServer) GetParcel(context.Context, *GetParcelRequest) (*GetParcelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetParcel not implemented")
}
func (UnimplementedParcelsUserServer) GetCheckpoints(context.Context, *GetCheckpointsRequest) (*GetCheckpointsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCheckpoints not implemented")
}
func (UnimplementedParcelsUserServer) AddSubscription(context.Context, *AddSubscriptionRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSubscription not implemented")
}
func (UnimplementedParcelsUserServer) DeleteSubscription(context.Context, *DeleteSubscriptionRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSubscription not implemented")
}
func (UnimplementedParcelsUserServer) mustEmbedUnimplementedParcelsUserServer() {}
func (UnimplementedParcelsUserServer) testEmbeddedByValue()                     {}

// UnsafeParcelsUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ParcelsUserServer will
// result in compilation errors.
type UnsafeParcelsUserServer interface {
	mustEmbedUnimplementedParcelsUserServer()
}

func RegisterParcelsUserServer(s grpc.ServiceRegistrar, srv ParcelsUserServer) {
	// If the following call pancis, it indicates UnimplementedParcelsUserServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ParcelsUser_ServiceDesc, srv)
}

func _ParcelsUser_GetParcel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetParcelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParcelsUserServer).GetParcel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ParcelsUser_GetParcel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParcelsUserServer).GetParcel(ctx, req.(*GetParcelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParcelsUser_GetCheckpoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCheckpointsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParcelsUserServer).GetCheckpoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ParcelsUser_GetCheckpoints_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParcelsUserServer).GetCheckpoints(ctx, req.(*GetCheckpointsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParcelsUser_AddSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSubscriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParcelsUserServer).AddSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ParcelsUser_AddSubscription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParcelsUserServer).AddSubscription(ctx, req.(*AddSubscriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParcelsUser_DeleteSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSubscriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParcelsUserServer).DeleteSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ParcelsUser_DeleteSubscription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParcelsUserServer).DeleteSubscription(ctx, req.(*DeleteSubscriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ParcelsUser_ServiceDesc is the grpc.ServiceDesc for ParcelsUser service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ParcelsUser_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pupb.ParcelsUser",
	HandlerType: (*ParcelsUserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetParcel",
			Handler:    _ParcelsUser_GetParcel_Handler,
		},
		{
			MethodName: "GetCheckpoints",
			Handler:    _ParcelsUser_GetCheckpoints_Handler,
		},
		{
			MethodName: "AddSubscription",
			Handler:    _ParcelsUser_AddSubscription_Handler,
		},
		{
			MethodName: "DeleteSubscription",
			Handler:    _ParcelsUser_DeleteSubscription_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pu/pu.proto",
}

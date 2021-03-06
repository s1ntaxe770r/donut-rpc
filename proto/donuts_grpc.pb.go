// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package donut_rpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DonutShopClient is the client API for DonutShop service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DonutShopClient interface {
	GetDonut(ctx context.Context, in *DonutRequest, opts ...grpc.CallOption) (*Donut, error)
	GetDonuts(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Donuts, error)
	MakeDonut(ctx context.Context, in *Donut, opts ...grpc.CallOption) (*DonutRequest, error)
	GetVersion(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Version, error)
}

type donutShopClient struct {
	cc grpc.ClientConnInterface
}

func NewDonutShopClient(cc grpc.ClientConnInterface) DonutShopClient {
	return &donutShopClient{cc}
}

func (c *donutShopClient) GetDonut(ctx context.Context, in *DonutRequest, opts ...grpc.CallOption) (*Donut, error) {
	out := new(Donut)
	err := c.cc.Invoke(ctx, "/DonutShop/GetDonut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *donutShopClient) GetDonuts(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Donuts, error) {
	out := new(Donuts)
	err := c.cc.Invoke(ctx, "/DonutShop/GetDonuts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *donutShopClient) MakeDonut(ctx context.Context, in *Donut, opts ...grpc.CallOption) (*DonutRequest, error) {
	out := new(DonutRequest)
	err := c.cc.Invoke(ctx, "/DonutShop/MakeDonut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *donutShopClient) GetVersion(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Version, error) {
	out := new(Version)
	err := c.cc.Invoke(ctx, "/DonutShop/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DonutShopServer is the server API for DonutShop service.
// All implementations must embed UnimplementedDonutShopServer
// for forward compatibility
type DonutShopServer interface {
	GetDonut(context.Context, *DonutRequest) (*Donut, error)
	GetDonuts(context.Context, *emptypb.Empty) (*Donuts, error)
	MakeDonut(context.Context, *Donut) (*DonutRequest, error)
	GetVersion(context.Context, *emptypb.Empty) (*Version, error)
	mustEmbedUnimplementedDonutShopServer()
}

// UnimplementedDonutShopServer must be embedded to have forward compatible implementations.
type UnimplementedDonutShopServer struct {
}

func (UnimplementedDonutShopServer) GetDonut(context.Context, *DonutRequest) (*Donut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDonut not implemented")
}
func (UnimplementedDonutShopServer) GetDonuts(context.Context, *emptypb.Empty) (*Donuts, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDonuts not implemented")
}
func (UnimplementedDonutShopServer) MakeDonut(context.Context, *Donut) (*DonutRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakeDonut not implemented")
}
func (UnimplementedDonutShopServer) GetVersion(context.Context, *emptypb.Empty) (*Version, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVersion not implemented")
}
func (UnimplementedDonutShopServer) mustEmbedUnimplementedDonutShopServer() {}

// UnsafeDonutShopServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DonutShopServer will
// result in compilation errors.
type UnsafeDonutShopServer interface {
	mustEmbedUnimplementedDonutShopServer()
}

func RegisterDonutShopServer(s grpc.ServiceRegistrar, srv DonutShopServer) {
	s.RegisterService(&DonutShop_ServiceDesc, srv)
}

func _DonutShop_GetDonut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DonutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DonutShopServer).GetDonut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DonutShop/GetDonut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DonutShopServer).GetDonut(ctx, req.(*DonutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DonutShop_GetDonuts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DonutShopServer).GetDonuts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DonutShop/GetDonuts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DonutShopServer).GetDonuts(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DonutShop_MakeDonut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Donut)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DonutShopServer).MakeDonut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DonutShop/MakeDonut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DonutShopServer).MakeDonut(ctx, req.(*Donut))
	}
	return interceptor(ctx, in, info, handler)
}

func _DonutShop_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DonutShopServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DonutShop/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DonutShopServer).GetVersion(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// DonutShop_ServiceDesc is the grpc.ServiceDesc for DonutShop service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DonutShop_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "DonutShop",
	HandlerType: (*DonutShopServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDonut",
			Handler:    _DonutShop_GetDonut_Handler,
		},
		{
			MethodName: "GetDonuts",
			Handler:    _DonutShop_GetDonuts_Handler,
		},
		{
			MethodName: "MakeDonut",
			Handler:    _DonutShop_MakeDonut_Handler,
		},
		{
			MethodName: "GetVersion",
			Handler:    _DonutShop_GetVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/donuts.proto",
}

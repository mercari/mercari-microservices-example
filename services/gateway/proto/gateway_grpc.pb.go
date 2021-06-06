// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	proto "github.com/mercari/go-conference-2021-spring-office-hour/services/authority/proto"
	proto1 "github.com/mercari/go-conference-2021-spring-office-hour/services/catalog/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GatewayServiceClient is the client API for GatewayService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayServiceClient interface {
	Signup(ctx context.Context, in *proto.SignupRequest, opts ...grpc.CallOption) (*proto.SignupResponse, error)
	Signin(ctx context.Context, in *proto.SigninRequest, opts ...grpc.CallOption) (*proto.SigninResponse, error)
	CreateItem(ctx context.Context, in *proto1.CreateItemRequest, opts ...grpc.CallOption) (*proto1.CreateItemResponse, error)
	GetItem(ctx context.Context, in *proto1.GetItemRequest, opts ...grpc.CallOption) (*proto1.GetItemResponse, error)
	ListItems(ctx context.Context, in *proto1.ListItemsRequest, opts ...grpc.CallOption) (*proto1.ListItemsResponse, error)
}

type gatewayServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayServiceClient(cc grpc.ClientConnInterface) GatewayServiceClient {
	return &gatewayServiceClient{cc}
}

func (c *gatewayServiceClient) Signup(ctx context.Context, in *proto.SignupRequest, opts ...grpc.CallOption) (*proto.SignupResponse, error) {
	out := new(proto.SignupResponse)
	err := c.cc.Invoke(ctx, "/mercari.go_conference_2021_spring_office_hour.gateway.GatewayService/Signup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) Signin(ctx context.Context, in *proto.SigninRequest, opts ...grpc.CallOption) (*proto.SigninResponse, error) {
	out := new(proto.SigninResponse)
	err := c.cc.Invoke(ctx, "/mercari.go_conference_2021_spring_office_hour.gateway.GatewayService/Signin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) CreateItem(ctx context.Context, in *proto1.CreateItemRequest, opts ...grpc.CallOption) (*proto1.CreateItemResponse, error) {
	out := new(proto1.CreateItemResponse)
	err := c.cc.Invoke(ctx, "/mercari.go_conference_2021_spring_office_hour.gateway.GatewayService/CreateItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) GetItem(ctx context.Context, in *proto1.GetItemRequest, opts ...grpc.CallOption) (*proto1.GetItemResponse, error) {
	out := new(proto1.GetItemResponse)
	err := c.cc.Invoke(ctx, "/mercari.go_conference_2021_spring_office_hour.gateway.GatewayService/GetItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) ListItems(ctx context.Context, in *proto1.ListItemsRequest, opts ...grpc.CallOption) (*proto1.ListItemsResponse, error) {
	out := new(proto1.ListItemsResponse)
	err := c.cc.Invoke(ctx, "/mercari.go_conference_2021_spring_office_hour.gateway.GatewayService/ListItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServiceServer is the server API for GatewayService service.
// All implementations must embed UnimplementedGatewayServiceServer
// for forward compatibility
type GatewayServiceServer interface {
	Signup(context.Context, *proto.SignupRequest) (*proto.SignupResponse, error)
	Signin(context.Context, *proto.SigninRequest) (*proto.SigninResponse, error)
	CreateItem(context.Context, *proto1.CreateItemRequest) (*proto1.CreateItemResponse, error)
	GetItem(context.Context, *proto1.GetItemRequest) (*proto1.GetItemResponse, error)
	ListItems(context.Context, *proto1.ListItemsRequest) (*proto1.ListItemsResponse, error)
	mustEmbedUnimplementedGatewayServiceServer()
}

// UnimplementedGatewayServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayServiceServer struct {
}

func (UnimplementedGatewayServiceServer) Signup(context.Context, *proto.SignupRequest) (*proto.SignupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Signup not implemented")
}
func (UnimplementedGatewayServiceServer) Signin(context.Context, *proto.SigninRequest) (*proto.SigninResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Signin not implemented")
}
func (UnimplementedGatewayServiceServer) CreateItem(context.Context, *proto1.CreateItemRequest) (*proto1.CreateItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateItem not implemented")
}
func (UnimplementedGatewayServiceServer) GetItem(context.Context, *proto1.GetItemRequest) (*proto1.GetItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetItem not implemented")
}
func (UnimplementedGatewayServiceServer) ListItems(context.Context, *proto1.ListItemsRequest) (*proto1.ListItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListItems not implemented")
}
func (UnimplementedGatewayServiceServer) mustEmbedUnimplementedGatewayServiceServer() {}

// UnsafeGatewayServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GatewayServiceServer will
// result in compilation errors.
type UnsafeGatewayServiceServer interface {
	mustEmbedUnimplementedGatewayServiceServer()
}

func RegisterGatewayServiceServer(s grpc.ServiceRegistrar, srv GatewayServiceServer) {
	s.RegisterService(&GatewayService_ServiceDesc, srv)
}

func _GatewayService_Signup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto.SignupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).Signup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mercari.go_conference_2021_spring_office_hour.gateway.GatewayService/Signup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).Signup(ctx, req.(*proto.SignupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_Signin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto.SigninRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).Signin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mercari.go_conference_2021_spring_office_hour.gateway.GatewayService/Signin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).Signin(ctx, req.(*proto.SigninRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_CreateItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto1.CreateItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).CreateItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mercari.go_conference_2021_spring_office_hour.gateway.GatewayService/CreateItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).CreateItem(ctx, req.(*proto1.CreateItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_GetItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto1.GetItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).GetItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mercari.go_conference_2021_spring_office_hour.gateway.GatewayService/GetItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).GetItem(ctx, req.(*proto1.GetItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_ListItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto1.ListItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).ListItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mercari.go_conference_2021_spring_office_hour.gateway.GatewayService/ListItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).ListItems(ctx, req.(*proto1.ListItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GatewayService_ServiceDesc is the grpc.ServiceDesc for GatewayService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GatewayService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mercari.go_conference_2021_spring_office_hour.gateway.GatewayService",
	HandlerType: (*GatewayServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Signup",
			Handler:    _GatewayService_Signup_Handler,
		},
		{
			MethodName: "Signin",
			Handler:    _GatewayService_Signin_Handler,
		},
		{
			MethodName: "CreateItem",
			Handler:    _GatewayService_CreateItem_Handler,
		},
		{
			MethodName: "GetItem",
			Handler:    _GatewayService_GetItem_Handler,
		},
		{
			MethodName: "ListItems",
			Handler:    _GatewayService_ListItems_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/gateway/proto/gateway.proto",
}
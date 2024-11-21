// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: gluon/contract/query.proto

package contract

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

const (
	Query_Params_FullMethodName         = "/gluon.contract.Query/Params"
	Query_Order_FullMethodName          = "/gluon.contract.Query/Order"
	Query_OrderAll_FullMethodName       = "/gluon.contract.Query/OrderAll"
	Query_SortedOrder_FullMethodName    = "/gluon.contract.Query/SortedOrder"
	Query_SortedOrderAll_FullMethodName = "/gluon.contract.Query/SortedOrderAll"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a list of Order items.
	Order(ctx context.Context, in *QueryGetOrderRequest, opts ...grpc.CallOption) (*QueryGetOrderResponse, error)
	OrderAll(ctx context.Context, in *QueryAllOrderRequest, opts ...grpc.CallOption) (*QueryAllOrderResponse, error)
	// Queries a list of SortedOrder items.
	SortedOrder(ctx context.Context, in *QueryGetSortedOrderRequest, opts ...grpc.CallOption) (*QueryGetSortedOrderResponse, error)
	SortedOrderAll(ctx context.Context, in *QueryAllSortedOrderRequest, opts ...grpc.CallOption) (*QueryAllSortedOrderResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, Query_Params_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Order(ctx context.Context, in *QueryGetOrderRequest, opts ...grpc.CallOption) (*QueryGetOrderResponse, error) {
	out := new(QueryGetOrderResponse)
	err := c.cc.Invoke(ctx, Query_Order_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) OrderAll(ctx context.Context, in *QueryAllOrderRequest, opts ...grpc.CallOption) (*QueryAllOrderResponse, error) {
	out := new(QueryAllOrderResponse)
	err := c.cc.Invoke(ctx, Query_OrderAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) SortedOrder(ctx context.Context, in *QueryGetSortedOrderRequest, opts ...grpc.CallOption) (*QueryGetSortedOrderResponse, error) {
	out := new(QueryGetSortedOrderResponse)
	err := c.cc.Invoke(ctx, Query_SortedOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) SortedOrderAll(ctx context.Context, in *QueryAllSortedOrderRequest, opts ...grpc.CallOption) (*QueryAllSortedOrderResponse, error) {
	out := new(QueryAllSortedOrderResponse)
	err := c.cc.Invoke(ctx, Query_SortedOrderAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a list of Order items.
	Order(context.Context, *QueryGetOrderRequest) (*QueryGetOrderResponse, error)
	OrderAll(context.Context, *QueryAllOrderRequest) (*QueryAllOrderResponse, error)
	// Queries a list of SortedOrder items.
	SortedOrder(context.Context, *QueryGetSortedOrderRequest) (*QueryGetSortedOrderResponse, error)
	SortedOrderAll(context.Context, *QueryAllSortedOrderRequest) (*QueryAllSortedOrderResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) Order(context.Context, *QueryGetOrderRequest) (*QueryGetOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Order not implemented")
}
func (UnimplementedQueryServer) OrderAll(context.Context, *QueryAllOrderRequest) (*QueryAllOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderAll not implemented")
}
func (UnimplementedQueryServer) SortedOrder(context.Context, *QueryGetSortedOrderRequest) (*QueryGetSortedOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SortedOrder not implemented")
}
func (UnimplementedQueryServer) SortedOrderAll(context.Context, *QueryAllSortedOrderRequest) (*QueryAllSortedOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SortedOrderAll not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Params_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Order_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Order(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Order_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Order(ctx, req.(*QueryGetOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_OrderAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).OrderAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_OrderAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).OrderAll(ctx, req.(*QueryAllOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_SortedOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetSortedOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SortedOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_SortedOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).SortedOrder(ctx, req.(*QueryGetSortedOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_SortedOrderAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllSortedOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SortedOrderAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_SortedOrderAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).SortedOrderAll(ctx, req.(*QueryAllSortedOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gluon.contract.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "Order",
			Handler:    _Query_Order_Handler,
		},
		{
			MethodName: "OrderAll",
			Handler:    _Query_OrderAll_Handler,
		},
		{
			MethodName: "SortedOrder",
			Handler:    _Query_SortedOrder_Handler,
		},
		{
			MethodName: "SortedOrderAll",
			Handler:    _Query_SortedOrderAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gluon/contract/query.proto",
}

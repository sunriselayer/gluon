// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: gluon/perp/query.proto

package perp

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
	Query_Params_FullMethodName                  = "/gluon.perp.Query/Params"
	Query_Position_FullMethodName                = "/gluon.perp.Query/Position"
	Query_Positions_FullMethodName               = "/gluon.perp.Query/Positions"
	Query_PositionPriceQuantity_FullMethodName   = "/gluon.perp.Query/PositionPriceQuantity"
	Query_PositionPriceQuantities_FullMethodName = "/gluon.perp.Query/PositionPriceQuantities"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Query defines the gRPC querier service.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Position
	Position(ctx context.Context, in *QueryPositionRequest, opts ...grpc.CallOption) (*QueryPositionResponse, error)
	// Positions
	Positions(ctx context.Context, in *QueryPositionsRequest, opts ...grpc.CallOption) (*QueryPositionsResponse, error)
	// PositionPriceQuantity
	PositionPriceQuantity(ctx context.Context, in *QueryPositionPriceQuantityRequest, opts ...grpc.CallOption) (*QueryPositionPriceQuantityResponse, error)
	// PositionPriceQuantities
	PositionPriceQuantities(ctx context.Context, in *QueryPositionPriceQuantitiesRequest, opts ...grpc.CallOption) (*QueryPositionPriceQuantitiesResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, Query_Params_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Position(ctx context.Context, in *QueryPositionRequest, opts ...grpc.CallOption) (*QueryPositionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueryPositionResponse)
	err := c.cc.Invoke(ctx, Query_Position_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Positions(ctx context.Context, in *QueryPositionsRequest, opts ...grpc.CallOption) (*QueryPositionsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueryPositionsResponse)
	err := c.cc.Invoke(ctx, Query_Positions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) PositionPriceQuantity(ctx context.Context, in *QueryPositionPriceQuantityRequest, opts ...grpc.CallOption) (*QueryPositionPriceQuantityResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueryPositionPriceQuantityResponse)
	err := c.cc.Invoke(ctx, Query_PositionPriceQuantity_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) PositionPriceQuantities(ctx context.Context, in *QueryPositionPriceQuantitiesRequest, opts ...grpc.CallOption) (*QueryPositionPriceQuantitiesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueryPositionPriceQuantitiesResponse)
	err := c.cc.Invoke(ctx, Query_PositionPriceQuantities_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility.
//
// Query defines the gRPC querier service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Position
	Position(context.Context, *QueryPositionRequest) (*QueryPositionResponse, error)
	// Positions
	Positions(context.Context, *QueryPositionsRequest) (*QueryPositionsResponse, error)
	// PositionPriceQuantity
	PositionPriceQuantity(context.Context, *QueryPositionPriceQuantityRequest) (*QueryPositionPriceQuantityResponse, error)
	// PositionPriceQuantities
	PositionPriceQuantities(context.Context, *QueryPositionPriceQuantitiesRequest) (*QueryPositionPriceQuantitiesResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedQueryServer struct{}

func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) Position(context.Context, *QueryPositionRequest) (*QueryPositionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Position not implemented")
}
func (UnimplementedQueryServer) Positions(context.Context, *QueryPositionsRequest) (*QueryPositionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Positions not implemented")
}
func (UnimplementedQueryServer) PositionPriceQuantity(context.Context, *QueryPositionPriceQuantityRequest) (*QueryPositionPriceQuantityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PositionPriceQuantity not implemented")
}
func (UnimplementedQueryServer) PositionPriceQuantities(context.Context, *QueryPositionPriceQuantitiesRequest) (*QueryPositionPriceQuantitiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PositionPriceQuantities not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}
func (UnimplementedQueryServer) testEmbeddedByValue()               {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	// If the following call pancis, it indicates UnimplementedQueryServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
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

func _Query_Position_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPositionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Position(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Position_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Position(ctx, req.(*QueryPositionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Positions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPositionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Positions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Positions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Positions(ctx, req.(*QueryPositionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_PositionPriceQuantity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPositionPriceQuantityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).PositionPriceQuantity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_PositionPriceQuantity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).PositionPriceQuantity(ctx, req.(*QueryPositionPriceQuantityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_PositionPriceQuantities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPositionPriceQuantitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).PositionPriceQuantities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_PositionPriceQuantities_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).PositionPriceQuantities(ctx, req.(*QueryPositionPriceQuantitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gluon.perp.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "Position",
			Handler:    _Query_Position_Handler,
		},
		{
			MethodName: "Positions",
			Handler:    _Query_Positions_Handler,
		},
		{
			MethodName: "PositionPriceQuantity",
			Handler:    _Query_PositionPriceQuantity_Handler,
		},
		{
			MethodName: "PositionPriceQuantities",
			Handler:    _Query_PositionPriceQuantities_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gluon/perp/query.proto",
}

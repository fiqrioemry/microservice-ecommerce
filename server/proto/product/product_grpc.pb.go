// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.30.2
// source: proto/product/product.proto

package product

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
	ProductService_GetProductSnapshot_FullMethodName          = "/product.ProductService/GetProductSnapshot"
	ProductService_GetMultipleProductSnapshots_FullMethodName = "/product.ProductService/GetMultipleProductSnapshots"
	ProductService_CheckProductAvailability_FullMethodName    = "/product.ProductService/CheckProductAvailability"
	ProductService_UpdateProductStock_FullMethodName          = "/product.ProductService/UpdateProductStock"
)

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// === SERVICE ===
type ProductServiceClient interface {
	GetProductSnapshot(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*ProductSnapshotResponse, error)
	GetMultipleProductSnapshots(ctx context.Context, in *GetMultipleProductRequest, opts ...grpc.CallOption) (*MultipleProductSnapshotResponse, error)
	CheckProductAvailability(ctx context.Context, in *CheckAvailabilityRequest, opts ...grpc.CallOption) (*CheckAvailabilityResponse, error)
	UpdateProductStock(ctx context.Context, in *UpdateStockRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
}

type productServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductServiceClient(cc grpc.ClientConnInterface) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) GetProductSnapshot(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*ProductSnapshotResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProductSnapshotResponse)
	err := c.cc.Invoke(ctx, ProductService_GetProductSnapshot_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetMultipleProductSnapshots(ctx context.Context, in *GetMultipleProductRequest, opts ...grpc.CallOption) (*MultipleProductSnapshotResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MultipleProductSnapshotResponse)
	err := c.cc.Invoke(ctx, ProductService_GetMultipleProductSnapshots_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) CheckProductAvailability(ctx context.Context, in *CheckAvailabilityRequest, opts ...grpc.CallOption) (*CheckAvailabilityResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckAvailabilityResponse)
	err := c.cc.Invoke(ctx, ProductService_CheckProductAvailability_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) UpdateProductStock(ctx context.Context, in *UpdateStockRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, ProductService_UpdateProductStock_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServiceServer is the server API for ProductService service.
// All implementations must embed UnimplementedProductServiceServer
// for forward compatibility.
//
// === SERVICE ===
type ProductServiceServer interface {
	GetProductSnapshot(context.Context, *GetProductRequest) (*ProductSnapshotResponse, error)
	GetMultipleProductSnapshots(context.Context, *GetMultipleProductRequest) (*MultipleProductSnapshotResponse, error)
	CheckProductAvailability(context.Context, *CheckAvailabilityRequest) (*CheckAvailabilityResponse, error)
	UpdateProductStock(context.Context, *UpdateStockRequest) (*EmptyResponse, error)
	mustEmbedUnimplementedProductServiceServer()
}

// UnimplementedProductServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedProductServiceServer struct{}

func (UnimplementedProductServiceServer) GetProductSnapshot(context.Context, *GetProductRequest) (*ProductSnapshotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductSnapshot not implemented")
}
func (UnimplementedProductServiceServer) GetMultipleProductSnapshots(context.Context, *GetMultipleProductRequest) (*MultipleProductSnapshotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMultipleProductSnapshots not implemented")
}
func (UnimplementedProductServiceServer) CheckProductAvailability(context.Context, *CheckAvailabilityRequest) (*CheckAvailabilityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckProductAvailability not implemented")
}
func (UnimplementedProductServiceServer) UpdateProductStock(context.Context, *UpdateStockRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProductStock not implemented")
}
func (UnimplementedProductServiceServer) mustEmbedUnimplementedProductServiceServer() {}
func (UnimplementedProductServiceServer) testEmbeddedByValue()                        {}

// UnsafeProductServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductServiceServer will
// result in compilation errors.
type UnsafeProductServiceServer interface {
	mustEmbedUnimplementedProductServiceServer()
}

func RegisterProductServiceServer(s grpc.ServiceRegistrar, srv ProductServiceServer) {
	// If the following call pancis, it indicates UnimplementedProductServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ProductService_ServiceDesc, srv)
}

func _ProductService_GetProductSnapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetProductSnapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_GetProductSnapshot_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetProductSnapshot(ctx, req.(*GetProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetMultipleProductSnapshots_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMultipleProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetMultipleProductSnapshots(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_GetMultipleProductSnapshots_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetMultipleProductSnapshots(ctx, req.(*GetMultipleProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_CheckProductAvailability_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckAvailabilityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).CheckProductAvailability(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_CheckProductAvailability_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).CheckProductAvailability(ctx, req.(*CheckAvailabilityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_UpdateProductStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).UpdateProductStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_UpdateProductStock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).UpdateProductStock(ctx, req.(*UpdateStockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductService_ServiceDesc is the grpc.ServiceDesc for ProductService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "product.ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProductSnapshot",
			Handler:    _ProductService_GetProductSnapshot_Handler,
		},
		{
			MethodName: "GetMultipleProductSnapshots",
			Handler:    _ProductService_GetMultipleProductSnapshots_Handler,
		},
		{
			MethodName: "CheckProductAvailability",
			Handler:    _ProductService_CheckProductAvailability_Handler,
		},
		{
			MethodName: "UpdateProductStock",
			Handler:    _ProductService_UpdateProductStock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/product/product.proto",
}

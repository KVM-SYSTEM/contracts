// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.0
// source: core.proto

package core_server

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
	MaterialService_Add_FullMethodName     = "/measure.MaterialService/Add"
	MaterialService_Get_FullMethodName     = "/measure.MaterialService/Get"
	MaterialService_GetMany_FullMethodName = "/measure.MaterialService/GetMany"
	MaterialService_Delete_FullMethodName  = "/measure.MaterialService/Delete"
)

// MaterialServiceClient is the client API for MaterialService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MaterialServiceClient interface {
	Add(ctx context.Context, in *AddMaterialReq, opts ...grpc.CallOption) (*AddMaterialRes, error)
	Get(ctx context.Context, in *GetMaterialReq, opts ...grpc.CallOption) (*GetMaterialRes, error)
	GetMany(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetManyMaterialRes, error)
	Delete(ctx context.Context, in *DeleteMaterialReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type materialServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMaterialServiceClient(cc grpc.ClientConnInterface) MaterialServiceClient {
	return &materialServiceClient{cc}
}

func (c *materialServiceClient) Add(ctx context.Context, in *AddMaterialReq, opts ...grpc.CallOption) (*AddMaterialRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddMaterialRes)
	err := c.cc.Invoke(ctx, MaterialService_Add_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *materialServiceClient) Get(ctx context.Context, in *GetMaterialReq, opts ...grpc.CallOption) (*GetMaterialRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMaterialRes)
	err := c.cc.Invoke(ctx, MaterialService_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *materialServiceClient) GetMany(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetManyMaterialRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetManyMaterialRes)
	err := c.cc.Invoke(ctx, MaterialService_GetMany_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *materialServiceClient) Delete(ctx context.Context, in *DeleteMaterialReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, MaterialService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MaterialServiceServer is the server API for MaterialService service.
// All implementations must embed UnimplementedMaterialServiceServer
// for forward compatibility.
type MaterialServiceServer interface {
	Add(context.Context, *AddMaterialReq) (*AddMaterialRes, error)
	Get(context.Context, *GetMaterialReq) (*GetMaterialRes, error)
	GetMany(context.Context, *emptypb.Empty) (*GetManyMaterialRes, error)
	Delete(context.Context, *DeleteMaterialReq) (*emptypb.Empty, error)
	mustEmbedUnimplementedMaterialServiceServer()
}

// UnimplementedMaterialServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMaterialServiceServer struct{}

func (UnimplementedMaterialServiceServer) Add(context.Context, *AddMaterialReq) (*AddMaterialRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedMaterialServiceServer) Get(context.Context, *GetMaterialReq) (*GetMaterialRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedMaterialServiceServer) GetMany(context.Context, *emptypb.Empty) (*GetManyMaterialRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMany not implemented")
}
func (UnimplementedMaterialServiceServer) Delete(context.Context, *DeleteMaterialReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedMaterialServiceServer) mustEmbedUnimplementedMaterialServiceServer() {}
func (UnimplementedMaterialServiceServer) testEmbeddedByValue()                         {}

// UnsafeMaterialServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MaterialServiceServer will
// result in compilation errors.
type UnsafeMaterialServiceServer interface {
	mustEmbedUnimplementedMaterialServiceServer()
}

func RegisterMaterialServiceServer(s grpc.ServiceRegistrar, srv MaterialServiceServer) {
	// If the following call pancis, it indicates UnimplementedMaterialServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MaterialService_ServiceDesc, srv)
}

func _MaterialService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMaterialReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MaterialServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MaterialService_Add_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MaterialServiceServer).Add(ctx, req.(*AddMaterialReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MaterialService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMaterialReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MaterialServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MaterialService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MaterialServiceServer).Get(ctx, req.(*GetMaterialReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MaterialService_GetMany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MaterialServiceServer).GetMany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MaterialService_GetMany_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MaterialServiceServer).GetMany(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MaterialService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMaterialReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MaterialServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MaterialService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MaterialServiceServer).Delete(ctx, req.(*DeleteMaterialReq))
	}
	return interceptor(ctx, in, info, handler)
}

// MaterialService_ServiceDesc is the grpc.ServiceDesc for MaterialService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MaterialService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "measure.MaterialService",
	HandlerType: (*MaterialServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _MaterialService_Add_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _MaterialService_Get_Handler,
		},
		{
			MethodName: "GetMany",
			Handler:    _MaterialService_GetMany_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _MaterialService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "core.proto",
}

const (
	PriceService_Get_FullMethodName     = "/measure.PriceService/Get"
	PriceService_GetMany_FullMethodName = "/measure.PriceService/GetMany"
)

// PriceServiceClient is the client API for PriceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PriceServiceClient interface {
	Get(ctx context.Context, in *GetPriceReq, opts ...grpc.CallOption) (*GetPriceRes, error)
	GetMany(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetManyPricesRes, error)
}

type priceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPriceServiceClient(cc grpc.ClientConnInterface) PriceServiceClient {
	return &priceServiceClient{cc}
}

func (c *priceServiceClient) Get(ctx context.Context, in *GetPriceReq, opts ...grpc.CallOption) (*GetPriceRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPriceRes)
	err := c.cc.Invoke(ctx, PriceService_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *priceServiceClient) GetMany(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetManyPricesRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetManyPricesRes)
	err := c.cc.Invoke(ctx, PriceService_GetMany_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PriceServiceServer is the server API for PriceService service.
// All implementations must embed UnimplementedPriceServiceServer
// for forward compatibility.
type PriceServiceServer interface {
	Get(context.Context, *GetPriceReq) (*GetPriceRes, error)
	GetMany(context.Context, *emptypb.Empty) (*GetManyPricesRes, error)
	mustEmbedUnimplementedPriceServiceServer()
}

// UnimplementedPriceServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPriceServiceServer struct{}

func (UnimplementedPriceServiceServer) Get(context.Context, *GetPriceReq) (*GetPriceRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedPriceServiceServer) GetMany(context.Context, *emptypb.Empty) (*GetManyPricesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMany not implemented")
}
func (UnimplementedPriceServiceServer) mustEmbedUnimplementedPriceServiceServer() {}
func (UnimplementedPriceServiceServer) testEmbeddedByValue()                      {}

// UnsafePriceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PriceServiceServer will
// result in compilation errors.
type UnsafePriceServiceServer interface {
	mustEmbedUnimplementedPriceServiceServer()
}

func RegisterPriceServiceServer(s grpc.ServiceRegistrar, srv PriceServiceServer) {
	// If the following call pancis, it indicates UnimplementedPriceServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PriceService_ServiceDesc, srv)
}

func _PriceService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPriceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PriceServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PriceService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PriceServiceServer).Get(ctx, req.(*GetPriceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PriceService_GetMany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PriceServiceServer).GetMany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PriceService_GetMany_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PriceServiceServer).GetMany(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// PriceService_ServiceDesc is the grpc.ServiceDesc for PriceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PriceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "measure.PriceService",
	HandlerType: (*PriceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _PriceService_Get_Handler,
		},
		{
			MethodName: "GetMany",
			Handler:    _PriceService_GetMany_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "core.proto",
}

const (
	DivideService_Add_FullMethodName     = "/measure.DivideService/Add"
	DivideService_GetMany_FullMethodName = "/measure.DivideService/GetMany"
	DivideService_Get_FullMethodName     = "/measure.DivideService/Get"
	DivideService_Delete_FullMethodName  = "/measure.DivideService/Delete"
)

// DivideServiceClient is the client API for DivideService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DivideServiceClient interface {
	Add(ctx context.Context, in *AddDivideReq, opts ...grpc.CallOption) (*AddDivideRes, error)
	GetMany(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetManyDividesRes, error)
	Get(ctx context.Context, in *GetDivideReq, opts ...grpc.CallOption) (*GetDivideRes, error)
	Delete(ctx context.Context, in *DeleteDivideReq, opts ...grpc.CallOption) (*DeleteDivideRes, error)
}

type divideServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDivideServiceClient(cc grpc.ClientConnInterface) DivideServiceClient {
	return &divideServiceClient{cc}
}

func (c *divideServiceClient) Add(ctx context.Context, in *AddDivideReq, opts ...grpc.CallOption) (*AddDivideRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddDivideRes)
	err := c.cc.Invoke(ctx, DivideService_Add_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *divideServiceClient) GetMany(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetManyDividesRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetManyDividesRes)
	err := c.cc.Invoke(ctx, DivideService_GetMany_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *divideServiceClient) Get(ctx context.Context, in *GetDivideReq, opts ...grpc.CallOption) (*GetDivideRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetDivideRes)
	err := c.cc.Invoke(ctx, DivideService_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *divideServiceClient) Delete(ctx context.Context, in *DeleteDivideReq, opts ...grpc.CallOption) (*DeleteDivideRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteDivideRes)
	err := c.cc.Invoke(ctx, DivideService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DivideServiceServer is the server API for DivideService service.
// All implementations must embed UnimplementedDivideServiceServer
// for forward compatibility.
type DivideServiceServer interface {
	Add(context.Context, *AddDivideReq) (*AddDivideRes, error)
	GetMany(context.Context, *emptypb.Empty) (*GetManyDividesRes, error)
	Get(context.Context, *GetDivideReq) (*GetDivideRes, error)
	Delete(context.Context, *DeleteDivideReq) (*DeleteDivideRes, error)
	mustEmbedUnimplementedDivideServiceServer()
}

// UnimplementedDivideServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDivideServiceServer struct{}

func (UnimplementedDivideServiceServer) Add(context.Context, *AddDivideReq) (*AddDivideRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedDivideServiceServer) GetMany(context.Context, *emptypb.Empty) (*GetManyDividesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMany not implemented")
}
func (UnimplementedDivideServiceServer) Get(context.Context, *GetDivideReq) (*GetDivideRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedDivideServiceServer) Delete(context.Context, *DeleteDivideReq) (*DeleteDivideRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedDivideServiceServer) mustEmbedUnimplementedDivideServiceServer() {}
func (UnimplementedDivideServiceServer) testEmbeddedByValue()                       {}

// UnsafeDivideServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DivideServiceServer will
// result in compilation errors.
type UnsafeDivideServiceServer interface {
	mustEmbedUnimplementedDivideServiceServer()
}

func RegisterDivideServiceServer(s grpc.ServiceRegistrar, srv DivideServiceServer) {
	// If the following call pancis, it indicates UnimplementedDivideServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DivideService_ServiceDesc, srv)
}

func _DivideService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDivideReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DivideServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DivideService_Add_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DivideServiceServer).Add(ctx, req.(*AddDivideReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DivideService_GetMany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DivideServiceServer).GetMany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DivideService_GetMany_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DivideServiceServer).GetMany(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DivideService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDivideReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DivideServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DivideService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DivideServiceServer).Get(ctx, req.(*GetDivideReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DivideService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDivideReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DivideServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DivideService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DivideServiceServer).Delete(ctx, req.(*DeleteDivideReq))
	}
	return interceptor(ctx, in, info, handler)
}

// DivideService_ServiceDesc is the grpc.ServiceDesc for DivideService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DivideService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "measure.DivideService",
	HandlerType: (*DivideServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _DivideService_Add_Handler,
		},
		{
			MethodName: "GetMany",
			Handler:    _DivideService_GetMany_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _DivideService_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _DivideService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "core.proto",
}

const (
	SectionService_GetMany_FullMethodName = "/measure.SectionService/GetMany"
	SectionService_Get_FullMethodName     = "/measure.SectionService/Get"
)

// SectionServiceClient is the client API for SectionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SectionServiceClient interface {
	GetMany(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetManySectionsRes, error)
	Get(ctx context.Context, in *GetSectionReq, opts ...grpc.CallOption) (*GetSectionRes, error)
}

type sectionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSectionServiceClient(cc grpc.ClientConnInterface) SectionServiceClient {
	return &sectionServiceClient{cc}
}

func (c *sectionServiceClient) GetMany(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetManySectionsRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetManySectionsRes)
	err := c.cc.Invoke(ctx, SectionService_GetMany_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sectionServiceClient) Get(ctx context.Context, in *GetSectionReq, opts ...grpc.CallOption) (*GetSectionRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSectionRes)
	err := c.cc.Invoke(ctx, SectionService_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SectionServiceServer is the server API for SectionService service.
// All implementations must embed UnimplementedSectionServiceServer
// for forward compatibility.
type SectionServiceServer interface {
	GetMany(context.Context, *emptypb.Empty) (*GetManySectionsRes, error)
	Get(context.Context, *GetSectionReq) (*GetSectionRes, error)
	mustEmbedUnimplementedSectionServiceServer()
}

// UnimplementedSectionServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSectionServiceServer struct{}

func (UnimplementedSectionServiceServer) GetMany(context.Context, *emptypb.Empty) (*GetManySectionsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMany not implemented")
}
func (UnimplementedSectionServiceServer) Get(context.Context, *GetSectionReq) (*GetSectionRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedSectionServiceServer) mustEmbedUnimplementedSectionServiceServer() {}
func (UnimplementedSectionServiceServer) testEmbeddedByValue()                        {}

// UnsafeSectionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SectionServiceServer will
// result in compilation errors.
type UnsafeSectionServiceServer interface {
	mustEmbedUnimplementedSectionServiceServer()
}

func RegisterSectionServiceServer(s grpc.ServiceRegistrar, srv SectionServiceServer) {
	// If the following call pancis, it indicates UnimplementedSectionServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SectionService_ServiceDesc, srv)
}

func _SectionService_GetMany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SectionServiceServer).GetMany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SectionService_GetMany_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SectionServiceServer).GetMany(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SectionService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSectionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SectionServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SectionService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SectionServiceServer).Get(ctx, req.(*GetSectionReq))
	}
	return interceptor(ctx, in, info, handler)
}

// SectionService_ServiceDesc is the grpc.ServiceDesc for SectionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SectionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "measure.SectionService",
	HandlerType: (*SectionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMany",
			Handler:    _SectionService_GetMany_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _SectionService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "core.proto",
}

const (
	UserService_CheckAuth_FullMethodName = "/measure.UserService/CheckAuth"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	CheckAuth(ctx context.Context, in *CheckAuthReq, opts ...grpc.CallOption) (*CheckAuthRes, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) CheckAuth(ctx context.Context, in *CheckAuthReq, opts ...grpc.CallOption) (*CheckAuthRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckAuthRes)
	err := c.cc.Invoke(ctx, UserService_CheckAuth_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility.
type UserServiceServer interface {
	CheckAuth(context.Context, *CheckAuthReq) (*CheckAuthRes, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserServiceServer struct{}

func (UnimplementedUserServiceServer) CheckAuth(context.Context, *CheckAuthReq) (*CheckAuthRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckAuth not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}
func (UnimplementedUserServiceServer) testEmbeddedByValue()                     {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	// If the following call pancis, it indicates UnimplementedUserServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_CheckAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckAuthReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CheckAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_CheckAuth_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CheckAuth(ctx, req.(*CheckAuthReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "measure.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckAuth",
			Handler:    _UserService_CheckAuth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "core.proto",
}
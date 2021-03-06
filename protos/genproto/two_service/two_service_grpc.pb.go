// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.14.0
// source: two_service_protos/two_service.proto

package two_service

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

// TwoMicroserviceClient is the client API for TwoMicroservice service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TwoMicroserviceClient interface {
	Search(ctx context.Context, in *ReadAllPostReq, opts ...grpc.CallOption) (*ReadAllPostRes, error)
	Delete(ctx context.Context, in *DeletePostReq, opts ...grpc.CallOption) (*DeletePostRes, error)
	GetByID(ctx context.Context, in *GetByIDPostReq, opts ...grpc.CallOption) (*GetByIDPostRes, error)
	Update(ctx context.Context, in *UpdatePostReq, opts ...grpc.CallOption) (*UpdatePostRes, error)
}

type twoMicroserviceClient struct {
	cc grpc.ClientConnInterface
}

func NewTwoMicroserviceClient(cc grpc.ClientConnInterface) TwoMicroserviceClient {
	return &twoMicroserviceClient{cc}
}

func (c *twoMicroserviceClient) Search(ctx context.Context, in *ReadAllPostReq, opts ...grpc.CallOption) (*ReadAllPostRes, error) {
	out := new(ReadAllPostRes)
	err := c.cc.Invoke(ctx, "/two_service.TwoMicroservice/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *twoMicroserviceClient) Delete(ctx context.Context, in *DeletePostReq, opts ...grpc.CallOption) (*DeletePostRes, error) {
	out := new(DeletePostRes)
	err := c.cc.Invoke(ctx, "/two_service.TwoMicroservice/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *twoMicroserviceClient) GetByID(ctx context.Context, in *GetByIDPostReq, opts ...grpc.CallOption) (*GetByIDPostRes, error) {
	out := new(GetByIDPostRes)
	err := c.cc.Invoke(ctx, "/two_service.TwoMicroservice/GetByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *twoMicroserviceClient) Update(ctx context.Context, in *UpdatePostReq, opts ...grpc.CallOption) (*UpdatePostRes, error) {
	out := new(UpdatePostRes)
	err := c.cc.Invoke(ctx, "/two_service.TwoMicroservice/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TwoMicroserviceServer is the server API for TwoMicroservice service.
// All implementations should embed UnimplementedTwoMicroserviceServer
// for forward compatibility
type TwoMicroserviceServer interface {
	Search(context.Context, *ReadAllPostReq) (*ReadAllPostRes, error)
	Delete(context.Context, *DeletePostReq) (*DeletePostRes, error)
	GetByID(context.Context, *GetByIDPostReq) (*GetByIDPostRes, error)
	Update(context.Context, *UpdatePostReq) (*UpdatePostRes, error)
}

// UnimplementedTwoMicroserviceServer should be embedded to have forward compatible implementations.
type UnimplementedTwoMicroserviceServer struct {
}

func (UnimplementedTwoMicroserviceServer) Search(context.Context, *ReadAllPostReq) (*ReadAllPostRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedTwoMicroserviceServer) Delete(context.Context, *DeletePostReq) (*DeletePostRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedTwoMicroserviceServer) GetByID(context.Context, *GetByIDPostReq) (*GetByIDPostRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (UnimplementedTwoMicroserviceServer) Update(context.Context, *UpdatePostReq) (*UpdatePostRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}

// UnsafeTwoMicroserviceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TwoMicroserviceServer will
// result in compilation errors.
type UnsafeTwoMicroserviceServer interface {
	mustEmbedUnimplementedTwoMicroserviceServer()
}

func RegisterTwoMicroserviceServer(s grpc.ServiceRegistrar, srv TwoMicroserviceServer) {
	s.RegisterService(&TwoMicroservice_ServiceDesc, srv)
}

func _TwoMicroservice_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadAllPostReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TwoMicroserviceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/two_service.TwoMicroservice/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TwoMicroserviceServer).Search(ctx, req.(*ReadAllPostReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TwoMicroservice_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePostReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TwoMicroserviceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/two_service.TwoMicroservice/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TwoMicroserviceServer).Delete(ctx, req.(*DeletePostReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TwoMicroservice_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIDPostReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TwoMicroserviceServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/two_service.TwoMicroservice/GetByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TwoMicroserviceServer).GetByID(ctx, req.(*GetByIDPostReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TwoMicroservice_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePostReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TwoMicroserviceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/two_service.TwoMicroservice/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TwoMicroserviceServer).Update(ctx, req.(*UpdatePostReq))
	}
	return interceptor(ctx, in, info, handler)
}

// TwoMicroservice_ServiceDesc is the grpc.ServiceDesc for TwoMicroservice service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TwoMicroservice_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "two_service.TwoMicroservice",
	HandlerType: (*TwoMicroserviceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _TwoMicroservice_Search_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _TwoMicroservice_Delete_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _TwoMicroservice_GetByID_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _TwoMicroservice_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "two_service_protos/two_service.proto",
}

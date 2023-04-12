// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: model/sinegmapay.proto

package service

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

// SinegmaPaymentClient is the client API for SinegmaPayment service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SinegmaPaymentClient interface {
	CheckBalance(ctx context.Context, in *CheckBalanceMessage, opts ...grpc.CallOption) (*ResultMessage, error)
	DoPayment(ctx context.Context, in *PaymentMessage, opts ...grpc.CallOption) (*ResultMessage, error)
}

type sinegmaPaymentClient struct {
	cc grpc.ClientConnInterface
}

func NewSinegmaPaymentClient(cc grpc.ClientConnInterface) SinegmaPaymentClient {
	return &sinegmaPaymentClient{cc}
}

func (c *sinegmaPaymentClient) CheckBalance(ctx context.Context, in *CheckBalanceMessage, opts ...grpc.CallOption) (*ResultMessage, error) {
	out := new(ResultMessage)
	err := c.cc.Invoke(ctx, "/service.SinegmaPayment/CheckBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sinegmaPaymentClient) DoPayment(ctx context.Context, in *PaymentMessage, opts ...grpc.CallOption) (*ResultMessage, error) {
	out := new(ResultMessage)
	err := c.cc.Invoke(ctx, "/service.SinegmaPayment/DoPayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SinegmaPaymentServer is the server API for SinegmaPayment service.
// All implementations must embed UnimplementedSinegmaPaymentServer
// for forward compatibility
type SinegmaPaymentServer interface {
	CheckBalance(context.Context, *CheckBalanceMessage) (*ResultMessage, error)
	DoPayment(context.Context, *PaymentMessage) (*ResultMessage, error)
	mustEmbedUnimplementedSinegmaPaymentServer()
}

// UnimplementedSinegmaPaymentServer must be embedded to have forward compatible implementations.
type UnimplementedSinegmaPaymentServer struct {
}

func (UnimplementedSinegmaPaymentServer) CheckBalance(context.Context, *CheckBalanceMessage) (*ResultMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckBalance not implemented")
}
func (UnimplementedSinegmaPaymentServer) DoPayment(context.Context, *PaymentMessage) (*ResultMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoPayment not implemented")
}
func (UnimplementedSinegmaPaymentServer) mustEmbedUnimplementedSinegmaPaymentServer() {}

// UnsafeSinegmaPaymentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SinegmaPaymentServer will
// result in compilation errors.
type UnsafeSinegmaPaymentServer interface {
	mustEmbedUnimplementedSinegmaPaymentServer()
}

func RegisterSinegmaPaymentServer(s grpc.ServiceRegistrar, srv SinegmaPaymentServer) {
	s.RegisterService(&SinegmaPayment_ServiceDesc, srv)
}

func _SinegmaPayment_CheckBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckBalanceMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SinegmaPaymentServer).CheckBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.SinegmaPayment/CheckBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SinegmaPaymentServer).CheckBalance(ctx, req.(*CheckBalanceMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _SinegmaPayment_DoPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaymentMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SinegmaPaymentServer).DoPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.SinegmaPayment/DoPayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SinegmaPaymentServer).DoPayment(ctx, req.(*PaymentMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// SinegmaPayment_ServiceDesc is the grpc.ServiceDesc for SinegmaPayment service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SinegmaPayment_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.SinegmaPayment",
	HandlerType: (*SinegmaPaymentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckBalance",
			Handler:    _SinegmaPayment_CheckBalance_Handler,
		},
		{
			MethodName: "DoPayment",
			Handler:    _SinegmaPayment_DoPayment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "model/sinegmapay.proto",
}
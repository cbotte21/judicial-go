// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: judicial.proto

package pb

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
	JudicialService_Ban_FullMethodName       = "/JudicialService/Ban"
	JudicialService_Unban_FullMethodName     = "/JudicialService/Unban"
	JudicialService_Integrity_FullMethodName = "/JudicialService/Integrity"
)

// JudicialServiceClient is the client API for JudicialService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JudicialServiceClient interface {
	Ban(ctx context.Context, in *BanRequest, opts ...grpc.CallOption) (*BanResponse, error)
	Unban(ctx context.Context, in *UnbanRequest, opts ...grpc.CallOption) (*UnbanResponse, error)
	Integrity(ctx context.Context, in *IntegrityRequest, opts ...grpc.CallOption) (*IntegrityResponse, error)
}

type judicialServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewJudicialServiceClient(cc grpc.ClientConnInterface) JudicialServiceClient {
	return &judicialServiceClient{cc}
}

func (c *judicialServiceClient) Ban(ctx context.Context, in *BanRequest, opts ...grpc.CallOption) (*BanResponse, error) {
	out := new(BanResponse)
	err := c.cc.Invoke(ctx, JudicialService_Ban_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *judicialServiceClient) Unban(ctx context.Context, in *UnbanRequest, opts ...grpc.CallOption) (*UnbanResponse, error) {
	out := new(UnbanResponse)
	err := c.cc.Invoke(ctx, JudicialService_Unban_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *judicialServiceClient) Integrity(ctx context.Context, in *IntegrityRequest, opts ...grpc.CallOption) (*IntegrityResponse, error) {
	out := new(IntegrityResponse)
	err := c.cc.Invoke(ctx, JudicialService_Integrity_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JudicialServiceServer is the server API for JudicialService service.
// All implementations must embed UnimplementedJudicialServiceServer
// for forward compatibility
type JudicialServiceServer interface {
	Ban(context.Context, *BanRequest) (*BanResponse, error)
	Unban(context.Context, *UnbanRequest) (*UnbanResponse, error)
	Integrity(context.Context, *IntegrityRequest) (*IntegrityResponse, error)
	mustEmbedUnimplementedJudicialServiceServer()
}

// UnimplementedJudicialServiceServer must be embedded to have forward compatible implementations.
type UnimplementedJudicialServiceServer struct {
}

func (UnimplementedJudicialServiceServer) Ban(context.Context, *BanRequest) (*BanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ban not implemented")
}
func (UnimplementedJudicialServiceServer) Unban(context.Context, *UnbanRequest) (*UnbanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unban not implemented")
}
func (UnimplementedJudicialServiceServer) Integrity(context.Context, *IntegrityRequest) (*IntegrityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Integrity not implemented")
}
func (UnimplementedJudicialServiceServer) mustEmbedUnimplementedJudicialServiceServer() {}

// UnsafeJudicialServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JudicialServiceServer will
// result in compilation errors.
type UnsafeJudicialServiceServer interface {
	mustEmbedUnimplementedJudicialServiceServer()
}

func RegisterJudicialServiceServer(s grpc.ServiceRegistrar, srv JudicialServiceServer) {
	s.RegisterService(&JudicialService_ServiceDesc, srv)
}

func _JudicialService_Ban_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JudicialServiceServer).Ban(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JudicialService_Ban_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JudicialServiceServer).Ban(ctx, req.(*BanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JudicialService_Unban_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnbanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JudicialServiceServer).Unban(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JudicialService_Unban_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JudicialServiceServer).Unban(ctx, req.(*UnbanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JudicialService_Integrity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IntegrityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JudicialServiceServer).Integrity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JudicialService_Integrity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JudicialServiceServer).Integrity(ctx, req.(*IntegrityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// JudicialService_ServiceDesc is the grpc.ServiceDesc for JudicialService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var JudicialService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "JudicialService",
	HandlerType: (*JudicialServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ban",
			Handler:    _JudicialService_Ban_Handler,
		},
		{
			MethodName: "Unban",
			Handler:    _JudicialService_Unban_Handler,
		},
		{
			MethodName: "Integrity",
			Handler:    _JudicialService_Integrity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "judicial.proto",
}

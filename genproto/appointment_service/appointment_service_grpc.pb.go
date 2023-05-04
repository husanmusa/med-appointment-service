// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package appointment_service

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

// AppointmentServiceClient is the client API for AppointmentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AppointmentServiceClient interface {
	Create(ctx context.Context, in *Appointment, opts ...grpc.CallOption) (*Appointment, error)
	Get(ctx context.Context, in *AppointmentId, opts ...grpc.CallOption) (*Appointment, error)
	Cancel(ctx context.Context, in *AppointmentId, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type appointmentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAppointmentServiceClient(cc grpc.ClientConnInterface) AppointmentServiceClient {
	return &appointmentServiceClient{cc}
}

func (c *appointmentServiceClient) Create(ctx context.Context, in *Appointment, opts ...grpc.CallOption) (*Appointment, error) {
	out := new(Appointment)
	err := c.cc.Invoke(ctx, "/appointment_service.AppointmentService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appointmentServiceClient) Get(ctx context.Context, in *AppointmentId, opts ...grpc.CallOption) (*Appointment, error) {
	out := new(Appointment)
	err := c.cc.Invoke(ctx, "/appointment_service.AppointmentService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appointmentServiceClient) Cancel(ctx context.Context, in *AppointmentId, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/appointment_service.AppointmentService/Cancel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppointmentServiceServer is the server API for AppointmentService service.
// All implementations must embed UnimplementedAppointmentServiceServer
// for forward compatibility
type AppointmentServiceServer interface {
	Create(context.Context, *Appointment) (*Appointment, error)
	Get(context.Context, *AppointmentId) (*Appointment, error)
	Cancel(context.Context, *AppointmentId) (*emptypb.Empty, error)
	mustEmbedUnimplementedAppointmentServiceServer()
}

// UnimplementedAppointmentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAppointmentServiceServer struct {
}

func (UnimplementedAppointmentServiceServer) Create(context.Context, *Appointment) (*Appointment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedAppointmentServiceServer) Get(context.Context, *AppointmentId) (*Appointment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedAppointmentServiceServer) Cancel(context.Context, *AppointmentId) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cancel not implemented")
}
func (UnimplementedAppointmentServiceServer) mustEmbedUnimplementedAppointmentServiceServer() {}

// UnsafeAppointmentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AppointmentServiceServer will
// result in compilation errors.
type UnsafeAppointmentServiceServer interface {
	mustEmbedUnimplementedAppointmentServiceServer()
}

func RegisterAppointmentServiceServer(s grpc.ServiceRegistrar, srv AppointmentServiceServer) {
	s.RegisterService(&AppointmentService_ServiceDesc, srv)
}

func _AppointmentService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Appointment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppointmentServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appointment_service.AppointmentService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppointmentServiceServer).Create(ctx, req.(*Appointment))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppointmentService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppointmentId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppointmentServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appointment_service.AppointmentService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppointmentServiceServer).Get(ctx, req.(*AppointmentId))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppointmentService_Cancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppointmentId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppointmentServiceServer).Cancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appointment_service.AppointmentService/Cancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppointmentServiceServer).Cancel(ctx, req.(*AppointmentId))
	}
	return interceptor(ctx, in, info, handler)
}

// AppointmentService_ServiceDesc is the grpc.ServiceDesc for AppointmentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AppointmentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "appointment_service.AppointmentService",
	HandlerType: (*AppointmentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _AppointmentService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _AppointmentService_Get_Handler,
		},
		{
			MethodName: "Cancel",
			Handler:    _AppointmentService_Cancel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "appointment_service.proto",
}
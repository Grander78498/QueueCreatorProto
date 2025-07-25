// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.31.1
// source: creator.proto

package QueueCreatorProto

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
	QueueCreatorAPI_SaveUser_FullMethodName = "/QueueCreatorAPI/SaveUser"
)

// QueueCreatorAPIClient is the client API for QueueCreatorAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueueCreatorAPIClient interface {
	SaveUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
}

type queueCreatorAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewQueueCreatorAPIClient(cc grpc.ClientConnInterface) QueueCreatorAPIClient {
	return &queueCreatorAPIClient{cc}
}

func (c *queueCreatorAPIClient) SaveUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, QueueCreatorAPI_SaveUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueueCreatorAPIServer is the server API for QueueCreatorAPI service.
// All implementations must embed UnimplementedQueueCreatorAPIServer
// for forward compatibility.
type QueueCreatorAPIServer interface {
	SaveUser(context.Context, *UserRequest) (*UserResponse, error)
	mustEmbedUnimplementedQueueCreatorAPIServer()
}

// UnimplementedQueueCreatorAPIServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedQueueCreatorAPIServer struct{}

func (UnimplementedQueueCreatorAPIServer) SaveUser(context.Context, *UserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveUser not implemented")
}
func (UnimplementedQueueCreatorAPIServer) mustEmbedUnimplementedQueueCreatorAPIServer() {}
func (UnimplementedQueueCreatorAPIServer) testEmbeddedByValue()                         {}

// UnsafeQueueCreatorAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueueCreatorAPIServer will
// result in compilation errors.
type UnsafeQueueCreatorAPIServer interface {
	mustEmbedUnimplementedQueueCreatorAPIServer()
}

func RegisterQueueCreatorAPIServer(s grpc.ServiceRegistrar, srv QueueCreatorAPIServer) {
	// If the following call pancis, it indicates UnimplementedQueueCreatorAPIServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&QueueCreatorAPI_ServiceDesc, srv)
}

func _QueueCreatorAPI_SaveUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueCreatorAPIServer).SaveUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QueueCreatorAPI_SaveUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueCreatorAPIServer).SaveUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// QueueCreatorAPI_ServiceDesc is the grpc.ServiceDesc for QueueCreatorAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QueueCreatorAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "QueueCreatorAPI",
	HandlerType: (*QueueCreatorAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveUser",
			Handler:    _QueueCreatorAPI_SaveUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "creator.proto",
}

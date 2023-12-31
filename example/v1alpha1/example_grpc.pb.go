// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package examplev1alpha1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// ExampleServiceClient is the client API for ExampleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExampleServiceClient interface {
	// Heartbeat
	//
	// Provides information about the service health
	Heartbeat(ctx context.Context, in *HeartbeatRequest, opts ...grpc.CallOption) (*HeartbeatResponse, error)
	// ExamplePost
	//
	// Sets up a POST endpoint
	ExamplePost(ctx context.Context, in *ExamplePostRequest, opts ...grpc.CallOption) (*ExamplePostResponse, error)
	// ExampleGet
	//
	// Sets up a GET endpoint
	ExampleGet(ctx context.Context, in *ExampleGetRequest, opts ...grpc.CallOption) (*ExampleGetResponse, error)
	// ExampleDelete
	//
	// Sets up a DELETE endpoint
	ExampleDelete(ctx context.Context, in *ExampleDeleteRequest, opts ...grpc.CallOption) (*ExampleDeleteResponse, error)
}

type exampleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExampleServiceClient(cc grpc.ClientConnInterface) ExampleServiceClient {
	return &exampleServiceClient{cc}
}

func (c *exampleServiceClient) Heartbeat(ctx context.Context, in *HeartbeatRequest, opts ...grpc.CallOption) (*HeartbeatResponse, error) {
	out := new(HeartbeatResponse)
	err := c.cc.Invoke(ctx, "/example.v1alpha1.ExampleService/Heartbeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleServiceClient) ExamplePost(ctx context.Context, in *ExamplePostRequest, opts ...grpc.CallOption) (*ExamplePostResponse, error) {
	out := new(ExamplePostResponse)
	err := c.cc.Invoke(ctx, "/example.v1alpha1.ExampleService/ExamplePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleServiceClient) ExampleGet(ctx context.Context, in *ExampleGetRequest, opts ...grpc.CallOption) (*ExampleGetResponse, error) {
	out := new(ExampleGetResponse)
	err := c.cc.Invoke(ctx, "/example.v1alpha1.ExampleService/ExampleGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleServiceClient) ExampleDelete(ctx context.Context, in *ExampleDeleteRequest, opts ...grpc.CallOption) (*ExampleDeleteResponse, error) {
	out := new(ExampleDeleteResponse)
	err := c.cc.Invoke(ctx, "/example.v1alpha1.ExampleService/ExampleDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExampleServiceServer is the server API for ExampleService service.
// All implementations should embed UnimplementedExampleServiceServer
// for forward compatibility
type ExampleServiceServer interface {
	// Heartbeat
	//
	// Provides information about the service health
	Heartbeat(context.Context, *HeartbeatRequest) (*HeartbeatResponse, error)
	// ExamplePost
	//
	// Sets up a POST endpoint
	ExamplePost(context.Context, *ExamplePostRequest) (*ExamplePostResponse, error)
	// ExampleGet
	//
	// Sets up a GET endpoint
	ExampleGet(context.Context, *ExampleGetRequest) (*ExampleGetResponse, error)
	// ExampleDelete
	//
	// Sets up a DELETE endpoint
	ExampleDelete(context.Context, *ExampleDeleteRequest) (*ExampleDeleteResponse, error)
}

// UnimplementedExampleServiceServer should be embedded to have forward compatible implementations.
type UnimplementedExampleServiceServer struct {
}

func (UnimplementedExampleServiceServer) Heartbeat(context.Context, *HeartbeatRequest) (*HeartbeatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Heartbeat not implemented")
}
func (UnimplementedExampleServiceServer) ExamplePost(context.Context, *ExamplePostRequest) (*ExamplePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExamplePost not implemented")
}
func (UnimplementedExampleServiceServer) ExampleGet(context.Context, *ExampleGetRequest) (*ExampleGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExampleGet not implemented")
}
func (UnimplementedExampleServiceServer) ExampleDelete(context.Context, *ExampleDeleteRequest) (*ExampleDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExampleDelete not implemented")
}

// UnsafeExampleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExampleServiceServer will
// result in compilation errors.
type UnsafeExampleServiceServer interface {
	mustEmbedUnimplementedExampleServiceServer()
}

func RegisterExampleServiceServer(s *grpc.Server, srv ExampleServiceServer) {
	s.RegisterService(&_ExampleService_serviceDesc, srv)
}

func _ExampleService_Heartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartbeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServiceServer).Heartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.v1alpha1.ExampleService/Heartbeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServiceServer).Heartbeat(ctx, req.(*HeartbeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExampleService_ExamplePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExamplePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServiceServer).ExamplePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.v1alpha1.ExampleService/ExamplePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServiceServer).ExamplePost(ctx, req.(*ExamplePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExampleService_ExampleGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExampleGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServiceServer).ExampleGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.v1alpha1.ExampleService/ExampleGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServiceServer).ExampleGet(ctx, req.(*ExampleGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExampleService_ExampleDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExampleDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServiceServer).ExampleDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.v1alpha1.ExampleService/ExampleDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServiceServer).ExampleDelete(ctx, req.(*ExampleDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ExampleService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "example.v1alpha1.ExampleService",
	HandlerType: (*ExampleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Heartbeat",
			Handler:    _ExampleService_Heartbeat_Handler,
		},
		{
			MethodName: "ExamplePost",
			Handler:    _ExampleService_ExamplePost_Handler,
		},
		{
			MethodName: "ExampleGet",
			Handler:    _ExampleService_ExampleGet_Handler,
		},
		{
			MethodName: "ExampleDelete",
			Handler:    _ExampleService_ExampleDelete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example/v1alpha1/example.proto",
}

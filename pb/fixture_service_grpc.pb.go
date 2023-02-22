// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: proto/fixture_service.proto

package proto

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

// FixtureServiceClient is the client API for FixtureService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FixtureServiceClient interface {
	CreateFixture(ctx context.Context, in *CreateFixtureRequest, opts ...grpc.CallOption) (*CreateFixtureResponse, error)
	SearchFixture(ctx context.Context, in *SearchFixtureRequest, opts ...grpc.CallOption) (FixtureService_SearchFixtureClient, error)
}

type fixtureServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFixtureServiceClient(cc grpc.ClientConnInterface) FixtureServiceClient {
	return &fixtureServiceClient{cc}
}

func (c *fixtureServiceClient) CreateFixture(ctx context.Context, in *CreateFixtureRequest, opts ...grpc.CallOption) (*CreateFixtureResponse, error) {
	out := new(CreateFixtureResponse)
	err := c.cc.Invoke(ctx, "/fixtures.FixtureService/CreateFixture", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fixtureServiceClient) SearchFixture(ctx context.Context, in *SearchFixtureRequest, opts ...grpc.CallOption) (FixtureService_SearchFixtureClient, error) {
	stream, err := c.cc.NewStream(ctx, &FixtureService_ServiceDesc.Streams[0], "/fixtures.FixtureService/SearchFixture", opts...)
	if err != nil {
		return nil, err
	}
	x := &fixtureServiceSearchFixtureClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FixtureService_SearchFixtureClient interface {
	Recv() (*SearchFixtureResponse, error)
	grpc.ClientStream
}

type fixtureServiceSearchFixtureClient struct {
	grpc.ClientStream
}

func (x *fixtureServiceSearchFixtureClient) Recv() (*SearchFixtureResponse, error) {
	m := new(SearchFixtureResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FixtureServiceServer is the server API for FixtureService service.
// All implementations must embed UnimplementedFixtureServiceServer
// for forward compatibility
type FixtureServiceServer interface {
	CreateFixture(context.Context, *CreateFixtureRequest) (*CreateFixtureResponse, error)
	SearchFixture(*SearchFixtureRequest, FixtureService_SearchFixtureServer) error
	mustEmbedUnimplementedFixtureServiceServer()
}

// UnimplementedFixtureServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFixtureServiceServer struct {
}

func (UnimplementedFixtureServiceServer) CreateFixture(context.Context, *CreateFixtureRequest) (*CreateFixtureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFixture not implemented")
}
func (UnimplementedFixtureServiceServer) SearchFixture(*SearchFixtureRequest, FixtureService_SearchFixtureServer) error {
	return status.Errorf(codes.Unimplemented, "method SearchFixture not implemented")
}
func (UnimplementedFixtureServiceServer) mustEmbedUnimplementedFixtureServiceServer() {}

// UnsafeFixtureServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FixtureServiceServer will
// result in compilation errors.
type UnsafeFixtureServiceServer interface {
	mustEmbedUnimplementedFixtureServiceServer()
}

func RegisterFixtureServiceServer(s grpc.ServiceRegistrar, srv FixtureServiceServer) {
	s.RegisterService(&FixtureService_ServiceDesc, srv)
}

func _FixtureService_CreateFixture_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFixtureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FixtureServiceServer).CreateFixture(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fixtures.FixtureService/CreateFixture",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FixtureServiceServer).CreateFixture(ctx, req.(*CreateFixtureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FixtureService_SearchFixture_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SearchFixtureRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FixtureServiceServer).SearchFixture(m, &fixtureServiceSearchFixtureServer{stream})
}

type FixtureService_SearchFixtureServer interface {
	Send(*SearchFixtureResponse) error
	grpc.ServerStream
}

type fixtureServiceSearchFixtureServer struct {
	grpc.ServerStream
}

func (x *fixtureServiceSearchFixtureServer) Send(m *SearchFixtureResponse) error {
	return x.ServerStream.SendMsg(m)
}

// FixtureService_ServiceDesc is the grpc.ServiceDesc for FixtureService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FixtureService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fixtures.FixtureService",
	HandlerType: (*FixtureServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFixture",
			Handler:    _FixtureService_CreateFixture_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SearchFixture",
			Handler:       _FixtureService_SearchFixture_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/fixture_service.proto",
}
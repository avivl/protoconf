// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protoconfservice

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

// ProtoconfServiceClient is the client API for ProtoconfService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProtoconfServiceClient interface {
	SubscribeForConfig(ctx context.Context, in *ConfigSubscriptionRequest, opts ...grpc.CallOption) (ProtoconfService_SubscribeForConfigClient, error)
}

type protoconfServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProtoconfServiceClient(cc grpc.ClientConnInterface) ProtoconfServiceClient {
	return &protoconfServiceClient{cc}
}

func (c *protoconfServiceClient) SubscribeForConfig(ctx context.Context, in *ConfigSubscriptionRequest, opts ...grpc.CallOption) (ProtoconfService_SubscribeForConfigClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProtoconfService_ServiceDesc.Streams[0], "/v1.ProtoconfService/SubscribeForConfig", opts...)
	if err != nil {
		return nil, err
	}
	x := &protoconfServiceSubscribeForConfigClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ProtoconfService_SubscribeForConfigClient interface {
	Recv() (*ConfigUpdate, error)
	grpc.ClientStream
}

type protoconfServiceSubscribeForConfigClient struct {
	grpc.ClientStream
}

func (x *protoconfServiceSubscribeForConfigClient) Recv() (*ConfigUpdate, error) {
	m := new(ConfigUpdate)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProtoconfServiceServer is the server API for ProtoconfService service.
// All implementations must embed UnimplementedProtoconfServiceServer
// for forward compatibility
type ProtoconfServiceServer interface {
	SubscribeForConfig(*ConfigSubscriptionRequest, ProtoconfService_SubscribeForConfigServer) error
	mustEmbedUnimplementedProtoconfServiceServer()
}

// UnimplementedProtoconfServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProtoconfServiceServer struct {
}

func (UnimplementedProtoconfServiceServer) SubscribeForConfig(*ConfigSubscriptionRequest, ProtoconfService_SubscribeForConfigServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeForConfig not implemented")
}
func (UnimplementedProtoconfServiceServer) mustEmbedUnimplementedProtoconfServiceServer() {}

// UnsafeProtoconfServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProtoconfServiceServer will
// result in compilation errors.
type UnsafeProtoconfServiceServer interface {
	mustEmbedUnimplementedProtoconfServiceServer()
}

func RegisterProtoconfServiceServer(s grpc.ServiceRegistrar, srv ProtoconfServiceServer) {
	s.RegisterService(&ProtoconfService_ServiceDesc, srv)
}

func _ProtoconfService_SubscribeForConfig_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ConfigSubscriptionRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProtoconfServiceServer).SubscribeForConfig(m, &protoconfServiceSubscribeForConfigServer{stream})
}

type ProtoconfService_SubscribeForConfigServer interface {
	Send(*ConfigUpdate) error
	grpc.ServerStream
}

type protoconfServiceSubscribeForConfigServer struct {
	grpc.ServerStream
}

func (x *protoconfServiceSubscribeForConfigServer) Send(m *ConfigUpdate) error {
	return x.ServerStream.SendMsg(m)
}

// ProtoconfService_ServiceDesc is the grpc.ServiceDesc for ProtoconfService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProtoconfService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.ProtoconfService",
	HandlerType: (*ProtoconfServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeForConfig",
			Handler:       _ProtoconfService_SubscribeForConfig_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protoconf_service.proto",
}

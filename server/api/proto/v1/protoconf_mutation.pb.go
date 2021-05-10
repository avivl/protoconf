// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.9.0
// source: server/api/proto/v1/protoconf_mutation.proto

package v1

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	v1 "github.com/protoconf/protoconf/datatypes/proto/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ConfigMutationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path           string             `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Value          *v1.ProtoconfValue `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	ScriptMetadata string             `protobuf:"bytes,3,opt,name=script_metadata,json=scriptMetadata,proto3" json:"script_metadata,omitempty"`
}

func (x *ConfigMutationRequest) Reset() {
	*x = ConfigMutationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_api_proto_v1_protoconf_mutation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigMutationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigMutationRequest) ProtoMessage() {}

func (x *ConfigMutationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_api_proto_v1_protoconf_mutation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigMutationRequest.ProtoReflect.Descriptor instead.
func (*ConfigMutationRequest) Descriptor() ([]byte, []int) {
	return file_server_api_proto_v1_protoconf_mutation_proto_rawDescGZIP(), []int{0}
}

func (x *ConfigMutationRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ConfigMutationRequest) GetValue() *v1.ProtoconfValue {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *ConfigMutationRequest) GetScriptMetadata() string {
	if x != nil {
		return x.ScriptMetadata
	}
	return ""
}

type ConfigMutationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConfigMutationResponse) Reset() {
	*x = ConfigMutationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_api_proto_v1_protoconf_mutation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigMutationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigMutationResponse) ProtoMessage() {}

func (x *ConfigMutationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_server_api_proto_v1_protoconf_mutation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigMutationResponse.ProtoReflect.Descriptor instead.
func (*ConfigMutationResponse) Descriptor() ([]byte, []int) {
	return file_server_api_proto_v1_protoconf_mutation_proto_rawDescGZIP(), []int{1}
}

var File_server_api_proto_v1_protoconf_mutation_proto protoreflect.FileDescriptor

var file_server_api_proto_v1_protoconf_mutation_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6e, 0x66, 0x5f,
	0x6d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x76, 0x31, 0x1a, 0x28, 0x64, 0x61, 0x74, 0x61, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6e, 0x66,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7e, 0x0a, 0x15,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x28, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6e, 0x66, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x5f, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x18, 0x0a, 0x16,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x61, 0x0a, 0x18, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6e, 0x66, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x45, 0x0a, 0x0c, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x19, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4d, 0x75,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1d, 0x0a, 0x1b, 0x63, 0x6f, 0x6d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_server_api_proto_v1_protoconf_mutation_proto_rawDescOnce sync.Once
	file_server_api_proto_v1_protoconf_mutation_proto_rawDescData = file_server_api_proto_v1_protoconf_mutation_proto_rawDesc
)

func file_server_api_proto_v1_protoconf_mutation_proto_rawDescGZIP() []byte {
	file_server_api_proto_v1_protoconf_mutation_proto_rawDescOnce.Do(func() {
		file_server_api_proto_v1_protoconf_mutation_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_api_proto_v1_protoconf_mutation_proto_rawDescData)
	})
	return file_server_api_proto_v1_protoconf_mutation_proto_rawDescData
}

var file_server_api_proto_v1_protoconf_mutation_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_server_api_proto_v1_protoconf_mutation_proto_goTypes = []interface{}{
	(*ConfigMutationRequest)(nil),  // 0: v1.ConfigMutationRequest
	(*ConfigMutationResponse)(nil), // 1: v1.ConfigMutationResponse
	(*v1.ProtoconfValue)(nil),      // 2: v1.ProtoconfValue
}
var file_server_api_proto_v1_protoconf_mutation_proto_depIdxs = []int32{
	2, // 0: v1.ConfigMutationRequest.value:type_name -> v1.ProtoconfValue
	0, // 1: v1.ProtoconfMutationService.MutateConfig:input_type -> v1.ConfigMutationRequest
	1, // 2: v1.ProtoconfMutationService.MutateConfig:output_type -> v1.ConfigMutationResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_server_api_proto_v1_protoconf_mutation_proto_init() }
func file_server_api_proto_v1_protoconf_mutation_proto_init() {
	if File_server_api_proto_v1_protoconf_mutation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_server_api_proto_v1_protoconf_mutation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigMutationRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_server_api_proto_v1_protoconf_mutation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigMutationResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_server_api_proto_v1_protoconf_mutation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_server_api_proto_v1_protoconf_mutation_proto_goTypes,
		DependencyIndexes: file_server_api_proto_v1_protoconf_mutation_proto_depIdxs,
		MessageInfos:      file_server_api_proto_v1_protoconf_mutation_proto_msgTypes,
	}.Build()
	File_server_api_proto_v1_protoconf_mutation_proto = out.File
	file_server_api_proto_v1_protoconf_mutation_proto_rawDesc = nil
	file_server_api_proto_v1_protoconf_mutation_proto_goTypes = nil
	file_server_api_proto_v1_protoconf_mutation_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProtoconfMutationServiceClient is the client API for ProtoconfMutationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProtoconfMutationServiceClient interface {
	MutateConfig(ctx context.Context, in *ConfigMutationRequest, opts ...grpc.CallOption) (*ConfigMutationResponse, error)
}

type protoconfMutationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProtoconfMutationServiceClient(cc grpc.ClientConnInterface) ProtoconfMutationServiceClient {
	return &protoconfMutationServiceClient{cc}
}

func (c *protoconfMutationServiceClient) MutateConfig(ctx context.Context, in *ConfigMutationRequest, opts ...grpc.CallOption) (*ConfigMutationResponse, error) {
	out := new(ConfigMutationResponse)
	err := c.cc.Invoke(ctx, "/v1.ProtoconfMutationService/MutateConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProtoconfMutationServiceServer is the server API for ProtoconfMutationService service.
type ProtoconfMutationServiceServer interface {
	MutateConfig(context.Context, *ConfigMutationRequest) (*ConfigMutationResponse, error)
}

// UnimplementedProtoconfMutationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProtoconfMutationServiceServer struct {
}

func (*UnimplementedProtoconfMutationServiceServer) MutateConfig(context.Context, *ConfigMutationRequest) (*ConfigMutationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateConfig not implemented")
}

func RegisterProtoconfMutationServiceServer(s *grpc.Server, srv ProtoconfMutationServiceServer) {
	s.RegisterService(&_ProtoconfMutationService_serviceDesc, srv)
}

func _ProtoconfMutationService_MutateConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigMutationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtoconfMutationServiceServer).MutateConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ProtoconfMutationService/MutateConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtoconfMutationServiceServer).MutateConfig(ctx, req.(*ConfigMutationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProtoconfMutationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.ProtoconfMutationService",
	HandlerType: (*ProtoconfMutationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateConfig",
			Handler:    _ProtoconfMutationService_MutateConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server/api/proto/v1/protoconf_mutation.proto",
}
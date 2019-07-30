// Code generated by protoc-gen-go. DO NOT EDIT.
// source: example.proto

package v1 // import "github.com/dio/grpctranscoding/api/example/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ExampleClient is the client API for Example service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExampleClient interface {
	Sample(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
}

type exampleClient struct {
	cc *grpc.ClientConn
}

func NewExampleClient(cc *grpc.ClientConn) ExampleClient {
	return &exampleClient{cc}
}

func (c *exampleClient) Sample(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/grpctranscoding.api.example.v1.Example/Sample", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExampleServer is the server API for Example service.
type ExampleServer interface {
	Sample(context.Context, *empty.Empty) (*empty.Empty, error)
}

func RegisterExampleServer(s *grpc.Server, srv ExampleServer) {
	s.RegisterService(&_Example_serviceDesc, srv)
}

func _Example_Sample_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServer).Sample(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpctranscoding.api.example.v1.Example/Sample",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServer).Sample(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Example_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpctranscoding.api.example.v1.Example",
	HandlerType: (*ExampleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sample",
			Handler:    _Example_Sample_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example.proto",
}

func init() { proto.RegisterFile("example.proto", fileDescriptor_example_21efed594635f660) }

var fileDescriptor_example_21efed594635f660 = []byte{
	// 184 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x4b, 0x2f, 0x2a, 0x48, 0x2e,
	0x29, 0x4a, 0xcc, 0x2b, 0x4e, 0xce, 0x4f, 0xc9, 0xcc, 0x4b, 0xd7, 0x4b, 0x2c, 0xc8, 0xd4, 0x83,
	0x29, 0x29, 0x33, 0x94, 0x92, 0x49, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4,
	0x4f, 0xcc, 0xcb, 0xcb, 0x2f, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0x86, 0xe8, 0x96, 0x92, 0x86,
	0xca, 0x82, 0x79, 0x49, 0xa5, 0x69, 0xfa, 0xa9, 0xb9, 0x05, 0x25, 0x95, 0x10, 0x49, 0xa3, 0x10,
	0x2e, 0x76, 0x57, 0x88, 0x41, 0x42, 0x9e, 0x5c, 0x6c, 0xc1, 0x10, 0x96, 0x98, 0x1e, 0x44, 0x8b,
	0x1e, 0x4c, 0x8b, 0x9e, 0x2b, 0x48, 0x8b, 0x14, 0x0e, 0x71, 0x25, 0xfe, 0xa6, 0xcb, 0x4f, 0x26,
	0x33, 0x71, 0x0a, 0xb1, 0xeb, 0x17, 0x83, 0x0d, 0x70, 0xd2, 0x8f, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xc9, 0xcc, 0xd7, 0x47, 0xf3, 0x01, 0xd8, 0xa9,
	0x50, 0x1f, 0xe8, 0x97, 0x19, 0x26, 0xb1, 0x81, 0x4d, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x02, 0xed, 0x18, 0xcb, 0xf9, 0x00, 0x00, 0x00,
}

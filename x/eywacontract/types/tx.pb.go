// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: eywacontract/eywacontract/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("eywacontract/eywacontract/tx.proto", fileDescriptor_1f1d0daca0114d22)
}

var fileDescriptor_1f1d0daca0114d22 = []byte{
	// 109 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4a, 0xad, 0x2c, 0x4f,
	0x4c, 0xce, 0xcf, 0x2b, 0x29, 0x4a, 0x4c, 0x2e, 0xd1, 0x47, 0xe1, 0x94, 0x54, 0xe8, 0x15, 0x14,
	0xe5, 0x97, 0xe4, 0x0b, 0x49, 0x22, 0x0b, 0xeb, 0x21, 0x73, 0x8c, 0x58, 0xb9, 0x98, 0x7d, 0x8b,
	0xd3, 0x9d, 0x6c, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6,
	0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x0a, 0x6c, 0xbe,
	0x2e, 0xdc, 0xcc, 0x0a, 0x34, 0x2b, 0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0xd6, 0x18, 0x03,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x3c, 0xb5, 0xf5, 0xe1, 0x8c, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "eywacontract.eywacontract.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "eywacontract/eywacontract/tx.proto",
}

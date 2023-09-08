// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: eywacontract/eywacontract/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
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

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_90dddaedbed5ef05, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_90dddaedbed5ef05, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

type QueryGetRegisterTypeRequest struct {
	Index string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
}

func (m *QueryGetRegisterTypeRequest) Reset()         { *m = QueryGetRegisterTypeRequest{} }
func (m *QueryGetRegisterTypeRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetRegisterTypeRequest) ProtoMessage()    {}
func (*QueryGetRegisterTypeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_90dddaedbed5ef05, []int{2}
}
func (m *QueryGetRegisterTypeRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetRegisterTypeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetRegisterTypeRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetRegisterTypeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetRegisterTypeRequest.Merge(m, src)
}
func (m *QueryGetRegisterTypeRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetRegisterTypeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetRegisterTypeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetRegisterTypeRequest proto.InternalMessageInfo

func (m *QueryGetRegisterTypeRequest) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

type QueryGetRegisterTypeResponse struct {
	RegisterType RegisterType `protobuf:"bytes,1,opt,name=registerType,proto3" json:"registerType"`
}

func (m *QueryGetRegisterTypeResponse) Reset()         { *m = QueryGetRegisterTypeResponse{} }
func (m *QueryGetRegisterTypeResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetRegisterTypeResponse) ProtoMessage()    {}
func (*QueryGetRegisterTypeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_90dddaedbed5ef05, []int{3}
}
func (m *QueryGetRegisterTypeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetRegisterTypeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetRegisterTypeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetRegisterTypeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetRegisterTypeResponse.Merge(m, src)
}
func (m *QueryGetRegisterTypeResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetRegisterTypeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetRegisterTypeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetRegisterTypeResponse proto.InternalMessageInfo

func (m *QueryGetRegisterTypeResponse) GetRegisterType() RegisterType {
	if m != nil {
		return m.RegisterType
	}
	return RegisterType{}
}

type QueryAllRegisterTypeRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllRegisterTypeRequest) Reset()         { *m = QueryAllRegisterTypeRequest{} }
func (m *QueryAllRegisterTypeRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllRegisterTypeRequest) ProtoMessage()    {}
func (*QueryAllRegisterTypeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_90dddaedbed5ef05, []int{4}
}
func (m *QueryAllRegisterTypeRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllRegisterTypeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllRegisterTypeRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllRegisterTypeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllRegisterTypeRequest.Merge(m, src)
}
func (m *QueryAllRegisterTypeRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllRegisterTypeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllRegisterTypeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllRegisterTypeRequest proto.InternalMessageInfo

func (m *QueryAllRegisterTypeRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllRegisterTypeResponse struct {
	RegisterType []RegisterType      `protobuf:"bytes,1,rep,name=registerType,proto3" json:"registerType"`
	Pagination   *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllRegisterTypeResponse) Reset()         { *m = QueryAllRegisterTypeResponse{} }
func (m *QueryAllRegisterTypeResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllRegisterTypeResponse) ProtoMessage()    {}
func (*QueryAllRegisterTypeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_90dddaedbed5ef05, []int{5}
}
func (m *QueryAllRegisterTypeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllRegisterTypeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllRegisterTypeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllRegisterTypeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllRegisterTypeResponse.Merge(m, src)
}
func (m *QueryAllRegisterTypeResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllRegisterTypeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllRegisterTypeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllRegisterTypeResponse proto.InternalMessageInfo

func (m *QueryAllRegisterTypeResponse) GetRegisterType() []RegisterType {
	if m != nil {
		return m.RegisterType
	}
	return nil
}

func (m *QueryAllRegisterTypeResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "eywacontract.eywacontract.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "eywacontract.eywacontract.QueryParamsResponse")
	proto.RegisterType((*QueryGetRegisterTypeRequest)(nil), "eywacontract.eywacontract.QueryGetRegisterTypeRequest")
	proto.RegisterType((*QueryGetRegisterTypeResponse)(nil), "eywacontract.eywacontract.QueryGetRegisterTypeResponse")
	proto.RegisterType((*QueryAllRegisterTypeRequest)(nil), "eywacontract.eywacontract.QueryAllRegisterTypeRequest")
	proto.RegisterType((*QueryAllRegisterTypeResponse)(nil), "eywacontract.eywacontract.QueryAllRegisterTypeResponse")
}

func init() {
	proto.RegisterFile("eywacontract/eywacontract/query.proto", fileDescriptor_90dddaedbed5ef05)
}

var fileDescriptor_90dddaedbed5ef05 = []byte{
	// 494 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0x4f, 0x6b, 0x14, 0x31,
	0x18, 0xc6, 0x27, 0xb5, 0x5d, 0x30, 0x16, 0x84, 0xb8, 0x07, 0x1d, 0xcb, 0xa8, 0x41, 0xab, 0x56,
	0x36, 0x61, 0xba, 0xa0, 0x08, 0x82, 0xb4, 0x07, 0x7b, 0x6d, 0x07, 0xf1, 0xe0, 0x45, 0xb2, 0x6b,
	0x18, 0x06, 0xa6, 0x93, 0xd9, 0x49, 0xaa, 0x1d, 0xc4, 0x8b, 0x9f, 0x40, 0xec, 0x67, 0xf1, 0xd2,
	0x4f, 0xd0, 0x63, 0xc1, 0x8b, 0x27, 0x91, 0x5d, 0x3f, 0x88, 0x4c, 0x92, 0x62, 0x86, 0xce, 0x9f,
	0x4a, 0x6f, 0x9b, 0xc9, 0xfb, 0x3c, 0xcf, 0x2f, 0x79, 0xdf, 0x2c, 0x7c, 0xc0, 0xcb, 0x8f, 0x6c,
	0x2a, 0x32, 0x55, 0xb0, 0xa9, 0xa2, 0xb5, 0xc5, 0xec, 0x80, 0x17, 0x25, 0xc9, 0x0b, 0xa1, 0x04,
	0xba, 0xe5, 0xee, 0x10, 0x77, 0xe1, 0x0f, 0x63, 0x11, 0x0b, 0x5d, 0x45, 0xab, 0x5f, 0x46, 0xe0,
	0xaf, 0xc5, 0x42, 0xc4, 0x29, 0xa7, 0x2c, 0x4f, 0x28, 0xcb, 0x32, 0xa1, 0x98, 0x4a, 0x44, 0x26,
	0xed, 0xee, 0xc6, 0x54, 0xc8, 0x7d, 0x21, 0xe9, 0x84, 0x49, 0x6e, 0x72, 0xe8, 0x87, 0x70, 0xc2,
	0x15, 0x0b, 0x69, 0xce, 0xe2, 0x24, 0xd3, 0xc5, 0xb6, 0x76, 0xbd, 0x9d, 0x30, 0x67, 0x05, 0xdb,
	0x3f, 0xf3, 0x1c, 0xb5, 0xd7, 0x15, 0x3c, 0x4e, 0xa4, 0xe2, 0xc5, 0x3b, 0x55, 0xe6, 0xdc, 0x94,
	0xe3, 0x21, 0x44, 0x7b, 0x55, 0xf0, 0xae, 0xf6, 0x88, 0xf8, 0xec, 0x80, 0x4b, 0x85, 0xdf, 0xc0,
	0x1b, 0xb5, 0xaf, 0x32, 0x17, 0x99, 0xe4, 0xe8, 0x25, 0x1c, 0x98, 0xac, 0x9b, 0xe0, 0x2e, 0x78,
	0x74, 0x6d, 0xf3, 0x1e, 0x69, 0xbd, 0x0f, 0x62, 0xa4, 0xdb, 0xcb, 0x27, 0xbf, 0xee, 0x78, 0x91,
	0x95, 0xe1, 0x31, 0xbc, 0xad, 0x7d, 0x77, 0xb8, 0x8a, 0x2c, 0xcc, 0xeb, 0x32, 0xe7, 0x36, 0x16,
	0x0d, 0xe1, 0x4a, 0x92, 0xbd, 0xe7, 0x87, 0xda, 0xfe, 0x6a, 0x64, 0x16, 0x78, 0x06, 0xd7, 0x9a,
	0x45, 0x96, 0x6a, 0x0f, 0xae, 0x16, 0xce, 0x77, 0xcb, 0xf6, 0xb0, 0x83, 0xcd, 0xb5, 0xb1, 0x84,
	0x35, 0x0b, 0xcc, 0x2d, 0xe7, 0x56, 0x9a, 0x36, 0x71, 0xbe, 0x82, 0xf0, 0x5f, 0x7f, 0x6c, 0xde,
	0x3a, 0x31, 0xcd, 0x24, 0x55, 0x33, 0x89, 0x19, 0x1a, 0xdb, 0x4c, 0xb2, 0xcb, 0xe2, 0x33, 0x6d,
	0xe4, 0x28, 0xf1, 0x31, 0xb0, 0x47, 0x3b, 0x97, 0xd3, 0x7a, 0xb4, 0x2b, 0x97, 0x3c, 0x1a, 0xda,
	0xa9, 0xb1, 0x2f, 0xd9, 0xbb, 0xea, 0x63, 0x37, 0x3c, 0x2e, 0xfc, 0xe6, 0xb7, 0x65, 0xb8, 0xa2,
	0xe1, 0xd1, 0x11, 0x80, 0x03, 0xd3, 0x6e, 0x34, 0xea, 0x40, 0x3b, 0x3f, 0x67, 0x3e, 0xb9, 0x68,
	0xb9, 0xc9, 0xc7, 0x1b, 0x5f, 0x7e, 0xfc, 0x39, 0x5a, 0xba, 0x8f, 0xb0, 0x1e, 0xec, 0x51, 0xd7,
	0x73, 0x40, 0xc7, 0x00, 0xae, 0xba, 0xb7, 0x81, 0x9e, 0xf6, 0x85, 0x35, 0x4f, 0xa5, 0xff, 0xec,
	0xbf, 0x75, 0x96, 0xf6, 0xb9, 0xa6, 0x1d, 0xa3, 0xb0, 0x8b, 0xb6, 0xf6, 0x28, 0xe9, 0x27, 0x3d,
	0xf2, 0x9f, 0xd1, 0x77, 0x00, 0xaf, 0xbb, 0x9e, 0x5b, 0x69, 0xda, 0xcf, 0xdf, 0x3c, 0xad, 0xfd,
	0xfc, 0x2d, 0xd3, 0x87, 0x43, 0xcd, 0xff, 0x04, 0x3d, 0xbe, 0x30, 0xff, 0xf6, 0x8b, 0x93, 0x79,
	0x00, 0x4e, 0xe7, 0x01, 0xf8, 0x3d, 0x0f, 0xc0, 0xd7, 0x45, 0xe0, 0x9d, 0x2e, 0x02, 0xef, 0xe7,
	0x22, 0xf0, 0xde, 0xe2, 0xba, 0xc7, 0x61, 0xdd, 0xa5, 0x12, 0xcb, 0xc9, 0x40, 0xff, 0x27, 0x8d,
	0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x5e, 0x0a, 0x2c, 0x3a, 0x8e, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a list of RegisterType items.
	RegisterType(ctx context.Context, in *QueryGetRegisterTypeRequest, opts ...grpc.CallOption) (*QueryGetRegisterTypeResponse, error)
	RegisterTypeAll(ctx context.Context, in *QueryAllRegisterTypeRequest, opts ...grpc.CallOption) (*QueryAllRegisterTypeResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/eywacontract.eywacontract.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) RegisterType(ctx context.Context, in *QueryGetRegisterTypeRequest, opts ...grpc.CallOption) (*QueryGetRegisterTypeResponse, error) {
	out := new(QueryGetRegisterTypeResponse)
	err := c.cc.Invoke(ctx, "/eywacontract.eywacontract.Query/RegisterType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) RegisterTypeAll(ctx context.Context, in *QueryAllRegisterTypeRequest, opts ...grpc.CallOption) (*QueryAllRegisterTypeResponse, error) {
	out := new(QueryAllRegisterTypeResponse)
	err := c.cc.Invoke(ctx, "/eywacontract.eywacontract.Query/RegisterTypeAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a list of RegisterType items.
	RegisterType(context.Context, *QueryGetRegisterTypeRequest) (*QueryGetRegisterTypeResponse, error)
	RegisterTypeAll(context.Context, *QueryAllRegisterTypeRequest) (*QueryAllRegisterTypeResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) RegisterType(ctx context.Context, req *QueryGetRegisterTypeRequest) (*QueryGetRegisterTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterType not implemented")
}
func (*UnimplementedQueryServer) RegisterTypeAll(ctx context.Context, req *QueryAllRegisterTypeRequest) (*QueryAllRegisterTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterTypeAll not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eywacontract.eywacontract.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_RegisterType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetRegisterTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).RegisterType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eywacontract.eywacontract.Query/RegisterType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).RegisterType(ctx, req.(*QueryGetRegisterTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_RegisterTypeAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllRegisterTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).RegisterTypeAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eywacontract.eywacontract.Query/RegisterTypeAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).RegisterTypeAll(ctx, req.(*QueryAllRegisterTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "eywacontract.eywacontract.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "RegisterType",
			Handler:    _Query_RegisterType_Handler,
		},
		{
			MethodName: "RegisterTypeAll",
			Handler:    _Query_RegisterTypeAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eywacontract/eywacontract/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryGetRegisterTypeRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetRegisterTypeRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetRegisterTypeRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetRegisterTypeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetRegisterTypeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetRegisterTypeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.RegisterType.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryAllRegisterTypeRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllRegisterTypeRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllRegisterTypeRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllRegisterTypeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllRegisterTypeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllRegisterTypeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.RegisterType) > 0 {
		for iNdEx := len(m.RegisterType) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RegisterType[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryGetRegisterTypeRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetRegisterTypeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.RegisterType.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryAllRegisterTypeRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllRegisterTypeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.RegisterType) > 0 {
		for _, e := range m.RegisterType {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGetRegisterTypeRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetRegisterTypeRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetRegisterTypeRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGetRegisterTypeResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetRegisterTypeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetRegisterTypeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegisterType", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RegisterType.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAllRegisterTypeRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAllRegisterTypeRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllRegisterTypeRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAllRegisterTypeResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAllRegisterTypeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllRegisterTypeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegisterType", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RegisterType = append(m.RegisterType, RegisterType{})
			if err := m.RegisterType[len(m.RegisterType)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)

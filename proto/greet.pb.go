// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/greet.proto

package greetpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GreetReq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetReq) Reset()         { *m = GreetReq{} }
func (m *GreetReq) String() string { return proto.CompactTextString(m) }
func (*GreetReq) ProtoMessage()    {}
func (*GreetReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_95ca2ad3f55a58e3, []int{0}
}

func (m *GreetReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetReq.Unmarshal(m, b)
}
func (m *GreetReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetReq.Marshal(b, m, deterministic)
}
func (m *GreetReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetReq.Merge(m, src)
}
func (m *GreetReq) XXX_Size() int {
	return xxx_messageInfo_GreetReq.Size(m)
}
func (m *GreetReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetReq.DiscardUnknown(m)
}

var xxx_messageInfo_GreetReq proto.InternalMessageInfo

func (m *GreetReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GreetRes struct {
	Greeting             string   `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetRes) Reset()         { *m = GreetRes{} }
func (m *GreetRes) String() string { return proto.CompactTextString(m) }
func (*GreetRes) ProtoMessage()    {}
func (*GreetRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_95ca2ad3f55a58e3, []int{1}
}

func (m *GreetRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetRes.Unmarshal(m, b)
}
func (m *GreetRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetRes.Marshal(b, m, deterministic)
}
func (m *GreetRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetRes.Merge(m, src)
}
func (m *GreetRes) XXX_Size() int {
	return xxx_messageInfo_GreetRes.Size(m)
}
func (m *GreetRes) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetRes.DiscardUnknown(m)
}

var xxx_messageInfo_GreetRes proto.InternalMessageInfo

func (m *GreetRes) GetGreeting() string {
	if m != nil {
		return m.Greeting
	}
	return ""
}

func init() {
	proto.RegisterType((*GreetReq)(nil), "greet.GreetReq")
	proto.RegisterType((*GreetRes)(nil), "greet.GreetRes")
}

func init() { proto.RegisterFile("proto/greet.proto", fileDescriptor_95ca2ad3f55a58e3) }

var fileDescriptor_95ca2ad3f55a58e3 = []byte{
	// 133 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0x2f, 0x4a, 0x4d, 0x2d, 0xd1, 0x03, 0xb3, 0x85, 0x58, 0xc1, 0x1c, 0x25, 0x39,
	0x2e, 0x0e, 0x77, 0x10, 0x23, 0x28, 0xb5, 0x50, 0x48, 0x88, 0x8b, 0x25, 0x2f, 0x31, 0x37, 0x55,
	0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xcc, 0x56, 0x52, 0x83, 0xcb, 0x17, 0x0b, 0x49, 0x71,
	0x71, 0x80, 0x35, 0x65, 0xe6, 0xa5, 0x43, 0xd5, 0xc0, 0xf9, 0x46, 0xd6, 0x5c, 0x3c, 0x60, 0x75,
	0xc1, 0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9, 0x42, 0xda, 0x5c, 0xac, 0x60, 0xbe, 0x10, 0xbf, 0x1e,
	0xc4, 0x56, 0x98, 0x2d, 0x52, 0x68, 0x02, 0xc5, 0x4a, 0x0c, 0x4e, 0x9c, 0x51, 0xec, 0x60, 0xb1,
	0x82, 0xa4, 0x24, 0x36, 0xb0, 0xeb, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x0d, 0xe3, 0x37,
	0xc0, 0xb2, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreetServiceClient is the client API for GreetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreetServiceClient interface {
	Greet(ctx context.Context, in *GreetReq, opts ...grpc.CallOption) (*GreetRes, error)
}

type greetServiceClient struct {
	cc *grpc.ClientConn
}

func NewGreetServiceClient(cc *grpc.ClientConn) GreetServiceClient {
	return &greetServiceClient{cc}
}

func (c *greetServiceClient) Greet(ctx context.Context, in *GreetReq, opts ...grpc.CallOption) (*GreetRes, error) {
	out := new(GreetRes)
	err := c.cc.Invoke(ctx, "/greet.GreetService/Greet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreetServiceServer is the server API for GreetService service.
type GreetServiceServer interface {
	Greet(context.Context, *GreetReq) (*GreetRes, error)
}

// UnimplementedGreetServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGreetServiceServer struct {
}

func (*UnimplementedGreetServiceServer) Greet(ctx context.Context, req *GreetReq) (*GreetRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Greet not implemented")
}

func RegisterGreetServiceServer(s *grpc.Server, srv GreetServiceServer) {
	s.RegisterService(&_GreetService_serviceDesc, srv)
}

func _GreetService_Greet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetServiceServer).Greet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greet.GreetService/Greet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetServiceServer).Greet(ctx, req.(*GreetReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _GreetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "greet.GreetService",
	HandlerType: (*GreetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greet",
			Handler:    _GreetService_Greet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/greet.proto",
}

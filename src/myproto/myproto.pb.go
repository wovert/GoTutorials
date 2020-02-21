// Code generated by protoc-gen-go. DO NOT EDIT.
// source: myproto.proto

package myproto

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

// 客户端发送给服务端
type HelloReq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloReq) Reset()         { *m = HelloReq{} }
func (m *HelloReq) String() string { return proto.CompactTextString(m) }
func (*HelloReq) ProtoMessage()    {}
func (*HelloReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_04877df457807402, []int{0}
}

func (m *HelloReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReq.Unmarshal(m, b)
}
func (m *HelloReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReq.Marshal(b, m, deterministic)
}
func (m *HelloReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReq.Merge(m, src)
}
func (m *HelloReq) XXX_Size() int {
	return xxx_messageInfo_HelloReq.Size(m)
}
func (m *HelloReq) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReq.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReq proto.InternalMessageInfo

func (m *HelloReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// 服务端返回给客户端
type HelloRes struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRes) Reset()         { *m = HelloRes{} }
func (m *HelloRes) String() string { return proto.CompactTextString(m) }
func (*HelloRes) ProtoMessage()    {}
func (*HelloRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_04877df457807402, []int{1}
}

func (m *HelloRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRes.Unmarshal(m, b)
}
func (m *HelloRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRes.Marshal(b, m, deterministic)
}
func (m *HelloRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRes.Merge(m, src)
}
func (m *HelloRes) XXX_Size() int {
	return xxx_messageInfo_HelloRes.Size(m)
}
func (m *HelloRes) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRes.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRes proto.InternalMessageInfo

func (m *HelloRes) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

// 客户端发送给服务端
type NameReq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NameReq) Reset()         { *m = NameReq{} }
func (m *NameReq) String() string { return proto.CompactTextString(m) }
func (*NameReq) ProtoMessage()    {}
func (*NameReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_04877df457807402, []int{2}
}

func (m *NameReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NameReq.Unmarshal(m, b)
}
func (m *NameReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NameReq.Marshal(b, m, deterministic)
}
func (m *NameReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NameReq.Merge(m, src)
}
func (m *NameReq) XXX_Size() int {
	return xxx_messageInfo_NameReq.Size(m)
}
func (m *NameReq) XXX_DiscardUnknown() {
	xxx_messageInfo_NameReq.DiscardUnknown(m)
}

var xxx_messageInfo_NameReq proto.InternalMessageInfo

func (m *NameReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// 服务端返回给客户端
type NameRes struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NameRes) Reset()         { *m = NameRes{} }
func (m *NameRes) String() string { return proto.CompactTextString(m) }
func (*NameRes) ProtoMessage()    {}
func (*NameRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_04877df457807402, []int{3}
}

func (m *NameRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NameRes.Unmarshal(m, b)
}
func (m *NameRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NameRes.Marshal(b, m, deterministic)
}
func (m *NameRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NameRes.Merge(m, src)
}
func (m *NameRes) XXX_Size() int {
	return xxx_messageInfo_NameRes.Size(m)
}
func (m *NameRes) XXX_DiscardUnknown() {
	xxx_messageInfo_NameRes.DiscardUnknown(m)
}

var xxx_messageInfo_NameRes proto.InternalMessageInfo

func (m *NameRes) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloReq)(nil), "myproto.HelloReq")
	proto.RegisterType((*HelloRes)(nil), "myproto.HelloRes")
	proto.RegisterType((*NameReq)(nil), "myproto.NameReq")
	proto.RegisterType((*NameRes)(nil), "myproto.NameRes")
}

func init() { proto.RegisterFile("myproto.proto", fileDescriptor_04877df457807402) }

var fileDescriptor_04877df457807402 = []byte{
	// 160 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0xad, 0x2c, 0x28,
	0xca, 0x2f, 0xc9, 0xd7, 0x03, 0x93, 0x42, 0xec, 0x50, 0xae, 0x92, 0x1c, 0x17, 0x87, 0x47, 0x6a,
	0x4e, 0x4e, 0x7e, 0x50, 0x6a, 0xa1, 0x90, 0x10, 0x17, 0x4b, 0x5e, 0x62, 0x6e, 0xaa, 0x04, 0xa3,
	0x02, 0xa3, 0x06, 0x67, 0x10, 0x98, 0xad, 0x24, 0x03, 0x97, 0x2f, 0x16, 0x12, 0xe0, 0x62, 0xce,
	0x2d, 0x4e, 0x87, 0x4a, 0x83, 0x98, 0x4a, 0xb2, 0x5c, 0xec, 0x7e, 0x89, 0xb9, 0xa9, 0xb8, 0x34,
	0x4b, 0xc3, 0xa4, 0xb1, 0xe8, 0x35, 0x2a, 0xe2, 0xe2, 0x06, 0x9b, 0x1c, 0x9c, 0x5a, 0x54, 0x96,
	0x5a, 0x24, 0x64, 0xc4, 0xc5, 0x11, 0x9c, 0x58, 0x09, 0x16, 0x11, 0x12, 0xd4, 0x83, 0xb9, 0x16,
	0xe6, 0x36, 0x29, 0x0c, 0xa1, 0x62, 0x25, 0x06, 0x21, 0x7d, 0x2e, 0xf6, 0xe0, 0xc4, 0x4a, 0x90,
	0x15, 0x42, 0x02, 0x70, 0x79, 0xa8, 0x83, 0xa4, 0xd0, 0x45, 0x8a, 0x95, 0x18, 0x92, 0xd8, 0xc0,
	0x02, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa5, 0xf7, 0x7b, 0x7f, 0x0e, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HelloServerClient is the client API for HelloServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloServerClient interface {
	// 第一个接口：打招呼的服务
	SayHello(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*HelloRes, error)
	// 第二个接口：说名字的服务
	SayName(ctx context.Context, in *NameReq, opts ...grpc.CallOption) (*NameRes, error)
}

type helloServerClient struct {
	cc *grpc.ClientConn
}

func NewHelloServerClient(cc *grpc.ClientConn) HelloServerClient {
	return &helloServerClient{cc}
}

func (c *helloServerClient) SayHello(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*HelloRes, error) {
	out := new(HelloRes)
	err := c.cc.Invoke(ctx, "/myproto.HelloServer/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloServerClient) SayName(ctx context.Context, in *NameReq, opts ...grpc.CallOption) (*NameRes, error) {
	out := new(NameRes)
	err := c.cc.Invoke(ctx, "/myproto.HelloServer/SayName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloServerServer is the server API for HelloServer service.
type HelloServerServer interface {
	// 第一个接口：打招呼的服务
	SayHello(context.Context, *HelloReq) (*HelloRes, error)
	// 第二个接口：说名字的服务
	SayName(context.Context, *NameReq) (*NameRes, error)
}

// UnimplementedHelloServerServer can be embedded to have forward compatible implementations.
type UnimplementedHelloServerServer struct {
}

func (*UnimplementedHelloServerServer) SayHello(ctx context.Context, req *HelloReq) (*HelloRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (*UnimplementedHelloServerServer) SayName(ctx context.Context, req *NameReq) (*NameRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayName not implemented")
}

func RegisterHelloServerServer(s *grpc.Server, srv HelloServerServer) {
	s.RegisterService(&_HelloServer_serviceDesc, srv)
}

func _HelloServer_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServerServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/myproto.HelloServer/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServerServer).SayHello(ctx, req.(*HelloReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloServer_SayName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NameReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServerServer).SayName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/myproto.HelloServer/SayName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServerServer).SayName(ctx, req.(*NameReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _HelloServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "myproto.HelloServer",
	HandlerType: (*HelloServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _HelloServer_SayHello_Handler,
		},
		{
			MethodName: "SayName",
			Handler:    _HelloServer_SayName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "myproto.proto",
}
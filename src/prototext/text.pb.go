// Code generated by protoc-gen-go. DO NOT EDIT.
// source: text.proto

package prototext

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// 熊猫 发送的消息
type PandaRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Height               int32    `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	Weight               []int32  `protobuf:"varint,3,rep,packed,name=weight,proto3" json:"weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PandaRequest) Reset()         { *m = PandaRequest{} }
func (m *PandaRequest) String() string { return proto.CompactTextString(m) }
func (*PandaRequest) ProtoMessage()    {}
func (*PandaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8e73d1ce47f9297, []int{0}
}

func (m *PandaRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PandaRequest.Unmarshal(m, b)
}
func (m *PandaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PandaRequest.Marshal(b, m, deterministic)
}
func (m *PandaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PandaRequest.Merge(m, src)
}
func (m *PandaRequest) XXX_Size() int {
	return xxx_messageInfo_PandaRequest.Size(m)
}
func (m *PandaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PandaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PandaRequest proto.InternalMessageInfo

func (m *PandaRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PandaRequest) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *PandaRequest) GetWeight() []int32 {
	if m != nil {
		return m.Weight
	}
	return nil
}

// 熊猫 返回的消息
type PandaResponse struct {
	Errno                int32    `protobuf:"varint,1,opt,name=errno,proto3" json:"errno,omitempty"`
	Errmessage           string   `protobuf:"bytes,2,opt,name=errmessage,proto3" json:"errmessage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PandaResponse) Reset()         { *m = PandaResponse{} }
func (m *PandaResponse) String() string { return proto.CompactTextString(m) }
func (*PandaResponse) ProtoMessage()    {}
func (*PandaResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8e73d1ce47f9297, []int{1}
}

func (m *PandaResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PandaResponse.Unmarshal(m, b)
}
func (m *PandaResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PandaResponse.Marshal(b, m, deterministic)
}
func (m *PandaResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PandaResponse.Merge(m, src)
}
func (m *PandaResponse) XXX_Size() int {
	return xxx_messageInfo_PandaResponse.Size(m)
}
func (m *PandaResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PandaResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PandaResponse proto.InternalMessageInfo

func (m *PandaResponse) GetErrno() int32 {
	if m != nil {
		return m.Errno
	}
	return 0
}

func (m *PandaResponse) GetErrmessage() string {
	if m != nil {
		return m.Errmessage
	}
	return ""
}

func init() {
	proto.RegisterType((*PandaRequest)(nil), "prototext.PandaRequest")
	proto.RegisterType((*PandaResponse)(nil), "prototext.PandaResponse")
}

func init() { proto.RegisterFile("text.proto", fileDescriptor_a8e73d1ce47f9297) }

var fileDescriptor_a8e73d1ce47f9297 = []byte{
	// 155 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8e, 0xb1, 0x0a, 0xc2, 0x30,
	0x10, 0x86, 0x89, 0x35, 0x85, 0x1c, 0xba, 0x1c, 0x22, 0x9d, 0x24, 0x74, 0xea, 0xe4, 0xe2, 0x33,
	0xb8, 0xcb, 0xbd, 0x41, 0xc4, 0xa3, 0x75, 0x68, 0x52, 0x93, 0x88, 0x3e, 0xbe, 0xe4, 0xd2, 0xc1,
	0xe9, 0xee, 0xfb, 0x86, 0x8f, 0x1f, 0x20, 0xf3, 0x37, 0x9f, 0x97, 0x18, 0x72, 0x40, 0x23, 0xa7,
	0x88, 0x9e, 0x60, 0x77, 0x73, 0xfe, 0xe1, 0x88, 0x5f, 0x6f, 0x4e, 0x19, 0x11, 0xb6, 0xde, 0xcd,
	0xdc, 0x29, 0xab, 0x06, 0x43, 0xf2, 0xe3, 0x11, 0xda, 0x89, 0x9f, 0xe3, 0x94, 0xbb, 0x8d, 0x55,
	0x83, 0xa6, 0x95, 0x8a, 0xff, 0x54, 0xdf, 0xd8, 0xa6, 0xf8, 0x4a, 0xfd, 0x15, 0xf6, 0x6b, 0x33,
	0x2d, 0xc1, 0x27, 0xc6, 0x03, 0x68, 0x8e, 0xd1, 0x07, 0xa9, 0x6a, 0xaa, 0x80, 0x27, 0x00, 0x8e,
	0x71, 0xe6, 0x94, 0xdc, 0xc8, 0x92, 0x36, 0xf4, 0x67, 0xee, 0xad, 0xac, 0xbc, 0xfc, 0x02, 0x00,
	0x00, 0xff, 0xff, 0xfa, 0x8c, 0x2d, 0x5c, 0xba, 0x00, 0x00, 0x00,
}
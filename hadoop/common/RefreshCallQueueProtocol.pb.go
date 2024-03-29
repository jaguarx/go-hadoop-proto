// Code generated by protoc-gen-go. DO NOT EDIT.
// source: RefreshCallQueueProtocol.proto

package hadoop_common

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

//*
//  Refresh callqueue request.
type RefreshCallQueueRequestProto struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RefreshCallQueueRequestProto) Reset()         { *m = RefreshCallQueueRequestProto{} }
func (m *RefreshCallQueueRequestProto) String() string { return proto.CompactTextString(m) }
func (*RefreshCallQueueRequestProto) ProtoMessage()    {}
func (*RefreshCallQueueRequestProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_bcdf9c0f468ba6fc, []int{0}
}

func (m *RefreshCallQueueRequestProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RefreshCallQueueRequestProto.Unmarshal(m, b)
}
func (m *RefreshCallQueueRequestProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RefreshCallQueueRequestProto.Marshal(b, m, deterministic)
}
func (m *RefreshCallQueueRequestProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RefreshCallQueueRequestProto.Merge(m, src)
}
func (m *RefreshCallQueueRequestProto) XXX_Size() int {
	return xxx_messageInfo_RefreshCallQueueRequestProto.Size(m)
}
func (m *RefreshCallQueueRequestProto) XXX_DiscardUnknown() {
	xxx_messageInfo_RefreshCallQueueRequestProto.DiscardUnknown(m)
}

var xxx_messageInfo_RefreshCallQueueRequestProto proto.InternalMessageInfo

//*
// void response.
type RefreshCallQueueResponseProto struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RefreshCallQueueResponseProto) Reset()         { *m = RefreshCallQueueResponseProto{} }
func (m *RefreshCallQueueResponseProto) String() string { return proto.CompactTextString(m) }
func (*RefreshCallQueueResponseProto) ProtoMessage()    {}
func (*RefreshCallQueueResponseProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_bcdf9c0f468ba6fc, []int{1}
}

func (m *RefreshCallQueueResponseProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RefreshCallQueueResponseProto.Unmarshal(m, b)
}
func (m *RefreshCallQueueResponseProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RefreshCallQueueResponseProto.Marshal(b, m, deterministic)
}
func (m *RefreshCallQueueResponseProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RefreshCallQueueResponseProto.Merge(m, src)
}
func (m *RefreshCallQueueResponseProto) XXX_Size() int {
	return xxx_messageInfo_RefreshCallQueueResponseProto.Size(m)
}
func (m *RefreshCallQueueResponseProto) XXX_DiscardUnknown() {
	xxx_messageInfo_RefreshCallQueueResponseProto.DiscardUnknown(m)
}

var xxx_messageInfo_RefreshCallQueueResponseProto proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RefreshCallQueueRequestProto)(nil), "hadoop.common.RefreshCallQueueRequestProto")
	proto.RegisterType((*RefreshCallQueueResponseProto)(nil), "hadoop.common.RefreshCallQueueResponseProto")
}

func init() { proto.RegisterFile("RefreshCallQueueProtocol.proto", fileDescriptor_bcdf9c0f468ba6fc) }

var fileDescriptor_bcdf9c0f468ba6fc = []byte{
	// 168 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x0b, 0x4a, 0x4d, 0x2b,
	0x4a, 0x2d, 0xce, 0x70, 0x4e, 0xcc, 0xc9, 0x09, 0x2c, 0x4d, 0x2d, 0x4d, 0x0d, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0xce, 0xcf, 0xd1, 0x2b, 0x00, 0x31, 0x84, 0x78, 0x33, 0x12, 0x53, 0xf2, 0xf3, 0x0b,
	0xf4, 0x92, 0xf3, 0x73, 0x73, 0xf3, 0xf3, 0x94, 0xe4, 0xb8, 0x64, 0xd0, 0x35, 0x04, 0xa5, 0x16,
	0x96, 0xa6, 0x16, 0x97, 0x80, 0xf5, 0x29, 0xc9, 0x73, 0xc9, 0x62, 0xca, 0x17, 0x17, 0xe4, 0xe7,
	0x15, 0x43, 0x0c, 0x36, 0x9a, 0xc0, 0xc8, 0x25, 0x8f, 0xcb, 0xca, 0xe0, 0xd4, 0xa2, 0xb2, 0xcc,
	0xe4, 0x54, 0xa1, 0x5c, 0x2e, 0x81, 0x22, 0x34, 0x25, 0x42, 0xda, 0x7a, 0x28, 0x0e, 0xd1, 0xc3,
	0xe7, 0x0a, 0x29, 0x1d, 0x82, 0x8a, 0x91, 0x9c, 0xe4, 0xe4, 0xcc, 0x25, 0x9d, 0x5f, 0x94, 0xae,
	0x97, 0x58, 0x90, 0x98, 0x9c, 0x91, 0x0a, 0xd3, 0x99, 0x59, 0x90, 0x0c, 0x09, 0x01, 0x27, 0x9c,
	0x21, 0x04, 0xa6, 0x8b, 0x3b, 0x18, 0x19, 0x17, 0x30, 0x32, 0x02, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x5d, 0x22, 0xcd, 0xe0, 0x48, 0x01, 0x00, 0x00,
}

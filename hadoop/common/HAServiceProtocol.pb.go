// Code generated by protoc-gen-go. DO NOT EDIT.
// source: HAServiceProtocol.proto

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

type HAServiceStateProto int32

const (
	HAServiceStateProto_INITIALIZING HAServiceStateProto = 0
	HAServiceStateProto_ACTIVE       HAServiceStateProto = 1
	HAServiceStateProto_STANDBY      HAServiceStateProto = 2
	HAServiceStateProto_OBSERVER     HAServiceStateProto = 3
)

var HAServiceStateProto_name = map[int32]string{
	0: "INITIALIZING",
	1: "ACTIVE",
	2: "STANDBY",
	3: "OBSERVER",
}

var HAServiceStateProto_value = map[string]int32{
	"INITIALIZING": 0,
	"ACTIVE":       1,
	"STANDBY":      2,
	"OBSERVER":     3,
}

func (x HAServiceStateProto) Enum() *HAServiceStateProto {
	p := new(HAServiceStateProto)
	*p = x
	return p
}

func (x HAServiceStateProto) String() string {
	return proto.EnumName(HAServiceStateProto_name, int32(x))
}

func (x *HAServiceStateProto) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(HAServiceStateProto_value, data, "HAServiceStateProto")
	if err != nil {
		return err
	}
	*x = HAServiceStateProto(value)
	return nil
}

func (HAServiceStateProto) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3fb056deb41c1807, []int{0}
}

type HARequestSource int32

const (
	HARequestSource_REQUEST_BY_USER        HARequestSource = 0
	HARequestSource_REQUEST_BY_USER_FORCED HARequestSource = 1
	HARequestSource_REQUEST_BY_ZKFC        HARequestSource = 2
)

var HARequestSource_name = map[int32]string{
	0: "REQUEST_BY_USER",
	1: "REQUEST_BY_USER_FORCED",
	2: "REQUEST_BY_ZKFC",
}

var HARequestSource_value = map[string]int32{
	"REQUEST_BY_USER":        0,
	"REQUEST_BY_USER_FORCED": 1,
	"REQUEST_BY_ZKFC":        2,
}

func (x HARequestSource) Enum() *HARequestSource {
	p := new(HARequestSource)
	*p = x
	return p
}

func (x HARequestSource) String() string {
	return proto.EnumName(HARequestSource_name, int32(x))
}

func (x *HARequestSource) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(HARequestSource_value, data, "HARequestSource")
	if err != nil {
		return err
	}
	*x = HARequestSource(value)
	return nil
}

func (HARequestSource) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3fb056deb41c1807, []int{1}
}

type HAStateChangeRequestInfoProto struct {
	ReqSource            *HARequestSource `protobuf:"varint,1,req,name=reqSource,enum=hadoop.common.HARequestSource" json:"reqSource,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *HAStateChangeRequestInfoProto) Reset()         { *m = HAStateChangeRequestInfoProto{} }
func (m *HAStateChangeRequestInfoProto) String() string { return proto.CompactTextString(m) }
func (*HAStateChangeRequestInfoProto) ProtoMessage()    {}
func (*HAStateChangeRequestInfoProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_3fb056deb41c1807, []int{0}
}

func (m *HAStateChangeRequestInfoProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HAStateChangeRequestInfoProto.Unmarshal(m, b)
}
func (m *HAStateChangeRequestInfoProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HAStateChangeRequestInfoProto.Marshal(b, m, deterministic)
}
func (m *HAStateChangeRequestInfoProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HAStateChangeRequestInfoProto.Merge(m, src)
}
func (m *HAStateChangeRequestInfoProto) XXX_Size() int {
	return xxx_messageInfo_HAStateChangeRequestInfoProto.Size(m)
}
func (m *HAStateChangeRequestInfoProto) XXX_DiscardUnknown() {
	xxx_messageInfo_HAStateChangeRequestInfoProto.DiscardUnknown(m)
}

var xxx_messageInfo_HAStateChangeRequestInfoProto proto.InternalMessageInfo

func (m *HAStateChangeRequestInfoProto) GetReqSource() HARequestSource {
	if m != nil && m.ReqSource != nil {
		return *m.ReqSource
	}
	return HARequestSource_REQUEST_BY_USER
}

//*
// void request
type MonitorHealthRequestProto struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MonitorHealthRequestProto) Reset()         { *m = MonitorHealthRequestProto{} }
func (m *MonitorHealthRequestProto) String() string { return proto.CompactTextString(m) }
func (*MonitorHealthRequestProto) ProtoMessage()    {}
func (*MonitorHealthRequestProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_3fb056deb41c1807, []int{1}
}

func (m *MonitorHealthRequestProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MonitorHealthRequestProto.Unmarshal(m, b)
}
func (m *MonitorHealthRequestProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MonitorHealthRequestProto.Marshal(b, m, deterministic)
}
func (m *MonitorHealthRequestProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MonitorHealthRequestProto.Merge(m, src)
}
func (m *MonitorHealthRequestProto) XXX_Size() int {
	return xxx_messageInfo_MonitorHealthRequestProto.Size(m)
}
func (m *MonitorHealthRequestProto) XXX_DiscardUnknown() {
	xxx_messageInfo_MonitorHealthRequestProto.DiscardUnknown(m)
}

var xxx_messageInfo_MonitorHealthRequestProto proto.InternalMessageInfo

//*
// void response
type MonitorHealthResponseProto struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MonitorHealthResponseProto) Reset()         { *m = MonitorHealthResponseProto{} }
func (m *MonitorHealthResponseProto) String() string { return proto.CompactTextString(m) }
func (*MonitorHealthResponseProto) ProtoMessage()    {}
func (*MonitorHealthResponseProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_3fb056deb41c1807, []int{2}
}

func (m *MonitorHealthResponseProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MonitorHealthResponseProto.Unmarshal(m, b)
}
func (m *MonitorHealthResponseProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MonitorHealthResponseProto.Marshal(b, m, deterministic)
}
func (m *MonitorHealthResponseProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MonitorHealthResponseProto.Merge(m, src)
}
func (m *MonitorHealthResponseProto) XXX_Size() int {
	return xxx_messageInfo_MonitorHealthResponseProto.Size(m)
}
func (m *MonitorHealthResponseProto) XXX_DiscardUnknown() {
	xxx_messageInfo_MonitorHealthResponseProto.DiscardUnknown(m)
}

var xxx_messageInfo_MonitorHealthResponseProto proto.InternalMessageInfo

//*
// void request
type TransitionToActiveRequestProto struct {
	ReqInfo              *HAStateChangeRequestInfoProto `protobuf:"bytes,1,req,name=reqInfo" json:"reqInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *TransitionToActiveRequestProto) Reset()         { *m = TransitionToActiveRequestProto{} }
func (m *TransitionToActiveRequestProto) String() string { return proto.CompactTextString(m) }
func (*TransitionToActiveRequestProto) ProtoMessage()    {}
func (*TransitionToActiveRequestProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_3fb056deb41c1807, []int{3}
}

func (m *TransitionToActiveRequestProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransitionToActiveRequestProto.Unmarshal(m, b)
}
func (m *TransitionToActiveRequestProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransitionToActiveRequestProto.Marshal(b, m, deterministic)
}
func (m *TransitionToActiveRequestProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransitionToActiveRequestProto.Merge(m, src)
}
func (m *TransitionToActiveRequestProto) XXX_Size() int {
	return xxx_messageInfo_TransitionToActiveRequestProto.Size(m)
}
func (m *TransitionToActiveRequestProto) XXX_DiscardUnknown() {
	xxx_messageInfo_TransitionToActiveRequestProto.DiscardUnknown(m)
}

var xxx_messageInfo_TransitionToActiveRequestProto proto.InternalMessageInfo

func (m *TransitionToActiveRequestProto) GetReqInfo() *HAStateChangeRequestInfoProto {
	if m != nil {
		return m.ReqInfo
	}
	return nil
}

//*
// void response
type TransitionToActiveResponseProto struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransitionToActiveResponseProto) Reset()         { *m = TransitionToActiveResponseProto{} }
func (m *TransitionToActiveResponseProto) String() string { return proto.CompactTextString(m) }
func (*TransitionToActiveResponseProto) ProtoMessage()    {}
func (*TransitionToActiveResponseProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_3fb056deb41c1807, []int{4}
}

func (m *TransitionToActiveResponseProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransitionToActiveResponseProto.Unmarshal(m, b)
}
func (m *TransitionToActiveResponseProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransitionToActiveResponseProto.Marshal(b, m, deterministic)
}
func (m *TransitionToActiveResponseProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransitionToActiveResponseProto.Merge(m, src)
}
func (m *TransitionToActiveResponseProto) XXX_Size() int {
	return xxx_messageInfo_TransitionToActiveResponseProto.Size(m)
}
func (m *TransitionToActiveResponseProto) XXX_DiscardUnknown() {
	xxx_messageInfo_TransitionToActiveResponseProto.DiscardUnknown(m)
}

var xxx_messageInfo_TransitionToActiveResponseProto proto.InternalMessageInfo

//*
// void request
type TransitionToStandbyRequestProto struct {
	ReqInfo              *HAStateChangeRequestInfoProto `protobuf:"bytes,1,req,name=reqInfo" json:"reqInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *TransitionToStandbyRequestProto) Reset()         { *m = TransitionToStandbyRequestProto{} }
func (m *TransitionToStandbyRequestProto) String() string { return proto.CompactTextString(m) }
func (*TransitionToStandbyRequestProto) ProtoMessage()    {}
func (*TransitionToStandbyRequestProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_3fb056deb41c1807, []int{5}
}

func (m *TransitionToStandbyRequestProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransitionToStandbyRequestProto.Unmarshal(m, b)
}
func (m *TransitionToStandbyRequestProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransitionToStandbyRequestProto.Marshal(b, m, deterministic)
}
func (m *TransitionToStandbyRequestProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransitionToStandbyRequestProto.Merge(m, src)
}
func (m *TransitionToStandbyRequestProto) XXX_Size() int {
	return xxx_messageInfo_TransitionToStandbyRequestProto.Size(m)
}
func (m *TransitionToStandbyRequestProto) XXX_DiscardUnknown() {
	xxx_messageInfo_TransitionToStandbyRequestProto.DiscardUnknown(m)
}

var xxx_messageInfo_TransitionToStandbyRequestProto proto.InternalMessageInfo

func (m *TransitionToStandbyRequestProto) GetReqInfo() *HAStateChangeRequestInfoProto {
	if m != nil {
		return m.ReqInfo
	}
	return nil
}

//*
// void response
type TransitionToStandbyResponseProto struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransitionToStandbyResponseProto) Reset()         { *m = TransitionToStandbyResponseProto{} }
func (m *TransitionToStandbyResponseProto) String() string { return proto.CompactTextString(m) }
func (*TransitionToStandbyResponseProto) ProtoMessage()    {}
func (*TransitionToStandbyResponseProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_3fb056deb41c1807, []int{6}
}

func (m *TransitionToStandbyResponseProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransitionToStandbyResponseProto.Unmarshal(m, b)
}
func (m *TransitionToStandbyResponseProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransitionToStandbyResponseProto.Marshal(b, m, deterministic)
}
func (m *TransitionToStandbyResponseProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransitionToStandbyResponseProto.Merge(m, src)
}
func (m *TransitionToStandbyResponseProto) XXX_Size() int {
	return xxx_messageInfo_TransitionToStandbyResponseProto.Size(m)
}
func (m *TransitionToStandbyResponseProto) XXX_DiscardUnknown() {
	xxx_messageInfo_TransitionToStandbyResponseProto.DiscardUnknown(m)
}

var xxx_messageInfo_TransitionToStandbyResponseProto proto.InternalMessageInfo

//*
// void request
type TransitionToObserverRequestProto struct {
	ReqInfo              *HAStateChangeRequestInfoProto `protobuf:"bytes,1,req,name=reqInfo" json:"reqInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *TransitionToObserverRequestProto) Reset()         { *m = TransitionToObserverRequestProto{} }
func (m *TransitionToObserverRequestProto) String() string { return proto.CompactTextString(m) }
func (*TransitionToObserverRequestProto) ProtoMessage()    {}
func (*TransitionToObserverRequestProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_3fb056deb41c1807, []int{7}
}

func (m *TransitionToObserverRequestProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransitionToObserverRequestProto.Unmarshal(m, b)
}
func (m *TransitionToObserverRequestProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransitionToObserverRequestProto.Marshal(b, m, deterministic)
}
func (m *TransitionToObserverRequestProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransitionToObserverRequestProto.Merge(m, src)
}
func (m *TransitionToObserverRequestProto) XXX_Size() int {
	return xxx_messageInfo_TransitionToObserverRequestProto.Size(m)
}
func (m *TransitionToObserverRequestProto) XXX_DiscardUnknown() {
	xxx_messageInfo_TransitionToObserverRequestProto.DiscardUnknown(m)
}

var xxx_messageInfo_TransitionToObserverRequestProto proto.InternalMessageInfo

func (m *TransitionToObserverRequestProto) GetReqInfo() *HAStateChangeRequestInfoProto {
	if m != nil {
		return m.ReqInfo
	}
	return nil
}

//*
// void response
type TransitionToObserverResponseProto struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransitionToObserverResponseProto) Reset()         { *m = TransitionToObserverResponseProto{} }
func (m *TransitionToObserverResponseProto) String() string { return proto.CompactTextString(m) }
func (*TransitionToObserverResponseProto) ProtoMessage()    {}
func (*TransitionToObserverResponseProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_3fb056deb41c1807, []int{8}
}

func (m *TransitionToObserverResponseProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransitionToObserverResponseProto.Unmarshal(m, b)
}
func (m *TransitionToObserverResponseProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransitionToObserverResponseProto.Marshal(b, m, deterministic)
}
func (m *TransitionToObserverResponseProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransitionToObserverResponseProto.Merge(m, src)
}
func (m *TransitionToObserverResponseProto) XXX_Size() int {
	return xxx_messageInfo_TransitionToObserverResponseProto.Size(m)
}
func (m *TransitionToObserverResponseProto) XXX_DiscardUnknown() {
	xxx_messageInfo_TransitionToObserverResponseProto.DiscardUnknown(m)
}

var xxx_messageInfo_TransitionToObserverResponseProto proto.InternalMessageInfo

//*
// void request
type GetServiceStatusRequestProto struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetServiceStatusRequestProto) Reset()         { *m = GetServiceStatusRequestProto{} }
func (m *GetServiceStatusRequestProto) String() string { return proto.CompactTextString(m) }
func (*GetServiceStatusRequestProto) ProtoMessage()    {}
func (*GetServiceStatusRequestProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_3fb056deb41c1807, []int{9}
}

func (m *GetServiceStatusRequestProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetServiceStatusRequestProto.Unmarshal(m, b)
}
func (m *GetServiceStatusRequestProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetServiceStatusRequestProto.Marshal(b, m, deterministic)
}
func (m *GetServiceStatusRequestProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetServiceStatusRequestProto.Merge(m, src)
}
func (m *GetServiceStatusRequestProto) XXX_Size() int {
	return xxx_messageInfo_GetServiceStatusRequestProto.Size(m)
}
func (m *GetServiceStatusRequestProto) XXX_DiscardUnknown() {
	xxx_messageInfo_GetServiceStatusRequestProto.DiscardUnknown(m)
}

var xxx_messageInfo_GetServiceStatusRequestProto proto.InternalMessageInfo

//*
// Returns the state of the service
type GetServiceStatusResponseProto struct {
	State *HAServiceStateProto `protobuf:"varint,1,req,name=state,enum=hadoop.common.HAServiceStateProto" json:"state,omitempty"`
	// If state is STANDBY, indicate whether it is
	// ready to become active.
	ReadyToBecomeActive *bool `protobuf:"varint,2,opt,name=readyToBecomeActive" json:"readyToBecomeActive,omitempty"`
	// If not ready to become active, a textual explanation of why not
	NotReadyReason       *string  `protobuf:"bytes,3,opt,name=notReadyReason" json:"notReadyReason,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetServiceStatusResponseProto) Reset()         { *m = GetServiceStatusResponseProto{} }
func (m *GetServiceStatusResponseProto) String() string { return proto.CompactTextString(m) }
func (*GetServiceStatusResponseProto) ProtoMessage()    {}
func (*GetServiceStatusResponseProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_3fb056deb41c1807, []int{10}
}

func (m *GetServiceStatusResponseProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetServiceStatusResponseProto.Unmarshal(m, b)
}
func (m *GetServiceStatusResponseProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetServiceStatusResponseProto.Marshal(b, m, deterministic)
}
func (m *GetServiceStatusResponseProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetServiceStatusResponseProto.Merge(m, src)
}
func (m *GetServiceStatusResponseProto) XXX_Size() int {
	return xxx_messageInfo_GetServiceStatusResponseProto.Size(m)
}
func (m *GetServiceStatusResponseProto) XXX_DiscardUnknown() {
	xxx_messageInfo_GetServiceStatusResponseProto.DiscardUnknown(m)
}

var xxx_messageInfo_GetServiceStatusResponseProto proto.InternalMessageInfo

func (m *GetServiceStatusResponseProto) GetState() HAServiceStateProto {
	if m != nil && m.State != nil {
		return *m.State
	}
	return HAServiceStateProto_INITIALIZING
}

func (m *GetServiceStatusResponseProto) GetReadyToBecomeActive() bool {
	if m != nil && m.ReadyToBecomeActive != nil {
		return *m.ReadyToBecomeActive
	}
	return false
}

func (m *GetServiceStatusResponseProto) GetNotReadyReason() string {
	if m != nil && m.NotReadyReason != nil {
		return *m.NotReadyReason
	}
	return ""
}

func init() {
	proto.RegisterEnum("hadoop.common.HAServiceStateProto", HAServiceStateProto_name, HAServiceStateProto_value)
	proto.RegisterEnum("hadoop.common.HARequestSource", HARequestSource_name, HARequestSource_value)
	proto.RegisterType((*HAStateChangeRequestInfoProto)(nil), "hadoop.common.HAStateChangeRequestInfoProto")
	proto.RegisterType((*MonitorHealthRequestProto)(nil), "hadoop.common.MonitorHealthRequestProto")
	proto.RegisterType((*MonitorHealthResponseProto)(nil), "hadoop.common.MonitorHealthResponseProto")
	proto.RegisterType((*TransitionToActiveRequestProto)(nil), "hadoop.common.TransitionToActiveRequestProto")
	proto.RegisterType((*TransitionToActiveResponseProto)(nil), "hadoop.common.TransitionToActiveResponseProto")
	proto.RegisterType((*TransitionToStandbyRequestProto)(nil), "hadoop.common.TransitionToStandbyRequestProto")
	proto.RegisterType((*TransitionToStandbyResponseProto)(nil), "hadoop.common.TransitionToStandbyResponseProto")
	proto.RegisterType((*TransitionToObserverRequestProto)(nil), "hadoop.common.TransitionToObserverRequestProto")
	proto.RegisterType((*TransitionToObserverResponseProto)(nil), "hadoop.common.TransitionToObserverResponseProto")
	proto.RegisterType((*GetServiceStatusRequestProto)(nil), "hadoop.common.GetServiceStatusRequestProto")
	proto.RegisterType((*GetServiceStatusResponseProto)(nil), "hadoop.common.GetServiceStatusResponseProto")
}

func init() { proto.RegisterFile("HAServiceProtocol.proto", fileDescriptor_3fb056deb41c1807) }

var fileDescriptor_3fb056deb41c1807 = []byte{
	// 574 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0x6d, 0x6f, 0xd2, 0x40,
	0x1c, 0x5f, 0x99, 0xba, 0xed, 0xbf, 0xa7, 0xe6, 0x30, 0x8a, 0x75, 0x43, 0x56, 0x13, 0x83, 0x73,
	0xd6, 0x85, 0x57, 0x26, 0xfa, 0xa6, 0x65, 0x65, 0x34, 0x2a, 0xe8, 0xb5, 0x9b, 0xd9, 0x12, 0x43,
	0xba, 0x72, 0xd2, 0x9a, 0xd1, 0x83, 0xeb, 0x41, 0xc2, 0x37, 0xf0, 0x63, 0xf8, 0x1d, 0xfc, 0x30,
	0x7e, 0x1d, 0x03, 0x65, 0xd9, 0xd1, 0xd6, 0xca, 0x9b, 0xbd, 0x82, 0xde, 0xfd, 0x9e, 0xda, 0xff,
	0xfd, 0x72, 0xf0, 0xb8, 0xa9, 0xdb, 0x84, 0x8d, 0x03, 0x8f, 0x7c, 0x66, 0x94, 0x53, 0x8f, 0x5e,
	0x6b, 0x83, 0xe9, 0x1f, 0xb4, 0xed, 0xbb, 0x5d, 0x4a, 0x07, 0x9a, 0x47, 0xfb, 0x7d, 0x1a, 0xaa,
	0xdf, 0x60, 0xbf, 0xa9, 0xdb, 0xdc, 0xe5, 0xa4, 0xee, 0xbb, 0x61, 0x8f, 0x60, 0x32, 0x1c, 0x91,
	0x88, 0x5b, 0xe1, 0x77, 0x3a, 0x23, 0xa2, 0xf7, 0xb0, 0xc1, 0xc8, 0xd0, 0xa6, 0x23, 0xe6, 0x91,
	0x92, 0x54, 0x29, 0x54, 0x77, 0x6a, 0x65, 0x6d, 0x41, 0x43, 0x6b, 0xea, 0x73, 0x56, 0x8c, 0xc2,
	0xb7, 0x04, 0xf5, 0x29, 0x3c, 0xf9, 0x44, 0xc3, 0x80, 0x53, 0xd6, 0x24, 0xee, 0x35, 0xf7, 0xe7,
	0xc0, 0x99, 0xb4, 0xba, 0x07, 0x4a, 0x62, 0x33, 0x1a, 0xd0, 0x30, 0x8a, 0x13, 0xab, 0x3e, 0x94,
	0x1d, 0xe6, 0x86, 0x51, 0xc0, 0x03, 0x1a, 0x3a, 0x54, 0xf7, 0x78, 0x30, 0x26, 0x22, 0x1f, 0x35,
	0x60, 0x8d, 0x91, 0xe1, 0x34, 0xea, 0x2c, 0xd8, 0x66, 0xed, 0x28, 0x15, 0x2c, 0xe7, 0xcd, 0xf0,
	0x0d, 0x59, 0x3d, 0x80, 0x67, 0x59, 0x4e, 0x62, 0x98, 0x60, 0x11, 0x62, 0x73, 0x37, 0xec, 0x5e,
	0x4d, 0xee, 0x24, 0x8d, 0x0a, 0x95, 0x4c, 0x2b, 0x31, 0xce, 0x8f, 0x45, 0x4c, 0xfb, 0x2a, 0x22,
	0x6c, 0x4c, 0xd8, 0x9d, 0xe4, 0x79, 0x0e, 0x07, 0xd9, 0x5e, 0x62, 0xa0, 0x32, 0xec, 0x9d, 0x12,
	0x3e, 0x3f, 0x71, 0x53, 0xd9, 0x51, 0xb4, 0x30, 0xea, 0xdf, 0x12, 0xec, 0xa7, 0x01, 0x82, 0x02,
	0x7a, 0x0b, 0xf7, 0xa3, 0x69, 0x9c, 0xf9, 0x19, 0x53, 0xd3, 0x61, 0x6f, 0xb9, 0x31, 0x05, 0xc7,
	0x04, 0x74, 0x0c, 0x45, 0x46, 0xdc, 0xee, 0xc4, 0xa1, 0x06, 0xf1, 0x68, 0x9f, 0xc4, 0xf3, 0x2b,
	0x15, 0x2a, 0x52, 0x75, 0x1d, 0x67, 0x6d, 0xa1, 0x17, 0xb0, 0x13, 0x52, 0x8e, 0xa7, 0x3b, 0x98,
	0xb8, 0x11, 0x0d, 0x4b, 0xab, 0x15, 0xa9, 0xba, 0x81, 0x13, 0xab, 0x87, 0x2d, 0x28, 0x66, 0xf8,
	0x22, 0x19, 0xb6, 0xac, 0x96, 0xe5, 0x58, 0xfa, 0x47, 0xeb, 0xd2, 0x6a, 0x9d, 0xca, 0x2b, 0x08,
	0xe0, 0x81, 0x5e, 0x77, 0xac, 0x73, 0x53, 0x96, 0xd0, 0x26, 0xac, 0xd9, 0x8e, 0xde, 0x3a, 0x31,
	0x2e, 0xe4, 0x02, 0xda, 0x82, 0xf5, 0xb6, 0x61, 0x9b, 0xf8, 0xdc, 0xc4, 0xf2, 0xea, 0xe1, 0x57,
	0xd8, 0x4d, 0x74, 0x05, 0x15, 0x61, 0x17, 0x9b, 0x5f, 0xce, 0x4c, 0xdb, 0xe9, 0x18, 0x17, 0x9d,
	0x33, 0xdb, 0xc4, 0xf2, 0x0a, 0x52, 0xe0, 0x51, 0x62, 0xb1, 0xd3, 0x68, 0xe3, 0xba, 0x79, 0x22,
	0x4b, 0x09, 0xc2, 0xe5, 0x87, 0x46, 0x5d, 0x2e, 0xd4, 0xfe, 0xdc, 0x83, 0x52, 0xaa, 0xf0, 0xf3,
	0x47, 0xd4, 0x85, 0xed, 0xbe, 0x58, 0x33, 0x54, 0x4d, 0x7c, 0xdb, 0x7f, 0x36, 0x54, 0x79, 0x99,
	0x8f, 0x14, 0xe7, 0x17, 0x01, 0xe2, 0xa9, 0x12, 0xa1, 0xd7, 0x09, 0x81, 0xfc, 0x46, 0x2b, 0xda,
	0x12, 0x70, 0xd1, 0x74, 0x0c, 0x45, 0x9e, 0xee, 0x0a, 0xca, 0x93, 0xc9, 0xa8, 0xae, 0xf2, 0x66,
	0x19, 0xbc, 0xe8, 0x3b, 0x81, 0x87, 0x3c, 0xa3, 0x13, 0x28, 0x4f, 0x28, 0xab, 0xa4, 0xca, 0xf1,
	0x52, 0x04, 0xd1, 0xba, 0x0f, 0x72, 0x2f, 0x51, 0x24, 0xf4, 0x2a, 0xa1, 0x92, 0x57, 0x45, 0xe5,
	0xe8, 0xbf, 0x60, 0xc1, 0xce, 0x78, 0x07, 0x0a, 0x65, 0x3d, 0xcd, 0x1d, 0xb8, 0x9e, 0x4f, 0x6e,
	0x98, 0xbe, 0x1b, 0x5f, 0x26, 0x46, 0xfa, 0x96, 0x99, 0xfd, 0x46, 0x3f, 0x25, 0xe9, 0x97, 0x24,
	0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x30, 0x57, 0xef, 0x05, 0x85, 0x06, 0x00, 0x00,
}

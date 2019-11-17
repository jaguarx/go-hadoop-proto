// Code generated by protoc-gen-go. DO NOT EDIT.
// source: acl.proto

package hadoop_hdfs

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

type AclEntryProto_AclEntryScopeProto int32

const (
	AclEntryProto_ACCESS  AclEntryProto_AclEntryScopeProto = 0
	AclEntryProto_DEFAULT AclEntryProto_AclEntryScopeProto = 1
)

var AclEntryProto_AclEntryScopeProto_name = map[int32]string{
	0: "ACCESS",
	1: "DEFAULT",
}

var AclEntryProto_AclEntryScopeProto_value = map[string]int32{
	"ACCESS":  0,
	"DEFAULT": 1,
}

func (x AclEntryProto_AclEntryScopeProto) Enum() *AclEntryProto_AclEntryScopeProto {
	p := new(AclEntryProto_AclEntryScopeProto)
	*p = x
	return p
}

func (x AclEntryProto_AclEntryScopeProto) String() string {
	return proto.EnumName(AclEntryProto_AclEntryScopeProto_name, int32(x))
}

func (x *AclEntryProto_AclEntryScopeProto) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(AclEntryProto_AclEntryScopeProto_value, data, "AclEntryProto_AclEntryScopeProto")
	if err != nil {
		return err
	}
	*x = AclEntryProto_AclEntryScopeProto(value)
	return nil
}

func (AclEntryProto_AclEntryScopeProto) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a452f070aeef01eb, []int{1, 0}
}

type AclEntryProto_AclEntryTypeProto int32

const (
	AclEntryProto_USER  AclEntryProto_AclEntryTypeProto = 0
	AclEntryProto_GROUP AclEntryProto_AclEntryTypeProto = 1
	AclEntryProto_MASK  AclEntryProto_AclEntryTypeProto = 2
	AclEntryProto_OTHER AclEntryProto_AclEntryTypeProto = 3
)

var AclEntryProto_AclEntryTypeProto_name = map[int32]string{
	0: "USER",
	1: "GROUP",
	2: "MASK",
	3: "OTHER",
}

var AclEntryProto_AclEntryTypeProto_value = map[string]int32{
	"USER":  0,
	"GROUP": 1,
	"MASK":  2,
	"OTHER": 3,
}

func (x AclEntryProto_AclEntryTypeProto) Enum() *AclEntryProto_AclEntryTypeProto {
	p := new(AclEntryProto_AclEntryTypeProto)
	*p = x
	return p
}

func (x AclEntryProto_AclEntryTypeProto) String() string {
	return proto.EnumName(AclEntryProto_AclEntryTypeProto_name, int32(x))
}

func (x *AclEntryProto_AclEntryTypeProto) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(AclEntryProto_AclEntryTypeProto_value, data, "AclEntryProto_AclEntryTypeProto")
	if err != nil {
		return err
	}
	*x = AclEntryProto_AclEntryTypeProto(value)
	return nil
}

func (AclEntryProto_AclEntryTypeProto) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a452f070aeef01eb, []int{1, 1}
}

type AclEntryProto_FsActionProto int32

const (
	AclEntryProto_NONE          AclEntryProto_FsActionProto = 0
	AclEntryProto_EXECUTE       AclEntryProto_FsActionProto = 1
	AclEntryProto_WRITE         AclEntryProto_FsActionProto = 2
	AclEntryProto_WRITE_EXECUTE AclEntryProto_FsActionProto = 3
	AclEntryProto_READ          AclEntryProto_FsActionProto = 4
	AclEntryProto_READ_EXECUTE  AclEntryProto_FsActionProto = 5
	AclEntryProto_READ_WRITE    AclEntryProto_FsActionProto = 6
	AclEntryProto_PERM_ALL      AclEntryProto_FsActionProto = 7
)

var AclEntryProto_FsActionProto_name = map[int32]string{
	0: "NONE",
	1: "EXECUTE",
	2: "WRITE",
	3: "WRITE_EXECUTE",
	4: "READ",
	5: "READ_EXECUTE",
	6: "READ_WRITE",
	7: "PERM_ALL",
}

var AclEntryProto_FsActionProto_value = map[string]int32{
	"NONE":          0,
	"EXECUTE":       1,
	"WRITE":         2,
	"WRITE_EXECUTE": 3,
	"READ":          4,
	"READ_EXECUTE":  5,
	"READ_WRITE":    6,
	"PERM_ALL":      7,
}

func (x AclEntryProto_FsActionProto) Enum() *AclEntryProto_FsActionProto {
	p := new(AclEntryProto_FsActionProto)
	*p = x
	return p
}

func (x AclEntryProto_FsActionProto) String() string {
	return proto.EnumName(AclEntryProto_FsActionProto_name, int32(x))
}

func (x *AclEntryProto_FsActionProto) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(AclEntryProto_FsActionProto_value, data, "AclEntryProto_FsActionProto")
	if err != nil {
		return err
	}
	*x = AclEntryProto_FsActionProto(value)
	return nil
}

func (AclEntryProto_FsActionProto) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a452f070aeef01eb, []int{1, 2}
}

//*
// File or Directory permision - same spec as posix
type FsPermissionProto struct {
	Perm                 *uint32  `protobuf:"varint,1,req,name=perm" json:"perm,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FsPermissionProto) Reset()         { *m = FsPermissionProto{} }
func (m *FsPermissionProto) String() string { return proto.CompactTextString(m) }
func (*FsPermissionProto) ProtoMessage()    {}
func (*FsPermissionProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_a452f070aeef01eb, []int{0}
}

func (m *FsPermissionProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FsPermissionProto.Unmarshal(m, b)
}
func (m *FsPermissionProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FsPermissionProto.Marshal(b, m, deterministic)
}
func (m *FsPermissionProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FsPermissionProto.Merge(m, src)
}
func (m *FsPermissionProto) XXX_Size() int {
	return xxx_messageInfo_FsPermissionProto.Size(m)
}
func (m *FsPermissionProto) XXX_DiscardUnknown() {
	xxx_messageInfo_FsPermissionProto.DiscardUnknown(m)
}

var xxx_messageInfo_FsPermissionProto proto.InternalMessageInfo

func (m *FsPermissionProto) GetPerm() uint32 {
	if m != nil && m.Perm != nil {
		return *m.Perm
	}
	return 0
}

type AclEntryProto struct {
	Type                 *AclEntryProto_AclEntryTypeProto  `protobuf:"varint,1,req,name=type,enum=hadoop.hdfs.AclEntryProto_AclEntryTypeProto" json:"type,omitempty"`
	Scope                *AclEntryProto_AclEntryScopeProto `protobuf:"varint,2,req,name=scope,enum=hadoop.hdfs.AclEntryProto_AclEntryScopeProto" json:"scope,omitempty"`
	Permissions          *AclEntryProto_FsActionProto      `protobuf:"varint,3,req,name=permissions,enum=hadoop.hdfs.AclEntryProto_FsActionProto" json:"permissions,omitempty"`
	Name                 *string                           `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *AclEntryProto) Reset()         { *m = AclEntryProto{} }
func (m *AclEntryProto) String() string { return proto.CompactTextString(m) }
func (*AclEntryProto) ProtoMessage()    {}
func (*AclEntryProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_a452f070aeef01eb, []int{1}
}

func (m *AclEntryProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AclEntryProto.Unmarshal(m, b)
}
func (m *AclEntryProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AclEntryProto.Marshal(b, m, deterministic)
}
func (m *AclEntryProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AclEntryProto.Merge(m, src)
}
func (m *AclEntryProto) XXX_Size() int {
	return xxx_messageInfo_AclEntryProto.Size(m)
}
func (m *AclEntryProto) XXX_DiscardUnknown() {
	xxx_messageInfo_AclEntryProto.DiscardUnknown(m)
}

var xxx_messageInfo_AclEntryProto proto.InternalMessageInfo

func (m *AclEntryProto) GetType() AclEntryProto_AclEntryTypeProto {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return AclEntryProto_USER
}

func (m *AclEntryProto) GetScope() AclEntryProto_AclEntryScopeProto {
	if m != nil && m.Scope != nil {
		return *m.Scope
	}
	return AclEntryProto_ACCESS
}

func (m *AclEntryProto) GetPermissions() AclEntryProto_FsActionProto {
	if m != nil && m.Permissions != nil {
		return *m.Permissions
	}
	return AclEntryProto_NONE
}

func (m *AclEntryProto) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

type AclStatusProto struct {
	Owner                *string            `protobuf:"bytes,1,req,name=owner" json:"owner,omitempty"`
	Group                *string            `protobuf:"bytes,2,req,name=group" json:"group,omitempty"`
	Sticky               *bool              `protobuf:"varint,3,req,name=sticky" json:"sticky,omitempty"`
	Entries              []*AclEntryProto   `protobuf:"bytes,4,rep,name=entries" json:"entries,omitempty"`
	Permission           *FsPermissionProto `protobuf:"bytes,5,opt,name=permission" json:"permission,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AclStatusProto) Reset()         { *m = AclStatusProto{} }
func (m *AclStatusProto) String() string { return proto.CompactTextString(m) }
func (*AclStatusProto) ProtoMessage()    {}
func (*AclStatusProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_a452f070aeef01eb, []int{2}
}

func (m *AclStatusProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AclStatusProto.Unmarshal(m, b)
}
func (m *AclStatusProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AclStatusProto.Marshal(b, m, deterministic)
}
func (m *AclStatusProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AclStatusProto.Merge(m, src)
}
func (m *AclStatusProto) XXX_Size() int {
	return xxx_messageInfo_AclStatusProto.Size(m)
}
func (m *AclStatusProto) XXX_DiscardUnknown() {
	xxx_messageInfo_AclStatusProto.DiscardUnknown(m)
}

var xxx_messageInfo_AclStatusProto proto.InternalMessageInfo

func (m *AclStatusProto) GetOwner() string {
	if m != nil && m.Owner != nil {
		return *m.Owner
	}
	return ""
}

func (m *AclStatusProto) GetGroup() string {
	if m != nil && m.Group != nil {
		return *m.Group
	}
	return ""
}

func (m *AclStatusProto) GetSticky() bool {
	if m != nil && m.Sticky != nil {
		return *m.Sticky
	}
	return false
}

func (m *AclStatusProto) GetEntries() []*AclEntryProto {
	if m != nil {
		return m.Entries
	}
	return nil
}

func (m *AclStatusProto) GetPermission() *FsPermissionProto {
	if m != nil {
		return m.Permission
	}
	return nil
}

type ModifyAclEntriesRequestProto struct {
	Src                  *string          `protobuf:"bytes,1,req,name=src" json:"src,omitempty"`
	AclSpec              []*AclEntryProto `protobuf:"bytes,2,rep,name=aclSpec" json:"aclSpec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ModifyAclEntriesRequestProto) Reset()         { *m = ModifyAclEntriesRequestProto{} }
func (m *ModifyAclEntriesRequestProto) String() string { return proto.CompactTextString(m) }
func (*ModifyAclEntriesRequestProto) ProtoMessage()    {}
func (*ModifyAclEntriesRequestProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_a452f070aeef01eb, []int{3}
}

func (m *ModifyAclEntriesRequestProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModifyAclEntriesRequestProto.Unmarshal(m, b)
}
func (m *ModifyAclEntriesRequestProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModifyAclEntriesRequestProto.Marshal(b, m, deterministic)
}
func (m *ModifyAclEntriesRequestProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModifyAclEntriesRequestProto.Merge(m, src)
}
func (m *ModifyAclEntriesRequestProto) XXX_Size() int {
	return xxx_messageInfo_ModifyAclEntriesRequestProto.Size(m)
}
func (m *ModifyAclEntriesRequestProto) XXX_DiscardUnknown() {
	xxx_messageInfo_ModifyAclEntriesRequestProto.DiscardUnknown(m)
}

var xxx_messageInfo_ModifyAclEntriesRequestProto proto.InternalMessageInfo

func (m *ModifyAclEntriesRequestProto) GetSrc() string {
	if m != nil && m.Src != nil {
		return *m.Src
	}
	return ""
}

func (m *ModifyAclEntriesRequestProto) GetAclSpec() []*AclEntryProto {
	if m != nil {
		return m.AclSpec
	}
	return nil
}

type ModifyAclEntriesResponseProto struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ModifyAclEntriesResponseProto) Reset()         { *m = ModifyAclEntriesResponseProto{} }
func (m *ModifyAclEntriesResponseProto) String() string { return proto.CompactTextString(m) }
func (*ModifyAclEntriesResponseProto) ProtoMessage()    {}
func (*ModifyAclEntriesResponseProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_a452f070aeef01eb, []int{4}
}

func (m *ModifyAclEntriesResponseProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModifyAclEntriesResponseProto.Unmarshal(m, b)
}
func (m *ModifyAclEntriesResponseProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModifyAclEntriesResponseProto.Marshal(b, m, deterministic)
}
func (m *ModifyAclEntriesResponseProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModifyAclEntriesResponseProto.Merge(m, src)
}
func (m *ModifyAclEntriesResponseProto) XXX_Size() int {
	return xxx_messageInfo_ModifyAclEntriesResponseProto.Size(m)
}
func (m *ModifyAclEntriesResponseProto) XXX_DiscardUnknown() {
	xxx_messageInfo_ModifyAclEntriesResponseProto.DiscardUnknown(m)
}

var xxx_messageInfo_ModifyAclEntriesResponseProto proto.InternalMessageInfo

type RemoveAclRequestProto struct {
	Src                  *string  `protobuf:"bytes,1,req,name=src" json:"src,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveAclRequestProto) Reset()         { *m = RemoveAclRequestProto{} }
func (m *RemoveAclRequestProto) String() string { return proto.CompactTextString(m) }
func (*RemoveAclRequestProto) ProtoMessage()    {}
func (*RemoveAclRequestProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_a452f070aeef01eb, []int{5}
}

func (m *RemoveAclRequestProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveAclRequestProto.Unmarshal(m, b)
}
func (m *RemoveAclRequestProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveAclRequestProto.Marshal(b, m, deterministic)
}
func (m *RemoveAclRequestProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveAclRequestProto.Merge(m, src)
}
func (m *RemoveAclRequestProto) XXX_Size() int {
	return xxx_messageInfo_RemoveAclRequestProto.Size(m)
}
func (m *RemoveAclRequestProto) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveAclRequestProto.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveAclRequestProto proto.InternalMessageInfo

func (m *RemoveAclRequestProto) GetSrc() string {
	if m != nil && m.Src != nil {
		return *m.Src
	}
	return ""
}

type RemoveAclResponseProto struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveAclResponseProto) Reset()         { *m = RemoveAclResponseProto{} }
func (m *RemoveAclResponseProto) String() string { return proto.CompactTextString(m) }
func (*RemoveAclResponseProto) ProtoMessage()    {}
func (*RemoveAclResponseProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_a452f070aeef01eb, []int{6}
}

func (m *RemoveAclResponseProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveAclResponseProto.Unmarshal(m, b)
}
func (m *RemoveAclResponseProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveAclResponseProto.Marshal(b, m, deterministic)
}
func (m *RemoveAclResponseProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveAclResponseProto.Merge(m, src)
}
func (m *RemoveAclResponseProto) XXX_Size() int {
	return xxx_messageInfo_RemoveAclResponseProto.Size(m)
}
func (m *RemoveAclResponseProto) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveAclResponseProto.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveAclResponseProto proto.InternalMessageInfo

type RemoveAclEntriesRequestProto struct {
	Src                  *string          `protobuf:"bytes,1,req,name=src" json:"src,omitempty"`
	AclSpec              []*AclEntryProto `protobuf:"bytes,2,rep,name=aclSpec" json:"aclSpec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *RemoveAclEntriesRequestProto) Reset()         { *m = RemoveAclEntriesRequestProto{} }
func (m *RemoveAclEntriesRequestProto) String() string { return proto.CompactTextString(m) }
func (*RemoveAclEntriesRequestProto) ProtoMessage()    {}
func (*RemoveAclEntriesRequestProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_a452f070aeef01eb, []int{7}
}

func (m *RemoveAclEntriesRequestProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveAclEntriesRequestProto.Unmarshal(m, b)
}
func (m *RemoveAclEntriesRequestProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveAclEntriesRequestProto.Marshal(b, m, deterministic)
}
func (m *RemoveAclEntriesRequestProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveAclEntriesRequestProto.Merge(m, src)
}
func (m *RemoveAclEntriesRequestProto) XXX_Size() int {
	return xxx_messageInfo_RemoveAclEntriesRequestProto.Size(m)
}
func (m *RemoveAclEntriesRequestProto) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveAclEntriesRequestProto.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveAclEntriesRequestProto proto.InternalMessageInfo

func (m *RemoveAclEntriesRequestProto) GetSrc() string {
	if m != nil && m.Src != nil {
		return *m.Src
	}
	return ""
}

func (m *RemoveAclEntriesRequestProto) GetAclSpec() []*AclEntryProto {
	if m != nil {
		return m.AclSpec
	}
	return nil
}

type RemoveAclEntriesResponseProto struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveAclEntriesResponseProto) Reset()         { *m = RemoveAclEntriesResponseProto{} }
func (m *RemoveAclEntriesResponseProto) String() string { return proto.CompactTextString(m) }
func (*RemoveAclEntriesResponseProto) ProtoMessage()    {}
func (*RemoveAclEntriesResponseProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_a452f070aeef01eb, []int{8}
}

func (m *RemoveAclEntriesResponseProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveAclEntriesResponseProto.Unmarshal(m, b)
}
func (m *RemoveAclEntriesResponseProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveAclEntriesResponseProto.Marshal(b, m, deterministic)
}
func (m *RemoveAclEntriesResponseProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveAclEntriesResponseProto.Merge(m, src)
}
func (m *RemoveAclEntriesResponseProto) XXX_Size() int {
	return xxx_messageInfo_RemoveAclEntriesResponseProto.Size(m)
}
func (m *RemoveAclEntriesResponseProto) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveAclEntriesResponseProto.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveAclEntriesResponseProto proto.InternalMessageInfo

type RemoveDefaultAclRequestProto struct {
	Src                  *string  `protobuf:"bytes,1,req,name=src" json:"src,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveDefaultAclRequestProto) Reset()         { *m = RemoveDefaultAclRequestProto{} }
func (m *RemoveDefaultAclRequestProto) String() string { return proto.CompactTextString(m) }
func (*RemoveDefaultAclRequestProto) ProtoMessage()    {}
func (*RemoveDefaultAclRequestProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_a452f070aeef01eb, []int{9}
}

func (m *RemoveDefaultAclRequestProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveDefaultAclRequestProto.Unmarshal(m, b)
}
func (m *RemoveDefaultAclRequestProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveDefaultAclRequestProto.Marshal(b, m, deterministic)
}
func (m *RemoveDefaultAclRequestProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveDefaultAclRequestProto.Merge(m, src)
}
func (m *RemoveDefaultAclRequestProto) XXX_Size() int {
	return xxx_messageInfo_RemoveDefaultAclRequestProto.Size(m)
}
func (m *RemoveDefaultAclRequestProto) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveDefaultAclRequestProto.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveDefaultAclRequestProto proto.InternalMessageInfo

func (m *RemoveDefaultAclRequestProto) GetSrc() string {
	if m != nil && m.Src != nil {
		return *m.Src
	}
	return ""
}

type RemoveDefaultAclResponseProto struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveDefaultAclResponseProto) Reset()         { *m = RemoveDefaultAclResponseProto{} }
func (m *RemoveDefaultAclResponseProto) String() string { return proto.CompactTextString(m) }
func (*RemoveDefaultAclResponseProto) ProtoMessage()    {}
func (*RemoveDefaultAclResponseProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_a452f070aeef01eb, []int{10}
}

func (m *RemoveDefaultAclResponseProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveDefaultAclResponseProto.Unmarshal(m, b)
}
func (m *RemoveDefaultAclResponseProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveDefaultAclResponseProto.Marshal(b, m, deterministic)
}
func (m *RemoveDefaultAclResponseProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveDefaultAclResponseProto.Merge(m, src)
}
func (m *RemoveDefaultAclResponseProto) XXX_Size() int {
	return xxx_messageInfo_RemoveDefaultAclResponseProto.Size(m)
}
func (m *RemoveDefaultAclResponseProto) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveDefaultAclResponseProto.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveDefaultAclResponseProto proto.InternalMessageInfo

type SetAclRequestProto struct {
	Src                  *string          `protobuf:"bytes,1,req,name=src" json:"src,omitempty"`
	AclSpec              []*AclEntryProto `protobuf:"bytes,2,rep,name=aclSpec" json:"aclSpec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SetAclRequestProto) Reset()         { *m = SetAclRequestProto{} }
func (m *SetAclRequestProto) String() string { return proto.CompactTextString(m) }
func (*SetAclRequestProto) ProtoMessage()    {}
func (*SetAclRequestProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_a452f070aeef01eb, []int{11}
}

func (m *SetAclRequestProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetAclRequestProto.Unmarshal(m, b)
}
func (m *SetAclRequestProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetAclRequestProto.Marshal(b, m, deterministic)
}
func (m *SetAclRequestProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetAclRequestProto.Merge(m, src)
}
func (m *SetAclRequestProto) XXX_Size() int {
	return xxx_messageInfo_SetAclRequestProto.Size(m)
}
func (m *SetAclRequestProto) XXX_DiscardUnknown() {
	xxx_messageInfo_SetAclRequestProto.DiscardUnknown(m)
}

var xxx_messageInfo_SetAclRequestProto proto.InternalMessageInfo

func (m *SetAclRequestProto) GetSrc() string {
	if m != nil && m.Src != nil {
		return *m.Src
	}
	return ""
}

func (m *SetAclRequestProto) GetAclSpec() []*AclEntryProto {
	if m != nil {
		return m.AclSpec
	}
	return nil
}

type SetAclResponseProto struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetAclResponseProto) Reset()         { *m = SetAclResponseProto{} }
func (m *SetAclResponseProto) String() string { return proto.CompactTextString(m) }
func (*SetAclResponseProto) ProtoMessage()    {}
func (*SetAclResponseProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_a452f070aeef01eb, []int{12}
}

func (m *SetAclResponseProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetAclResponseProto.Unmarshal(m, b)
}
func (m *SetAclResponseProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetAclResponseProto.Marshal(b, m, deterministic)
}
func (m *SetAclResponseProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetAclResponseProto.Merge(m, src)
}
func (m *SetAclResponseProto) XXX_Size() int {
	return xxx_messageInfo_SetAclResponseProto.Size(m)
}
func (m *SetAclResponseProto) XXX_DiscardUnknown() {
	xxx_messageInfo_SetAclResponseProto.DiscardUnknown(m)
}

var xxx_messageInfo_SetAclResponseProto proto.InternalMessageInfo

type GetAclStatusRequestProto struct {
	Src                  *string  `protobuf:"bytes,1,req,name=src" json:"src,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAclStatusRequestProto) Reset()         { *m = GetAclStatusRequestProto{} }
func (m *GetAclStatusRequestProto) String() string { return proto.CompactTextString(m) }
func (*GetAclStatusRequestProto) ProtoMessage()    {}
func (*GetAclStatusRequestProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_a452f070aeef01eb, []int{13}
}

func (m *GetAclStatusRequestProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAclStatusRequestProto.Unmarshal(m, b)
}
func (m *GetAclStatusRequestProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAclStatusRequestProto.Marshal(b, m, deterministic)
}
func (m *GetAclStatusRequestProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAclStatusRequestProto.Merge(m, src)
}
func (m *GetAclStatusRequestProto) XXX_Size() int {
	return xxx_messageInfo_GetAclStatusRequestProto.Size(m)
}
func (m *GetAclStatusRequestProto) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAclStatusRequestProto.DiscardUnknown(m)
}

var xxx_messageInfo_GetAclStatusRequestProto proto.InternalMessageInfo

func (m *GetAclStatusRequestProto) GetSrc() string {
	if m != nil && m.Src != nil {
		return *m.Src
	}
	return ""
}

type GetAclStatusResponseProto struct {
	Result               *AclStatusProto `protobuf:"bytes,1,req,name=result" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GetAclStatusResponseProto) Reset()         { *m = GetAclStatusResponseProto{} }
func (m *GetAclStatusResponseProto) String() string { return proto.CompactTextString(m) }
func (*GetAclStatusResponseProto) ProtoMessage()    {}
func (*GetAclStatusResponseProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_a452f070aeef01eb, []int{14}
}

func (m *GetAclStatusResponseProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAclStatusResponseProto.Unmarshal(m, b)
}
func (m *GetAclStatusResponseProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAclStatusResponseProto.Marshal(b, m, deterministic)
}
func (m *GetAclStatusResponseProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAclStatusResponseProto.Merge(m, src)
}
func (m *GetAclStatusResponseProto) XXX_Size() int {
	return xxx_messageInfo_GetAclStatusResponseProto.Size(m)
}
func (m *GetAclStatusResponseProto) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAclStatusResponseProto.DiscardUnknown(m)
}

var xxx_messageInfo_GetAclStatusResponseProto proto.InternalMessageInfo

func (m *GetAclStatusResponseProto) GetResult() *AclStatusProto {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterEnum("hadoop.hdfs.AclEntryProto_AclEntryScopeProto", AclEntryProto_AclEntryScopeProto_name, AclEntryProto_AclEntryScopeProto_value)
	proto.RegisterEnum("hadoop.hdfs.AclEntryProto_AclEntryTypeProto", AclEntryProto_AclEntryTypeProto_name, AclEntryProto_AclEntryTypeProto_value)
	proto.RegisterEnum("hadoop.hdfs.AclEntryProto_FsActionProto", AclEntryProto_FsActionProto_name, AclEntryProto_FsActionProto_value)
	proto.RegisterType((*FsPermissionProto)(nil), "hadoop.hdfs.FsPermissionProto")
	proto.RegisterType((*AclEntryProto)(nil), "hadoop.hdfs.AclEntryProto")
	proto.RegisterType((*AclStatusProto)(nil), "hadoop.hdfs.AclStatusProto")
	proto.RegisterType((*ModifyAclEntriesRequestProto)(nil), "hadoop.hdfs.ModifyAclEntriesRequestProto")
	proto.RegisterType((*ModifyAclEntriesResponseProto)(nil), "hadoop.hdfs.ModifyAclEntriesResponseProto")
	proto.RegisterType((*RemoveAclRequestProto)(nil), "hadoop.hdfs.RemoveAclRequestProto")
	proto.RegisterType((*RemoveAclResponseProto)(nil), "hadoop.hdfs.RemoveAclResponseProto")
	proto.RegisterType((*RemoveAclEntriesRequestProto)(nil), "hadoop.hdfs.RemoveAclEntriesRequestProto")
	proto.RegisterType((*RemoveAclEntriesResponseProto)(nil), "hadoop.hdfs.RemoveAclEntriesResponseProto")
	proto.RegisterType((*RemoveDefaultAclRequestProto)(nil), "hadoop.hdfs.RemoveDefaultAclRequestProto")
	proto.RegisterType((*RemoveDefaultAclResponseProto)(nil), "hadoop.hdfs.RemoveDefaultAclResponseProto")
	proto.RegisterType((*SetAclRequestProto)(nil), "hadoop.hdfs.SetAclRequestProto")
	proto.RegisterType((*SetAclResponseProto)(nil), "hadoop.hdfs.SetAclResponseProto")
	proto.RegisterType((*GetAclStatusRequestProto)(nil), "hadoop.hdfs.GetAclStatusRequestProto")
	proto.RegisterType((*GetAclStatusResponseProto)(nil), "hadoop.hdfs.GetAclStatusResponseProto")
}

func init() { proto.RegisterFile("acl.proto", fileDescriptor_a452f070aeef01eb) }

var fileDescriptor_a452f070aeef01eb = []byte{
	// 614 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xed, 0x6e, 0xd3, 0x30,
	0x14, 0x5d, 0x9a, 0xb4, 0x5b, 0x6e, 0xd6, 0xc9, 0x33, 0x6c, 0x0a, 0x30, 0xa0, 0x8a, 0x84, 0x28,
	0xd2, 0x56, 0xa1, 0x02, 0x3f, 0x41, 0x64, 0x9d, 0x37, 0x3e, 0xf6, 0x51, 0x39, 0xad, 0xe0, 0x07,
	0xd2, 0x14, 0xb9, 0xee, 0x1a, 0x91, 0xc6, 0x21, 0x4e, 0x41, 0xfd, 0xc3, 0xb3, 0xf0, 0x3c, 0x3c,
	0x07, 0x0f, 0x82, 0xe2, 0xa4, 0x25, 0xa5, 0x62, 0x45, 0x48, 0xfc, 0xbb, 0xbe, 0x3e, 0xe7, 0xdc,
	0x73, 0x74, 0x13, 0x83, 0xe9, 0xb3, 0xb0, 0x15, 0x27, 0x22, 0x15, 0xd8, 0x1a, 0xf9, 0x03, 0x21,
	0xe2, 0xd6, 0x68, 0x30, 0x94, 0xce, 0x43, 0xd8, 0x3e, 0x96, 0x5d, 0x9e, 0x8c, 0x03, 0x29, 0x03,
	0x11, 0x75, 0x15, 0x02, 0x83, 0x11, 0xf3, 0x64, 0x6c, 0x6b, 0x8d, 0x4a, 0xb3, 0x4e, 0x55, 0xed,
	0xfc, 0xd0, 0xa1, 0xee, 0xb2, 0x90, 0x44, 0x69, 0x32, 0xcd, 0x51, 0x2f, 0xc1, 0x48, 0xa7, 0x31,
	0x57, 0xa8, 0xad, 0xf6, 0x7e, 0xab, 0x24, 0xdb, 0x5a, 0x40, 0xce, 0x4f, 0xbd, 0x69, 0xcc, 0x55,
	0x87, 0x2a, 0x26, 0xee, 0x40, 0x55, 0x32, 0x11, 0x73, 0xbb, 0xa2, 0x24, 0x0e, 0xfe, 0x42, 0xc2,
	0xcb, 0xf0, 0xb9, 0x46, 0xce, 0xc5, 0x6f, 0xc0, 0x8a, 0xe7, 0xfe, 0xa5, 0xad, 0x2b, 0xa9, 0xe6,
	0x35, 0x52, 0xc7, 0xd2, 0x65, 0xe9, 0x2c, 0x2b, 0x2d, 0x93, 0xb3, 0xe0, 0x91, 0x3f, 0xe6, 0xb6,
	0xd1, 0xd0, 0x9a, 0x26, 0x55, 0xb5, 0x73, 0x00, 0x78, 0x79, 0x38, 0x06, 0xa8, 0xb9, 0x9d, 0x0e,
	0xf1, 0x3c, 0xb4, 0x86, 0x2d, 0x58, 0x3f, 0x22, 0xc7, 0x6e, 0xff, 0xb4, 0x87, 0x34, 0xe7, 0x39,
	0x6c, 0x2f, 0xc5, 0xc5, 0x1b, 0x60, 0xf4, 0x3d, 0x42, 0xd1, 0x1a, 0x36, 0xa1, 0x7a, 0x42, 0x2f,
	0xfa, 0x5d, 0xa4, 0x65, 0xcd, 0x33, 0xd7, 0x7b, 0x8b, 0x2a, 0x59, 0xf3, 0xa2, 0xf7, 0x8a, 0x50,
	0xa4, 0x3b, 0x5f, 0xa1, 0xbe, 0xe0, 0x2f, 0x43, 0x9d, 0x5f, 0x9c, 0x93, 0x7c, 0x0c, 0x79, 0x4f,
	0x3a, 0xfd, 0x1e, 0x41, 0x5a, 0x46, 0x79, 0x47, 0x5f, 0xf7, 0x08, 0xaa, 0xe0, 0x6d, 0xa8, 0xab,
	0xf2, 0x72, 0x76, 0xab, 0x67, 0x24, 0x4a, 0xdc, 0x23, 0x64, 0x60, 0x04, 0x9b, 0x59, 0x35, 0xbf,
	0xab, 0xe2, 0x2d, 0x00, 0xd5, 0xc9, 0xe9, 0x35, 0xbc, 0x09, 0x1b, 0x5d, 0x42, 0xcf, 0x2e, 0xdd,
	0xd3, 0x53, 0xb4, 0xee, 0x7c, 0xd7, 0x60, 0xcb, 0x65, 0xa1, 0x97, 0xfa, 0xe9, 0x44, 0xe6, 0x0e,
	0x6e, 0x42, 0x55, 0x7c, 0x89, 0x78, 0xa2, 0x16, 0x6d, 0xd2, 0xfc, 0x90, 0x75, 0xaf, 0x12, 0x31,
	0x89, 0xd5, 0xee, 0x4c, 0x9a, 0x1f, 0xf0, 0x2e, 0xd4, 0x64, 0x1a, 0xb0, 0x8f, 0x53, 0xb5, 0x87,
	0x0d, 0x5a, 0x9c, 0xf0, 0x53, 0x58, 0xe7, 0x51, 0x9a, 0x04, 0x5c, 0xda, 0x46, 0x43, 0x6f, 0x5a,
	0xed, 0xdb, 0x7f, 0x5e, 0x10, 0x9d, 0x41, 0xf1, 0x0b, 0x80, 0x5f, 0xdb, 0xb1, 0xab, 0x0d, 0xad,
	0x69, 0xb5, 0xef, 0x2d, 0x10, 0x97, 0xbe, 0x5d, 0x5a, 0x62, 0x38, 0x43, 0xd8, 0x3b, 0x13, 0x83,
	0x60, 0x38, 0x2d, 0xf4, 0x03, 0x2e, 0x29, 0xff, 0x34, 0xe1, 0x32, 0xcd, 0x93, 0x21, 0xd0, 0x65,
	0xc2, 0x8a, 0x5c, 0x59, 0x99, 0xf9, 0xf4, 0x59, 0xe8, 0xc5, 0x9c, 0xd9, 0x95, 0xd5, 0x3e, 0x0b,
	0xa8, 0x73, 0x1f, 0xee, 0x2e, 0xcf, 0x91, 0xb1, 0x88, 0x64, 0xbe, 0x7f, 0xe7, 0x11, 0xec, 0x50,
	0x3e, 0x16, 0x9f, 0xb9, 0xcb, 0xc2, 0xeb, 0x1d, 0x38, 0x36, 0xec, 0x96, 0xa0, 0x65, 0x91, 0x21,
	0xec, 0xcd, 0x6f, 0xfe, 0x73, 0x9a, 0xe5, 0x39, 0x65, 0x23, 0x8f, 0x67, 0x46, 0x8e, 0xf8, 0xd0,
	0x9f, 0x84, 0xe9, 0xea, 0x50, 0x73, 0xc9, 0x32, 0xa3, 0x2c, 0xf9, 0x01, 0xb0, 0xc7, 0x57, 0x0b,
	0xfd, 0x63, 0xa2, 0x1d, 0xb8, 0x31, 0x53, 0x2f, 0x0f, 0xdd, 0x07, 0xfb, 0x44, 0xb5, 0xf3, 0xaf,
	0x7d, 0x45, 0x86, 0x2e, 0xdc, 0x5a, 0x44, 0x97, 0xa4, 0xf0, 0x13, 0xa8, 0x25, 0x5c, 0x4e, 0xc2,
	0x54, 0x31, 0xac, 0xf6, 0x9d, 0xdf, 0x6d, 0x95, 0x7e, 0x28, 0x5a, 0x40, 0x0f, 0x9f, 0xc1, 0x03,
	0x91, 0x5c, 0xb5, 0xfc, 0xd8, 0x67, 0x23, 0xbe, 0x40, 0x50, 0x0f, 0x35, 0x13, 0xc5, 0x8b, 0x7d,
	0x68, 0xba, 0x2c, 0x54, 0x54, 0xf9, 0x4d, 0xd3, 0x7e, 0x06, 0x00, 0x00, 0xff, 0xff, 0x0b, 0xdb,
	0x1b, 0x41, 0xcb, 0x05, 0x00, 0x00,
}
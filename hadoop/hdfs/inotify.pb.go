// Code generated by protoc-gen-go. DO NOT EDIT.
// source: inotify.proto

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

type EventType int32

const (
	EventType_EVENT_CREATE   EventType = 0
	EventType_EVENT_CLOSE    EventType = 1
	EventType_EVENT_APPEND   EventType = 2
	EventType_EVENT_RENAME   EventType = 3
	EventType_EVENT_METADATA EventType = 4
	EventType_EVENT_UNLINK   EventType = 5
	EventType_EVENT_TRUNCATE EventType = 6
)

var EventType_name = map[int32]string{
	0: "EVENT_CREATE",
	1: "EVENT_CLOSE",
	2: "EVENT_APPEND",
	3: "EVENT_RENAME",
	4: "EVENT_METADATA",
	5: "EVENT_UNLINK",
	6: "EVENT_TRUNCATE",
}

var EventType_value = map[string]int32{
	"EVENT_CREATE":   0,
	"EVENT_CLOSE":    1,
	"EVENT_APPEND":   2,
	"EVENT_RENAME":   3,
	"EVENT_METADATA": 4,
	"EVENT_UNLINK":   5,
	"EVENT_TRUNCATE": 6,
}

func (x EventType) Enum() *EventType {
	p := new(EventType)
	*p = x
	return p
}

func (x EventType) String() string {
	return proto.EnumName(EventType_name, int32(x))
}

func (x *EventType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EventType_value, data, "EventType")
	if err != nil {
		return err
	}
	*x = EventType(value)
	return nil
}

func (EventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5fe8964429f4f5e5, []int{0}
}

type INodeType int32

const (
	INodeType_I_TYPE_FILE      INodeType = 0
	INodeType_I_TYPE_DIRECTORY INodeType = 1
	INodeType_I_TYPE_SYMLINK   INodeType = 2
)

var INodeType_name = map[int32]string{
	0: "I_TYPE_FILE",
	1: "I_TYPE_DIRECTORY",
	2: "I_TYPE_SYMLINK",
}

var INodeType_value = map[string]int32{
	"I_TYPE_FILE":      0,
	"I_TYPE_DIRECTORY": 1,
	"I_TYPE_SYMLINK":   2,
}

func (x INodeType) Enum() *INodeType {
	p := new(INodeType)
	*p = x
	return p
}

func (x INodeType) String() string {
	return proto.EnumName(INodeType_name, int32(x))
}

func (x *INodeType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(INodeType_value, data, "INodeType")
	if err != nil {
		return err
	}
	*x = INodeType(value)
	return nil
}

func (INodeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5fe8964429f4f5e5, []int{1}
}

type MetadataUpdateType int32

const (
	MetadataUpdateType_META_TYPE_TIMES       MetadataUpdateType = 0
	MetadataUpdateType_META_TYPE_REPLICATION MetadataUpdateType = 1
	MetadataUpdateType_META_TYPE_OWNER       MetadataUpdateType = 2
	MetadataUpdateType_META_TYPE_PERMS       MetadataUpdateType = 3
	MetadataUpdateType_META_TYPE_ACLS        MetadataUpdateType = 4
	MetadataUpdateType_META_TYPE_XATTRS      MetadataUpdateType = 5
)

var MetadataUpdateType_name = map[int32]string{
	0: "META_TYPE_TIMES",
	1: "META_TYPE_REPLICATION",
	2: "META_TYPE_OWNER",
	3: "META_TYPE_PERMS",
	4: "META_TYPE_ACLS",
	5: "META_TYPE_XATTRS",
}

var MetadataUpdateType_value = map[string]int32{
	"META_TYPE_TIMES":       0,
	"META_TYPE_REPLICATION": 1,
	"META_TYPE_OWNER":       2,
	"META_TYPE_PERMS":       3,
	"META_TYPE_ACLS":        4,
	"META_TYPE_XATTRS":      5,
}

func (x MetadataUpdateType) Enum() *MetadataUpdateType {
	p := new(MetadataUpdateType)
	*p = x
	return p
}

func (x MetadataUpdateType) String() string {
	return proto.EnumName(MetadataUpdateType_name, int32(x))
}

func (x *MetadataUpdateType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MetadataUpdateType_value, data, "MetadataUpdateType")
	if err != nil {
		return err
	}
	*x = MetadataUpdateType(value)
	return nil
}

func (MetadataUpdateType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5fe8964429f4f5e5, []int{2}
}

type EventProto struct {
	Type                 *EventType `protobuf:"varint,1,req,name=type,enum=hadoop.hdfs.EventType" json:"type,omitempty"`
	Contents             []byte     `protobuf:"bytes,2,req,name=contents" json:"contents,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *EventProto) Reset()         { *m = EventProto{} }
func (m *EventProto) String() string { return proto.CompactTextString(m) }
func (*EventProto) ProtoMessage()    {}
func (*EventProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fe8964429f4f5e5, []int{0}
}

func (m *EventProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventProto.Unmarshal(m, b)
}
func (m *EventProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventProto.Marshal(b, m, deterministic)
}
func (m *EventProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventProto.Merge(m, src)
}
func (m *EventProto) XXX_Size() int {
	return xxx_messageInfo_EventProto.Size(m)
}
func (m *EventProto) XXX_DiscardUnknown() {
	xxx_messageInfo_EventProto.DiscardUnknown(m)
}

var xxx_messageInfo_EventProto proto.InternalMessageInfo

func (m *EventProto) GetType() EventType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return EventType_EVENT_CREATE
}

func (m *EventProto) GetContents() []byte {
	if m != nil {
		return m.Contents
	}
	return nil
}

type EventBatchProto struct {
	Txid                 *int64        `protobuf:"varint,1,req,name=txid" json:"txid,omitempty"`
	Events               []*EventProto `protobuf:"bytes,2,rep,name=events" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *EventBatchProto) Reset()         { *m = EventBatchProto{} }
func (m *EventBatchProto) String() string { return proto.CompactTextString(m) }
func (*EventBatchProto) ProtoMessage()    {}
func (*EventBatchProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fe8964429f4f5e5, []int{1}
}

func (m *EventBatchProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventBatchProto.Unmarshal(m, b)
}
func (m *EventBatchProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventBatchProto.Marshal(b, m, deterministic)
}
func (m *EventBatchProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventBatchProto.Merge(m, src)
}
func (m *EventBatchProto) XXX_Size() int {
	return xxx_messageInfo_EventBatchProto.Size(m)
}
func (m *EventBatchProto) XXX_DiscardUnknown() {
	xxx_messageInfo_EventBatchProto.DiscardUnknown(m)
}

var xxx_messageInfo_EventBatchProto proto.InternalMessageInfo

func (m *EventBatchProto) GetTxid() int64 {
	if m != nil && m.Txid != nil {
		return *m.Txid
	}
	return 0
}

func (m *EventBatchProto) GetEvents() []*EventProto {
	if m != nil {
		return m.Events
	}
	return nil
}

type CreateEventProto struct {
	Type                 *INodeType         `protobuf:"varint,1,req,name=type,enum=hadoop.hdfs.INodeType" json:"type,omitempty"`
	Path                 *string            `protobuf:"bytes,2,req,name=path" json:"path,omitempty"`
	Ctime                *int64             `protobuf:"varint,3,req,name=ctime" json:"ctime,omitempty"`
	OwnerName            *string            `protobuf:"bytes,4,req,name=ownerName" json:"ownerName,omitempty"`
	GroupName            *string            `protobuf:"bytes,5,req,name=groupName" json:"groupName,omitempty"`
	Perms                *FsPermissionProto `protobuf:"bytes,6,req,name=perms" json:"perms,omitempty"`
	Replication          *int32             `protobuf:"varint,7,opt,name=replication" json:"replication,omitempty"`
	SymlinkTarget        *string            `protobuf:"bytes,8,opt,name=symlinkTarget" json:"symlinkTarget,omitempty"`
	Overwrite            *bool              `protobuf:"varint,9,opt,name=overwrite" json:"overwrite,omitempty"`
	DefaultBlockSize     *int64             `protobuf:"varint,10,opt,name=defaultBlockSize,def=0" json:"defaultBlockSize,omitempty"`
	ErasureCoded         *bool              `protobuf:"varint,11,opt,name=erasureCoded" json:"erasureCoded,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CreateEventProto) Reset()         { *m = CreateEventProto{} }
func (m *CreateEventProto) String() string { return proto.CompactTextString(m) }
func (*CreateEventProto) ProtoMessage()    {}
func (*CreateEventProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fe8964429f4f5e5, []int{2}
}

func (m *CreateEventProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateEventProto.Unmarshal(m, b)
}
func (m *CreateEventProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateEventProto.Marshal(b, m, deterministic)
}
func (m *CreateEventProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateEventProto.Merge(m, src)
}
func (m *CreateEventProto) XXX_Size() int {
	return xxx_messageInfo_CreateEventProto.Size(m)
}
func (m *CreateEventProto) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateEventProto.DiscardUnknown(m)
}

var xxx_messageInfo_CreateEventProto proto.InternalMessageInfo

const Default_CreateEventProto_DefaultBlockSize int64 = 0

func (m *CreateEventProto) GetType() INodeType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return INodeType_I_TYPE_FILE
}

func (m *CreateEventProto) GetPath() string {
	if m != nil && m.Path != nil {
		return *m.Path
	}
	return ""
}

func (m *CreateEventProto) GetCtime() int64 {
	if m != nil && m.Ctime != nil {
		return *m.Ctime
	}
	return 0
}

func (m *CreateEventProto) GetOwnerName() string {
	if m != nil && m.OwnerName != nil {
		return *m.OwnerName
	}
	return ""
}

func (m *CreateEventProto) GetGroupName() string {
	if m != nil && m.GroupName != nil {
		return *m.GroupName
	}
	return ""
}

func (m *CreateEventProto) GetPerms() *FsPermissionProto {
	if m != nil {
		return m.Perms
	}
	return nil
}

func (m *CreateEventProto) GetReplication() int32 {
	if m != nil && m.Replication != nil {
		return *m.Replication
	}
	return 0
}

func (m *CreateEventProto) GetSymlinkTarget() string {
	if m != nil && m.SymlinkTarget != nil {
		return *m.SymlinkTarget
	}
	return ""
}

func (m *CreateEventProto) GetOverwrite() bool {
	if m != nil && m.Overwrite != nil {
		return *m.Overwrite
	}
	return false
}

func (m *CreateEventProto) GetDefaultBlockSize() int64 {
	if m != nil && m.DefaultBlockSize != nil {
		return *m.DefaultBlockSize
	}
	return Default_CreateEventProto_DefaultBlockSize
}

func (m *CreateEventProto) GetErasureCoded() bool {
	if m != nil && m.ErasureCoded != nil {
		return *m.ErasureCoded
	}
	return false
}

type CloseEventProto struct {
	Path                 *string  `protobuf:"bytes,1,req,name=path" json:"path,omitempty"`
	FileSize             *int64   `protobuf:"varint,2,req,name=fileSize" json:"fileSize,omitempty"`
	Timestamp            *int64   `protobuf:"varint,3,req,name=timestamp" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CloseEventProto) Reset()         { *m = CloseEventProto{} }
func (m *CloseEventProto) String() string { return proto.CompactTextString(m) }
func (*CloseEventProto) ProtoMessage()    {}
func (*CloseEventProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fe8964429f4f5e5, []int{3}
}

func (m *CloseEventProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloseEventProto.Unmarshal(m, b)
}
func (m *CloseEventProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloseEventProto.Marshal(b, m, deterministic)
}
func (m *CloseEventProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloseEventProto.Merge(m, src)
}
func (m *CloseEventProto) XXX_Size() int {
	return xxx_messageInfo_CloseEventProto.Size(m)
}
func (m *CloseEventProto) XXX_DiscardUnknown() {
	xxx_messageInfo_CloseEventProto.DiscardUnknown(m)
}

var xxx_messageInfo_CloseEventProto proto.InternalMessageInfo

func (m *CloseEventProto) GetPath() string {
	if m != nil && m.Path != nil {
		return *m.Path
	}
	return ""
}

func (m *CloseEventProto) GetFileSize() int64 {
	if m != nil && m.FileSize != nil {
		return *m.FileSize
	}
	return 0
}

func (m *CloseEventProto) GetTimestamp() int64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

type TruncateEventProto struct {
	Path                 *string  `protobuf:"bytes,1,req,name=path" json:"path,omitempty"`
	FileSize             *int64   `protobuf:"varint,2,req,name=fileSize" json:"fileSize,omitempty"`
	Timestamp            *int64   `protobuf:"varint,3,req,name=timestamp" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TruncateEventProto) Reset()         { *m = TruncateEventProto{} }
func (m *TruncateEventProto) String() string { return proto.CompactTextString(m) }
func (*TruncateEventProto) ProtoMessage()    {}
func (*TruncateEventProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fe8964429f4f5e5, []int{4}
}

func (m *TruncateEventProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TruncateEventProto.Unmarshal(m, b)
}
func (m *TruncateEventProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TruncateEventProto.Marshal(b, m, deterministic)
}
func (m *TruncateEventProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TruncateEventProto.Merge(m, src)
}
func (m *TruncateEventProto) XXX_Size() int {
	return xxx_messageInfo_TruncateEventProto.Size(m)
}
func (m *TruncateEventProto) XXX_DiscardUnknown() {
	xxx_messageInfo_TruncateEventProto.DiscardUnknown(m)
}

var xxx_messageInfo_TruncateEventProto proto.InternalMessageInfo

func (m *TruncateEventProto) GetPath() string {
	if m != nil && m.Path != nil {
		return *m.Path
	}
	return ""
}

func (m *TruncateEventProto) GetFileSize() int64 {
	if m != nil && m.FileSize != nil {
		return *m.FileSize
	}
	return 0
}

func (m *TruncateEventProto) GetTimestamp() int64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

type AppendEventProto struct {
	Path                 *string  `protobuf:"bytes,1,req,name=path" json:"path,omitempty"`
	NewBlock             *bool    `protobuf:"varint,2,opt,name=newBlock,def=0" json:"newBlock,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppendEventProto) Reset()         { *m = AppendEventProto{} }
func (m *AppendEventProto) String() string { return proto.CompactTextString(m) }
func (*AppendEventProto) ProtoMessage()    {}
func (*AppendEventProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fe8964429f4f5e5, []int{5}
}

func (m *AppendEventProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppendEventProto.Unmarshal(m, b)
}
func (m *AppendEventProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppendEventProto.Marshal(b, m, deterministic)
}
func (m *AppendEventProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppendEventProto.Merge(m, src)
}
func (m *AppendEventProto) XXX_Size() int {
	return xxx_messageInfo_AppendEventProto.Size(m)
}
func (m *AppendEventProto) XXX_DiscardUnknown() {
	xxx_messageInfo_AppendEventProto.DiscardUnknown(m)
}

var xxx_messageInfo_AppendEventProto proto.InternalMessageInfo

const Default_AppendEventProto_NewBlock bool = false

func (m *AppendEventProto) GetPath() string {
	if m != nil && m.Path != nil {
		return *m.Path
	}
	return ""
}

func (m *AppendEventProto) GetNewBlock() bool {
	if m != nil && m.NewBlock != nil {
		return *m.NewBlock
	}
	return Default_AppendEventProto_NewBlock
}

type RenameEventProto struct {
	SrcPath              *string  `protobuf:"bytes,1,req,name=srcPath" json:"srcPath,omitempty"`
	DestPath             *string  `protobuf:"bytes,2,req,name=destPath" json:"destPath,omitempty"`
	Timestamp            *int64   `protobuf:"varint,3,req,name=timestamp" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RenameEventProto) Reset()         { *m = RenameEventProto{} }
func (m *RenameEventProto) String() string { return proto.CompactTextString(m) }
func (*RenameEventProto) ProtoMessage()    {}
func (*RenameEventProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fe8964429f4f5e5, []int{6}
}

func (m *RenameEventProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RenameEventProto.Unmarshal(m, b)
}
func (m *RenameEventProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RenameEventProto.Marshal(b, m, deterministic)
}
func (m *RenameEventProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RenameEventProto.Merge(m, src)
}
func (m *RenameEventProto) XXX_Size() int {
	return xxx_messageInfo_RenameEventProto.Size(m)
}
func (m *RenameEventProto) XXX_DiscardUnknown() {
	xxx_messageInfo_RenameEventProto.DiscardUnknown(m)
}

var xxx_messageInfo_RenameEventProto proto.InternalMessageInfo

func (m *RenameEventProto) GetSrcPath() string {
	if m != nil && m.SrcPath != nil {
		return *m.SrcPath
	}
	return ""
}

func (m *RenameEventProto) GetDestPath() string {
	if m != nil && m.DestPath != nil {
		return *m.DestPath
	}
	return ""
}

func (m *RenameEventProto) GetTimestamp() int64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

type MetadataUpdateEventProto struct {
	Path                 *string             `protobuf:"bytes,1,req,name=path" json:"path,omitempty"`
	Type                 *MetadataUpdateType `protobuf:"varint,2,req,name=type,enum=hadoop.hdfs.MetadataUpdateType" json:"type,omitempty"`
	Mtime                *int64              `protobuf:"varint,3,opt,name=mtime" json:"mtime,omitempty"`
	Atime                *int64              `protobuf:"varint,4,opt,name=atime" json:"atime,omitempty"`
	Replication          *int32              `protobuf:"varint,5,opt,name=replication" json:"replication,omitempty"`
	OwnerName            *string             `protobuf:"bytes,6,opt,name=ownerName" json:"ownerName,omitempty"`
	GroupName            *string             `protobuf:"bytes,7,opt,name=groupName" json:"groupName,omitempty"`
	Perms                *FsPermissionProto  `protobuf:"bytes,8,opt,name=perms" json:"perms,omitempty"`
	Acls                 []*AclEntryProto    `protobuf:"bytes,9,rep,name=acls" json:"acls,omitempty"`
	XAttrs               []*XAttrProto       `protobuf:"bytes,10,rep,name=xAttrs" json:"xAttrs,omitempty"`
	XAttrsRemoved        *bool               `protobuf:"varint,11,opt,name=xAttrsRemoved" json:"xAttrsRemoved,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *MetadataUpdateEventProto) Reset()         { *m = MetadataUpdateEventProto{} }
func (m *MetadataUpdateEventProto) String() string { return proto.CompactTextString(m) }
func (*MetadataUpdateEventProto) ProtoMessage()    {}
func (*MetadataUpdateEventProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fe8964429f4f5e5, []int{7}
}

func (m *MetadataUpdateEventProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetadataUpdateEventProto.Unmarshal(m, b)
}
func (m *MetadataUpdateEventProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetadataUpdateEventProto.Marshal(b, m, deterministic)
}
func (m *MetadataUpdateEventProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetadataUpdateEventProto.Merge(m, src)
}
func (m *MetadataUpdateEventProto) XXX_Size() int {
	return xxx_messageInfo_MetadataUpdateEventProto.Size(m)
}
func (m *MetadataUpdateEventProto) XXX_DiscardUnknown() {
	xxx_messageInfo_MetadataUpdateEventProto.DiscardUnknown(m)
}

var xxx_messageInfo_MetadataUpdateEventProto proto.InternalMessageInfo

func (m *MetadataUpdateEventProto) GetPath() string {
	if m != nil && m.Path != nil {
		return *m.Path
	}
	return ""
}

func (m *MetadataUpdateEventProto) GetType() MetadataUpdateType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return MetadataUpdateType_META_TYPE_TIMES
}

func (m *MetadataUpdateEventProto) GetMtime() int64 {
	if m != nil && m.Mtime != nil {
		return *m.Mtime
	}
	return 0
}

func (m *MetadataUpdateEventProto) GetAtime() int64 {
	if m != nil && m.Atime != nil {
		return *m.Atime
	}
	return 0
}

func (m *MetadataUpdateEventProto) GetReplication() int32 {
	if m != nil && m.Replication != nil {
		return *m.Replication
	}
	return 0
}

func (m *MetadataUpdateEventProto) GetOwnerName() string {
	if m != nil && m.OwnerName != nil {
		return *m.OwnerName
	}
	return ""
}

func (m *MetadataUpdateEventProto) GetGroupName() string {
	if m != nil && m.GroupName != nil {
		return *m.GroupName
	}
	return ""
}

func (m *MetadataUpdateEventProto) GetPerms() *FsPermissionProto {
	if m != nil {
		return m.Perms
	}
	return nil
}

func (m *MetadataUpdateEventProto) GetAcls() []*AclEntryProto {
	if m != nil {
		return m.Acls
	}
	return nil
}

func (m *MetadataUpdateEventProto) GetXAttrs() []*XAttrProto {
	if m != nil {
		return m.XAttrs
	}
	return nil
}

func (m *MetadataUpdateEventProto) GetXAttrsRemoved() bool {
	if m != nil && m.XAttrsRemoved != nil {
		return *m.XAttrsRemoved
	}
	return false
}

type UnlinkEventProto struct {
	Path                 *string  `protobuf:"bytes,1,req,name=path" json:"path,omitempty"`
	Timestamp            *int64   `protobuf:"varint,2,req,name=timestamp" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnlinkEventProto) Reset()         { *m = UnlinkEventProto{} }
func (m *UnlinkEventProto) String() string { return proto.CompactTextString(m) }
func (*UnlinkEventProto) ProtoMessage()    {}
func (*UnlinkEventProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fe8964429f4f5e5, []int{8}
}

func (m *UnlinkEventProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnlinkEventProto.Unmarshal(m, b)
}
func (m *UnlinkEventProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnlinkEventProto.Marshal(b, m, deterministic)
}
func (m *UnlinkEventProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnlinkEventProto.Merge(m, src)
}
func (m *UnlinkEventProto) XXX_Size() int {
	return xxx_messageInfo_UnlinkEventProto.Size(m)
}
func (m *UnlinkEventProto) XXX_DiscardUnknown() {
	xxx_messageInfo_UnlinkEventProto.DiscardUnknown(m)
}

var xxx_messageInfo_UnlinkEventProto proto.InternalMessageInfo

func (m *UnlinkEventProto) GetPath() string {
	if m != nil && m.Path != nil {
		return *m.Path
	}
	return ""
}

func (m *UnlinkEventProto) GetTimestamp() int64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

type EventsListProto struct {
	Events               []*EventProto      `protobuf:"bytes,1,rep,name=events" json:"events,omitempty"`
	FirstTxid            *int64             `protobuf:"varint,2,req,name=firstTxid" json:"firstTxid,omitempty"`
	LastTxid             *int64             `protobuf:"varint,3,req,name=lastTxid" json:"lastTxid,omitempty"`
	SyncTxid             *int64             `protobuf:"varint,4,req,name=syncTxid" json:"syncTxid,omitempty"`
	Batch                []*EventBatchProto `protobuf:"bytes,5,rep,name=batch" json:"batch,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *EventsListProto) Reset()         { *m = EventsListProto{} }
func (m *EventsListProto) String() string { return proto.CompactTextString(m) }
func (*EventsListProto) ProtoMessage()    {}
func (*EventsListProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fe8964429f4f5e5, []int{9}
}

func (m *EventsListProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventsListProto.Unmarshal(m, b)
}
func (m *EventsListProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventsListProto.Marshal(b, m, deterministic)
}
func (m *EventsListProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventsListProto.Merge(m, src)
}
func (m *EventsListProto) XXX_Size() int {
	return xxx_messageInfo_EventsListProto.Size(m)
}
func (m *EventsListProto) XXX_DiscardUnknown() {
	xxx_messageInfo_EventsListProto.DiscardUnknown(m)
}

var xxx_messageInfo_EventsListProto proto.InternalMessageInfo

func (m *EventsListProto) GetEvents() []*EventProto {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *EventsListProto) GetFirstTxid() int64 {
	if m != nil && m.FirstTxid != nil {
		return *m.FirstTxid
	}
	return 0
}

func (m *EventsListProto) GetLastTxid() int64 {
	if m != nil && m.LastTxid != nil {
		return *m.LastTxid
	}
	return 0
}

func (m *EventsListProto) GetSyncTxid() int64 {
	if m != nil && m.SyncTxid != nil {
		return *m.SyncTxid
	}
	return 0
}

func (m *EventsListProto) GetBatch() []*EventBatchProto {
	if m != nil {
		return m.Batch
	}
	return nil
}

func init() {
	proto.RegisterEnum("hadoop.hdfs.EventType", EventType_name, EventType_value)
	proto.RegisterEnum("hadoop.hdfs.INodeType", INodeType_name, INodeType_value)
	proto.RegisterEnum("hadoop.hdfs.MetadataUpdateType", MetadataUpdateType_name, MetadataUpdateType_value)
	proto.RegisterType((*EventProto)(nil), "hadoop.hdfs.EventProto")
	proto.RegisterType((*EventBatchProto)(nil), "hadoop.hdfs.EventBatchProto")
	proto.RegisterType((*CreateEventProto)(nil), "hadoop.hdfs.CreateEventProto")
	proto.RegisterType((*CloseEventProto)(nil), "hadoop.hdfs.CloseEventProto")
	proto.RegisterType((*TruncateEventProto)(nil), "hadoop.hdfs.TruncateEventProto")
	proto.RegisterType((*AppendEventProto)(nil), "hadoop.hdfs.AppendEventProto")
	proto.RegisterType((*RenameEventProto)(nil), "hadoop.hdfs.RenameEventProto")
	proto.RegisterType((*MetadataUpdateEventProto)(nil), "hadoop.hdfs.MetadataUpdateEventProto")
	proto.RegisterType((*UnlinkEventProto)(nil), "hadoop.hdfs.UnlinkEventProto")
	proto.RegisterType((*EventsListProto)(nil), "hadoop.hdfs.EventsListProto")
}

func init() { proto.RegisterFile("inotify.proto", fileDescriptor_5fe8964429f4f5e5) }

var fileDescriptor_5fe8964429f4f5e5 = []byte{
	// 929 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0x4f, 0x6f, 0xe3, 0x44,
	0x14, 0x5f, 0x3b, 0x71, 0xdb, 0xbc, 0xb4, 0x5b, 0x6b, 0x28, 0x60, 0xa2, 0x15, 0x84, 0x08, 0xa4,
	0xa8, 0x12, 0x01, 0x15, 0x2e, 0xec, 0xcd, 0x4d, 0x5d, 0xc9, 0x22, 0x71, 0xa3, 0x89, 0xbb, 0x6c,
	0x4f, 0xd5, 0xac, 0x3d, 0x69, 0xac, 0x75, 0x6c, 0x6b, 0x66, 0xfa, 0x27, 0x7c, 0x06, 0x6e, 0x5c,
	0x38, 0x72, 0xe6, 0xb3, 0xf0, 0xa1, 0xd0, 0xcc, 0xb8, 0x76, 0xbc, 0x59, 0xd4, 0x1e, 0xb8, 0xf9,
	0xfd, 0xde, 0x6f, 0xde, 0xfb, 0x79, 0xde, 0x1f, 0x1b, 0x0e, 0x92, 0x2c, 0x17, 0xc9, 0x62, 0x3d,
	0x2a, 0x58, 0x2e, 0x72, 0xd4, 0x5d, 0x92, 0x38, 0xcf, 0x8b, 0xd1, 0x32, 0x5e, 0xf0, 0x5e, 0x87,
	0x44, 0xa9, 0xc6, 0x7b, 0xdd, 0x07, 0x22, 0x04, 0x2b, 0x0d, 0x90, 0x5e, 0xfd, 0x3c, 0x08, 0x01,
	0xbc, 0x3b, 0x9a, 0x89, 0x99, 0x3a, 0x7e, 0x0c, 0x6d, 0xb1, 0x2e, 0xa8, 0x63, 0xf4, 0xcd, 0xe1,
	0xcb, 0x93, 0xcf, 0x46, 0x1b, 0xd1, 0x46, 0x8a, 0x16, 0xae, 0x0b, 0x8a, 0x15, 0x07, 0xf5, 0x60,
	0x2f, 0xca, 0x33, 0x41, 0x33, 0xc1, 0x1d, 0xb3, 0x6f, 0x0e, 0xf7, 0x71, 0x65, 0x0f, 0xde, 0xc0,
	0xa1, 0xa2, 0x9f, 0x12, 0x11, 0x2d, 0x75, 0x68, 0x04, 0x6d, 0xf1, 0x90, 0xc4, 0x2a, 0x74, 0x0b,
	0xab, 0x67, 0xf4, 0x3d, 0xec, 0xd0, 0xbb, 0x32, 0x40, 0x6b, 0xd8, 0x3d, 0xf9, 0x7c, 0x3b, 0xa1,
	0x3a, 0x8c, 0x4b, 0xda, 0xe0, 0x8f, 0x16, 0xd8, 0x63, 0x46, 0x89, 0xa0, 0xcf, 0x14, 0xed, 0x07,
	0x79, 0x4c, 0x37, 0x44, 0x23, 0x68, 0x17, 0x44, 0x2c, 0x95, 0xe0, 0x0e, 0x56, 0xcf, 0xe8, 0x08,
	0xac, 0x48, 0x24, 0x2b, 0xea, 0xb4, 0x94, 0x34, 0x6d, 0xa0, 0x57, 0xd0, 0xc9, 0xef, 0x33, 0xca,
	0x02, 0xb2, 0xa2, 0x4e, 0x5b, 0xd1, 0x6b, 0x40, 0x7a, 0x6f, 0x58, 0x7e, 0x5b, 0x28, 0xaf, 0xa5,
	0xbd, 0x15, 0x80, 0x7e, 0x02, 0xab, 0xa0, 0x6c, 0xc5, 0x9d, 0x9d, 0xbe, 0x39, 0xec, 0x9e, 0x7c,
	0xd9, 0x90, 0x74, 0xce, 0x67, 0x94, 0xad, 0x12, 0xce, 0x93, 0x3c, 0xd3, 0x6f, 0xa7, 0xc9, 0xa8,
	0x0f, 0x5d, 0x46, 0x8b, 0x34, 0x89, 0x88, 0x48, 0xf2, 0xcc, 0xd9, 0xed, 0x1b, 0x43, 0x0b, 0x6f,
	0x42, 0xe8, 0x1b, 0x38, 0xe0, 0xeb, 0x55, 0x9a, 0x64, 0xef, 0x43, 0xc2, 0x6e, 0xa8, 0x70, 0xf6,
	0xfa, 0xc6, 0xb0, 0x83, 0x9b, 0xa0, 0x52, 0x7e, 0x47, 0xd9, 0x3d, 0x4b, 0x04, 0x75, 0x3a, 0x7d,
	0x63, 0xb8, 0x87, 0x6b, 0x00, 0x7d, 0x07, 0x76, 0x4c, 0x17, 0xe4, 0x36, 0x15, 0xa7, 0x69, 0x1e,
	0xbd, 0x9f, 0x27, 0xbf, 0x51, 0x07, 0xfa, 0xc6, 0xb0, 0xf5, 0xda, 0xf8, 0x01, 0x6f, 0xb9, 0xd0,
	0x00, 0xf6, 0x29, 0x23, 0xfc, 0x96, 0xd1, 0x71, 0x1e, 0xd3, 0xd8, 0xe9, 0xaa, 0x78, 0x0d, 0x6c,
	0x70, 0x0d, 0x87, 0xe3, 0x34, 0xe7, 0x9b, 0x35, 0x79, 0xbc, 0x67, 0x63, 0xe3, 0x9e, 0x7b, 0xb0,
	0xb7, 0x48, 0x52, 0xaa, 0x32, 0x9a, 0xea, 0xaa, 0x2b, 0x5b, 0x6a, 0x96, 0xb7, 0xce, 0x05, 0x59,
	0x15, 0x65, 0x1d, 0x6a, 0x60, 0xf0, 0x0e, 0x50, 0xc8, 0x6e, 0xb3, 0xa8, 0x59, 0xf7, 0xff, 0x37,
	0x87, 0x0f, 0xb6, 0x5b, 0x14, 0x34, 0x8b, 0x9f, 0xc8, 0xf0, 0x35, 0xec, 0x65, 0xf4, 0x5e, 0x5d,
	0x90, 0x63, 0xca, 0xcb, 0x78, 0x6d, 0x2d, 0x48, 0xca, 0x29, 0xae, 0xe0, 0xc1, 0x02, 0x6c, 0x4c,
	0x33, 0xb2, 0xda, 0x14, 0xeb, 0xc0, 0x2e, 0x67, 0xd1, 0xac, 0x8e, 0xf6, 0x68, 0x4a, 0xc9, 0x31,
	0xe5, 0x62, 0x56, 0xb7, 0x65, 0x65, 0x3f, 0x21, 0xf9, 0xef, 0x16, 0x38, 0x53, 0x2a, 0x48, 0x4c,
	0x04, 0xb9, 0x2c, 0xe2, 0xa7, 0x6f, 0xe7, 0xc7, 0x72, 0x52, 0x4c, 0x35, 0x29, 0x5f, 0x35, 0xda,
	0xb2, 0x19, 0x68, 0x63, 0x64, 0x8e, 0xc0, 0x5a, 0x95, 0xe3, 0x61, 0xc8, 0xf1, 0x50, 0x86, 0x44,
	0x89, 0x42, 0xdb, 0x1a, 0x55, 0xc6, 0x87, 0x2d, 0x6c, 0x6d, 0xb7, 0x70, 0x63, 0xac, 0x76, 0x54,
	0xfb, 0xfe, 0xd7, 0x58, 0xed, 0x6a, 0xef, 0x47, 0xc6, 0x4a, 0xb6, 0xfd, 0xb3, 0xc7, 0x6a, 0x04,
	0x6d, 0x12, 0xa5, 0xdc, 0xe9, 0xa8, 0x15, 0xd3, 0x6b, 0x1c, 0x72, 0xa3, 0xd4, 0xcb, 0x04, 0x5b,
	0xeb, 0x03, 0x8a, 0x27, 0x97, 0xd2, 0x83, 0x2b, 0x04, 0xe3, 0x0e, 0x7c, 0x64, 0x29, 0xbd, 0x95,
	0xae, 0x72, 0x29, 0x69, 0x9a, 0x9c, 0x4a, 0xfd, 0x84, 0xe9, 0x2a, 0xbf, 0xab, 0x66, 0xa4, 0x09,
	0x0e, 0xce, 0xc0, 0xbe, 0xcc, 0xe4, 0x94, 0x3e, 0x51, 0xa3, 0x46, 0xc9, 0xcd, 0x0f, 0x4b, 0xfe,
	0x8f, 0x51, 0x6e, 0x56, 0x3e, 0x49, 0x78, 0x19, 0xa5, 0xde, 0xa2, 0xc6, 0xb3, 0xb6, 0xa8, 0x4c,
	0xb1, 0x48, 0x18, 0x17, 0xa1, 0xdc, 0xc7, 0x65, 0x8a, 0x0a, 0x90, 0xfd, 0x98, 0x92, 0xd2, 0xa9,
	0x5b, 0xae, 0xb2, 0xa5, 0x8f, 0xaf, 0xb3, 0x48, 0xf9, 0xda, 0xda, 0xf7, 0x68, 0xa3, 0x13, 0xb0,
	0xde, 0xc9, 0x75, 0xef, 0x58, 0x4a, 0xc5, 0xab, 0x6d, 0x15, 0xf5, 0xd7, 0x00, 0x6b, 0xea, 0xf1,
	0xef, 0x06, 0x74, 0xaa, 0xef, 0x0a, 0xb2, 0x61, 0xdf, 0x7b, 0xe3, 0x05, 0xe1, 0xf5, 0x18, 0x7b,
	0x6e, 0xe8, 0xd9, 0x2f, 0xd0, 0x21, 0x74, 0x4b, 0x64, 0x72, 0x31, 0xf7, 0x6c, 0xa3, 0xa6, 0xb8,
	0xb3, 0x99, 0x17, 0x9c, 0xd9, 0x66, 0x8d, 0x60, 0x2f, 0x70, 0xa7, 0x9e, 0xdd, 0x42, 0x08, 0x5e,
	0x6a, 0x64, 0xea, 0x85, 0xee, 0x99, 0x1b, 0xba, 0x76, 0xbb, 0x66, 0x5d, 0x06, 0x13, 0x3f, 0xf8,
	0xc5, 0xb6, 0x6a, 0x56, 0x88, 0x2f, 0x83, 0xb1, 0x4c, 0xb7, 0x73, 0x7c, 0x0e, 0x9d, 0xea, 0x83,
	0x21, 0x73, 0xfb, 0xd7, 0xe1, 0xd5, 0xcc, 0xbb, 0x3e, 0xf7, 0x27, 0x52, 0xcc, 0x11, 0xd8, 0x25,
	0x70, 0xe6, 0x63, 0x6f, 0x1c, 0x5e, 0xe0, 0x2b, 0xdb, 0x90, 0x71, 0x4a, 0x74, 0x7e, 0x35, 0x55,
	0xb1, 0xcd, 0xe3, 0x3f, 0x0d, 0x40, 0xdb, 0xf3, 0x84, 0x3e, 0x81, 0x43, 0x29, 0x49, 0xb3, 0x43,
	0x7f, 0xea, 0xcd, 0xed, 0x17, 0xe8, 0x0b, 0xf8, 0xb4, 0x06, 0xb1, 0x37, 0x9b, 0xf8, 0x63, 0x37,
	0xf4, 0x2f, 0x02, 0xdb, 0x68, 0xf2, 0x2f, 0x7e, 0x0d, 0x3c, 0x6c, 0x9b, 0x4d, 0x70, 0xe6, 0xe1,
	0xe9, 0x5c, 0xbf, 0x72, 0x0d, 0xba, 0xe3, 0xc9, 0xdc, 0x6e, 0x4b, 0xb9, 0x35, 0xf6, 0xd6, 0x0d,
	0x43, 0x3c, 0xb7, 0xad, 0xd3, 0x9f, 0xe1, 0xdb, 0x9c, 0xdd, 0x8c, 0x48, 0x41, 0xa2, 0x25, 0x6d,
	0x94, 0x48, 0xfd, 0x0f, 0x44, 0x79, 0xf9, 0xc7, 0x70, 0x7a, 0xe0, 0xeb, 0x1f, 0x0b, 0x55, 0x2f,
	0xfe, 0x97, 0x61, 0xfc, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x8e, 0xb9, 0x7c, 0x72, 0x6b, 0x08, 0x00,
	0x00,
}

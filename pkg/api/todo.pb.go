// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/proto/todo.proto

package api

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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3ce579ad8bb256f, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type Data struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Body                 string   `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Data) Reset()         { *m = Data{} }
func (m *Data) String() string { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()    {}
func (*Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3ce579ad8bb256f, []int{1}
}

func (m *Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Data.Unmarshal(m, b)
}
func (m *Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Data.Marshal(b, m, deterministic)
}
func (m *Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Data.Merge(m, src)
}
func (m *Data) XXX_Size() int {
	return xxx_messageInfo_Data.Size(m)
}
func (m *Data) XXX_DiscardUnknown() {
	xxx_messageInfo_Data.DiscardUnknown(m)
}

var xxx_messageInfo_Data proto.InternalMessageInfo

func (m *Data) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Data) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

type TaskId struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskId) Reset()         { *m = TaskId{} }
func (m *TaskId) String() string { return proto.CompactTextString(m) }
func (*TaskId) ProtoMessage()    {}
func (*TaskId) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3ce579ad8bb256f, []int{2}
}

func (m *TaskId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskId.Unmarshal(m, b)
}
func (m *TaskId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskId.Marshal(b, m, deterministic)
}
func (m *TaskId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskId.Merge(m, src)
}
func (m *TaskId) XXX_Size() int {
	return xxx_messageInfo_TaskId.Size(m)
}
func (m *TaskId) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskId.DiscardUnknown(m)
}

var xxx_messageInfo_TaskId proto.InternalMessageInfo

func (m *TaskId) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Task struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Body                 string   `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	Done                 bool     `protobuf:"varint,4,opt,name=done,proto3" json:"done,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3ce579ad8bb256f, []int{3}
}

func (m *Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Task.Unmarshal(m, b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Task.Marshal(b, m, deterministic)
}
func (m *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(m, src)
}
func (m *Task) XXX_Size() int {
	return xxx_messageInfo_Task.Size(m)
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

func (m *Task) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Task) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Task) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *Task) GetDone() bool {
	if m != nil {
		return m.Done
	}
	return false
}

type TaskList struct {
	Task                 []*Task  `protobuf:"bytes,1,rep,name=task,proto3" json:"task,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskList) Reset()         { *m = TaskList{} }
func (m *TaskList) String() string { return proto.CompactTextString(m) }
func (*TaskList) ProtoMessage()    {}
func (*TaskList) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3ce579ad8bb256f, []int{4}
}

func (m *TaskList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskList.Unmarshal(m, b)
}
func (m *TaskList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskList.Marshal(b, m, deterministic)
}
func (m *TaskList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskList.Merge(m, src)
}
func (m *TaskList) XXX_Size() int {
	return xxx_messageInfo_TaskList.Size(m)
}
func (m *TaskList) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskList.DiscardUnknown(m)
}

var xxx_messageInfo_TaskList proto.InternalMessageInfo

func (m *TaskList) GetTask() []*Task {
	if m != nil {
		return m.Task
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "api.Empty")
	proto.RegisterType((*Data)(nil), "api.Data")
	proto.RegisterType((*TaskId)(nil), "api.TaskId")
	proto.RegisterType((*Task)(nil), "api.Task")
	proto.RegisterType((*TaskList)(nil), "api.TaskList")
}

func init() { proto.RegisterFile("api/proto/todo.proto", fileDescriptor_b3ce579ad8bb256f) }

var fileDescriptor_b3ce579ad8bb256f = []byte{
	// 286 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x97, 0x36, 0xeb, 0xb6, 0x37, 0xf4, 0xf0, 0xd8, 0xa1, 0x0c, 0x84, 0x12, 0x61, 0xd4,
	0x4b, 0x27, 0xf3, 0x13, 0xa8, 0x15, 0x19, 0xe8, 0xa5, 0xd4, 0x8b, 0xb7, 0x8c, 0xe4, 0x10, 0x3a,
	0x97, 0xd0, 0xbe, 0xcb, 0x3e, 0xae, 0xdf, 0x44, 0x92, 0xc2, 0xea, 0x86, 0x78, 0x7b, 0xef, 0xf7,
	0xfe, 0xf9, 0x85, 0x97, 0xc0, 0x42, 0x3a, 0xb3, 0x76, 0xad, 0x25, 0xbb, 0x26, 0xab, 0x6c, 0x11,
	0x4a, 0x8c, 0xa5, 0x33, 0x62, 0x02, 0xe3, 0x97, 0x2f, 0x47, 0x47, 0x71, 0x0f, 0xbc, 0x94, 0x24,
	0x71, 0x01, 0x63, 0x32, 0xb4, 0xd7, 0x29, 0xcb, 0x58, 0x3e, 0xab, 0xfa, 0x06, 0x11, 0xf8, 0xce,
	0xaa, 0x63, 0x1a, 0x05, 0x18, 0x6a, 0x91, 0x42, 0x52, 0xcb, 0xae, 0xd9, 0x2a, 0xbc, 0x86, 0xc8,
	0xa8, 0x70, 0x60, 0x5c, 0x45, 0x46, 0x89, 0x1a, 0xb8, 0x9f, 0x5c, 0xf2, 0xc1, 0x1d, 0xfd, 0xe5,
	0x8e, 0x07, 0xb7, 0x67, 0xca, 0x1e, 0x74, 0xca, 0x33, 0x96, 0x4f, 0xab, 0x50, 0x8b, 0x3b, 0x98,
	0x7a, 0xeb, 0x9b, 0xe9, 0x08, 0x6f, 0x80, 0x93, 0xec, 0x9a, 0x94, 0x65, 0x71, 0x3e, 0xdf, 0xcc,
	0x0a, 0xe9, 0x4c, 0xe1, 0x87, 0x55, 0xc0, 0x9b, 0x6f, 0x06, 0xbc, 0xb6, 0xa5, 0xc5, 0x0c, 0x92,
	0xe7, 0x56, 0x4b, 0xd2, 0xd8, 0x67, 0xfc, 0x8a, 0xcb, 0x21, 0x2e, 0x46, 0x98, 0x01, 0xaf, 0xb4,
	0x54, 0x38, 0x3f, 0xc1, 0xad, 0xba, 0x4c, 0x24, 0x1f, 0x4e, 0x0d, 0x0e, 0x8f, 0xcf, 0x13, 0xb7,
	0x90, 0x94, 0x7a, 0xaf, 0x49, 0x9f, 0x5b, 0x20, 0x34, 0xfd, 0xf3, 0x8e, 0x70, 0x05, 0xf0, 0x2e,
	0xdb, 0xe6, 0xb1, 0x2b, 0xed, 0x41, 0xff, 0x73, 0xdd, 0x0a, 0x26, 0xaf, 0x9a, 0xc2, 0x96, 0xbf,
	0x04, 0xcb, 0xab, 0x53, 0xc6, 0x8f, 0xc4, 0xe8, 0x29, 0xfe, 0x64, 0xc5, 0x2e, 0x09, 0x5f, 0xf9,
	0xf0, 0x13, 0x00, 0x00, 0xff, 0xff, 0x41, 0x55, 0x68, 0x65, 0xe2, 0x01, 0x00, 0x00,
}

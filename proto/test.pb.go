// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

package test

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

type Test1 struct {
	A                    int64    `protobuf:"zigzag64,1,opt,name=a,proto3" json:"a,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Test1) Reset()         { *m = Test1{} }
func (m *Test1) String() string { return proto.CompactTextString(m) }
func (*Test1) ProtoMessage()    {}
func (*Test1) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{0}
}

func (m *Test1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Test1.Unmarshal(m, b)
}
func (m *Test1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Test1.Marshal(b, m, deterministic)
}
func (m *Test1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Test1.Merge(m, src)
}
func (m *Test1) XXX_Size() int {
	return xxx_messageInfo_Test1.Size(m)
}
func (m *Test1) XXX_DiscardUnknown() {
	xxx_messageInfo_Test1.DiscardUnknown(m)
}

var xxx_messageInfo_Test1 proto.InternalMessageInfo

func (m *Test1) GetA() int64 {
	if m != nil {
		return m.A
	}
	return 0
}

func init() {
	proto.RegisterType((*Test1)(nil), "Test1")
}

func init() { proto.RegisterFile("test.proto", fileDescriptor_c161fcfdc0c3ff1e) }

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messages.proto

package cosipbft

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

type Prepare struct {
	Proposal             *any.Any `protobuf:"bytes,1,opt,name=proposal,proto3" json:"proposal,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Prepare) Reset()         { *m = Prepare{} }
func (m *Prepare) String() string { return proto.CompactTextString(m) }
func (*Prepare) ProtoMessage()    {}
func (*Prepare) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{0}
}

func (m *Prepare) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Prepare.Unmarshal(m, b)
}
func (m *Prepare) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Prepare.Marshal(b, m, deterministic)
}
func (m *Prepare) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Prepare.Merge(m, src)
}
func (m *Prepare) XXX_Size() int {
	return xxx_messageInfo_Prepare.Size(m)
}
func (m *Prepare) XXX_DiscardUnknown() {
	xxx_messageInfo_Prepare.DiscardUnknown(m)
}

var xxx_messageInfo_Prepare proto.InternalMessageInfo

func (m *Prepare) GetProposal() *any.Any {
	if m != nil {
		return m.Proposal
	}
	return nil
}

type ForwardLinkProto struct {
	From                 []byte   `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To                   []byte   `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Prepare              *any.Any `protobuf:"bytes,3,opt,name=prepare,proto3" json:"prepare,omitempty"`
	Commit               *any.Any `protobuf:"bytes,4,opt,name=commit,proto3" json:"commit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ForwardLinkProto) Reset()         { *m = ForwardLinkProto{} }
func (m *ForwardLinkProto) String() string { return proto.CompactTextString(m) }
func (*ForwardLinkProto) ProtoMessage()    {}
func (*ForwardLinkProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{1}
}

func (m *ForwardLinkProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ForwardLinkProto.Unmarshal(m, b)
}
func (m *ForwardLinkProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ForwardLinkProto.Marshal(b, m, deterministic)
}
func (m *ForwardLinkProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ForwardLinkProto.Merge(m, src)
}
func (m *ForwardLinkProto) XXX_Size() int {
	return xxx_messageInfo_ForwardLinkProto.Size(m)
}
func (m *ForwardLinkProto) XXX_DiscardUnknown() {
	xxx_messageInfo_ForwardLinkProto.DiscardUnknown(m)
}

var xxx_messageInfo_ForwardLinkProto proto.InternalMessageInfo

func (m *ForwardLinkProto) GetFrom() []byte {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *ForwardLinkProto) GetTo() []byte {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *ForwardLinkProto) GetPrepare() *any.Any {
	if m != nil {
		return m.Prepare
	}
	return nil
}

func (m *ForwardLinkProto) GetCommit() *any.Any {
	if m != nil {
		return m.Commit
	}
	return nil
}

type Commit struct {
	ForwardLink          *ForwardLinkProto `protobuf:"bytes,1,opt,name=forwardLink,proto3" json:"forwardLink,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Commit) Reset()         { *m = Commit{} }
func (m *Commit) String() string { return proto.CompactTextString(m) }
func (*Commit) ProtoMessage()    {}
func (*Commit) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{2}
}

func (m *Commit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Commit.Unmarshal(m, b)
}
func (m *Commit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Commit.Marshal(b, m, deterministic)
}
func (m *Commit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Commit.Merge(m, src)
}
func (m *Commit) XXX_Size() int {
	return xxx_messageInfo_Commit.Size(m)
}
func (m *Commit) XXX_DiscardUnknown() {
	xxx_messageInfo_Commit.DiscardUnknown(m)
}

var xxx_messageInfo_Commit proto.InternalMessageInfo

func (m *Commit) GetForwardLink() *ForwardLinkProto {
	if m != nil {
		return m.ForwardLink
	}
	return nil
}

func init() {
	proto.RegisterType((*Prepare)(nil), "cosipbft.Prepare")
	proto.RegisterType((*ForwardLinkProto)(nil), "cosipbft.ForwardLinkProto")
	proto.RegisterType((*Commit)(nil), "cosipbft.Commit")
}

func init() { proto.RegisterFile("messages.proto", fileDescriptor_4dc296cbfe5ffcd5) }

var fileDescriptor_4dc296cbfe5ffcd5 = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x4f, 0x4b, 0x87, 0x30,
	0x18, 0xc7, 0xd1, 0x44, 0xe5, 0x31, 0x24, 0x46, 0x87, 0xe5, 0x29, 0x3c, 0x75, 0x88, 0x19, 0x75,
	0xac, 0x4b, 0x04, 0x9e, 0x3a, 0x88, 0xef, 0x60, 0xda, 0x26, 0x92, 0xf3, 0x19, 0xdb, 0x22, 0x7c,
	0x1f, 0xbd, 0xe0, 0x60, 0x6a, 0xfd, 0xf8, 0x1d, 0xbc, 0x3d, 0x7c, 0xf9, 0x7c, 0xff, 0x6c, 0x90,
	0x2b, 0x61, 0x2d, 0x1f, 0x84, 0x65, 0xda, 0xa0, 0x43, 0x92, 0xf6, 0x68, 0x47, 0xdd, 0x49, 0x57,
	0xdc, 0x0c, 0x88, 0xc3, 0x24, 0x2a, 0xaf, 0x77, 0x5f, 0xb2, 0xe2, 0xf3, 0xb2, 0x42, 0xe5, 0x33,
	0x24, 0x8d, 0x11, 0x9a, 0x1b, 0x41, 0x1e, 0x20, 0xd5, 0x06, 0x35, 0x5a, 0x3e, 0xd1, 0xe0, 0x36,
	0xb8, 0xcb, 0x1e, 0xaf, 0xd9, 0x6a, 0x64, 0xbb, 0x91, 0xbd, 0xce, 0x4b, 0xfb, 0x47, 0x95, 0x3f,
	0x01, 0x5c, 0xd5, 0x68, 0xbe, 0xb9, 0xf9, 0x78, 0x1f, 0xe7, 0xcf, 0xc6, 0xd7, 0x12, 0x88, 0xa4,
	0x41, 0xe5, 0x23, 0x2e, 0x5b, 0x7f, 0x93, 0x1c, 0x42, 0x87, 0x34, 0xf4, 0x4a, 0xe8, 0x90, 0x30,
	0x48, 0xf4, 0xda, 0x4a, 0x2f, 0x0e, 0x9a, 0x76, 0x88, 0xdc, 0x43, 0xdc, 0xa3, 0x52, 0xa3, 0xa3,
	0xd1, 0x01, 0xbe, 0x31, 0x65, 0x0d, 0xf1, 0x9b, 0xbf, 0xc8, 0x0b, 0x64, 0xf2, 0x7f, 0xdf, 0xf6,
	0xaa, 0x82, 0xed, 0x1f, 0xc3, 0xce, 0xc7, 0xb7, 0xa7, 0x78, 0x17, 0xfb, 0xf4, 0xa7, 0xdf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xfd, 0xb8, 0xc3, 0xd5, 0x59, 0x01, 0x00, 0x00,
}

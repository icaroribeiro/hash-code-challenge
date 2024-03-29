// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/proto/entities/promotion.proto

package entities

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

type Promotion struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Code                 string   `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	MaxDiscountPct       float32  `protobuf:"fixed32,5,opt,name=max_discount_pct,json=maxDiscountPct,proto3" json:"max_discount_pct,omitempty"`
	Products             []string `protobuf:"bytes,6,rep,name=products,proto3" json:"products,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Promotion) Reset()         { *m = Promotion{} }
func (m *Promotion) String() string { return proto.CompactTextString(m) }
func (*Promotion) ProtoMessage()    {}
func (*Promotion) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c452e958b5c66b7, []int{0}
}

func (m *Promotion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Promotion.Unmarshal(m, b)
}
func (m *Promotion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Promotion.Marshal(b, m, deterministic)
}
func (m *Promotion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Promotion.Merge(m, src)
}
func (m *Promotion) XXX_Size() int {
	return xxx_messageInfo_Promotion.Size(m)
}
func (m *Promotion) XXX_DiscardUnknown() {
	xxx_messageInfo_Promotion.DiscardUnknown(m)
}

var xxx_messageInfo_Promotion proto.InternalMessageInfo

func (m *Promotion) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Promotion) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Promotion) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Promotion) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Promotion) GetMaxDiscountPct() float32 {
	if m != nil {
		return m.MaxDiscountPct
	}
	return 0
}

func (m *Promotion) GetProducts() []string {
	if m != nil {
		return m.Products
	}
	return nil
}

func init() {
	proto.RegisterType((*Promotion)(nil), "entities.Promotion")
}

func init() {
	proto.RegisterFile("github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/proto/entities/promotion.proto", fileDescriptor_2c452e958b5c66b7)
}

var fileDescriptor_2c452e958b5c66b7 = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x90, 0xbd, 0x6a, 0xec, 0x30,
	0x10, 0x46, 0xb1, 0xf7, 0x87, 0xb5, 0x2e, 0x2c, 0x17, 0x91, 0x42, 0xa4, 0x32, 0xa9, 0xdc, 0x78,
	0x55, 0xe4, 0x0d, 0x42, 0xea, 0x60, 0x5c, 0xa6, 0x31, 0xf2, 0x48, 0xd8, 0x43, 0x6c, 0x8d, 0x91,
	0xc6, 0xb0, 0x2f, 0x95, 0x77, 0x0c, 0x56, 0xb2, 0x4b, 0xfa, 0x74, 0xfa, 0xce, 0x01, 0x0d, 0x1c,
	0xd1, 0x0d, 0xc8, 0xe3, 0xda, 0x5f, 0x80, 0x66, 0x8d, 0x60, 0x02, 0x05, 0xec, 0x1d, 0x06, 0xd2,
	0xa3, 0x89, 0x63, 0x0d, 0x64, 0x5d, 0x0d, 0xa3, 0x99, 0x26, 0xe7, 0x07, 0xa7, 0x7b, 0x03, 0x1f,
	0xb5, 0xf3, 0x56, 0x0f, 0xa4, 0xd1, 0xb3, 0x0b, 0xde, 0x4c, 0x7a, 0x09, 0xc4, 0xa4, 0x9d, 0x67,
	0x64, 0x74, 0x71, 0x9b, 0x33, 0x31, 0x92, 0xbf, 0x24, 0x21, 0x4f, 0x37, 0xf3, 0xf4, 0x99, 0x89,
	0xa2, 0xb9, 0x59, 0x79, 0x16, 0x39, 0x5a, 0x95, 0x95, 0x59, 0x55, 0xb4, 0x39, 0x5a, 0x29, 0xc5,
	0x7e, 0xbb, 0xa6, 0xf2, 0x44, 0xd2, 0x5b, 0x3e, 0x88, 0x03, 0x23, 0x4f, 0x4e, 0xed, 0x12, 0xfc,
	0x1e, 0xb2, 0x14, 0xff, 0xac, 0x8b, 0x10, 0x70, 0xd9, 0x3e, 0x52, 0xfb, 0xe4, 0x7e, 0x23, 0x59,
	0x89, 0xff, 0xb3, 0xb9, 0x76, 0x16, 0x23, 0xd0, 0xea, 0xb9, 0x5b, 0x80, 0xd5, 0xa1, 0xcc, 0xaa,
	0xbc, 0x3d, 0xcf, 0xe6, 0xfa, 0xfa, 0x83, 0x1b, 0x60, 0xf9, 0x28, 0x4e, 0x4b, 0x20, 0xbb, 0x02,
	0x47, 0x75, 0x2c, 0x77, 0x55, 0xd1, 0xde, 0xf7, 0x4b, 0xf3, 0xfe, 0xf6, 0xe7, 0x38, 0x43, 0x58,
	0xe0, 0xde, 0xa6, 0x3f, 0xa6, 0x24, 0xcf, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x2c, 0x5f, 0xcf,
	0x30, 0x75, 0x01, 0x00, 0x00,
}

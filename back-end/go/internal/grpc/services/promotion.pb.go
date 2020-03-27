// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/proto/services/promotion.proto

package services

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	entities "github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/grpc/entities"
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

type GetAllPromotionsResponse struct {
	Promotions           []*entities.Promotion `protobuf:"bytes,1,rep,name=promotions,proto3" json:"promotions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *GetAllPromotionsResponse) Reset()         { *m = GetAllPromotionsResponse{} }
func (m *GetAllPromotionsResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllPromotionsResponse) ProtoMessage()    {}
func (*GetAllPromotionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7930af44f91217f, []int{0}
}

func (m *GetAllPromotionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllPromotionsResponse.Unmarshal(m, b)
}
func (m *GetAllPromotionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllPromotionsResponse.Marshal(b, m, deterministic)
}
func (m *GetAllPromotionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllPromotionsResponse.Merge(m, src)
}
func (m *GetAllPromotionsResponse) XXX_Size() int {
	return xxx_messageInfo_GetAllPromotionsResponse.Size(m)
}
func (m *GetAllPromotionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllPromotionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllPromotionsResponse proto.InternalMessageInfo

func (m *GetAllPromotionsResponse) GetPromotions() []*entities.Promotion {
	if m != nil {
		return m.Promotions
	}
	return nil
}

type CreatePromotionRequest struct {
	Promotion            *entities.Promotion `protobuf:"bytes,1,opt,name=promotion,proto3" json:"promotion,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *CreatePromotionRequest) Reset()         { *m = CreatePromotionRequest{} }
func (m *CreatePromotionRequest) String() string { return proto.CompactTextString(m) }
func (*CreatePromotionRequest) ProtoMessage()    {}
func (*CreatePromotionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7930af44f91217f, []int{1}
}

func (m *CreatePromotionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePromotionRequest.Unmarshal(m, b)
}
func (m *CreatePromotionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePromotionRequest.Marshal(b, m, deterministic)
}
func (m *CreatePromotionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePromotionRequest.Merge(m, src)
}
func (m *CreatePromotionRequest) XXX_Size() int {
	return xxx_messageInfo_CreatePromotionRequest.Size(m)
}
func (m *CreatePromotionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePromotionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePromotionRequest proto.InternalMessageInfo

func (m *CreatePromotionRequest) GetPromotion() *entities.Promotion {
	if m != nil {
		return m.Promotion
	}
	return nil
}

type GetPromotionRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPromotionRequest) Reset()         { *m = GetPromotionRequest{} }
func (m *GetPromotionRequest) String() string { return proto.CompactTextString(m) }
func (*GetPromotionRequest) ProtoMessage()    {}
func (*GetPromotionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7930af44f91217f, []int{2}
}

func (m *GetPromotionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPromotionRequest.Unmarshal(m, b)
}
func (m *GetPromotionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPromotionRequest.Marshal(b, m, deterministic)
}
func (m *GetPromotionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPromotionRequest.Merge(m, src)
}
func (m *GetPromotionRequest) XXX_Size() int {
	return xxx_messageInfo_GetPromotionRequest.Size(m)
}
func (m *GetPromotionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPromotionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPromotionRequest proto.InternalMessageInfo

func (m *GetPromotionRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type UpdatePromotionRequest struct {
	Id                   string              `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Promotion            *entities.Promotion `protobuf:"bytes,2,opt,name=promotion,proto3" json:"promotion,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *UpdatePromotionRequest) Reset()         { *m = UpdatePromotionRequest{} }
func (m *UpdatePromotionRequest) String() string { return proto.CompactTextString(m) }
func (*UpdatePromotionRequest) ProtoMessage()    {}
func (*UpdatePromotionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7930af44f91217f, []int{3}
}

func (m *UpdatePromotionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePromotionRequest.Unmarshal(m, b)
}
func (m *UpdatePromotionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePromotionRequest.Marshal(b, m, deterministic)
}
func (m *UpdatePromotionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePromotionRequest.Merge(m, src)
}
func (m *UpdatePromotionRequest) XXX_Size() int {
	return xxx_messageInfo_UpdatePromotionRequest.Size(m)
}
func (m *UpdatePromotionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePromotionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePromotionRequest proto.InternalMessageInfo

func (m *UpdatePromotionRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdatePromotionRequest) GetPromotion() *entities.Promotion {
	if m != nil {
		return m.Promotion
	}
	return nil
}

type DeletePromotionRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletePromotionRequest) Reset()         { *m = DeletePromotionRequest{} }
func (m *DeletePromotionRequest) String() string { return proto.CompactTextString(m) }
func (*DeletePromotionRequest) ProtoMessage()    {}
func (*DeletePromotionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7930af44f91217f, []int{4}
}

func (m *DeletePromotionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePromotionRequest.Unmarshal(m, b)
}
func (m *DeletePromotionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePromotionRequest.Marshal(b, m, deterministic)
}
func (m *DeletePromotionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePromotionRequest.Merge(m, src)
}
func (m *DeletePromotionRequest) XXX_Size() int {
	return xxx_messageInfo_DeletePromotionRequest.Size(m)
}
func (m *DeletePromotionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePromotionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePromotionRequest proto.InternalMessageInfo

func (m *DeletePromotionRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type IncludeProductToPromotionRequest struct {
	Id                   string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Product              *entities.Product `protobuf:"bytes,2,opt,name=product,proto3" json:"product,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *IncludeProductToPromotionRequest) Reset()         { *m = IncludeProductToPromotionRequest{} }
func (m *IncludeProductToPromotionRequest) String() string { return proto.CompactTextString(m) }
func (*IncludeProductToPromotionRequest) ProtoMessage()    {}
func (*IncludeProductToPromotionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7930af44f91217f, []int{5}
}

func (m *IncludeProductToPromotionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IncludeProductToPromotionRequest.Unmarshal(m, b)
}
func (m *IncludeProductToPromotionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IncludeProductToPromotionRequest.Marshal(b, m, deterministic)
}
func (m *IncludeProductToPromotionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IncludeProductToPromotionRequest.Merge(m, src)
}
func (m *IncludeProductToPromotionRequest) XXX_Size() int {
	return xxx_messageInfo_IncludeProductToPromotionRequest.Size(m)
}
func (m *IncludeProductToPromotionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IncludeProductToPromotionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IncludeProductToPromotionRequest proto.InternalMessageInfo

func (m *IncludeProductToPromotionRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *IncludeProductToPromotionRequest) GetProduct() *entities.Product {
	if m != nil {
		return m.Product
	}
	return nil
}

type ExcludeProductFromPromotionRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ProductId            string   `protobuf:"bytes,2,opt,name=productId,proto3" json:"productId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExcludeProductFromPromotionRequest) Reset()         { *m = ExcludeProductFromPromotionRequest{} }
func (m *ExcludeProductFromPromotionRequest) String() string { return proto.CompactTextString(m) }
func (*ExcludeProductFromPromotionRequest) ProtoMessage()    {}
func (*ExcludeProductFromPromotionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7930af44f91217f, []int{6}
}

func (m *ExcludeProductFromPromotionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExcludeProductFromPromotionRequest.Unmarshal(m, b)
}
func (m *ExcludeProductFromPromotionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExcludeProductFromPromotionRequest.Marshal(b, m, deterministic)
}
func (m *ExcludeProductFromPromotionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExcludeProductFromPromotionRequest.Merge(m, src)
}
func (m *ExcludeProductFromPromotionRequest) XXX_Size() int {
	return xxx_messageInfo_ExcludeProductFromPromotionRequest.Size(m)
}
func (m *ExcludeProductFromPromotionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExcludeProductFromPromotionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExcludeProductFromPromotionRequest proto.InternalMessageInfo

func (m *ExcludeProductFromPromotionRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ExcludeProductFromPromotionRequest) GetProductId() string {
	if m != nil {
		return m.ProductId
	}
	return ""
}

func init() {
	proto.RegisterType((*GetAllPromotionsResponse)(nil), "services.GetAllPromotionsResponse")
	proto.RegisterType((*CreatePromotionRequest)(nil), "services.CreatePromotionRequest")
	proto.RegisterType((*GetPromotionRequest)(nil), "services.GetPromotionRequest")
	proto.RegisterType((*UpdatePromotionRequest)(nil), "services.UpdatePromotionRequest")
	proto.RegisterType((*DeletePromotionRequest)(nil), "services.DeletePromotionRequest")
	proto.RegisterType((*IncludeProductToPromotionRequest)(nil), "services.IncludeProductToPromotionRequest")
	proto.RegisterType((*ExcludeProductFromPromotionRequest)(nil), "services.ExcludeProductFromPromotionRequest")
}

func init() {
	proto.RegisterFile("github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/proto/services/promotion.proto", fileDescriptor_f7930af44f91217f)
}

var fileDescriptor_f7930af44f91217f = []byte{
	// 440 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0x5f, 0x6b, 0xd4, 0x40,
	0x14, 0xc5, 0xd9, 0x2d, 0xa8, 0x7b, 0x15, 0x5b, 0x53, 0x58, 0x96, 0xa8, 0xb0, 0x0c, 0x08, 0x05,
	0xd9, 0x19, 0x6c, 0x3f, 0x81, 0xd5, 0xda, 0x16, 0xa1, 0x96, 0xa8, 0x2f, 0x8a, 0x94, 0x64, 0xe6,
	0x9a, 0x1d, 0x9c, 0x9d, 0x89, 0x33, 0x13, 0xd1, 0x77, 0x3f, 0xb8, 0x24, 0x69, 0xfe, 0x6c, 0x9a,
	0x65, 0x0b, 0xea, 0xe3, 0xee, 0x3d, 0xf3, 0x3b, 0x27, 0x99, 0x73, 0x03, 0x57, 0xa9, 0xf4, 0xcb,
	0x3c, 0xa1, 0xdc, 0xac, 0x98, 0xe4, 0xb1, 0x35, 0x56, 0x26, 0x28, 0xad, 0x61, 0xcb, 0xd8, 0x2d,
	0x17, 0xdc, 0x08, 0x5c, 0xf0, 0x65, 0xac, 0x14, 0xea, 0x14, 0x59, 0x12, 0xf3, 0x6f, 0x0b, 0xd4,
	0x82, 0xa5, 0x86, 0x49, 0xed, 0xd1, 0xea, 0x58, 0xb1, 0xcc, 0x1a, 0x6f, 0x98, 0x43, 0xfb, 0x43,
	0x72, 0x74, 0xc5, 0xcf, 0x95, 0xf1, 0xd2, 0x68, 0x5a, 0x0e, 0x82, 0x7b, 0xf5, 0x24, 0x7c, 0x9c,
	0x1a, 0x93, 0x2a, 0xac, 0x0e, 0x24, 0xf9, 0x57, 0x86, 0xab, 0xcc, 0xff, 0xaa, 0x64, 0xe1, 0xbf,
	0xca, 0x81, 0xda, 0x4b, 0x2f, 0x6f, 0xe6, 0x08, 0xbf, 0xfc, 0x07, 0x03, 0x91, 0x73, 0x5f, 0xe1,
	0xc9, 0x3b, 0x98, 0x9d, 0xa2, 0x7f, 0xa9, 0xd4, 0x65, 0xed, 0xeb, 0x22, 0x74, 0x99, 0xd1, 0x0e,
	0x83, 0x23, 0x80, 0x26, 0x8d, 0x9b, 0x8d, 0xe6, 0x3b, 0x07, 0xf7, 0x0f, 0xf7, 0x69, 0x0d, 0xa2,
	0xcd, 0x89, 0xa8, 0x23, 0x23, 0x6f, 0x61, 0xfa, 0xca, 0x62, 0xec, 0xb1, 0x1d, 0xe3, 0xf7, 0x1c,
	0x9d, 0x0f, 0x5e, 0xc0, 0xa4, 0xd1, 0xcd, 0x46, 0xf3, 0xd1, 0x26, 0x5a, 0xab, 0x22, 0xcf, 0x60,
	0xff, 0x14, 0xfd, 0x0d, 0xd2, 0x43, 0x18, 0x4b, 0x51, 0x22, 0x26, 0xd1, 0x58, 0x0a, 0xf2, 0x19,
	0xa6, 0x1f, 0x33, 0x31, 0xe4, 0xd9, 0x53, 0xae, 0x67, 0x18, 0xdf, 0x2a, 0xc3, 0x01, 0x4c, 0x5f,
	0xa3, 0xc2, 0xed, 0x70, 0x72, 0x05, 0xf3, 0x73, 0xcd, 0x55, 0x2e, 0x0a, 0x69, 0xf1, 0x8e, 0x3f,
	0x98, 0xad, 0x81, 0x9e, 0xc3, 0xdd, 0xeb, 0x0b, 0xb9, 0x8e, 0xf3, 0x68, 0x2d, 0x4e, 0x31, 0x88,
	0x6a, 0x05, 0x89, 0x80, 0x9c, 0xfc, 0xec, 0x1a, 0xbc, 0xb1, 0x66, 0xb5, 0xd5, 0xe2, 0x49, 0xf9,
	0xcc, 0x85, 0xfc, 0x5c, 0x94, 0x26, 0x93, 0xa8, 0xfd, 0xe3, 0xf0, 0xf7, 0x0e, 0xec, 0x35, 0x88,
	0xf7, 0x55, 0xe7, 0x83, 0x0b, 0xd8, 0xeb, 0xb7, 0x22, 0x98, 0xd2, 0x6a, 0x0f, 0x68, 0xbd, 0x07,
	0xf4, 0xa4, 0xd8, 0x83, 0x90, 0xd0, 0x7a, 0x53, 0xe8, 0xc6, 0x26, 0x9d, 0xc1, 0x6e, 0xaf, 0x14,
	0xc1, 0xbc, 0x3d, 0x36, 0xdc, 0x97, 0x70, 0xe8, 0x62, 0x82, 0x63, 0x78, 0xd0, 0x6d, 0x44, 0xf0,
	0x74, 0xcd, 0xfd, 0x76, 0x8c, 0x33, 0xd8, 0xed, 0xd5, 0xa5, 0x9b, 0x66, 0xb8, 0x49, 0x1b, 0x49,
	0xbd, 0x6e, 0x74, 0x49, 0xc3, 0xb5, 0x19, 0x24, 0x1d, 0x5f, 0x7e, 0xba, 0xf8, 0xeb, 0x45, 0x4f,
	0x6d, 0xc6, 0x9b, 0x0f, 0x5a, 0x72, 0xa7, 0xbc, 0xa7, 0xa3, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x41, 0x32, 0x1c, 0xb4, 0x2a, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PromotionServiceClient is the client API for PromotionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PromotionServiceClient interface {
	// Get (read) the list of all promotions.
	GetAllPromotions(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetAllPromotionsResponse, error)
	// Create a new promotion.
	CreatePromotion(ctx context.Context, in *CreatePromotionRequest, opts ...grpc.CallOption) (*entities.Promotion, error)
	// Get (read) a specific promotion by its id.
	GetPromotion(ctx context.Context, in *GetPromotionRequest, opts ...grpc.CallOption) (*entities.Promotion, error)
	// Update a specific promotion by its id.
	UpdatePromotion(ctx context.Context, in *UpdatePromotionRequest, opts ...grpc.CallOption) (*entities.Promotion, error)
	// Delete a specific promotion by its id.
	DeletePromotion(ctx context.Context, in *DeletePromotionRequest, opts ...grpc.CallOption) (*entities.Promotion, error)
}

type promotionServiceClient struct {
	cc *grpc.ClientConn
}

func NewPromotionServiceClient(cc *grpc.ClientConn) PromotionServiceClient {
	return &promotionServiceClient{cc}
}

func (c *promotionServiceClient) GetAllPromotions(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetAllPromotionsResponse, error) {
	out := new(GetAllPromotionsResponse)
	err := c.cc.Invoke(ctx, "/services.PromotionService/GetAllPromotions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *promotionServiceClient) CreatePromotion(ctx context.Context, in *CreatePromotionRequest, opts ...grpc.CallOption) (*entities.Promotion, error) {
	out := new(entities.Promotion)
	err := c.cc.Invoke(ctx, "/services.PromotionService/CreatePromotion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *promotionServiceClient) GetPromotion(ctx context.Context, in *GetPromotionRequest, opts ...grpc.CallOption) (*entities.Promotion, error) {
	out := new(entities.Promotion)
	err := c.cc.Invoke(ctx, "/services.PromotionService/GetPromotion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *promotionServiceClient) UpdatePromotion(ctx context.Context, in *UpdatePromotionRequest, opts ...grpc.CallOption) (*entities.Promotion, error) {
	out := new(entities.Promotion)
	err := c.cc.Invoke(ctx, "/services.PromotionService/UpdatePromotion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *promotionServiceClient) DeletePromotion(ctx context.Context, in *DeletePromotionRequest, opts ...grpc.CallOption) (*entities.Promotion, error) {
	out := new(entities.Promotion)
	err := c.cc.Invoke(ctx, "/services.PromotionService/DeletePromotion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PromotionServiceServer is the server API for PromotionService service.
type PromotionServiceServer interface {
	// Get (read) the list of all promotions.
	GetAllPromotions(context.Context, *empty.Empty) (*GetAllPromotionsResponse, error)
	// Create a new promotion.
	CreatePromotion(context.Context, *CreatePromotionRequest) (*entities.Promotion, error)
	// Get (read) a specific promotion by its id.
	GetPromotion(context.Context, *GetPromotionRequest) (*entities.Promotion, error)
	// Update a specific promotion by its id.
	UpdatePromotion(context.Context, *UpdatePromotionRequest) (*entities.Promotion, error)
	// Delete a specific promotion by its id.
	DeletePromotion(context.Context, *DeletePromotionRequest) (*entities.Promotion, error)
}

// UnimplementedPromotionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPromotionServiceServer struct {
}

func (*UnimplementedPromotionServiceServer) GetAllPromotions(ctx context.Context, req *empty.Empty) (*GetAllPromotionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllPromotions not implemented")
}
func (*UnimplementedPromotionServiceServer) CreatePromotion(ctx context.Context, req *CreatePromotionRequest) (*entities.Promotion, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePromotion not implemented")
}
func (*UnimplementedPromotionServiceServer) GetPromotion(ctx context.Context, req *GetPromotionRequest) (*entities.Promotion, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPromotion not implemented")
}
func (*UnimplementedPromotionServiceServer) UpdatePromotion(ctx context.Context, req *UpdatePromotionRequest) (*entities.Promotion, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePromotion not implemented")
}
func (*UnimplementedPromotionServiceServer) DeletePromotion(ctx context.Context, req *DeletePromotionRequest) (*entities.Promotion, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePromotion not implemented")
}

func RegisterPromotionServiceServer(s *grpc.Server, srv PromotionServiceServer) {
	s.RegisterService(&_PromotionService_serviceDesc, srv)
}

func _PromotionService_GetAllPromotions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromotionServiceServer).GetAllPromotions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.PromotionService/GetAllPromotions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromotionServiceServer).GetAllPromotions(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PromotionService_CreatePromotion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePromotionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromotionServiceServer).CreatePromotion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.PromotionService/CreatePromotion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromotionServiceServer).CreatePromotion(ctx, req.(*CreatePromotionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PromotionService_GetPromotion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPromotionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromotionServiceServer).GetPromotion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.PromotionService/GetPromotion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromotionServiceServer).GetPromotion(ctx, req.(*GetPromotionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PromotionService_UpdatePromotion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePromotionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromotionServiceServer).UpdatePromotion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.PromotionService/UpdatePromotion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromotionServiceServer).UpdatePromotion(ctx, req.(*UpdatePromotionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PromotionService_DeletePromotion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePromotionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromotionServiceServer).DeletePromotion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.PromotionService/DeletePromotion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromotionServiceServer).DeletePromotion(ctx, req.(*DeletePromotionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PromotionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "services.PromotionService",
	HandlerType: (*PromotionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllPromotions",
			Handler:    _PromotionService_GetAllPromotions_Handler,
		},
		{
			MethodName: "CreatePromotion",
			Handler:    _PromotionService_CreatePromotion_Handler,
		},
		{
			MethodName: "GetPromotion",
			Handler:    _PromotionService_GetPromotion_Handler,
		},
		{
			MethodName: "UpdatePromotion",
			Handler:    _PromotionService_UpdatePromotion_Handler,
		},
		{
			MethodName: "DeletePromotion",
			Handler:    _PromotionService_DeletePromotion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/proto/services/promotion.proto",
}
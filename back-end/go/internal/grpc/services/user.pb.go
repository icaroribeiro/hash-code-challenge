// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/proto/services/user.proto

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

type GetAllUsersResponse struct {
	Users                []*entities.User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GetAllUsersResponse) Reset()         { *m = GetAllUsersResponse{} }
func (m *GetAllUsersResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllUsersResponse) ProtoMessage()    {}
func (*GetAllUsersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1f04b4dcd711334, []int{0}
}

func (m *GetAllUsersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllUsersResponse.Unmarshal(m, b)
}
func (m *GetAllUsersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllUsersResponse.Marshal(b, m, deterministic)
}
func (m *GetAllUsersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllUsersResponse.Merge(m, src)
}
func (m *GetAllUsersResponse) XXX_Size() int {
	return xxx_messageInfo_GetAllUsersResponse.Size(m)
}
func (m *GetAllUsersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllUsersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllUsersResponse proto.InternalMessageInfo

func (m *GetAllUsersResponse) GetUsers() []*entities.User {
	if m != nil {
		return m.Users
	}
	return nil
}

type CreateUserRequest struct {
	User                 *entities.User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1f04b4dcd711334, []int{1}
}

func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (m *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(m, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetUser() *entities.User {
	if m != nil {
		return m.User
	}
	return nil
}

type GetUserRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserRequest) Reset()         { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()    {}
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1f04b4dcd711334, []int{2}
}

func (m *GetUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserRequest.Unmarshal(m, b)
}
func (m *GetUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserRequest.Marshal(b, m, deterministic)
}
func (m *GetUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserRequest.Merge(m, src)
}
func (m *GetUserRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserRequest.Size(m)
}
func (m *GetUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserRequest proto.InternalMessageInfo

func (m *GetUserRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type UpdateUserRequest struct {
	Id                   string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	User                 *entities.User `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *UpdateUserRequest) Reset()         { *m = UpdateUserRequest{} }
func (m *UpdateUserRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateUserRequest) ProtoMessage()    {}
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1f04b4dcd711334, []int{3}
}

func (m *UpdateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserRequest.Unmarshal(m, b)
}
func (m *UpdateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserRequest.Marshal(b, m, deterministic)
}
func (m *UpdateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserRequest.Merge(m, src)
}
func (m *UpdateUserRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateUserRequest.Size(m)
}
func (m *UpdateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserRequest proto.InternalMessageInfo

func (m *UpdateUserRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateUserRequest) GetUser() *entities.User {
	if m != nil {
		return m.User
	}
	return nil
}

type DeleteUserRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteUserRequest) Reset()         { *m = DeleteUserRequest{} }
func (m *DeleteUserRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteUserRequest) ProtoMessage()    {}
func (*DeleteUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1f04b4dcd711334, []int{4}
}

func (m *DeleteUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteUserRequest.Unmarshal(m, b)
}
func (m *DeleteUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteUserRequest.Marshal(b, m, deterministic)
}
func (m *DeleteUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteUserRequest.Merge(m, src)
}
func (m *DeleteUserRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteUserRequest.Size(m)
}
func (m *DeleteUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteUserRequest proto.InternalMessageInfo

func (m *DeleteUserRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*GetAllUsersResponse)(nil), "services.GetAllUsersResponse")
	proto.RegisterType((*CreateUserRequest)(nil), "services.CreateUserRequest")
	proto.RegisterType((*GetUserRequest)(nil), "services.GetUserRequest")
	proto.RegisterType((*UpdateUserRequest)(nil), "services.UpdateUserRequest")
	proto.RegisterType((*DeleteUserRequest)(nil), "services.DeleteUserRequest")
}

func init() {
	proto.RegisterFile("github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/proto/services/user.proto", fileDescriptor_e1f04b4dcd711334)
}

var fileDescriptor_e1f04b4dcd711334 = []byte{
	// 358 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x69, 0xfc, 0x3f, 0x85, 0x42, 0x57, 0x90, 0x92, 0x22, 0x94, 0xe8, 0xc1, 0x4b, 0x77,
	0xa1, 0x3d, 0x48, 0xf1, 0xa4, 0x56, 0x7a, 0x13, 0xa9, 0xf4, 0xd2, 0x5b, 0xfe, 0x8c, 0xc9, 0xe2,
	0x36, 0x1b, 0x77, 0x37, 0x82, 0x1f, 0xcd, 0x6f, 0x27, 0x49, 0x4c, 0x93, 0xa6, 0xa5, 0x17, 0xbd,
	0xce, 0x9b, 0xf7, 0x5e, 0xf8, 0x4d, 0x16, 0x96, 0x21, 0x37, 0x51, 0xea, 0x51, 0x5f, 0xae, 0x18,
	0xf7, 0x5d, 0x25, 0x15, 0xf7, 0x90, 0x2b, 0xc9, 0x22, 0x57, 0x47, 0x43, 0x5f, 0x06, 0x38, 0xf4,
	0x23, 0x57, 0x08, 0x8c, 0x43, 0x64, 0x9e, 0xeb, 0xbf, 0x0f, 0x31, 0x0e, 0x58, 0x28, 0x19, 0x8f,
	0x0d, 0xaa, 0xd8, 0x15, 0x2c, 0x51, 0xd2, 0x48, 0xa6, 0x51, 0x7d, 0x72, 0x1f, 0x35, 0x4b, 0x35,
	0x2a, 0x9a, 0xcf, 0xc8, 0x69, 0x39, 0xb4, 0xff, 0xab, 0x05, 0x63, 0xc3, 0x0d, 0xdf, 0x68, 0xb1,
	0xfb, 0xa1, 0x94, 0xa1, 0xc0, 0x62, 0xc3, 0x4b, 0xdf, 0x18, 0xae, 0x12, 0xf3, 0x55, 0x88, 0xce,
	0x1d, 0x9c, 0xcf, 0xd0, 0xdc, 0x0b, 0xb1, 0xd0, 0xa8, 0xf4, 0x1c, 0x75, 0x22, 0x63, 0x8d, 0xe4,
	0x1a, 0x8e, 0xb2, 0x04, 0xdd, 0x6b, 0x0d, 0x0e, 0x6e, 0xda, 0xa3, 0x0e, 0x2d, 0x83, 0x69, 0xb6,
	0x37, 0x2f, 0x44, 0xe7, 0x16, 0xba, 0x8f, 0x0a, 0x5d, 0x83, 0xf9, 0x10, 0x3f, 0x52, 0xd4, 0x86,
	0x38, 0x70, 0x98, 0xa9, 0xbd, 0xd6, 0xa0, 0xb5, 0xc3, 0x99, 0x6b, 0xce, 0x00, 0x3a, 0x33, 0x34,
	0x75, 0x57, 0x07, 0x2c, 0x1e, 0xe4, 0x9e, 0xb3, 0xb9, 0xc5, 0x03, 0x67, 0x06, 0xdd, 0x45, 0x12,
	0x34, 0xa2, 0x1b, 0x4b, 0xeb, 0x2a, 0x6b, 0x4f, 0xd5, 0x15, 0x74, 0xa7, 0x28, 0x70, 0x6f, 0xd0,
	0xe8, 0xdb, 0x82, 0x76, 0xa6, 0xbf, 0x16, 0xf7, 0x20, 0x53, 0x68, 0xd7, 0xa8, 0x90, 0x0b, 0x5a,
	0x20, 0xa4, 0x25, 0x42, 0xfa, 0x94, 0x21, 0xb4, 0x2f, 0x69, 0x79, 0x40, 0xba, 0x0b, 0xe2, 0x04,
	0xa0, 0xc2, 0x43, 0xfa, 0xd5, 0xf2, 0x16, 0x34, 0xbb, 0xf1, 0xed, 0x64, 0x0c, 0x27, 0xbf, 0x80,
	0x48, 0x6f, 0xa3, 0x64, 0x9f, 0x69, 0x02, 0x50, 0x31, 0xab, 0xf7, 0x6d, 0x91, 0xdc, 0x65, 0xad,
	0x28, 0xd5, 0xad, 0x5b, 0xec, 0x9a, 0xd6, 0x87, 0x97, 0xe5, 0xf3, 0x9f, 0x7f, 0xde, 0x50, 0x25,
	0xfe, 0xfa, 0x85, 0x78, 0xc7, 0x39, 0xe6, 0xf1, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x80, 0x96,
	0xe7, 0xe0, 0x7b, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	// Get (read) the list of all users.
	GetAllUsers(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetAllUsersResponse, error)
	// Create a new user.
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*entities.User, error)
	// Get (read) a specific user by its id.
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*entities.User, error)
	// Update a specific user by its id.
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*entities.User, error)
	// Delete a specific user by its id.
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*entities.User, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetAllUsers(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetAllUsersResponse, error) {
	out := new(GetAllUsersResponse)
	err := c.cc.Invoke(ctx, "/services.UserService/GetAllUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*entities.User, error) {
	out := new(entities.User)
	err := c.cc.Invoke(ctx, "/services.UserService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*entities.User, error) {
	out := new(entities.User)
	err := c.cc.Invoke(ctx, "/services.UserService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*entities.User, error) {
	out := new(entities.User)
	err := c.cc.Invoke(ctx, "/services.UserService/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*entities.User, error) {
	out := new(entities.User)
	err := c.cc.Invoke(ctx, "/services.UserService/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	// Get (read) the list of all users.
	GetAllUsers(context.Context, *empty.Empty) (*GetAllUsersResponse, error)
	// Create a new user.
	CreateUser(context.Context, *CreateUserRequest) (*entities.User, error)
	// Get (read) a specific user by its id.
	GetUser(context.Context, *GetUserRequest) (*entities.User, error)
	// Update a specific user by its id.
	UpdateUser(context.Context, *UpdateUserRequest) (*entities.User, error)
	// Delete a specific user by its id.
	DeleteUser(context.Context, *DeleteUserRequest) (*entities.User, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) GetAllUsers(ctx context.Context, req *empty.Empty) (*GetAllUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllUsers not implemented")
}
func (*UnimplementedUserServiceServer) CreateUser(ctx context.Context, req *CreateUserRequest) (*entities.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (*UnimplementedUserServiceServer) GetUser(ctx context.Context, req *GetUserRequest) (*entities.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (*UnimplementedUserServiceServer) UpdateUser(ctx context.Context, req *UpdateUserRequest) (*entities.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (*UnimplementedUserServiceServer) DeleteUser(ctx context.Context, req *DeleteUserRequest) (*entities.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_GetAllUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetAllUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.UserService/GetAllUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetAllUsers(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.UserService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.UserService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.UserService/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.UserService/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "services.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllUsers",
			Handler:    _UserService_GetAllUsers_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _UserService_GetUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserService_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserService_DeleteUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/icaroribeiro/hash-code-challenge/back-end/go/internal/proto/services/user.proto",
}

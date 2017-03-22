// Code generated by protoc-gen-go.
// source: grpcuser.proto
// DO NOT EDIT!

/*
Package grpcuser is a generated protocol buffer package.

It is generated from these files:
	grpcuser.proto

It has these top-level messages:
	LoginRequest
	RegisterUser
	Response
	UserByGuid
*/
package grpcuser

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type StatusResponse int32

const (
	StatusResponse_OK               StatusResponse = 0
	StatusResponse_CREDENTIAL_EXIST StatusResponse = 2
	StatusResponse_FAILED           StatusResponse = 3
	StatusResponse_SERVER_ERROR     StatusResponse = 4
)

var StatusResponse_name = map[int32]string{
	0: "OK",
	2: "CREDENTIAL_EXIST",
	3: "FAILED",
	4: "SERVER_ERROR",
}
var StatusResponse_value = map[string]int32{
	"OK":               0,
	"CREDENTIAL_EXIST": 2,
	"FAILED":           3,
	"SERVER_ERROR":     4,
}

func (x StatusResponse) String() string {
	return proto.EnumName(StatusResponse_name, int32(x))
}
func (StatusResponse) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type LoginRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *LoginRequest) Reset()                    { *m = LoginRequest{} }
func (m *LoginRequest) String() string            { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()               {}
func (*LoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type RegisterUser struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	Email    string `protobuf:"bytes,4,opt,name=email" json:"email,omitempty"`
}

func (m *RegisterUser) Reset()                    { *m = RegisterUser{} }
func (m *RegisterUser) String() string            { return proto.CompactTextString(m) }
func (*RegisterUser) ProtoMessage()               {}
func (*RegisterUser) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RegisterUser) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *RegisterUser) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *RegisterUser) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type Response struct {
	Token  string         `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	Status StatusResponse `protobuf:"varint,2,opt,name=status,enum=StatusResponse" json:"status,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Response) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Response) GetStatus() StatusResponse {
	if m != nil {
		return m.Status
	}
	return StatusResponse_OK
}

type UserByGuid struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *UserByGuid) Reset()                    { *m = UserByGuid{} }
func (m *UserByGuid) String() string            { return proto.CompactTextString(m) }
func (*UserByGuid) ProtoMessage()               {}
func (*UserByGuid) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *UserByGuid) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "LoginRequest")
	proto.RegisterType((*RegisterUser)(nil), "RegisterUser")
	proto.RegisterType((*Response)(nil), "Response")
	proto.RegisterType((*UserByGuid)(nil), "UserByGuid")
	proto.RegisterEnum("StatusResponse", StatusResponse_name, StatusResponse_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for User service

type UserClient interface {
	UserLogin(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*Response, error)
	NewUser(ctx context.Context, in *RegisterUser, opts ...grpc.CallOption) (*Response, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) UserLogin(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/User/UserLogin", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) NewUser(ctx context.Context, in *RegisterUser, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/User/NewUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for User service

type UserServer interface {
	UserLogin(context.Context, *LoginRequest) (*Response, error)
	NewUser(context.Context, *RegisterUser) (*Response, error)
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_UserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/UserLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserLogin(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_NewUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).NewUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User/NewUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).NewUser(ctx, req.(*RegisterUser))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserLogin",
			Handler:    _User_UserLogin_Handler,
		},
		{
			MethodName: "NewUser",
			Handler:    _User_NewUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpcuser.proto",
}

func init() { proto.RegisterFile("grpcuser.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x91, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xb7, 0x39, 0xeb, 0xf6, 0xe8, 0x6a, 0x08, 0x3b, 0x94, 0x9d, 0x24, 0x20, 0x13, 0x0f,
	0x3d, 0xcc, 0x4f, 0x30, 0x5d, 0x26, 0xc5, 0xb2, 0x41, 0x3a, 0x87, 0x07, 0x61, 0x54, 0xf7, 0x28,
	0x45, 0xd7, 0xd4, 0x24, 0x65, 0xf8, 0xed, 0xa5, 0xe9, 0x2a, 0x1b, 0x78, 0xf2, 0x14, 0x7e, 0xfc,
	0xde, 0xfb, 0x27, 0xfc, 0x03, 0x5e, 0xaa, 0x8a, 0xf7, 0x52, 0xa3, 0x0a, 0x0a, 0x25, 0x8d, 0x64,
	0x73, 0x70, 0x23, 0x99, 0x66, 0xb9, 0xc0, 0xaf, 0x12, 0xb5, 0xa1, 0x23, 0xe8, 0x55, 0x36, 0x4f,
	0x76, 0xe8, 0xb7, 0xaf, 0xda, 0x37, 0x7d, 0xf1, 0xcb, 0x95, 0x2b, 0x12, 0xad, 0xf7, 0x52, 0x6d,
	0xfd, 0x4e, 0xed, 0x1a, 0x66, 0xaf, 0xe0, 0x0a, 0x4c, 0x33, 0x6d, 0x50, 0x3d, 0x6b, 0x54, 0xff,
	0xcd, 0xa1, 0x43, 0x38, 0xc7, 0x5d, 0x92, 0x7d, 0xfa, 0x5d, 0x2b, 0x6a, 0x60, 0x21, 0xf4, 0x04,
	0xea, 0x42, 0xe6, 0x1a, 0xab, 0x09, 0x23, 0x3f, 0x30, 0x3f, 0xc4, 0xd6, 0x40, 0xc7, 0xe0, 0x68,
	0x93, 0x98, 0x52, 0xdb, 0x44, 0x6f, 0x72, 0x19, 0xc4, 0x16, 0x9b, 0x35, 0x71, 0xd0, 0x8c, 0x01,
	0x54, 0x0f, 0xbc, 0xff, 0x7e, 0x2c, 0xb3, 0xed, 0xdf, 0x61, 0xb7, 0x11, 0x78, 0xa7, 0xdb, 0xd4,
	0x81, 0xce, 0xf2, 0x89, 0xb4, 0xe8, 0x10, 0xc8, 0x83, 0xe0, 0x33, 0xbe, 0x58, 0x85, 0xd3, 0x68,
	0xc3, 0x5f, 0xc2, 0x78, 0x45, 0x3a, 0x14, 0xc0, 0x99, 0x4f, 0xc3, 0x88, 0xcf, 0xc8, 0x19, 0x25,
	0xe0, 0xc6, 0x5c, 0xac, 0xb9, 0xd8, 0x70, 0x21, 0x96, 0x82, 0x74, 0x27, 0x6b, 0xe8, 0xda, 0x4a,
	0xc6, 0xd0, 0xaf, 0x4e, 0x5b, 0x37, 0x1d, 0x04, 0xc7, 0xb5, 0x8f, 0xfa, 0x41, 0x73, 0x15, 0x6b,
	0xd1, 0x6b, 0xb8, 0x58, 0xe0, 0xde, 0xee, 0x0c, 0x82, 0xe3, 0x56, 0x4f, 0xc6, 0xde, 0x1c, 0xfb,
	0x83, 0x77, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xcc, 0xd7, 0x5d, 0x03, 0xd3, 0x01, 0x00, 0x00,
}

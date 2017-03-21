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
	StatusResponse_OK             StatusResponse = 0
	StatusResponse_USERNAME_EXIST StatusResponse = 1
	StatusResponse_EMAIL_EXIST    StatusResponse = 2
	StatusResponse_FAILED         StatusResponse = 3
	StatusResponse_SERVER_ERROR   StatusResponse = 4
)

var StatusResponse_name = map[int32]string{
	0: "OK",
	1: "USERNAME_EXIST",
	2: "EMAIL_EXIST",
	3: "FAILED",
	4: "SERVER_ERROR",
}
var StatusResponse_value = map[string]int32{
	"OK":             0,
	"USERNAME_EXIST": 1,
	"EMAIL_EXIST":    2,
	"FAILED":         3,
	"SERVER_ERROR":   4,
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
	// 297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x91, 0x51, 0x4b, 0xc3, 0x30,
	0x14, 0x85, 0xb7, 0x39, 0xeb, 0x76, 0xdd, 0xba, 0x12, 0x7c, 0x28, 0x7b, 0x92, 0x80, 0x4c, 0x7c,
	0xe8, 0xc3, 0xfc, 0x05, 0x15, 0x33, 0x29, 0x76, 0x1b, 0xa4, 0x6e, 0x08, 0x0a, 0xa3, 0xba, 0x4b,
	0x29, 0xba, 0xa6, 0x26, 0x29, 0xc3, 0x7f, 0x2f, 0x4d, 0x5b, 0xd9, 0xc0, 0x27, 0x9f, 0xc2, 0xc9,
	0x77, 0xcf, 0x49, 0x38, 0x17, 0xec, 0x44, 0xe6, 0xef, 0x85, 0x42, 0xe9, 0xe5, 0x52, 0x68, 0x41,
	0x67, 0x30, 0x08, 0x45, 0x92, 0x66, 0x1c, 0xbf, 0x0a, 0x54, 0x9a, 0x8c, 0xa1, 0x57, 0xd2, 0x2c,
	0xde, 0xa1, 0xdb, 0xbe, 0x6c, 0x5f, 0xf7, 0xf9, 0xaf, 0x2e, 0x59, 0x1e, 0x2b, 0xb5, 0x17, 0x72,
	0xeb, 0x76, 0x2a, 0xd6, 0x68, 0xfa, 0x0a, 0x03, 0x8e, 0x49, 0xaa, 0x34, 0xca, 0x95, 0x42, 0xf9,
	0xdf, 0x1c, 0x72, 0x01, 0xa7, 0xb8, 0x8b, 0xd3, 0x4f, 0xb7, 0x6b, 0x40, 0x25, 0x68, 0x00, 0x3d,
	0x8e, 0x2a, 0x17, 0x99, 0xc2, 0x72, 0x42, 0x8b, 0x0f, 0xcc, 0xea, 0xd8, 0x4a, 0x90, 0x09, 0x58,
	0x4a, 0xc7, 0xba, 0x50, 0x26, 0xd1, 0x9e, 0x8e, 0xbc, 0xc8, 0xc8, 0xc6, 0xc6, 0x6b, 0x4c, 0x29,
	0x40, 0xf9, 0xc1, 0xbb, 0xef, 0x87, 0x22, 0xdd, 0xfe, 0x1d, 0x76, 0xf3, 0x02, 0xf6, 0xb1, 0x9b,
	0x58, 0xd0, 0x59, 0x3e, 0x3a, 0x2d, 0x42, 0xc0, 0x5e, 0x45, 0x8c, 0x2f, 0xfc, 0x39, 0xdb, 0xb0,
	0xe7, 0x20, 0x7a, 0x72, 0xda, 0x64, 0x04, 0xe7, 0x6c, 0xee, 0x07, 0x61, 0x7d, 0xd1, 0x21, 0x00,
	0xd6, 0xcc, 0x0f, 0x42, 0x76, 0xef, 0x9c, 0x10, 0x07, 0x06, 0x11, 0xe3, 0x6b, 0xc6, 0x37, 0x8c,
	0xf3, 0x25, 0x77, 0xba, 0xd3, 0x35, 0x74, 0x4d, 0x43, 0x13, 0xe8, 0x97, 0xa7, 0x69, 0x9f, 0x0c,
	0xbd, 0xc3, 0x2d, 0x8c, 0xfb, 0x5e, 0xf3, 0x32, 0x6d, 0x91, 0x2b, 0x38, 0x5b, 0xe0, 0xde, 0x78,
	0x86, 0xde, 0x61, 0xc9, 0x47, 0x63, 0x6f, 0x96, 0x59, 0xe8, 0xed, 0x4f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xdf, 0x62, 0x3c, 0x3e, 0xe2, 0x01, 0x00, 0x00,
}

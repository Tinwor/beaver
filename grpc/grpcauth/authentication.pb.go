// Code generated by protoc-gen-go.
// source: authentication.proto
// DO NOT EDIT!

/*
Package authentication is a generated protocol buffer package.

It is generated from these files:
	authentication.proto

It has these top-level messages:
	LoginRequest
	RegisterRequest
	AuthenticationResponse
*/
package grpcauth

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
	StatusResponse_OK           StatusResponse = 0
	StatusResponse_SERVER_ERROR StatusResponse = 1
	StatusResponse_FAILURE      StatusResponse = 2
	StatusResponse_USER_EXIST   StatusResponse = 3
	StatusResponse_EMAIL_EXIST  StatusResponse = 4
)

var StatusResponse_name = map[int32]string{
	0: "OK",
	1: "SERVER_ERROR",
	2: "FAILURE",
	3: "USER_EXIST",
	4: "EMAIL_EXIST",
}
var StatusResponse_value = map[string]int32{
	"OK":           0,
	"SERVER_ERROR": 1,
	"FAILURE":      2,
	"USER_EXIST":   3,
	"EMAIL_EXIST":  4,
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

type RegisterRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	Email    string `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
}

func (m *RegisterRequest) Reset()                    { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string            { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()               {}
func (*RegisterRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RegisterRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *RegisterRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *RegisterRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type AuthenticationResponse struct {
	Status StatusResponse `protobuf:"varint,1,opt,name=status,enum=StatusResponse" json:"status,omitempty"`
	Token  string         `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
}

func (m *AuthenticationResponse) Reset()                    { *m = AuthenticationResponse{} }
func (m *AuthenticationResponse) String() string            { return proto.CompactTextString(m) }
func (*AuthenticationResponse) ProtoMessage()               {}
func (*AuthenticationResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AuthenticationResponse) GetStatus() StatusResponse {
	if m != nil {
		return m.Status
	}
	return StatusResponse_OK
}

func (m *AuthenticationResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "LoginRequest")
	proto.RegisterType((*RegisterRequest)(nil), "RegisterRequest")
	proto.RegisterType((*AuthenticationResponse)(nil), "AuthenticationResponse")
	proto.RegisterEnum("StatusResponse", StatusResponse_name, StatusResponse_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Authentication service

type AuthenticationClient interface {
	RegisterNewUser(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*AuthenticationResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*AuthenticationResponse, error)
}

type authenticationClient struct {
	cc *grpc.ClientConn
}

func NewAuthenticationClient(cc *grpc.ClientConn) AuthenticationClient {
	return &authenticationClient{cc}
}

func (c *authenticationClient) RegisterNewUser(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*AuthenticationResponse, error) {
	out := new(AuthenticationResponse)
	err := grpc.Invoke(ctx, "/Authentication/RegisterNewUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*AuthenticationResponse, error) {
	out := new(AuthenticationResponse)
	err := grpc.Invoke(ctx, "/Authentication/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Authentication service

type AuthenticationServer interface {
	RegisterNewUser(context.Context, *RegisterRequest) (*AuthenticationResponse, error)
	Login(context.Context, *LoginRequest) (*AuthenticationResponse, error)
}

func RegisterAuthenticationServer(s *grpc.Server, srv AuthenticationServer) {
	s.RegisterService(&_Authentication_serviceDesc, srv)
}

func _Authentication_RegisterNewUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).RegisterNewUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Authentication/RegisterNewUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).RegisterNewUser(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Authentication/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Authentication_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Authentication",
	HandlerType: (*AuthenticationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterNewUser",
			Handler:    _Authentication_RegisterNewUser_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Authentication_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authentication.proto",
}

func init() { proto.RegisterFile("authentication.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x91, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0x87, 0x9b, 0xd6, 0x56, 0x9d, 0xd6, 0x74, 0x19, 0x8a, 0x96, 0x9e, 0x24, 0x17, 0xc5, 0x43,
	0xc0, 0x7a, 0x17, 0x72, 0x48, 0x21, 0x18, 0x2d, 0x6c, 0x8c, 0xf6, 0x56, 0x56, 0x1d, 0x6a, 0xd0,
	0xee, 0xc6, 0xec, 0x86, 0x3e, 0x80, 0x2f, 0x2e, 0xf9, 0x53, 0x35, 0x22, 0x5e, 0x3c, 0x7e, 0xf9,
	0x26, 0x33, 0x3b, 0xf3, 0x83, 0x91, 0xc8, 0xcd, 0x33, 0x49, 0x93, 0x3c, 0x0a, 0x93, 0x28, 0xe9,
	0xa6, 0x99, 0x32, 0xca, 0x99, 0xc1, 0x20, 0x54, 0xab, 0x44, 0x72, 0x7a, 0xcb, 0x49, 0x1b, 0x9c,
	0xc0, 0x5e, 0xae, 0x29, 0x93, 0x62, 0x4d, 0x63, 0xeb, 0xd8, 0x3a, 0xdd, 0xe7, 0x9f, 0x5c, 0xb8,
	0x54, 0x68, 0xbd, 0x51, 0xd9, 0xd3, 0xb8, 0x5d, 0xb9, 0x2d, 0x3b, 0x4b, 0x18, 0x72, 0x5a, 0x25,
	0xda, 0x50, 0xf6, 0xcf, 0x56, 0x38, 0x82, 0x2e, 0xad, 0x45, 0xf2, 0x3a, 0xee, 0x94, 0xa2, 0x02,
	0xe7, 0x1e, 0x0e, 0xbd, 0xc6, 0x02, 0x9c, 0x74, 0xaa, 0xa4, 0x26, 0x3c, 0x81, 0x9e, 0x36, 0xc2,
	0xe4, 0xba, 0x9c, 0x62, 0x4f, 0x87, 0x6e, 0x54, 0xe2, 0xb6, 0x80, 0xd7, 0xba, 0x68, 0x6c, 0xd4,
	0x0b, 0xc9, 0x7a, 0x62, 0x05, 0x67, 0x0b, 0xb0, 0x9b, 0xf5, 0xd8, 0x83, 0xf6, 0xfc, 0x8a, 0xb5,
	0x90, 0xc1, 0x20, 0xf2, 0xf9, 0x9d, 0xcf, 0x97, 0x3e, 0xe7, 0x73, 0xce, 0x2c, 0xec, 0xc3, 0xee,
	0xcc, 0x0b, 0xc2, 0x98, 0xfb, 0xac, 0x8d, 0x36, 0x40, 0x1c, 0x15, 0x72, 0x11, 0x44, 0xb7, 0xac,
	0x83, 0x43, 0xe8, 0xfb, 0xd7, 0x5e, 0x10, 0xd6, 0x1f, 0x76, 0xa6, 0xef, 0x16, 0xd8, 0xcd, 0x37,
	0xe3, 0xe5, 0xd7, 0x99, 0x6e, 0x68, 0x13, 0x6b, 0xca, 0x90, 0xb9, 0x3f, 0x0e, 0x37, 0x39, 0x72,
	0x7f, 0xdf, 0xd4, 0x69, 0xe1, 0x39, 0x74, 0xcb, 0xb8, 0xf0, 0xc0, 0xfd, 0x1e, 0xdb, 0x1f, 0xbf,
	0x3c, 0xf4, 0xca, 0xa0, 0x2f, 0x3e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x4e, 0xa2, 0x3e, 0x53, 0x00,
	0x02, 0x00, 0x00,
}
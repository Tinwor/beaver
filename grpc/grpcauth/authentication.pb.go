// Code generated by protoc-gen-go.
// source: authentication.proto
// DO NOT EDIT!

/*
Package authentication is a generated protocol buffer package.

It is generated from these files:
	authentication.proto

It has these top-level messages:
	AuthenticationRequest
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

type AuthenticationStatusResponse int32

const (
	AuthenticationStatusResponse_OK           AuthenticationStatusResponse = 0
	AuthenticationStatusResponse_SERVER_ERROR AuthenticationStatusResponse = 1
	AuthenticationStatusResponse_FAILURE      AuthenticationStatusResponse = 2
	AuthenticationStatusResponse_USER_EXIST   AuthenticationStatusResponse = 3
	AuthenticationStatusResponse_EMAIL_EXIST  AuthenticationStatusResponse = 4
)

var AuthenticationStatusResponse_name = map[int32]string{
	0: "OK",
	1: "SERVER_ERROR",
	2: "FAILURE",
	3: "USER_EXIST",
	4: "EMAIL_EXIST",
}
var AuthenticationStatusResponse_value = map[string]int32{
	"OK":           0,
	"SERVER_ERROR": 1,
	"FAILURE":      2,
	"USER_EXIST":   3,
	"EMAIL_EXIST":  4,
}

func (x AuthenticationStatusResponse) String() string {
	return proto.EnumName(AuthenticationStatusResponse_name, int32(x))
}
func (AuthenticationStatusResponse) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type AuthenticationRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *AuthenticationRequest) Reset()                    { *m = AuthenticationRequest{} }
func (m *AuthenticationRequest) String() string            { return proto.CompactTextString(m) }
func (*AuthenticationRequest) ProtoMessage()               {}
func (*AuthenticationRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AuthenticationRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AuthenticationRequest) GetPassword() string {
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
	Status AuthenticationStatusResponse `protobuf:"varint,1,opt,name=status,enum=AuthenticationStatusResponse" json:"status,omitempty"`
	Token  string                       `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
}

func (m *AuthenticationResponse) Reset()                    { *m = AuthenticationResponse{} }
func (m *AuthenticationResponse) String() string            { return proto.CompactTextString(m) }
func (*AuthenticationResponse) ProtoMessage()               {}
func (*AuthenticationResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AuthenticationResponse) GetStatus() AuthenticationStatusResponse {
	if m != nil {
		return m.Status
	}
	return AuthenticationStatusResponse_OK
}

func (m *AuthenticationResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*AuthenticationRequest)(nil), "AuthenticationRequest")
	proto.RegisterType((*RegisterRequest)(nil), "RegisterRequest")
	proto.RegisterType((*AuthenticationResponse)(nil), "AuthenticationResponse")
	proto.RegisterEnum("AuthenticationStatusResponse", AuthenticationStatusResponse_name, AuthenticationStatusResponse_value)
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
	Login(ctx context.Context, in *AuthenticationRequest, opts ...grpc.CallOption) (*AuthenticationResponse, error)
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

func (c *authenticationClient) Login(ctx context.Context, in *AuthenticationRequest, opts ...grpc.CallOption) (*AuthenticationResponse, error) {
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
	Login(context.Context, *AuthenticationRequest) (*AuthenticationResponse, error)
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
	in := new(AuthenticationRequest)
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
		return srv.(AuthenticationServer).Login(ctx, req.(*AuthenticationRequest))
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
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x92, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x9b, 0xd6, 0x56, 0x9d, 0x4a, 0xbb, 0x0c, 0xb5, 0x96, 0xa2, 0x20, 0x39, 0x89, 0x87,
	0x1c, 0x2a, 0x5e, 0x3c, 0x08, 0x3d, 0x44, 0x08, 0x46, 0x03, 0x1b, 0x23, 0xde, 0xca, 0xaa, 0x63,
	0x0d, 0xda, 0xdd, 0x98, 0xdd, 0xd0, 0x2f, 0xe1, 0x87, 0x96, 0xfc, 0xa9, 0x92, 0x52, 0x7a, 0xf1,
	0xf8, 0xe6, 0xf7, 0x66, 0xdf, 0xf2, 0x76, 0x61, 0x20, 0x32, 0xf3, 0x4e, 0xd2, 0xc4, 0x2f, 0xc2,
	0xc4, 0x4a, 0x3a, 0x49, 0xaa, 0x8c, 0xb2, 0x03, 0x38, 0x9c, 0xd6, 0xe6, 0x9c, 0xbe, 0x32, 0xd2,
	0x06, 0xc7, 0xb0, 0x97, 0x69, 0x4a, 0xa5, 0x58, 0xd0, 0xc8, 0x3a, 0xb5, 0xce, 0xf6, 0xf9, 0xaf,
	0xce, 0x59, 0x22, 0xb4, 0x5e, 0xaa, 0xf4, 0x75, 0xd4, 0x2c, 0xd9, 0x4a, 0xdb, 0x33, 0xe8, 0x73,
	0x9a, 0xc7, 0xda, 0x50, 0xfa, 0xcf, 0xa3, 0x70, 0x00, 0x6d, 0x5a, 0x88, 0xf8, 0x73, 0xd4, 0x2a,
	0x40, 0x29, 0x6c, 0x82, 0xe1, 0xfa, 0x8d, 0x75, 0xa2, 0xa4, 0x26, 0xbc, 0x84, 0x8e, 0x36, 0xc2,
	0x64, 0xba, 0x48, 0xe9, 0x4d, 0x4e, 0x9c, 0xba, 0x31, 0x2c, 0xe0, 0xca, 0xce, 0x2b, 0x73, 0x1e,
	0x63, 0xd4, 0x07, 0xc9, 0x2a, 0xbf, 0x14, 0xe7, 0x6f, 0x70, 0xbc, 0x6d, 0x1b, 0x3b, 0xd0, 0x0c,
	0x6e, 0x59, 0x03, 0x19, 0x1c, 0x84, 0x2e, 0x7f, 0x74, 0xf9, 0xcc, 0xe5, 0x3c, 0xe0, 0xcc, 0xc2,
	0x2e, 0xec, 0xde, 0x4c, 0x3d, 0x3f, 0xe2, 0x2e, 0x6b, 0x62, 0x0f, 0x20, 0x0a, 0x73, 0xf8, 0xe4,
	0x85, 0x0f, 0xac, 0x85, 0x7d, 0xe8, 0xba, 0x77, 0x53, 0xcf, 0xaf, 0x06, 0x3b, 0x93, 0x6f, 0x0b,
	0x7a, 0xf5, 0x20, 0xbc, 0xfe, 0xab, 0xf0, 0x9e, 0x96, 0x91, 0xa6, 0x14, 0x99, 0xb3, 0x56, 0xea,
	0xf8, 0xc8, 0xd9, 0xdc, 0x82, 0xdd, 0xc0, 0x2b, 0x68, 0xfb, 0x6a, 0x1e, 0x4b, 0x1c, 0x3a, 0x1b,
	0xdf, 0x76, 0xcb, 0xee, 0x73, 0xa7, 0xf8, 0x16, 0x17, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe4,
	0x3d, 0x21, 0x42, 0x2e, 0x02, 0x00, 0x00,
}

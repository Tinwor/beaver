// Code generated by protoc-gen-go.
// source: authorization.proto
// DO NOT EDIT!

/*
Package authorization is a generated protocol buffer package.

It is generated from these files:
	authorization.proto

It has these top-level messages:
	TokenRefreshRequest
	TokenRequest
	TokenResponse
	AuthorizationRequest
	AuthorizationResponse
*/
package authorization

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

type AuthorizationStatusResponse int32

const (
	AuthorizationStatusResponse_OK           AuthorizationStatusResponse = 0
	AuthorizationStatusResponse_RENEW        AuthorizationStatusResponse = 1
	AuthorizationStatusResponse_TIMEOUT      AuthorizationStatusResponse = 2
	AuthorizationStatusResponse_ERROR        AuthorizationStatusResponse = 3
	AuthorizationStatusResponse_SERVER_ERROR AuthorizationStatusResponse = 4
)

var AuthorizationStatusResponse_name = map[int32]string{
	0: "OK",
	1: "RENEW",
	2: "TIMEOUT",
	3: "ERROR",
	4: "SERVER_ERROR",
}
var AuthorizationStatusResponse_value = map[string]int32{
	"OK":           0,
	"RENEW":        1,
	"TIMEOUT":      2,
	"ERROR":        3,
	"SERVER_ERROR": 4,
}

func (x AuthorizationStatusResponse) String() string {
	return proto.EnumName(AuthorizationStatusResponse_name, int32(x))
}
func (AuthorizationStatusResponse) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type TokenRefreshRequest struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *TokenRefreshRequest) Reset()                    { *m = TokenRefreshRequest{} }
func (m *TokenRefreshRequest) String() string            { return proto.CompactTextString(m) }
func (*TokenRefreshRequest) ProtoMessage()               {}
func (*TokenRefreshRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TokenRefreshRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type TokenRequest struct {
	Guid string `protobuf:"bytes,1,opt,name=guid" json:"guid,omitempty"`
}

func (m *TokenRequest) Reset()                    { *m = TokenRequest{} }
func (m *TokenRequest) String() string            { return proto.CompactTextString(m) }
func (*TokenRequest) ProtoMessage()               {}
func (*TokenRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TokenRequest) GetGuid() string {
	if m != nil {
		return m.Guid
	}
	return ""
}

type TokenResponse struct {
	Token    string                      `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	Response AuthorizationStatusResponse `protobuf:"varint,2,opt,name=response,enum=AuthorizationStatusResponse" json:"response,omitempty"`
}

func (m *TokenResponse) Reset()                    { *m = TokenResponse{} }
func (m *TokenResponse) String() string            { return proto.CompactTextString(m) }
func (*TokenResponse) ProtoMessage()               {}
func (*TokenResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *TokenResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *TokenResponse) GetResponse() AuthorizationStatusResponse {
	if m != nil {
		return m.Response
	}
	return AuthorizationStatusResponse_OK
}

type AuthorizationRequest struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *AuthorizationRequest) Reset()                    { *m = AuthorizationRequest{} }
func (m *AuthorizationRequest) String() string            { return proto.CompactTextString(m) }
func (*AuthorizationRequest) ProtoMessage()               {}
func (*AuthorizationRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AuthorizationRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type AuthorizationResponse struct {
	Response AuthorizationStatusResponse `protobuf:"varint,1,opt,name=response,enum=AuthorizationStatusResponse" json:"response,omitempty"`
	Guid     string                      `protobuf:"bytes,2,opt,name=guid" json:"guid,omitempty"`
}

func (m *AuthorizationResponse) Reset()                    { *m = AuthorizationResponse{} }
func (m *AuthorizationResponse) String() string            { return proto.CompactTextString(m) }
func (*AuthorizationResponse) ProtoMessage()               {}
func (*AuthorizationResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *AuthorizationResponse) GetResponse() AuthorizationStatusResponse {
	if m != nil {
		return m.Response
	}
	return AuthorizationStatusResponse_OK
}

func (m *AuthorizationResponse) GetGuid() string {
	if m != nil {
		return m.Guid
	}
	return ""
}

func init() {
	proto.RegisterType((*TokenRefreshRequest)(nil), "TokenRefreshRequest")
	proto.RegisterType((*TokenRequest)(nil), "TokenRequest")
	proto.RegisterType((*TokenResponse)(nil), "TokenResponse")
	proto.RegisterType((*AuthorizationRequest)(nil), "AuthorizationRequest")
	proto.RegisterType((*AuthorizationResponse)(nil), "AuthorizationResponse")
	proto.RegisterEnum("AuthorizationStatusResponse", AuthorizationStatusResponse_name, AuthorizationStatusResponse_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Authorization service

type AuthorizationClient interface {
	NewToken(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*TokenResponse, error)
	RefreshToken(ctx context.Context, in *TokenRefreshRequest, opts ...grpc.CallOption) (*TokenResponse, error)
	AuthorizeUser(ctx context.Context, in *AuthorizationRequest, opts ...grpc.CallOption) (*AuthorizationResponse, error)
}

type authorizationClient struct {
	cc *grpc.ClientConn
}

func NewAuthorizationClient(cc *grpc.ClientConn) AuthorizationClient {
	return &authorizationClient{cc}
}

func (c *authorizationClient) NewToken(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*TokenResponse, error) {
	out := new(TokenResponse)
	err := grpc.Invoke(ctx, "/Authorization/NewToken", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationClient) RefreshToken(ctx context.Context, in *TokenRefreshRequest, opts ...grpc.CallOption) (*TokenResponse, error) {
	out := new(TokenResponse)
	err := grpc.Invoke(ctx, "/Authorization/RefreshToken", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationClient) AuthorizeUser(ctx context.Context, in *AuthorizationRequest, opts ...grpc.CallOption) (*AuthorizationResponse, error) {
	out := new(AuthorizationResponse)
	err := grpc.Invoke(ctx, "/Authorization/AuthorizeUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Authorization service

type AuthorizationServer interface {
	NewToken(context.Context, *TokenRequest) (*TokenResponse, error)
	RefreshToken(context.Context, *TokenRefreshRequest) (*TokenResponse, error)
	AuthorizeUser(context.Context, *AuthorizationRequest) (*AuthorizationResponse, error)
}

func RegisterAuthorizationServer(s *grpc.Server, srv AuthorizationServer) {
	s.RegisterService(&_Authorization_serviceDesc, srv)
}

func _Authorization_NewToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).NewToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Authorization/NewToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).NewToken(ctx, req.(*TokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authorization_RefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenRefreshRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).RefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Authorization/RefreshToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).RefreshToken(ctx, req.(*TokenRefreshRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authorization_AuthorizeUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).AuthorizeUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Authorization/AuthorizeUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).AuthorizeUser(ctx, req.(*AuthorizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Authorization_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Authorization",
	HandlerType: (*AuthorizationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewToken",
			Handler:    _Authorization_NewToken_Handler,
		},
		{
			MethodName: "RefreshToken",
			Handler:    _Authorization_RefreshToken_Handler,
		},
		{
			MethodName: "AuthorizeUser",
			Handler:    _Authorization_AuthorizeUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authorization.proto",
}

func init() { proto.RegisterFile("authorization.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0x9b, 0xba, 0xcd, 0xed, 0xd9, 0x8e, 0x92, 0x75, 0x32, 0xa6, 0x87, 0x91, 0xd3, 0x70,
	0x92, 0xc3, 0x04, 0xf1, 0xa8, 0x87, 0x1c, 0x44, 0x5c, 0x21, 0xeb, 0x14, 0xbc, 0x8c, 0x8a, 0xd1,
	0x15, 0xa1, 0x99, 0x4d, 0x8a, 0xe0, 0x1f, 0xe6, 0xdf, 0x27, 0x6b, 0xe3, 0xd6, 0x4a, 0x19, 0x78,
	0x4b, 0x1e, 0xbf, 0xf7, 0xf2, 0xbd, 0xef, 0x0b, 0xf4, 0xa2, 0x4c, 0xaf, 0x64, 0x1a, 0x7f, 0x45,
	0x3a, 0x96, 0x09, 0x5d, 0xa7, 0x52, 0x4b, 0x32, 0x81, 0x5e, 0x28, 0xdf, 0x45, 0xc2, 0xc5, 0x6b,
	0x2a, 0xd4, 0x8a, 0x8b, 0x8f, 0x4c, 0x28, 0x8d, 0x7d, 0x68, 0xea, 0x4d, 0x79, 0x80, 0x46, 0x68,
	0xdc, 0xe1, 0xc5, 0x85, 0x10, 0x70, 0x0c, 0x5c, 0x50, 0x18, 0x1a, 0x6f, 0x59, 0xfc, 0x62, 0xa0,
	0xfc, 0x4c, 0x96, 0xe0, 0x1a, 0x46, 0xad, 0x65, 0xa2, 0x44, 0xfd, 0x28, 0x7c, 0x05, 0xed, 0xd4,
	0x10, 0x03, 0x7b, 0x84, 0xc6, 0xdd, 0xe9, 0x29, 0xbd, 0x29, 0xeb, 0x9b, 0xeb, 0x48, 0x67, 0xea,
	0x77, 0x0a, 0xdf, 0xd2, 0xe4, 0x1c, 0xfc, 0x0a, 0xb8, 0x5f, 0xb2, 0x80, 0xfe, 0x1f, 0xda, 0xc8,
	0x2a, 0x0b, 0x40, 0xff, 0x11, 0xb0, 0xdd, 0xda, 0xde, 0x6d, 0x7d, 0xf6, 0x04, 0x27, 0x7b, 0x9a,
	0x71, 0x0b, 0xec, 0xe0, 0xce, 0xb3, 0x70, 0x07, 0x9a, 0x9c, 0xcd, 0xd8, 0xa3, 0x87, 0xf0, 0x11,
	0x1c, 0x86, 0xb7, 0xf7, 0x2c, 0x58, 0x84, 0x9e, 0xbd, 0xa9, 0x33, 0xce, 0x03, 0xee, 0x1d, 0x60,
	0x0f, 0x9c, 0x39, 0xe3, 0x0f, 0x8c, 0x2f, 0x8b, 0x4a, 0x63, 0xfa, 0x8d, 0xc0, 0xad, 0x0c, 0xc7,
	0x13, 0x68, 0xcf, 0xc4, 0x67, 0x6e, 0x33, 0x76, 0x69, 0x39, 0x92, 0x61, 0x97, 0x56, 0xdc, 0x27,
	0x16, 0xbe, 0x04, 0xc7, 0x84, 0x5b, 0x34, 0xf8, 0xb4, 0x26, 0xf0, 0x9a, 0xbe, 0xeb, 0xdd, 0xab,
	0x62, 0xa1, 0x44, 0x8a, 0xfb, 0xb4, 0xce, 0xf7, 0xe1, 0x31, 0xad, 0x35, 0x98, 0x58, 0xcf, 0xad,
	0xfc, 0x8b, 0x5d, 0xfc, 0x04, 0x00, 0x00, 0xff, 0xff, 0x12, 0xc8, 0xa9, 0xc7, 0x79, 0x02, 0x00,
	0x00,
}

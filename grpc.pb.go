// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc.proto

package root

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type AccountRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountRequest) Reset()         { *m = AccountRequest{} }
func (m *AccountRequest) String() string { return proto.CompactTextString(m) }
func (*AccountRequest) ProtoMessage()    {}
func (*AccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{0}
}

func (m *AccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountRequest.Unmarshal(m, b)
}
func (m *AccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountRequest.Marshal(b, m, deterministic)
}
func (m *AccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountRequest.Merge(m, src)
}
func (m *AccountRequest) XXX_Size() int {
	return xxx_messageInfo_AccountRequest.Size(m)
}
func (m *AccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AccountRequest proto.InternalMessageInfo

func (m *AccountRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AccountRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type AccountResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountResponse) Reset()         { *m = AccountResponse{} }
func (m *AccountResponse) String() string { return proto.CompactTextString(m) }
func (*AccountResponse) ProtoMessage()    {}
func (*AccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{1}
}

func (m *AccountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountResponse.Unmarshal(m, b)
}
func (m *AccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountResponse.Marshal(b, m, deterministic)
}
func (m *AccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountResponse.Merge(m, src)
}
func (m *AccountResponse) XXX_Size() int {
	return xxx_messageInfo_AccountResponse.Size(m)
}
func (m *AccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AccountResponse proto.InternalMessageInfo

func (m *AccountResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AccountResponse) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func init() {
	proto.RegisterType((*AccountRequest)(nil), "root.AccountRequest")
	proto.RegisterType((*AccountResponse)(nil), "root.AccountResponse")
}

func init() { proto.RegisterFile("grpc.proto", fileDescriptor_bedfbfc9b54e5600) }

var fileDescriptor_bedfbfc9b54e5600 = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x2f, 0x2a, 0x48,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x29, 0xca, 0xcf, 0x2f, 0x51, 0xf2, 0xe0, 0xe2,
	0x73, 0x4c, 0x4e, 0xce, 0x2f, 0xcd, 0x2b, 0x09, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x92,
	0xe2, 0xe2, 0x28, 0x2d, 0x4e, 0x2d, 0xca, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0,
	0x0c, 0x82, 0xf3, 0x41, 0x72, 0x05, 0x89, 0xc5, 0xc5, 0xe5, 0xf9, 0x45, 0x29, 0x12, 0x4c, 0x10,
	0x39, 0x18, 0x5f, 0xc9, 0x96, 0x8b, 0x1f, 0x6e, 0x52, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x10,
	0x1f, 0x17, 0x53, 0x66, 0x0a, 0xd4, 0x10, 0xa6, 0xcc, 0x14, 0x14, 0xa3, 0x99, 0x50, 0x8d, 0x36,
	0xf2, 0xe2, 0xe2, 0xf0, 0x4c, 0x49, 0xcd, 0x2b, 0xc9, 0x2c, 0xa9, 0x14, 0xb2, 0xe3, 0xe2, 0x75,
	0x2e, 0x4a, 0x4d, 0x2c, 0x49, 0x85, 0x1a, 0x28, 0x24, 0xa2, 0x07, 0x72, 0xac, 0x1e, 0xaa, 0x4b,
	0xa5, 0x44, 0xd1, 0x44, 0x21, 0xb6, 0x2a, 0x31, 0x24, 0xb1, 0x81, 0x7d, 0x68, 0x0c, 0x08, 0x00,
	0x00, 0xff, 0xff, 0x7b, 0xb7, 0x13, 0xb2, 0xef, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IdentityClient is the client API for Identity service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IdentityClient interface {
	CreateAccount(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountResponse, error)
}

type identityClient struct {
	cc *grpc.ClientConn
}

func NewIdentityClient(cc *grpc.ClientConn) IdentityClient {
	return &identityClient{cc}
}

func (c *identityClient) CreateAccount(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountResponse, error) {
	out := new(AccountResponse)
	err := c.cc.Invoke(ctx, "/root.Identity/CreateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IdentityServer is the server API for Identity service.
type IdentityServer interface {
	CreateAccount(context.Context, *AccountRequest) (*AccountResponse, error)
}

// UnimplementedIdentityServer can be embedded to have forward compatible implementations.
type UnimplementedIdentityServer struct {
}

func (*UnimplementedIdentityServer) CreateAccount(ctx context.Context, req *AccountRequest) (*AccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}

func RegisterIdentityServer(s *grpc.Server, srv IdentityServer) {
	s.RegisterService(&_Identity_serviceDesc, srv)
}

func _Identity_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/root.Identity/CreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServer).CreateAccount(ctx, req.(*AccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Identity_serviceDesc = grpc.ServiceDesc{
	ServiceName: "root.Identity",
	HandlerType: (*IdentityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAccount",
			Handler:    _Identity_CreateAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc.proto",
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.ext.proto

package pb

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

type SignInReq struct {
	PhoneNumber          string   `protobuf:"bytes,1,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Code                 string   `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	DeviceId             int64    `protobuf:"varint,3,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignInReq) Reset()         { *m = SignInReq{} }
func (m *SignInReq) String() string { return proto.CompactTextString(m) }
func (*SignInReq) ProtoMessage()    {}
func (*SignInReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f85764c91792e488, []int{0}
}

func (m *SignInReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignInReq.Unmarshal(m, b)
}
func (m *SignInReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignInReq.Marshal(b, m, deterministic)
}
func (m *SignInReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignInReq.Merge(m, src)
}
func (m *SignInReq) XXX_Size() int {
	return xxx_messageInfo_SignInReq.Size(m)
}
func (m *SignInReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SignInReq.DiscardUnknown(m)
}

var xxx_messageInfo_SignInReq proto.InternalMessageInfo

func (m *SignInReq) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *SignInReq) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *SignInReq) GetDeviceId() int64 {
	if m != nil {
		return m.DeviceId
	}
	return 0
}

type SignInResp struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignInResp) Reset()         { *m = SignInResp{} }
func (m *SignInResp) String() string { return proto.CompactTextString(m) }
func (*SignInResp) ProtoMessage()    {}
func (*SignInResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_f85764c91792e488, []int{1}
}

func (m *SignInResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignInResp.Unmarshal(m, b)
}
func (m *SignInResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignInResp.Marshal(b, m, deterministic)
}
func (m *SignInResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignInResp.Merge(m, src)
}
func (m *SignInResp) XXX_Size() int {
	return xxx_messageInfo_SignInResp.Size(m)
}
func (m *SignInResp) XXX_DiscardUnknown() {
	xxx_messageInfo_SignInResp.DiscardUnknown(m)
}

var xxx_messageInfo_SignInResp proto.InternalMessageInfo

func (m *SignInResp) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *SignInResp) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type UpdateUserReq struct {
	Nickname             string   `protobuf:"bytes,1,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Sex                  int32    `protobuf:"varint,2,opt,name=sex,proto3" json:"sex,omitempty"`
	AvatarUrl            string   `protobuf:"bytes,3,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	Extra                string   `protobuf:"bytes,4,opt,name=extra,proto3" json:"extra,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserReq) Reset()         { *m = UpdateUserReq{} }
func (m *UpdateUserReq) String() string { return proto.CompactTextString(m) }
func (*UpdateUserReq) ProtoMessage()    {}
func (*UpdateUserReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f85764c91792e488, []int{2}
}

func (m *UpdateUserReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserReq.Unmarshal(m, b)
}
func (m *UpdateUserReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserReq.Marshal(b, m, deterministic)
}
func (m *UpdateUserReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserReq.Merge(m, src)
}
func (m *UpdateUserReq) XXX_Size() int {
	return xxx_messageInfo_UpdateUserReq.Size(m)
}
func (m *UpdateUserReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserReq proto.InternalMessageInfo

func (m *UpdateUserReq) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *UpdateUserReq) GetSex() int32 {
	if m != nil {
		return m.Sex
	}
	return 0
}

func (m *UpdateUserReq) GetAvatarUrl() string {
	if m != nil {
		return m.AvatarUrl
	}
	return ""
}

func (m *UpdateUserReq) GetExtra() string {
	if m != nil {
		return m.Extra
	}
	return ""
}

type UpdateUserResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserResp) Reset()         { *m = UpdateUserResp{} }
func (m *UpdateUserResp) String() string { return proto.CompactTextString(m) }
func (*UpdateUserResp) ProtoMessage()    {}
func (*UpdateUserResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_f85764c91792e488, []int{3}
}

func (m *UpdateUserResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserResp.Unmarshal(m, b)
}
func (m *UpdateUserResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserResp.Marshal(b, m, deterministic)
}
func (m *UpdateUserResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserResp.Merge(m, src)
}
func (m *UpdateUserResp) XXX_Size() int {
	return xxx_messageInfo_UpdateUserResp.Size(m)
}
func (m *UpdateUserResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserResp.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserResp proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SignInReq)(nil), "pb.SignInReq")
	proto.RegisterType((*SignInResp)(nil), "pb.SignInResp")
	proto.RegisterType((*UpdateUserReq)(nil), "pb.UpdateUserReq")
	proto.RegisterType((*UpdateUserResp)(nil), "pb.UpdateUserResp")
}

func init() { proto.RegisterFile("user.ext.proto", fileDescriptor_f85764c91792e488) }

var fileDescriptor_f85764c91792e488 = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0x3f, 0x4f, 0xf3, 0x30,
	0x10, 0xc6, 0x95, 0xfe, 0xcf, 0xbd, 0x6f, 0x4b, 0x39, 0x21, 0x11, 0x05, 0x21, 0x95, 0x2c, 0x54,
	0x0c, 0x19, 0xe8, 0xc8, 0x8c, 0x50, 0x17, 0x86, 0xa0, 0xce, 0x91, 0x13, 0x9f, 0x20, 0x6a, 0xeb,
	0x18, 0xdb, 0xa9, 0xf2, 0x3d, 0xf8, 0xc2, 0xc8, 0x76, 0x28, 0x2d, 0xdb, 0xfd, 0x9e, 0xbb, 0xdc,
	0xf3, 0x5c, 0x0c, 0xb3, 0x46, 0x93, 0x4a, 0xa9, 0x35, 0xa9, 0x54, 0xb5, 0xa9, 0xb1, 0x27, 0x8b,
	0xd8, 0x6b, 0x95, 0xe8, 0xb4, 0x24, 0x87, 0xf0, 0xad, 0x7a, 0x17, 0x6b, 0x91, 0xd1, 0x27, 0xde,
	0xc1, 0x7f, 0xf9, 0x51, 0x0b, 0xca, 0x45, 0xb3, 0x2f, 0x48, 0x45, 0xc1, 0x22, 0x58, 0x86, 0xd9,
	0x3f, 0xa7, 0xbd, 0x3a, 0x09, 0x11, 0x06, 0x65, 0xcd, 0x29, 0xea, 0xb9, 0x96, 0xab, 0xf1, 0x06,
	0x42, 0x4e, 0x87, 0xaa, 0xa4, 0xbc, 0xe2, 0x51, 0x7f, 0x11, 0x2c, 0xfb, 0xd9, 0xc4, 0x0b, 0x6b,
	0x9e, 0x3c, 0x01, 0xfc, 0x18, 0x68, 0x89, 0xd7, 0x30, 0xb6, 0x01, 0xec, 0x60, 0xe0, 0x06, 0x47,
	0x16, 0xd7, 0x1c, 0xaf, 0x60, 0x68, 0xea, 0x2d, 0x89, 0x6e, 0xb1, 0x87, 0x44, 0xc1, 0x74, 0x23,
	0x39, 0x33, 0xb4, 0xd1, 0xa4, 0x6c, 0xc2, 0x18, 0x26, 0xa2, 0x2a, 0xb7, 0x82, 0xed, 0xa9, 0x4b,
	0x77, 0x64, 0x9c, 0x43, 0x5f, 0x53, 0xeb, 0x16, 0x0c, 0x33, 0x5b, 0xe2, 0x2d, 0x00, 0x3b, 0x30,
	0xc3, 0x54, 0xde, 0xa8, 0x9d, 0x4b, 0x16, 0x66, 0xa1, 0x57, 0x36, 0x6a, 0x67, 0x3d, 0xa9, 0x35,
	0x8a, 0x45, 0x03, 0xef, 0xe9, 0x20, 0x99, 0xc3, 0xec, 0xd4, 0x53, 0xcb, 0xc7, 0xaf, 0x00, 0xc6,
	0x16, 0x9e, 0x5b, 0x83, 0xf7, 0x30, 0xf2, 0xe7, 0xe0, 0x34, 0x95, 0x45, 0x7a, 0xfc, 0x77, 0xf1,
	0xec, 0x14, 0xb5, 0xc4, 0x07, 0x18, 0xbf, 0x90, 0xb1, 0x9f, 0xa1, 0x6b, 0x75, 0x60, 0x47, 0x2f,
	0xce, 0x58, 0x4b, 0x5c, 0x01, 0xfc, 0x5a, 0xe2, 0xa5, 0x6d, 0x9f, 0x9d, 0x1d, 0xe3, 0x5f, 0x49,
	0xcb, 0x62, 0xe4, 0x1e, 0x70, 0xf5, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x21, 0xb9, 0xe2, 0xd0, 0xe6,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserExtClient is the client API for UserExt service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserExtClient interface {
	// 登录
	SignIn(ctx context.Context, in *SignInReq, opts ...grpc.CallOption) (*SignInResp, error)
	// 获取用户信息
	GetUser(ctx context.Context, in *GetUserReq, opts ...grpc.CallOption) (*GetUserResp, error)
	// 更新用户信息
	UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*UpdateUserResp, error)
}

type userExtClient struct {
	cc *grpc.ClientConn
}

func NewUserExtClient(cc *grpc.ClientConn) UserExtClient {
	return &userExtClient{cc}
}

func (c *userExtClient) SignIn(ctx context.Context, in *SignInReq, opts ...grpc.CallOption) (*SignInResp, error) {
	out := new(SignInResp)
	err := c.cc.Invoke(ctx, "/pb.UserExt/SignIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userExtClient) GetUser(ctx context.Context, in *GetUserReq, opts ...grpc.CallOption) (*GetUserResp, error) {
	out := new(GetUserResp)
	err := c.cc.Invoke(ctx, "/pb.UserExt/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userExtClient) UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*UpdateUserResp, error) {
	out := new(UpdateUserResp)
	err := c.cc.Invoke(ctx, "/pb.UserExt/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserExtServer is the server API for UserExt service.
type UserExtServer interface {
	// 登录
	SignIn(context.Context, *SignInReq) (*SignInResp, error)
	// 获取用户信息
	GetUser(context.Context, *GetUserReq) (*GetUserResp, error)
	// 更新用户信息
	UpdateUser(context.Context, *UpdateUserReq) (*UpdateUserResp, error)
}

// UnimplementedUserExtServer can be embedded to have forward compatible implementations.
type UnimplementedUserExtServer struct {
}

func (*UnimplementedUserExtServer) SignIn(ctx context.Context, req *SignInReq) (*SignInResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}
func (*UnimplementedUserExtServer) GetUser(ctx context.Context, req *GetUserReq) (*GetUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (*UnimplementedUserExtServer) UpdateUser(ctx context.Context, req *UpdateUserReq) (*UpdateUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}

func RegisterUserExtServer(s *grpc.Server, srv UserExtServer) {
	s.RegisterService(&_UserExt_serviceDesc, srv)
}

func _UserExt_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserExtServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserExt/SignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserExtServer).SignIn(ctx, req.(*SignInReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserExt_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserExtServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserExt/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserExtServer).GetUser(ctx, req.(*GetUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserExt_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserExtServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserExt/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserExtServer).UpdateUser(ctx, req.(*UpdateUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserExt_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UserExt",
	HandlerType: (*UserExtServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignIn",
			Handler:    _UserExt_SignIn_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _UserExt_GetUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserExt_UpdateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.ext.proto",
}

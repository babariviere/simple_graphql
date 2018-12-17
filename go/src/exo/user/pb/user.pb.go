// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type UsersRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UsersRequest) Reset()         { *m = UsersRequest{} }
func (m *UsersRequest) String() string { return proto.CompactTextString(m) }
func (*UsersRequest) ProtoMessage()    {}
func (*UsersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *UsersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UsersRequest.Unmarshal(m, b)
}
func (m *UsersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UsersRequest.Marshal(b, m, deterministic)
}
func (m *UsersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UsersRequest.Merge(m, src)
}
func (m *UsersRequest) XXX_Size() int {
	return xxx_messageInfo_UsersRequest.Size(m)
}
func (m *UsersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UsersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UsersRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*User)(nil), "pb.User")
	proto.RegisterType((*UsersRequest)(nil), "pb.UsersRequest")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 134 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x2d, 0x4e, 0x2d,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0xd2, 0xe1, 0x62, 0x09, 0x2d,
	0x4e, 0x2d, 0x12, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62,
	0xca, 0x4c, 0x11, 0x12, 0xe1, 0x62, 0x4d, 0xcd, 0x4d, 0xcc, 0xcc, 0x91, 0x60, 0x02, 0x0b, 0x41,
	0x38, 0x4a, 0x7c, 0x5c, 0x3c, 0x20, 0xd5, 0xc5, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x46,
	0xe6, 0x5c, 0xdc, 0x20, 0x7e, 0x70, 0x6a, 0x51, 0x59, 0x66, 0x72, 0xaa, 0x90, 0x06, 0x17, 0x87,
	0x63, 0x4e, 0x0e, 0x58, 0x85, 0x90, 0x80, 0x5e, 0x41, 0x92, 0x1e, 0xb2, 0x62, 0x29, 0x0e, 0x98,
	0x88, 0x01, 0x63, 0x12, 0x1b, 0xd8, 0x05, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x93, 0xf3,
	0x34, 0xad, 0x8f, 0x00, 0x00, 0x00,
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
	AllUsers(ctx context.Context, in *UsersRequest, opts ...grpc.CallOption) (UserService_AllUsersClient, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) AllUsers(ctx context.Context, in *UsersRequest, opts ...grpc.CallOption) (UserService_AllUsersClient, error) {
	stream, err := c.cc.NewStream(ctx, &_UserService_serviceDesc.Streams[0], "/pb.UserService/AllUsers", opts...)
	if err != nil {
		return nil, err
	}
	x := &userServiceAllUsersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UserService_AllUsersClient interface {
	Recv() (*User, error)
	grpc.ClientStream
}

type userServiceAllUsersClient struct {
	grpc.ClientStream
}

func (x *userServiceAllUsersClient) Recv() (*User, error) {
	m := new(User)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	AllUsers(*UsersRequest, UserService_AllUsersServer) error
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_AllUsers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(UsersRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UserServiceServer).AllUsers(m, &userServiceAllUsersServer{stream})
}

type UserService_AllUsersServer interface {
	Send(*User) error
	grpc.ServerStream
}

type userServiceAllUsersServer struct {
	grpc.ServerStream
}

func (x *userServiceAllUsersServer) Send(m *User) error {
	return x.ServerStream.SendMsg(m)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "AllUsers",
			Handler:       _UserService_AllUsers_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "user.proto",
}

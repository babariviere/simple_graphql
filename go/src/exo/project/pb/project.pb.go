// Code generated by protoc-gen-go. DO NOT EDIT.
// source: project.proto

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

type Project struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Project) Reset()         { *m = Project{} }
func (m *Project) String() string { return proto.CompactTextString(m) }
func (*Project) ProtoMessage()    {}
func (*Project) Descriptor() ([]byte, []int) {
	return fileDescriptor_8340e6318dfdfac2, []int{0}
}

func (m *Project) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Project.Unmarshal(m, b)
}
func (m *Project) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Project.Marshal(b, m, deterministic)
}
func (m *Project) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Project.Merge(m, src)
}
func (m *Project) XXX_Size() int {
	return xxx_messageInfo_Project.Size(m)
}
func (m *Project) XXX_DiscardUnknown() {
	xxx_messageInfo_Project.DiscardUnknown(m)
}

var xxx_messageInfo_Project proto.InternalMessageInfo

func (m *Project) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Project) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ProjectsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProjectsRequest) Reset()         { *m = ProjectsRequest{} }
func (m *ProjectsRequest) String() string { return proto.CompactTextString(m) }
func (*ProjectsRequest) ProtoMessage()    {}
func (*ProjectsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8340e6318dfdfac2, []int{1}
}

func (m *ProjectsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectsRequest.Unmarshal(m, b)
}
func (m *ProjectsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectsRequest.Marshal(b, m, deterministic)
}
func (m *ProjectsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectsRequest.Merge(m, src)
}
func (m *ProjectsRequest) XXX_Size() int {
	return xxx_messageInfo_ProjectsRequest.Size(m)
}
func (m *ProjectsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectsRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Project)(nil), "pb.Project")
	proto.RegisterType((*ProjectsRequest)(nil), "pb.ProjectsRequest")
}

func init() { proto.RegisterFile("project.proto", fileDescriptor_8340e6318dfdfac2) }

var fileDescriptor_8340e6318dfdfac2 = []byte{
	// 135 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x28, 0xca, 0xcf,
	0x4a, 0x4d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0xd2, 0xe5,
	0x62, 0x0f, 0x80, 0x08, 0x0a, 0xf1, 0x71, 0x31, 0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70,
	0x06, 0x31, 0x65, 0xa6, 0x08, 0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x81, 0x45,
	0xc0, 0x6c, 0x25, 0x41, 0x2e, 0x7e, 0xa8, 0xf2, 0xe2, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12,
	0x23, 0x67, 0x2e, 0x3e, 0xa8, 0x50, 0x70, 0x6a, 0x51, 0x59, 0x66, 0x72, 0xaa, 0x90, 0x21, 0x17,
	0xb7, 0x63, 0x4e, 0x0e, 0x4c, 0x9d, 0x90, 0xb0, 0x5e, 0x41, 0x92, 0x1e, 0x9a, 0x2e, 0x29, 0x6e,
	0x24, 0x41, 0x03, 0xc6, 0x24, 0x36, 0xb0, 0x8b, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc2,
	0x6e, 0x3a, 0x64, 0xa2, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProjectServiceClient is the client API for ProjectService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProjectServiceClient interface {
	AllProjects(ctx context.Context, in *ProjectsRequest, opts ...grpc.CallOption) (ProjectService_AllProjectsClient, error)
}

type projectServiceClient struct {
	cc *grpc.ClientConn
}

func NewProjectServiceClient(cc *grpc.ClientConn) ProjectServiceClient {
	return &projectServiceClient{cc}
}

func (c *projectServiceClient) AllProjects(ctx context.Context, in *ProjectsRequest, opts ...grpc.CallOption) (ProjectService_AllProjectsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ProjectService_serviceDesc.Streams[0], "/pb.ProjectService/AllProjects", opts...)
	if err != nil {
		return nil, err
	}
	x := &projectServiceAllProjectsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ProjectService_AllProjectsClient interface {
	Recv() (*Project, error)
	grpc.ClientStream
}

type projectServiceAllProjectsClient struct {
	grpc.ClientStream
}

func (x *projectServiceAllProjectsClient) Recv() (*Project, error) {
	m := new(Project)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProjectServiceServer is the server API for ProjectService service.
type ProjectServiceServer interface {
	AllProjects(*ProjectsRequest, ProjectService_AllProjectsServer) error
}

func RegisterProjectServiceServer(s *grpc.Server, srv ProjectServiceServer) {
	s.RegisterService(&_ProjectService_serviceDesc, srv)
}

func _ProjectService_AllProjects_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ProjectsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProjectServiceServer).AllProjects(m, &projectServiceAllProjectsServer{stream})
}

type ProjectService_AllProjectsServer interface {
	Send(*Project) error
	grpc.ServerStream
}

type projectServiceAllProjectsServer struct {
	grpc.ServerStream
}

func (x *projectServiceAllProjectsServer) Send(m *Project) error {
	return x.ServerStream.SendMsg(m)
}

var _ProjectService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ProjectService",
	HandlerType: (*ProjectServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "AllProjects",
			Handler:       _ProjectService_AllProjects_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "project.proto",
}

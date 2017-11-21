// Code generated by protoc-gen-go. DO NOT EDIT.
// source: disneyland.proto

/*
Package disneyland is a generated protocol buffer package.

It is generated from these files:
	disneyland.proto

It has these top-level messages:
	Job
	ListOfJobs
	RequestWithId
	ListJobsRequest
*/
package disneyland

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

type Job_Status int32

const (
	Job_PENDING   Job_Status = 0
	Job_PULLED    Job_Status = 1
	Job_RUNNING   Job_Status = 2
	Job_FAILED    Job_Status = 3
	Job_COMPLETED Job_Status = 4
)

var Job_Status_name = map[int32]string{
	0: "PENDING",
	1: "PULLED",
	2: "RUNNING",
	3: "FAILED",
	4: "COMPLETED",
}
var Job_Status_value = map[string]int32{
	"PENDING":   0,
	"PULLED":    1,
	"RUNNING":   2,
	"FAILED":    3,
	"COMPLETED": 4,
}

func (x Job_Status) String() string {
	return proto.EnumName(Job_Status_name, int32(x))
}
func (Job_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type Job struct {
	Project  string     `protobuf:"bytes,1,opt,name=project" json:"project,omitempty"`
	Id       uint64     `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	Kind     string     `protobuf:"bytes,3,opt,name=kind" json:"kind,omitempty"`
	Status   Job_Status `protobuf:"varint,4,opt,name=status,enum=Job_Status" json:"status,omitempty"`
	Input    string     `protobuf:"bytes,5,opt,name=input" json:"input,omitempty"`
	Output   string     `protobuf:"bytes,6,opt,name=output" json:"output,omitempty"`
	Metadata string     `protobuf:"bytes,7,opt,name=metadata" json:"metadata,omitempty"`
}

func (m *Job) Reset()                    { *m = Job{} }
func (m *Job) String() string            { return proto.CompactTextString(m) }
func (*Job) ProtoMessage()               {}
func (*Job) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Job) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *Job) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Job) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *Job) GetStatus() Job_Status {
	if m != nil {
		return m.Status
	}
	return Job_PENDING
}

func (m *Job) GetInput() string {
	if m != nil {
		return m.Input
	}
	return ""
}

func (m *Job) GetOutput() string {
	if m != nil {
		return m.Output
	}
	return ""
}

func (m *Job) GetMetadata() string {
	if m != nil {
		return m.Metadata
	}
	return ""
}

type ListOfJobs struct {
	Jobs []*Job `protobuf:"bytes,1,rep,name=jobs" json:"jobs,omitempty"`
}

func (m *ListOfJobs) Reset()                    { *m = ListOfJobs{} }
func (m *ListOfJobs) String() string            { return proto.CompactTextString(m) }
func (*ListOfJobs) ProtoMessage()               {}
func (*ListOfJobs) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ListOfJobs) GetJobs() []*Job {
	if m != nil {
		return m.Jobs
	}
	return nil
}

type RequestWithId struct {
	Id uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *RequestWithId) Reset()                    { *m = RequestWithId{} }
func (m *RequestWithId) String() string            { return proto.CompactTextString(m) }
func (*RequestWithId) ProtoMessage()               {}
func (*RequestWithId) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RequestWithId) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ListJobsRequest struct {
	HowMany uint32 `protobuf:"varint,1,opt,name=how_many,json=howMany" json:"how_many,omitempty"`
	Project string `protobuf:"bytes,2,opt,name=project" json:"project,omitempty"`
	Kind    string `protobuf:"bytes,3,opt,name=kind" json:"kind,omitempty"`
}

func (m *ListJobsRequest) Reset()                    { *m = ListJobsRequest{} }
func (m *ListJobsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListJobsRequest) ProtoMessage()               {}
func (*ListJobsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ListJobsRequest) GetHowMany() uint32 {
	if m != nil {
		return m.HowMany
	}
	return 0
}

func (m *ListJobsRequest) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *ListJobsRequest) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func init() {
	proto.RegisterType((*Job)(nil), "Job")
	proto.RegisterType((*ListOfJobs)(nil), "ListOfJobs")
	proto.RegisterType((*RequestWithId)(nil), "RequestWithId")
	proto.RegisterType((*ListJobsRequest)(nil), "ListJobsRequest")
	proto.RegisterEnum("Job_Status", Job_Status_name, Job_Status_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Disneyland service

type DisneylandClient interface {
	CreateJob(ctx context.Context, in *Job, opts ...grpc.CallOption) (*Job, error)
	GetJob(ctx context.Context, in *RequestWithId, opts ...grpc.CallOption) (*Job, error)
	ListJobs(ctx context.Context, in *ListJobsRequest, opts ...grpc.CallOption) (*ListOfJobs, error)
	CreateMultipeJobs(ctx context.Context, in *ListOfJobs, opts ...grpc.CallOption) (*ListOfJobs, error)
	ModifyJob(ctx context.Context, in *Job, opts ...grpc.CallOption) (*Job, error)
	PullPendingJobs(ctx context.Context, in *ListJobsRequest, opts ...grpc.CallOption) (*ListOfJobs, error)
	DeleteJob(ctx context.Context, in *RequestWithId, opts ...grpc.CallOption) (*Job, error)
}

type disneylandClient struct {
	cc *grpc.ClientConn
}

func NewDisneylandClient(cc *grpc.ClientConn) DisneylandClient {
	return &disneylandClient{cc}
}

func (c *disneylandClient) CreateJob(ctx context.Context, in *Job, opts ...grpc.CallOption) (*Job, error) {
	out := new(Job)
	err := grpc.Invoke(ctx, "/Disneyland/CreateJob", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *disneylandClient) GetJob(ctx context.Context, in *RequestWithId, opts ...grpc.CallOption) (*Job, error) {
	out := new(Job)
	err := grpc.Invoke(ctx, "/Disneyland/GetJob", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *disneylandClient) ListJobs(ctx context.Context, in *ListJobsRequest, opts ...grpc.CallOption) (*ListOfJobs, error) {
	out := new(ListOfJobs)
	err := grpc.Invoke(ctx, "/Disneyland/ListJobs", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *disneylandClient) CreateMultipeJobs(ctx context.Context, in *ListOfJobs, opts ...grpc.CallOption) (*ListOfJobs, error) {
	out := new(ListOfJobs)
	err := grpc.Invoke(ctx, "/Disneyland/CreateMultipeJobs", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *disneylandClient) ModifyJob(ctx context.Context, in *Job, opts ...grpc.CallOption) (*Job, error) {
	out := new(Job)
	err := grpc.Invoke(ctx, "/Disneyland/ModifyJob", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *disneylandClient) PullPendingJobs(ctx context.Context, in *ListJobsRequest, opts ...grpc.CallOption) (*ListOfJobs, error) {
	out := new(ListOfJobs)
	err := grpc.Invoke(ctx, "/Disneyland/PullPendingJobs", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *disneylandClient) DeleteJob(ctx context.Context, in *RequestWithId, opts ...grpc.CallOption) (*Job, error) {
	out := new(Job)
	err := grpc.Invoke(ctx, "/Disneyland/DeleteJob", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Disneyland service

type DisneylandServer interface {
	CreateJob(context.Context, *Job) (*Job, error)
	GetJob(context.Context, *RequestWithId) (*Job, error)
	ListJobs(context.Context, *ListJobsRequest) (*ListOfJobs, error)
	CreateMultipeJobs(context.Context, *ListOfJobs) (*ListOfJobs, error)
	ModifyJob(context.Context, *Job) (*Job, error)
	PullPendingJobs(context.Context, *ListJobsRequest) (*ListOfJobs, error)
	DeleteJob(context.Context, *RequestWithId) (*Job, error)
}

func RegisterDisneylandServer(s *grpc.Server, srv DisneylandServer) {
	s.RegisterService(&_Disneyland_serviceDesc, srv)
}

func _Disneyland_CreateJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Job)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DisneylandServer).CreateJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Disneyland/CreateJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DisneylandServer).CreateJob(ctx, req.(*Job))
	}
	return interceptor(ctx, in, info, handler)
}

func _Disneyland_GetJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestWithId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DisneylandServer).GetJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Disneyland/GetJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DisneylandServer).GetJob(ctx, req.(*RequestWithId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Disneyland_ListJobs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListJobsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DisneylandServer).ListJobs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Disneyland/ListJobs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DisneylandServer).ListJobs(ctx, req.(*ListJobsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Disneyland_CreateMultipeJobs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOfJobs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DisneylandServer).CreateMultipeJobs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Disneyland/CreateMultipeJobs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DisneylandServer).CreateMultipeJobs(ctx, req.(*ListOfJobs))
	}
	return interceptor(ctx, in, info, handler)
}

func _Disneyland_ModifyJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Job)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DisneylandServer).ModifyJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Disneyland/ModifyJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DisneylandServer).ModifyJob(ctx, req.(*Job))
	}
	return interceptor(ctx, in, info, handler)
}

func _Disneyland_PullPendingJobs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListJobsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DisneylandServer).PullPendingJobs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Disneyland/PullPendingJobs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DisneylandServer).PullPendingJobs(ctx, req.(*ListJobsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Disneyland_DeleteJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestWithId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DisneylandServer).DeleteJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Disneyland/DeleteJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DisneylandServer).DeleteJob(ctx, req.(*RequestWithId))
	}
	return interceptor(ctx, in, info, handler)
}

var _Disneyland_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Disneyland",
	HandlerType: (*DisneylandServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateJob",
			Handler:    _Disneyland_CreateJob_Handler,
		},
		{
			MethodName: "GetJob",
			Handler:    _Disneyland_GetJob_Handler,
		},
		{
			MethodName: "ListJobs",
			Handler:    _Disneyland_ListJobs_Handler,
		},
		{
			MethodName: "CreateMultipeJobs",
			Handler:    _Disneyland_CreateMultipeJobs_Handler,
		},
		{
			MethodName: "ModifyJob",
			Handler:    _Disneyland_ModifyJob_Handler,
		},
		{
			MethodName: "PullPendingJobs",
			Handler:    _Disneyland_PullPendingJobs_Handler,
		},
		{
			MethodName: "DeleteJob",
			Handler:    _Disneyland_DeleteJob_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "disneyland.proto",
}

func init() { proto.RegisterFile("disneyland.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 433 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x61, 0x6b, 0xdb, 0x30,
	0x10, 0x8d, 0x1d, 0xd7, 0x89, 0x2f, 0x24, 0xf5, 0x8e, 0x31, 0x94, 0x7c, 0x99, 0x71, 0x61, 0x04,
	0x06, 0x1e, 0x64, 0xbf, 0x60, 0xd4, 0x59, 0x49, 0x48, 0x52, 0xe3, 0xad, 0x0c, 0xf6, 0x65, 0xc8,
	0x93, 0xba, 0xa8, 0x73, 0xa5, 0x2c, 0x96, 0x29, 0xf9, 0x2d, 0xfb, 0xad, 0x83, 0x21, 0xc5, 0xc9,
	0xda, 0x6e, 0x83, 0x7e, 0x31, 0xbe, 0xf7, 0xde, 0x9d, 0xde, 0x3d, 0x09, 0x42, 0x26, 0x2a, 0xc9,
	0x77, 0x25, 0x95, 0x2c, 0xd9, 0x6c, 0x95, 0x56, 0xf1, 0x2f, 0x07, 0xda, 0x73, 0x55, 0x20, 0x81,
	0xce, 0x66, 0xab, 0x6e, 0xf8, 0x57, 0x4d, 0x9c, 0xc8, 0x19, 0x07, 0xf9, 0xa1, 0xc4, 0x01, 0xb8,
	0x82, 0x11, 0x37, 0x72, 0xc6, 0x5e, 0xee, 0x0a, 0x86, 0x08, 0xde, 0x77, 0x21, 0x19, 0x69, 0x5b,
	0x99, 0xfd, 0xc7, 0x33, 0xf0, 0x2b, 0x4d, 0x75, 0x5d, 0x11, 0x2f, 0x72, 0xc6, 0x83, 0x49, 0x2f,
	0x99, 0xab, 0x22, 0xf9, 0x60, 0xa1, 0xbc, 0xa1, 0xf0, 0x39, 0x9c, 0x08, 0xb9, 0xa9, 0x35, 0x39,
	0xb1, 0x9d, 0xfb, 0x02, 0x5f, 0x80, 0xaf, 0x6a, 0x6d, 0x60, 0xdf, 0xc2, 0x4d, 0x85, 0x23, 0xe8,
	0xde, 0x72, 0x4d, 0x19, 0xd5, 0x94, 0x74, 0x2c, 0x73, 0xac, 0xe3, 0x19, 0xf8, 0xfb, 0xd9, 0xd8,
	0x83, 0x4e, 0x36, 0x5d, 0xa5, 0xb3, 0xd5, 0x45, 0xd8, 0x42, 0x00, 0x3f, 0xbb, 0x5a, 0x2c, 0xa6,
	0x69, 0xe8, 0x18, 0x22, 0xbf, 0x5a, 0xad, 0x0c, 0xe1, 0x1a, 0xe2, 0xfd, 0xbb, 0x99, 0x21, 0xda,
	0xd8, 0x87, 0xe0, 0xfc, 0x72, 0x99, 0x2d, 0xa6, 0x1f, 0xa7, 0x69, 0xe8, 0xc5, 0xaf, 0x00, 0x16,
	0xa2, 0xd2, 0x97, 0xd7, 0x73, 0x55, 0x54, 0x48, 0xc0, 0xbb, 0x51, 0x45, 0x45, 0x9c, 0xa8, 0x3d,
	0xee, 0x4d, 0x3c, 0xb3, 0x45, 0x6e, 0x91, 0xf8, 0x25, 0xf4, 0x73, 0xfe, 0xa3, 0xe6, 0x95, 0xfe,
	0x24, 0xf4, 0x7a, 0xc6, 0x9a, 0x58, 0x9c, 0x43, 0x2c, 0xf1, 0x67, 0x38, 0x35, 0x83, 0xcc, 0x98,
	0x46, 0x88, 0x43, 0xe8, 0xae, 0xd5, 0xdd, 0x97, 0x5b, 0x2a, 0x77, 0x56, 0xd8, 0xcf, 0x3b, 0x6b,
	0x75, 0xb7, 0xa4, 0x72, 0x77, 0x3f, 0x6e, 0xf7, 0x61, 0xdc, 0xff, 0x88, 0x77, 0xf2, 0xd3, 0x05,
	0x48, 0x8f, 0x37, 0x87, 0x43, 0x08, 0xce, 0xb7, 0x9c, 0x6a, 0x6e, 0x2e, 0xce, 0x9a, 0x1c, 0xd9,
	0x6f, 0xdc, 0xc2, 0x08, 0xfc, 0x0b, 0x6e, 0x4c, 0xe0, 0x20, 0x79, 0xe0, 0xf7, 0xa8, 0x78, 0x0d,
	0xdd, 0x83, 0x4f, 0x0c, 0x93, 0x47, 0x96, 0x47, 0xbd, 0xe4, 0x4f, 0x1a, 0x71, 0x0b, 0xdf, 0xc0,
	0xb3, 0xfd, 0x49, 0xcb, 0xba, 0xd4, 0x62, 0xc3, 0x6d, 0xd7, 0x7d, 0xcd, 0xe3, 0x86, 0x21, 0x04,
	0x4b, 0xc5, 0xc4, 0xf5, 0xee, 0x6f, 0x6b, 0x13, 0x38, 0xcd, 0xea, 0xb2, 0xcc, 0xb8, 0x64, 0x42,
	0x7e, 0x7b, 0xda, 0xf9, 0x67, 0x10, 0xa4, 0xbc, 0xe4, 0xfb, 0x4d, 0xff, 0xb3, 0x51, 0xe1, 0xdb,
	0x97, 0xfc, 0xf6, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd8, 0x14, 0x1f, 0xe5, 0xdd, 0x02, 0x00,
	0x00,
}

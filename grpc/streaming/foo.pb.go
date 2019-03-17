// Code generated by protoc-gen-go. DO NOT EDIT.
// source: foo.proto

package foo

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
//const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Data struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Data) Reset()         { *m = Data{} }
func (m *Data) String() string { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()    {}
func (*Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ce1e2eec643ca48, []int{0}
}

func (m *Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Data.Unmarshal(m, b)
}
func (m *Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Data.Marshal(b, m, deterministic)
}
func (m *Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Data.Merge(m, src)
}
func (m *Data) XXX_Size() int {
	return xxx_messageInfo_Data.Size(m)
}
func (m *Data) XXX_DiscardUnknown() {
	xxx_messageInfo_Data.DiscardUnknown(m)
}

var xxx_messageInfo_Data proto.InternalMessageInfo

func (m *Data) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Data)(nil), "foo.Data")
}

func init() { proto.RegisterFile("foo.proto", fileDescriptor_7ce1e2eec643ca48) }

var fileDescriptor_7ce1e2eec643ca48 = []byte{
	// 109 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0xcb, 0xcf, 0xd7,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0xcb, 0xcf, 0x57, 0x92, 0xe0, 0x62, 0x71, 0x49,
	0x2c, 0x49, 0x14, 0x12, 0xe0, 0x62, 0xce, 0x2d, 0x4e, 0x97, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c,
	0x02, 0x31, 0x8d, 0x8c, 0xb8, 0xb8, 0xdc, 0xf2, 0xf3, 0x83, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53,
	0x85, 0x54, 0xb8, 0xd8, 0xdc, 0xf2, 0xf3, 0x83, 0x02, 0x9c, 0x85, 0x38, 0xf5, 0x40, 0x46, 0x80,
	0x34, 0x49, 0x21, 0x98, 0x4a, 0x0c, 0x1a, 0x8c, 0x06, 0x8c, 0x49, 0x6c, 0x60, 0x93, 0x8d, 0x01,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x68, 0x17, 0xca, 0x55, 0x66, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FooServiceClient is the client API for FooService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FooServiceClient interface {
	FooRPC(ctx context.Context, opts ...grpc.CallOption) (FooService_FooRPCClient, error)
}

type fooServiceClient struct {
	cc *grpc.ClientConn
}

func NewFooServiceClient(cc *grpc.ClientConn) FooServiceClient {
	return &fooServiceClient{cc}
}

func (c *fooServiceClient) FooRPC(ctx context.Context, opts ...grpc.CallOption) (FooService_FooRPCClient, error) {
	stream, err := c.cc.NewStream(ctx, &_FooService_serviceDesc.Streams[0], "/foo.FooService/FooRPC", opts...)
	if err != nil {
		return nil, err
	}
	x := &fooServiceFooRPCClient{stream}
	return x, nil
}

type FooService_FooRPCClient interface {
	Send(*Data) error
	Recv() (*Data, error)
	grpc.ClientStream
}

type fooServiceFooRPCClient struct {
	grpc.ClientStream
}

func (x *fooServiceFooRPCClient) Send(m *Data) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fooServiceFooRPCClient) Recv() (*Data, error) {
	m := new(Data)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FooServiceServer is the server API for FooService service.
type FooServiceServer interface {
	FooRPC(FooService_FooRPCServer) error
}

func RegisterFooServiceServer(s *grpc.Server, srv FooServiceServer) {
	s.RegisterService(&_FooService_serviceDesc, srv)
}

func _FooService_FooRPC_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FooServiceServer).FooRPC(&fooServiceFooRPCServer{stream})
}

type FooService_FooRPCServer interface {
	Send(*Data) error
	Recv() (*Data, error)
	grpc.ServerStream
}

type fooServiceFooRPCServer struct {
	grpc.ServerStream
}

func (x *fooServiceFooRPCServer) Send(m *Data) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fooServiceFooRPCServer) Recv() (*Data, error) {
	m := new(Data)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _FooService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "foo.FooService",
	HandlerType: (*FooServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "FooRPC",
			Handler:       _FooService_FooRPC_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "foo.proto",
}

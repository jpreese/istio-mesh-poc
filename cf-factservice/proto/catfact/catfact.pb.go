// Code generated by protoc-gen-go. DO NOT EDIT.
// source: catfact.proto

package catfact

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type CatFactRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CatFactRequest) Reset()         { *m = CatFactRequest{} }
func (m *CatFactRequest) String() string { return proto.CompactTextString(m) }
func (*CatFactRequest) ProtoMessage()    {}
func (*CatFactRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd5f0c7b71a509a6, []int{0}
}

func (m *CatFactRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CatFactRequest.Unmarshal(m, b)
}
func (m *CatFactRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CatFactRequest.Marshal(b, m, deterministic)
}
func (m *CatFactRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CatFactRequest.Merge(m, src)
}
func (m *CatFactRequest) XXX_Size() int {
	return xxx_messageInfo_CatFactRequest.Size(m)
}
func (m *CatFactRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CatFactRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CatFactRequest proto.InternalMessageInfo

type CatFactResponse struct {
	Fact                 string   `protobuf:"bytes,1,opt,name=fact,proto3" json:"fact,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CatFactResponse) Reset()         { *m = CatFactResponse{} }
func (m *CatFactResponse) String() string { return proto.CompactTextString(m) }
func (*CatFactResponse) ProtoMessage()    {}
func (*CatFactResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd5f0c7b71a509a6, []int{1}
}

func (m *CatFactResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CatFactResponse.Unmarshal(m, b)
}
func (m *CatFactResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CatFactResponse.Marshal(b, m, deterministic)
}
func (m *CatFactResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CatFactResponse.Merge(m, src)
}
func (m *CatFactResponse) XXX_Size() int {
	return xxx_messageInfo_CatFactResponse.Size(m)
}
func (m *CatFactResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CatFactResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CatFactResponse proto.InternalMessageInfo

func (m *CatFactResponse) GetFact() string {
	if m != nil {
		return m.Fact
	}
	return ""
}

func init() {
	proto.RegisterType((*CatFactRequest)(nil), "CatFactRequest")
	proto.RegisterType((*CatFactResponse)(nil), "CatFactResponse")
}

func init() { proto.RegisterFile("catfact.proto", fileDescriptor_dd5f0c7b71a509a6) }

var fileDescriptor_dd5f0c7b71a509a6 = []byte{
	// 119 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x4e, 0x2c, 0x49,
	0x4b, 0x4c, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x12, 0xe0, 0xe2, 0x73, 0x4e, 0x2c,
	0x71, 0x4b, 0x4c, 0x2e, 0x09, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x51, 0x52, 0xe5, 0xe2, 0x87,
	0x8b, 0x14, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0x09, 0x71, 0xb1, 0x80, 0xb4, 0x48, 0x30, 0x2a,
	0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x46, 0x36, 0x70, 0x8d, 0xc1, 0xa9, 0x45, 0x65, 0x99, 0xc9,
	0xa9, 0x42, 0x5a, 0x5c, 0xcc, 0xee, 0xa9, 0x25, 0x42, 0xfc, 0x7a, 0xa8, 0x06, 0x4a, 0x09, 0xe8,
	0xa1, 0x99, 0xa7, 0xc4, 0x90, 0xc4, 0x06, 0xb6, 0xdd, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x9d,
	0xa5, 0x93, 0x11, 0x8e, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CatFactServiceClient is the client API for CatFactService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CatFactServiceClient interface {
	Get(ctx context.Context, in *CatFactRequest, opts ...grpc.CallOption) (*CatFactResponse, error)
}

type catFactServiceClient struct {
	cc *grpc.ClientConn
}

func NewCatFactServiceClient(cc *grpc.ClientConn) CatFactServiceClient {
	return &catFactServiceClient{cc}
}

func (c *catFactServiceClient) Get(ctx context.Context, in *CatFactRequest, opts ...grpc.CallOption) (*CatFactResponse, error) {
	out := new(CatFactResponse)
	err := c.cc.Invoke(ctx, "/CatFactService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CatFactServiceServer is the server API for CatFactService service.
type CatFactServiceServer interface {
	Get(context.Context, *CatFactRequest) (*CatFactResponse, error)
}

// UnimplementedCatFactServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCatFactServiceServer struct {
}

func (*UnimplementedCatFactServiceServer) Get(ctx context.Context, req *CatFactRequest) (*CatFactResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func RegisterCatFactServiceServer(s *grpc.Server, srv CatFactServiceServer) {
	s.RegisterService(&_CatFactService_serviceDesc, srv)
}

func _CatFactService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CatFactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatFactServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CatFactService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatFactServiceServer).Get(ctx, req.(*CatFactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CatFactService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "CatFactService",
	HandlerType: (*CatFactServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _CatFactService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "catfact.proto",
}

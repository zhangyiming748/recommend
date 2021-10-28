// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rank.proto

package protoFile

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

// Status 枚举状态
type RankReply_Status int32

const (
	RankReply_OK   RankReply_Status = 0
	RankReply_FAIL RankReply_Status = 1
)

var RankReply_Status_name = map[int32]string{
	0: "OK",
	1: "FAIL",
}

var RankReply_Status_value = map[string]int32{
	"OK":   0,
	"FAIL": 1,
}

func (x RankReply_Status) String() string {
	return proto.EnumName(RankReply_Status_name, int32(x))
}

func (RankReply_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_28127d302aca29e8, []int{1, 0}
}

//请求体
type RankRequest struct {
	Method               string            `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Param                map[string]string `protobuf:"bytes,2,rep,name=param,proto3" json:"param,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RankRequest) Reset()         { *m = RankRequest{} }
func (m *RankRequest) String() string { return proto.CompactTextString(m) }
func (*RankRequest) ProtoMessage()    {}
func (*RankRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_28127d302aca29e8, []int{0}
}

func (m *RankRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RankRequest.Unmarshal(m, b)
}
func (m *RankRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RankRequest.Marshal(b, m, deterministic)
}
func (m *RankRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RankRequest.Merge(m, src)
}
func (m *RankRequest) XXX_Size() int {
	return xxx_messageInfo_RankRequest.Size(m)
}
func (m *RankRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RankRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RankRequest proto.InternalMessageInfo

func (m *RankRequest) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *RankRequest) GetParam() map[string]string {
	if m != nil {
		return m.Param
	}
	return nil
}

//返回数据
type RankReply struct {
	Code                 RankReply_Status `protobuf:"varint,3,opt,name=code,proto3,enum=protoFile.RankReply_Status" json:"code,omitempty"`
	Message              string           `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	Data                 []*Article       `protobuf:"bytes,5,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *RankReply) Reset()         { *m = RankReply{} }
func (m *RankReply) String() string { return proto.CompactTextString(m) }
func (*RankReply) ProtoMessage()    {}
func (*RankReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_28127d302aca29e8, []int{1}
}

func (m *RankReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RankReply.Unmarshal(m, b)
}
func (m *RankReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RankReply.Marshal(b, m, deterministic)
}
func (m *RankReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RankReply.Merge(m, src)
}
func (m *RankReply) XXX_Size() int {
	return xxx_messageInfo_RankReply.Size(m)
}
func (m *RankReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RankReply.DiscardUnknown(m)
}

var xxx_messageInfo_RankReply proto.InternalMessageInfo

func (m *RankReply) GetCode() RankReply_Status {
	if m != nil {
		return m.Code
	}
	return RankReply_OK
}

func (m *RankReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *RankReply) GetData() []*Article {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterEnum("protoFile.RankReply_Status", RankReply_Status_name, RankReply_Status_value)
	proto.RegisterType((*RankRequest)(nil), "protoFile.RankRequest")
	proto.RegisterMapType((map[string]string)(nil), "protoFile.RankRequest.ParamEntry")
	proto.RegisterType((*RankReply)(nil), "protoFile.RankReply")
}

func init() { proto.RegisterFile("rank.proto", fileDescriptor_28127d302aca29e8) }

var fileDescriptor_28127d302aca29e8 = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x5f, 0x4b, 0xf3, 0x30,
	0x14, 0xc6, 0x97, 0xae, 0xeb, 0xfb, 0xee, 0x8c, 0xc9, 0x38, 0x8c, 0x11, 0xea, 0xcd, 0xec, 0x85,
	0xec, 0xaa, 0xc2, 0xbc, 0x70, 0x78, 0x23, 0x53, 0x1c, 0x88, 0xa2, 0xd2, 0xdd, 0x79, 0x17, 0xbb,
	0x83, 0x96, 0xfe, 0x49, 0x4d, 0xd3, 0x41, 0xbf, 0x89, 0xf8, 0x69, 0x65, 0xc9, 0x74, 0x03, 0x77,
	0x95, 0xf3, 0x9c, 0xf3, 0xe4, 0x97, 0xe4, 0x09, 0x80, 0x12, 0x45, 0x1a, 0x96, 0x4a, 0x6a, 0x89,
	0x5d, 0xb3, 0x2c, 0x92, 0x8c, 0xfc, 0xbe, 0x50, 0x3a, 0x89, 0x33, 0xb2, 0x93, 0xe0, 0x93, 0x41,
	0x2f, 0x12, 0x45, 0x1a, 0xd1, 0x47, 0x4d, 0x95, 0xc6, 0x11, 0x78, 0x39, 0xe9, 0x77, 0xb9, 0xe2,
	0x6c, 0xcc, 0x26, 0xdd, 0x68, 0xab, 0xf0, 0x02, 0x3a, 0xa5, 0x50, 0x22, 0xe7, 0xce, 0xb8, 0x3d,
	0xe9, 0x4d, 0x4f, 0xc2, 0x5f, 0x62, 0xb8, 0xb7, 0x3d, 0x7c, 0xde, 0x78, 0x6e, 0x0b, 0xad, 0x9a,
	0xc8, 0xfa, 0xfd, 0x19, 0xc0, 0xae, 0x89, 0x03, 0x68, 0xa7, 0xd4, 0x6c, 0xd9, 0x9b, 0x12, 0x87,
	0xd0, 0x59, 0x8b, 0xac, 0x26, 0xee, 0x98, 0x9e, 0x15, 0x97, 0xce, 0x8c, 0x05, 0x5f, 0x0c, 0xba,
	0x96, 0x5d, 0x66, 0x0d, 0x9e, 0x81, 0x1b, 0xcb, 0x15, 0xf1, 0xf6, 0x98, 0x4d, 0x8e, 0xa6, 0xc7,
	0x7f, 0xce, 0x2f, 0xb3, 0x26, 0x5c, 0x6a, 0xa1, 0xeb, 0x2a, 0x32, 0x46, 0xe4, 0xf0, 0x2f, 0xa7,
	0xaa, 0x12, 0x6f, 0xc4, 0x5d, 0x83, 0xfe, 0x91, 0x78, 0x0a, 0xee, 0x4a, 0x68, 0xc1, 0x3b, 0xe6,
	0x29, 0xb8, 0x87, 0x9a, 0xdb, 0x6c, 0x22, 0x33, 0x0f, 0x7c, 0xf0, 0x2c, 0x11, 0x3d, 0x70, 0x9e,
	0xee, 0x07, 0x2d, 0xfc, 0x0f, 0xee, 0x62, 0x7e, 0xf7, 0x30, 0x60, 0xd3, 0x47, 0x1b, 0xdb, 0x92,
	0xd4, 0x3a, 0x89, 0x09, 0xaf, 0xa0, 0x7f, 0x23, 0xf3, 0xbc, 0x2e, 0x92, 0x58, 0xe8, 0x44, 0x16,
	0x38, 0x3a, 0x1c, 0x90, 0x3f, 0x3c, 0x74, 0xf1, 0xa0, 0x75, 0xdd, 0x7b, 0xd9, 0xfd, 0xd1, 0xab,
	0x67, 0xca, 0xf3, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x12, 0x8c, 0x3e, 0x7a, 0xc3, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RankServiceClient is the client API for RankService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RankServiceClient interface {
	Communication(ctx context.Context, in *RankRequest, opts ...grpc.CallOption) (*RankReply, error)
}

type rankServiceClient struct {
	cc *grpc.ClientConn
}

func NewRankServiceClient(cc *grpc.ClientConn) RankServiceClient {
	return &rankServiceClient{cc}
}

func (c *rankServiceClient) Communication(ctx context.Context, in *RankRequest, opts ...grpc.CallOption) (*RankReply, error) {
	out := new(RankReply)
	err := c.cc.Invoke(ctx, "/protoFile.RankService/Communication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RankServiceServer is the server API for RankService service.
type RankServiceServer interface {
	Communication(context.Context, *RankRequest) (*RankReply, error)
}

func RegisterRankServiceServer(s *grpc.Server, srv RankServiceServer) {
	s.RegisterService(&_RankService_serviceDesc, srv)
}

func _RankService_Communication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RankRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RankServiceServer).Communication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protoFile.RankService/Communication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RankServiceServer).Communication(ctx, req.(*RankRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RankService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protoFile.RankService",
	HandlerType: (*RankServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Communication",
			Handler:    _RankService_Communication_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rank.proto",
}

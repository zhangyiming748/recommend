// Code generated by protoc-gen-go. DO NOT EDIT.
// source: usercf.proto

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
type UserCfReply_Status int32

const (
	UserCfReply_OK   UserCfReply_Status = 0
	UserCfReply_FAIL UserCfReply_Status = 1
)

var UserCfReply_Status_name = map[int32]string{
	0: "OK",
	1: "FAIL",
}

var UserCfReply_Status_value = map[string]int32{
	"OK":   0,
	"FAIL": 1,
}

func (x UserCfReply_Status) String() string {
	return proto.EnumName(UserCfReply_Status_name, int32(x))
}

func (UserCfReply_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3c14f35651f2db32, []int{1, 0}
}

// Status 枚举状态
type RecReply_Status int32

const (
	RecReply_OK   RecReply_Status = 0
	RecReply_FAIL RecReply_Status = 1
)

var RecReply_Status_name = map[int32]string{
	0: "OK",
	1: "FAIL",
}

var RecReply_Status_value = map[string]int32{
	"OK":   0,
	"FAIL": 1,
}

func (x RecReply_Status) String() string {
	return proto.EnumName(RecReply_Status_name, int32(x))
}

func (RecReply_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3c14f35651f2db32, []int{3, 0}
}

//请求体
type UserCfRequest struct {
	Method               string            `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Param                map[string]string `protobuf:"bytes,2,rep,name=param,proto3" json:"param,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *UserCfRequest) Reset()         { *m = UserCfRequest{} }
func (m *UserCfRequest) String() string { return proto.CompactTextString(m) }
func (*UserCfRequest) ProtoMessage()    {}
func (*UserCfRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c14f35651f2db32, []int{0}
}

func (m *UserCfRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserCfRequest.Unmarshal(m, b)
}
func (m *UserCfRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserCfRequest.Marshal(b, m, deterministic)
}
func (m *UserCfRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserCfRequest.Merge(m, src)
}
func (m *UserCfRequest) XXX_Size() int {
	return xxx_messageInfo_UserCfRequest.Size(m)
}
func (m *UserCfRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserCfRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserCfRequest proto.InternalMessageInfo

func (m *UserCfRequest) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *UserCfRequest) GetParam() map[string]string {
	if m != nil {
		return m.Param
	}
	return nil
}

//返回数据
type UserCfReply struct {
	Code                 UserCfReply_Status `protobuf:"varint,3,opt,name=code,proto3,enum=protoFile.UserCfReply_Status" json:"code,omitempty"`
	Message              string             `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	Data                 []*Article         `protobuf:"bytes,5,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *UserCfReply) Reset()         { *m = UserCfReply{} }
func (m *UserCfReply) String() string { return proto.CompactTextString(m) }
func (*UserCfReply) ProtoMessage()    {}
func (*UserCfReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c14f35651f2db32, []int{1}
}

func (m *UserCfReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserCfReply.Unmarshal(m, b)
}
func (m *UserCfReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserCfReply.Marshal(b, m, deterministic)
}
func (m *UserCfReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserCfReply.Merge(m, src)
}
func (m *UserCfReply) XXX_Size() int {
	return xxx_messageInfo_UserCfReply.Size(m)
}
func (m *UserCfReply) XXX_DiscardUnknown() {
	xxx_messageInfo_UserCfReply.DiscardUnknown(m)
}

var xxx_messageInfo_UserCfReply proto.InternalMessageInfo

func (m *UserCfReply) GetCode() UserCfReply_Status {
	if m != nil {
		return m.Code
	}
	return UserCfReply_OK
}

func (m *UserCfReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *UserCfReply) GetData() []*Article {
	if m != nil {
		return m.Data
	}
	return nil
}

type RecReqeust struct {
	Uid                  string          `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	K                    int32           `protobuf:"varint,2,opt,name=K,proto3" json:"K,omitempty"`
	N                    int32           `protobuf:"varint,3,opt,name=N,proto3" json:"N,omitempty"`
	FilterMap            map[string]bool `protobuf:"bytes,4,rep,name=filterMap,proto3" json:"filterMap,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *RecReqeust) Reset()         { *m = RecReqeust{} }
func (m *RecReqeust) String() string { return proto.CompactTextString(m) }
func (*RecReqeust) ProtoMessage()    {}
func (*RecReqeust) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c14f35651f2db32, []int{2}
}

func (m *RecReqeust) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecReqeust.Unmarshal(m, b)
}
func (m *RecReqeust) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecReqeust.Marshal(b, m, deterministic)
}
func (m *RecReqeust) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecReqeust.Merge(m, src)
}
func (m *RecReqeust) XXX_Size() int {
	return xxx_messageInfo_RecReqeust.Size(m)
}
func (m *RecReqeust) XXX_DiscardUnknown() {
	xxx_messageInfo_RecReqeust.DiscardUnknown(m)
}

var xxx_messageInfo_RecReqeust proto.InternalMessageInfo

func (m *RecReqeust) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *RecReqeust) GetK() int32 {
	if m != nil {
		return m.K
	}
	return 0
}

func (m *RecReqeust) GetN() int32 {
	if m != nil {
		return m.N
	}
	return 0
}

func (m *RecReqeust) GetFilterMap() map[string]bool {
	if m != nil {
		return m.FilterMap
	}
	return nil
}

type RecReply struct {
	Code                 RecReply_Status   `protobuf:"varint,5,opt,name=code,proto3,enum=protoFile.RecReply_Status" json:"code,omitempty"`
	Message              string            `protobuf:"bytes,6,opt,name=message,proto3" json:"message,omitempty"`
	Data                 map[string]string `protobuf:"bytes,7,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RecReply) Reset()         { *m = RecReply{} }
func (m *RecReply) String() string { return proto.CompactTextString(m) }
func (*RecReply) ProtoMessage()    {}
func (*RecReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c14f35651f2db32, []int{3}
}

func (m *RecReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecReply.Unmarshal(m, b)
}
func (m *RecReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecReply.Marshal(b, m, deterministic)
}
func (m *RecReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecReply.Merge(m, src)
}
func (m *RecReply) XXX_Size() int {
	return xxx_messageInfo_RecReply.Size(m)
}
func (m *RecReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RecReply.DiscardUnknown(m)
}

var xxx_messageInfo_RecReply proto.InternalMessageInfo

func (m *RecReply) GetCode() RecReply_Status {
	if m != nil {
		return m.Code
	}
	return RecReply_OK
}

func (m *RecReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *RecReply) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterEnum("protoFile.UserCfReply_Status", UserCfReply_Status_name, UserCfReply_Status_value)
	proto.RegisterEnum("protoFile.RecReply_Status", RecReply_Status_name, RecReply_Status_value)
	proto.RegisterType((*UserCfRequest)(nil), "protoFile.UserCfRequest")
	proto.RegisterMapType((map[string]string)(nil), "protoFile.UserCfRequest.ParamEntry")
	proto.RegisterType((*UserCfReply)(nil), "protoFile.UserCfReply")
	proto.RegisterType((*RecReqeust)(nil), "protoFile.RecReqeust")
	proto.RegisterMapType((map[string]bool)(nil), "protoFile.RecReqeust.FilterMapEntry")
	proto.RegisterType((*RecReply)(nil), "protoFile.RecReply")
	proto.RegisterMapType((map[string]string)(nil), "protoFile.RecReply.DataEntry")
}

func init() { proto.RegisterFile("usercf.proto", fileDescriptor_3c14f35651f2db32) }

var fileDescriptor_3c14f35651f2db32 = []byte{
	// 460 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xb3, 0x89, 0xed, 0xc6, 0x93, 0xa6, 0x8a, 0x06, 0xa8, 0x56, 0x96, 0x2a, 0x45, 0x06,
	0xa1, 0x9c, 0x2c, 0x35, 0x1c, 0x68, 0x11, 0x97, 0x36, 0x10, 0x09, 0x05, 0x0a, 0xda, 0x8a, 0x0b,
	0xb7, 0xc5, 0x99, 0x80, 0x85, 0x1d, 0xbb, 0xeb, 0x75, 0xa5, 0x3c, 0x05, 0x2f, 0x80, 0x78, 0x12,
	0x1e, 0x85, 0x87, 0x41, 0x5e, 0x27, 0x71, 0x82, 0x4c, 0x51, 0x4f, 0xd9, 0xf9, 0xb3, 0x5f, 0xe6,
	0xfb, 0xed, 0x18, 0x0e, 0x8b, 0x9c, 0x54, 0xb8, 0x08, 0x32, 0x95, 0xea, 0x14, 0x5d, 0xf3, 0x33,
	0x8d, 0x62, 0xf2, 0xfa, 0x52, 0xe9, 0x28, 0x8c, 0xa9, 0xaa, 0xf8, 0x3f, 0x18, 0xf4, 0x3f, 0xe6,
	0xa4, 0x26, 0x0b, 0x41, 0x37, 0x05, 0xe5, 0x1a, 0x8f, 0xc1, 0x49, 0x48, 0x7f, 0x4d, 0xe7, 0x9c,
	0x0d, 0xd9, 0xc8, 0x15, 0xeb, 0x08, 0xcf, 0xc1, 0xce, 0xa4, 0x92, 0x09, 0x6f, 0x0f, 0x3b, 0xa3,
	0xde, 0xf8, 0x71, 0xb0, 0xd5, 0x0c, 0xf6, 0x04, 0x82, 0x0f, 0x65, 0xd7, 0xeb, 0xa5, 0x56, 0x2b,
	0x51, 0xdd, 0xf0, 0xce, 0x00, 0xea, 0x24, 0x0e, 0xa0, 0xf3, 0x8d, 0x56, 0x6b, 0xf5, 0xf2, 0x88,
	0x0f, 0xc1, 0xbe, 0x95, 0x71, 0x41, 0xbc, 0x6d, 0x72, 0x55, 0xf0, 0xa2, 0x7d, 0xc6, 0xfc, 0x9f,
	0x0c, 0x7a, 0x1b, 0xf5, 0x2c, 0x5e, 0xe1, 0x29, 0x58, 0x61, 0x3a, 0x27, 0xde, 0x19, 0xb2, 0xd1,
	0xd1, 0xf8, 0xa4, 0x61, 0x86, 0x2c, 0x5e, 0x05, 0xd7, 0x5a, 0xea, 0x22, 0x17, 0xa6, 0x15, 0x39,
	0x1c, 0x24, 0x94, 0xe7, 0xf2, 0x0b, 0x71, 0xcb, 0xc8, 0x6f, 0x42, 0x7c, 0x0a, 0xd6, 0x5c, 0x6a,
	0xc9, 0x6d, 0x63, 0x08, 0x77, 0xc4, 0x2e, 0x2a, 0x46, 0xc2, 0xd4, 0x7d, 0x0f, 0x9c, 0x4a, 0x11,
	0x1d, 0x68, 0xbf, 0x9f, 0x0d, 0x5a, 0xd8, 0x05, 0x6b, 0x7a, 0xf1, 0xe6, 0xed, 0x80, 0xf9, 0xbf,
	0x18, 0x80, 0xa0, 0x50, 0xd0, 0x0d, 0x15, 0xb9, 0x2e, 0xbd, 0x15, 0xd1, 0x86, 0x5c, 0x79, 0xc4,
	0x43, 0x60, 0x33, 0xe3, 0xcb, 0x16, 0x6c, 0x56, 0x46, 0x57, 0x66, 0x78, 0x5b, 0xb0, 0x2b, 0xbc,
	0x04, 0x77, 0x11, 0xc5, 0x9a, 0xd4, 0x3b, 0x99, 0x71, 0xcb, 0x4c, 0xf1, 0x64, 0x67, 0x8a, 0x5a,
	0x37, 0x98, 0x6e, 0xda, 0x2a, 0xae, 0xf5, 0x35, 0xef, 0x25, 0x1c, 0xed, 0x17, 0xff, 0xc7, 0xb7,
	0xbb, 0xcb, 0xf7, 0x37, 0x83, 0xae, 0xf9, 0x9b, 0x12, 0x6e, 0xb0, 0x86, 0x6b, 0x1b, 0xb8, 0xde,
	0xdf, 0x93, 0xdc, 0x45, 0xd6, 0xd9, 0x27, 0x7b, 0xba, 0x26, 0x7b, 0x60, 0x3c, 0x9d, 0x34, 0x29,
	0xbd, 0x92, 0x5a, 0x56, 0x66, 0x4c, 0xab, 0xf7, 0x1c, 0xdc, 0x6d, 0xea, 0x5e, 0x2b, 0x72, 0xc7,
	0xeb, 0x8c, 0xbf, 0x6f, 0xb7, 0xfb, 0x9a, 0xd4, 0x6d, 0x14, 0x12, 0x4e, 0xa0, 0x3f, 0x49, 0x93,
	0xa4, 0x58, 0x46, 0xa1, 0xd4, 0x51, 0xba, 0x44, 0xfe, 0xaf, 0x3d, 0xf6, 0x8e, 0x9b, 0xb7, 0xcb,
	0x6f, 0xe1, 0x39, 0xb8, 0x82, 0xc2, 0x34, 0x49, 0x68, 0x39, 0xc7, 0x47, 0x8d, 0x2f, 0xe6, 0x3d,
	0x68, 0x30, 0xed, 0xb7, 0x2e, 0x7b, 0x9f, 0xea, 0x6f, 0xf1, 0xb3, 0x63, 0x8e, 0xcf, 0xfe, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x76, 0x80, 0x44, 0xfe, 0xad, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserCfServiceClient is the client API for UserCfService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserCfServiceClient interface {
	Communication(ctx context.Context, in *UserCfRequest, opts ...grpc.CallOption) (*UserCfReply, error)
	Recommend(ctx context.Context, in *RecReqeust, opts ...grpc.CallOption) (*RecReply, error)
}

type userCfServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserCfServiceClient(cc *grpc.ClientConn) UserCfServiceClient {
	return &userCfServiceClient{cc}
}

func (c *userCfServiceClient) Communication(ctx context.Context, in *UserCfRequest, opts ...grpc.CallOption) (*UserCfReply, error) {
	out := new(UserCfReply)
	err := c.cc.Invoke(ctx, "/protoFile.UserCfService/Communication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCfServiceClient) Recommend(ctx context.Context, in *RecReqeust, opts ...grpc.CallOption) (*RecReply, error) {
	out := new(RecReply)
	err := c.cc.Invoke(ctx, "/protoFile.UserCfService/Recommend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserCfServiceServer is the server API for UserCfService service.
type UserCfServiceServer interface {
	Communication(context.Context, *UserCfRequest) (*UserCfReply, error)
	Recommend(context.Context, *RecReqeust) (*RecReply, error)
}

func RegisterUserCfServiceServer(s *grpc.Server, srv UserCfServiceServer) {
	s.RegisterService(&_UserCfService_serviceDesc, srv)
}

func _UserCfService_Communication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserCfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCfServiceServer).Communication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protoFile.UserCfService/Communication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCfServiceServer).Communication(ctx, req.(*UserCfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCfService_Recommend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecReqeust)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCfServiceServer).Recommend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protoFile.UserCfService/Recommend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCfServiceServer).Recommend(ctx, req.(*RecReqeust))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserCfService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protoFile.UserCfService",
	HandlerType: (*UserCfServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Communication",
			Handler:    _UserCfService_Communication_Handler,
		},
		{
			MethodName: "Recommend",
			Handler:    _UserCfService_Recommend_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "usercf.proto",
}

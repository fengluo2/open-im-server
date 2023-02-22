// Code generated by protoc-gen-go. DO NOT EDIT.
// source: push/push.proto

package pbPush // import "Open_IM/pkg/proto/push"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import sdkws "Open_IM/pkg/proto/sdkws"

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

type PushMsgReq struct {
	MsgData              *sdkws.MsgData `protobuf:"bytes,1,opt,name=msgData" json:"msgData,omitempty"`
	SourceID             string         `protobuf:"bytes,2,opt,name=sourceID" json:"sourceID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *PushMsgReq) Reset()         { *m = PushMsgReq{} }
func (m *PushMsgReq) String() string { return proto.CompactTextString(m) }
func (*PushMsgReq) ProtoMessage()    {}
func (*PushMsgReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_push_534b6eabcdc561f2, []int{0}
}
func (m *PushMsgReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushMsgReq.Unmarshal(m, b)
}
func (m *PushMsgReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushMsgReq.Marshal(b, m, deterministic)
}
func (dst *PushMsgReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushMsgReq.Merge(dst, src)
}
func (m *PushMsgReq) XXX_Size() int {
	return xxx_messageInfo_PushMsgReq.Size(m)
}
func (m *PushMsgReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PushMsgReq.DiscardUnknown(m)
}

var xxx_messageInfo_PushMsgReq proto.InternalMessageInfo

func (m *PushMsgReq) GetMsgData() *sdkws.MsgData {
	if m != nil {
		return m.MsgData
	}
	return nil
}

func (m *PushMsgReq) GetSourceID() string {
	if m != nil {
		return m.SourceID
	}
	return ""
}

type PushMsgResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PushMsgResp) Reset()         { *m = PushMsgResp{} }
func (m *PushMsgResp) String() string { return proto.CompactTextString(m) }
func (*PushMsgResp) ProtoMessage()    {}
func (*PushMsgResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_push_534b6eabcdc561f2, []int{1}
}
func (m *PushMsgResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushMsgResp.Unmarshal(m, b)
}
func (m *PushMsgResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushMsgResp.Marshal(b, m, deterministic)
}
func (dst *PushMsgResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushMsgResp.Merge(dst, src)
}
func (m *PushMsgResp) XXX_Size() int {
	return xxx_messageInfo_PushMsgResp.Size(m)
}
func (m *PushMsgResp) XXX_DiscardUnknown() {
	xxx_messageInfo_PushMsgResp.DiscardUnknown(m)
}

var xxx_messageInfo_PushMsgResp proto.InternalMessageInfo

type DelUserPushTokenReq struct {
	UserID               string   `protobuf:"bytes,1,opt,name=userID" json:"userID,omitempty"`
	PlatformID           int32    `protobuf:"varint,2,opt,name=platformID" json:"platformID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DelUserPushTokenReq) Reset()         { *m = DelUserPushTokenReq{} }
func (m *DelUserPushTokenReq) String() string { return proto.CompactTextString(m) }
func (*DelUserPushTokenReq) ProtoMessage()    {}
func (*DelUserPushTokenReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_push_534b6eabcdc561f2, []int{2}
}
func (m *DelUserPushTokenReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelUserPushTokenReq.Unmarshal(m, b)
}
func (m *DelUserPushTokenReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelUserPushTokenReq.Marshal(b, m, deterministic)
}
func (dst *DelUserPushTokenReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelUserPushTokenReq.Merge(dst, src)
}
func (m *DelUserPushTokenReq) XXX_Size() int {
	return xxx_messageInfo_DelUserPushTokenReq.Size(m)
}
func (m *DelUserPushTokenReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DelUserPushTokenReq.DiscardUnknown(m)
}

var xxx_messageInfo_DelUserPushTokenReq proto.InternalMessageInfo

func (m *DelUserPushTokenReq) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *DelUserPushTokenReq) GetPlatformID() int32 {
	if m != nil {
		return m.PlatformID
	}
	return 0
}

type DelUserPushTokenResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DelUserPushTokenResp) Reset()         { *m = DelUserPushTokenResp{} }
func (m *DelUserPushTokenResp) String() string { return proto.CompactTextString(m) }
func (*DelUserPushTokenResp) ProtoMessage()    {}
func (*DelUserPushTokenResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_push_534b6eabcdc561f2, []int{3}
}
func (m *DelUserPushTokenResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelUserPushTokenResp.Unmarshal(m, b)
}
func (m *DelUserPushTokenResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelUserPushTokenResp.Marshal(b, m, deterministic)
}
func (dst *DelUserPushTokenResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelUserPushTokenResp.Merge(dst, src)
}
func (m *DelUserPushTokenResp) XXX_Size() int {
	return xxx_messageInfo_DelUserPushTokenResp.Size(m)
}
func (m *DelUserPushTokenResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DelUserPushTokenResp.DiscardUnknown(m)
}

var xxx_messageInfo_DelUserPushTokenResp proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PushMsgReq)(nil), "push.PushMsgReq")
	proto.RegisterType((*PushMsgResp)(nil), "push.PushMsgResp")
	proto.RegisterType((*DelUserPushTokenReq)(nil), "push.DelUserPushTokenReq")
	proto.RegisterType((*DelUserPushTokenResp)(nil), "push.DelUserPushTokenResp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for PushMsgService service

type PushMsgServiceClient interface {
	PushMsg(ctx context.Context, in *PushMsgReq, opts ...grpc.CallOption) (*PushMsgResp, error)
	DelUserPushToken(ctx context.Context, in *DelUserPushTokenReq, opts ...grpc.CallOption) (*DelUserPushTokenResp, error)
}

type pushMsgServiceClient struct {
	cc *grpc.ClientConn
}

func NewPushMsgServiceClient(cc *grpc.ClientConn) PushMsgServiceClient {
	return &pushMsgServiceClient{cc}
}

func (c *pushMsgServiceClient) PushMsg(ctx context.Context, in *PushMsgReq, opts ...grpc.CallOption) (*PushMsgResp, error) {
	out := new(PushMsgResp)
	err := grpc.Invoke(ctx, "/push.PushMsgService/PushMsg", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pushMsgServiceClient) DelUserPushToken(ctx context.Context, in *DelUserPushTokenReq, opts ...grpc.CallOption) (*DelUserPushTokenResp, error) {
	out := new(DelUserPushTokenResp)
	err := grpc.Invoke(ctx, "/push.PushMsgService/DelUserPushToken", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PushMsgService service

type PushMsgServiceServer interface {
	PushMsg(context.Context, *PushMsgReq) (*PushMsgResp, error)
	DelUserPushToken(context.Context, *DelUserPushTokenReq) (*DelUserPushTokenResp, error)
}

func RegisterPushMsgServiceServer(s *grpc.Server, srv PushMsgServiceServer) {
	s.RegisterService(&_PushMsgService_serviceDesc, srv)
}

func _PushMsgService_PushMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushMsgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PushMsgServiceServer).PushMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/push.PushMsgService/PushMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PushMsgServiceServer).PushMsg(ctx, req.(*PushMsgReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PushMsgService_DelUserPushToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelUserPushTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PushMsgServiceServer).DelUserPushToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/push.PushMsgService/DelUserPushToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PushMsgServiceServer).DelUserPushToken(ctx, req.(*DelUserPushTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _PushMsgService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "push.PushMsgService",
	HandlerType: (*PushMsgServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PushMsg",
			Handler:    _PushMsgService_PushMsg_Handler,
		},
		{
			MethodName: "DelUserPushToken",
			Handler:    _PushMsgService_DelUserPushToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "push/push.proto",
}

func init() { proto.RegisterFile("push/push.proto", fileDescriptor_push_534b6eabcdc561f2) }

var fileDescriptor_push_534b6eabcdc561f2 = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x41, 0x4f, 0xc2, 0x40,
	0x10, 0x85, 0x53, 0xa3, 0x20, 0x43, 0x44, 0x5c, 0x0d, 0xc1, 0x26, 0x2a, 0xe9, 0x45, 0x2e, 0xb4,
	0x09, 0x1e, 0xbd, 0x99, 0x5e, 0x7a, 0xd8, 0x68, 0x56, 0xbd, 0x78, 0x31, 0x05, 0xc7, 0x62, 0x0a,
	0xec, 0xb8, 0xd3, 0xca, 0x8f, 0xf0, 0x4f, 0x9b, 0xdd, 0x56, 0x24, 0xa8, 0x97, 0xa6, 0xf3, 0xcd,
	0xcb, 0x7b, 0x79, 0xb3, 0x70, 0x48, 0x25, 0xcf, 0x22, 0xfb, 0x09, 0xc9, 0xe8, 0x42, 0x8b, 0x5d,
	0xfb, 0xef, 0x5f, 0xde, 0x12, 0x2e, 0x47, 0x89, 0x1c, 0xdd, 0xa3, 0xf9, 0x40, 0x13, 0x51, 0x9e,
	0x45, 0x6e, 0x1f, 0xf1, 0x4b, 0xbe, 0xe2, 0x68, 0xc5, 0x95, 0x3c, 0x50, 0x00, 0x77, 0x25, 0xcf,
	0x24, 0x67, 0x0a, 0xdf, 0xc5, 0x10, 0x9a, 0x0b, 0xce, 0xe2, 0xb4, 0x48, 0xfb, 0xde, 0xc0, 0x1b,
	0xb6, 0xc7, 0x9d, 0xd0, 0xe9, 0x43, 0x59, 0x51, 0xf5, 0xbd, 0x16, 0x3e, 0xec, 0xb3, 0x2e, 0xcd,
	0x14, 0x93, 0xb8, 0xbf, 0x33, 0xf0, 0x86, 0x2d, 0xb5, 0x9e, 0x83, 0x03, 0x68, 0xaf, 0x3d, 0x99,
	0x02, 0x09, 0xc7, 0x31, 0xce, 0x1f, 0x19, 0x8d, 0xa5, 0x0f, 0x3a, 0xc7, 0xa5, 0xcd, 0xea, 0x41,
	0xa3, 0x64, 0x34, 0x49, 0xec, 0xa2, 0x5a, 0xaa, 0x9e, 0xc4, 0x39, 0x00, 0xcd, 0xd3, 0xe2, 0x55,
	0x9b, 0x45, 0xed, 0xbd, 0xa7, 0x36, 0x48, 0xd0, 0x83, 0x93, 0xdf, 0x76, 0x4c, 0xe3, 0x4f, 0x0f,
	0x3a, 0x75, 0xac, 0x2d, 0xfd, 0x36, 0x45, 0x11, 0x42, 0xb3, 0x26, 0xa2, 0x1b, 0xba, 0x1b, 0xfd,
	0x74, 0xf5, 0x8f, 0xb6, 0x08, 0x93, 0x48, 0xa0, 0xbb, 0x6d, 0x2d, 0x4e, 0x2b, 0xd9, 0x1f, 0x0d,
	0x7c, 0xff, 0xbf, 0x15, 0xd3, 0xcd, 0xc5, 0xd3, 0x99, 0x7d, 0x82, 0xe7, 0x44, 0x6e, 0xdc, 0xde,
	0xca, 0xaf, 0x69, 0x62, 0x95, 0x93, 0x86, 0x43, 0x57, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x04,
	0x7d, 0xb3, 0xcd, 0xc1, 0x01, 0x00, 0x00,
}

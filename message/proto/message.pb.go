// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	message.proto

It has these top-level messages:
	MsgFilter
	FilterResp
	MsgSaveReq
	MsgSaveResp
	DelSetReq
	DelSetResp
	RegCliReq
	RegCliResp
	DelCliReq
	DelCliResp
	SerType
	AddrResp
	MsgPingReq
	MsgPingResp
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type MsgFilter struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
}

func (m *MsgFilter) Reset()                    { *m = MsgFilter{} }
func (m *MsgFilter) String() string            { return proto1.CompactTextString(m) }
func (*MsgFilter) ProtoMessage()               {}
func (*MsgFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *MsgFilter) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type FilterResp struct {
	Content []*MsgSaveReq `protobuf:"bytes,1,rep,name=content" json:"content,omitempty"`
}

func (m *FilterResp) Reset()                    { *m = FilterResp{} }
func (m *FilterResp) String() string            { return proto1.CompactTextString(m) }
func (*FilterResp) ProtoMessage()               {}
func (*FilterResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *FilterResp) GetContent() []*MsgSaveReq {
	if m != nil {
		return m.Content
	}
	return nil
}

type MsgSaveReq struct {
	DevId    string `protobuf:"bytes,1,opt,name=devId" json:"devId,omitempty"`
	Expire   string `protobuf:"bytes,2,opt,name=expire" json:"expire,omitempty"`
	Platform string `protobuf:"bytes,3,opt,name=platform" json:"platform,omitempty"`
	Phone    string `protobuf:"bytes,4,opt,name=phone" json:"phone,omitempty"`
	Msg      string `protobuf:"bytes,5,opt,name=msg" json:"msg,omitempty"`
	PushId   string `protobuf:"bytes,6,opt,name=pushId" json:"pushId,omitempty"`
	Callback string `protobuf:"bytes,7,opt,name=callback" json:"callback,omitempty"`
	Code     string `protobuf:"bytes,8,opt,name=code" json:"code,omitempty"`
}

func (m *MsgSaveReq) Reset()                    { *m = MsgSaveReq{} }
func (m *MsgSaveReq) String() string            { return proto1.CompactTextString(m) }
func (*MsgSaveReq) ProtoMessage()               {}
func (*MsgSaveReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *MsgSaveReq) GetDevId() string {
	if m != nil {
		return m.DevId
	}
	return ""
}

func (m *MsgSaveReq) GetExpire() string {
	if m != nil {
		return m.Expire
	}
	return ""
}

func (m *MsgSaveReq) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

func (m *MsgSaveReq) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *MsgSaveReq) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *MsgSaveReq) GetPushId() string {
	if m != nil {
		return m.PushId
	}
	return ""
}

func (m *MsgSaveReq) GetCallback() string {
	if m != nil {
		return m.Callback
	}
	return ""
}

func (m *MsgSaveReq) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type MsgSaveResp struct {
}

func (m *MsgSaveResp) Reset()                    { *m = MsgSaveResp{} }
func (m *MsgSaveResp) String() string            { return proto1.CompactTextString(m) }
func (*MsgSaveResp) ProtoMessage()               {}
func (*MsgSaveResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type DelSetReq struct {
	Key    string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Member string `protobuf:"bytes,2,opt,name=member" json:"member,omitempty"`
}

func (m *DelSetReq) Reset()                    { *m = DelSetReq{} }
func (m *DelSetReq) String() string            { return proto1.CompactTextString(m) }
func (*DelSetReq) ProtoMessage()               {}
func (*DelSetReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *DelSetReq) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *DelSetReq) GetMember() string {
	if m != nil {
		return m.Member
	}
	return ""
}

type DelSetResp struct {
}

func (m *DelSetResp) Reset()                    { *m = DelSetResp{} }
func (m *DelSetResp) String() string            { return proto1.CompactTextString(m) }
func (*DelSetResp) ProtoMessage()               {}
func (*DelSetResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type RegCliReq struct {
	Key        string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	ConectAddr string `protobuf:"bytes,2,opt,name=conectAddr" json:"conectAddr,omitempty"`
}

func (m *RegCliReq) Reset()                    { *m = RegCliReq{} }
func (m *RegCliReq) String() string            { return proto1.CompactTextString(m) }
func (*RegCliReq) ProtoMessage()               {}
func (*RegCliReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *RegCliReq) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *RegCliReq) GetConectAddr() string {
	if m != nil {
		return m.ConectAddr
	}
	return ""
}

type RegCliResp struct {
}

func (m *RegCliResp) Reset()                    { *m = RegCliResp{} }
func (m *RegCliResp) String() string            { return proto1.CompactTextString(m) }
func (*RegCliResp) ProtoMessage()               {}
func (*RegCliResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type DelCliReq struct {
	Key        string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	ConectAddr string `protobuf:"bytes,2,opt,name=conectAddr" json:"conectAddr,omitempty"`
}

func (m *DelCliReq) Reset()                    { *m = DelCliReq{} }
func (m *DelCliReq) String() string            { return proto1.CompactTextString(m) }
func (*DelCliReq) ProtoMessage()               {}
func (*DelCliReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *DelCliReq) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *DelCliReq) GetConectAddr() string {
	if m != nil {
		return m.ConectAddr
	}
	return ""
}

type DelCliResp struct {
}

func (m *DelCliResp) Reset()                    { *m = DelCliResp{} }
func (m *DelCliResp) String() string            { return proto1.CompactTextString(m) }
func (*DelCliResp) ProtoMessage()               {}
func (*DelCliResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type SerType struct {
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
}

func (m *SerType) Reset()                    { *m = SerType{} }
func (m *SerType) String() string            { return proto1.CompactTextString(m) }
func (*SerType) ProtoMessage()               {}
func (*SerType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *SerType) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type AddrResp struct {
	Addrs []string `protobuf:"bytes,1,rep,name=addrs" json:"addrs,omitempty"`
}

func (m *AddrResp) Reset()                    { *m = AddrResp{} }
func (m *AddrResp) String() string            { return proto1.CompactTextString(m) }
func (*AddrResp) ProtoMessage()               {}
func (*AddrResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *AddrResp) GetAddrs() []string {
	if m != nil {
		return m.Addrs
	}
	return nil
}

type MsgPingReq struct {
}

func (m *MsgPingReq) Reset()                    { *m = MsgPingReq{} }
func (m *MsgPingReq) String() string            { return proto1.CompactTextString(m) }
func (*MsgPingReq) ProtoMessage()               {}
func (*MsgPingReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

type MsgPingResp struct {
}

func (m *MsgPingResp) Reset()                    { *m = MsgPingResp{} }
func (m *MsgPingResp) String() string            { return proto1.CompactTextString(m) }
func (*MsgPingResp) ProtoMessage()               {}
func (*MsgPingResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func init() {
	proto1.RegisterType((*MsgFilter)(nil), "proto.MsgFilter")
	proto1.RegisterType((*FilterResp)(nil), "proto.FilterResp")
	proto1.RegisterType((*MsgSaveReq)(nil), "proto.MsgSaveReq")
	proto1.RegisterType((*MsgSaveResp)(nil), "proto.MsgSaveResp")
	proto1.RegisterType((*DelSetReq)(nil), "proto.DelSetReq")
	proto1.RegisterType((*DelSetResp)(nil), "proto.DelSetResp")
	proto1.RegisterType((*RegCliReq)(nil), "proto.RegCliReq")
	proto1.RegisterType((*RegCliResp)(nil), "proto.RegCliResp")
	proto1.RegisterType((*DelCliReq)(nil), "proto.DelCliReq")
	proto1.RegisterType((*DelCliResp)(nil), "proto.DelCliResp")
	proto1.RegisterType((*SerType)(nil), "proto.SerType")
	proto1.RegisterType((*AddrResp)(nil), "proto.AddrResp")
	proto1.RegisterType((*MsgPingReq)(nil), "proto.MsgPingReq")
	proto1.RegisterType((*MsgPingResp)(nil), "proto.MsgPingResp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for MsgSaveService service

type MsgSaveServiceClient interface {
	GetMsgByKey(ctx context.Context, in *MsgFilter, opts ...grpc.CallOption) (*FilterResp, error)
	MsgSave(ctx context.Context, in *MsgSaveReq, opts ...grpc.CallOption) (*MsgSaveResp, error)
	DelSetByMem(ctx context.Context, in *DelSetReq, opts ...grpc.CallOption) (*DelSetResp, error)
	RegisterCli(ctx context.Context, in *RegCliReq, opts ...grpc.CallOption) (*RegCliResp, error)
	DelCli(ctx context.Context, in *DelCliReq, opts ...grpc.CallOption) (*DelCliResp, error)
	GetSerAddr(ctx context.Context, in *SerType, opts ...grpc.CallOption) (*AddrResp, error)
	Ping(ctx context.Context, in *MsgPingReq, opts ...grpc.CallOption) (*MsgPingResp, error)
}

type msgSaveServiceClient struct {
	cc *grpc.ClientConn
}

func NewMsgSaveServiceClient(cc *grpc.ClientConn) MsgSaveServiceClient {
	return &msgSaveServiceClient{cc}
}

func (c *msgSaveServiceClient) GetMsgByKey(ctx context.Context, in *MsgFilter, opts ...grpc.CallOption) (*FilterResp, error) {
	out := new(FilterResp)
	err := grpc.Invoke(ctx, "/proto.msgSaveService/GetMsgByKey", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgSaveServiceClient) MsgSave(ctx context.Context, in *MsgSaveReq, opts ...grpc.CallOption) (*MsgSaveResp, error) {
	out := new(MsgSaveResp)
	err := grpc.Invoke(ctx, "/proto.msgSaveService/MsgSave", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgSaveServiceClient) DelSetByMem(ctx context.Context, in *DelSetReq, opts ...grpc.CallOption) (*DelSetResp, error) {
	out := new(DelSetResp)
	err := grpc.Invoke(ctx, "/proto.msgSaveService/DelSetByMem", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgSaveServiceClient) RegisterCli(ctx context.Context, in *RegCliReq, opts ...grpc.CallOption) (*RegCliResp, error) {
	out := new(RegCliResp)
	err := grpc.Invoke(ctx, "/proto.msgSaveService/RegisterCli", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgSaveServiceClient) DelCli(ctx context.Context, in *DelCliReq, opts ...grpc.CallOption) (*DelCliResp, error) {
	out := new(DelCliResp)
	err := grpc.Invoke(ctx, "/proto.msgSaveService/DelCli", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgSaveServiceClient) GetSerAddr(ctx context.Context, in *SerType, opts ...grpc.CallOption) (*AddrResp, error) {
	out := new(AddrResp)
	err := grpc.Invoke(ctx, "/proto.msgSaveService/GetSerAddr", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgSaveServiceClient) Ping(ctx context.Context, in *MsgPingReq, opts ...grpc.CallOption) (*MsgPingResp, error) {
	out := new(MsgPingResp)
	err := grpc.Invoke(ctx, "/proto.msgSaveService/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MsgSaveService service

type MsgSaveServiceServer interface {
	GetMsgByKey(context.Context, *MsgFilter) (*FilterResp, error)
	MsgSave(context.Context, *MsgSaveReq) (*MsgSaveResp, error)
	DelSetByMem(context.Context, *DelSetReq) (*DelSetResp, error)
	RegisterCli(context.Context, *RegCliReq) (*RegCliResp, error)
	DelCli(context.Context, *DelCliReq) (*DelCliResp, error)
	GetSerAddr(context.Context, *SerType) (*AddrResp, error)
	Ping(context.Context, *MsgPingReq) (*MsgPingResp, error)
}

func RegisterMsgSaveServiceServer(s *grpc.Server, srv MsgSaveServiceServer) {
	s.RegisterService(&_MsgSaveService_serviceDesc, srv)
}

func _MsgSaveService_GetMsgByKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgSaveServiceServer).GetMsgByKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.msgSaveService/GetMsgByKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgSaveServiceServer).GetMsgByKey(ctx, req.(*MsgFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgSaveService_MsgSave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSaveReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgSaveServiceServer).MsgSave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.msgSaveService/MsgSave",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgSaveServiceServer).MsgSave(ctx, req.(*MsgSaveReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgSaveService_DelSetByMem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelSetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgSaveServiceServer).DelSetByMem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.msgSaveService/DelSetByMem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgSaveServiceServer).DelSetByMem(ctx, req.(*DelSetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgSaveService_RegisterCli_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegCliReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgSaveServiceServer).RegisterCli(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.msgSaveService/RegisterCli",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgSaveServiceServer).RegisterCli(ctx, req.(*RegCliReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgSaveService_DelCli_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelCliReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgSaveServiceServer).DelCli(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.msgSaveService/DelCli",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgSaveServiceServer).DelCli(ctx, req.(*DelCliReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgSaveService_GetSerAddr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SerType)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgSaveServiceServer).GetSerAddr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.msgSaveService/GetSerAddr",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgSaveServiceServer).GetSerAddr(ctx, req.(*SerType))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgSaveService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgPingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgSaveServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.msgSaveService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgSaveServiceServer).Ping(ctx, req.(*MsgPingReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _MsgSaveService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.msgSaveService",
	HandlerType: (*MsgSaveServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMsgByKey",
			Handler:    _MsgSaveService_GetMsgByKey_Handler,
		},
		{
			MethodName: "MsgSave",
			Handler:    _MsgSaveService_MsgSave_Handler,
		},
		{
			MethodName: "DelSetByMem",
			Handler:    _MsgSaveService_DelSetByMem_Handler,
		},
		{
			MethodName: "RegisterCli",
			Handler:    _MsgSaveService_RegisterCli_Handler,
		},
		{
			MethodName: "DelCli",
			Handler:    _MsgSaveService_DelCli_Handler,
		},
		{
			MethodName: "GetSerAddr",
			Handler:    _MsgSaveService_GetSerAddr_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _MsgSaveService_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.proto",
}

func init() { proto1.RegisterFile("message.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 494 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x5d, 0x6b, 0xd4, 0x40,
	0x14, 0x86, 0xbb, 0xdd, 0xcf, 0x9c, 0xb5, 0xb5, 0x0e, 0x45, 0x86, 0x85, 0xca, 0x32, 0x57, 0x05,
	0xa1, 0xc2, 0x5a, 0x2f, 0xbc, 0xf0, 0xc2, 0xb6, 0x58, 0x8a, 0x04, 0x24, 0xf1, 0x0f, 0x64, 0x93,
	0xe3, 0x34, 0x34, 0x1f, 0xd3, 0x99, 0x71, 0x31, 0x7f, 0xd0, 0xff, 0xe4, 0x9d, 0xcc, 0x57, 0x36,
	0xb4, 0xde, 0x79, 0x95, 0x79, 0x4f, 0xf2, 0x3e, 0x93, 0xf3, 0x9e, 0x03, 0x47, 0x35, 0x2a, 0x95,
	0x71, 0xbc, 0x10, 0xb2, 0xd5, 0x2d, 0x99, 0xda, 0x07, 0x3b, 0x83, 0x28, 0x56, 0xfc, 0x4b, 0x59,
	0x69, 0x94, 0xe4, 0x04, 0xc6, 0x0f, 0xd8, 0xd1, 0xd1, 0x7a, 0x74, 0x1e, 0x25, 0xe6, 0xc8, 0x3e,
	0x02, 0xb8, 0x77, 0x09, 0x2a, 0x41, 0xde, 0xc2, 0x3c, 0x6f, 0x1b, 0x8d, 0x8d, 0xa6, 0xa3, 0xf5,
	0xf8, 0x7c, 0xb9, 0x79, 0xe5, 0x60, 0x17, 0xb1, 0xe2, 0x69, 0xb6, 0xc3, 0x04, 0x1f, 0x93, 0xf0,
	0x05, 0xfb, 0x3d, 0x02, 0xd8, 0xd7, 0xc9, 0x29, 0x4c, 0x0b, 0xdc, 0xdd, 0x15, 0x9e, 0xee, 0x04,
	0x79, 0x0d, 0x33, 0xfc, 0x25, 0x4a, 0x89, 0xf4, 0xd0, 0x96, 0xbd, 0x22, 0x2b, 0x58, 0x88, 0x2a,
	0xd3, 0x3f, 0x5a, 0x59, 0xd3, 0xb1, 0x7d, 0xd3, 0x6b, 0x43, 0x12, 0xf7, 0x6d, 0x83, 0x74, 0xe2,
	0x48, 0x56, 0x98, 0x7f, 0xaf, 0x15, 0xa7, 0x53, 0xf7, 0xef, 0xb5, 0xe2, 0x86, 0x2d, 0x7e, 0xaa,
	0xfb, 0xbb, 0x82, 0xce, 0x1c, 0xdb, 0x29, 0xc3, 0xce, 0xb3, 0xaa, 0xda, 0x66, 0xf9, 0x03, 0x9d,
	0x3b, 0x76, 0xd0, 0x84, 0xc0, 0x24, 0x6f, 0x0b, 0xa4, 0x0b, 0x5b, 0xb7, 0x67, 0x76, 0x04, 0xcb,
	0xbe, 0x0f, 0x25, 0xd8, 0x07, 0x88, 0x6e, 0xb0, 0x4a, 0x51, 0x9b, 0xae, 0x9e, 0x25, 0x66, 0x6e,
	0xad, 0xb1, 0xde, 0xa2, 0x0c, 0x1d, 0x39, 0xc5, 0x5e, 0x00, 0x04, 0x9b, 0x12, 0xec, 0x13, 0x44,
	0x09, 0xf2, 0xeb, 0xaa, 0xfc, 0x37, 0xe4, 0x0d, 0x40, 0xde, 0x36, 0x98, 0xeb, 0xcf, 0x45, 0x11,
	0x40, 0x83, 0x8a, 0x81, 0x05, 0xbb, 0x83, 0xdd, 0x60, 0xf5, 0x3f, 0xb0, 0x60, 0x57, 0x82, 0x9d,
	0xc1, 0x3c, 0x45, 0xf9, 0xbd, 0x13, 0x68, 0xc2, 0xd0, 0x9d, 0x40, 0xcf, 0xb2, 0x67, 0xb6, 0x86,
	0x85, 0x31, 0xd9, 0x75, 0x38, 0x85, 0x69, 0x56, 0x14, 0x52, 0xd9, 0x65, 0x88, 0x12, 0x27, 0x0c,
	0x2e, 0x56, 0xfc, 0x5b, 0xd9, 0xf0, 0x04, 0x1f, 0x7d, 0x78, 0x4e, 0x29, 0xb1, 0xf9, 0x73, 0x08,
	0xc7, 0xb5, 0x0b, 0x33, 0x45, 0xb9, 0x2b, 0x73, 0x24, 0x97, 0xb0, 0xbc, 0x45, 0x1d, 0x2b, 0x7e,
	0xd5, 0x7d, 0xc5, 0x8e, 0x9c, 0xec, 0x57, 0xca, 0x6d, 0xde, 0x2a, 0x2c, 0xd9, 0x7e, 0x11, 0xd9,
	0x01, 0xd9, 0xc0, 0xdc, 0x0f, 0x85, 0x3c, 0x5f, 0xc2, 0x15, 0x79, 0x5a, 0xb2, 0x9e, 0x4b, 0x58,
	0xba, 0x11, 0x5c, 0x75, 0x31, 0xd6, 0xfd, 0x4d, 0xfd, 0x34, 0xfb, 0x9b, 0x06, 0x83, 0xb2, 0xae,
	0x04, 0x79, 0xa9, 0x34, 0xca, 0xeb, 0xaa, 0xec, 0x5d, 0xfd, 0xf8, 0x7a, 0xd7, 0x60, 0x22, 0x07,
	0xe4, 0x1d, 0xcc, 0x5c, 0xa8, 0xc3, 0x6b, 0x9e, 0x18, 0x06, 0xa9, 0x1b, 0x03, 0xdc, 0xa2, 0x4e,
	0x51, 0x9a, 0x78, 0xc9, 0xb1, 0xff, 0xc4, 0x8f, 0x62, 0xf5, 0xd2, 0xeb, 0x90, 0xbd, 0x35, 0x4c,
	0x4c, 0xac, 0xc3, 0xf6, 0x7d, 0xe8, 0xc3, 0xf6, 0x43, 0xf2, 0xec, 0x60, 0x3b, 0xb3, 0xc5, 0xf7,
	0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x97, 0x71, 0x56, 0x46, 0x09, 0x04, 0x00, 0x00,
}

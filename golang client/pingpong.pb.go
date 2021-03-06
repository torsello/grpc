// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pingpong.proto

package pingpong

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

type Request struct {
	Cadena               string   `protobuf:"bytes,1,opt,name=cadena,proto3" json:"cadena,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cfbf639ab46154b, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetCadena() string {
	if m != nil {
		return m.Cadena
	}
	return ""
}

type Reply struct {
	Cadena               string   `protobuf:"bytes,1,opt,name=cadena,proto3" json:"cadena,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Reply) Reset()         { *m = Reply{} }
func (m *Reply) String() string { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()    {}
func (*Reply) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cfbf639ab46154b, []int{1}
}

func (m *Reply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Reply.Unmarshal(m, b)
}
func (m *Reply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Reply.Marshal(b, m, deterministic)
}
func (m *Reply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Reply.Merge(m, src)
}
func (m *Reply) XXX_Size() int {
	return xxx_messageInfo_Reply.Size(m)
}
func (m *Reply) XXX_DiscardUnknown() {
	xxx_messageInfo_Reply.DiscardUnknown(m)
}

var xxx_messageInfo_Reply proto.InternalMessageInfo

func (m *Reply) GetCadena() string {
	if m != nil {
		return m.Cadena
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "Request")
	proto.RegisterType((*Reply)(nil), "Reply")
}

func init() { proto.RegisterFile("pingpong.proto", fileDescriptor_1cfbf639ab46154b) }

var fileDescriptor_1cfbf639ab46154b = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0xd0, 0xbf, 0x6a, 0xc3, 0x30,
	0x10, 0xc7, 0xf1, 0x7a, 0xa8, 0xdb, 0x5e, 0x4b, 0x4d, 0x55, 0x28, 0xa5, 0x4b, 0x6b, 0x81, 0x69,
	0x27, 0x0f, 0xed, 0x13, 0x94, 0xcc, 0x59, 0x9c, 0x3f, 0xfb, 0x45, 0x3e, 0x8c, 0x88, 0x23, 0x29,
	0x27, 0x29, 0x90, 0x27, 0xcf, 0x1a, 0x50, 0x3c, 0x3a, 0x1e, 0x0f, 0x3e, 0x7c, 0x39, 0x7e, 0xf0,
	0xec, 0xb4, 0xe9, 0x9c, 0x35, 0x5d, 0xed, 0xd8, 0x06, 0x2b, 0x4b, 0xb8, 0x6b, 0x68, 0x1f, 0xc9,
	0x07, 0xf1, 0x06, 0xb9, 0xc2, 0x96, 0x0c, 0xbe, 0x67, 0x5f, 0xd9, 0xcf, 0x43, 0x33, 0x5c, 0xf2,
	0x13, 0x6e, 0x1b, 0x72, 0xfd, 0xf1, 0x1a, 0xf8, 0x3d, 0x65, 0xf0, 0x3a, 0x44, 0x12, 0x5c, 0x10,
	0x1f, 0xb4, 0x22, 0x51, 0x41, 0x11, 0x18, 0x89, 0xff, 0x89, 0xad, 0x8b, 0xc4, 0xc1, 0x8a, 0xfb,
	0x7a, 0x80, 0x1f, 0x79, 0x9d, 0xac, 0xbc, 0x11, 0x12, 0x9e, 0x12, 0x5b, 0xf9, 0x88, 0xac, 0xc7,
	0x4d, 0x09, 0x8f, 0xc9, 0xac, 0x23, 0xf5, 0xd6, 0x4f, 0x66, 0x96, 0x5a, 0x6d, 0x29, 0x8c, 0x9b,
	0x0a, 0x0a, 0x85, 0x26, 0xe8, 0x16, 0xdb, 0x29, 0xf6, 0x0d, 0x2f, 0xf1, 0xf2, 0xcc, 0x1c, 0xfd,
	0xcc, 0xee, 0x1c, 0xe3, 0x28, 0xdc, 0xe4, 0x69, 0xc4, 0xbf, 0x73, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x0a, 0x9e, 0xa1, 0x09, 0x56, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RequestReplyServiceClient is the client API for RequestReplyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RequestReplyServiceClient interface {
	TraerAeropuerto(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error)
	TraerUsuario(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error)
	TraerVuelos(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error)
	TraerTickets(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error)
	CantidadTickets(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error)
	UsuarioMasCompras(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error)
}

type requestReplyServiceClient struct {
	cc *grpc.ClientConn
}

func NewRequestReplyServiceClient(cc *grpc.ClientConn) RequestReplyServiceClient {
	return &requestReplyServiceClient{cc}
}

func (c *requestReplyServiceClient) TraerAeropuerto(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/RequestReplyService/traerAeropuerto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requestReplyServiceClient) TraerUsuario(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/RequestReplyService/traerUsuario", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requestReplyServiceClient) TraerVuelos(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/RequestReplyService/traerVuelos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requestReplyServiceClient) TraerTickets(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/RequestReplyService/traerTickets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requestReplyServiceClient) CantidadTickets(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/RequestReplyService/cantidadTickets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requestReplyServiceClient) UsuarioMasCompras(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/RequestReplyService/usuarioMasCompras", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RequestReplyServiceServer is the server API for RequestReplyService service.
type RequestReplyServiceServer interface {
	TraerAeropuerto(context.Context, *Request) (*Reply, error)
	TraerUsuario(context.Context, *Request) (*Reply, error)
	TraerVuelos(context.Context, *Request) (*Reply, error)
	TraerTickets(context.Context, *Request) (*Reply, error)
	CantidadTickets(context.Context, *Request) (*Reply, error)
	UsuarioMasCompras(context.Context, *Request) (*Reply, error)
}

// UnimplementedRequestReplyServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRequestReplyServiceServer struct {
}

func (*UnimplementedRequestReplyServiceServer) TraerAeropuerto(ctx context.Context, req *Request) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TraerAeropuerto not implemented")
}
func (*UnimplementedRequestReplyServiceServer) TraerUsuario(ctx context.Context, req *Request) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TraerUsuario not implemented")
}
func (*UnimplementedRequestReplyServiceServer) TraerVuelos(ctx context.Context, req *Request) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TraerVuelos not implemented")
}
func (*UnimplementedRequestReplyServiceServer) TraerTickets(ctx context.Context, req *Request) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TraerTickets not implemented")
}
func (*UnimplementedRequestReplyServiceServer) CantidadTickets(ctx context.Context, req *Request) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CantidadTickets not implemented")
}
func (*UnimplementedRequestReplyServiceServer) UsuarioMasCompras(ctx context.Context, req *Request) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UsuarioMasCompras not implemented")
}

func RegisterRequestReplyServiceServer(s *grpc.Server, srv RequestReplyServiceServer) {
	s.RegisterService(&_RequestReplyService_serviceDesc, srv)
}

func _RequestReplyService_TraerAeropuerto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestReplyServiceServer).TraerAeropuerto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RequestReplyService/TraerAeropuerto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestReplyServiceServer).TraerAeropuerto(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _RequestReplyService_TraerUsuario_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestReplyServiceServer).TraerUsuario(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RequestReplyService/TraerUsuario",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestReplyServiceServer).TraerUsuario(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _RequestReplyService_TraerVuelos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestReplyServiceServer).TraerVuelos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RequestReplyService/TraerVuelos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestReplyServiceServer).TraerVuelos(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _RequestReplyService_TraerTickets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestReplyServiceServer).TraerTickets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RequestReplyService/TraerTickets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestReplyServiceServer).TraerTickets(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _RequestReplyService_CantidadTickets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestReplyServiceServer).CantidadTickets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RequestReplyService/CantidadTickets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestReplyServiceServer).CantidadTickets(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _RequestReplyService_UsuarioMasCompras_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestReplyServiceServer).UsuarioMasCompras(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RequestReplyService/UsuarioMasCompras",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestReplyServiceServer).UsuarioMasCompras(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _RequestReplyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "RequestReplyService",
	HandlerType: (*RequestReplyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "traerAeropuerto",
			Handler:    _RequestReplyService_TraerAeropuerto_Handler,
		},
		{
			MethodName: "traerUsuario",
			Handler:    _RequestReplyService_TraerUsuario_Handler,
		},
		{
			MethodName: "traerVuelos",
			Handler:    _RequestReplyService_TraerVuelos_Handler,
		},
		{
			MethodName: "traerTickets",
			Handler:    _RequestReplyService_TraerTickets_Handler,
		},
		{
			MethodName: "cantidadTickets",
			Handler:    _RequestReplyService_CantidadTickets_Handler,
		},
		{
			MethodName: "usuarioMasCompras",
			Handler:    _RequestReplyService_UsuarioMasCompras_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pingpong.proto",
}

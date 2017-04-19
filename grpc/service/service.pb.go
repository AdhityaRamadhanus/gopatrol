// Code generated by protoc-gen-go.
// source: service.proto
// DO NOT EDIT!

/*
Package service is a generated protocol buffer package.

It is generated from these files:
	service.proto

It has these top-level messages:
	GenericEndpointRequest
	AddHttpEndpointRequest
	AddTcpEndpointRequest
	AddDNSEndpointRequest
	DeleteEndpointRequest
	EndpointResponse
	ListEndpointRequest
	ListEndpointResponse
*/
package service

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

type GenericEndpointRequest struct {
	Name         string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Url          string `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
	Attempts     int32  `protobuf:"varint,3,opt,name=attempts" json:"attempts,omitempty"`
	Thresholdrtt int64  `protobuf:"varint,4,opt,name=thresholdrtt" json:"thresholdrtt,omitempty"`
}

func (m *GenericEndpointRequest) Reset()                    { *m = GenericEndpointRequest{} }
func (m *GenericEndpointRequest) String() string            { return proto.CompactTextString(m) }
func (*GenericEndpointRequest) ProtoMessage()               {}
func (*GenericEndpointRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GenericEndpointRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GenericEndpointRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *GenericEndpointRequest) GetAttempts() int32 {
	if m != nil {
		return m.Attempts
	}
	return 0
}

func (m *GenericEndpointRequest) GetThresholdrtt() int64 {
	if m != nil {
		return m.Thresholdrtt
	}
	return 0
}

type AddHttpEndpointRequest struct {
	Endpoint       *GenericEndpointRequest `protobuf:"bytes,1,opt,name=endpoint" json:"endpoint,omitempty"`
	Upstatus       int32                   `protobuf:"varint,2,opt,name=upstatus" json:"upstatus,omitempty"`
	MustContain    string                  `protobuf:"bytes,3,opt,name=must_contain,json=mustContain" json:"must_contain,omitempty"`
	MustNotContain string                  `protobuf:"bytes,4,opt,name=must_not_contain,json=mustNotContain" json:"must_not_contain,omitempty"`
	Headers        string                  `protobuf:"bytes,5,opt,name=headers" json:"headers,omitempty"`
}

func (m *AddHttpEndpointRequest) Reset()                    { *m = AddHttpEndpointRequest{} }
func (m *AddHttpEndpointRequest) String() string            { return proto.CompactTextString(m) }
func (*AddHttpEndpointRequest) ProtoMessage()               {}
func (*AddHttpEndpointRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AddHttpEndpointRequest) GetEndpoint() *GenericEndpointRequest {
	if m != nil {
		return m.Endpoint
	}
	return nil
}

func (m *AddHttpEndpointRequest) GetUpstatus() int32 {
	if m != nil {
		return m.Upstatus
	}
	return 0
}

func (m *AddHttpEndpointRequest) GetMustContain() string {
	if m != nil {
		return m.MustContain
	}
	return ""
}

func (m *AddHttpEndpointRequest) GetMustNotContain() string {
	if m != nil {
		return m.MustNotContain
	}
	return ""
}

func (m *AddHttpEndpointRequest) GetHeaders() string {
	if m != nil {
		return m.Headers
	}
	return ""
}

type AddTcpEndpointRequest struct {
	Endpoint      *GenericEndpointRequest `protobuf:"bytes,1,opt,name=endpoint" json:"endpoint,omitempty"`
	TlsEnabled    bool                    `protobuf:"varint,2,opt,name=tls_enabled,json=tlsEnabled" json:"tls_enabled,omitempty"`
	TlsSkipVerify bool                    `protobuf:"varint,3,opt,name=tls_skip_verify,json=tlsSkipVerify" json:"tls_skip_verify,omitempty"`
	TlsCaFile     string                  `protobuf:"bytes,4,opt,name=tls_ca_file,json=tlsCaFile" json:"tls_ca_file,omitempty"`
	Timeout       int64                   `protobuf:"varint,5,opt,name=timeout" json:"timeout,omitempty"`
}

func (m *AddTcpEndpointRequest) Reset()                    { *m = AddTcpEndpointRequest{} }
func (m *AddTcpEndpointRequest) String() string            { return proto.CompactTextString(m) }
func (*AddTcpEndpointRequest) ProtoMessage()               {}
func (*AddTcpEndpointRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AddTcpEndpointRequest) GetEndpoint() *GenericEndpointRequest {
	if m != nil {
		return m.Endpoint
	}
	return nil
}

func (m *AddTcpEndpointRequest) GetTlsEnabled() bool {
	if m != nil {
		return m.TlsEnabled
	}
	return false
}

func (m *AddTcpEndpointRequest) GetTlsSkipVerify() bool {
	if m != nil {
		return m.TlsSkipVerify
	}
	return false
}

func (m *AddTcpEndpointRequest) GetTlsCaFile() string {
	if m != nil {
		return m.TlsCaFile
	}
	return ""
}

func (m *AddTcpEndpointRequest) GetTimeout() int64 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

type AddDNSEndpointRequest struct {
	Endpoint *GenericEndpointRequest `protobuf:"bytes,1,opt,name=endpoint" json:"endpoint,omitempty"`
	Hostname string                  `protobuf:"bytes,2,opt,name=hostname" json:"hostname,omitempty"`
	Timeout  int64                   `protobuf:"varint,5,opt,name=timeout" json:"timeout,omitempty"`
}

func (m *AddDNSEndpointRequest) Reset()                    { *m = AddDNSEndpointRequest{} }
func (m *AddDNSEndpointRequest) String() string            { return proto.CompactTextString(m) }
func (*AddDNSEndpointRequest) ProtoMessage()               {}
func (*AddDNSEndpointRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AddDNSEndpointRequest) GetEndpoint() *GenericEndpointRequest {
	if m != nil {
		return m.Endpoint
	}
	return nil
}

func (m *AddDNSEndpointRequest) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *AddDNSEndpointRequest) GetTimeout() int64 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

type DeleteEndpointRequest struct {
	Url string `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
}

func (m *DeleteEndpointRequest) Reset()                    { *m = DeleteEndpointRequest{} }
func (m *DeleteEndpointRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteEndpointRequest) ProtoMessage()               {}
func (*DeleteEndpointRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *DeleteEndpointRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type EndpointResponse struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *EndpointResponse) Reset()                    { *m = EndpointResponse{} }
func (m *EndpointResponse) String() string            { return proto.CompactTextString(m) }
func (*EndpointResponse) ProtoMessage()               {}
func (*EndpointResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *EndpointResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ListEndpointRequest struct {
	Check bool `protobuf:"varint,1,opt,name=check" json:"check,omitempty"`
}

func (m *ListEndpointRequest) Reset()                    { *m = ListEndpointRequest{} }
func (m *ListEndpointRequest) String() string            { return proto.CompactTextString(m) }
func (*ListEndpointRequest) ProtoMessage()               {}
func (*ListEndpointRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ListEndpointRequest) GetCheck() bool {
	if m != nil {
		return m.Check
	}
	return false
}

type ListEndpointResponse struct {
	Endpoints []*ListEndpointResponse_Endpoint `protobuf:"bytes,1,rep,name=endpoints" json:"endpoints,omitempty"`
}

func (m *ListEndpointResponse) Reset()                    { *m = ListEndpointResponse{} }
func (m *ListEndpointResponse) String() string            { return proto.CompactTextString(m) }
func (*ListEndpointResponse) ProtoMessage()               {}
func (*ListEndpointResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ListEndpointResponse) GetEndpoints() []*ListEndpointResponse_Endpoint {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

type ListEndpointResponse_Endpoint struct {
	Name   string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Url    string `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
	Status string `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
}

func (m *ListEndpointResponse_Endpoint) Reset()         { *m = ListEndpointResponse_Endpoint{} }
func (m *ListEndpointResponse_Endpoint) String() string { return proto.CompactTextString(m) }
func (*ListEndpointResponse_Endpoint) ProtoMessage()    {}
func (*ListEndpointResponse_Endpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{7, 0}
}

func (m *ListEndpointResponse_Endpoint) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ListEndpointResponse_Endpoint) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *ListEndpointResponse_Endpoint) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*GenericEndpointRequest)(nil), "service.GenericEndpointRequest")
	proto.RegisterType((*AddHttpEndpointRequest)(nil), "service.AddHttpEndpointRequest")
	proto.RegisterType((*AddTcpEndpointRequest)(nil), "service.AddTcpEndpointRequest")
	proto.RegisterType((*AddDNSEndpointRequest)(nil), "service.AddDNSEndpointRequest")
	proto.RegisterType((*DeleteEndpointRequest)(nil), "service.DeleteEndpointRequest")
	proto.RegisterType((*EndpointResponse)(nil), "service.EndpointResponse")
	proto.RegisterType((*ListEndpointRequest)(nil), "service.ListEndpointRequest")
	proto.RegisterType((*ListEndpointResponse)(nil), "service.ListEndpointResponse")
	proto.RegisterType((*ListEndpointResponse_Endpoint)(nil), "service.ListEndpointResponse.Endpoint")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Checkup service

type CheckupClient interface {
	// Add Endpoint Function
	AddHTTPEndpoint(ctx context.Context, in *AddHttpEndpointRequest, opts ...grpc.CallOption) (*EndpointResponse, error)
	AddTCPEndpoint(ctx context.Context, in *AddTcpEndpointRequest, opts ...grpc.CallOption) (*EndpointResponse, error)
	AddDNSEndpoint(ctx context.Context, in *AddDNSEndpointRequest, opts ...grpc.CallOption) (*EndpointResponse, error)
	// Delete Endpoint Function
	DeleteEndpoint(ctx context.Context, in *DeleteEndpointRequest, opts ...grpc.CallOption) (*EndpointResponse, error)
	// Endpoint Inspection Function
	ListEndpoint(ctx context.Context, in *ListEndpointRequest, opts ...grpc.CallOption) (*ListEndpointResponse, error)
}

type checkupClient struct {
	cc *grpc.ClientConn
}

func NewCheckupClient(cc *grpc.ClientConn) CheckupClient {
	return &checkupClient{cc}
}

func (c *checkupClient) AddHTTPEndpoint(ctx context.Context, in *AddHttpEndpointRequest, opts ...grpc.CallOption) (*EndpointResponse, error) {
	out := new(EndpointResponse)
	err := grpc.Invoke(ctx, "/service.checkup/AddHTTPEndpoint", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkupClient) AddTCPEndpoint(ctx context.Context, in *AddTcpEndpointRequest, opts ...grpc.CallOption) (*EndpointResponse, error) {
	out := new(EndpointResponse)
	err := grpc.Invoke(ctx, "/service.checkup/AddTCPEndpoint", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkupClient) AddDNSEndpoint(ctx context.Context, in *AddDNSEndpointRequest, opts ...grpc.CallOption) (*EndpointResponse, error) {
	out := new(EndpointResponse)
	err := grpc.Invoke(ctx, "/service.checkup/AddDNSEndpoint", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkupClient) DeleteEndpoint(ctx context.Context, in *DeleteEndpointRequest, opts ...grpc.CallOption) (*EndpointResponse, error) {
	out := new(EndpointResponse)
	err := grpc.Invoke(ctx, "/service.checkup/DeleteEndpoint", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkupClient) ListEndpoint(ctx context.Context, in *ListEndpointRequest, opts ...grpc.CallOption) (*ListEndpointResponse, error) {
	out := new(ListEndpointResponse)
	err := grpc.Invoke(ctx, "/service.checkup/ListEndpoint", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Checkup service

type CheckupServer interface {
	// Add Endpoint Function
	AddHTTPEndpoint(context.Context, *AddHttpEndpointRequest) (*EndpointResponse, error)
	AddTCPEndpoint(context.Context, *AddTcpEndpointRequest) (*EndpointResponse, error)
	AddDNSEndpoint(context.Context, *AddDNSEndpointRequest) (*EndpointResponse, error)
	// Delete Endpoint Function
	DeleteEndpoint(context.Context, *DeleteEndpointRequest) (*EndpointResponse, error)
	// Endpoint Inspection Function
	ListEndpoint(context.Context, *ListEndpointRequest) (*ListEndpointResponse, error)
}

func RegisterCheckupServer(s *grpc.Server, srv CheckupServer) {
	s.RegisterService(&_Checkup_serviceDesc, srv)
}

func _Checkup_AddHTTPEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddHttpEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckupServer).AddHTTPEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.checkup/AddHTTPEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckupServer).AddHTTPEndpoint(ctx, req.(*AddHttpEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Checkup_AddTCPEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTcpEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckupServer).AddTCPEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.checkup/AddTCPEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckupServer).AddTCPEndpoint(ctx, req.(*AddTcpEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Checkup_AddDNSEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDNSEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckupServer).AddDNSEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.checkup/AddDNSEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckupServer).AddDNSEndpoint(ctx, req.(*AddDNSEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Checkup_DeleteEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckupServer).DeleteEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.checkup/DeleteEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckupServer).DeleteEndpoint(ctx, req.(*DeleteEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Checkup_ListEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckupServer).ListEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.checkup/ListEndpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckupServer).ListEndpoint(ctx, req.(*ListEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Checkup_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.checkup",
	HandlerType: (*CheckupServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddHTTPEndpoint",
			Handler:    _Checkup_AddHTTPEndpoint_Handler,
		},
		{
			MethodName: "AddTCPEndpoint",
			Handler:    _Checkup_AddTCPEndpoint_Handler,
		},
		{
			MethodName: "AddDNSEndpoint",
			Handler:    _Checkup_AddDNSEndpoint_Handler,
		},
		{
			MethodName: "DeleteEndpoint",
			Handler:    _Checkup_DeleteEndpoint_Handler,
		},
		{
			MethodName: "ListEndpoint",
			Handler:    _Checkup_ListEndpoint_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

func init() { proto.RegisterFile("service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 552 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x18, 0xac, 0x71, 0xd2, 0x38, 0x5f, 0xd2, 0x36, 0x5a, 0xda, 0xc8, 0x44, 0xd0, 0x06, 0x1f, 0xaa,
	0x20, 0x50, 0x0e, 0xe1, 0xc8, 0xa9, 0x4a, 0x0a, 0x3d, 0x40, 0x41, 0xdb, 0x88, 0x6b, 0xe4, 0xda,
	0x5f, 0xc9, 0x2a, 0x8e, 0x6d, 0xbc, 0x9f, 0x2b, 0x71, 0xe1, 0x01, 0x78, 0x00, 0x1e, 0x81, 0x57,
	0x82, 0xc7, 0x41, 0x5e, 0xdb, 0x9b, 0x9f, 0xa6, 0x3f, 0x87, 0xde, 0x3c, 0xa3, 0xd9, 0xd1, 0x4c,
	0x76, 0x36, 0xb0, 0x23, 0x31, 0xb9, 0x16, 0x1e, 0xf6, 0xe3, 0x24, 0xa2, 0x88, 0xd5, 0x0a, 0xe8,
	0xfc, 0x84, 0xf6, 0x07, 0x0c, 0x31, 0x11, 0xde, 0x69, 0xe8, 0xc7, 0x91, 0x08, 0x89, 0xe3, 0xf7,
	0x14, 0x25, 0x31, 0x06, 0x95, 0xd0, 0x9d, 0xa3, 0x6d, 0x74, 0x8d, 0x5e, 0x9d, 0xab, 0x6f, 0xd6,
	0x02, 0x33, 0x4d, 0x02, 0xfb, 0x89, 0xa2, 0xb2, 0x4f, 0xd6, 0x01, 0xcb, 0x25, 0xc2, 0x79, 0x4c,
	0xd2, 0x36, 0xbb, 0x46, 0xaf, 0xca, 0x35, 0x66, 0x0e, 0x34, 0x69, 0x9a, 0xa0, 0x9c, 0x46, 0x81,
	0x9f, 0x10, 0xd9, 0x95, 0xae, 0xd1, 0x33, 0xf9, 0x0a, 0xe7, 0xfc, 0x33, 0xa0, 0x7d, 0xe2, 0xfb,
	0x67, 0x44, 0xf1, 0x7a, 0x80, 0x77, 0x60, 0x61, 0x41, 0xa9, 0x10, 0x8d, 0xc1, 0x51, 0xbf, 0x6c,
	0xb1, 0x39, 0x33, 0xd7, 0x07, 0xb2, 0x5c, 0x69, 0x2c, 0xc9, 0xa5, 0x54, 0xaa, 0xb8, 0x55, 0xae,
	0x31, 0x7b, 0x09, 0xcd, 0x79, 0x2a, 0x69, 0xe2, 0x45, 0x21, 0xb9, 0x22, 0x54, 0xb9, 0xeb, 0xbc,
	0x91, 0x71, 0xc3, 0x9c, 0x62, 0x3d, 0x68, 0x29, 0x49, 0x18, 0x2d, 0x64, 0x15, 0x25, 0xdb, 0xcd,
	0xf8, 0xf3, 0x48, 0x2b, 0x6d, 0xa8, 0x4d, 0xd1, 0xf5, 0x31, 0x91, 0x76, 0x55, 0x09, 0x4a, 0xe8,
	0xfc, 0x35, 0xe0, 0xe0, 0xc4, 0xf7, 0xc7, 0xde, 0xe3, 0x36, 0x3b, 0x82, 0x06, 0x05, 0x72, 0x82,
	0xa1, 0x7b, 0x19, 0xa0, 0xaf, 0xca, 0x59, 0x1c, 0x28, 0x90, 0xa7, 0x39, 0xc3, 0x8e, 0x61, 0x2f,
	0x13, 0xc8, 0x99, 0x88, 0x27, 0xd7, 0x98, 0x88, 0xab, 0x1f, 0xaa, 0xa1, 0xc5, 0x77, 0x28, 0x90,
	0x17, 0x33, 0x11, 0x7f, 0x55, 0x24, 0x3b, 0xcc, 0x8d, 0x3c, 0x77, 0x72, 0x25, 0x02, 0x2c, 0xea,
	0xd5, 0x29, 0x90, 0x43, 0xf7, 0xbd, 0x08, 0x30, 0x6b, 0x46, 0x62, 0x8e, 0x51, 0x4a, 0xaa, 0x99,
	0xc9, 0x4b, 0xe8, 0xfc, 0xca, 0x9b, 0x8d, 0xce, 0x2f, 0x1e, 0xfb, 0xce, 0xa6, 0x91, 0x24, 0xb5,
	0xba, 0x7c, 0x62, 0x1a, 0xdf, 0x11, 0xe6, 0x15, 0x1c, 0x8c, 0x30, 0x40, 0xc2, 0xf5, 0x2c, 0xc5,
	0x58, 0x0d, 0x3d, 0x56, 0xe7, 0x0d, 0xb4, 0x16, 0x22, 0x19, 0x47, 0xa1, 0x54, 0xc6, 0x73, 0x94,
	0xd2, 0xfd, 0x56, 0x2e, 0xbd, 0x84, 0xce, 0x6b, 0x78, 0xfa, 0x51, 0x48, 0x5a, 0xb7, 0xdd, 0x87,
	0xaa, 0x37, 0x45, 0x6f, 0xa6, 0xe4, 0x16, 0xcf, 0x81, 0xf3, 0xc7, 0x80, 0xfd, 0x55, 0x75, 0xe1,
	0x3f, 0x82, 0x7a, 0x59, 0x50, 0xda, 0x46, 0xd7, 0xec, 0x35, 0x06, 0xc7, 0xfa, 0x27, 0xd9, 0x74,
	0xa2, 0xaf, 0x89, 0xc5, 0xc1, 0xce, 0x19, 0x58, 0x25, 0xfd, 0xc0, 0x87, 0xd9, 0x86, 0xed, 0x62,
	0xfe, 0xf9, 0xbc, 0x0b, 0x34, 0xf8, 0x6d, 0x42, 0x4d, 0x45, 0x4e, 0x63, 0xf6, 0x19, 0xf6, 0xb2,
	0xb7, 0x37, 0x1e, 0x7f, 0xd1, 0xe6, 0x8b, 0xeb, 0xda, 0xfc, 0x2a, 0x3b, 0xcf, 0xb4, 0x60, 0x3d,
	0xb8, 0xb3, 0xc5, 0x3e, 0xc1, 0x6e, 0xb6, 0xf8, 0xe1, 0xc2, 0xef, 0x70, 0xd9, 0xef, 0xe6, 0x53,
	0x78, 0x88, 0xdd, 0xd2, 0xcc, 0x56, 0xed, 0x6e, 0xee, 0xef, 0x5e, 0xbb, 0xd5, 0xa5, 0x2c, 0xd9,
	0x6d, 0x9c, 0xd0, 0x7d, 0x76, 0xcd, 0xe5, 0xfb, 0x63, 0xcf, 0x6f, 0xb9, 0xd6, 0xdc, 0xea, 0xc5,
	0x9d, 0x97, 0xee, 0x6c, 0x5d, 0x6e, 0xab, 0x7f, 0xe6, 0xb7, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff,
	0xb5, 0xf8, 0xbe, 0x7c, 0xaa, 0x05, 0x00, 0x00,
}

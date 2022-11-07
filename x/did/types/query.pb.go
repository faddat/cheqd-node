// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cheqd/did/v2/query.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type QueryGetDidDocRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *QueryGetDidDocRequest) Reset()         { *m = QueryGetDidDocRequest{} }
func (m *QueryGetDidDocRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetDidDocRequest) ProtoMessage()    {}
func (*QueryGetDidDocRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d818263856d0dc9, []int{0}
}
func (m *QueryGetDidDocRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetDidDocRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetDidDocRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetDidDocRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetDidDocRequest.Merge(m, src)
}
func (m *QueryGetDidDocRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetDidDocRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetDidDocRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetDidDocRequest proto.InternalMessageInfo

func (m *QueryGetDidDocRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type QueryGetDidDocResponse struct {
	Value *DidDocWithMetadata `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *QueryGetDidDocResponse) Reset()         { *m = QueryGetDidDocResponse{} }
func (m *QueryGetDidDocResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetDidDocResponse) ProtoMessage()    {}
func (*QueryGetDidDocResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d818263856d0dc9, []int{1}
}
func (m *QueryGetDidDocResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetDidDocResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetDidDocResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetDidDocResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetDidDocResponse.Merge(m, src)
}
func (m *QueryGetDidDocResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetDidDocResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetDidDocResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetDidDocResponse proto.InternalMessageInfo

func (m *QueryGetDidDocResponse) GetValue() *DidDocWithMetadata {
	if m != nil {
		return m.Value
	}
	return nil
}

type QueryGetDidDocVersionRequest struct {
	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (m *QueryGetDidDocVersionRequest) Reset()         { *m = QueryGetDidDocVersionRequest{} }
func (m *QueryGetDidDocVersionRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetDidDocVersionRequest) ProtoMessage()    {}
func (*QueryGetDidDocVersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d818263856d0dc9, []int{2}
}
func (m *QueryGetDidDocVersionRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetDidDocVersionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetDidDocVersionRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetDidDocVersionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetDidDocVersionRequest.Merge(m, src)
}
func (m *QueryGetDidDocVersionRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetDidDocVersionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetDidDocVersionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetDidDocVersionRequest proto.InternalMessageInfo

func (m *QueryGetDidDocVersionRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *QueryGetDidDocVersionRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type QueryGetDidDocVersionResponse struct {
	Value *DidDocWithMetadata `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *QueryGetDidDocVersionResponse) Reset()         { *m = QueryGetDidDocVersionResponse{} }
func (m *QueryGetDidDocVersionResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetDidDocVersionResponse) ProtoMessage()    {}
func (*QueryGetDidDocVersionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d818263856d0dc9, []int{3}
}
func (m *QueryGetDidDocVersionResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetDidDocVersionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetDidDocVersionResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetDidDocVersionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetDidDocVersionResponse.Merge(m, src)
}
func (m *QueryGetDidDocVersionResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetDidDocVersionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetDidDocVersionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetDidDocVersionResponse proto.InternalMessageInfo

func (m *QueryGetDidDocVersionResponse) GetValue() *DidDocWithMetadata {
	if m != nil {
		return m.Value
	}
	return nil
}

type QueryGetAllDidDocVersionsRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *QueryGetAllDidDocVersionsRequest) Reset()         { *m = QueryGetAllDidDocVersionsRequest{} }
func (m *QueryGetAllDidDocVersionsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetAllDidDocVersionsRequest) ProtoMessage()    {}
func (*QueryGetAllDidDocVersionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d818263856d0dc9, []int{4}
}
func (m *QueryGetAllDidDocVersionsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetAllDidDocVersionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetAllDidDocVersionsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetAllDidDocVersionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetAllDidDocVersionsRequest.Merge(m, src)
}
func (m *QueryGetAllDidDocVersionsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetAllDidDocVersionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetAllDidDocVersionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetAllDidDocVersionsRequest proto.InternalMessageInfo

func (m *QueryGetAllDidDocVersionsRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type QueryGetAllDidDocVersionsResponse struct {
	Versions []*Metadata `protobuf:"bytes,1,rep,name=versions,proto3" json:"versions,omitempty"`
}

func (m *QueryGetAllDidDocVersionsResponse) Reset()         { *m = QueryGetAllDidDocVersionsResponse{} }
func (m *QueryGetAllDidDocVersionsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetAllDidDocVersionsResponse) ProtoMessage()    {}
func (*QueryGetAllDidDocVersionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d818263856d0dc9, []int{5}
}
func (m *QueryGetAllDidDocVersionsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetAllDidDocVersionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetAllDidDocVersionsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetAllDidDocVersionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetAllDidDocVersionsResponse.Merge(m, src)
}
func (m *QueryGetAllDidDocVersionsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetAllDidDocVersionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetAllDidDocVersionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetAllDidDocVersionsResponse proto.InternalMessageInfo

func (m *QueryGetAllDidDocVersionsResponse) GetVersions() []*Metadata {
	if m != nil {
		return m.Versions
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryGetDidDocRequest)(nil), "cheqd.did.v2.QueryGetDidDocRequest")
	proto.RegisterType((*QueryGetDidDocResponse)(nil), "cheqd.did.v2.QueryGetDidDocResponse")
	proto.RegisterType((*QueryGetDidDocVersionRequest)(nil), "cheqd.did.v2.QueryGetDidDocVersionRequest")
	proto.RegisterType((*QueryGetDidDocVersionResponse)(nil), "cheqd.did.v2.QueryGetDidDocVersionResponse")
	proto.RegisterType((*QueryGetAllDidDocVersionsRequest)(nil), "cheqd.did.v2.QueryGetAllDidDocVersionsRequest")
	proto.RegisterType((*QueryGetAllDidDocVersionsResponse)(nil), "cheqd.did.v2.QueryGetAllDidDocVersionsResponse")
}

func init() { proto.RegisterFile("cheqd/did/v2/query.proto", fileDescriptor_8d818263856d0dc9) }

var fileDescriptor_8d818263856d0dc9 = []byte{
	// 426 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xc1, 0x8a, 0xd3, 0x40,
	0x18, 0xc7, 0x3b, 0x29, 0xad, 0x3a, 0x55, 0xc1, 0x01, 0x4b, 0x1a, 0x6b, 0x48, 0x53, 0xa5, 0xc5,
	0x62, 0x06, 0x22, 0x78, 0xaf, 0x14, 0xf4, 0x22, 0x68, 0x0f, 0x16, 0xbc, 0xa5, 0x99, 0xa1, 0x1d,
	0x88, 0x99, 0xb4, 0x33, 0x09, 0x96, 0xd2, 0x8b, 0x4f, 0x20, 0x88, 0xf8, 0x1c, 0xbe, 0x85, 0xc7,
	0x82, 0x97, 0x3d, 0x2e, 0xed, 0x3e, 0xc8, 0xb2, 0x93, 0xb4, 0x6c, 0x76, 0x9b, 0x52, 0xd8, 0x4b,
	0x20, 0xf9, 0xfe, 0xdf, 0x2f, 0xbf, 0x7c, 0x5f, 0x06, 0xea, 0xfe, 0x94, 0xce, 0x08, 0x26, 0x8c,
	0xe0, 0xc4, 0xc5, 0xb3, 0x98, 0xce, 0x17, 0x4e, 0x34, 0xe7, 0x92, 0xa3, 0x87, 0xaa, 0xe2, 0x10,
	0x46, 0x9c, 0xc4, 0x35, 0x9a, 0x13, 0xce, 0x27, 0x01, 0xc5, 0x5e, 0xc4, 0xb0, 0x17, 0x86, 0x5c,
	0x7a, 0x92, 0xf1, 0x50, 0xa4, 0x59, 0xa3, 0x91, 0xa3, 0x10, 0x46, 0x08, 0xf7, 0xd3, 0x92, 0xdd,
	0x81, 0x4f, 0x3f, 0x5f, 0x51, 0xdf, 0x53, 0x39, 0x60, 0x64, 0xc0, 0xfd, 0x21, 0x9d, 0xc5, 0x54,
	0x48, 0xf4, 0x18, 0x6a, 0x8c, 0xe8, 0xc0, 0x02, 0xdd, 0x07, 0x43, 0x8d, 0x11, 0xfb, 0x13, 0xac,
	0xdf, 0x0c, 0x8a, 0x88, 0x87, 0x82, 0xa2, 0xb7, 0xb0, 0x92, 0x78, 0x41, 0x4c, 0x55, 0xb8, 0xe6,
	0x5a, 0xce, 0x75, 0x33, 0x27, 0x0d, 0x8f, 0x98, 0x9c, 0x7e, 0xa4, 0xd2, 0x23, 0x9e, 0xf4, 0x86,
	0x69, 0xdc, 0xfe, 0x00, 0x9b, 0x79, 0xe2, 0x17, 0x3a, 0x17, 0x8c, 0x87, 0x05, 0x06, 0x48, 0x87,
	0xf7, 0x92, 0x34, 0xa1, 0x6b, 0xea, 0xe1, 0xee, 0xd6, 0x1e, 0xc1, 0xe7, 0x05, 0xa4, 0x3b, 0x2a,
	0xba, 0xd0, 0xda, 0x81, 0xfb, 0x41, 0x90, 0x63, 0x8b, 0xa2, 0x41, 0x8d, 0x60, 0xeb, 0x48, 0x4f,
	0x26, 0xe4, 0xc2, 0xfb, 0x99, 0xbc, 0xd0, 0x81, 0x55, 0xee, 0xd6, 0xdc, 0x7a, 0xde, 0x69, 0x6f,
	0xb2, 0xcf, 0xb9, 0x7f, 0xcb, 0xb0, 0xa2, 0xc8, 0x28, 0x81, 0xd5, 0x94, 0x8b, 0xda, 0xf9, 0xae,
	0x83, 0xab, 0x34, 0x5e, 0x1c, 0x0f, 0xa5, 0x4a, 0x76, 0xeb, 0xc7, 0xff, 0x8b, 0x5f, 0xda, 0x33,
	0xd4, 0xc0, 0x07, 0xfe, 0x16, 0xbc, 0x64, 0x64, 0x85, 0x7e, 0x03, 0xf8, 0x28, 0xf7, 0x41, 0xe8,
	0xd5, 0x31, 0x74, 0x7e, 0x9f, 0x46, 0xef, 0xa4, 0x6c, 0x66, 0xd3, 0x53, 0x36, 0x2f, 0x51, 0xbb,
	0xd0, 0x06, 0x2f, 0xb3, 0xc9, 0xac, 0xd0, 0x1f, 0x00, 0x9f, 0xdc, 0x9a, 0x35, 0x72, 0x0e, 0xbf,
	0xaf, 0x68, 0x91, 0x06, 0x3e, 0x39, 0x7f, 0xf2, 0xc4, 0xde, 0xf5, 0xff, 0x6d, 0x4c, 0xb0, 0xde,
	0x98, 0xe0, 0x7c, 0x63, 0x82, 0x9f, 0x5b, 0xb3, 0xb4, 0xde, 0x9a, 0xa5, 0xb3, 0xad, 0x59, 0xfa,
	0xda, 0x99, 0x30, 0x39, 0x8d, 0xc7, 0x8e, 0xcf, 0xbf, 0x65, 0xed, 0xea, 0xfa, 0x3a, 0xe4, 0x84,
	0xe2, 0xef, 0x8a, 0x25, 0x17, 0x11, 0x15, 0xe3, 0xaa, 0x3a, 0xa8, 0x6f, 0x2e, 0x03, 0x00, 0x00,
	0xff, 0xff, 0x3e, 0x17, 0x8f, 0x97, 0x0b, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	DidDoc(ctx context.Context, in *QueryGetDidDocRequest, opts ...grpc.CallOption) (*QueryGetDidDocResponse, error)
	DidDocVersion(ctx context.Context, in *QueryGetDidDocVersionRequest, opts ...grpc.CallOption) (*QueryGetDidDocVersionResponse, error)
	AllDidDocVersions(ctx context.Context, in *QueryGetAllDidDocVersionsRequest, opts ...grpc.CallOption) (*QueryGetAllDidDocVersionsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) DidDoc(ctx context.Context, in *QueryGetDidDocRequest, opts ...grpc.CallOption) (*QueryGetDidDocResponse, error) {
	out := new(QueryGetDidDocResponse)
	err := c.cc.Invoke(ctx, "/cheqd.did.v2.Query/DidDoc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DidDocVersion(ctx context.Context, in *QueryGetDidDocVersionRequest, opts ...grpc.CallOption) (*QueryGetDidDocVersionResponse, error) {
	out := new(QueryGetDidDocVersionResponse)
	err := c.cc.Invoke(ctx, "/cheqd.did.v2.Query/DidDocVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) AllDidDocVersions(ctx context.Context, in *QueryGetAllDidDocVersionsRequest, opts ...grpc.CallOption) (*QueryGetAllDidDocVersionsResponse, error) {
	out := new(QueryGetAllDidDocVersionsResponse)
	err := c.cc.Invoke(ctx, "/cheqd.did.v2.Query/AllDidDocVersions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	DidDoc(context.Context, *QueryGetDidDocRequest) (*QueryGetDidDocResponse, error)
	DidDocVersion(context.Context, *QueryGetDidDocVersionRequest) (*QueryGetDidDocVersionResponse, error)
	AllDidDocVersions(context.Context, *QueryGetAllDidDocVersionsRequest) (*QueryGetAllDidDocVersionsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) DidDoc(ctx context.Context, req *QueryGetDidDocRequest) (*QueryGetDidDocResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DidDoc not implemented")
}
func (*UnimplementedQueryServer) DidDocVersion(ctx context.Context, req *QueryGetDidDocVersionRequest) (*QueryGetDidDocVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DidDocVersion not implemented")
}
func (*UnimplementedQueryServer) AllDidDocVersions(ctx context.Context, req *QueryGetAllDidDocVersionsRequest) (*QueryGetAllDidDocVersionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllDidDocVersions not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_DidDoc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetDidDocRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DidDoc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cheqd.did.v2.Query/DidDoc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DidDoc(ctx, req.(*QueryGetDidDocRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DidDocVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetDidDocVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DidDocVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cheqd.did.v2.Query/DidDocVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DidDocVersion(ctx, req.(*QueryGetDidDocVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_AllDidDocVersions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetAllDidDocVersionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AllDidDocVersions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cheqd.did.v2.Query/AllDidDocVersions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AllDidDocVersions(ctx, req.(*QueryGetAllDidDocVersionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cheqd.did.v2.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DidDoc",
			Handler:    _Query_DidDoc_Handler,
		},
		{
			MethodName: "DidDocVersion",
			Handler:    _Query_DidDocVersion_Handler,
		},
		{
			MethodName: "AllDidDocVersions",
			Handler:    _Query_AllDidDocVersions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cheqd/did/v2/query.proto",
}

func (m *QueryGetDidDocRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetDidDocRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetDidDocRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetDidDocResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetDidDocResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetDidDocResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Value != nil {
		{
			size, err := m.Value.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetDidDocVersionRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetDidDocVersionRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetDidDocVersionRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Version) > 0 {
		i -= len(m.Version)
		copy(dAtA[i:], m.Version)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Version)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetDidDocVersionResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetDidDocVersionResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetDidDocVersionResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Value != nil {
		{
			size, err := m.Value.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetAllDidDocVersionsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetAllDidDocVersionsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetAllDidDocVersionsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetAllDidDocVersionsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetAllDidDocVersionsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetAllDidDocVersionsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Versions) > 0 {
		for iNdEx := len(m.Versions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Versions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryGetDidDocRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetDidDocResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Value != nil {
		l = m.Value.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetDidDocVersionRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetDidDocVersionResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Value != nil {
		l = m.Value.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetAllDidDocVersionsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetAllDidDocVersionsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Versions) > 0 {
		for _, e := range m.Versions {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryGetDidDocRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetDidDocRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetDidDocRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGetDidDocResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetDidDocResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetDidDocResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Value == nil {
				m.Value = &DidDocWithMetadata{}
			}
			if err := m.Value.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGetDidDocVersionRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetDidDocVersionRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetDidDocVersionRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGetDidDocVersionResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetDidDocVersionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetDidDocVersionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Value == nil {
				m.Value = &DidDocWithMetadata{}
			}
			if err := m.Value.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGetAllDidDocVersionsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetAllDidDocVersionsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetAllDidDocVersionsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGetAllDidDocVersionsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetAllDidDocVersionsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetAllDidDocVersionsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Versions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Versions = append(m.Versions, &Metadata{})
			if err := m.Versions[len(m.Versions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)

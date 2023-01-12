// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: cheqd/did/v2/query.proto

package didv2

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Fetch latest version of a DID Document for a given DID
	DidDoc(ctx context.Context, in *QueryDidDocRequest, opts ...grpc.CallOption) (*QueryDidDocResponse, error)
	// Fetch specific version of a DID Document for a given DID
	DidDocVersion(ctx context.Context, in *QueryDidDocVersionRequest, opts ...grpc.CallOption) (*QueryDidDocVersionResponse, error)
	// Fetch list of all versions of DID Documents for a given DID
	AllDidDocVersionsMetadata(ctx context.Context, in *QueryAllDidDocVersionsMetadataRequest, opts ...grpc.CallOption) (*QueryAllDidDocVersionsMetadataResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) DidDoc(ctx context.Context, in *QueryDidDocRequest, opts ...grpc.CallOption) (*QueryDidDocResponse, error) {
	out := new(QueryDidDocResponse)
	err := c.cc.Invoke(ctx, "/cheqd.did.v2.Query/DidDoc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DidDocVersion(ctx context.Context, in *QueryDidDocVersionRequest, opts ...grpc.CallOption) (*QueryDidDocVersionResponse, error) {
	out := new(QueryDidDocVersionResponse)
	err := c.cc.Invoke(ctx, "/cheqd.did.v2.Query/DidDocVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) AllDidDocVersionsMetadata(ctx context.Context, in *QueryAllDidDocVersionsMetadataRequest, opts ...grpc.CallOption) (*QueryAllDidDocVersionsMetadataResponse, error) {
	out := new(QueryAllDidDocVersionsMetadataResponse)
	err := c.cc.Invoke(ctx, "/cheqd.did.v2.Query/AllDidDocVersionsMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Fetch latest version of a DID Document for a given DID
	DidDoc(context.Context, *QueryDidDocRequest) (*QueryDidDocResponse, error)
	// Fetch specific version of a DID Document for a given DID
	DidDocVersion(context.Context, *QueryDidDocVersionRequest) (*QueryDidDocVersionResponse, error)
	// Fetch list of all versions of DID Documents for a given DID
	AllDidDocVersionsMetadata(context.Context, *QueryAllDidDocVersionsMetadataRequest) (*QueryAllDidDocVersionsMetadataResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) DidDoc(context.Context, *QueryDidDocRequest) (*QueryDidDocResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DidDoc not implemented")
}
func (UnimplementedQueryServer) DidDocVersion(context.Context, *QueryDidDocVersionRequest) (*QueryDidDocVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DidDocVersion not implemented")
}
func (UnimplementedQueryServer) AllDidDocVersionsMetadata(context.Context, *QueryAllDidDocVersionsMetadataRequest) (*QueryAllDidDocVersionsMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllDidDocVersionsMetadata not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_DidDoc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDidDocRequest)
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
		return srv.(QueryServer).DidDoc(ctx, req.(*QueryDidDocRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DidDocVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDidDocVersionRequest)
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
		return srv.(QueryServer).DidDocVersion(ctx, req.(*QueryDidDocVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_AllDidDocVersionsMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllDidDocVersionsMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AllDidDocVersionsMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cheqd.did.v2.Query/AllDidDocVersionsMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AllDidDocVersionsMetadata(ctx, req.(*QueryAllDidDocVersionsMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
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
			MethodName: "AllDidDocVersionsMetadata",
			Handler:    _Query_AllDidDocVersionsMetadata_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cheqd/did/v2/query.proto",
}

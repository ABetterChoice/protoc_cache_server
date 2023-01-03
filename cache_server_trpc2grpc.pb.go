// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.14.0
// source: cache_server.proto

package protoc_cache_server

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

// APIServerClientV2 is the client API for APIServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type APIServerClientV2 interface {
	GetTabConfig(ctx context.Context, in *GetTabConfigReq, opts ...grpc.CallOption) (*GetTabConfigResp, error)
	BatchGetExperimentBucket(ctx context.Context, in *BatchGetExperimentBucketReq, opts ...grpc.CallOption) (*BatchGetExperimentBucketResp, error)
	BatchGetGroupBucket(ctx context.Context, in *BatchGetGroupBucketReq, opts ...grpc.CallOption) (*BatchGetGroupBucketResp, error)
}

type aPIServerClientV2 struct {
	cc grpc.ClientConnInterface
}

func NewAPIServerClientV2(cc grpc.ClientConnInterface) APIServerClientV2 {
	return &aPIServerClientV2{cc}
}

func (c *aPIServerClientV2) GetTabConfig(ctx context.Context, in *GetTabConfigReq, opts ...grpc.CallOption) (*GetTabConfigResp, error) {
	out := new(GetTabConfigResp)
	err := c.cc.Invoke(ctx, "/opensource.tab.cache_server.APIServer/GetTabConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIServerClientV2) BatchGetExperimentBucket(ctx context.Context, in *BatchGetExperimentBucketReq, opts ...grpc.CallOption) (*BatchGetExperimentBucketResp, error) {
	out := new(BatchGetExperimentBucketResp)
	err := c.cc.Invoke(ctx, "/opensource.tab.cache_server.APIServer/BatchGetExperimentBucket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIServerClientV2) BatchGetGroupBucket(ctx context.Context, in *BatchGetGroupBucketReq, opts ...grpc.CallOption) (*BatchGetGroupBucketResp, error) {
	out := new(BatchGetGroupBucketResp)
	err := c.cc.Invoke(ctx, "/opensource.tab.cache_server.APIServer/BatchGetGroupBucket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// APIServerServerV2 is the server API for APIServer service.
// All implementations should embed UnimplementedAPIServerServerV2
// for forward compatibility
type APIServerServerV2 interface {
	GetTabConfig(context.Context, *GetTabConfigReq, *GetTabConfigResp) error
	BatchGetExperimentBucket(context.Context, *BatchGetExperimentBucketReq, *BatchGetExperimentBucketResp) error
	BatchGetGroupBucket(context.Context, *BatchGetGroupBucketReq, *BatchGetGroupBucketResp) error
}

// UnimplementedAPIServerServerV2 should be embedded to have forward compatible implementations.
type UnimplementedAPIServerServerV2 struct {
}

func (UnimplementedAPIServerServerV2) GetTabConfig(context.Context, *GetTabConfigReq, *GetTabConfigResp) error {
	return status.Errorf(codes.Unimplemented, "method GetTabConfig not implemented")
}
func (UnimplementedAPIServerServerV2) BatchGetExperimentBucket(context.Context, *BatchGetExperimentBucketReq, *BatchGetExperimentBucketResp) error {
	return status.Errorf(codes.Unimplemented, "method BatchGetExperimentBucket not implemented")
}
func (UnimplementedAPIServerServerV2) BatchGetGroupBucket(context.Context, *BatchGetGroupBucketReq, *BatchGetGroupBucketResp) error {
	return status.Errorf(codes.Unimplemented, "method BatchGetGroupBucket not implemented")
}

// UnsafeAPIServerServerV2 may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to APIServerServerV2 will
// result in compilation errors.
type UnsafeAPIServerServerV2 interface {
	mustEmbedUnimplementedAPIServerServerV2()
}

func RegisterAPIServerServerV2(s grpc.ServiceRegistrar, srv APIServerServerV2) {
	s.RegisterService(&APIServer_ServiceDescV2, srv)
}

func _APIServer_GetTabConfig_HandlerV2(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTabConfigReq)
	out := new(GetTabConfigResp)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		err := srv.(APIServerServerV2).GetTabConfig(ctx, in, out)
		return out, err
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/opensource.tab.cache_server.APIServer/GetTabConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		rsp := new(GetTabConfigResp)
		err := srv.(APIServerServerV2).GetTabConfig(ctx, req.(*GetTabConfigReq), rsp)
		return rsp, err
	}
	return interceptor(ctx, in, info, handler)
}

func _APIServer_BatchGetExperimentBucket_HandlerV2(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchGetExperimentBucketReq)
	out := new(BatchGetExperimentBucketResp)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		err := srv.(APIServerServerV2).BatchGetExperimentBucket(ctx, in, out)
		return out, err
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/opensource.tab.cache_server.APIServer/BatchGetExperimentBucket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		rsp := new(BatchGetExperimentBucketResp)
		err := srv.(APIServerServerV2).BatchGetExperimentBucket(ctx, req.(*BatchGetExperimentBucketReq), rsp)
		return rsp, err
	}
	return interceptor(ctx, in, info, handler)
}

func _APIServer_BatchGetGroupBucket_HandlerV2(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchGetGroupBucketReq)
	out := new(BatchGetGroupBucketResp)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		err := srv.(APIServerServerV2).BatchGetGroupBucket(ctx, in, out)
		return out, err
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/opensource.tab.cache_server.APIServer/BatchGetGroupBucket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		rsp := new(BatchGetGroupBucketResp)
		err := srv.(APIServerServerV2).BatchGetGroupBucket(ctx, req.(*BatchGetGroupBucketReq), rsp)
		return rsp, err
	}
	return interceptor(ctx, in, info, handler)
}

// APIServer_ServiceDescV2 is the grpc.ServiceDesc for APIServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var APIServer_ServiceDescV2 = grpc.ServiceDesc{
	ServiceName: "opensource.tab.cache_server.APIServer",
	HandlerType: (*APIServerServerV2)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTabConfig",
			Handler:    _APIServer_GetTabConfig_HandlerV2,
		},
		{
			MethodName: "BatchGetExperimentBucket",
			Handler:    _APIServer_BatchGetExperimentBucket_HandlerV2,
		},
		{
			MethodName: "BatchGetGroupBucket",
			Handler:    _APIServer_BatchGetGroupBucket_HandlerV2,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cache_server.proto",
}

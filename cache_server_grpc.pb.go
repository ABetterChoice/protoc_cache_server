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

// APIServerClient is the client API for APIServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type APIServerClient interface {
	GetTabConfig(ctx context.Context, in *GetTabConfigReq, opts ...grpc.CallOption) (*GetTabConfigResp, error)
	BatchGetExperimentBucket(ctx context.Context, in *BatchGetExperimentBucketReq, opts ...grpc.CallOption) (*BatchGetExperimentBucketResp, error)
	BatchGetGroupBucket(ctx context.Context, in *BatchGetGroupBucketReq, opts ...grpc.CallOption) (*BatchGetGroupBucketResp, error)
}

type aPIServerClient struct {
	cc grpc.ClientConnInterface
}

func NewAPIServerClient(cc grpc.ClientConnInterface) APIServerClient {
	return &aPIServerClient{cc}
}

func (c *aPIServerClient) GetTabConfig(ctx context.Context, in *GetTabConfigReq, opts ...grpc.CallOption) (*GetTabConfigResp, error) {
	out := new(GetTabConfigResp)
	err := c.cc.Invoke(ctx, "/opensource.tab.cache_server.APIServer/GetTabConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIServerClient) BatchGetExperimentBucket(ctx context.Context, in *BatchGetExperimentBucketReq, opts ...grpc.CallOption) (*BatchGetExperimentBucketResp, error) {
	out := new(BatchGetExperimentBucketResp)
	err := c.cc.Invoke(ctx, "/opensource.tab.cache_server.APIServer/BatchGetExperimentBucket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIServerClient) BatchGetGroupBucket(ctx context.Context, in *BatchGetGroupBucketReq, opts ...grpc.CallOption) (*BatchGetGroupBucketResp, error) {
	out := new(BatchGetGroupBucketResp)
	err := c.cc.Invoke(ctx, "/opensource.tab.cache_server.APIServer/BatchGetGroupBucket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// APIServerServer is the server API for APIServer service.
// All implementations should embed UnimplementedAPIServerServer
// for forward compatibility
type APIServerServer interface {
	GetTabConfig(context.Context, *GetTabConfigReq) (*GetTabConfigResp, error)
	BatchGetExperimentBucket(context.Context, *BatchGetExperimentBucketReq) (*BatchGetExperimentBucketResp, error)
	BatchGetGroupBucket(context.Context, *BatchGetGroupBucketReq) (*BatchGetGroupBucketResp, error)
}

// UnimplementedAPIServerServer should be embedded to have forward compatible implementations.
type UnimplementedAPIServerServer struct {
}

func (UnimplementedAPIServerServer) GetTabConfig(context.Context, *GetTabConfigReq) (*GetTabConfigResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTabConfig not implemented")
}
func (UnimplementedAPIServerServer) BatchGetExperimentBucket(context.Context, *BatchGetExperimentBucketReq) (*BatchGetExperimentBucketResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchGetExperimentBucket not implemented")
}
func (UnimplementedAPIServerServer) BatchGetGroupBucket(context.Context, *BatchGetGroupBucketReq) (*BatchGetGroupBucketResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchGetGroupBucket not implemented")
}

// UnsafeAPIServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to APIServerServer will
// result in compilation errors.
type UnsafeAPIServerServer interface {
	mustEmbedUnimplementedAPIServerServer()
}

func RegisterAPIServerServer(s grpc.ServiceRegistrar, srv APIServerServer) {
	s.RegisterService(&APIServer_ServiceDesc, srv)
}

func _APIServer_GetTabConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTabConfigReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServerServer).GetTabConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/opensource.tab.cache_server.APIServer/GetTabConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServerServer).GetTabConfig(ctx, req.(*GetTabConfigReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _APIServer_BatchGetExperimentBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchGetExperimentBucketReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServerServer).BatchGetExperimentBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/opensource.tab.cache_server.APIServer/BatchGetExperimentBucket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServerServer).BatchGetExperimentBucket(ctx, req.(*BatchGetExperimentBucketReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _APIServer_BatchGetGroupBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchGetGroupBucketReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServerServer).BatchGetGroupBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/opensource.tab.cache_server.APIServer/BatchGetGroupBucket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServerServer).BatchGetGroupBucket(ctx, req.(*BatchGetGroupBucketReq))
	}
	return interceptor(ctx, in, info, handler)
}

// APIServer_ServiceDesc is the grpc.ServiceDesc for APIServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var APIServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "opensource.tab.cache_server.APIServer",
	HandlerType: (*APIServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTabConfig",
			Handler:    _APIServer_GetTabConfig_Handler,
		},
		{
			MethodName: "BatchGetExperimentBucket",
			Handler:    _APIServer_BatchGetExperimentBucket_Handler,
		},
		{
			MethodName: "BatchGetGroupBucket",
			Handler:    _APIServer_BatchGetGroupBucket_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cache_server.proto",
}

// EventServerClient is the client API for EventServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventServerClient interface {
	LogExposureGroup(ctx context.Context, in *ExposureGroup, opts ...grpc.CallOption) (*CommonResp, error)
	LogEventGroup(ctx context.Context, in *EventGroup, opts ...grpc.CallOption) (*CommonResp, error)
}

type eventServerClient struct {
	cc grpc.ClientConnInterface
}

func NewEventServerClient(cc grpc.ClientConnInterface) EventServerClient {
	return &eventServerClient{cc}
}

func (c *eventServerClient) LogExposureGroup(ctx context.Context, in *ExposureGroup, opts ...grpc.CallOption) (*CommonResp, error) {
	out := new(CommonResp)
	err := c.cc.Invoke(ctx, "/opensource.tab.cache_server.EventServer/LogExposureGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventServerClient) LogEventGroup(ctx context.Context, in *EventGroup, opts ...grpc.CallOption) (*CommonResp, error) {
	out := new(CommonResp)
	err := c.cc.Invoke(ctx, "/opensource.tab.cache_server.EventServer/LogEventGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventServerServer is the server API for EventServer service.
// All implementations should embed UnimplementedEventServerServer
// for forward compatibility
type EventServerServer interface {
	LogExposureGroup(context.Context, *ExposureGroup) (*CommonResp, error)
	LogEventGroup(context.Context, *EventGroup) (*CommonResp, error)
}

// UnimplementedEventServerServer should be embedded to have forward compatible implementations.
type UnimplementedEventServerServer struct {
}

func (UnimplementedEventServerServer) LogExposureGroup(context.Context, *ExposureGroup) (*CommonResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogExposureGroup not implemented")
}
func (UnimplementedEventServerServer) LogEventGroup(context.Context, *EventGroup) (*CommonResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogEventGroup not implemented")
}

// UnsafeEventServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventServerServer will
// result in compilation errors.
type UnsafeEventServerServer interface {
	mustEmbedUnimplementedEventServerServer()
}

func RegisterEventServerServer(s grpc.ServiceRegistrar, srv EventServerServer) {
	s.RegisterService(&EventServer_ServiceDesc, srv)
}

func _EventServer_LogExposureGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExposureGroup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServerServer).LogExposureGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/opensource.tab.cache_server.EventServer/LogExposureGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServerServer).LogExposureGroup(ctx, req.(*ExposureGroup))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventServer_LogEventGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventGroup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServerServer).LogEventGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/opensource.tab.cache_server.EventServer/LogEventGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServerServer).LogEventGroup(ctx, req.(*EventGroup))
	}
	return interceptor(ctx, in, info, handler)
}

// EventServer_ServiceDesc is the grpc.ServiceDesc for EventServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EventServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "opensource.tab.cache_server.EventServer",
	HandlerType: (*EventServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LogExposureGroup",
			Handler:    _EventServer_LogExposureGroup_Handler,
		},
		{
			MethodName: "LogEventGroup",
			Handler:    _EventServer_LogEventGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cache_server.proto",
}

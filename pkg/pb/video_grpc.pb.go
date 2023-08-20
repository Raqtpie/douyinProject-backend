// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.0--rc2
// source: video.proto

package pb

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

const (
	VideoCenter_Feed_FullMethodName          = "/douyin.core.video_center/Feed"
	VideoCenter_PublishList_FullMethodName   = "/douyin.core.video_center/PublishList"
	VideoCenter_PublishAction_FullMethodName = "/douyin.core.video_center/PublishAction"
)

// VideoCenterClient is the client API for VideoCenter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VideoCenterClient interface {
	// 视频流
	Feed(ctx context.Context, in *DouyinFeedRequest, opts ...grpc.CallOption) (*DouyinFeedResponse, error)
	// 发布列表
	PublishList(ctx context.Context, in *DouyinPublishListRequest, opts ...grpc.CallOption) (*DouyinPublishListResponse, error)
	// 视频投稿
	PublishAction(ctx context.Context, in *DouyinPublishActionRequest, opts ...grpc.CallOption) (*DouyinPublishActionResponse, error)
}

type videoCenterClient struct {
	cc grpc.ClientConnInterface
}

func NewVideoCenterClient(cc grpc.ClientConnInterface) VideoCenterClient {
	return &videoCenterClient{cc}
}

func (c *videoCenterClient) Feed(ctx context.Context, in *DouyinFeedRequest, opts ...grpc.CallOption) (*DouyinFeedResponse, error) {
	out := new(DouyinFeedResponse)
	err := c.cc.Invoke(ctx, VideoCenter_Feed_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoCenterClient) PublishList(ctx context.Context, in *DouyinPublishListRequest, opts ...grpc.CallOption) (*DouyinPublishListResponse, error) {
	out := new(DouyinPublishListResponse)
	err := c.cc.Invoke(ctx, VideoCenter_PublishList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoCenterClient) PublishAction(ctx context.Context, in *DouyinPublishActionRequest, opts ...grpc.CallOption) (*DouyinPublishActionResponse, error) {
	out := new(DouyinPublishActionResponse)
	err := c.cc.Invoke(ctx, VideoCenter_PublishAction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VideoCenterServer is the server API for VideoCenter service.
// All implementations must embed UnimplementedVideoCenterServer
// for forward compatibility
type VideoCenterServer interface {
	// 视频流
	Feed(context.Context, *DouyinFeedRequest) (*DouyinFeedResponse, error)
	// 发布列表
	PublishList(context.Context, *DouyinPublishListRequest) (*DouyinPublishListResponse, error)
	// 视频投稿
	PublishAction(context.Context, *DouyinPublishActionRequest) (*DouyinPublishActionResponse, error)
	mustEmbedUnimplementedVideoCenterServer()
}

// UnimplementedVideoCenterServer must be embedded to have forward compatible implementations.
type UnimplementedVideoCenterServer struct {
}

func (UnimplementedVideoCenterServer) Feed(context.Context, *DouyinFeedRequest) (*DouyinFeedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Feed not implemented")
}
func (UnimplementedVideoCenterServer) PublishList(context.Context, *DouyinPublishListRequest) (*DouyinPublishListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishList not implemented")
}
func (UnimplementedVideoCenterServer) PublishAction(context.Context, *DouyinPublishActionRequest) (*DouyinPublishActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishAction not implemented")
}
func (UnimplementedVideoCenterServer) mustEmbedUnimplementedVideoCenterServer() {}

// UnsafeVideoCenterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VideoCenterServer will
// result in compilation errors.
type UnsafeVideoCenterServer interface {
	mustEmbedUnimplementedVideoCenterServer()
}

func RegisterVideoCenterServer(s grpc.ServiceRegistrar, srv VideoCenterServer) {
	s.RegisterService(&VideoCenter_ServiceDesc, srv)
}

func _VideoCenter_Feed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinFeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoCenterServer).Feed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VideoCenter_Feed_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoCenterServer).Feed(ctx, req.(*DouyinFeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoCenter_PublishList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinPublishListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoCenterServer).PublishList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VideoCenter_PublishList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoCenterServer).PublishList(ctx, req.(*DouyinPublishListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoCenter_PublishAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinPublishActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoCenterServer).PublishAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VideoCenter_PublishAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoCenterServer).PublishAction(ctx, req.(*DouyinPublishActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VideoCenter_ServiceDesc is the grpc.ServiceDesc for VideoCenter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VideoCenter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "douyin.core.video_center",
	HandlerType: (*VideoCenterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Feed",
			Handler:    _VideoCenter_Feed_Handler,
		},
		{
			MethodName: "PublishList",
			Handler:    _VideoCenter_PublishList_Handler,
		},
		{
			MethodName: "PublishAction",
			Handler:    _VideoCenter_PublishAction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "video.proto",
}
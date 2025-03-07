// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: bitcolibri/birdFeed/v1/query.proto

package birdFeedv1

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
	Query_GetTweet_FullMethodName         = "/bitcolibri.birdFeed.v1.Query/GetTweet"
	Query_GetAuthorTweets_FullMethodName  = "/bitcolibri.birdFeed.v1.Query/GetAuthorTweets"
	Query_GetTweetLikes_FullMethodName    = "/bitcolibri.birdFeed.v1.Query/GetTweetLikes"
	Query_GetUser_FullMethodName          = "/bitcolibri.birdFeed.v1.Query/GetUser"
	Query_GetUserFollowers_FullMethodName = "/bitcolibri.birdFeed.v1.Query/GetUserFollowers"
	Query_GetUserFollows_FullMethodName   = "/bitcolibri.birdFeed.v1.Query/GetUserFollows"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	GetTweet(ctx context.Context, in *QueryGetTweetRequest, opts ...grpc.CallOption) (*QueryGetTweetResponse, error)
	GetAuthorTweets(ctx context.Context, in *QueryGetAuthorTweetsRequest, opts ...grpc.CallOption) (*QueryGetAuthorTweetsResponse, error)
	GetTweetLikes(ctx context.Context, in *QueryGetTweetLikesRequest, opts ...grpc.CallOption) (*QueryGetTweetLikesResponse, error)
	GetUser(ctx context.Context, in *QueryGetUserRequest, opts ...grpc.CallOption) (*QueryGetUserResponse, error)
	GetUserFollowers(ctx context.Context, in *QueryGetUserFollowersRequest, opts ...grpc.CallOption) (*QueryGetUserFollowersResponse, error)
	GetUserFollows(ctx context.Context, in *QueryGetUserFollowsRequest, opts ...grpc.CallOption) (*QueryGetUserFollowsResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) GetTweet(ctx context.Context, in *QueryGetTweetRequest, opts ...grpc.CallOption) (*QueryGetTweetResponse, error) {
	out := new(QueryGetTweetResponse)
	err := c.cc.Invoke(ctx, Query_GetTweet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetAuthorTweets(ctx context.Context, in *QueryGetAuthorTweetsRequest, opts ...grpc.CallOption) (*QueryGetAuthorTweetsResponse, error) {
	out := new(QueryGetAuthorTweetsResponse)
	err := c.cc.Invoke(ctx, Query_GetAuthorTweets_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetTweetLikes(ctx context.Context, in *QueryGetTweetLikesRequest, opts ...grpc.CallOption) (*QueryGetTweetLikesResponse, error) {
	out := new(QueryGetTweetLikesResponse)
	err := c.cc.Invoke(ctx, Query_GetTweetLikes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetUser(ctx context.Context, in *QueryGetUserRequest, opts ...grpc.CallOption) (*QueryGetUserResponse, error) {
	out := new(QueryGetUserResponse)
	err := c.cc.Invoke(ctx, Query_GetUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetUserFollowers(ctx context.Context, in *QueryGetUserFollowersRequest, opts ...grpc.CallOption) (*QueryGetUserFollowersResponse, error) {
	out := new(QueryGetUserFollowersResponse)
	err := c.cc.Invoke(ctx, Query_GetUserFollowers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetUserFollows(ctx context.Context, in *QueryGetUserFollowsRequest, opts ...grpc.CallOption) (*QueryGetUserFollowsResponse, error) {
	out := new(QueryGetUserFollowsResponse)
	err := c.cc.Invoke(ctx, Query_GetUserFollows_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	GetTweet(context.Context, *QueryGetTweetRequest) (*QueryGetTweetResponse, error)
	GetAuthorTweets(context.Context, *QueryGetAuthorTweetsRequest) (*QueryGetAuthorTweetsResponse, error)
	GetTweetLikes(context.Context, *QueryGetTweetLikesRequest) (*QueryGetTweetLikesResponse, error)
	GetUser(context.Context, *QueryGetUserRequest) (*QueryGetUserResponse, error)
	GetUserFollowers(context.Context, *QueryGetUserFollowersRequest) (*QueryGetUserFollowersResponse, error)
	GetUserFollows(context.Context, *QueryGetUserFollowsRequest) (*QueryGetUserFollowsResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) GetTweet(context.Context, *QueryGetTweetRequest) (*QueryGetTweetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTweet not implemented")
}
func (UnimplementedQueryServer) GetAuthorTweets(context.Context, *QueryGetAuthorTweetsRequest) (*QueryGetAuthorTweetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthorTweets not implemented")
}
func (UnimplementedQueryServer) GetTweetLikes(context.Context, *QueryGetTweetLikesRequest) (*QueryGetTweetLikesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTweetLikes not implemented")
}
func (UnimplementedQueryServer) GetUser(context.Context, *QueryGetUserRequest) (*QueryGetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedQueryServer) GetUserFollowers(context.Context, *QueryGetUserFollowersRequest) (*QueryGetUserFollowersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserFollowers not implemented")
}
func (UnimplementedQueryServer) GetUserFollows(context.Context, *QueryGetUserFollowsRequest) (*QueryGetUserFollowsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserFollows not implemented")
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

func _Query_GetTweet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetTweetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetTweet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GetTweet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetTweet(ctx, req.(*QueryGetTweetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetAuthorTweets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetAuthorTweetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetAuthorTweets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GetAuthorTweets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetAuthorTweets(ctx, req.(*QueryGetAuthorTweetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetTweetLikes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetTweetLikesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetTweetLikes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GetTweetLikes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetTweetLikes(ctx, req.(*QueryGetTweetLikesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetUser(ctx, req.(*QueryGetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetUserFollowers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetUserFollowersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetUserFollowers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GetUserFollowers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetUserFollowers(ctx, req.(*QueryGetUserFollowersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetUserFollows_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetUserFollowsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetUserFollows(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GetUserFollows_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetUserFollows(ctx, req.(*QueryGetUserFollowsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bitcolibri.birdFeed.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTweet",
			Handler:    _Query_GetTweet_Handler,
		},
		{
			MethodName: "GetAuthorTweets",
			Handler:    _Query_GetAuthorTweets_Handler,
		},
		{
			MethodName: "GetTweetLikes",
			Handler:    _Query_GetTweetLikes_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _Query_GetUser_Handler,
		},
		{
			MethodName: "GetUserFollowers",
			Handler:    _Query_GetUserFollowers_Handler,
		},
		{
			MethodName: "GetUserFollows",
			Handler:    _Query_GetUserFollows_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bitcolibri/birdFeed/v1/query.proto",
}

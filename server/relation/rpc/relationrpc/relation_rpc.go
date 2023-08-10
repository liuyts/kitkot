// Code generated by goctl. DO NOT EDIT.
// Source: relation.proto

package relationrpc

import (
	"context"

	"kitkot/server/relation/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	FollowActionRequest          = pb.FollowActionRequest
	FollowActionResponse         = pb.FollowActionResponse
	FriendUser                   = pb.FriendUser
	GetFollowListRequest         = pb.GetFollowListRequest
	GetFollowListResponse        = pb.GetFollowListResponse
	GetFollowerListRequest       = pb.GetFollowerListRequest
	GetFollowerListResponse      = pb.GetFollowerListResponse
	GetFriendListRequest         = pb.GetFriendListRequest
	GetFriendListResponse        = pb.GetFriendListResponse
	GetUserFollowCountRequest    = pb.GetUserFollowCountRequest
	GetUserFollowCountResponse   = pb.GetUserFollowCountResponse
	GetUserFollowerCountRequest  = pb.GetUserFollowerCountRequest
	GetUserFollowerCountResponse = pb.GetUserFollowerCountResponse
	IsFollowRequest              = pb.IsFollowRequest
	IsFollowResponse             = pb.IsFollowResponse
	User                         = pb.User

	RelationRpc interface {
		FollowAction(ctx context.Context, in *FollowActionRequest, opts ...grpc.CallOption) (*FollowActionResponse, error)
		GetFollowList(ctx context.Context, in *GetFollowListRequest, opts ...grpc.CallOption) (*GetFollowListResponse, error)
		GetFollowerList(ctx context.Context, in *GetFollowerListRequest, opts ...grpc.CallOption) (*GetFollowerListResponse, error)
		GetUserFollowCount(ctx context.Context, in *GetUserFollowCountRequest, opts ...grpc.CallOption) (*GetUserFollowCountResponse, error)
		GetUserFollowerCount(ctx context.Context, in *GetUserFollowerCountRequest, opts ...grpc.CallOption) (*GetUserFollowerCountResponse, error)
		GetFriendList(ctx context.Context, in *GetFriendListRequest, opts ...grpc.CallOption) (*GetFriendListResponse, error)
		IsFollow(ctx context.Context, in *IsFollowRequest, opts ...grpc.CallOption) (*IsFollowResponse, error)
	}

	defaultRelationRpc struct {
		cli zrpc.Client
	}
)

func NewRelationRpc(cli zrpc.Client) RelationRpc {
	return &defaultRelationRpc{
		cli: cli,
	}
}

func (m *defaultRelationRpc) FollowAction(ctx context.Context, in *FollowActionRequest, opts ...grpc.CallOption) (*FollowActionResponse, error) {
	client := pb.NewRelationRpcClient(m.cli.Conn())
	return client.FollowAction(ctx, in, opts...)
}

func (m *defaultRelationRpc) GetFollowList(ctx context.Context, in *GetFollowListRequest, opts ...grpc.CallOption) (*GetFollowListResponse, error) {
	client := pb.NewRelationRpcClient(m.cli.Conn())
	return client.GetFollowList(ctx, in, opts...)
}

func (m *defaultRelationRpc) GetFollowerList(ctx context.Context, in *GetFollowerListRequest, opts ...grpc.CallOption) (*GetFollowerListResponse, error) {
	client := pb.NewRelationRpcClient(m.cli.Conn())
	return client.GetFollowerList(ctx, in, opts...)
}

func (m *defaultRelationRpc) GetUserFollowCount(ctx context.Context, in *GetUserFollowCountRequest, opts ...grpc.CallOption) (*GetUserFollowCountResponse, error) {
	client := pb.NewRelationRpcClient(m.cli.Conn())
	return client.GetUserFollowCount(ctx, in, opts...)
}

func (m *defaultRelationRpc) GetUserFollowerCount(ctx context.Context, in *GetUserFollowerCountRequest, opts ...grpc.CallOption) (*GetUserFollowerCountResponse, error) {
	client := pb.NewRelationRpcClient(m.cli.Conn())
	return client.GetUserFollowerCount(ctx, in, opts...)
}

func (m *defaultRelationRpc) GetFriendList(ctx context.Context, in *GetFriendListRequest, opts ...grpc.CallOption) (*GetFriendListResponse, error) {
	client := pb.NewRelationRpcClient(m.cli.Conn())
	return client.GetFriendList(ctx, in, opts...)
}

func (m *defaultRelationRpc) IsFollow(ctx context.Context, in *IsFollowRequest, opts ...grpc.CallOption) (*IsFollowResponse, error) {
	client := pb.NewRelationRpcClient(m.cli.Conn())
	return client.IsFollow(ctx, in, opts...)
}

package mock

import (
	"context"
	"google.golang.org/grpc"
	"kitkot/server/relation/rpc/relationrpc"
)

type RelationRpc struct {
}

func (r RelationRpc) AddFollow(ctx context.Context, in *relationrpc.AddFollowRequest, opts ...grpc.CallOption) (*relationrpc.AddFollowResponse, error) {
	return &relationrpc.AddFollowResponse{}, nil
}

func (r RelationRpc) DelFollow(ctx context.Context, in *relationrpc.DelFollowRequest, opts ...grpc.CallOption) (*relationrpc.DelFollowResponse, error) {
	return &relationrpc.DelFollowResponse{}, nil
}

func (r RelationRpc) GetFollowList(ctx context.Context, in *relationrpc.GetFollowListRequest, opts ...grpc.CallOption) (*relationrpc.GetFollowListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r RelationRpc) GetFollowerList(ctx context.Context, in *relationrpc.GetFollowerListRequest, opts ...grpc.CallOption) (*relationrpc.GetFollowerListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r RelationRpc) GetUserFollowCount(ctx context.Context, in *relationrpc.GetUserFollowCountRequest, opts ...grpc.CallOption) (*relationrpc.GetUserFollowCountResponse, error) {
	return &relationrpc.GetUserFollowCountResponse{Count: 100}, nil
}

func (r RelationRpc) GetUserFollowerCount(ctx context.Context, in *relationrpc.GetUserFollowerCountRequest, opts ...grpc.CallOption) (*relationrpc.GetUserFollowerCountResponse, error) {
	return &relationrpc.GetUserFollowerCountResponse{Count: 200}, nil
}

func (r RelationRpc) GetFriendList(ctx context.Context, in *relationrpc.GetFriendListRequest, opts ...grpc.CallOption) (*relationrpc.GetFriendListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r RelationRpc) IsFollow(ctx context.Context, in *relationrpc.IsFollowRequest, opts ...grpc.CallOption) (*relationrpc.IsFollowResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewRelationRpc() *RelationRpc {
	return &RelationRpc{}
}

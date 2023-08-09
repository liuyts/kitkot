package mock

import (
	"context"
	"google.golang.org/grpc"
	"kitkot/server/favorite/rpc/favoriterpc"
)

type FavoriteRpc struct {
}

func (f FavoriteRpc) AddFavorite(ctx context.Context, in *favoriterpc.AddFavoriteRequest, opts ...grpc.CallOption) (*favoriterpc.AddFavoriteResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (f FavoriteRpc) DelFavorite(ctx context.Context, in *favoriterpc.DelFavoriteRequest, opts ...grpc.CallOption) (*favoriterpc.DelFavoriteResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (f FavoriteRpc) GetVideoFavoriteCount(ctx context.Context, in *favoriterpc.GetVideoFavoriteCountRequest, opts ...grpc.CallOption) (*favoriterpc.GetVideoFavoriteCountResponse, error) {
	return &favoriterpc.GetVideoFavoriteCountResponse{Count: 10}, nil
}

func (f FavoriteRpc) GetUserFavoriteCount(ctx context.Context, in *favoriterpc.GetUserFavoriteCountRequest, opts ...grpc.CallOption) (*favoriterpc.GetUserFavoriteCountResponse, error) {
	return &favoriterpc.GetUserFavoriteCountResponse{Count: 50}, nil
}

func (f FavoriteRpc) GetUserFavoritedCount(ctx context.Context, in *favoriterpc.GetUserFavoritedCountRequest, opts ...grpc.CallOption) (*favoriterpc.GetUserFavoritedCountResponse, error) {
	return &favoriterpc.GetUserFavoritedCountResponse{Count: 60}, nil
}

func (f FavoriteRpc) IsFavorite(ctx context.Context, in *favoriterpc.IsFavoriteRequest, opts ...grpc.CallOption) (*favoriterpc.IsFavoriteResponse, error) {
	return &favoriterpc.IsFavoriteResponse{IsFavorite: true}, nil
}

func NewFavoriteRpc() *FavoriteRpc {
	return &FavoriteRpc{}
}

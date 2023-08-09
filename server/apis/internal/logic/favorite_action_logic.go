package logic

import (
	"context"
	"kitkot/common/consts"
	"kitkot/server/favorite/rpc/favoriterpc"

	"kitkot/server/apis/internal/svc"
	"kitkot/server/apis/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteActionLogic {
	return &FavoriteActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteActionLogic) FavoriteAction(req *types.FavoriteActionRequest) (resp *types.FavoriteActionResponse, err error) {
	userId := l.ctx.Value(consts.UserId).(int64)

	if req.ActionType == consts.FavoriteAdd {
		_, err = l.svcCtx.FavoriteRpc.AddFavorite(l.ctx, &favoriterpc.AddFavoriteRequest{
			UserId:  userId,
			VideoId: req.VideoId,
		})
		if err != nil {
			l.Errorf("AddFavorite error: %v", err)
			return
		}

	} else {
		_, err = l.svcCtx.FavoriteRpc.DelFavorite(l.ctx, &favoriterpc.DelFavoriteRequest{
			UserId:  userId,
			VideoId: req.VideoId,
		})
		if err != nil {
			l.Errorf("DelFavorite error: %v", err)
			return
		}
	}

	resp = new(types.FavoriteActionResponse)
	resp.Message = "success"

	return
}

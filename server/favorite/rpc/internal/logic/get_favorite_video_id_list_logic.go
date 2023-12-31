package logic

import (
	"context"
	"kitkot/common/consts"
	"strconv"

	"kitkot/server/favorite/rpc/internal/svc"
	"kitkot/server/favorite/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFavoriteVideoIdListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFavoriteVideoIdListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFavoriteVideoIdListLogic {
	return &GetFavoriteVideoIdListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFavoriteVideoIdListLogic) GetFavoriteVideoIdList(in *pb.GetFavoriteVideoIdListRequest) (resp *pb.GetFavoriteVideoListIdResponse, err error) {
	resp = new(pb.GetFavoriteVideoListIdResponse)
	idListStr, err := l.svcCtx.RedisClient.ZrevrangeCtx(l.ctx, consts.UserFavoriteIdPrefix+strconv.Itoa(int(in.UserId)), 0, -1)
	if err != nil {
		l.Errorf("RedisClient ZrangeCtx error: %v", err)
		return
	}

	idList := make([]int64, 0, len(idListStr))
	for _, idStr := range idListStr {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			l.Errorf("strconv.Atoi error: %v", err)
			return nil, err
		}
		idList = append(idList, int64(id))
	}
	resp.VideoIdList = idList

	return
}

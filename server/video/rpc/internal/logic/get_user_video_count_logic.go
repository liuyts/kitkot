package logic

import (
	"context"
	"kitkot/common/consts"
	"strconv"

	"kitkot/server/video/rpc/internal/svc"
	"kitkot/server/video/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserVideoCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserVideoCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserVideoCountLogic {
	return &GetUserVideoCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserVideoCountLogic) GetUserVideoCount(in *pb.GetUserVideoCountRequest) (resp *pb.GetUserVideoCountResponse, err error) {
	userIdStr := strconv.Itoa(int(in.UserId))
	count, err := l.svcCtx.RedisClient.ZcardCtx(l.ctx, consts.UserVideoRankPrefix+userIdStr)
	if err != nil {
		l.Errorf("Redis Get user video count error: %v", err)
		return nil, err
	}
	resp = new(pb.GetUserVideoCountResponse)
	resp.Count = int64(count)

	return
}

package logic

import (
	"context"

	"kitkot/server/apis/internal/svc"
	"kitkot/server/apis/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RelationActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRelationActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationActionLogic {
	return &RelationActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RelationActionLogic) RelationAction(req *types.RelationActionRequest) (resp *types.RelationActionResponse, err error) {
	// todo: add your logic here and delete this line

	resp = new(types.RelationActionResponse)

	return
}

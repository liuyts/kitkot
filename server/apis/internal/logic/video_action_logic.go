package logic

import (
	"context"
	"kitkot/common/utils"
	"mime/multipart"

	"kitkot/server/apis/internal/svc"
	"kitkot/server/apis/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type VideoActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVideoActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VideoActionLogic {
	return &VideoActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VideoActionLogic) VideoAction(req *types.VideoActionRequest, file *multipart.FileHeader) (resp *types.VideoActionResponse, err error) {
	// todo: add your logic here and delete this line
	err = utils.SaveUploadedFile(file, "D:/Desktop/testFile"+"/"+file.Filename)
	if err != nil {
		return nil, err
	}
	resp = new(types.VideoActionResponse)
	resp.Message = "success"

	return
}

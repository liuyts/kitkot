package logic

import (
	"context"
	"errors"
	"github.com/minio/minio-go/v7"
	"kitkot/common/consts"
	"kitkot/common/utils"
	"kitkot/server/apis/internal/svc"
	"kitkot/server/apis/internal/types"
	"kitkot/server/video/rpc/videorpc"
	"mime/multipart"
	"os"
	"path"

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
	//// 计算文件hash
	//fileHash, err := utils.CalcFileHash(file)
	//if err != nil {
	//	return nil, err
	//}
	//l.Infof("VideoHash:%v", fileHash)
	//// 从redis中查询文件是否存在

	ext := path.Ext(file.Filename)
	if ext != ".mp4" {
		return nil, errors.New("只支持mp4格式")
	}

	playName := utils.UUID() + ext
	coverName := utils.UUID() + ".jpg"
	VideoTmpPath := consts.FileTmpPath + "/" + playName
	CoverTmpPath := consts.FileTmpPath + "/" + coverName
	// 暂存到本地
	err = utils.SaveUploadedFile(file, VideoTmpPath)
	if err != nil {
		return nil, err
	}
	// 截取封面
	err = l.CaptureCover(VideoTmpPath, CoverTmpPath)
	if err != nil {
		return nil, err
	}
	// 将视频和封面上传到minio
	_, err = l.svcCtx.MinioClient.FPutObject(context.Background(), l.svcCtx.Config.MinioConf.BucketName, playName, VideoTmpPath, minio.PutObjectOptions{})
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.MinioClient.FPutObject(context.Background(), l.svcCtx.Config.MinioConf.BucketName, coverName, CoverTmpPath, minio.PutObjectOptions{})
	if err != nil {
		return nil, err
	}
	// 删除本地暂存
	err = os.Remove(VideoTmpPath)
	if err != nil {
		return nil, err
	}
	err = os.Remove(CoverTmpPath)
	if err != nil {
		return nil, err
	}

	playURL := l.svcCtx.MinioClient.EndpointURL().String() + "/" + l.svcCtx.Config.MinioConf.BucketName + "/" + playName
	CoverURL := l.svcCtx.MinioClient.EndpointURL().String() + "/" + l.svcCtx.Config.MinioConf.BucketName + "/" + coverName

	userId := l.ctx.Value(consts.UserId).(int64)
	_, err = l.svcCtx.VideoRpc.PublishVideo(l.ctx, &videorpc.PublishVideoRequest{
		AuthorId: userId,
		Title:    req.Title,
		PlayUrl:  playURL,
		CoverUrl: CoverURL,
	})
	if err != nil {
		return nil, err
	}

	resp = new(types.VideoActionResponse)
	resp.Message = "发布成功"

	return
}

func (l *VideoActionLogic) CaptureCover(VideoTmpPath, CoverTmpPath string) error {
	return utils.CmdWithDirNoOut(consts.FileTmpPath,
		"ffmpeg",
		"-i", VideoTmpPath,
		"-y",
		"-f", "image2",
		"-frames", "1",
		"-ss", "1",
		CoverTmpPath,
	)
}

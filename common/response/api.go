package response

import (
	"context"
	"kitkot/common/response/errcode"
	"net/http"
)

type ErrrResponse struct {
	Code int    `json:"status_code"`
	Msg  string `json:"status_msg"`
}

func ErrHandlerCtx() func(context.Context, error) (int, any) {
	return func(ctx context.Context, err error) (int, any) {
		// 一律返回200
		errResp := &ErrrResponse{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		}
		if errcode.IsGrpcError(err) {
			errResp.Code = errcode.CodeFromGrpcError(err)
			return http.StatusOK, errResp
		}

		return http.StatusOK, errResp
	}
}

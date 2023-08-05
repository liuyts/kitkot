package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"kitkot/server/apis/internal/logic"
	"kitkot/server/apis/internal/svc"
	"kitkot/server/apis/internal/types"
)

func VideoActionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.VideoActionRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		_, fileHeader, err := r.FormFile("data")
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewVideoActionLogic(r.Context(), svcCtx)
		resp, err := l.VideoAction(&req, fileHeader)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

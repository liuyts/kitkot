package handler

import (
	"kitkot/server/chat/api/internal/logic"
	"kitkot/server/chat/api/internal/svc"
	"kitkot/server/chat/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func MessageActionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MessageActionRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewMessageActionLogic(r.Context(), svcCtx)
		resp, err := l.MessageAction(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

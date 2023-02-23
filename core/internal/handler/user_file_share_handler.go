package handler

import (
	"net/http"

	"CloudStorage/core/internal/logic"
	"CloudStorage/core/internal/svc"
	"CloudStorage/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserFileShareHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserFileShareRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUserFileShareLogic(r.Context(), svcCtx)
		resp, err := l.UserFileShare(&req, r.Header.Get("UserIdentity"))
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

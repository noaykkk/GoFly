package handler

import (
	"net/http"

	"CloudStorage/core/internal/logic"
	"CloudStorage/core/internal/svc"
	"CloudStorage/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserFileRenameHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserFileRenameRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUserFileRenameLogic(r.Context(), svcCtx)
		resp, err := l.UserFileRename(&req, r.Header.Get("UserIdentity"))
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

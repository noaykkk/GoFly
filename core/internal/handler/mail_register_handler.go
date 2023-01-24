package handler

import (
	"net/http"

	"CloudStorage/core/internal/logic"
	"CloudStorage/core/internal/svc"
	"CloudStorage/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func MailRegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MailVerificationCodeRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewMailRegisterLogic(r.Context(), svcCtx)
		resp, err := l.MailRegister(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

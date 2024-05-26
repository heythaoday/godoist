package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"godoist/internal/logic"
	"godoist/internal/svc"
)

func UserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUserLogic(r.Context(), svcCtx)
		resp, err := l.User()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

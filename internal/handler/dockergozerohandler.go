package handler

import (
	"net/http"

	"docker_gozero/internal/logic"
	"docker_gozero/internal/svc"
	"docker_gozero/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func Docker_gozeroHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewDocker_gozeroLogic(r.Context(), svcCtx)
		resp, err := l.Docker_gozero(&req)

		
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

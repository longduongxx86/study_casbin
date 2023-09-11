package handler

import (
	"context"
	"docker_gozero/internal/logic"
	"docker_gozero/internal/svc"
	"docker_gozero/internal/types"
	"fmt"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserLogin

		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		user, err := l.Login(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			token, errWhenCreateToken := CreateJWT(user)
			if errWhenCreateToken != nil {
				httpx.ErrorCtx(r.Context(), w, errWhenCreateToken)
			}

			go func() {
				svcCtx.RedisDB.SetCtx(
					context.Background(),
					fmt.Sprintf("1|%s|%d", user.RoleID, user.Id),
					token,
				)
			}()

			httpx.OkJsonCtx(r.Context(), w, types.Response{Message: token})
		}
	}
}

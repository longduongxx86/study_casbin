package handler

import (
	"context"
	"docker_gozero/internal/logic"
	"docker_gozero/internal/svc"
	"docker_gozero/internal/types"
	"errors"
	"fmt"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			http.Error(w, "Authorization header is missing", http.StatusUnauthorized)
			return
		}

		token := auth[len("Bearer "):]

		i, err := ParseJWT(token)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		}

		user := i.(types.UserIden)

		session, err := svcCtx.RedisDB.GetCtx(
			context.Background(),
			fmt.Sprintf("1|%s|%d", user.RoleID, user.Id),
		)

		if err != nil {
			httpx.ErrorCtx(r.Context(), w, errors.New("khonghople"))
		}

		l := logic.NewGetLogic(r.Context(), svcCtx)
		l.Get(user.Id)

		if session != "" {
			httpx.ErrorCtx(r.Context(), w, errors.New("khonghople"))
		}

		itf, err := ParseJWT(token)

		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJson(w, itf)
		}
	}
}

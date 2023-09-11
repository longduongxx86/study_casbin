package logic

import (
	"context"
	"docker_gozero/internal/svc"
	"docker_gozero/internal/types"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(userLogin *types.UserLogin) (*types.UserIden, error) {
	var user types.UserIden

	err := l.svcCtx.MysqlDB.QueryRow(
		&user,
		"SELECT id, role_id FROM accounts WHERE username=? AND password=?",
		userLogin.Username, userLogin.Password,
	)

	if err != nil {
		fmt.Println("error:", err)
		return nil, err
	}

	return &user, nil
}

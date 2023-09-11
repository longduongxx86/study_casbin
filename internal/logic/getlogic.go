package logic

import (
	"context"
	"docker_gozero/internal/svc"
	"docker_gozero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogic {
	return &GetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLogic) Get(userId int) (resp *types.UserInfo, err error) {
	err = l.svcCtx.MysqlDB.QueryRow(&resp, "SELECT username, email, role_id FROM accounts WHERE userid = ?", userId)

	return resp, err
}

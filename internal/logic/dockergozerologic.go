package logic

import (
	"context"

	"docker_gozero/internal/svc"
	"docker_gozero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Docker_gozeroLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDocker_gozeroLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Docker_gozeroLogic {
	return &Docker_gozeroLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Docker_gozeroLogic) Docker_gozero(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}

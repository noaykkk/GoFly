package logic

import (
	"context"

	"CloudStorage/core/internal/svc"
	"CloudStorage/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CoreLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CoreLogic {
	return &CoreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CoreLogic) Core(req *types.Request) (resp *types.Response, err error) {
	resp = new(types.Response)
	resp.Message = "Hello Cloud Disk"
	return
}

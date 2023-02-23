package logic

import (
	"CloudStorage/core/models"
	"context"
	"errors"

	"CloudStorage/core/internal/svc"
	"CloudStorage/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileMoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileMoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileMoveLogic {
	return &UserFileMoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileMoveLogic) UserFileMove(req *types.UserFileMoveRequest, userIdentity string) (resp *types.UserFileMoveResponse, err error) {
	parentData := new(models.UserRepository)
	has, err := l.svcCtx.Engine.Where("identity = ? AND user_identity = ?", req.ParentIdnetity, userIdentity).Get(parentData)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("folder not exist")
	}

	_, err = l.svcCtx.Engine.Where("identity = ?", req.Idnetity).Update(models.UserRepository{
		ParentId: int64(parentData.Id),
	})

	return
}

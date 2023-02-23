package logic

import (
	"CloudStorage/core/helper"
	"CloudStorage/core/internal/svc"
	"CloudStorage/core/internal/types"
	"CloudStorage/core/models"
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileShareLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileShareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileShareLogic {
	return &UserFileShareLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileShareLogic) UserFileShare(req *types.UserFileShareRequest, userIdentity string) (resp *types.UserFileShareResponse, err error) {
	uuid := helper.GenerateUUID()
	ur := new(models.UserRepository)
	has, err := l.svcCtx.Engine.Where("identity = ?", req.UserRepositoryIdentity).Get(ur)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("user repository not found")
	}
	data := &models.ShareBasic{
		Identity:               uuid,
		UserIdentity:           userIdentity,
		UserRepositoryIdentity: req.UserRepositoryIdentity,
		RepositoryIdentity:     ur.RepositoryIdentity,
		ExpiredTime:            req.ExpiredTime,
	}
	_, err = l.svcCtx.Engine.Insert(data)
	if err != nil {
		return
	}
	resp = &types.UserFileShareResponse{
		Identity: uuid,
	}

	return
}

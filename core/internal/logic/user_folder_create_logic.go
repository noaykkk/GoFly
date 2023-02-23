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

type UserFolderCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFolderCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFolderCreateLogic {
	return &UserFolderCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFolderCreateLogic) UserFolderCreate(req *types.UserFolderCreateRequest, userIdentity string) (resp *types.UserFolderCreateResponse, err error) {
	cnt, err := l.svcCtx.Engine.Where("name = ? AND parent_id = ?", req.Name, req.ParentId).Count(new(models.UserRepository))
	if err != nil {
		return nil, err
	}
	if cnt > 0 {
		return nil, errors.New("name exist")
	}
	data := &models.UserRepository{
		Identity:     helper.GenerateUUID(),
		UserIdentity: userIdentity,
		ParentId:     req.ParentId,
		Name:         req.Name,
	}
	_, err = l.svcCtx.Engine.Insert(data)
	if err != nil {
		return
	}
	return
}

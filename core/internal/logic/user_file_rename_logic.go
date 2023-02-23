package logic

import (
	"CloudStorage/core/models"
	"context"
	"errors"

	"CloudStorage/core/internal/svc"
	"CloudStorage/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileRenameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileRenameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileRenameLogic {
	return &UserFileRenameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileRenameLogic) UserFileRename(req *types.UserFileRenameRequest, userIdentity string) (resp *types.UserFileRenameResponse, err error) {
	cnt, err := l.svcCtx.Engine.Where("name = ? AND parent_id = (SELECT parent_id FROM user_repository ur WHERE ur.identity = ?)", req.Name, req.Identity).Count(new(models.UserRepository))
	if err != nil {
		return nil, err
	}
	if cnt > 0 {
		return nil, errors.New("name exist")
	}

	data := &models.UserRepository{Name: req.Name}
	_, err = l.svcCtx.Engine.Where("identity = ? AND user_identity = ? ", req.Identity, userIdentity).Update(data)
	if err != nil {
		return
	}

	return
}

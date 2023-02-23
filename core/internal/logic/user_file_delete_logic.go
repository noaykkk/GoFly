package logic

import (
	"CloudStorage/core/models"
	"context"
	"fmt"

	"CloudStorage/core/internal/svc"
	"CloudStorage/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileDeleteLogic {
	return &UserFileDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileDeleteLogic) UserFileDelete(req *types.UserFileDeleteRequest, userIdentity string) (resp *types.UserFileDeleteResponse, err error) {
	rp := new(models.RepositoryPool)
	_, err = l.svcCtx.Engine.Select("size").Where("identity = (SELECT repository_identity FROM user_repository WHERE identity = ? LIMIT 1)", req.Identity).Get(rp)
	if err != nil {
		fmt.Println(err)
		return
	}

	if rp.Size > 0 {
		_, err = l.svcCtx.Engine.Exec("UPDATE user_basic SET now_volume = now_volume - ? WHERE identity = ?", rp.Size, userIdentity)
		if err != nil {
			return
		}
	}

	_, err = l.svcCtx.Engine.Where("user_identity = ? AND identity = ?", userIdentity, req.Identity).Delete(new(models.UserRepository))
	return
}

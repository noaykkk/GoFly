package logic

import (
	"CloudStorage/core/helper"
	"CloudStorage/core/models"
	"context"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"log"

	"CloudStorage/core/internal/svc"
	"CloudStorage/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterRequest) (resp *types.UserRegisterResponse, err error) {
	// todo: add your logic here and delete this line
	code, err := l.svcCtx.RDB.Get(l.ctx, req.Email).Result()
	if err != nil {
		return nil, errors.New("Empty Verification Code")
	}
	if code != req.Code {
		err = errors.New("Wrong Verification Code")
	}
	cnt, err := l.svcCtx.Engine.Where("username = ?", req.Name).Count(new(models.UserBasic))
	if err != nil {
		return nil, err
	}
	if cnt > 0 {
		err = errors.New("Username already exists")
		return
	}
	user := &models.UserBasic{
		Identity: helper.GenerateUUID(),
		Name:     req.Name,
		Password: helper.Md5(req.Password),
		Email:    req.Email,
	}
	insert, err := l.svcCtx.Engine.Insert(user)
	if err != nil {
		return nil, err
	}
	log.Println("insert user:", insert)
	return
}

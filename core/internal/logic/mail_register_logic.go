package logic

import (
	"CloudStorage/core/define"
	"CloudStorage/core/helper"
	"CloudStorage/core/models"
	"context"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"time"

	"CloudStorage/core/internal/svc"
	"CloudStorage/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MailRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMailRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MailRegisterLogic {
	return &MailRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MailRegisterLogic) MailRegister(req *types.MailVerificationCodeRequest) (resp *types.MailVerificationCodeResponse, err error) {
	// todo: add your logic here and delete this line
	cnt, err := l.svcCtx.Engine.Where("email = ?", req.Email).Count(new(models.UserBasic))
	if err != nil {
		return
	}
	if cnt > 0 {
		err = errors.New("This e-mail address has been already registered")
		return
	}
	code := helper.RadomCode
	l.svcCtx.RDB.Set(l.ctx, req.Email, code, time.Second*time.Duration(define.CodeExpire))
	err = helper.MailSend(req.Email, "123456")
	if err != nil {
		return nil, err
	}
	return
}

package user

import (
	"context"

	"github.com/PIPIKAI/fake-ins-v2/service/usercenter/api/internal/svc"
	"github.com/PIPIKAI/fake-ins-v2/service/usercenter/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditInfoLogic {
	return &EditInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditInfoLogic) EditInfo() (resp *types.UserDto, err error) {
	// todo: add your logic here and delete this line

	return
}

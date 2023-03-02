package user

import (
	"context"

	"github.com/PIPIKAI/fake-ins-v2/service/usercenter/api/internal/svc"
	"github.com/PIPIKAI/fake-ins-v2/service/usercenter/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetByUidLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetByUidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetByUidLogic {
	return &GetByUidLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetByUidLogic) GetByUid() (resp *types.UserDto, err error) {
	// todo: add your logic here and delete this line

	return
}

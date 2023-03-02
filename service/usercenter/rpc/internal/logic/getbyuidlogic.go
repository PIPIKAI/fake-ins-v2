package logic

import (
	"context"

	"github.com/PIPIKAI/fake-ins-v2/service/usercenter/rpc/gen/pb"
	"github.com/PIPIKAI/fake-ins-v2/service/usercenter/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetByUidLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetByUidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetByUidLogic {
	return &GetByUidLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetByUidLogic) GetByUid(in *pb.Message) (*pb.UserDto, error) {
	// todo: add your logic here and delete this line

	return &pb.UserDto{}, nil
}

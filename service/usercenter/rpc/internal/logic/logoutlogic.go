package logic

import (
	"context"

	"github.com/PIPIKAI/fake-ins-v2/service/usercenter/rpc/gen/pb"
	"github.com/PIPIKAI/fake-ins-v2/service/usercenter/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LogoutLogic) Logout(in *pb.Message) (*pb.Message, error) {
	// todo: add your logic here and delete this line

	return &pb.Message{}, nil
}

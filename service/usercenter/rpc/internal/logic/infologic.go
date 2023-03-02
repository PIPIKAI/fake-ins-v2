package logic

import (
	"context"

	"github.com/PIPIKAI/fake-ins-v2/service/usercenter/rpc/gen/pb"
	"github.com/PIPIKAI/fake-ins-v2/service/usercenter/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type InfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InfoLogic {
	return &InfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InfoLogic) Info(in *pb.Message) (*pb.UserDto, error) {
	// todo: add your logic here and delete this line

	return &pb.UserDto{}, nil
}

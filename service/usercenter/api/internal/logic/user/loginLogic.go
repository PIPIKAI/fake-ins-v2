package user

import (
	"context"
	"log"

	"github.com/PIPIKAI/fake-ins-v2/service/usercenter/api/internal/svc"
	"github.com/PIPIKAI/fake-ins-v2/service/usercenter/api/internal/types"
	"github.com/PIPIKAI/fake-ins-v2/service/usercenter/rpc/usercenter"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// todo: add your logic here and delete this line
	rpcRes, err := l.svcCtx.UserCenterRpc.Login(l.ctx, &usercenter.LoginReq{
		EmailOrPhoneOrUsername: req.EmailOrPhoneOrUsername,
		Password:               req.Password,
	})
	if err != nil {
		log.Println("err: ", err)
		return nil, err
	}
	var res types.LoginResp
	copier.Copy(&res, rpcRes)
	return &res, nil
}

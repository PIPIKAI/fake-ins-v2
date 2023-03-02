package logic

import (
	"context"

	"github.com/PIPIKAI/fake-ins-v2/common/errorx"
	"github.com/PIPIKAI/fake-ins-v2/service/usercenter/model"
	"github.com/PIPIKAI/fake-ins-v2/service/usercenter/rpc/gen/pb"
	"github.com/PIPIKAI/fake-ins-v2/service/usercenter/rpc/internal/svc"
	"github.com/PIPIKAI/fake-ins-v2/service/usercenter/rpc/usercenter"
	"golang.org/x/crypto/bcrypt"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.LoginReq) (*pb.LoginResp, error) {
	// todo: add your logic here and delete this line
	user, err := l.svcCtx.UserModel.FindOneByUserName(l.ctx, in.EmailOrPhoneOrUsername)
	if err != nil && err != model.ErrNotFound {
		return nil, errorx.NewSystemError("数据库查询失败")
	}

	if user == nil {
		return nil, errorx.NewDefaultError("用户不存在")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PassWord), []byte(in.Password)); err != nil {
		return nil, errorx.NewDefaultError("密码错误")
	}

	// 产生gwt
	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	tokenResp, err := generateTokenLogic.GenerateToken(&usercenter.GenerateTokenReq{
		UserId: user.Id,
	})
	if err != nil {
		return nil, errorx.NewDefaultError("GenerateToken userId")
	}
	return &pb.LoginResp{
		AccessToken:  tokenResp.AccessToken,
		AccessExpire: tokenResp.AccessExpire,
		RefreshAfter: tokenResp.RefreshAfter,
	}, nil
}

// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"github.com/PIPIKAI/fake-ins-v2/service/usercenter/rpc/gen/pb"
	"github.com/PIPIKAI/fake-ins-v2/service/usercenter/rpc/internal/logic"
	"github.com/PIPIKAI/fake-ins-v2/service/usercenter/rpc/internal/svc"
)

type UsercenterServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedUsercenterServer
}

func NewUsercenterServer(svcCtx *svc.ServiceContext) *UsercenterServer {
	return &UsercenterServer{
		svcCtx: svcCtx,
	}
}

func (s *UsercenterServer) Login(ctx context.Context, in *pb.LoginReq) (*pb.LoginResp, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *UsercenterServer) Register(ctx context.Context, in *pb.RegisterReq) (*pb.RegisterResp, error) {
	l := logic.NewRegisterLogic(ctx, s.svcCtx)
	return l.Register(in)
}

func (s *UsercenterServer) Info(ctx context.Context, in *pb.Message) (*pb.UserDto, error) {
	l := logic.NewInfoLogic(ctx, s.svcCtx)
	return l.Info(in)
}

func (s *UsercenterServer) Logout(ctx context.Context, in *pb.Message) (*pb.Message, error) {
	l := logic.NewLogoutLogic(ctx, s.svcCtx)
	return l.Logout(in)
}

func (s *UsercenterServer) GetByUserName(ctx context.Context, in *pb.Message) (*pb.UserDto, error) {
	l := logic.NewGetByUserNameLogic(ctx, s.svcCtx)
	return l.GetByUserName(in)
}

func (s *UsercenterServer) GetByUid(ctx context.Context, in *pb.Message) (*pb.UserDto, error) {
	l := logic.NewGetByUidLogic(ctx, s.svcCtx)
	return l.GetByUid(in)
}

func (s *UsercenterServer) GenerateToken(ctx context.Context, in *pb.GenerateTokenReq) (*pb.GenerateTokenResp, error) {
	l := logic.NewGenerateTokenLogic(ctx, s.svcCtx)
	return l.GenerateToken(in)
}

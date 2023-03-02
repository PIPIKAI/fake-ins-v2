package svc

import (
	"github.com/PIPIKAI/fake-ins-v2/service/usercenter/api/internal/config"
	"github.com/PIPIKAI/fake-ins-v2/service/usercenter/rpc/usercenter"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	UserCenterRpc usercenter.Usercenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UserCenterRpc: usercenter.NewUsercenter(zrpc.MustNewClient(c.UserCenterRpcConf)),
	}
}

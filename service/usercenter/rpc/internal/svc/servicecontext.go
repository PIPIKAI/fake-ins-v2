package svc

import (
	"github.com/PIPIKAI/fake-ins-v2/service/usercenter/model"
	"github.com/PIPIKAI/fake-ins-v2/service/usercenter/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis

	UserModel model.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config: c,
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
		UserModel: model.NewUsersModel(sqlConn, c.Cache),
	}
}

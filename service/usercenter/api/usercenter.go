package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"

	"github.com/PIPIKAI/fake-ins-v2/common/errorx"
	"github.com/PIPIKAI/fake-ins-v2/service/usercenter/api/internal/config"
	"github.com/PIPIKAI/fake-ins-v2/service/usercenter/api/internal/handler"
	"github.com/PIPIKAI/fake-ins-v2/service/usercenter/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/usercenter.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// 自定义错误
	httpx.SetErrorHandlerCtx(func(ctx context.Context, err error) (int, interface{}) {
		switch e := err.(type) {
		case *errorx.CodeError:
			return http.StatusOK, e.Data()
		default:
			return http.StatusInternalServerError, nil
		}
	})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

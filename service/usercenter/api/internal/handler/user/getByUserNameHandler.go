package user

import (
	"net/http"

	"github.com/PIPIKAI/fake-ins-v2/service/usercenter/api/internal/logic/user"
	"github.com/PIPIKAI/fake-ins-v2/service/usercenter/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetByUserNameHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewGetByUserNameLogic(r.Context(), svcCtx)
		resp, err := l.GetByUserName()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

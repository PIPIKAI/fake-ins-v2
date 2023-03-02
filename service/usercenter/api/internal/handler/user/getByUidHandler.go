package user

import (
	"net/http"

	"github.com/PIPIKAI/fake-ins-v2/service/usercenter/api/internal/logic/user"
	"github.com/PIPIKAI/fake-ins-v2/service/usercenter/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetByUidHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewGetByUidLogic(r.Context(), svcCtx)
		resp, err := l.GetByUid()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

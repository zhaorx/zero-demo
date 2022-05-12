package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"user-api/internal/logic/user"
	"user-api/internal/svc"
	"user-api/internal/types"
)

func UserUpateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserUpdateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewUserUpateLogic(r.Context(), svcCtx)
		resp, err := l.UserUpate(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

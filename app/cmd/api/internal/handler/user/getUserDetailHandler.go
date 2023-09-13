package user

import (
	"github.com/wuqianaer/go-zero-oj/app/common/response"
	"net/http"

	"github.com/wuqianaer/go-zero-oj/app/cmd/api/internal/logic/user"
	"github.com/wuqianaer/go-zero-oj/app/cmd/api/internal/svc"
	"github.com/wuqianaer/go-zero-oj/app/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetUserDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserDeitalReq
		if err := httpx.Parse(r, &req); err != nil {
			response.JsonBaseResponseCtx(r.Context(), w, err)
			return
		}

		l := user.NewGetUserDetailLogic(r.Context(), svcCtx)
		resp, err := l.GetUserDetail(&req)
		if err != nil {
			response.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			response.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}

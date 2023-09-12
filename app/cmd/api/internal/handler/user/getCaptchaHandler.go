package user

import (
	"github.com/wuqianaer/go-zero-oj/app/common/pkg/utils"
	"github.com/wuqianaer/go-zero-oj/app/common/response"
	"net/http"

	"github.com/wuqianaer/go-zero-oj/app/cmd/api/internal/logic/user"
	"github.com/wuqianaer/go-zero-oj/app/cmd/api/internal/svc"
)

func GetCaptchaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewGetCaptchaLogic(r.Context(), svcCtx)
		ip := utils.GetIP(r)
		resp, err := l.GetCaptcha(ip)
		if err != nil {
			response.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			response.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}

package problem

import (
	"github.com/wuqianaer/go-zero-oj/app/common/response"
	"net/http"

	"github.com/wuqianaer/go-zero-oj/app/cmd/api/internal/logic/problem"
	"github.com/wuqianaer/go-zero-oj/app/cmd/api/internal/svc"
	"github.com/wuqianaer/go-zero-oj/app/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetProblemDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ProblemDeitalReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := problem.NewGetProblemDetailLogic(r.Context(), svcCtx)
		resp, err := l.GetProblemDetail(&req)
		if err != nil {
			response.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			response.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}

package problem

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/wuqianaer/go-zero-oj/app/common/global"
	"github.com/wuqianaer/go-zero-oj/app/common/response"
	"net/http"

	"github.com/wuqianaer/go-zero-oj/app/cmd/api/internal/logic/problem"
	"github.com/wuqianaer/go-zero-oj/app/cmd/api/internal/svc"
	"github.com/wuqianaer/go-zero-oj/app/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SubmitListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ProblemSubmitListReq
		if err := httpx.Parse(r, &req); err != nil {
			response.JsonBaseResponseCtx(r.Context(), w, err)
			return
		}

		if validateErr := global.Validate.StructCtx(r.Context(), &req); validateErr != nil {
			for _, err := range validateErr.(validator.ValidationErrors) {
				response.JsonBaseResponseCtx(r.Context(), w, errors.New(err.Translate(global.Translator)))
				return
			}
		}

		l := problem.NewSubmitListLogic(r.Context(), svcCtx)
		resp, err := l.SubmitList(&req)
		if err != nil {
			response.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			response.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}

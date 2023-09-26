package admin

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/wuqianaer/go-zero-oj/app/common/consts"
	"github.com/wuqianaer/go-zero-oj/app/common/global"
	"github.com/wuqianaer/go-zero-oj/app/common/response"
	"net/http"

	"github.com/wuqianaer/go-zero-oj/app/cmd/api/internal/logic/admin"
	"github.com/wuqianaer/go-zero-oj/app/cmd/api/internal/svc"
	"github.com/wuqianaer/go-zero-oj/app/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateProblemHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateProblemReq
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
		l := admin.NewUpdateProblemLogic(r.Context(), svcCtx)
		err := l.UpdateProblem(&req)
		if err != nil {
			response.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			response.JsonBaseResponseCtx(r.Context(), w, response.CodeResponse{
				Code: consts.Code_Success,
				Msg:  "修改成功", // TODO 添加你需要返回的数据
			})
		}
	}
}

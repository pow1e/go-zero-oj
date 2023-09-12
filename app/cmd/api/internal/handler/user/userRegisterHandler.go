package user

import (
	"github.com/wuqianaer/go-zero-oj/app/common/consts"
	"github.com/wuqianaer/go-zero-oj/app/common/pkg/utils"
	"github.com/wuqianaer/go-zero-oj/app/common/response"
	"net/http"

	"github.com/wuqianaer/go-zero-oj/app/cmd/api/internal/logic/user"
	"github.com/wuqianaer/go-zero-oj/app/cmd/api/internal/svc"
	"github.com/wuqianaer/go-zero-oj/app/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserRegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserRegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 校验密码
		if resp := verifyPassword(req.Password, req.ConfirmPassword); resp != nil {
			response.JsonBaseResponseCtx(r.Context(), w, resp)
			return
		}

		if resp := verifyUserName(req.UserName); resp != nil {
			response.JsonBaseResponseCtx(r.Context(), w, resp)
			return
		}

		l := user.NewUserRegisterLogic(r.Context(), svcCtx)
		err := l.UserRegister(&req, utils.GetIP(r))
		if err != nil {
			response.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			response.JsonBaseResponseCtx(r.Context(), w, response.CodeResponse{
				Code: consts.Code_Success,
				Msg:  "注册成功",
			})
		}
	}
}

// verifyPassword 校验密码
func verifyPassword(password, confirmPassword string) *response.CodeResponse {
	resp := &response.CodeResponse{}
	resp.Code = consts.Code_Invaild
	if len(password) < consts.MinPasswordLength && len(password) > consts.MaxPasswordLength {
		resp.Msg = consts.ErrPasswordLength
		return resp
	}

	if len(confirmPassword) < 6 || len(confirmPassword) > 11 {
		resp.Msg = consts.ErrPasswordLength
		return resp
	}

	if password != confirmPassword {
		resp.Msg = consts.ErrConfirmPassword
		return resp
	}
	return nil
}

// verifyUserName 校验用户名
func verifyUserName(username string) *response.CodeResponse {
	resp := &response.CodeResponse{}
	resp.Code = consts.Code_Invaild
	name := []rune(username)
	// 判断长度
	if len(name) < consts.MinUserNameLength || len(name) > consts.MaxUserNameLength {
		resp.Msg = consts.ErrUserNameLength
		return resp
	}
	return nil
}

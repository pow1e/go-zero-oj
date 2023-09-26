package middleware

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/wuqianaer/go-zero-oj/app/cmd/api/internal/config"
	"github.com/wuqianaer/go-zero-oj/app/common/consts"
	"github.com/wuqianaer/go-zero-oj/app/common/global"
	"github.com/wuqianaer/go-zero-oj/app/common/response"
	"net/http"
	"strings"
)

type AuthorizationMiddleware struct {
	authConfig config.Auth
}

func NewAuthorizationMiddleware(authConfig config.Auth) *AuthorizationMiddleware {
	return &AuthorizationMiddleware{authConfig: authConfig}
}

func (m *AuthorizationMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var token string
		var resp response.CodeResponse
		// 1.判断请求头是否为空
		if token = r.Header.Get(consts.AuthorizationHeader); token == "" {
			resp.Code = consts.Code_Invaild
			resp.Msg = consts.ErrAuthorization
			response.JsonBaseResponseCtx(r.Context(), w, resp)
			return
		}

		// 2.判断token是否合法
		clams, err := m.parseToken(token)
		if err != nil {
			resp.Code = consts.Code_ErrAuthorization
			resp.Msg = consts.ErrTokenInvalid
			response.JsonBaseResponseCtx(r.Context(), w, err)
			return
		}

		// 判断是否是管理员路由
		path := r.URL.Path
		lastIndex := strings.LastIndex(path, "/")
		// 如果不存在
		route := path[lastIndex+1:]
		if _, ok := consts.AdminRouteMap[route]; !ok {
			resp.Code = consts.Code_ErrAuthorization
			resp.Msg = consts.ErrPermissions
			response.JsonBaseResponseCtx(r.Context(), w, resp)
			return
		} else {
			if clams.BaseClaim.ID != consts.AdminID {
				resp.Code = consts.Code_ErrAuthorization
				resp.Msg = consts.ErrPermissions
				response.JsonBaseResponseCtx(r.Context(), w, resp)
				return
			}
		}

		ctxWithValue := context.WithValue(r.Context(), consts.UserInfo, clams)

		// 放行
		next(w, r.WithContext(ctxWithValue))
	}
}

// AuthorizationMiddleware 解析token
func (m *AuthorizationMiddleware) parseToken(tokenString string) (*global.CustomClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(tokenString, &global.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(m.authConfig.AccessSecret), nil
	})

	if err != nil {
		var valueErr *jwt.ValidationError
		if errors.As(err, &valueErr) {
			if valueErr.Errors == jwt.ValidationErrorExpired {
				return nil, errors.New(consts.ErrTokenExpired)
			} else {
				return nil, errors.New(consts.ErrTokenInvalid)
			}
		}
	}

	if err != nil {
		return nil, err
	}
	if jwtToken != nil {
		// 对claims进行断言
		if claims, ok := jwtToken.Claims.(*global.CustomClaims); ok {
			return claims, nil
		}
		return nil, errors.New(consts.ErrTokenInvalid)
	}
	return nil, nil
}

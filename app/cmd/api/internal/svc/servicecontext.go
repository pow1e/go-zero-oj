package svc

import (
	"github.com/wuqianaer/go-zero-oj/app/cmd/api/internal/config"
	"github.com/wuqianaer/go-zero-oj/app/cmd/api/internal/middleware"
	"github.com/wuqianaer/go-zero-oj/app/cmd/api/internal/repo"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config            config.Config
	Repository        *repo.Repository
	CaptchaThrottling rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	repository := repo.NewRepository(c)
	return &ServiceContext{
		Config:     c,
		Repository: repository,
		CaptchaThrottling: middleware.NewCaptchaThrottlingMiddleware(
			repository.Redis,
			c.Captcha.MaxTime,
		).Handle,
	}
}

package user

import (
	"context"
	"errors"
	"github.com/mojocn/base64Captcha"
	"github.com/wuqianaer/go-zero-oj/app/common/consts"

	"github.com/wuqianaer/go-zero-oj/app/cmd/api/internal/svc"
	"github.com/wuqianaer/go-zero-oj/app/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

var store = base64Captcha.DefaultMemStore

type GetCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCaptchaLogic {
	return &GetCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCaptchaLogic) GetCaptcha(clientIP string) (resp *types.CaptchaResp, err error) {
	var key = consts.GetCaptchaCachePrefix + clientIP
	// 删除上一次的验证码
	if _, err = l.svcCtx.Repository.Redis.DelCtx(l.ctx, key); err != nil {
		l.Logger.Info("删除验证码失败", err.Error())
		return nil, errors.New(consts.ErrCaptcha)
	}

	captchaConfig := l.svcCtx.Config.Captcha
	// 生成验证码
	driver := base64Captcha.NewDriverDigit(captchaConfig.Height, captchaConfig.Width, captchaConfig.KeyLong, 0.7, 80)
	c := base64Captcha.NewCaptcha(driver, store)

	captchaID, b64s, err := c.Generate()
	if err != nil {
		l.Logger.Info("获取验证码失败！", err.Error())
		return nil, errors.New(consts.ErrCreateCaptcha)
	}

	code := store.Get(captchaID, true)
	err = l.svcCtx.Repository.Redis.Setex(consts.GetCaptchaCachePrefix+clientIP, code, captchaConfig.Timeout)
	if err != nil {
		l.Logger.Info("验证码存放redis失败", err.Error())
		return nil, errors.New(consts.ErrCreateCaptcha)
	}

	return &types.CaptchaResp{
		CaptchaID:   captchaID,
		PictureData: b64s,
	}, nil
}

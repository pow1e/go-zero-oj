package user

import (
	"context"
	"errors"
	"github.com/wuqianaer/go-zero-oj/app/cmd/api/internal/dal/model"
	"github.com/wuqianaer/go-zero-oj/app/cmd/api/internal/repo"
	"github.com/wuqianaer/go-zero-oj/app/common/consts"
	"github.com/wuqianaer/go-zero-oj/app/common/pkg/utils"
	"gorm.io/gorm"

	"github.com/wuqianaer/go-zero-oj/app/cmd/api/internal/svc"
	"github.com/wuqianaer/go-zero-oj/app/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterReq, ip string) error {
	var key = consts.GetCaptchaCachePrefix + ip
	if len(req.CaptchaCode) != l.svcCtx.Config.Captcha.KeyLong {
		return errors.New(consts.ErrCaptcha)
	}

	// 从redis中获取验证码
	captchaCode, err := repo.GlobalRepository.Redis.GetCtx(l.ctx, key)
	if err != nil {
		l.Logger.Error("从redis中获取验证码失败", err.Error())
		return errors.New(consts.ErrUnKnow)
	}

	// 验证码错误则刷新验证码 并返回
	if captchaCode != req.CaptchaCode {
		return errors.New(consts.ErrCaptcha)
	}

	u := l.svcCtx.Repository.Model.User
	findUser, err := u.WithContext(l.ctx).Where(u.Name.Eq(req.UserName)).First()

	// 用户不存在
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		req.Password, err = utils.Encrypt(req.Password)
		err = u.WithContext(l.ctx).Create(&model.User{Name: req.UserName, Password: req.Password})
		if err != nil {
			l.Logger.Debug("创建用户失败", err.Error())
			return errors.New(consts.ErrUnKnow)
		}
		return nil
	}

	// 用户已存在
	if findUser != nil {
		return errors.New(consts.ErrUserExist)
	}

	// 注册成功就将redis删除
	if err = l.DeleteCaptchaCache(key); err != nil {
		l.Logger.Info("删除redis中的验证码失败", err.Error())
		return errors.New(consts.ErrUnKnow)
	}

	return nil
}

// DeleteCaptchaCache 删除redis
func (l *UserRegisterLogic) DeleteCaptchaCache(key string) error {
	_, err := l.svcCtx.Repository.Redis.DelCtx(l.ctx, key)
	if err != nil {
		return err
	}
	return nil
}

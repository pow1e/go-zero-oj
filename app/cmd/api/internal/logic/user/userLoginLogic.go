package user

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/wuqianaer/go-zero-oj/app/cmd/api/internal/repo/dao"
	"github.com/wuqianaer/go-zero-oj/app/cmd/api/internal/svc"
	"github.com/wuqianaer/go-zero-oj/app/cmd/api/internal/types"
	"github.com/wuqianaer/go-zero-oj/app/common/consts"
	"github.com/wuqianaer/go-zero-oj/app/common/global"

	"github.com/wuqianaer/go-zero-oj/app/common/pkg/utils"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"time"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.UserLoginReq, ip string) (resp *types.UserLoginResp, err error) {
	var key = consts.GetCaptchaCachePrefix + ip
	if len(req.CaptchaCode) != l.svcCtx.Config.Captcha.KeyLong {
		return nil, errors.New(consts.ErrCaptcha)
	}

	// 从redis中获取验证码
	captchaCode, err := l.svcCtx.Repository.Redis.GetCtx(l.ctx, key)
	if err != nil {
		l.Logger.Error("从redis中获取验证码失败", err.Error())
		return nil, errors.New(consts.ErrUnKnow)
	}

	if captchaCode != req.CaptchaCode {
		return nil, errors.New(consts.ErrCaptcha)
	}

	// 查看当前用户
	userDao := dao.NewUserDao(l.svcCtx.Repository.Model)
	user, err := userDao.FindByUserName(l.ctx, req.UserName)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		l.Logger.Debug("查询失败", err.Error())
		return nil, errors.New(consts.ErrUnKnow)
	}

	// 用户不已存在
	if user == nil {
		return nil, errors.New(consts.ErrUserLogin)
	}

	// 比较密码
	if err = utils.ComparePassword(user.Password, req.Password); err != nil {
		return nil, errors.New(consts.ErrUserLogin)
	}

	// 颁发token
	acToken, reToken, err := l.CreateJwtToken(global.BaseClaim{
		ID:       user.ID,
		UserName: user.Name,
	})
	if err != nil {
		l.Logger.Debug("创建token失败", err.Error())
		return nil, errors.New(consts.ErrCreatToken)
	}

	// 登陆成功后就将验证码删除
	if err = l.DeleteCaptchaCache(key); err != nil {
		l.Logger.Info("删除redis失败", err.Error())
		return nil, errors.New(consts.ErrUnKnow)
	}

	return &types.UserLoginResp{
		AccessToken:  acToken,
		RefreshToken: reToken,
	}, nil
}

func (l *UserLoginLogic) CreateJwtToken(baseClaim global.BaseClaim) (accessToken, refreshToken string, err error) {
	unixTime := time.Now().Unix()
	accessExpiresTime := unixTime + l.svcCtx.Config.Auth.AccessExpire
	refreshExpiresTime := unixTime + l.svcCtx.Config.Auth.RefreshExpire
	claims := global.CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),                      // 签发时间
			ExpiresAt: jwt.NewNumericDate(time.Unix(accessExpiresTime, 0)), // 过期时间
			Issuer:    "oj",
		},
		BaseClaim: baseClaim,
	}
	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(l.svcCtx.Config.Auth.AccessSecret))
	if err != nil {
		return "", "", err
	}
	claims.ExpiresAt = jwt.NewNumericDate(time.Unix(refreshExpiresTime, 0))
	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(l.svcCtx.Config.Auth.AccessSecret))
	return
}

func (l *UserLoginLogic) DeleteCaptchaCache(key string) error {
	_, err := l.svcCtx.Repository.Redis.DelCtx(l.ctx, key)
	if err != nil {
		return err
	}
	return nil
}

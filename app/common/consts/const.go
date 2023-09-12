package consts

const (
	MinUserNameLength = 0
	MaxUserNameLength = 12

	MinPasswordLength = 6
	MaxPasswordLength = 12
)

const (
	GetCaptchaCachePrefix                 = "cache:oj:GetCaptcha:"          // 当前ip对应的验证码缓存
	GetCurrentIPCaptchaMaxTimeCachePrefix = "cache:oj:IPGetCaptchaMaxTime:" // 当前ip请求验证码接口一天的最大值缓存
)

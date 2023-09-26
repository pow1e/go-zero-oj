package consts

const (
	MinUserNameLength = 0
	MaxUserNameLength = 12

	MinPasswordLength = 6
	MaxPasswordLength = 12
)

// 用户缓存
const (
	GetCaptchaCachePrefix                 = "cache:oj:GetCaptcha:"          // 当前ip对应的验证码缓存
	GetCurrentIPCaptchaMaxTimeCachePrefix = "cache:oj:IPGetCaptchaMaxTime:" // 当前ip请求验证码接口一天的最大值缓存
)

// 用户
const (
	Goland = iota
	Java
	Rust
	Php
	Cpp
	C
)

const (
	AdminID             = 1
	AuthorizationHeader = "Authorization"
	UserInfo            = "userInfo"
)

var AdminRouteMap = map[string]struct{}{
	"publish-problem": {},
	"update-problem":  {},
	"update-category": {},
	"delete-problem":  {},
	// 继续添加需要的键值对
}

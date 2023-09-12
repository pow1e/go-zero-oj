package consts

const (
	ErrSearch                  = "查询失败"
	ErrUnKnow                  = "内部错误，请联系工作人员！"
	ErrConfirmPassword         = "两次密码不一致，请重试"
	ErrPasswordLength          = "密码长度不合法"
	ErrUserNameLength          = "用户名长度不合法"
	ErrCreatToken              = "颁发token失败"
	ErrCreateCaptcha           = "生成验证码失败"
	ErrOverTakeCaptchaMaxTimes = "超过验证码请求的最大次数，请稍后重试"
)

// user常量
const (
	ErrUserExist    = "当前用户已存在"
	ErrUserNotExist = "当前用户不已存在"
	ErrUserLogin    = "用户名或密码错误，请重试"
	ErrCaptcha      = "验证码错误，请重试"
)

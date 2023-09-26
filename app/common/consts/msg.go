package consts

const (
	ErrSearch                  = "查询失败"
	ErrUnKnow                  = "内部错误，请联系工作人员！"
	ErrConfirmPassword         = "两次密码不一致，请重试"
	ErrPasswordLength          = "密码长度不合法"
	ErrUserNameLength          = "用户名长度不合法"
	ErrOverTakeCaptchaMaxTimes = "超过验证码请求的最大次数，请稍后重试"
)

// user常量
const (
	ErrUserExist    = "当前用户已存在"
	ErrUserNotExist = "当前用户不已存在"
	ErrUserLogin    = "用户名或密码错误，请重试"
	ErrCaptcha      = "验证码错误，请重试"
)

const (
	ErrPhoneFormat = "电话格式错误，请重试"
	ErrEmailFormat = "邮箱格式错误，请重试"
)

// token以及校验
const (
	ErrCreatToken    = "颁发token失败"
	ErrCreateCaptcha = "生成验证码失败"
	ErrAuthorization = "用户认证失败,请重试"
	ErrPermissions   = "权限不足，请重试"
	ErrTokenExpired  = "当前请求token已过期，请重试"
	ErrTokenInvalid  = "当前请求token不合法，请重试"
)

// 问题创建
const (
	ErrExistProblemTitle = "该标题已存在，请重试"
	ErrCreateProblem     = "创建问题失败，请重试"
	ErrNotExistCategory  = "当前种类不存在，请重试"
	ErrNotExistProblem   = "当前问题不存在，请重试"
)

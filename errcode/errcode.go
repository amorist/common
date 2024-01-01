package errcode

var (
	// 服务级错误码
	ErrServer    = NewError(10001, "服务异常，请联系管理员")
	ErrParam     = NewError(10002, "参数错误")
	ErrSignParam = NewError(10003, "签名参数错误")

	// 模块级错误码 - 账号模块
	ErrAccountPhone         = NewError(20101, "用户手机号不合法")
	ErrAccountDuplicate     = NewError(20102, "账号已存在")
	ErrAccountPhoneParam    = NewError(20103, "参数错误：手机号码不能为空")
	ErrAccountPasswordParam = NewError(20104, "参数错误：密码不能为空")
	ErrAccountSmsParam      = NewError(20105, "参数错误：短信验证码不能为空")
	ErrAccountCaptcha       = NewError(20106, "用户验证码有误")
	ErrAccountPassword      = NewError(20107, "密码错误")
)

// Error .
type Error struct {
	Code int
	Msg  string
}

// NewError .
func NewError(code int, msg string) error {
	return &Error{
		Code: code,
		Msg:  msg,
	}
}

func (e Error) Error() string {
	return e.Msg
}

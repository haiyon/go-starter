package ecode

// All pkg ecode. ref: bilibili
var (
	OK = 0 // 成功

	AppKeyInvalid           = -1   // 应用程序不存在或已被封禁
	AccessKeyErr            = -2   // Access Key错误
	SignCheckErr            = -3   // API校验密匙错误
	MethodNoPermission      = -4   // 调用方对该Method没有权限
	NoLogin                 = -101 // 账号未登录
	UserDisabled            = -102 // 账号被封停
	LackOfScores            = -103 // 积分不足
	LackOfCoins             = -104 // 硬币不足
	CaptchaErr              = -105 // 验证码错误
	UserInactive            = -106 // 账号未激活
	UserNoMember            = -107 // 账号非正式会员或在适应期
	AppDenied               = -108 // 应用不存在或者被封禁
	MobileNoVerfiy          = -110 // 未绑定手机
	CsrfNotMatchErr         = -111 // csrf 校验失败
	ServiceUpdate           = -112 // 系统升级中
	UserIDCheckInvalid      = -113 // 账号尚未实名认证
	UserIDCheckInvalidPhone = -114 // 请先绑定手机
	UserIDCheckInvalidCard  = -115 // 请先完成实名认证

	NotModified           = -304 // 木有改动
	TemporaryRedirect     = -307 // 撞车跳转
	RequestErr            = -400 // 请求错误
	Unauthorized          = -401 // 未认证
	AccessDenied          = -403 // 访问权限不足
	NothingFound          = -404 // 啥都木有
	MethodNotAllowed      = -405 // 不支持该方法
	Conflict              = -409 // 冲突
	ServerErr             = -500 // 服务器错误
	ServiceUnavailable    = -503 // 过载保护,服务暂不可用
	Deadline              = -504 // 服务调用超时
	LimitExceed           = -509 // 超出限制
	FileNotExists         = -616 // 上传文件不存在
	FileTooLarge          = -617 // 上传文件太大
	FailedTooManyTimes    = -625 // 登录失败次数太多
	UserNotExist          = -626 // 用户不存在
	PasswordTooLeak       = -628 // 密码太弱
	UsernameOrPasswordErr = -629 // 用户名或密码错误
	TargetNumberLimit     = -632 // 操作对象数量限制
	TargetBlocked         = -643 // 被锁定
	UserLevelLow          = -650 // 用户等级太低
	UserDuplicate         = -652 // 重复的用户
	AccessTokenExpires    = -658 // Token 过期
	PasswordHashExpires   = -662 // 密码时间戳过期
	AreaLimit             = -688 // 地理区域限制
	CopyrightLimit        = -689 // 版权限制
	FailToAddMoral        = -701 // 扣节操失败

	Degrade     = -1200 // 被降级过滤的请求
	RPCNoClient = -1201 // rpc服务的client都不可用
	RPCNoAuth   = -1202 // rpc服务的client没有授权
)

var ecodeText = map[int]string{
	OK:           "成功",
	NothingFound: "啥都没有",
	RequestErr:   "请求错误",
}

// Text returns text
func Text(code int) string {
	return ecodeText[code]
}

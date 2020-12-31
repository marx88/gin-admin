package constants

const (
	// UTF8 字符集
	UTF8 = "UTF-8"

	// GBK 字符集
	GBK = "GBK"

	// HTTP http请求
	HTTP = "http://"

	// HTTPS https请求
	HTTPS = "https://"

	// Success 通用成功标识
	Success = "0"

	// Fail 通用失败标识
	Fail = "1"

	// LoginSuccess 登录成功
	LoginSuccess = "Success"

	// Logout 注销
	Logout = "Logout"

	// LoginFail 登录失败
	LoginFail = "Error"

	// CaptchaCodeKey 验证码 redis key
	CaptchaCodeKey = "captcha_codes:"

	// LoginTokenKey 登录用户 redis key
	LoginTokenKey = "login_tokens:"

	// RepeatSubmitKey 防重提交 redis key
	RepeatSubmitKey = "repeat_submit:"

	// CaptchaExpiration 验证码有效期（分钟）
	CaptchaExpiration = 2

	// Token 令牌
	Token = "token"

	// TokenPrefix 令牌前缀
	TokenPrefix = "Bearer "

	// LoginUserKey 令牌前缀
	LoginUserKey = "login_user_key"

	// JwtUserID 用户ID
	JwtUserID = "userid"

	// JwtUserName 用户名称
	JwtUserName = "sub"

	// JwtAvatar 用户头像
	JwtAvatar = "avatar"

	// JwtCreated 创建时间
	JwtCreated = "created"

	// JwtAuthorities 用户权限
	JwtAuthorities = "authorities"

	// SysConfigKey 参数管理 cache key
	SysConfigKey = "sys_config:"

	// SysDictKey 字典管理 cache key
	SysDictKey = "sys_dict:"

	// ResourcePrefix 资源映射路径 前缀
	ResourcePrefix = "/profile"
)

// @Description https://open.weibo.com/wiki/Error_code 参照微博的错误码
// @Author: toom1996 <1023150697@qq.com>
// @dateTime: 2021/3/1 15:25
package util

// 1 服务错误
// 2 系统错误
const (

	// 请求成功
	HttpCodeSuccess = 0

	// ---------- 验证器 ----------

	// 验证器错误
	ValidatorErrVerifyFailed = 11000

	// ---------- 用户 ----------

	// 用户token 过期
	UserErrTokenExpired = 12000

	Z_ERROR = 500

	ERROR_EXIST_TAG         = 10001
	ERROR_NOT_EXIST_TAG     = 10002
	ERROR_NOT_EXIST_ARTICLE = 10003

	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004
)

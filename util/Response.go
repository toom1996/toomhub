// @Description
// @Author: toom1996 <1023150697@qq.com>
// @dateTime: 2020/12/15 9:21
package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

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

	// token不合法
	UserErrBadToken = 12001

	// 获取用户信息失败
	UserErrGetInfoError = 12002

	Z_ERROR = 500

	ERROR_EXIST_TAG         = 10001
	ERROR_NOT_EXIST_TAG     = 10002
	ERROR_NOT_EXIST_ARTICLE = 10003

	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004
)

type ResponseData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (r *ResponseData) Error() string {
	return fmt.Sprintf("status code %d: %s", r.Code, r.Msg)
}

// response OK status code
func ResponseOk(context *gin.Context, message string, data interface{}) {
	BaseResponse(context, http.StatusOK, HttpCodeSuccess, message, data)
}

// response 400 status code
func ResponseError(context *gin.Context, err interface{}) {
	var errStr string
	var errCode int

	switch err.(type) {
	case validator.ValidationErrors:
		errStr = Translate(err.(validator.ValidationErrors))
		errCode = ValidatorErrVerifyFailed
	case *ResponseData:
		r := err.(*ResponseData)
		errStr = r.Msg
		errCode = r.Code
	}
	BaseResponse(context, http.StatusOK, errCode, errStr, "")
}

// base response
func BaseResponse(context *gin.Context, httpCode, errCode int, message string, data interface{}) {
	context.JSON(httpCode, ResponseData{
		Code: errCode,
		Msg:  message,
		Data: data,
	})
	return
}

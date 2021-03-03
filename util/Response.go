// @Description
// @Author: toom1996 <1023150697@qq.com>
// @dateTime: 2020/12/15 9:21
package util

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// response OK status code
func ResponseOk(context *gin.Context, message string, data interface{}) {
	BaseResponse(context, http.StatusOK, HttpCodeSuccess, message, data)
}

// response 400 status code
func ResponseError(context *gin.Context, err interface{}) {
	var errStr string
	switch err.(type) {
	case validator.ValidationErrors:
		errStr = Translate(err.(validator.ValidationErrors))
	case error:
		errStr = err.(error).Error()
	default:
		errStr = err.(string)

	}
	BaseResponse(context, http.StatusOK, ValidatorErrVerifyFailed, errStr, "")
}

// response 400 status code
func ResponseValidatorErr(context *gin.Context, err interface{}) {
	var errStr string
	errStr = Translate(err.(validator.ValidationErrors))
	BaseResponse(context, http.StatusOK, ValidatorErrVerifyFailed, errStr, "")
}

// base response
func BaseResponse(context *gin.Context, httpCode, errCode int, message string, data interface{}) {
	context.JSON(httpCode, Response{
		Code: errCode,
		Msg:  message,
		Data: data,
	})
	return
}

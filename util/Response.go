// @Description
// @Author: toom1996 <1023150697@qq.com>
// @dateTime: 2020/12/15 9:21
package util

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"toomhub/pkg"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// response 200 status code
func ResponseOk(context *gin.Context, message string, data interface{}) {
	BaseResponse(context, http.StatusOK, 200, message, data)
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
	BaseResponse(context, http.StatusOK, pkg.Z_BAD_REQUEST, errStr, "")
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

package v1

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
	rules "toomhub/rules/user/v1"
	service "toomhub/service/user/v1"
	"toomhub/util"
)

// @summary 用户注册接口
// @title Swagger Example API
// @tags  用户类接口
// @description  用户注册接口
// @produce  json
// @param mobile body string true "123456789"
// @param code body	int true "短信验证码"
// @success 200 {string} json "{"code":200,"data":"data","msg":"ok"}"
// @router /api/v1/user/register [post]
func Register(context *gin.Context) {
	var formValidator rules.V1UserRegister
	err := context.ShouldBind(&formValidator)
	if err != nil {
		util.ResponseError(context, err)
		return
	}

	info, err := service.V1UserRegister(&formValidator)
	if err != nil {
		util.ResponseError(context, err)
		return
	}

	util.ResponseOk(context, "注册成功", info)
}

func SmsSend(context *gin.Context) {
	var formValidator rules.V1UserSmsSend
	err := context.ShouldBind(&formValidator)
	if err != nil {
		util.ResponseError(context, err)
		return
	}

	//存redis
	rand.Seed(time.Now().UnixNano())
	_, err = util.SendRegisterSms(formValidator.Mobile, rand.Intn(999999))
	if err != nil {
		util.ResponseError(context, "验证码发送失败")
		return
	}
}

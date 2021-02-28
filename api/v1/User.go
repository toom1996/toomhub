package v1

import (
	"github.com/gin-gonic/gin"
	"toomhub/logic"
	rules "toomhub/rules/user/v1"
	"toomhub/util"
	//"math/rand"
	//"time"
	//rules "toomhub/rules/user/v1"
	//service "toomhub/service/user/v1"
	//"toomhub/util"
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

	formLogic := logic.UserLogic{}

	r, err := formLogic.Register(&formValidator)

	if err != nil {
		util.ResponseError(context, err)
		return
	}

	util.ResponseOk(context, "登陆成功", r)
}

// @summary 发送短信验证码接口
// @title Swagger Example API
// @tags  用户类接口
// @description  发送短信验证码接口
// @produce  json
// @param mobile body string true "123456789"
// @success 200 {string} json "{"code":200,"data":"data","msg":"ok"}"
// @router /api/v1/user/register [post]
func SmsSend(context *gin.Context) {
	var formValidator rules.V1UserSmsSend
	err := context.ShouldBind(&formValidator)
	if err != nil {
		util.ResponseError(context, err)
		return
	}

	formLogic := logic.UserLogic{}

	_, err = formLogic.SmsSend(&formValidator)
	if err != nil {
		util.ResponseError(context, err)
		return
	}
	util.ResponseOk(context, "验证码发送成功", "")
}

//github OAuth登陆
func GithubOAuth(context *gin.Context) {
	var formValidator rules.V1UserGithubOAuth
	err := context.ShouldBind(&formValidator)
	if err != nil {
		util.ResponseError(context, err)
		return
	}

	formLogic := logic.UserLogic{}

	info, err := formLogic.GithubOAuthLogic(&formValidator)

	if err != nil {
		util.ResponseError(context, err)
		return
	}
	util.ResponseOk(context, "OK", info)
}

//github OAuth登陆
func Login(context *gin.Context) {
	var formValidator rules.V1UserGithubOAuth
	err := context.ShouldBind(&formValidator)
	if err != nil {
		util.ResponseError(context, err)
		return
	}

	formLogic := logic.UserLogic{}

	info, err := formLogic.GithubOAuthLogic(&formValidator)

	if err != nil {
		util.ResponseError(context, err)
		return
	}
	util.ResponseOk(context, "OK", info)
}

//刷新token
func RefreshToken(context *gin.Context) {
	var formValidator rules.V1UserRefreshToken
	err := context.ShouldBind(&formValidator)
	if err != nil {
		util.ResponseError(context, err)
		return
	}

	formLogic := logic.UserLogic{}
	r, err := formLogic.RefreshToken(&formValidator, context)
	if err != nil {
		util.ResponseError(context, err)
		return
	}

	util.ResponseOk(context, "OK", r)
}

package v1

import (
	"github.com/gin-gonic/gin"
	rules "toomhub/rules/user/v1"
	service "toomhub/service/user/v1"
	"toomhub/util"
)

// @summary 用户注册接口
// @title Swagger Example API
// @tags  用户类接口
// @description  用户注册接口
// @produce  json
// @param mobile body string true "13502127317"
// @param code body	ReleaseTemplateAdd true "JSON数据"
// @success 200 {string} json "{"code":200,"data":"name","msg":"ok"}"
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

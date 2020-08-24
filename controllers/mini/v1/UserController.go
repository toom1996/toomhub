// @Description
// @Author    2020/8/19 15:59
package ControllersMiniV1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	LogicMiniV1 "toomhub/logic/mini/v1"
	ModelMiniV1 "toomhub/model/mini/v1"
)

type UserController struct {
	Name string `validate:"required"`
	Age  int    `validate:"gte=0,lte=100"`
}

type test struct {
	aa string
}

//当前控制器注册的路由
func (miniV1User *UserController) Register(engine *gin.Engine) {
	user := engine.Group("/v1/mini")
	{
		//小程序用户登陆接口
		user.POST("/login", miniV1User.LoginByV1)
	}
}

// @url 	localhost:8080/mini/login	POST
// @title    小程序用户登陆接口
// @description   初次登陆的用户将会入库, 非初次登陆的用户将会返回用户信息
// @auth	toom <1023150697@qq.com>
// @param     Context	*gin.Context
// @return
func (miniV1User *UserController) LoginByV1(Context *gin.Context) {
	//validator验证
	model := ModelMiniV1.LoginByV1Model{}
	err := Context.ShouldBind(&model)
	if err != nil {
		Context.String(http.StatusBadRequest, "参数错误:%s", err.Error())
		return
	}

	//逻辑验证
	logic := LogicMiniV1.UserLogic{}
	query, err := logic.Login(&model)

	if err != nil {
		Context.JSON(http.StatusOK, gin.H{
			"code":    400,
			"message": err.Error(),
			"data":    "",
		})
		return
	}

	Context.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "登陆成功",
		"data":    query,
	})
}
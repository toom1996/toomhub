// @Description
// @Author    2020/8/19 15:59
package ControllersMiniV1

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	LogicMiniV1 "toomhub/logic/mini/v1"
	validator2 "toomhub/validator"
	validatorMiniprogramV1 "toomhub/validator/miniprogram/v1"
)

type UserController struct {
	Name string `validate:"required"`
	Age  int    `validate:"gte=0,lte=100"`
}

type test struct {
	aa string
}

//当前控制器注册的路由
func (u *UserController) Register(engine *gin.Engine) {
	user := engine.Group("/v1/mini/user")
	{
		//小程序用户登陆接口
		user.POST("/login", u.Login)
		user.POST("/token-checker", u.tokenChecker)
	}
}

// @url 	localhost:8080/mini/login	POST
// @title    小程序用户登陆接口
// @description   初次登陆的用户将会入库并返回信息, 非初次登陆的用户将会返回用户信息
// @auth	toom <1023150697@qq.com>
// @param     Context	*gin.Context
// @return
func (u *UserController) Login(Context *gin.Context) {
	//validator验证
	validator := validatorMiniprogramV1.Login{}
	err := Context.ShouldBind(&validator)
	if err != nil {
		Context.String(http.StatusOK, "参数错误:%s", err.Error())
		return
	}

	//逻辑验证
	logic := LogicMiniV1.UserLogic{}
	query, err := logic.Login(&validator)

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

// @title	检查token接口
func (u *UserController) tokenChecker(Context *gin.Context) {
	var commonValidator validator2.CommonValidator

	formValidator := validatorMiniprogramV1.Refresh{}
	err := Context.ShouldBind(&formValidator)
	if err != nil {
		Context.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  commonValidator.TransError(err.(validator.ValidationErrors)),
		})
		return
	}

	formLogic := LogicMiniV1.UserLogic{}
	token, err := formLogic.Check(&formValidator, Context)
	if err != nil {
		Context.JSON(http.StatusOK, gin.H{
			"code": 401,
			"msg":  err.Error(),
		})
		return
	}
	Context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  token,
	})
}

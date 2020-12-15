package controllers

import (
	"github.com/gin-gonic/gin"
	"toomhub/middleware"
)

type RegisterController struct {
}

//当前控制器注册的路由
func (u *RegisterController) Register(engine *gin.Engine) {
	user := engine.Group("/v1/user")
	{
		//小程序用户登陆接口
		user.GET("/register", u.actionRegister)
	}
	user.Use(middleware.CheckIdentity())
	{
		//user.GET("/get-info", u.refreshInfo)
	}
}

func (*RegisterController) actionRegister(context *gin.Context) {
	context.JSON(200, gin.H{
		"code":    400,
		"message": "can't find code",
	})
	return
}

package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	v1 "toom/logic/miniprogram/v1"
	model2 "toom/model"
)

type UserController struct {
	Name string `validate:"required"`
	Age  int    `validate:"gte=0,lte=100"`
}

func (miniV1User *UserController) Router(engine *gin.Engine) {

	user := engine.Group("/v1/mini")
	{
		user.POST("/login", miniV1User.LoginByV1)
	}

}

// @url 	localhost:8080/mini/login
// @title    登录凭证校验。通过 wx.login 接口获得临时登录凭证 code 后传到开发者服务器调用此接口完成登录流程。
// @description   通过小程序传过来的 code 获取用户的 appid
// @auth	toom <1023150697@qq.com>
// @param     Context	*gin.Context
// @return
func (miniV1User *UserController) LoginByV1(Context *gin.Context) {
	model := model2.LoginByV1Model{}
	err := Context.ShouldBind(&model)
	if err != nil {
		Context.String(http.StatusBadRequest, "参数错误:%s", err.Error())
		return
	}

	logic := v1.UserLogic{}
	_, err = logic.Login(&model)
	if err != nil {
		Context.JSON(http.StatusOK, map[string]interface{}{
			"code":    400,
			"message": err.Error(),
			"data":    "",
		})
		return
	}

	Context.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"message": "登陆成功",
		"data":    "",
	})
}

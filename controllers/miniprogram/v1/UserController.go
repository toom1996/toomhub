package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	v1 "toom/logic/miniprogram/v1"
	validatorMiniprogramV1 "toom/validator/miniprogram/v1"
)

type UserController struct {
	Name string `validate:"required"`
	Age  int    `validate:"gte=0,lte=100"`
}

func (miniV1User *UserController) Router(engine *gin.Engine) {

	user := engine.Group("/mini")
	{
		user.POST("/login", miniV1User.Login)
	}

}

// @url 	localhost:8080/mini/login
// @title    登录凭证校验。通过 wx.login 接口获得临时登录凭证 code 后传到开发者服务器调用此接口完成登录流程。
// @description   通过小程序传过来的 code 获取用户的 appid
// @auth	toom <1023150697@qq.com>
// @param     Context	*gin.Context	gin上下文
// @return
func (miniV1User *UserController) Login(Context *gin.Context) {

	validate := validator.New()
	user := validatorMiniprogramV1.UserValidator{
		Code: Context.PostForm("code"),
		Tode: Context.PostForm("code"),
	}
	//验证器
	validatorError := validate.Struct(user)
	if validatorError != nil {
		Context.JSON(200, map[string]interface{}{
			"code": 400,
			"msg":  validatorError.Error(),
			"data": "",
		})
		return
	}

	logic := v1.UserLogic{}
	logic.Login()
	fmt.Println("success")
}

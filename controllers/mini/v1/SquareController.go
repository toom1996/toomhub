// @Description
// @Author    2020/8/26 9:18
package ControllersMiniV1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	LogicMiniV1 "toomhub/logic/mini/v1"
	validatorMiniprogramV1 "toomhub/validator/miniprogram/v1"
)

type SquareController struct {
}

//当前控制器注册的路由
func (square *SquareController) Register(engine *gin.Engine) {
	user := engine.Group("/v1/mini/sq")
	{
		//小程序用户登陆接口
		user.GET("/index", square.index)
	}
}

func (square *SquareController) index(Context *gin.Context) {
	//验证器
	validator := validatorMiniprogramV1.SquareIndex{}
	err := Context.BindQuery(&validator)
	if err != nil {
		Context.String(http.StatusBadRequest, "参数错误:%s", err.Error())
		return
	}

	//logic
	logic := LogicMiniV1.SquareLogic{}
	query, _ := logic.SquareIndex(&validator)

	Context.JSON(200, gin.H{
		"message": "OK",
		"code":    200,
		"data": map[string]interface{}{
			"list": query,
		},
	})
}

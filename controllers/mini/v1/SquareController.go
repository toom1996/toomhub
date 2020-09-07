// @Description
// @Author    2020/8/26 9:18
package ControllersMiniV1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	LogicMiniV1 "toomhub/logic/mini/v1"
	v1MiniMiddleware "toomhub/middware/mini/v1"
	validatorMiniprogramV1 "toomhub/validator/miniprogram/v1"
)

type SquareController struct {
}

//当前控制器注册的路由
func (square *SquareController) Register(engine *gin.Engine) {
	user := engine.Group("/v1/mini/sq")
	// 广场首页接口
	user.GET("/index", square.index)

	user.Use(v1MiniMiddleware.CheckIdentity())
	{
		// 发布一条广场信息
		user.POST("/create", square.create)
	}
}

func (square *SquareController) index(Context *gin.Context) {
	//验证器
	formValidator := validatorMiniprogramV1.SquareIndex{}
	err := Context.BindQuery(&formValidator)
	if err != nil {
		Context.String(http.StatusBadRequest, "参数错误:%s", err.Error())
		return
	}

	//logic
	logic := LogicMiniV1.SquareLogic{}
	query, _ := logic.SquareIndex(&formValidator)

	Context.JSON(200, gin.H{
		"message": "OK",
		"code":    200,
		"data": map[string]interface{}{
			"list": query,
		},
	})
}

// @title	创建一条广场消息
func (square *SquareController) create(Context *gin.Context) {
	//var commonValidator validator2.CommonValidator
	//验证器
	formValidator := validatorMiniprogramV1.SquareCreate{}
	err := Context.ShouldBind(&formValidator)

	if err != nil {
		Context.JSON(http.StatusOK, gin.H{
			"code": 400,
			//"message": commonValidator.TransError(err.(validator.ValidationErrors)),
			"message": err.Error(),
			"data":    "",
		})
		return
	}

	//logic
	logic := LogicMiniV1.SquareLogic{}
	_, err = logic.SquareCreate(&formValidator)

	if err != nil {
		Context.JSON(200, gin.H{
			"message": err,
			"code":    400,
		})
	}
	Context.JSON(200, gin.H{
		"message": "OK",
		"code":    200,
	})
}

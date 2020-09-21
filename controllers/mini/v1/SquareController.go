// @Description
// @Author    2020/8/26 9:18
package ControllersMiniV1

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
	LogicMiniV1 "toomhub/logic/mini/v1"
	v1MiniMiddleware "toomhub/middware/mini/v1"
	"toomhub/util"
	validatorMiniprogramV1 "toomhub/validator/miniprogram/v1"
)

type SquareController struct {
}

//当前控制器注册的路由
func (square *SquareController) Register(engine *gin.Engine) {

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("countValidate", validatorMiniprogramV1.CountValidate)
	}

	user := engine.Group("/v1/mini/sq")
	// 广场首页接口
	user.GET("/index", square.index)
	user.Use(v1MiniMiddleware.CheckIdentity())
	{
		// 发布一条广场信息
		user.POST("/create", square.create)
		//标签搜索
		user.GET("/tag-search", square.TagSearch)
		//广场消息点赞
		user.POST("/like", square.Like)
	}
}

// @title	广场首页
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

// @title	标签搜索
func (square *SquareController) TagSearch(Context *gin.Context) {

	util.EsGet("toomhub", "111")
}

// @title	点赞
func (square *SquareController) Like(Context *gin.Context) {

	formValidator := validatorMiniprogramV1.LikeValidator{}
	err := Context.ShouldBind(&formValidator)

	if err != nil {
		Context.JSON(http.StatusOK, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	logic := LogicMiniV1.SquareLogic{}
	_, err = logic.SquareLike(&formValidator)

	if err != nil {
		Context.JSON(http.StatusOK, gin.H{
			"code":    400,
			"message": "点赞失败",
		})
		return
	}

	Context.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "点赞成功",
	})

	return

}

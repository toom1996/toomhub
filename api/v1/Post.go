// @Description
// @Author: toom1996 <1023150697@qq.com>
// @dateTime: 2020/12/22 14:51
package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"toomhub/logic"
	rules "toomhub/rules/user/v1"
	"toomhub/util"
)

type PostLogic struct {
}

// 发布
func PublishPost(context *gin.Context) {
	var formValidator rules.V1PostPublishPost
	err := context.ShouldBind(&formValidator)
	if err != nil {
		util.ResponseError(context, err)
		return
	}
	fmt.Println("000000000000000")
	fmt.Println(context.Get("identity"))

	context.Get("identity")
	formLogic := logic.PostLogic{}
	formLogic.PublishPost(&formValidator, context)

	util.ResponseOk(context, "发布成功", "")
}

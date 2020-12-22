// @Description
// @Author: toom1996 <1023150697@qq.com>
// @dateTime: 2020/12/22 14:51
package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	rules "toomhub/rules/user/v1"
	"toomhub/util"
)

func PublishPost(context *gin.Context) {
	var formValidator rules.V1PostPublishPost
	err := context.ShouldBind(&formValidator)
	fmt.Println(err)
	fmt.Println(formValidator)
	util.ResponseOk(context, "发布成功", "")
}

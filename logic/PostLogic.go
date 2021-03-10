// @Description 发布消息逻辑层
// @Author: toom1996 <1023150697@qq.com>
// @dateTime: 2021/3/10 15:49
package logic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	rules "toomhub/rules/user/v1"
	"toomhub/util"
)

type PostLogic struct {
}

func (l *PostLogic) PublishPost(v *rules.V1PostPublishPost, c *gin.Context) {
	// TODO 判断用户状态
	fmt.Println(util.GetIdentity(c))
	fmt.Println(v)
}

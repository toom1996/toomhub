// @Description
// @Author    2020/9/2 15:33
package v1MiniMiddleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"toomhub/util"
)

// @title 	验证token
func CheckIdentity() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.GetHeader("Toomhub-Token")

		r, _ := util.ParseToken(token, c)

		fmt.Println(r)
		//c.JSON(200, map[string]string{
		//	"sdf": "000",
		//})
		//c.Abort()
	}
}

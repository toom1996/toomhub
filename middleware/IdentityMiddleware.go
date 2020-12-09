// @Description
// @Author    2020/9/2 15:33
package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
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

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

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
		token := c.GetHeader("Authorization")

		claims, err := util.ParseToken(token, c)
		c.Set("identity", claims)
		fmt.Println("18 ------> ideneity")
		fmt.Println(c.Get("identity"))

		if err != nil {
			c.JSON(util.Success, err.(*util.ResponseData))
			c.Abort()
		}
	}
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		c.Header("Access-Control-Expose-Headers", "Authorization, Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

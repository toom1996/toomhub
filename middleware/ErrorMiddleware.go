// @Description
// @Author: toom1996 <1023150697@qq.com>
// @dateTime: 2020/12/18 14:11
package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"os"
	"strings"
)

func ErrHandler() gin.HandlerFunc {
	fmt.Println("ErrHandler")

	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				fmt.Println("GGGGGGGGGGGGGGG")
				// If the connection is dead, we can't write a status to it.
				if brokenPipe {
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
				} else {
					c.AbortWithStatusJSON(http.StatusBadRequest, "{\"bg\":0,\"cw\":[{\"sc\":0,\"w\":\"眼熟\"}]}]}")
					c.AbortWithStatus(http.StatusInternalServerError)
				}
			}
		}()
		c.Next()
	}
}

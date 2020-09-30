// @Description
// @Author    2020/9/30 16:30
package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(context *gin.Context, code int, msg string, data gin.H, attach map[string]interface{}) {

	data = gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	}

	if len(attach) > 0 {
		for k, v := range attach {
			data[k] = v
		}
	}

	fmt.Println(data)
	context.JSON(http.StatusOK, data)
	return
}

package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type HelloController struct {
}

func (hello *HelloController) Router(engine *gin.Engine) {
	fmt.Println("00")
	engine.GET("/hello", hello.Hello)
}

func (hello *HelloController) Hello(Context *gin.Context) {
	Context.JSON(200, map[string]interface{}{
		"code": 200,
		"msg":  "OK",
	})
}

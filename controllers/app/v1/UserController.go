package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (hello *UserController) Router(engine *gin.Engine) {
	fmt.Println("00")
	engine.GET("/hello", hello.Hello)
}

func (hello *UserController) Hello(Context *gin.Context) {
	Context.JSON(200, map[string]interface{}{
		"code": 200,
		"msg":  "OK",
	})
}

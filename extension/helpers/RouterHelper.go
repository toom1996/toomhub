// @Description
// @Author    2020/8/19 15:27
package helpers

import (
	"github.com/gin-gonic/gin"
	ControllersMiniV1 "toomhub/controllers/mini/v1"
)

func RegisterRoutes(app *gin.Engine) {
	//v1小程序控制器
	new(ControllersMiniV1.UserController).Register(app)
}

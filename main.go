package main

import (
	"github.com/gin-gonic/gin"
	ControllersMiniV1 "toomhub/controllers/mini/v1"
	"toomhub/tool"
)

func main() {
	cfg, err := tool.Init("./config/app.json")

	if err != nil {
		panic(err.Error())
	}

	app := gin.Default()

	registerRouter(app)

	_ = app.Run(cfg.AppHost + ":" + cfg.AppPort)
}

//路由设置
func registerRouter(router *gin.Engine) {
	new(ControllersMiniV1.UserController).Register(router)
}

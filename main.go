package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"toomhub/controllers/common"
	ControllersMiniV1 "toomhub/controllers/mini/v1"
	"toomhub/util"
)

func main() {
	cfg, err := util.Init("./config/app.json")

	if err != nil {
		fmt.Println("000000000")
		panic(err.Error())
	}

	//初始化redis
	util.RedisInit()

	//初始化mysql
	util.MysqlInit()

	fmt.Println(cfg.AppPort)
	app := gin.Default()

	registerRouter(app)

	_ = app.Run(cfg.AppHost + ":" + cfg.AppPort)

}

//路由设置
func registerRouter(router *gin.Engine) {
	new(ControllersMiniV1.UserController).Register(router)
	new(ControllersMiniV1.SquareController).Register(router)
	new(common.ImageController).Register(router)
}

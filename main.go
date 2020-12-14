package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"toomhub/controllers"
	"toomhub/util"
)

func main() {
	cfg, err := util.Init("./config/app.json")

	if err != nil {
		panic(err.Error())
	}

	util.EsInit()

	//初始化redis
	util.RedisInit()

	//初始化mysql
	util.MysqlInit()

	//初始化zaplog
	util.ZapLogInit()

	fmt.Println(cfg.AppPort)
	app := gin.Default()
	registerRouter(app)

	a := []string{"test", "hello", "world"}
	util.Debug("output", a)
	_ = app.Run(cfg.AppHost + ":" + cfg.AppPort)
}

//路由设置
func registerRouter(router *gin.Engine) {
	new(controllers.UserController).Register(router)
	new(controllers.SquareController).Register(router)
	new(controllers.ImageController).Register(router)
	new(controllers.VideoController).Register(router)
}

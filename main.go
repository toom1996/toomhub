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
		fmt.Println("000000000")
		panic(err.Error())
	}

	util.EsInit()

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
	new(controllers.UserController).Register(router)
	new(controllers.SquareController).Register(router)
	new(controllers.ImageController).Register(router)
	new(controllers.VideoController).Register(router)
}

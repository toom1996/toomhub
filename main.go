package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"time"
	"toomhub/controllers"
	_ "toomhub/docs"
	"toomhub/util"
	"toomhub/validatorRules"
)

type Note struct {
	Content string
	Cities  []string
}

type Person struct {
	Name string
	Age  int `ini:"age"`
	Male bool
	Born time.Time
	Note
	Created time.Time `ini:"-"`
}

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

	validatorRules.InitVali() // 字段验证

	fmt.Println(cfg.AppPort)
	app := gin.Default()
	registerRouter(app)
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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

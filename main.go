package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"toom/tool"
)

func init() {

}

func main() {
	//初始化配置
	config, _ := tool.Load("./config/app.json")

	app := gin.Default()
	//注册路由
	tool.RegisterRoutes(app)
	_ = app.Run(config.AppHost + ":" + config.AppPort)
}

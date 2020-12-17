package main

import (
	_ "toomhub/docs"
	"toomhub/router"
	"toomhub/rules"
	"toomhub/util"
)

func main() {

	//util.EsInit()

	//初始化redis
	//util.RedisInit()

	//初始化mysql
	util.MysqlInit()

	//初始化zaplog
	util.ZapLogInit()

	//注册中文验证器
	rules.InitVali()

	//注册路由
	router.InitRouter()

}

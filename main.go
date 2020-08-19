package main

import (
	"github.com/gin-gonic/gin"
	"toomhub/extension/helpers"
)

func main() {
	cfg, err := helpers.Init("./config/app.json")

	if err != nil {
		panic(err.Error())
	}

	app := gin.Default()

	helpers.RegisterRoutes(app)

	_ = app.Run(cfg.AppHost + ":" + cfg.AppPort)
}

package tool

import (
	"github.com/gin-gonic/gin"
	miniV1 "toom/controllers/mini/v1"
)

func RegisterRoutes(app *gin.Engine) *gin.Engine {
	new(miniV1.UserController).Router(app)
	return app
}

package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	v1 "toomhub/api/v1"
	"toomhub/setting"
)

func InitRouter() {
	r := gin.Default()

	r.Use(gin.Logger())

	//r.Use(middleware.ErrHandler())

	r.Use(gin.Recovery())

	gin.SetMode(setting.ZConfig.App.RunMode)
	//registerRouter(r)

	apiV1 := r.Group("/api/v1")
	{
		//用户注册接口
		apiV1.POST("/user/register", v1.Register)

		//用户注册短信发送接口
		apiV1.POST("/user/sms-send", v1.SmsSend)

		//上传图片
		apiV1.GET("/upload/get-qiniu-access-token", v1.GetQiniuAccessToken)

		//发布主题
		apiV1.POST("/post/publish-post", v1.PublishPost)

		apiV1.GET("/user/github-callback", v1.GithubOAuth)
	}

	//swagger文档生成接口
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	_ = r.Run(setting.ZConfig.Server.HttpHost + ":" + setting.ZConfig.Server.HttpPort)

}

//路由设置
func registerRouter(router *gin.Engine) {
	//new(controllers.UserController).Register(router)
	//new(controllers.SquareController).Register(router)
	//new(controllers.ImageController).Register(router)
	//new(controllers.VideoController).Register(router)
}

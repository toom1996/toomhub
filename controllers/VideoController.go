// @Description
// @Author    2020/10/27 9:19
package controllers

import (
	"github.com/gin-gonic/gin"
	v1MiniMiddleware "toomhub/middware/mini/v1"
	"toomhub/service"
	"toomhub/util"
)

type VideoController struct {
}

//当前控制器注册的路由
func (image *VideoController) Register(engine *gin.Engine) {

	controller := engine.Group("/video")
	{
		controller.POST("/upload", image.Upload)
	}

	controller.Use(v1MiniMiddleware.CheckIdentity())
	{
		//controller.POST("/upload", image.Upload)
	}
}

func (*VideoController) Upload(context *gin.Context) {
	file, header, err := context.Request.FormFile("file")
	if err != nil {
		context.JSON(200, gin.H{
			"code": 400,
			"msg":  "上传失败!",
		})
		return
	}

	uploader := service.QiniuUploader{}

	res, err := uploader.VideoUpload(file, header.Filename)
	if err != nil {
		context.JSON(200, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	var mm interface{}
	mm = res["url"]
	name := mm.(string)

	url := util.GetConfig().Qiniu.FileServer + name
	context.JSON(200, gin.H{
		"code":    200,
		"message": "上传成功",
		"data": map[string]interface{}{
			"url":          url,
			"size":         res["size"],
			"extension":    res["extension"],
			"request_host": util.GetConfig().Qiniu.FileServer,
			"name":         res["url"],
		},
	})
}

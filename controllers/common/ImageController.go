// @Description
// @Author    2020/8/27 15:27
package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"toomhub/service"
	"toomhub/util"
)

type ImageController struct {
}

//当前控制器注册的路由
func (image *ImageController) Register(engine *gin.Engine) {
	user := engine.Group("/c/image")
	{
		//小程序用户登陆接口
		user.POST("/upload", image.Upload)
	}
}

// @title	广场图片上传接口
// @desc	图片对接到七牛云
func (*ImageController) Upload(context *gin.Context) {
	param := "?imageMogr2/auto-orient/format/webp"

	file, header, err := context.Request.FormFile("file")
	if err != nil {
		context.JSON(200, gin.H{
			"code": 400,
			"msg":  "上传失败!",
		})
		return
	}

	uploader := service.QiniuUploader{}
	//
	res, err := uploader.Upload(file, header.Filename)
	fmt.Println(file)

	if err != nil {
		context.JSON(200, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	fmt.Println(res)
	var mm interface{}
	mm = res["url"]
	name := mm.(string)

	url := util.GetConfig().Qiniu.FileServer + name + param
	context.JSON(200, gin.H{
		"code":    200,
		"message": "上传成功",
		"data": map[string]interface{}{
			"url":          url,
			"size":         res["size"],
			"extension":    res["extension"],
			"request_host": util.GetConfig().Qiniu.FileServer,
			"param":        param,
			"name":         res["url"],
		},
	})
}

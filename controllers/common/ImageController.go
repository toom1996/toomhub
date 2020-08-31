// @Description
// @Author    2020/8/27 15:27
package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"toomhub/service"
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

func (*ImageController) Upload(context *gin.Context) {

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
	url, err := uploader.Upload(file, header.Filename)
	fmt.Println(url)
	fmt.Println(err)
	//if err != nil {
	//	res["message"] = err.Error()
	//	return
	//}
	//_ = context.SaveUploadedFile(file, filepath)
	context.JSON(200, gin.H{
		"code": 200,
		"msg":  "上传成功!",
		"data": url,
	})
}

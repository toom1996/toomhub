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

	//defer context.JSON(, res)

	file, header, err := context.Request.FormFile("file")
	if err != nil {
		context.JSON(200, gin.H{
			"code": 400,
			"msg":  "上传失败!",
		})
		return
	}

	fmt.Println(header.Filename)
	fmt.Println(file)
	//文件后缀
	//fileExt := strings.ToLower(path.Ext(file.Filename))
	//
	//if fileExt != ".png" && fileExt != ".jpg" && fileExt != ".gif" && fileExt != ".jpeg" {
	//	context.JSON(200, gin.H{
	//		"code": 400,
	//		"msg":  "上传失败!只允许png,jpg,gif,jpeg文件",
	//	})
	//	return
	//}

	//fileName := fmt.Sprintf("%s%s", file.Filename, time.Now().String())
	//filedDir := fmt.Sprintf("%s%d%s/", 's', time.Now().Year(), time.Now().Month().String())
	//isExist, _ := tools.IsFileExist(filedDir)
	//if !isExist {
	//	_ = os.Mkdir(filedDir, os.ModePerm)
	//}

	//filepath := "C:\\Users\\EDZ\\go\\src\\awesomeProject\\test.jpg"
	//
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

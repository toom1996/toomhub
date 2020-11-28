// @Description
// @Author    2020/10/27 9:19
package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"strconv"
	"toomhub/middware"
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
		controller.POST("/upload1", image.Upload1)
	}

	controller.Use(middware.CheckIdentity())
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
func (*VideoController) Upload1(context *gin.Context) {
	chunkSizeQuery := context.Query("chunkSize")
	totalSizeQuery := context.Query("chunkSize")
	chunkSize, _ := strconv.Atoi(chunkSizeQuery)
	totalSize, _ := strconv.Atoi(totalSizeQuery)
	//rand.Seed(int64(uint64(time.Now().UnixNano())))
	//buf := new(bytes.Buffer)
	//_, _ = buf.ReadFrom(context.Request.Body)
	//fmt.Println(buf.String())
	//
	//fmt.Println()
	//randNum := strconv.Itoa(rand.Intn(1000))
	//fileName := "test" + randNum
	//fmt.Println(fileName)
	//f, err1 := os.Create(fileName) //创建文件
	//fmt.Println(err1)
	//
	//n, err1 := io.WriteString(f, buf.String()) //写入文件(字符串)
	//fmt.Println(err1)
	//fmt.Printf("写入 %d 个字节n", n)

	context.JSON(200, gin.H{
		"code":    200,
		"message": "上传成功",
	})
	config := util.GetConfig()
	putPolicy := storage.PutPolicy{
		Scope: "toomhub",
	}
	mac := qbox.NewMac(config.Qiniu.AccessKey, config.Qiniu.SecretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuabei
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	resumeUploader := storage.NewResumeUploader(&cfg)
	//ret := storage.PutRet{}
	ret1 := storage.BlkputRet{}
	//putExtra := storage.RputExtra{}
	mkblr := resumeUploader.Mkblk(util.Ctx, upToken, "http://up-z1.qiniup.com", &ret1, chunkSize, context.Request.Body, totalSize)
	//err := resumeUploader.PutFile(util.Ctx, &ret, upToken, "test", fileName, &putExtra)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	fmt.Println(mkblr)
	fmt.Println(ret1)

}

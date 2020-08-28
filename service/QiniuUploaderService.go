// @Description
// @Author    2020/8/27 17:09
package service

import (
	"context"
	"fmt"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"mime/multipart"
	"os"
	"toomhub/util"
)

// 获取文件大小接口
type Size interface {
	Size() int64
}

// 获取文件信息接口
type Stat interface {
	// 文件解析状态
	Stat() (os.FileInfo, error)
}

// 构造返回字段
type PutRet struct {
	Hash string `json:"hash"`
	Key  string `json:"key"`
}

type QiniuUploader struct {
}

// 加载
func (u QiniuUploader) Upload(file multipart.File, fileName string) (url string, err error) {
	config := util.GetConfig()
	var (
		ret  PutRet
		size int64
	)
	if statInterface, ok := file.(Stat); ok {
		// 获取文件状态
		fileInfo, _ := statInterface.Stat()
		size = fileInfo.Size()
		fmt.Println("0000000000")
		fmt.Println(fileInfo.Name())
	}
	if sizeInterface, ok := file.(Size); ok {
		size = sizeInterface.Size()
	}

	fmt.Println(size)
	putPolicy := storage.PutPolicy{
		Scope: config.Qiniu.Bucket,
	}
	//NewMac 构建一个新的拥有AK/SK的对象
	mac := qbox.NewMac(config.Qiniu.AccessKey, config.Qiniu.SecretKey)
	//UploadToken 方法用来进行上传凭证的生成
	token := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 上传的机房
	cfg.Zone = &storage.ZoneHuabei
	// 不启用https域名
	cfg.UseHTTPS = false
	// 不适用cnd加速
	cfg.UseCdnDomains = false
	// NewFormUploader 用来构建一个表单上传的对象
	uploader := storage.NewFormUploader(&cfg)
	//PutExtra 为表单上传的额外可选
	//项
	putExtra := storage.PutExtra{}
	//	 上传文件

	err = uploader.Put(context.Background(), &ret, token, fileName, file, size, &putExtra)
	if err != nil {
		fmt.Println(err)
		return
	}
	url = config.Qiniu.FileServer + ret.Key
	return

}

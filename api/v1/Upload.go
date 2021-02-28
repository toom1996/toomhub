// @Description
// @Author: toom1996 <1023150697@qq.com>
// @dateTime: 2020/12/21 14:02
package v1

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"toomhub/setting"
	"toomhub/util"
)

func GetQiniuAccessToken(context *gin.Context) {
	bucket := "toomhub"
	n := base64.StdEncoding.EncodeToString([]byte("zawazawa-image-thumb:$(key)"))
	fmt.Println(n)
	putPolicy := storage.PutPolicy{
		Scope:         bucket,
		ReturnBody:    `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"format":"$(imageInfo.format)","height":"$(imageInfo.height)", "width":"$(imageInfo.width)"}`,
		PersistentOps: "imageView2/2/w/200/h/200|saveas/" + n,
	}
	putPolicy.Expires = 100
	mac := qbox.NewMac(setting.ZConfig.Qiniu.AccessKey, setting.ZConfig.Qiniu.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	util.ResponseOk(context, "OK", upToken)
}

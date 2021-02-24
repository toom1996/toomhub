// @Description
// @Author: toom1996 <1023150697@qq.com>
// @dateTime: 2020/12/21 14:02
package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"toomhub/setting"
	"toomhub/util"
)

func GetQiniuAccessToken(context *gin.Context) {
	bucket := "zawazawa"
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	putPolicy.Expires = 100
	mac := qbox.NewMac(setting.ZConfig.Qiniu.AccessKey, setting.ZConfig.Qiniu.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	util.ResponseOk(context, "OK", upToken)
}

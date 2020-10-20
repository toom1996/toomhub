// @Description	小程序接口验证器
// @Author    2020/8/14 10:53
package validatorRules

import ModelMiniV1 "toomhub/model/mini/v1"

type Login struct {
	RawData       ModelMiniV1.V1MiniUserInfo `form:"rawData" binding:"required"`
	Iv            string                     `form:"iv" binding:"required"`
	EncryptedData string                     `form:"encryptedData" binding:"required"`
	Signature     string                     `form:"signature" binding:"required"`
	AuthKey       string                     `form:"authKey" binding:"required"`
}

type Refresh struct {
	Token        string `form:"token" binding:"required"`
	RefreshToken string `form:"refresh_token"`
}

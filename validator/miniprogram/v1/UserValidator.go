// @Description	小程序接口验证器
// @Author    2020/8/14 10:53
package validatorMiniprogramV1

import (
	ModelMiniV1 "toomhub/model/mini/v1"
)

type Login struct {
	Code     string                     `form:"code" binding:"required"`
	UserInfo ModelMiniV1.V1MiniUserInfo `form:"userInfo"`
}

type Refresh struct {
	Token        string `form:"token" binding:"required"`
	RefreshToken string `form:"refresh_token"`
}

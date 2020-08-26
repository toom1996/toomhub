// @Description	小程序接口验证器
// @Author    2020/8/14 10:53
package validatorMiniprogramV1

import "toomhub/model"

type Login struct {
	Code     string               `form:"code" binding:"required"`
	UserInfo model.V1MiniUserInfo `form:"userInfo"`
}

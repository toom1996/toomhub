// @Description
// @Author    2020/8/17 16:44
package service

import "toom/tool"

// @description   通过OPENID判断当前用户是否存在
// @auth	toom <1023150697@qq.com>
// @param    前端传递的OPENID
// @return BOOL
func V1MiniUserHasUser(openid string) bool {
	_, _ = DB.Raw("SELECT * FROM toomhub_user_mini WHERE 1").Rows()
	return false
}

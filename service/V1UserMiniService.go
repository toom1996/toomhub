// @Description
// @Author    2020/8/17 16:44
package service

import "fmt"

// @description   通过OPENID判断当前用户是否存在
// @auth	toom <1023150697@qq.com>
// @param    前端传递的OPENID
// @return BOOL
func V1MiniUserHasUser(openid string) bool {
	sql := "SELECT open_id FROM toomhub_user_mini WHERE openid = '" + openid + "'"
	query, err := DB.Raw(sql).Rows()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(query)
	return false
}

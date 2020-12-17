// @Description
// @Author: toom1996 <1023150697@qq.com>
// @dateTime: 2020/12/15 16:38
package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"toomhub/model"
	rules "toomhub/rules/user/v1"
	"toomhub/util"
)

func V1UserRegister(validator *rules.V1UserRegister) (map[string]interface{}, error) {
	m1 := model.ZawazawaUser{}
	query := model.ZawazawaUserMgr(util.DB).Select([]string{"id", "mobile"}).Where(&model.ZawazawaUser{
		Mobile: validator.Mobile,
	}).Take(&m1).RowsAffected

	fmt.Println(m1)
	fmt.Println(query)

	if query == 1 {
		return nil, errors.New("当前手机号码已被注册")
	}

	return gin.H{
		"name":   "toom",
		"age":    "11",
		"gender": "boy",
	}, nil
}

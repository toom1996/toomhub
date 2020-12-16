// @Description
// @Author: toom1996 <1023150697@qq.com>
// @dateTime: 2020/12/15 16:38
package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"toomhub/model"
	"toomhub/util"
)

func V1PcRegisterUser() (map[string]interface{}, error) {

	m1 := model.ZawazawaUser{}
	model.ZawazawaUserMgr(util.DB).Debug().Select([]string{"id"}).Where(&model.ZawazawaUser{
		Mobile: "13502127317",
	}).Take(&m1)
	fmt.Println(m1)

	return gin.H{}, nil
}

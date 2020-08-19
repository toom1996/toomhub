// @Description
// @Author    2020/8/19 14:38
package ServiceMiniV1

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"time"
	ModelMiniV1 "toomhub/model/mini/v1"
	"toomhub/service"
)

// @title	通过OPENID获取用户信息
// @description
// @auth	toom <1023150697@qq.com>
// @param    前端传递的OPENID
// @return BOOL
func GetUser(openid string) (interface{}, error) {
	DB := service.DB
	tableModel := ModelMiniV1.ToomhubUserMini{}
	//根据openid 查找用户
	err := DB.Where("openid = ?", openid).Take(&tableModel).Error
	if err != nil {
		return "", err
	}

	if gorm.IsRecordNotFoundError(err) {
		//TODO 插入一个新用户
		transaction := DB.Begin()
		user := ModelMiniV1.ToomhubUserMini{
			Openid:    openid,
			CreatedAt: time.Now().Unix(),
		}
		err := DB.Create(&user).Error
		if err != nil {
			transaction.Rollback()
			return "", err
		}

		return gin.H{
			"a": "000",
		}, err
	}

	return false, err
}

func V1MiniUserCreate() {

}

func GetToken() {

}

func SetRedis() {

}

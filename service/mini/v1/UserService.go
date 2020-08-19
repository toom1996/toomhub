// @Description
// @Author    2020/8/19 14:38
package ServiceMiniV1

import (
	"github.com/jinzhu/gorm"
	"time"
	ModelMiniV1 "toomhub/model/mini/v1"
	"toomhub/service"
)

// @title	通过OPENID获取用户信息
// @description
// @auth	toom <1023150697@qq.com>
func GetUser(openid string) (interface{}, error) {
	DB := service.DB
	tableModel := ModelMiniV1.ToomhubUserMini{}
	//根据openid 查找用户
	query := service.DB.Where("open_id = ?", openid).Take(&tableModel)

	//如果err不等于record not found错误，又不等于nil，那说明sql执行失败了。
	if gorm.IsRecordNotFoundError(query.Error) {
		//TODO 插入一个新用户
		res, err := UserCreate(openid, DB)
		if err != nil {
			return "", err
		}
		return res, nil
	} else if query.Error != nil {
		return "", query.Error
	}

	return query.Value, nil
}

// @title	创建一个新的小程序用户
// @auth toom <1023150697@qq.com>
func UserCreate(openid string, DB *gorm.DB) (interface{}, error) {
	transaction := DB.Begin()
	user := ModelMiniV1.ToomhubUserMini{
		OpenId:    openid,
		CreatedAt: time.Now().Unix(),
	}
	err := DB.Create(&user).Error
	if err != nil {
		transaction.Rollback()
		return "", err
	}
	return user, err
}

func GetToken() {

}

func SetRedis() {

}

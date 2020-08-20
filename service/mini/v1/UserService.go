// @Description
// @Author    2020/8/19 14:38
package ServiceMiniV1

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"strconv"
	"time"
	ModelMiniV1 "toomhub/model/mini/v1"
	"toomhub/util"
)

const UserCacheKey = "mini:user:"

// @title	通过OPENID获取用户信息
// @description
// @auth	toom <1023150697@qq.com>
func GetUser(openid string, validator *ModelMiniV1.LoginByV1Model) (interface{}, error) {
	DB := util.DB
	tableModel := ModelMiniV1.ToomhubUserMini{}
	//根据openid 查找用户
	query := util.DB.Where("open_id = ?", openid).Take(&tableModel)

	//如果err不等于record not found错误，又不等于nil，那说明sql执行失败了。
	if gorm.IsRecordNotFoundError(query.Error) {
		//TODO 插入一个新用户
		res, err := UserCreate(openid, DB, validator)
		if err != nil {
			return "", err
		}
		return res, nil
	} else if query.Error != nil {
		return "", query.Error
	}

	return query.Value, nil
}

type UserInfo struct {
	Id           string
	AccessToken  string
	RefreshToken string
}

type RedisInfo struct {
	OpenId    string
	CreatedAt int64
	NickName  string
	Gender    int8
	City      string
	Province  string
	Country   string
	AvatarUrl string
}

// @title	创建一个新的小程序用户
// @auth toom <1023150697@qq.com>
func UserCreate(openid string, DB *gorm.DB, validator *ModelMiniV1.LoginByV1Model) (interface{}, error) {
	createTime := time.Now().Unix()
	//开启事务
	transaction := DB.Begin()

	//插入到用户表
	userModel := ModelMiniV1.ToomhubUserMini{
		OpenId:    openid,
		CreatedAt: createTime,
	}
	userQuery := transaction.Create(&userModel).Scan(&userModel)
	if userQuery.Error != nil {
		transaction.Rollback()
		return "", userQuery.Error
	}

	//插入到用户信息表
	profileModel := ModelMiniV1.ToomhubUserMiniProfile{
		MiniId:    userModel.MiniId,
		NickName:  validator.UserInfo.Nickname,
		Gender:    validator.UserInfo.Gender,
		City:      validator.UserInfo.City,
		Province:  validator.UserInfo.Province,
		Country:   validator.UserInfo.Country,
		AvatarUrl: validator.UserInfo.AvatarUrl,
	}
	profileQuery := transaction.Create(&profileModel)
	if profileQuery.Error != nil {
		transaction.Rollback()
		return "", profileQuery.Error
	}

	//生成token
	token, err := util.GenerateToken(userModel.OpenId)
	if err != nil {
		fmt.Println(err)
	}
	tokenModel := ModelMiniV1.ToomhubUserMiniToken{
		MiniId:       userModel.MiniId,
		AccessToken:  token,
		RefreshToken: util.GetRandomString(64),
		CreatedAt:    createTime,
		UpdatedAt:    createTime,
	}
	err = transaction.Create(&tokenModel).Error
	if err != nil {
		fmt.Println(err)
	}

	//提交事务
	transaction.Commit()
	if transaction.Error != nil {
		fmt.Println(transaction.Error)
	}

	//塞入redis
	err = util.Rdb.HMSet(util.Ctx, UserCacheKey+strconv.Itoa(userModel.MiniId), map[string]interface{}{
		"open_id":    userModel.OpenId,
		"created_at": userModel.CreatedAt,
		"nick_name":  profileModel.NickName,
		"gender":     profileModel.Gender,
		"city":       profileModel.City,
		"province":   profileModel.Province,
		"country":    profileModel.Country,
		"avatar_url": profileModel.AvatarUrl,
	}).Err()
	if err != nil {
		fmt.Println(err)
	}
	return userModel, err
}

func GetToken() {

}

func SetRedis() {

}

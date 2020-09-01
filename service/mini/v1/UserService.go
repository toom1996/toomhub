// @Description
// @Author    2020/8/19 14:38
package ServiceMiniV1

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"strconv"
	"time"
	ModelMiniV1 "toomhub/model/mini/v1"
	"toomhub/util"
	validatorMiniprogramV1 "toomhub/validator/miniprogram/v1"
)

const UserCacheKey = "mini:user:"

// @title	通过OPENID获取用户信息
// @description
// @auth	toom <1023150697@qq.com>
func GetUser(openid string, validator *validatorMiniprogramV1.Login) (interface{}, error) {
	db := util.DB
	res, err := GetUserByOpenId(openid)
	//如果是未找到的openid
	if gorm.IsRecordNotFoundError(err) {
		//插入一个新用户
		res, err := UserCreate(openid, db, validator)
		fmt.Println("insert new user")
		if err != nil {
			return "", err
		}
		return res, nil
	}

	return res, nil
}

// @title 通过openid查找用户信息
func GetUserByOpenId(openid string) (interface{}, error) {
	db := util.DB
	tableModel := ModelMiniV1.ToomhubUserMini{}
	//根据openid 查找用户
	query := db.Where("open_id = ?", openid).Take(&tableModel)
	if query.Error != nil {
		fmt.Println(query.Error)
		return false, query.Error
	}

	fmt.Println("1111111")
	//通过用户id从REDIS中获取信息
	info, err := GetUserInfoByRedis(tableModel.MiniId)
	fmt.Println("22222")
	fmt.Println(info)
	fmt.Println(err)

	return info, nil
}

type UserInfo struct {
	Id           string
	AccessToken  string
	RefreshToken string
}

// @title	创建一个新的小程序用户
// @auth toom <1023150697@qq.com>
func UserCreate(openid string, DB *gorm.DB, validator *validatorMiniprogramV1.Login) (interface{}, error) {
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

	_, _ = SetUserInfoToRedis(userModel, profileModel, token)

	//提交事务
	transaction.Commit()
	if transaction.Error != nil {
		fmt.Println(transaction.Error)
	}

	return userModel, err
}

type RedisUserInfo struct {
	MiniId    int
	OpenId    string
	CreatedAt int64
	NickName  string
	Gender    int8
	City      string
	Province  string
	Country   string
	AvatarUrl string
}

// @title	从REDIS中获取用户信息
func GetUserInfoByRedis(userId int) (interface{}, error) {
	//从redis中获取
	id := strconv.Itoa(userId)
	query := util.Rdb.HMGet(util.Ctx, UserCacheKey+id, []string{
		"mini_id",
		"open_id",
		"avatar_url",
		"created_at",
		"nick_name",
		"gender",
		"city",
		"province",
		"country",
		"token",
	}...)
	if query.Err() != nil {
		fmt.Println(query.Err())
		return "", query.Err()
	}
	res, err := query.Result()
	if err != nil {
		fmt.Println(err)
		return "", query.Err()
	}

	fmt.Println("009999999")
	fmt.Println(res)
	if res[0] != nil {
		m := map[string]interface{}{
			"MiniId":    res[0],
			"OpenId":    res[1],
			"AvatarUrl": res[2],
			"CreatedAt": res[3],
			"NickName":  res[4],
			"Gender":    res[5],
			"City":      res[6],
			"Province":  res[7],
			"Country":   res[8],
			"Token":     res[9],
		}

		fmt.Println("GetUserInfoByRedis")
		fmt.Println(m)
		return m, nil
	}

	return "", errors.New("unknown error")
}

// @title	将用户信息塞入REDIS缓存
func SetUserInfoToRedis(userModel ModelMiniV1.ToomhubUserMini, profileModel ModelMiniV1.ToomhubUserMiniProfile, token string) (bool, error) {
	key := UserCacheKey + strconv.Itoa(userModel.MiniId)
	//塞入redis
	err := util.Rdb.HMSet(util.Ctx, key, map[string]interface{}{
		"mini_id":    userModel.MiniId,
		"open_id":    userModel.OpenId,
		"created_at": userModel.CreatedAt,
		"nick_name":  profileModel.NickName,
		"gender":     profileModel.Gender,
		"city":       profileModel.City,
		"province":   profileModel.Province,
		"country":    profileModel.Country,
		"avatar_url": profileModel.AvatarUrl,
		"token":      token,
	}).Err()

	////设置一周的过期时间
	//util.Rdb.Expire(util.Ctx, key, time.Second*60*60*24*7)

	if err != nil {
		fmt.Println(err)
	}
	return true, nil
}

func SetRedis() {

}

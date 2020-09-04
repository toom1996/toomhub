// @Description
// @Author    2020/8/14 17:03
package LogicMiniV1

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/medivhzhan/weapp/v2"
	ServiceMiniV1 "toomhub/service/mini/v1"
	"toomhub/util"
	validatorMiniprogramV1 "toomhub/validator/miniprogram/v1"
)

type UserLogic struct {
}

// @title	小程序登陆
func (logic *UserLogic) Login(validator *validatorMiniprogramV1.Login) (interface{}, error) {
	config := util.GetConfig()
	//微信接口验证
	res, err := weapp.Login(config.Mini.AppId, config.Mini.AppSecret, validator.Code)
	if err != nil {
		return map[string]string{}, err
	}
	if res.ErrCode != 0 {
		return map[string]string{}, errors.New(res.ErrMSG)
	}

	//数据库验证用户信息
	userInfo, err := ServiceMiniV1.GetUser(res.OpenID, validator)
	if err != nil {
		return "", err
	}

	fmt.Println(userInfo)
	return userInfo, err
}

func (logic *UserLogic) Refresh(validator *validatorMiniprogramV1.Refresh) (interface{}, error) {
	tokenClaims, err := jwt.ParseWithClaims(validator.Token, &util.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return util.GetConfig().Jwt.Secret, nil
	})

	if tokenClaims != nil {
		if tokenClaims.Valid {
			//不是过期的token不允许刷新
			if _, ok := tokenClaims.Claims.(*util.Claims); ok {
			}
		} else if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				//ServiceMiniV1.Refresh(validator)
				fmt.Println("refresh")
			}
		}
	}
	return "", errors.New("refresh handle error")
}

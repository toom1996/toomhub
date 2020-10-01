// @Description
// @Author    2020/8/14 17:03
package LogicMiniV1

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/medivhzhan/weapp/v2"
	"time"
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
	fmt.Println("we login ------>", res)
	//数据库验证用户信息
	userInfo, err := ServiceMiniV1.GetUser(res.OpenID, validator)
	if err != nil {
		return "", err
	}

	fmt.Println(userInfo)
	return userInfo, err
}

// @title	检查token
func (logic *UserLogic) Check(validator *validatorMiniprogramV1.Refresh, c *gin.Context) (interface{}, error) {
	tokenClaims, err := jwt.ParseWithClaims(validator.Token, &util.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(util.GetConfig().Jwt.Secret), nil
	})

	if tokenClaims != nil {
		if tokenClaims.Valid {
			if claims, ok := tokenClaims.Claims.(*util.Claims); ok {
				t := 60 * 60 * 24 * 5
				if time.Now().Unix()-claims.CreatedAt > int64(t) {
					r, _ := ServiceMiniV1.UpdateUserInfoToRedis(claims.MiniId)
					return r, nil
				}
				return nil, nil
			}
		} else if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				return nil, errors.New("token is expired")
			}
		}
	}
	return nil, errors.New("bad token")
}

// @Description
// @Author    2020/8/14 17:03
package LogicMiniV1

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
	ServiceMiniV1 "toomhub/service/mini/v1"
	"toomhub/util"
	"toomhub/validatorRules"
)

type UserLogic struct {
}

// @title	小程序登陆
func (logic *UserLogic) Login(validator *validatorRules.Login) (interface{}, error) {

	cacheInfo, err := util.Rdb.Get(util.Ctx, validator.AuthKey).Result()
	fmt.Println("cacheInfo -> ", cacheInfo)

	if err != nil {
		return "", err
	}

	sessionCache, err := util.JsonDecode(cacheInfo)
	if err != nil {
		return "", err
	}

	rawData, _ := util.JsonEncode(validator.RawData)

	sign := util.Sha1(rawData + sessionCache["session_key"].(string))

	//验证签名
	if sign != validator.Signature {
		return "", errors.New("signature validate fail")
	}
	//数据库验证用户信息
	userInfo, err := ServiceMiniV1.GetUser(sessionCache["openid"].(string), validator)
	if err != nil {
		return "", err
	}

	fmt.Println(userInfo)
	return userInfo, err
}

// @title	检查token
func (logic *UserLogic) Check(validator *validatorRules.Refresh, c *gin.Context) (interface{}, error) {
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

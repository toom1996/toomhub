// @Description
// @Author    2020/8/20 10:01
package util

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	"toomhub/setting"
)

//var jwtSecret = []byte("toomhub")

type Claims struct {
	Type string
	jwt.StandardClaims
}

var identity *Claims

func GenerateToken(id uint) (string, error) {
	nowTime := time.Now()
	expire, _ := strconv.Atoi(setting.ZConfig.Jwt.JwtExpire)
	expireTime := nowTime.Add(time.Duration(expire) * time.Second)

	claims := Claims{
		"token",
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "zawazawa",
			Id:        strconv.Itoa(int(id)),
			IssuedAt:  nowTime.Unix(),
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(setting.ZConfig.Jwt.JwtSecret))
	return token, err
}

func GenerateRefreshToken(id uint) (string, error) {
	nowTime := time.Now()
	expire, _ := strconv.Atoi(setting.ZConfig.Jwt.JwtExpire)
	expireTime := nowTime.Add(time.Duration(expire) * time.Second)

	claims := Claims{
		"refresh_token",
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "zawazawa",
			Id:        strconv.Itoa(int(id)),
			IssuedAt:  nowTime.Unix(),
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(setting.ZConfig.Jwt.JwtSecret))
	return token, err
}

func ParseToken(token string, c *gin.Context) (*Claims, error) {
	fmt.Println(token)
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(setting.ZConfig.Jwt.JwtSecret), nil
	})

	if tokenClaims != nil {
		if tokenClaims.Valid {
			if claims, ok := tokenClaims.Claims.(*Claims); ok {

				// 判断token是否为最新
				identity = &Claims{
					claims.Type,
					jwt.StandardClaims{
						Id:        claims.Id,
						IssuedAt:  claims.IssuedAt,
						ExpiresAt: claims.ExpiresAt,
						Issuer:    claims.Issuer,
					},
				}
				return claims, nil
			}
		} else if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				// token 过期了
				c.JSON(http.StatusOK, map[string]interface{}{
					"code": 401,
					"msg":  "登陆超时,请重新登陆",
				})
				c.Abort()
				return nil, nil
			}
		}
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"code": 401,
		"msg":  "bad token",
	})
	c.Abort()
	return nil, nil
}

// @title 获取用户信息
func GetIdentity() *Claims {
	return identity
}

// @title 获取用户信息
func Identity(ctx *gin.Context) *Claims {
	token := ctx.GetHeader("Toomhub-Token")
	r := &Claims{}
	if token != "" {
		r, _ = ParseToken(token, ctx)
	}
	return r
}

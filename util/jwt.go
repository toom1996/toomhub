// @Description
// @Author    2020/8/20 10:01
package util

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte("toomhub")

type Claims struct {
	MiniId string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(id int) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		fmt.Sprintf("%d", id),
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "toomhub",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if tokenClaims.Valid {
			if claims, ok := tokenClaims.Claims.(*Claims); ok {
				return claims, nil
			}
			fmt.Println("success")
		} else if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				fmt.Println("bad token")
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				// Token is either expired or not active yet
				fmt.Println("Time out ")
			} else {
				fmt.Println("Couldn't handle this token:", err)
			}
		}
	}

	//错误的token
	if err != nil {
		fmt.Println("错误的token")
		fmt.Println(err)
	}

	if tokenClaims != nil {
		if tokenClaims.Valid {
			fmt.Println("验证通过")
		}
		fmt.Println("验证不通过")
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

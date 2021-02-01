// @Description
// @Author: toom1996 <1023150697@qq.com>
// @dateTime: 2020/12/15 16:38
package service

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	rules "toomhub/rules/user/v1"
	"toomhub/setting"
)

func V1UserSmsSend() {

}

func V1UserRegister(validator *rules.V1UserRegister) (map[string]interface{}, error) {

	return nil, nil
	//DB := util.DB
	//
	//tx := DB.Begin()
	//zUser := model.ZawazawaUser{Mobile: validator.Mobile}
	//if result := tx.Create(&zUser); result.Error != nil {
	//	tx.Rollback()
	//	return nil, result.Error
	//}
	//accessToken, err := generateAccessToken(int64(zUser.ID))
	//if err != nil {
	//	tx.Rollback()
	//	return nil, err
	//}
	//zUserToken := model.ZawazawaUserToken{
	//	UId:          zUser.ID,
	//	Token:        accessToken,
	//	RefreshToken: util.GetRandomString(10),
	//	Type:         "PC",
	//}
	//if result := tx.Create(&zUserToken); result.Error != nil {
	//	tx.Rollback()
	//	return nil, err
	//}
	//tx.Commit()
	//
	//return gin.H{
	//	"id":           zUser.ID,
	//	"zToken":       zUserToken.Token,
	//	"refreshToken": zUserToken.RefreshToken,
	//	"avatar":       "http://himg.bdimg.com/sys/portrait/item/2332313032333135303639378a08.jpg", //头像
	//	"nickname":     "zawazawa" + fmt.Sprintf("%d", zUser.ID),                                   //头像
	//}, nil
}

type Claims struct {
	MiniId    int64
	CreatedAt int64
	jwt.StandardClaims
}

//生成AccessToken
func generateAccessToken(id int64) (string, error) {
	nowTime := time.Now()
	//七天有效期
	expireTime := nowTime.Add(60 * time.Minute * 24 * 7)

	claims := Claims{
		id,
		nowTime.Unix(),
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "zawazawa",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(setting.ZConfig.App.JwtSecret))
	return token, err
}

// @Description
// @Author    2020/8/14 17:03
package LogicMiniV1

import (
	"errors"
	"fmt"
	"github.com/medivhzhan/weapp/v2"
	ModelMiniV1 "toomhub/model/mini/v1"
	ServiceMiniV1 "toomhub/service/mini/v1"
	"toomhub/tool"
)

type UserLogic struct {
}

// @title	小程序登陆
func (logic *UserLogic) Login(validator *ModelMiniV1.LoginByV1Model) (interface{}, error) {
	config := tool.GetConfig()
	//微信接口验证
	res, err := weapp.Login(config.Mini.AppId, config.Mini.AppSecret, validator.Code)
	if err != nil {
		return map[string]string{}, err
	}
	if res.ErrCode != 0 {
		return map[string]string{}, errors.New(res.ErrMSG)
	}

	//数据库验证用户信息
	userInfo, err := ServiceMiniV1.GetUser(res.OpenID)
	if err != nil {
		return "", err
	}

	fmt.Println(userInfo)
	return userInfo, err
}

// @Description
// @Author: toom1996 <1023150697@qq.com>
// @dateTime: 2021/1/29 15:46
package logic

import (
	rules "toomhub/rules/user/v1"
	"toomhub/service"
)

type UserLogic struct {
}

type token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"` // 这个字段没用到
	Scope       string `json:"scope"`      // 这个字段也没用到
}

func (l *UserLogic) GithubOAuthLogic(validator *rules.V1UserGithubOAuth) (interface{}, error) {
	var ser service.UserService
	var saveInfo map[string]interface{}
	//获取github信息
	info, err := ser.GetGithubOAuthInfo(validator)
	if err != nil {
		return nil, err
	}

	//判断是否存在此用户
	if isNew := ser.IsNewUser(info.GitID); isNew == false {
		//存在,更新
		_, _ = ser.UpdateGithubOAuthInfo()
		return nil, nil
	} else {
		//不存在,新增
		saveInfo, err = ser.SaveGithubOAuthInfo(&info)
		if err != nil {
			return nil, err
		}

	}
	return map[string]interface{}{
		"avatar":   saveInfo["avatar"],
		"username": saveInfo["username"],
		"token":    "tokentest1234",
	}, nil
}

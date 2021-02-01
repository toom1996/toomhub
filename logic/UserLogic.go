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
	//获取github信息
	info, err := ser.GetGithubOAuthInfo(validator)
	if err != nil {
		return nil, err
	}

	//判断是否存在此用户
	if isNew := ser.IsNewUser(); isNew == false {
		//存在,更新
		_, _ = ser.UpdateGithubOAuthInfo()
		return nil, nil
	} else {
		//不存在,新增
		_, _ = ser.SaveGithubOAuthInfo(&info)

	}
	return map[string]interface{}{
		"avatar":   "http://47.105.189.195:2000/img/WechatIMG2.1b46015c.jpeg",
		"username": "toom",
	}, nil
}

// @Description
// @Author: toom1996 <1023150697@qq.com>
// @dateTime: 2021/1/29 15:46
package logic

import (
	"encoding/json"
	"fmt"
	"net/http"
	rules "toomhub/rules/user/v1"
	service2 "toomhub/service"
	"toomhub/setting"
)

type UserLogic struct {
}

type token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"` // 这个字段没用到
	Scope       string `json:"scope"`      // 这个字段也没用到
}

func (l *UserLogic) GithubOAuthLogic(validator *rules.V1UserGithubOAuth) (interface{}, error) {
	//编译好链接
	s := fmt.Sprintf(
		"https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s",
		setting.ZConfig.GithubOAuth.ClientId, setting.ZConfig.GithubOAuth.ClientSecret, validator.Code,
	)
	var err error
	// 形成请求
	var req *http.Request
	if req, err = http.NewRequest(http.MethodGet, s, nil); err != nil {
		return nil, err
	}
	req.Header.Set("accept", "application/json")

	// 发送请求并获得响应
	var httpClient = http.Client{}
	var res *http.Response
	if res, err = httpClient.Do(req); err != nil {
		return false, err
	}

	// 将响应体解析为 token，并返回
	var token token
	if err = json.NewDecoder(res.Body).Decode(&token); err != nil {
		return nil, err
	}
	fmt.Println(&token)

	// 通过token，获取用户信息
	var userInfo map[string]interface{}
	var service service2.UserService
	if userInfo, err = service.GetGithubOAuthInfo(token.AccessToken); err != nil {
		fmt.Println("获取用户信息失败，错误信息为:", err)
		return nil, err
	}
	return userInfo, nil
}

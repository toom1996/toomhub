// @Description
// @Author: toom1996 <1023150697@qq.com>
// @dateTime: 2021/1/29 16:26
package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"toomhub/model"
	rules "toomhub/rules/user/v1"
	"toomhub/setting"
	"toomhub/util"
)

type UserService struct {
}

type token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"` // 这个字段没用到
	Scope       string `json:"scope"`      // 这个字段也没用到
	State       string `json:"state"`      // 这个字段也没用到
}

type GithubOAuth struct {
	AvatarURL         string `json:"avatar_url"`          // 头像地址
	Bio               string `json:"bio"`                 // 个性签名
	Blog              string `json:"blog"`                // 博客地址好像是
	Company           string `json:"company"`             // 公司
	CreatedAt         string `json:"created_at"`          // 创建日期
	Email             string `json:"email"`               // 邮箱地址
	EventsURL         string `json:"events_url"`          // 不知道
	Followers         uint32 `json:"followers"`           // 粉丝数量
	FollowersURL      string `json:"followers_url"`       // 粉丝列表地址
	Following         uint32 `json:"following"`           // 关注用户
	FollowingURL      string `json:"following_url"`       // 关注用户列表地址
	GistsURL          string `json:"gists_url"`           // 不知道是什么
	Hireable          string `json:"hireable"`            // 不知道是什么
	HTMLURL           string `json:"html_url"`            // 主页地址
	GitID             uint32 `json:"id"`                  // github用户id
	Location          string `json:"location"`            // 定位??
	Login             string `json:"login"`               // git号
	Name              string `json:"name"`                // git昵称
	NodeID            string `json:"node_id"`             // 节点id??
	OrganizationsURL  string `json:"organizations_url"`   // 不知道
	PublicGists       uint32 `json:"public_gists"`        // 不知道
	PublicRepos       uint32 `json:"public_repos"`        // 开放的仓库数量
	ReceivedEventsURL string `json:"received_events_url"` // 不知道
	ReposURL          string `json:"repos_url"`           // 不知道
	StarredURL        string `json:"starred_url"`         // 不知道
	SubscriptionsURL  string `json:"subscriptions_url"`   // 仓库列表
	TwitterUsername   string `json:"twitter_username"`    // 推特用户名?
	Type              string `json:"type"`                // 不知道是什么类型
	URL               string `json:"url"`                 // 个人主页地址
}

func (service *UserService) GetGithubOAuthInfo(validator *rules.V1UserGithubOAuth) (GithubOAuth, error) {

	//编译好链接
	s := fmt.Sprintf(
		"https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s",
		setting.ZConfig.GithubOAuth.ClientId, setting.ZConfig.GithubOAuth.ClientSecret, validator.Code,
	)
	var err error
	// 形成请求
	var req *http.Request
	if req, err = http.NewRequest(http.MethodGet, s, nil); err != nil {
		return GithubOAuth{}, err
	}

	req.Header.Set("accept", "application/json")

	//发送请求并获得响应
	var httpClient = http.Client{}
	var res *http.Response
	if res, err = httpClient.Do(req); err != nil {
		return GithubOAuth{}, err
	}

	// 将响应体解析为 token，并返回
	var token token
	if err = json.NewDecoder(res.Body).Decode(&token); err != nil {
		return GithubOAuth{}, err
	}
	fmt.Println(&token)

	var userInfo = GithubOAuth{}
	// 形成请求
	var userInfoUrl = "https://api.github.com/user" // github用户信息获取接口

	if req, err = http.NewRequest(http.MethodGet, userInfoUrl, nil); err != nil {
		return userInfo, err
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("token %s", token.AccessToken))

	// 发送请求并获取响应
	var client = http.Client{}
	if res, err = client.Do(req); err != nil {
		return userInfo, err
	}

	// 将响应的数据写入 userInfo 中，并返回
	if err = json.NewDecoder(res.Body).Decode(&userInfo); err != nil {
		return userInfo, err
	}
	fmt.Println(userInfo.GitID)
	return userInfo, nil
}

//是否为新用户
//查数据库
func (service *UserService) IsNewUser() bool {

	return true
}

//存储github信息
func (service *UserService) SaveGithubOAuthInfo(info *GithubOAuth) (interface{}, error) {
	fmt.Println("SaveGithubOAuthInfo")
	//存入redis
	//存入数据库
	err := model.ZawazawaUserMgr(util.DB).Create(&info).Error
	if err != nil {
		fmt.Println(err)
	}

	return info, nil
}

//更新github信息
func (service *UserService) UpdateGithubOAuthInfo() (interface{}, error) {
	fmt.Println("UpdateGithubOAuthInfo")
	return nil, nil
}

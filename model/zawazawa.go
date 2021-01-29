package model

import (
	"time"
)

// ZawazawaPost [...]
type ZawazawaPost struct {
	ID        int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	Content   string `gorm:"column:content;type:varchar(255)" json:"content"`
	CreatedAt int    `gorm:"column:created_at;type:int(11)" json:"created_at"`
	CreatedBy int    `gorm:"column:created_by;type:int(11)" json:"created_by"`
}

// ZawazawaPostImage [...]
type ZawazawaPostImage struct {
	ID       int    `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	Rid      int    `gorm:"column:rid;type:int(11) unsigned;not null" json:"rid"`
	Host     string `gorm:"column:host;type:varchar(255);not null" json:"host"`
	Name     string `gorm:"column:name;type:varchar(255);not null" json:"name"`
	IsDelete bool   `gorm:"column:is_delete;type:tinyint(1) unsigned;not null" json:"is_delete"`
}

// ZawazawaUser [...]
type ZawazawaUser struct {
	ID                int       `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	AvatarURL         string    `gorm:"column:avatar_url;type:varchar(255)" json:"avatar_url"`                   // 头像地址
	Bio               string    `gorm:"column:bio;type:varchar(255)" json:"bio"`                                 // 个性签名
	Blog              string    `gorm:"column:blog;type:varchar(255)" json:"blog"`                               // 博客地址好像是
	Company           string    `gorm:"column:company;type:varchar(255)" json:"company"`                         // 公司
	CreatedAt         string    `gorm:"column:created_at;type:varchar(255)" json:"created_at"`                   // 创建日期
	Email             string    `gorm:"column:email;type:varchar(255)" json:"email"`                             // 邮箱地址
	EventsURL         string    `gorm:"column:events_url;type:varchar(255)" json:"events_url"`                   // 不知道
	Followers         int       `gorm:"column:followers;type:int(11)" json:"followers"`                          // 粉丝数量
	FollowersURL      string    `gorm:"column:followers_url;type:varchar(255)" json:"followers_url"`             // 粉丝列表地址
	Following         int       `gorm:"column:following;type:int(11)" json:"following"`                          // 关注用户
	FollowingURL      string    `gorm:"column:following_url;type:varchar(255)" json:"following_url"`             // 关注用户列表地址
	GistsURL          string    `gorm:"column:gists_url;type:varchar(255)" json:"gists_url"`                     // 不知道是什么
	GravatarID        int       `gorm:"column:gravatar_id;type:int(11)" json:"gravatar_id"`                      // 不知道是什么
	Hireable          string    `gorm:"column:hireable;type:varchar(255)" json:"hireable"`                       // 不知道是什么
	HTMLURL           string    `gorm:"column:html_url;type:varchar(255)" json:"html_url"`                       // 主页地址
	GitID             int       `gorm:"column:git_id;type:int(11)" json:"git_id"`                                // github用户id
	Location          string    `gorm:"column:location;type:varchar(255)" json:"location"`                       // 定位??
	Login             string    `gorm:"column:login;type:varchar(255)" json:"login"`                             // git号
	Name              string    `gorm:"column:name;type:varchar(255)" json:"name"`                               // git昵称
	NodeID            int       `gorm:"column:node_id;type:int(11)" json:"node_id"`                              // 节点id??
	OrganizationsURL  string    `gorm:"column:organizations_url;type:varchar(255)" json:"organizations_url"`     // 不知道
	PublicGists       string    `gorm:"column:public_gists;type:varchar(255)" json:"public_gists"`               // 不知道
	PublicRepos       int       `gorm:"column:public_repos;type:int(11)" json:"public_repos"`                    // 开放的仓库数量
	ReceivedEventsURL string    `gorm:"column:received_events_url;type:varchar(255)" json:"received_events_url"` // 不知道
	ReposURL          string    `gorm:"column:repos_url;type:varchar(255)" json:"repos_url"`                     // 不知道
	SiteAdmin         bool      `gorm:"column:site_admin;type:tinyint(1)" json:"site_admin"`                     // 网站管理员??
	StarredURL        string    `gorm:"column:starred_url;type:varchar(255)" json:"starred_url"`                 // 不知道
	SubscriptionsURL  string    `gorm:"column:subscriptions_url;type:varchar(255)" json:"subscriptions_url"`     // 仓库列表
	TwitterUsername   string    `gorm:"column:twitter_username;type:varchar(255)" json:"twitter_username"`       // 推特用户名?
	Type              string    `gorm:"column:type;type:varchar(64)" json:"type"`                                // 不知道是什么类型
	UpdatedAt         time.Time `gorm:"column:updated_at;type:datetime" json:"updated_at"`                       // 更新时间
	URL               string    `gorm:"column:url;type:varchar(255)" json:"url"`                                 // 个人主页地址
}

// ZawazawaUserToken [...]
type ZawazawaUserToken struct {
	ID           int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	UId          int    `gorm:"column:uid;type:int(11)" json:"uid"`
	Token        string `gorm:"column:token;type:varchar(255)" json:"token"`
	RefreshToken string `gorm:"column:refresh_token;type:varchar(255)" json:"refresh_token"`
	Type         string `gorm:"column:type;type:varchar(255)" json:"type"`
}

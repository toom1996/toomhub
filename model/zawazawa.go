package model

// ZawazawaPost [...]
type ZawazawaPost struct {
	ID        int    `gorm:"primary_key;column:id;type:int;not null" json:"-"`
	Content   string `gorm:"column:content;type:varchar(255)" json:"content"`
	CreatedAt int    `gorm:"column:created_at;type:int" json:"created_at"`
	CreatedBy int    `gorm:"column:created_by;type:int" json:"created_by"`
}

// ZawazawaPostImage [...]
type ZawazawaPostImage struct {
	ID       uint   `gorm:"primary_key;column:id;type:int unsigned;not null" json:"-"`
	Rid      uint   `gorm:"column:rid;type:int unsigned;not null" json:"rid"`
	Host     string `gorm:"column:host;type:varchar(255);not null" json:"host"`
	Name     string `gorm:"column:name;type:varchar(255);not null" json:"name"`
	IsDelete uint8  `gorm:"column:is_delete;type:tinyint unsigned;not null" json:"is_delete"`
}

// ZawazawaUser [...]
type ZawazawaUser struct {
	ID           uint   `gorm:"primary_key;column:id;type:int unsigned;not null" json:"-"`
	Nickname     string `gorm:"column:nickname;type:varchar(64);not null" json:"nickname"`            // 昵称
	Mobile       string `gorm:"column:mobile;type:varchar(16);not null" json:"mobile"`                // 手机号
	FansCount    uint   `gorm:"column:fans_count;type:int unsigned;not null" json:"fans_count"`       // 粉丝数量
	FollowCount  uint   `gorm:"column:follow_count;type:int unsigned;not null" json:"follow_count"`   // 关注数量
	CollectCount uint   `gorm:"column:collect_count;type:int unsigned;not null" json:"collect_count"` // 收藏数量
	OauthID      uint   `gorm:"column:oauth_id;type:int unsigned;not null" json:"oauth_id"`           // 授权id
	OauthType    uint16 `gorm:"column:oauth_type;type:smallint unsigned;not null" json:"oauth_type"`  // 授权登陆类型 0=>未授权 1=> github
}

// ZawazawaUserProfileGithub [...]
type ZawazawaUserProfileGithub struct {
	GitOauthId        uint   `gorm:"primary_key;column:id;type:int unsigned;not null" json:"-"`
	AvatarURL         string `gorm:"column:avatar_url;type:varchar(255);not null" json:"avatar_url"`                   // 头像地址
	Bio               string `gorm:"column:bio;type:varchar(255);not null" json:"bio"`                                 // 个性签名
	Blog              string `gorm:"column:blog;type:varchar(255);not null" json:"blog"`                               // 博客地址好像是
	Company           string `gorm:"column:company;type:varchar(255);not null" json:"company"`                         // 公司
	CreatedAt         string `gorm:"column:created_at;type:varchar(255);not null" json:"created_at"`                   // 创建日期
	Email             string `gorm:"column:email;type:varchar(255);not null" json:"email"`                             // 邮箱地址
	EventsURL         string `gorm:"column:events_url;type:varchar(255);not null" json:"events_url"`                   // 不知道
	Followers         uint   `gorm:"column:followers;type:int unsigned;not null" json:"followers"`                     // 粉丝数量
	FollowersURL      string `gorm:"column:followers_url;type:varchar(255);not null" json:"followers_url"`             // 粉丝列表地址
	Following         uint   `gorm:"column:following;type:int unsigned;not null" json:"following"`                     // 关注用户
	FollowingURL      string `gorm:"column:following_url;type:varchar(255);not null" json:"following_url"`             // 关注用户列表地址
	GistsURL          string `gorm:"column:gists_url;type:varchar(255);not null" json:"gists_url"`                     // 不知道是什么
	Hireable          string `gorm:"column:hireable;type:varchar(255);not null" json:"hireable"`                       // 不知道是什么
	HTMLURL           string `gorm:"column:html_url;type:varchar(255);not null" json:"html_url"`                       // 主页地址
	Id                uint   `gorm:"unique;unique;column:id;type:int unsigned;not null" json:"id"`                     // github用户id
	Location          string `gorm:"column:location;type:varchar(255);not null" json:"location"`                       // 定位??
	Login             string `gorm:"column:login;type:varchar(255);not null" json:"login"`                             // git号
	Name              string `gorm:"column:name;type:varchar(255);not null" json:"name"`                               // git昵称
	NodeID            string `gorm:"column:node_id;type:varchar(64);not null" json:"node_id"`                          // 节点id??
	OrganizationsURL  string `gorm:"column:organizations_url;type:varchar(255);not null" json:"organizations_url"`     // 不知道
	PublicGists       uint   `gorm:"column:public_gists;type:int unsigned;not null" json:"public_gists"`               // 不知道
	PublicRepos       uint   `gorm:"column:public_repos;type:int unsigned;not null" json:"public_repos"`               // 开放的仓库数量
	ReceivedEventsURL string `gorm:"column:received_events_url;type:varchar(255);not null" json:"received_events_url"` // 不知道
	ReposURL          string `gorm:"column:repos_url;type:varchar(255);not null" json:"repos_url"`                     // 不知道
	StarredURL        string `gorm:"column:starred_url;type:varchar(255);not null" json:"starred_url"`                 // 不知道
	SubscriptionsURL  string `gorm:"column:subscriptions_url;type:varchar(255);not null" json:"subscriptions_url"`     // 仓库列表
	TwitterUsername   string `gorm:"column:twitter_username;type:varchar(255);not null" json:"twitter_username"`       // 推特用户名?
	Type              string `gorm:"column:type;type:varchar(64);not null" json:"type"`                                // 不知道是什么类型
	URL               string `gorm:"column:url;type:varchar(255);not null" json:"url"`                                 // 个人主页地址
}

// ZawazawaUserToken [...]
type ZawazawaUserToken struct {
	ID           int    `gorm:"primary_key;column:id;type:int;not null" json:"-"`
	UId          int    `gorm:"column:uid;type:int" json:"uid"`
	Token        string `gorm:"column:token;type:varchar(255)" json:"token"`
	RefreshToken string `gorm:"column:refresh_token;type:varchar(255)" json:"refresh_token"`
	Type         string `gorm:"column:type;type:varchar(255)" json:"type"`
}

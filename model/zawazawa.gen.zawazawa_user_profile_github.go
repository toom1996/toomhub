package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _ZawazawaUserProfileGithubMgr struct {
	*_BaseMgr
}

// ZawazawaUserProfileGithubMgr open func
func ZawazawaUserProfileGithubMgr(db *gorm.DB) *_ZawazawaUserProfileGithubMgr {
	if db == nil {
		panic(fmt.Errorf("ZawazawaUserProfileGithubMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ZawazawaUserProfileGithubMgr{_BaseMgr: &_BaseMgr{DB: db.Table("zawazawa_user_profile_github"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ZawazawaUserProfileGithubMgr) GetTableName() string {
	return "zawazawa_user_profile_github"
}

// Get 获取
func (obj *_ZawazawaUserProfileGithubMgr) Get() (result ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ZawazawaUserProfileGithubMgr) Gets() (results []*ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithGitOauthID git_oauth_id获取
func (obj *_ZawazawaUserProfileGithubMgr) WithGitOauthID(gitOauthID uint) Option {
	return optionFunc(func(o *options) { o.query["git_oauth_id"] = gitOauthID })
}

// WithAvatarURL avatar_url获取 头像地址
func (obj *_ZawazawaUserProfileGithubMgr) WithAvatarURL(avatarURL string) Option {
	return optionFunc(func(o *options) { o.query["avatar_url"] = avatarURL })
}

// WithBio bio获取 个性签名
func (obj *_ZawazawaUserProfileGithubMgr) WithBio(bio string) Option {
	return optionFunc(func(o *options) { o.query["bio"] = bio })
}

// WithBlog blog获取 博客地址好像是
func (obj *_ZawazawaUserProfileGithubMgr) WithBlog(blog string) Option {
	return optionFunc(func(o *options) { o.query["blog"] = blog })
}

// WithCompany company获取 公司
func (obj *_ZawazawaUserProfileGithubMgr) WithCompany(company string) Option {
	return optionFunc(func(o *options) { o.query["company"] = company })
}

// WithCreatedAt created_at获取 创建日期
func (obj *_ZawazawaUserProfileGithubMgr) WithCreatedAt(createdAt string) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithEmail email获取 邮箱地址
func (obj *_ZawazawaUserProfileGithubMgr) WithEmail(email string) Option {
	return optionFunc(func(o *options) { o.query["email"] = email })
}

// WithEventsURL events_url获取 不知道
func (obj *_ZawazawaUserProfileGithubMgr) WithEventsURL(eventsURL string) Option {
	return optionFunc(func(o *options) { o.query["events_url"] = eventsURL })
}

// WithFollowers followers获取 粉丝数量
func (obj *_ZawazawaUserProfileGithubMgr) WithFollowers(followers uint) Option {
	return optionFunc(func(o *options) { o.query["followers"] = followers })
}

// WithFollowersURL followers_url获取 粉丝列表地址
func (obj *_ZawazawaUserProfileGithubMgr) WithFollowersURL(followersURL string) Option {
	return optionFunc(func(o *options) { o.query["followers_url"] = followersURL })
}

// WithFollowing following获取 关注用户
func (obj *_ZawazawaUserProfileGithubMgr) WithFollowing(following uint) Option {
	return optionFunc(func(o *options) { o.query["following"] = following })
}

// WithFollowingURL following_url获取 关注用户列表地址
func (obj *_ZawazawaUserProfileGithubMgr) WithFollowingURL(followingURL string) Option {
	return optionFunc(func(o *options) { o.query["following_url"] = followingURL })
}

// WithGistsURL gists_url获取 不知道是什么
func (obj *_ZawazawaUserProfileGithubMgr) WithGistsURL(gistsURL string) Option {
	return optionFunc(func(o *options) { o.query["gists_url"] = gistsURL })
}

// WithHireable hireable获取 不知道是什么
func (obj *_ZawazawaUserProfileGithubMgr) WithHireable(hireable string) Option {
	return optionFunc(func(o *options) { o.query["hireable"] = hireable })
}

// WithHTMLURL html_url获取 主页地址
func (obj *_ZawazawaUserProfileGithubMgr) WithHTMLURL(htmlURL string) Option {
	return optionFunc(func(o *options) { o.query["html_url"] = htmlURL })
}

// WithID id获取 github用户id
func (obj *_ZawazawaUserProfileGithubMgr) WithID(id uint) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithLocation location获取 定位??
func (obj *_ZawazawaUserProfileGithubMgr) WithLocation(location string) Option {
	return optionFunc(func(o *options) { o.query["location"] = location })
}

// WithLogin login获取 git号
func (obj *_ZawazawaUserProfileGithubMgr) WithLogin(login string) Option {
	return optionFunc(func(o *options) { o.query["login"] = login })
}

// WithName name获取 git昵称
func (obj *_ZawazawaUserProfileGithubMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithNodeID node_id获取 节点id??
func (obj *_ZawazawaUserProfileGithubMgr) WithNodeID(nodeID string) Option {
	return optionFunc(func(o *options) { o.query["node_id"] = nodeID })
}

// WithOrganizationsURL organizations_url获取 不知道
func (obj *_ZawazawaUserProfileGithubMgr) WithOrganizationsURL(organizationsURL string) Option {
	return optionFunc(func(o *options) { o.query["organizations_url"] = organizationsURL })
}

// WithPublicGists public_gists获取 不知道
func (obj *_ZawazawaUserProfileGithubMgr) WithPublicGists(publicGists uint) Option {
	return optionFunc(func(o *options) { o.query["public_gists"] = publicGists })
}

// WithPublicRepos public_repos获取 开放的仓库数量
func (obj *_ZawazawaUserProfileGithubMgr) WithPublicRepos(publicRepos uint) Option {
	return optionFunc(func(o *options) { o.query["public_repos"] = publicRepos })
}

// WithReceivedEventsURL received_events_url获取 不知道
func (obj *_ZawazawaUserProfileGithubMgr) WithReceivedEventsURL(receivedEventsURL string) Option {
	return optionFunc(func(o *options) { o.query["received_events_url"] = receivedEventsURL })
}

// WithReposURL repos_url获取 不知道
func (obj *_ZawazawaUserProfileGithubMgr) WithReposURL(reposURL string) Option {
	return optionFunc(func(o *options) { o.query["repos_url"] = reposURL })
}

// WithStarredURL starred_url获取 不知道
func (obj *_ZawazawaUserProfileGithubMgr) WithStarredURL(starredURL string) Option {
	return optionFunc(func(o *options) { o.query["starred_url"] = starredURL })
}

// WithSubscriptionsURL subscriptions_url获取 仓库列表
func (obj *_ZawazawaUserProfileGithubMgr) WithSubscriptionsURL(subscriptionsURL string) Option {
	return optionFunc(func(o *options) { o.query["subscriptions_url"] = subscriptionsURL })
}

// WithTwitterUsername twitter_username获取 推特用户名?
func (obj *_ZawazawaUserProfileGithubMgr) WithTwitterUsername(twitterUsername string) Option {
	return optionFunc(func(o *options) { o.query["twitter_username"] = twitterUsername })
}

// WithType type获取 不知道是什么类型
func (obj *_ZawazawaUserProfileGithubMgr) WithType(_type string) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithURL url获取 个人主页地址
func (obj *_ZawazawaUserProfileGithubMgr) WithURL(url string) Option {
	return optionFunc(func(o *options) { o.query["url"] = url })
}

// GetByOption 功能选项模式获取
func (obj *_ZawazawaUserProfileGithubMgr) GetByOption(opts ...Option) (result ZawazawaUserProfileGithub, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_ZawazawaUserProfileGithubMgr) GetByOptions(opts ...Option) (results []*ZawazawaUserProfileGithub, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromGitOauthID 通过git_oauth_id获取内容
func (obj *_ZawazawaUserProfileGithubMgr) GetFromGitOauthID(gitOauthID uint) (result ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("git_oauth_id = ?", gitOauthID).Find(&result).Error

	return
}

// GetBatchFromGitOauthID 批量唯一主键查找
func (obj *_ZawazawaUserProfileGithubMgr) GetBatchFromGitOauthID(gitOauthIDs []uint) (results []*ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("git_oauth_id IN (?)", gitOauthIDs).Find(&results).Error

	return
}

// GetFromAvatarURL 通过avatar_url获取内容 头像地址
func (obj *_ZawazawaUserProfileGithubMgr) GetFromAvatarURL(avatarURL string) (results []*ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("avatar_url = ?", avatarURL).Find(&results).Error

	return
}

// GetBatchFromAvatarURL 批量唯一主键查找 头像地址
func (obj *_ZawazawaUserProfileGithubMgr) GetBatchFromAvatarURL(avatarURLs []string) (results []*ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("avatar_url IN (?)", avatarURLs).Find(&results).Error

	return
}

// GetFromBio 通过bio获取内容 个性签名
func (obj *_ZawazawaUserProfileGithubMgr) GetFromBio(bio string) (results []*ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("bio = ?", bio).Find(&results).Error

	return
}

// GetBatchFromBio 批量唯一主键查找 个性签名
func (obj *_ZawazawaUserProfileGithubMgr) GetBatchFromBio(bios []string) (results []*ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("bio IN (?)", bios).Find(&results).Error

	return
}

// GetFromBlog 通过blog获取内容 博客地址好像是
func (obj *_ZawazawaUserProfileGithubMgr) GetFromBlog(blog string) (results []*ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("blog = ?", blog).Find(&results).Error

	return
}

// GetBatchFromBlog 批量唯一主键查找 博客地址好像是
func (obj *_ZawazawaUserProfileGithubMgr) GetBatchFromBlog(blogs []string) (results []*ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("blog IN (?)", blogs).Find(&results).Error

	return
}

// GetFromCompany 通过company获取内容 公司
func (obj *_ZawazawaUserProfileGithubMgr) GetFromCompany(company string) (results []*ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("company = ?", company).Find(&results).Error

	return
}

// GetBatchFromCompany 批量唯一主键查找 公司
func (obj *_ZawazawaUserProfileGithubMgr) GetBatchFromCompany(companys []string) (results []*ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("company IN (?)", companys).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建日期
func (obj *_ZawazawaUserProfileGithubMgr) GetFromCreatedAt(createdAt string) (results []*ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量唯一主键查找 创建日期
func (obj *_ZawazawaUserProfileGithubMgr) GetBatchFromCreatedAt(createdAts []string) (results []*ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromEmail 通过email获取内容 邮箱地址
func (obj *_ZawazawaUserProfileGithubMgr) GetFromEmail(email string) (results []*ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email = ?", email).Find(&results).Error

	return
}

// GetBatchFromEmail 批量唯一主键查找 邮箱地址
func (obj *_ZawazawaUserProfileGithubMgr) GetBatchFromEmail(emails []string) (results []*ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email IN (?)", emails).Find(&results).Error

	return
}

// GetFromEventsURL 通过events_url获取内容 不知道
func (obj *_ZawazawaUserProfileGithubMgr) GetFromEventsURL(eventsURL string) (results []*ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("events_url = ?", eventsURL).Find(&results).Error

	return
}

// GetBatchFromEventsURL 批量唯一主键查找 不知道
func (obj *_ZawazawaUserProfileGithubMgr) GetBatchFromEventsURL(eventsURLs []string) (results []*ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("events_url IN (?)", eventsURLs).Find(&results).Error

	return
}

// GetFromFollowers 通过followers获取内容 粉丝数量
func (obj *_ZawazawaUserProfileGithubMgr) GetFromFollowers(followers uint) (results []*ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("followers = ?", followers).Find(&results).Error

	return
}

// GetBatchFromFollowers 批量唯一主键查找 粉丝数量
func (obj *_ZawazawaUserProfileGithubMgr) GetBatchFromFollowers(followerss []uint) (results []*ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("followers IN (?)", followerss).Find(&results).Error

	return
}

// GetFromFollowersURL 通过followers_url获取内容 粉丝列表地址
func (obj *_ZawazawaUserProfileGithubMgr) GetFromFollowersURL(followersURL string) (results []*ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("followers_url = ?", followersURL).Find(&results).Error

	return
}

// GetBatchFromFollowersURL 批量唯一主键查找 粉丝列表地址
func (obj *_ZawazawaUserProfileGithubMgr) GetBatchFromFollowersURL(followersURLs []string) (results []*ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("followers_url IN (?)", followersURLs).Find(&results).Error

	return
}

// GetFromFollowing 通过following获取内容 关注用户
func (obj *_ZawazawaUserProfileGithubMgr) GetFromFollowing(following uint) (results []*ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("following = ?", following).Find(&results).Error

	return
}

// GetBatchFromFollowing 批量唯一主键查找 关注用户
func (obj *_ZawazawaUserProfileGithubMgr) GetBatchFromFollowing(followings []uint) (results []*ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("following IN (?)", followings).Find(&results).Error

	return
}

// GetFromFollowingURL 通过following_url获取内容 关注用户列表地址
func (obj *_ZawazawaUserProfileGithubMgr) GetFromFollowingURL(followingURL string) (results []*ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("following_url = ?", followingURL).Find(&results).Error

	return
}

// GetBatchFromFollowingURL 批量唯一主键查找 关注用户列表地址
func (obj *_ZawazawaUserProfileGithubMgr) GetBatchFromFollowingURL(followingURLs []string) (results []*ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("following_url IN (?)", followingURLs).Find(&results).Error

	return
}

// GetFromGistsURL 通过gists_url获取内容 不知道是什么
func (obj *_ZawazawaUserProfileGithubMgr) GetFromGistsURL(gistsURL string) (results []*ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("gists_url = ?", gistsURL).Find(&results).Error

	return
}

// GetBatchFromGistsURL 批量唯一主键查找 不知道是什么
func (obj *_ZawazawaUserProfileGithubMgr) GetBatchFromGistsURL(gistsURLs []string) (results []*ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("gists_url IN (?)", gistsURLs).Find(&results).Error

	return
}

// GetFromHireable 通过hireable获取内容 不知道是什么
func (obj *_ZawazawaUserProfileGithubMgr) GetFromHireable(hireable string) (results []*ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("hireable = ?", hireable).Find(&results).Error

	return
}

// GetBatchFromHireable 批量唯一主键查找 不知道是什么
func (obj *_ZawazawaUserProfileGithubMgr) GetBatchFromHireable(hireables []string) (results []*ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("hireable IN (?)", hireables).Find(&results).Error

	return
}

// GetFromHTMLURL 通过html_url获取内容 主页地址
func (obj *_ZawazawaUserProfileGithubMgr) GetFromHTMLURL(htmlURL string) (results []*ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("html_url = ?", htmlURL).Find(&results).Error

	return
}

// GetBatchFromHTMLURL 批量唯一主键查找 主页地址
func (obj *_ZawazawaUserProfileGithubMgr) GetBatchFromHTMLURL(htmlURLs []string) (results []*ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("html_url IN (?)", htmlURLs).Find(&results).Error

	return
}

// GetFromID 通过id获取内容 github用户id
func (obj *_ZawazawaUserProfileGithubMgr) GetFromID(id uint) (result ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找 github用户id
func (obj *_ZawazawaUserProfileGithubMgr) GetBatchFromID(ids []uint) (results []*ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromLocation 通过location获取内容 定位??
func (obj *_ZawazawaUserProfileGithubMgr) GetFromLocation(location string) (results []*ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("location = ?", location).Find(&results).Error

	return
}

// GetBatchFromLocation 批量唯一主键查找 定位??
func (obj *_ZawazawaUserProfileGithubMgr) GetBatchFromLocation(locations []string) (results []*ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("location IN (?)", locations).Find(&results).Error

	return
}

// GetFromLogin 通过login获取内容 git号
func (obj *_ZawazawaUserProfileGithubMgr) GetFromLogin(login string) (results []*ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("login = ?", login).Find(&results).Error

	return
}

// GetBatchFromLogin 批量唯一主键查找 git号
func (obj *_ZawazawaUserProfileGithubMgr) GetBatchFromLogin(logins []string) (results []*ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("login IN (?)", logins).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 git昵称
func (obj *_ZawazawaUserProfileGithubMgr) GetFromName(name string) (results []*ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找 git昵称
func (obj *_ZawazawaUserProfileGithubMgr) GetBatchFromName(names []string) (results []*ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromNodeID 通过node_id获取内容 节点id??
func (obj *_ZawazawaUserProfileGithubMgr) GetFromNodeID(nodeID string) (results []*ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("node_id = ?", nodeID).Find(&results).Error

	return
}

// GetBatchFromNodeID 批量唯一主键查找 节点id??
func (obj *_ZawazawaUserProfileGithubMgr) GetBatchFromNodeID(nodeIDs []string) (results []*ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("node_id IN (?)", nodeIDs).Find(&results).Error

	return
}

// GetFromOrganizationsURL 通过organizations_url获取内容 不知道
func (obj *_ZawazawaUserProfileGithubMgr) GetFromOrganizationsURL(organizationsURL string) (results []*ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("organizations_url = ?", organizationsURL).Find(&results).Error

	return
}

// GetBatchFromOrganizationsURL 批量唯一主键查找 不知道
func (obj *_ZawazawaUserProfileGithubMgr) GetBatchFromOrganizationsURL(organizationsURLs []string) (results []*ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("organizations_url IN (?)", organizationsURLs).Find(&results).Error

	return
}

// GetFromPublicGists 通过public_gists获取内容 不知道
func (obj *_ZawazawaUserProfileGithubMgr) GetFromPublicGists(publicGists uint) (results []*ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("public_gists = ?", publicGists).Find(&results).Error

	return
}

// GetBatchFromPublicGists 批量唯一主键查找 不知道
func (obj *_ZawazawaUserProfileGithubMgr) GetBatchFromPublicGists(publicGistss []uint) (results []*ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("public_gists IN (?)", publicGistss).Find(&results).Error

	return
}

// GetFromPublicRepos 通过public_repos获取内容 开放的仓库数量
func (obj *_ZawazawaUserProfileGithubMgr) GetFromPublicRepos(publicRepos uint) (results []*ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("public_repos = ?", publicRepos).Find(&results).Error

	return
}

// GetBatchFromPublicRepos 批量唯一主键查找 开放的仓库数量
func (obj *_ZawazawaUserProfileGithubMgr) GetBatchFromPublicRepos(publicReposs []uint) (results []*ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("public_repos IN (?)", publicReposs).Find(&results).Error

	return
}

// GetFromReceivedEventsURL 通过received_events_url获取内容 不知道
func (obj *_ZawazawaUserProfileGithubMgr) GetFromReceivedEventsURL(receivedEventsURL string) (results []*ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("received_events_url = ?", receivedEventsURL).Find(&results).Error

	return
}

// GetBatchFromReceivedEventsURL 批量唯一主键查找 不知道
func (obj *_ZawazawaUserProfileGithubMgr) GetBatchFromReceivedEventsURL(receivedEventsURLs []string) (results []*ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("received_events_url IN (?)", receivedEventsURLs).Find(&results).Error

	return
}

// GetFromReposURL 通过repos_url获取内容 不知道
func (obj *_ZawazawaUserProfileGithubMgr) GetFromReposURL(reposURL string) (results []*ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("repos_url = ?", reposURL).Find(&results).Error

	return
}

// GetBatchFromReposURL 批量唯一主键查找 不知道
func (obj *_ZawazawaUserProfileGithubMgr) GetBatchFromReposURL(reposURLs []string) (results []*ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("repos_url IN (?)", reposURLs).Find(&results).Error

	return
}

// GetFromStarredURL 通过starred_url获取内容 不知道
func (obj *_ZawazawaUserProfileGithubMgr) GetFromStarredURL(starredURL string) (results []*ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("starred_url = ?", starredURL).Find(&results).Error

	return
}

// GetBatchFromStarredURL 批量唯一主键查找 不知道
func (obj *_ZawazawaUserProfileGithubMgr) GetBatchFromStarredURL(starredURLs []string) (results []*ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("starred_url IN (?)", starredURLs).Find(&results).Error

	return
}

// GetFromSubscriptionsURL 通过subscriptions_url获取内容 仓库列表
func (obj *_ZawazawaUserProfileGithubMgr) GetFromSubscriptionsURL(subscriptionsURL string) (results []*ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("subscriptions_url = ?", subscriptionsURL).Find(&results).Error

	return
}

// GetBatchFromSubscriptionsURL 批量唯一主键查找 仓库列表
func (obj *_ZawazawaUserProfileGithubMgr) GetBatchFromSubscriptionsURL(subscriptionsURLs []string) (results []*ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("subscriptions_url IN (?)", subscriptionsURLs).Find(&results).Error

	return
}

// GetFromTwitterUsername 通过twitter_username获取内容 推特用户名?
func (obj *_ZawazawaUserProfileGithubMgr) GetFromTwitterUsername(twitterUsername string) (results []*ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("twitter_username = ?", twitterUsername).Find(&results).Error

	return
}

// GetBatchFromTwitterUsername 批量唯一主键查找 推特用户名?
func (obj *_ZawazawaUserProfileGithubMgr) GetBatchFromTwitterUsername(twitterUsernames []string) (results []*ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("twitter_username IN (?)", twitterUsernames).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 不知道是什么类型
func (obj *_ZawazawaUserProfileGithubMgr) GetFromType(_type string) (results []*ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量唯一主键查找 不知道是什么类型
func (obj *_ZawazawaUserProfileGithubMgr) GetBatchFromType(_types []string) (results []*ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type IN (?)", _types).Find(&results).Error

	return
}

// GetFromURL 通过url获取内容 个人主页地址
func (obj *_ZawazawaUserProfileGithubMgr) GetFromURL(url string) (results []*ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("url = ?", url).Find(&results).Error

	return
}

// GetBatchFromURL 批量唯一主键查找 个人主页地址
func (obj *_ZawazawaUserProfileGithubMgr) GetBatchFromURL(urls []string) (results []*ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("url IN (?)", urls).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ZawazawaUserProfileGithubMgr) FetchByPrimaryKey(gitOauthID uint) (result ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("git_oauth_id = ?", gitOauthID).Find(&result).Error

	return
}

// FetchUniqueByGitidUnique primay or index 获取唯一内容
func (obj *_ZawazawaUserProfileGithubMgr) FetchUniqueByGitidUnique(id uint) (result ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchUniqueByGitID primay or index 获取唯一内容
func (obj *_ZawazawaUserProfileGithubMgr) FetchUniqueByGitID(id uint) (result ZawazawaUserProfileGithub, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

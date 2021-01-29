package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _ZawazawaUserMgr struct {
	*_BaseMgr
}

// ZawazawaUserMgr open func
func ZawazawaUserMgr(db *gorm.DB) *_ZawazawaUserMgr {
	if db == nil {
		panic(fmt.Errorf("ZawazawaUserMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ZawazawaUserMgr{_BaseMgr: &_BaseMgr{DB: db.Table("zawazawa_user"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ZawazawaUserMgr) GetTableName() string {
	return "zawazawa_user"
}

// Get 获取
func (obj *_ZawazawaUserMgr) Get() (result ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ZawazawaUserMgr) Gets() (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_ZawazawaUserMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithAvatarURL avatar_url获取 头像地址
func (obj *_ZawazawaUserMgr) WithAvatarURL(avatarURL string) Option {
	return optionFunc(func(o *options) { o.query["avatar_url"] = avatarURL })
}

// WithBio bio获取 个性签名
func (obj *_ZawazawaUserMgr) WithBio(bio string) Option {
	return optionFunc(func(o *options) { o.query["bio"] = bio })
}

// WithBlog blog获取 博客地址好像是
func (obj *_ZawazawaUserMgr) WithBlog(blog string) Option {
	return optionFunc(func(o *options) { o.query["blog"] = blog })
}

// WithCompany company获取 公司
func (obj *_ZawazawaUserMgr) WithCompany(company string) Option {
	return optionFunc(func(o *options) { o.query["company"] = company })
}

// WithCreatedAt created_at获取 创建日期
func (obj *_ZawazawaUserMgr) WithCreatedAt(createdAt string) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithEmail email获取 邮箱地址
func (obj *_ZawazawaUserMgr) WithEmail(email string) Option {
	return optionFunc(func(o *options) { o.query["email"] = email })
}

// WithEventsURL events_url获取 不知道
func (obj *_ZawazawaUserMgr) WithEventsURL(eventsURL string) Option {
	return optionFunc(func(o *options) { o.query["events_url"] = eventsURL })
}

// WithFollowers followers获取 粉丝数量
func (obj *_ZawazawaUserMgr) WithFollowers(followers int) Option {
	return optionFunc(func(o *options) { o.query["followers"] = followers })
}

// WithFollowersURL followers_url获取 粉丝列表地址
func (obj *_ZawazawaUserMgr) WithFollowersURL(followersURL string) Option {
	return optionFunc(func(o *options) { o.query["followers_url"] = followersURL })
}

// WithFollowing following获取 关注用户
func (obj *_ZawazawaUserMgr) WithFollowing(following int) Option {
	return optionFunc(func(o *options) { o.query["following"] = following })
}

// WithFollowingURL following_url获取 关注用户列表地址
func (obj *_ZawazawaUserMgr) WithFollowingURL(followingURL string) Option {
	return optionFunc(func(o *options) { o.query["following_url"] = followingURL })
}

// WithGistsURL gists_url获取 不知道是什么
func (obj *_ZawazawaUserMgr) WithGistsURL(gistsURL string) Option {
	return optionFunc(func(o *options) { o.query["gists_url"] = gistsURL })
}

// WithGravatarID gravatar_id获取 不知道是什么
func (obj *_ZawazawaUserMgr) WithGravatarID(gravatarID int) Option {
	return optionFunc(func(o *options) { o.query["gravatar_id"] = gravatarID })
}

// WithHireable hireable获取 不知道是什么
func (obj *_ZawazawaUserMgr) WithHireable(hireable string) Option {
	return optionFunc(func(o *options) { o.query["hireable"] = hireable })
}

// WithHTMLURL html_url获取 主页地址
func (obj *_ZawazawaUserMgr) WithHTMLURL(htmlURL string) Option {
	return optionFunc(func(o *options) { o.query["html_url"] = htmlURL })
}

// WithGitID git_id获取 github用户id
func (obj *_ZawazawaUserMgr) WithGitID(gitID int) Option {
	return optionFunc(func(o *options) { o.query["git_id"] = gitID })
}

// WithLocation location获取 定位??
func (obj *_ZawazawaUserMgr) WithLocation(location string) Option {
	return optionFunc(func(o *options) { o.query["location"] = location })
}

// WithLogin login获取 git号
func (obj *_ZawazawaUserMgr) WithLogin(login string) Option {
	return optionFunc(func(o *options) { o.query["login"] = login })
}

// WithName name获取 git昵称
func (obj *_ZawazawaUserMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithNodeID node_id获取 节点id??
func (obj *_ZawazawaUserMgr) WithNodeID(nodeID int) Option {
	return optionFunc(func(o *options) { o.query["node_id"] = nodeID })
}

// WithOrganizationsURL organizations_url获取 不知道
func (obj *_ZawazawaUserMgr) WithOrganizationsURL(organizationsURL string) Option {
	return optionFunc(func(o *options) { o.query["organizations_url"] = organizationsURL })
}

// WithPublicGists public_gists获取 不知道
func (obj *_ZawazawaUserMgr) WithPublicGists(publicGists string) Option {
	return optionFunc(func(o *options) { o.query["public_gists"] = publicGists })
}

// WithPublicRepos public_repos获取 开放的仓库数量
func (obj *_ZawazawaUserMgr) WithPublicRepos(publicRepos int) Option {
	return optionFunc(func(o *options) { o.query["public_repos"] = publicRepos })
}

// WithReceivedEventsURL received_events_url获取 不知道
func (obj *_ZawazawaUserMgr) WithReceivedEventsURL(receivedEventsURL string) Option {
	return optionFunc(func(o *options) { o.query["received_events_url"] = receivedEventsURL })
}

// WithReposURL repos_url获取 不知道
func (obj *_ZawazawaUserMgr) WithReposURL(reposURL string) Option {
	return optionFunc(func(o *options) { o.query["repos_url"] = reposURL })
}

// WithSiteAdmin site_admin获取 网站管理员??
func (obj *_ZawazawaUserMgr) WithSiteAdmin(siteAdmin bool) Option {
	return optionFunc(func(o *options) { o.query["site_admin"] = siteAdmin })
}

// WithStarredURL starred_url获取 不知道
func (obj *_ZawazawaUserMgr) WithStarredURL(starredURL string) Option {
	return optionFunc(func(o *options) { o.query["starred_url"] = starredURL })
}

// WithSubscriptionsURL subscriptions_url获取 仓库列表
func (obj *_ZawazawaUserMgr) WithSubscriptionsURL(subscriptionsURL string) Option {
	return optionFunc(func(o *options) { o.query["subscriptions_url"] = subscriptionsURL })
}

// WithTwitterUsername twitter_username获取 推特用户名?
func (obj *_ZawazawaUserMgr) WithTwitterUsername(twitterUsername string) Option {
	return optionFunc(func(o *options) { o.query["twitter_username"] = twitterUsername })
}

// WithType type获取 不知道是什么类型
func (obj *_ZawazawaUserMgr) WithType(_type string) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithUpdatedAt updated_at获取 更新时间
func (obj *_ZawazawaUserMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithURL url获取 个人主页地址
func (obj *_ZawazawaUserMgr) WithURL(url string) Option {
	return optionFunc(func(o *options) { o.query["url"] = url })
}

// GetByOption 功能选项模式获取
func (obj *_ZawazawaUserMgr) GetByOption(opts ...Option) (result ZawazawaUser, err error) {
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
func (obj *_ZawazawaUserMgr) GetByOptions(opts ...Option) (results []*ZawazawaUser, err error) {
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

// GetFromID 通过id获取内容
func (obj *_ZawazawaUserMgr) GetFromID(id int) (result ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_ZawazawaUserMgr) GetBatchFromID(ids []int) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromAvatarURL 通过avatar_url获取内容 头像地址
func (obj *_ZawazawaUserMgr) GetFromAvatarURL(avatarURL string) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("avatar_url = ?", avatarURL).Find(&results).Error

	return
}

// GetBatchFromAvatarURL 批量唯一主键查找 头像地址
func (obj *_ZawazawaUserMgr) GetBatchFromAvatarURL(avatarURLs []string) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("avatar_url IN (?)", avatarURLs).Find(&results).Error

	return
}

// GetFromBio 通过bio获取内容 个性签名
func (obj *_ZawazawaUserMgr) GetFromBio(bio string) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("bio = ?", bio).Find(&results).Error

	return
}

// GetBatchFromBio 批量唯一主键查找 个性签名
func (obj *_ZawazawaUserMgr) GetBatchFromBio(bios []string) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("bio IN (?)", bios).Find(&results).Error

	return
}

// GetFromBlog 通过blog获取内容 博客地址好像是
func (obj *_ZawazawaUserMgr) GetFromBlog(blog string) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("blog = ?", blog).Find(&results).Error

	return
}

// GetBatchFromBlog 批量唯一主键查找 博客地址好像是
func (obj *_ZawazawaUserMgr) GetBatchFromBlog(blogs []string) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("blog IN (?)", blogs).Find(&results).Error

	return
}

// GetFromCompany 通过company获取内容 公司
func (obj *_ZawazawaUserMgr) GetFromCompany(company string) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("company = ?", company).Find(&results).Error

	return
}

// GetBatchFromCompany 批量唯一主键查找 公司
func (obj *_ZawazawaUserMgr) GetBatchFromCompany(companys []string) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("company IN (?)", companys).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 创建日期
func (obj *_ZawazawaUserMgr) GetFromCreatedAt(createdAt string) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量唯一主键查找 创建日期
func (obj *_ZawazawaUserMgr) GetBatchFromCreatedAt(createdAts []string) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromEmail 通过email获取内容 邮箱地址
func (obj *_ZawazawaUserMgr) GetFromEmail(email string) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email = ?", email).Find(&results).Error

	return
}

// GetBatchFromEmail 批量唯一主键查找 邮箱地址
func (obj *_ZawazawaUserMgr) GetBatchFromEmail(emails []string) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("email IN (?)", emails).Find(&results).Error

	return
}

// GetFromEventsURL 通过events_url获取内容 不知道
func (obj *_ZawazawaUserMgr) GetFromEventsURL(eventsURL string) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("events_url = ?", eventsURL).Find(&results).Error

	return
}

// GetBatchFromEventsURL 批量唯一主键查找 不知道
func (obj *_ZawazawaUserMgr) GetBatchFromEventsURL(eventsURLs []string) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("events_url IN (?)", eventsURLs).Find(&results).Error

	return
}

// GetFromFollowers 通过followers获取内容 粉丝数量
func (obj *_ZawazawaUserMgr) GetFromFollowers(followers int) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("followers = ?", followers).Find(&results).Error

	return
}

// GetBatchFromFollowers 批量唯一主键查找 粉丝数量
func (obj *_ZawazawaUserMgr) GetBatchFromFollowers(followerss []int) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("followers IN (?)", followerss).Find(&results).Error

	return
}

// GetFromFollowersURL 通过followers_url获取内容 粉丝列表地址
func (obj *_ZawazawaUserMgr) GetFromFollowersURL(followersURL string) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("followers_url = ?", followersURL).Find(&results).Error

	return
}

// GetBatchFromFollowersURL 批量唯一主键查找 粉丝列表地址
func (obj *_ZawazawaUserMgr) GetBatchFromFollowersURL(followersURLs []string) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("followers_url IN (?)", followersURLs).Find(&results).Error

	return
}

// GetFromFollowing 通过following获取内容 关注用户
func (obj *_ZawazawaUserMgr) GetFromFollowing(following int) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("following = ?", following).Find(&results).Error

	return
}

// GetBatchFromFollowing 批量唯一主键查找 关注用户
func (obj *_ZawazawaUserMgr) GetBatchFromFollowing(followings []int) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("following IN (?)", followings).Find(&results).Error

	return
}

// GetFromFollowingURL 通过following_url获取内容 关注用户列表地址
func (obj *_ZawazawaUserMgr) GetFromFollowingURL(followingURL string) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("following_url = ?", followingURL).Find(&results).Error

	return
}

// GetBatchFromFollowingURL 批量唯一主键查找 关注用户列表地址
func (obj *_ZawazawaUserMgr) GetBatchFromFollowingURL(followingURLs []string) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("following_url IN (?)", followingURLs).Find(&results).Error

	return
}

// GetFromGistsURL 通过gists_url获取内容 不知道是什么
func (obj *_ZawazawaUserMgr) GetFromGistsURL(gistsURL string) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("gists_url = ?", gistsURL).Find(&results).Error

	return
}

// GetBatchFromGistsURL 批量唯一主键查找 不知道是什么
func (obj *_ZawazawaUserMgr) GetBatchFromGistsURL(gistsURLs []string) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("gists_url IN (?)", gistsURLs).Find(&results).Error

	return
}

// GetFromGravatarID 通过gravatar_id获取内容 不知道是什么
func (obj *_ZawazawaUserMgr) GetFromGravatarID(gravatarID int) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("gravatar_id = ?", gravatarID).Find(&results).Error

	return
}

// GetBatchFromGravatarID 批量唯一主键查找 不知道是什么
func (obj *_ZawazawaUserMgr) GetBatchFromGravatarID(gravatarIDs []int) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("gravatar_id IN (?)", gravatarIDs).Find(&results).Error

	return
}

// GetFromHireable 通过hireable获取内容 不知道是什么
func (obj *_ZawazawaUserMgr) GetFromHireable(hireable string) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("hireable = ?", hireable).Find(&results).Error

	return
}

// GetBatchFromHireable 批量唯一主键查找 不知道是什么
func (obj *_ZawazawaUserMgr) GetBatchFromHireable(hireables []string) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("hireable IN (?)", hireables).Find(&results).Error

	return
}

// GetFromHTMLURL 通过html_url获取内容 主页地址
func (obj *_ZawazawaUserMgr) GetFromHTMLURL(htmlURL string) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("html_url = ?", htmlURL).Find(&results).Error

	return
}

// GetBatchFromHTMLURL 批量唯一主键查找 主页地址
func (obj *_ZawazawaUserMgr) GetBatchFromHTMLURL(htmlURLs []string) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("html_url IN (?)", htmlURLs).Find(&results).Error

	return
}

// GetFromGitID 通过git_id获取内容 github用户id
func (obj *_ZawazawaUserMgr) GetFromGitID(gitID int) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("git_id = ?", gitID).Find(&results).Error

	return
}

// GetBatchFromGitID 批量唯一主键查找 github用户id
func (obj *_ZawazawaUserMgr) GetBatchFromGitID(gitIDs []int) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("git_id IN (?)", gitIDs).Find(&results).Error

	return
}

// GetFromLocation 通过location获取内容 定位??
func (obj *_ZawazawaUserMgr) GetFromLocation(location string) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("location = ?", location).Find(&results).Error

	return
}

// GetBatchFromLocation 批量唯一主键查找 定位??
func (obj *_ZawazawaUserMgr) GetBatchFromLocation(locations []string) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("location IN (?)", locations).Find(&results).Error

	return
}

// GetFromLogin 通过login获取内容 git号
func (obj *_ZawazawaUserMgr) GetFromLogin(login string) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("login = ?", login).Find(&results).Error

	return
}

// GetBatchFromLogin 批量唯一主键查找 git号
func (obj *_ZawazawaUserMgr) GetBatchFromLogin(logins []string) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("login IN (?)", logins).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 git昵称
func (obj *_ZawazawaUserMgr) GetFromName(name string) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量唯一主键查找 git昵称
func (obj *_ZawazawaUserMgr) GetBatchFromName(names []string) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromNodeID 通过node_id获取内容 节点id??
func (obj *_ZawazawaUserMgr) GetFromNodeID(nodeID int) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("node_id = ?", nodeID).Find(&results).Error

	return
}

// GetBatchFromNodeID 批量唯一主键查找 节点id??
func (obj *_ZawazawaUserMgr) GetBatchFromNodeID(nodeIDs []int) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("node_id IN (?)", nodeIDs).Find(&results).Error

	return
}

// GetFromOrganizationsURL 通过organizations_url获取内容 不知道
func (obj *_ZawazawaUserMgr) GetFromOrganizationsURL(organizationsURL string) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("organizations_url = ?", organizationsURL).Find(&results).Error

	return
}

// GetBatchFromOrganizationsURL 批量唯一主键查找 不知道
func (obj *_ZawazawaUserMgr) GetBatchFromOrganizationsURL(organizationsURLs []string) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("organizations_url IN (?)", organizationsURLs).Find(&results).Error

	return
}

// GetFromPublicGists 通过public_gists获取内容 不知道
func (obj *_ZawazawaUserMgr) GetFromPublicGists(publicGists string) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("public_gists = ?", publicGists).Find(&results).Error

	return
}

// GetBatchFromPublicGists 批量唯一主键查找 不知道
func (obj *_ZawazawaUserMgr) GetBatchFromPublicGists(publicGistss []string) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("public_gists IN (?)", publicGistss).Find(&results).Error

	return
}

// GetFromPublicRepos 通过public_repos获取内容 开放的仓库数量
func (obj *_ZawazawaUserMgr) GetFromPublicRepos(publicRepos int) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("public_repos = ?", publicRepos).Find(&results).Error

	return
}

// GetBatchFromPublicRepos 批量唯一主键查找 开放的仓库数量
func (obj *_ZawazawaUserMgr) GetBatchFromPublicRepos(publicReposs []int) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("public_repos IN (?)", publicReposs).Find(&results).Error

	return
}

// GetFromReceivedEventsURL 通过received_events_url获取内容 不知道
func (obj *_ZawazawaUserMgr) GetFromReceivedEventsURL(receivedEventsURL string) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("received_events_url = ?", receivedEventsURL).Find(&results).Error

	return
}

// GetBatchFromReceivedEventsURL 批量唯一主键查找 不知道
func (obj *_ZawazawaUserMgr) GetBatchFromReceivedEventsURL(receivedEventsURLs []string) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("received_events_url IN (?)", receivedEventsURLs).Find(&results).Error

	return
}

// GetFromReposURL 通过repos_url获取内容 不知道
func (obj *_ZawazawaUserMgr) GetFromReposURL(reposURL string) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("repos_url = ?", reposURL).Find(&results).Error

	return
}

// GetBatchFromReposURL 批量唯一主键查找 不知道
func (obj *_ZawazawaUserMgr) GetBatchFromReposURL(reposURLs []string) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("repos_url IN (?)", reposURLs).Find(&results).Error

	return
}

// GetFromSiteAdmin 通过site_admin获取内容 网站管理员??
func (obj *_ZawazawaUserMgr) GetFromSiteAdmin(siteAdmin bool) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("site_admin = ?", siteAdmin).Find(&results).Error

	return
}

// GetBatchFromSiteAdmin 批量唯一主键查找 网站管理员??
func (obj *_ZawazawaUserMgr) GetBatchFromSiteAdmin(siteAdmins []bool) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("site_admin IN (?)", siteAdmins).Find(&results).Error

	return
}

// GetFromStarredURL 通过starred_url获取内容 不知道
func (obj *_ZawazawaUserMgr) GetFromStarredURL(starredURL string) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("starred_url = ?", starredURL).Find(&results).Error

	return
}

// GetBatchFromStarredURL 批量唯一主键查找 不知道
func (obj *_ZawazawaUserMgr) GetBatchFromStarredURL(starredURLs []string) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("starred_url IN (?)", starredURLs).Find(&results).Error

	return
}

// GetFromSubscriptionsURL 通过subscriptions_url获取内容 仓库列表
func (obj *_ZawazawaUserMgr) GetFromSubscriptionsURL(subscriptionsURL string) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("subscriptions_url = ?", subscriptionsURL).Find(&results).Error

	return
}

// GetBatchFromSubscriptionsURL 批量唯一主键查找 仓库列表
func (obj *_ZawazawaUserMgr) GetBatchFromSubscriptionsURL(subscriptionsURLs []string) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("subscriptions_url IN (?)", subscriptionsURLs).Find(&results).Error

	return
}

// GetFromTwitterUsername 通过twitter_username获取内容 推特用户名?
func (obj *_ZawazawaUserMgr) GetFromTwitterUsername(twitterUsername string) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("twitter_username = ?", twitterUsername).Find(&results).Error

	return
}

// GetBatchFromTwitterUsername 批量唯一主键查找 推特用户名?
func (obj *_ZawazawaUserMgr) GetBatchFromTwitterUsername(twitterUsernames []string) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("twitter_username IN (?)", twitterUsernames).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 不知道是什么类型
func (obj *_ZawazawaUserMgr) GetFromType(_type string) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量唯一主键查找 不知道是什么类型
func (obj *_ZawazawaUserMgr) GetBatchFromType(_types []string) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type IN (?)", _types).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容 更新时间
func (obj *_ZawazawaUserMgr) GetFromUpdatedAt(updatedAt time.Time) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_at = ?", updatedAt).Find(&results).Error

	return
}

// GetBatchFromUpdatedAt 批量唯一主键查找 更新时间
func (obj *_ZawazawaUserMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_at IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromURL 通过url获取内容 个人主页地址
func (obj *_ZawazawaUserMgr) GetFromURL(url string) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("url = ?", url).Find(&results).Error

	return
}

// GetBatchFromURL 批量唯一主键查找 个人主页地址
func (obj *_ZawazawaUserMgr) GetBatchFromURL(urls []string) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("url IN (?)", urls).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ZawazawaUserMgr) FetchByPrimaryKey(id int) (result ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

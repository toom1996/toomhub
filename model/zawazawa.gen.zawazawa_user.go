package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
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
func (obj *_ZawazawaUserMgr) WithID(id uint) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithNickname nickname获取 昵称
func (obj *_ZawazawaUserMgr) WithNickname(nickname string) Option {
	return optionFunc(func(o *options) { o.query["nickname"] = nickname })
}

// WithMobile mobile获取 手机号
func (obj *_ZawazawaUserMgr) WithMobile(mobile string) Option {
	return optionFunc(func(o *options) { o.query["mobile"] = mobile })
}

// WithFansCount fans_count获取 粉丝数量
func (obj *_ZawazawaUserMgr) WithFansCount(fansCount uint) Option {
	return optionFunc(func(o *options) { o.query["fans_count"] = fansCount })
}

// WithFollowCount follow_count获取 关注数量
func (obj *_ZawazawaUserMgr) WithFollowCount(followCount uint) Option {
	return optionFunc(func(o *options) { o.query["follow_count"] = followCount })
}

// WithCollectCount collect_count获取 收藏数量
func (obj *_ZawazawaUserMgr) WithCollectCount(collectCount uint) Option {
	return optionFunc(func(o *options) { o.query["collect_count"] = collectCount })
}

// WithOauthID oauth_id获取 授权id
func (obj *_ZawazawaUserMgr) WithOauthID(oauthID uint) Option {
	return optionFunc(func(o *options) { o.query["oauth_id"] = oauthID })
}

// WithOauthType oauth_type获取 授权登陆类型 0=>未授权 1=> github
func (obj *_ZawazawaUserMgr) WithOauthType(oauthType uint16) Option {
	return optionFunc(func(o *options) { o.query["oauth_type"] = oauthType })
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
func (obj *_ZawazawaUserMgr) GetFromID(id uint) (result ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_ZawazawaUserMgr) GetBatchFromID(ids []uint) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromNickname 通过nickname获取内容 昵称
func (obj *_ZawazawaUserMgr) GetFromNickname(nickname string) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("nickname = ?", nickname).Find(&results).Error

	return
}

// GetBatchFromNickname 批量唯一主键查找 昵称
func (obj *_ZawazawaUserMgr) GetBatchFromNickname(nicknames []string) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("nickname IN (?)", nicknames).Find(&results).Error

	return
}

// GetFromMobile 通过mobile获取内容 手机号
func (obj *_ZawazawaUserMgr) GetFromMobile(mobile string) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("mobile = ?", mobile).Find(&results).Error

	return
}

// GetBatchFromMobile 批量唯一主键查找 手机号
func (obj *_ZawazawaUserMgr) GetBatchFromMobile(mobiles []string) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("mobile IN (?)", mobiles).Find(&results).Error

	return
}

// GetFromFansCount 通过fans_count获取内容 粉丝数量
func (obj *_ZawazawaUserMgr) GetFromFansCount(fansCount uint) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("fans_count = ?", fansCount).Find(&results).Error

	return
}

// GetBatchFromFansCount 批量唯一主键查找 粉丝数量
func (obj *_ZawazawaUserMgr) GetBatchFromFansCount(fansCounts []uint) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("fans_count IN (?)", fansCounts).Find(&results).Error

	return
}

// GetFromFollowCount 通过follow_count获取内容 关注数量
func (obj *_ZawazawaUserMgr) GetFromFollowCount(followCount uint) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("follow_count = ?", followCount).Find(&results).Error

	return
}

// GetBatchFromFollowCount 批量唯一主键查找 关注数量
func (obj *_ZawazawaUserMgr) GetBatchFromFollowCount(followCounts []uint) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("follow_count IN (?)", followCounts).Find(&results).Error

	return
}

// GetFromCollectCount 通过collect_count获取内容 收藏数量
func (obj *_ZawazawaUserMgr) GetFromCollectCount(collectCount uint) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("collect_count = ?", collectCount).Find(&results).Error

	return
}

// GetBatchFromCollectCount 批量唯一主键查找 收藏数量
func (obj *_ZawazawaUserMgr) GetBatchFromCollectCount(collectCounts []uint) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("collect_count IN (?)", collectCounts).Find(&results).Error

	return
}

// GetFromOauthID 通过oauth_id获取内容 授权id
func (obj *_ZawazawaUserMgr) GetFromOauthID(oauthID uint) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("oauth_id = ?", oauthID).Find(&results).Error

	return
}

// GetBatchFromOauthID 批量唯一主键查找 授权id
func (obj *_ZawazawaUserMgr) GetBatchFromOauthID(oauthIDs []uint) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("oauth_id IN (?)", oauthIDs).Find(&results).Error

	return
}

// GetFromOauthType 通过oauth_type获取内容 授权登陆类型 0=>未授权 1=> github
func (obj *_ZawazawaUserMgr) GetFromOauthType(oauthType uint16) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("oauth_type = ?", oauthType).Find(&results).Error

	return
}

// GetBatchFromOauthType 批量唯一主键查找 授权登陆类型 0=>未授权 1=> github
func (obj *_ZawazawaUserMgr) GetBatchFromOauthType(oauthTypes []uint16) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("oauth_type IN (?)", oauthTypes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ZawazawaUserMgr) FetchByPrimaryKey(id uint) (result ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

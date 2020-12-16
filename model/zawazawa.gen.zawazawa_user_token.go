package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _ZawazawaUserTokenMgr struct {
	*_BaseMgr
}

// ZawazawaUserTokenMgr open func
func ZawazawaUserTokenMgr(db *gorm.DB) *_ZawazawaUserTokenMgr {
	if db == nil {
		panic(fmt.Errorf("ZawazawaUserTokenMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ZawazawaUserTokenMgr{_BaseMgr: &_BaseMgr{DB: db.Table("zawazawa_user_token"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ZawazawaUserTokenMgr) GetTableName() string {
	return "zawazawa_user_token"
}

// Get 获取
func (obj *_ZawazawaUserTokenMgr) Get() (result ZawazawaUserToken, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ZawazawaUserTokenMgr) Gets() (results []*ZawazawaUserToken, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_ZawazawaUserTokenMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUId uid获取
func (obj *_ZawazawaUserTokenMgr) WithUId(uid int) Option {
	return optionFunc(func(o *options) { o.query["uid"] = uid })
}

// WithToken token获取
func (obj *_ZawazawaUserTokenMgr) WithToken(token string) Option {
	return optionFunc(func(o *options) { o.query["token"] = token })
}

// WithRefreshToken refresh_token获取
func (obj *_ZawazawaUserTokenMgr) WithRefreshToken(refreshToken string) Option {
	return optionFunc(func(o *options) { o.query["refresh_token"] = refreshToken })
}

// WithType type获取
func (obj *_ZawazawaUserTokenMgr) WithType(_type string) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// GetByOption 功能选项模式获取
func (obj *_ZawazawaUserTokenMgr) GetByOption(opts ...Option) (result ZawazawaUserToken, err error) {
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
func (obj *_ZawazawaUserTokenMgr) GetByOptions(opts ...Option) (results []*ZawazawaUserToken, err error) {
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
func (obj *_ZawazawaUserTokenMgr) GetFromID(id int) (result ZawazawaUserToken, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_ZawazawaUserTokenMgr) GetBatchFromID(ids []int) (results []*ZawazawaUserToken, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromUId 通过uid获取内容
func (obj *_ZawazawaUserTokenMgr) GetFromUId(uid int) (results []*ZawazawaUserToken, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("uid = ?", uid).Find(&results).Error

	return
}

// GetBatchFromUId 批量唯一主键查找
func (obj *_ZawazawaUserTokenMgr) GetBatchFromUId(uids []int) (results []*ZawazawaUserToken, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("uid IN (?)", uids).Find(&results).Error

	return
}

// GetFromToken 通过token获取内容
func (obj *_ZawazawaUserTokenMgr) GetFromToken(token string) (results []*ZawazawaUserToken, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("token = ?", token).Find(&results).Error

	return
}

// GetBatchFromToken 批量唯一主键查找
func (obj *_ZawazawaUserTokenMgr) GetBatchFromToken(tokens []string) (results []*ZawazawaUserToken, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("token IN (?)", tokens).Find(&results).Error

	return
}

// GetFromRefreshToken 通过refresh_token获取内容
func (obj *_ZawazawaUserTokenMgr) GetFromRefreshToken(refreshToken string) (results []*ZawazawaUserToken, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("refresh_token = ?", refreshToken).Find(&results).Error

	return
}

// GetBatchFromRefreshToken 批量唯一主键查找
func (obj *_ZawazawaUserTokenMgr) GetBatchFromRefreshToken(refreshTokens []string) (results []*ZawazawaUserToken, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("refresh_token IN (?)", refreshTokens).Find(&results).Error

	return
}

// GetFromType 通过type获取内容
func (obj *_ZawazawaUserTokenMgr) GetFromType(_type string) (results []*ZawazawaUserToken, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量唯一主键查找
func (obj *_ZawazawaUserTokenMgr) GetBatchFromType(_types []string) (results []*ZawazawaUserToken, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type IN (?)", _types).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ZawazawaUserTokenMgr) FetchByPrimaryKey(id int) (result ZawazawaUserToken, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

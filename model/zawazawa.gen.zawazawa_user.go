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
func (obj *_ZawazawaUserMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithMobile mobile获取 手机号码
func (obj *_ZawazawaUserMgr) WithMobile(mobile string) Option {
	return optionFunc(func(o *options) { o.query["mobile"] = mobile })
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

// GetFromMobile 通过mobile获取内容 手机号码
func (obj *_ZawazawaUserMgr) GetFromMobile(mobile string) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("mobile = ?", mobile).Find(&results).Error

	return
}

// GetBatchFromMobile 批量唯一主键查找 手机号码
func (obj *_ZawazawaUserMgr) GetBatchFromMobile(mobiles []string) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("mobile IN (?)", mobiles).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ZawazawaUserMgr) FetchByPrimaryKey(id int) (result ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchIndexByMobile  获取多个内容
func (obj *_ZawazawaUserMgr) FetchIndexByMobile(mobile string) (results []*ZawazawaUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("mobile = ?", mobile).Find(&results).Error

	return
}

////GetUser	获取用户信息
//func (obj *_ZawazawaUserMgr) GetUser() {
//	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("mobile = ?", mobile).Find(&results).Error
//
//	return
//}

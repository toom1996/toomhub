package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _ZawazawaPostMgr struct {
	*_BaseMgr
}

// ZawazawaPostMgr open func
func ZawazawaPostMgr(db *gorm.DB) *_ZawazawaPostMgr {
	if db == nil {
		panic(fmt.Errorf("ZawazawaPostMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ZawazawaPostMgr{_BaseMgr: &_BaseMgr{DB: db.Table("zawazawa_post"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ZawazawaPostMgr) GetTableName() string {
	return "zawazawa_post"
}

// Get 获取
func (obj *_ZawazawaPostMgr) Get() (result ZawazawaPost, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ZawazawaPostMgr) Gets() (results []*ZawazawaPost, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_ZawazawaPostMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithContent content获取
func (obj *_ZawazawaPostMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}

// WithCreatedAt created_at获取
func (obj *_ZawazawaPostMgr) WithCreatedAt(createdAt int) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithCreatedBy created_by获取
func (obj *_ZawazawaPostMgr) WithCreatedBy(createdBy int) Option {
	return optionFunc(func(o *options) { o.query["created_by"] = createdBy })
}

// GetByOption 功能选项模式获取
func (obj *_ZawazawaPostMgr) GetByOption(opts ...Option) (result ZawazawaPost, err error) {
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
func (obj *_ZawazawaPostMgr) GetByOptions(opts ...Option) (results []*ZawazawaPost, err error) {
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
func (obj *_ZawazawaPostMgr) GetFromID(id int) (result ZawazawaPost, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量唯一主键查找
func (obj *_ZawazawaPostMgr) GetBatchFromID(ids []int) (results []*ZawazawaPost, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromContent 通过content获取内容
func (obj *_ZawazawaPostMgr) GetFromContent(content string) (results []*ZawazawaPost, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("content = ?", content).Find(&results).Error

	return
}

// GetBatchFromContent 批量唯一主键查找
func (obj *_ZawazawaPostMgr) GetBatchFromContent(contents []string) (results []*ZawazawaPost, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("content IN (?)", contents).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容
func (obj *_ZawazawaPostMgr) GetFromCreatedAt(createdAt int) (results []*ZawazawaPost, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at = ?", createdAt).Find(&results).Error

	return
}

// GetBatchFromCreatedAt 批量唯一主键查找
func (obj *_ZawazawaPostMgr) GetBatchFromCreatedAt(createdAts []int) (results []*ZawazawaPost, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_at IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromCreatedBy 通过created_by获取内容
func (obj *_ZawazawaPostMgr) GetFromCreatedBy(createdBy int) (results []*ZawazawaPost, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_by = ?", createdBy).Find(&results).Error

	return
}

// GetBatchFromCreatedBy 批量唯一主键查找
func (obj *_ZawazawaPostMgr) GetBatchFromCreatedBy(createdBys []int) (results []*ZawazawaPost, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_by IN (?)", createdBys).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_ZawazawaPostMgr) FetchByPrimaryKey(id int) (result ZawazawaPost, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

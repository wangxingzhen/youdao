package test

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _UcInshootMgr struct {
	*_BaseMgr
}

// UcInshootMgr open func
func UcInshootMgr(db *gorm.DB) *_UcInshootMgr {
	if db == nil {
		panic(fmt.Errorf("UcInshootMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_UcInshootMgr{_BaseMgr: &_BaseMgr{DB: db.Table("uc_inshoot"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UcInshootMgr) GetTableName() string {
	return "uc_inshoot"
}

// Get 获取
func (obj *_UcInshootMgr) Get() (result UcInshoot, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_UcInshootMgr) Gets() (results []*UcInshoot, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_UcInshootMgr) WithID(id uint32) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUcID uc_id获取 uc的用户id
func (obj *_UcInshootMgr) WithUcID(ucID int) Option {
	return optionFunc(func(o *options) { o.query["uc_id"] = ucID })
}

// WithInID in_id获取 inshoot的用户id
func (obj *_UcInshootMgr) WithInID(inID uint32) Option {
	return optionFunc(func(o *options) { o.query["in_id"] = inID })
}

// GetByOption 功能选项模式获取
func (obj *_UcInshootMgr) GetByOption(opts ...Option) (result UcInshoot, err error) {
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
func (obj *_UcInshootMgr) GetByOptions(opts ...Option) (results []*UcInshoot, err error) {
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
func (obj *_UcInshootMgr) GetFromID(id uint32) (result UcInshoot, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_UcInshootMgr) GetBatchFromID(ids []uint32) (results []*UcInshoot, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromUcID 通过uc_id获取内容 uc的用户id
func (obj *_UcInshootMgr) GetFromUcID(ucID int) (result UcInshoot, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`uc_id` = ?", ucID).Find(&result).Error

	return
}

// GetBatchFromUcID 批量查找 uc的用户id
func (obj *_UcInshootMgr) GetBatchFromUcID(ucIDs []int) (results []*UcInshoot, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`uc_id` IN (?)", ucIDs).Find(&results).Error

	return
}

// GetFromInID 通过in_id获取内容 inshoot的用户id
func (obj *_UcInshootMgr) GetFromInID(inID uint32) (results []*UcInshoot, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`in_id` = ?", inID).Find(&results).Error

	return
}

// GetBatchFromInID 批量查找 inshoot的用户id
func (obj *_UcInshootMgr) GetBatchFromInID(inIDs []uint32) (results []*UcInshoot, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`in_id` IN (?)", inIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_UcInshootMgr) FetchByPrimaryKey(id uint32) (result UcInshoot, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error

	return
}

// FetchUniqueByUcID primay or index 获取唯一内容
func (obj *_UcInshootMgr) FetchUniqueByUcID(ucID int) (result UcInshoot, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`uc_id` = ?", ucID).Find(&result).Error

	return
}

func (obj *_UcInshootMgr) InsertData(result UcInshoot) (err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Create(&result).Error

	return
}

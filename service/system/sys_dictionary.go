package system

import (
	"errors"

	"achilles/global"
	"achilles/models/system"
	"achilles/models/system/request"

	"gorm.io/gorm"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteSysDict
//@description: 创建字典数据
//@param: SysDict model.SysDict
//@return: err error

type DictionaryService struct{}

func (dictionaryService *DictionaryService) CreateSysDict(SysDict system.SysDict) (err error) {
	if (!errors.Is(global.GVA_DB.First(&system.SysDict{}, "type = ?", SysDict.Type).Error, gorm.ErrRecordNotFound)) {
		return errors.New("存在相同的 type，不允许创建")
	}
	err = global.GVA_DB.Create(&SysDict).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteSysDict
//@description: 删除字典数据
//@param: SysDict model.SysDict
//@return: err error

func (dictionaryService *DictionaryService) DeleteSysDict(SysDict system.SysDict) (err error) {
	err = global.GVA_DB.Delete(&SysDict).Delete(&SysDict.SysDictItems).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateSysDict
//@description: 更新字典数据
//@param: SysDict *model.SysDict
//@return: err error

func (dictionaryService *DictionaryService) UpdateSysDict(SysDict *system.SysDict) (err error) {
	var dict system.SysDict
	SysDictMap := map[string]interface{}{
		"Name":   SysDict.Name,
		"Type":   SysDict.Type,
		"Status": SysDict.Status,
		"Desc":   SysDict.Desc,
	}
	db := global.GVA_DB.Where("id = ?", SysDict.ID).First(&dict)
	if dict.Type != SysDict.Type {
		if !errors.Is(global.GVA_DB.First(&system.SysDict{}, "type = ?", SysDict.Type).Error, gorm.ErrRecordNotFound) {
			return errors.New("存在相同的 type，不允许创建")
		}
	}
	err = db.Updates(SysDictMap).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetSysDict
//@description: 根据 id 或者 type 获取字典单条数据
//@param: Type string, Id uint
//@return: err error, SysDict model.SysDict

func (dictionaryService *DictionaryService) GetSysDict(Type string, Id uint) (err error, SysDict system.SysDict) {
	err = global.GVA_DB.Where("type = ? OR id = ? and status = ?", Type, Id, true).Preload("SysDictDetails", "status = ?", true).First(&SysDict).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: GetSysDictInfoList
//@description: 分页获取字典列表
//@param: info request.SysDictSearch
//@return: err error, list interface{}, total int64

func (dictionaryService *DictionaryService) GetSysDictInfoList(info request.SysDictSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建 db
	db := global.GVA_DB.Model(&system.SysDict{})
	var SysDicts []system.SysDict
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("`name` LIKE ?", "%"+info.Name+"%")
	}
	if info.Type != "" {
		db = db.Where("`type` LIKE ?", "%"+info.Type+"%")
	}
	if info.Status != nil {
		db = db.Where("`status` = ?", info.Status)
	}
	if info.Desc != "" {
		db = db.Where("`desc` LIKE ?", "%"+info.Desc+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&SysDicts).Error
	return err, SysDicts, total
}

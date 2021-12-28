package system

import (
	"achilles/global"
	"achilles/models/system"
	"achilles/models/system/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateSysDictItem
//@description: 创建字典详情数据
//@param: sysDictionaryDetail model.SysDictItem
//@return: err error

type DictionaryDetailService struct{}

func (dictionaryDetailService *DictionaryDetailService) CreateSysDictionaryDetail(sysDictionaryDetail system.SysDictItem) (err error) {
	err = global.GVA_DB.Create(&sysDictionaryDetail).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteSysDictItem
//@description: 删除字典详情数据
//@param: sysDictionaryDetail model.SysDictItem
//@return: err error

func (dictionaryDetailService *DictionaryDetailService) DeleteSysDictionaryDetail(sysDictionaryDetail system.SysDictItem) (err error) {
	err = global.GVA_DB.Delete(&sysDictionaryDetail).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateSysDictItem
//@description: 更新字典详情数据
//@param: sysDictionaryDetail *model.SysDictItem
//@return: err error

func (dictionaryDetailService *DictionaryDetailService) UpdateSysDictionaryDetail(sysDictionaryDetail *system.SysDictItem) (err error) {
	err = global.GVA_DB.Save(sysDictionaryDetail).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetSysDictionaryDetail
//@description: 根据 id 获取字典详情单条数据
//@param: id uint
//@return: err error, sysDictionaryDetail model.SysDictItem

func (dictionaryDetailService *DictionaryDetailService) GetSysDictionaryDetail(id uint) (err error, sysDictionaryDetail system.SysDictItem) {
	err = global.GVA_DB.Where("id = ?", id).First(&sysDictionaryDetail).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetSysDictionaryDetailInfoList
//@description: 分页获取字典详情列表
//@param: info request.SysDictionaryDetailSearch
//@return: err error, list interface{}, total int64

func (dictionaryDetailService *DictionaryDetailService) GetSysDictionaryDetailInfoList(info request.SysDictionaryDetailSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建 db
	db := global.GVA_DB.Model(&system.SysDictItem{})
	var sysDictionaryDetails []system.SysDictItem
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Label != "" {
		db = db.Where("label LIKE ?", "%"+info.Label+"%")
	}
	if info.Value != 0 {
		db = db.Where("value = ?", info.Value)
	}
	if info.Status != nil {
		db = db.Where("status = ?", info.Status)
	}
	if info.SysDictID != 0 {
		db = db.Where("sys_dictionary_id = ?", info.SysDictID)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&sysDictionaryDetails).Error
	return err, sysDictionaryDetails, total
}

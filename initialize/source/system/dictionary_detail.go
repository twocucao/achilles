package system

import (
	"achilles/global"
	"achilles/models/system"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var DictionaryDetail = new(dictionaryDetail)

type dictionaryDetail struct{}

func (d *dictionaryDetail) TableName() string {
	return "sys_dictionary_details"
}

func (d *dictionaryDetail) Initialize() error {
	status := new(bool)
	*status = true
	entities := []system.SysDictItem{
		{Label: "男", Value: 1, Status: status, Sort: 1, SysDictID: 1},
		{Label: "女", Value: 2, Status: status, Sort: 2, SysDictID: 1},
		{Label: "smallint", Value: 1, Status: status, Sort: 1, SysDictID: 2},
		{Label: "mediumint", Value: 2, Status: status, Sort: 2, SysDictID: 2},
		{Label: "int", Value: 3, Status: status, Sort: 3, SysDictID: 2},
		{Label: "bigint", Value: 4, Status: status, Sort: 4, SysDictID: 2},
		{Label: "date", Status: status, SysDictID: 3},
		{Label: "time", Value: 1, Status: status, Sort: 1, SysDictID: 3},
		{Label: "year", Value: 2, Status: status, Sort: 2, SysDictID: 3},
		{Label: "datetime", Value: 3, Status: status, Sort: 3, SysDictID: 3},
		{Label: "timestamp", Value: 5, Status: status, Sort: 5, SysDictID: 3},
		{Label: "float", Status: status, SysDictID: 4},
		{Label: "double", Value: 1, Status: status, Sort: 1, SysDictID: 4},
		{Label: "decimal", Value: 2, Status: status, Sort: 2, SysDictID: 4},
		{Label: "char", Status: status, SysDictID: 5},
		{Label: "varchar", Value: 1, Status: status, Sort: 1, SysDictID: 5},
		{Label: "tinyblob", Value: 2, Status: status, Sort: 2, SysDictID: 5},
		{Label: "tinytext", Value: 3, Status: status, Sort: 3, SysDictID: 5},
		{Label: "text", Value: 4, Status: status, Sort: 4, SysDictID: 5},
		{Label: "blob", Value: 5, Status: status, Sort: 5, SysDictID: 5},
		{Label: "mediumblob", Value: 6, Status: status, Sort: 6, SysDictID: 5},
		{Label: "mediumtext", Value: 7, Status: status, Sort: 7, SysDictID: 5},
		{Label: "longblob", Value: 8, Status: status, Sort: 8, SysDictID: 5},
		{Label: "longtext", Value: 9, Status: status, Sort: 9, SysDictID: 5},
		{Label: "tinyint", Status: status, SysDictID: 6},
	}
	if err := global.GVA_DB.Create(&entities).Error; err != nil {
		return errors.Wrap(err, d.TableName()+"表数据初始化失败！")
	}
	return nil
}

func (d *dictionaryDetail) CheckDataExist() bool {
	if errors.Is(global.GVA_DB.Where("id = ?", 23).First(&system.SysDictItem{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}

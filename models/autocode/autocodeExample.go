// 自动生成模板 SysDictItem
package autocode

import (
	"achilles/global"
)

// 如果含有 time.Time 请自行 import time 包
type AutoCodeExample struct {
	global.GVA_MODEL
	AutoCodeExampleField string `json:"autoCodeExampleField" form:"autoCodeExampleField" gorm:"column:auto_code_example_field;comment:仅作示例条目无实际作用"` // 展示值
}

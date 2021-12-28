package system

import (
	"achilles/global"
)

type SysApi struct {
	global.GVA_MODEL
	Path        string `json:"path" gorm:"comment:api 路径"`             // api 路径
	Description string `json:"description" gorm:"comment:api 中文描述"`    // api 中文描述
	ApiGroup    string `json:"apiGroup" gorm:"comment:api 组"`          // api 组
	Method      string `json:"method" gorm:"default:POST;comment:方法"` // 方法：创建 POST(默认)| 查看 GET| 更新 PUT| 删除 DELETE
}

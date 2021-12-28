// 自动生成模板 SysDictionaryDetail
package request

import (
	"achilles/model/autocode"
	"achilles/model/common/request"
)

// 如果含有 time.Time 请自行 import time 包
type AutoCodeExampleSearch struct {
	autocode.AutoCodeExample
	request.PageInfo
}

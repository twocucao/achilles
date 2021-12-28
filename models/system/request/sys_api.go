package request

import (
	"achilles/models/common/request"
	"achilles/models/system"
)

// api 分页条件查询及排序结构体
type SearchApiParams struct {
	system.SysApi
	request.PageInfo
	OrderKey string `json:"orderKey"` // 排序
	Desc     bool   `json:"desc"`     // 排序方式：升序 false(默认)| 降序 true
}

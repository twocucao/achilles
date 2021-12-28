package request

import (
	"achilles/model/common/request"
	"achilles/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}

package request

import (
	"achilles/models/common/request"
	"achilles/models/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}

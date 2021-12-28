package request

import (
	"achilles/models/common/request"
	"achilles/models/system"
)

type SysDictionarySearch struct {
	system.SysDictionary
	request.PageInfo
}

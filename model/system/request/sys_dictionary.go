package request

import (
	"achilles/model/common/request"
	"achilles/model/system"
)

type SysDictionarySearch struct {
	system.SysDictionary
	request.PageInfo
}

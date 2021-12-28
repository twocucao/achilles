package request

import (
	"achilles/models/common/request"
	"achilles/models/system"
)

type SysDictSearch struct {
	system.SysDict
	request.PageInfo
}

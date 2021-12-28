package request

import (
	"achilles/model/common/request"
	"achilles/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}

package request

import (
	"achilles/models/common/request"
	"achilles/models/system"
)

type SysOperationRecordSearch struct {
	system.SysAuditLog
	request.PageInfo
}

package request

import (
	"achilles/model/autocode"
	"achilles/model/common/request"
)

type {{.StructName}}Search struct{
    autocode.{{.StructName}}
    request.PageInfo
}

package v1

import (
	"achilles/apps/api/v1/autocode"
	"achilles/apps/api/v1/example"
	"achilles/apps/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup   system.ApiGroup
	ExampleApiGroup  example.ApiGroup
	AutoCodeApiGroup autocode.ApiGroup
}

var ApiGroupApp = new(ApiGroup)

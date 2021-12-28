package admin

import (
	"achilles/api/autocode/autocodeView"
	"achilles/api/example/exampleView"
	"achilles/api/system"
)

type ApiGroup struct {
	SystemApiGroup   system.ApiGroup
	ExampleApiGroup  exampleView.ApiGroup
	AutoCodeApiGroup autocodeView.ApiGroup
}

var ApiGroupApp = new(ApiGroup)

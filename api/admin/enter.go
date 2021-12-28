package admin

import (
	"achilles/api/autocode/autocodeview"
	"achilles/api/example/exampleview"
	"achilles/api/system/views"
)

type ApiGroup struct {
	SystemApiGroup   views.ApiGroup
	ExampleApiGroup  exampleview.ApiGroup
	AutoCodeApiGroup autocodeview.ApiGroup
}

var ApiGroupApp = new(ApiGroup)

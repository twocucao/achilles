package router

import (
	"achilles/apps/router/autocode"
	"achilles/apps/router/example"
	"achilles/apps/router/system"
)

type RouterGroup struct {
	System   system.RouterGroup
	Example  example.RouterGroup
	Autocode autocode.RouterGroup
}

var RouterGroupApp = new(RouterGroup)

package api

import (
	"achilles/api/autocode"
	"achilles/api/example"
	"achilles/api/system"
)

type RouterGroup struct {
	System   system.RouterGroup
	Example  example.RouterGroup
	Autocode autocode.RouterGroup
}

var RouterGroupApp = new(RouterGroup)

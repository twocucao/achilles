package api

import (
	"achilles/api/autocode"
	"achilles/api/example"
)

type RouterGroup struct {
	Example  example.RouterGroup
	Autocode autocode.RouterGroup
}

var RouterGroupApp = new(RouterGroup)

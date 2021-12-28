package system

import (
	"achilles/global"
)

type JWTDenyList struct {
	global.GVA_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}

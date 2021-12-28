package system

import (
	"achilles/apps/api/v1"

	"github.com/gin-gonic/gin"
)

type InitRouter struct{}

func (s *InitRouter) InitInitRouter(Router *gin.RouterGroup) {
	initRouter := Router.Group("init")
	dbApi := v1.ApiGroupApp.SystemApiGroup.DBApi
	{
		initRouter.POST("initdb", dbApi.InitDB)   // 创建 Api
		initRouter.POST("checkdb", dbApi.CheckDB) // 创建 Api
	}
}

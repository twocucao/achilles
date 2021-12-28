package system

import (
	"achilles/api/admin"
	"github.com/gin-gonic/gin"
)

type InitRouter struct{}

func (s *InitRouter) InitInitRouter(Router *gin.RouterGroup) {
	initRouter := Router.Group("init")
	dbApi := admin.ApiGroupApp.SystemApiGroup.DBApi
	{
		initRouter.POST("initdb", dbApi.InitDB)   // 创建 Api
		initRouter.POST("checkdb", dbApi.CheckDB) // 创建 Api
	}
}

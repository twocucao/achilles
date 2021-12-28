package system

import (
	"achilles/api/admin"
	"github.com/gin-gonic/gin"
)

type AutoCodeRouter struct{}

func (s *AutoCodeRouter) InitAutoCodeRouter(Router *gin.RouterGroup) {
	autoCodeRouter := Router.Group("autoCode")
	autoCodeApi := admin.ApiGroupApp.SystemApiGroup.AutoCodeApi
	{
		autoCodeRouter.GET("getDB", autoCodeApi.GetDB)            // 获取数据库
		autoCodeRouter.GET("getTables", autoCodeApi.GetTables)    // 获取对应数据库的表
		autoCodeRouter.GET("getColumn", autoCodeApi.GetColumn)    // 获取指定表所有字段信息
		autoCodeRouter.POST("preview", autoCodeApi.PreviewTemp)   // 获取自动创建代码预览
		autoCodeRouter.POST("createTemp", autoCodeApi.CreateTemp) // 创建自动化代码
	}
}

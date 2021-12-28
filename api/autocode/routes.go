package autocode

import (
	"achilles/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(group *gin.RouterGroup) {
	initAutoCode(group)
}

func initAutoCode(group *gin.RouterGroup) {
	r := group.Group("/autoCodeExample").Use(middleware.SysAuditLog())
	rNoAudit := group.Group("/autoCodeExample")
	api := new(ExampleApi)
	{
		r.POST("/create", api.CreateAutoCodeExample) // 新建 AutoCodeExample
		r.POST("/delete", api.DeleteAutoCodeExample) // 删除 AutoCodeExample
		r.POST("/update", api.UpdateAutoCodeExample) // 更新 AutoCodeExample
	}
	{
		rNoAudit.GET("/detail", api.FindAutoCodeExample)  // 根据 ID 获取 AutoCodeExample
		rNoAudit.GET("/list", api.GetAutoCodeExampleList) // 获取 AutoCodeExample 列表
	}
}

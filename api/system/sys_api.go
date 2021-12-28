package system

import (
	"achilles/api/admin"
	"achilles/middleware"

	"github.com/gin-gonic/gin"
)

type ApiRouter struct{}

func (s *ApiRouter) InitApiRouter(group *gin.RouterGroup) {
	r := group.Group("api").Use(middleware.OperationRecord())
	apiRouterWithoutRecord := group.Group("api")
	apiRouterApi := admin.ApiGroupApp.SystemApiGroup.SystemApiApi
	{
		r.POST("createApi", apiRouterApi.CreateApi)               // 创建 Api
		r.POST("deleteApi", apiRouterApi.DeleteApi)               // 删除 Api
		r.POST("getApiById", apiRouterApi.GetApiById)             // 获取单条 Api 消息
		r.POST("updateApi", apiRouterApi.UpdateApi)               // 更新 api
		r.DELETE("deleteApisByIds", apiRouterApi.DeleteApisByIds) // 删除选中 api
	}
	{
		apiRouterWithoutRecord.POST("getAllApis", apiRouterApi.GetAllApis) // 获取所有 api
		apiRouterWithoutRecord.POST("getApiList", apiRouterApi.GetApiList) // 获取 Api 列表
	}
}

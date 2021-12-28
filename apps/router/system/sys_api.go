package system

import (
	"achilles/apps/api/v1"
	"achilles/middleware"

	"github.com/gin-gonic/gin"
)

type ApiRouter struct{}

func (s *ApiRouter) InitApiRouter(Router *gin.RouterGroup) {
	apiRouter := Router.Group("api").Use(middleware.OperationRecord())
	apiRouterWithoutRecord := Router.Group("api")
	apiRouterApi := v1.ApiGroupApp.SystemApiGroup.SystemApiApi
	{
		apiRouter.POST("createApi", apiRouterApi.CreateApi)               // 创建 Api
		apiRouter.POST("deleteApi", apiRouterApi.DeleteApi)               // 删除 Api
		apiRouter.POST("getApiById", apiRouterApi.GetApiById)             // 获取单条 Api 消息
		apiRouter.POST("updateApi", apiRouterApi.UpdateApi)               // 更新 api
		apiRouter.DELETE("deleteApisByIds", apiRouterApi.DeleteApisByIds) // 删除选中 api
	}
	{
		apiRouterWithoutRecord.POST("getAllApis", apiRouterApi.GetAllApis) // 获取所有 api
		apiRouterWithoutRecord.POST("getApiList", apiRouterApi.GetApiList) // 获取 Api 列表
	}
}

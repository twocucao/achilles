package system

import (
	"achilles/apps/api/v1"
	"achilles/middleware"

	"github.com/gin-gonic/gin"
)

type DictionaryDetailRouter struct{}

func (s *DictionaryDetailRouter) InitSysDictionaryDetailRouter(Router *gin.RouterGroup) {
	dictionaryDetailRouter := Router.Group("sysDictionaryDetail").Use(middleware.OperationRecord())
	dictionaryDetailRouterWithoutRecord := Router.Group("sysDictionaryDetail")
	sysDictionaryDetailApi := v1.ApiGroupApp.SystemApiGroup.DictionaryDetailApi
	{
		dictionaryDetailRouter.POST("createSysDictionaryDetail", sysDictionaryDetailApi.CreateSysDictionaryDetail)   // 新建 SysDictionaryDetail
		dictionaryDetailRouter.DELETE("deleteSysDictionaryDetail", sysDictionaryDetailApi.DeleteSysDictionaryDetail) // 删除 SysDictionaryDetail
		dictionaryDetailRouter.PUT("updateSysDictionaryDetail", sysDictionaryDetailApi.UpdateSysDictionaryDetail)    // 更新 SysDictionaryDetail
	}
	{
		dictionaryDetailRouterWithoutRecord.GET("findSysDictionaryDetail", sysDictionaryDetailApi.FindSysDictionaryDetail)       // 根据 ID 获取 SysDictionaryDetail
		dictionaryDetailRouterWithoutRecord.GET("getSysDictionaryDetailList", sysDictionaryDetailApi.GetSysDictionaryDetailList) // 获取 SysDictionaryDetail 列表
	}
}

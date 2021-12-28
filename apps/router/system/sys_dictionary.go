package system

import (
	"achilles/apps/api/v1"
	"achilles/middleware"

	"github.com/gin-gonic/gin"
)

type DictionaryRouter struct{}

func (s *DictionaryRouter) InitSysDictionaryRouter(Router *gin.RouterGroup) {
	sysDictionaryRouter := Router.Group("sysDictionary").Use(middleware.OperationRecord())
	sysDictionaryRouterWithoutRecord := Router.Group("sysDictionary")
	sysDictionaryApi := v1.ApiGroupApp.SystemApiGroup.DictionaryApi
	{
		sysDictionaryRouter.POST("createSysDictionary", sysDictionaryApi.CreateSysDictionary)   // 新建 SysDictionary
		sysDictionaryRouter.DELETE("deleteSysDictionary", sysDictionaryApi.DeleteSysDictionary) // 删除 SysDictionary
		sysDictionaryRouter.PUT("updateSysDictionary", sysDictionaryApi.UpdateSysDictionary)    // 更新 SysDictionary
	}
	{
		sysDictionaryRouterWithoutRecord.GET("findSysDictionary", sysDictionaryApi.FindSysDictionary)       // 根据 ID 获取 SysDictionary
		sysDictionaryRouterWithoutRecord.GET("getSysDictionaryList", sysDictionaryApi.GetSysDictionaryList) // 获取 SysDictionary 列表
	}
}

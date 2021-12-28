package system

import (
	"achilles/api/admin"
	"github.com/gin-gonic/gin"
)

type AutoCodeHistoryRouter struct{}

func (s *AutoCodeRouter) InitAutoCodeHistoryRouter(Router *gin.RouterGroup) {
	autoCodeHistoryRouter := Router.Group("autoCode")
	autoCodeHistoryApi := admin.ApiGroupApp.SystemApiGroup.AutoCodeHistoryApi
	{
		autoCodeHistoryRouter.POST("getMeta", autoCodeHistoryApi.First)         // 根据 id 获取 meta 信息
		autoCodeHistoryRouter.POST("rollback", autoCodeHistoryApi.RollBack)     // 回滚
		autoCodeHistoryRouter.POST("delSysHistory", autoCodeHistoryApi.Delete)  // 删除回滚记录
		autoCodeHistoryRouter.POST("getSysHistory", autoCodeHistoryApi.GetList) // 获取回滚记录分页
	}
}

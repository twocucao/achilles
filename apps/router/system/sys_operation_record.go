package system

import (
	"achilles/apps/api/v1"

	"github.com/gin-gonic/gin"
)

type OperationRecordRouter struct{}

func (s *OperationRecordRouter) InitSysOperationRecordRouter(Router *gin.RouterGroup) {
	operationRecordRouter := Router.Group("sysOperationRecord")
	authorityMenuApi := v1.ApiGroupApp.SystemApiGroup.OperationRecordApi
	{
		operationRecordRouter.POST("createSysOperationRecord", authorityMenuApi.CreateSysOperationRecord)             // 新建 SysOperationRecord
		operationRecordRouter.DELETE("deleteSysOperationRecord", authorityMenuApi.DeleteSysOperationRecord)           // 删除 SysOperationRecord
		operationRecordRouter.DELETE("deleteSysOperationRecordByIds", authorityMenuApi.DeleteSysOperationRecordByIds) // 批量删除 SysOperationRecord
		operationRecordRouter.GET("findSysOperationRecord", authorityMenuApi.FindSysOperationRecord)                  // 根据 ID 获取 SysOperationRecord
		operationRecordRouter.GET("getSysOperationRecordList", authorityMenuApi.GetSysOperationRecordList)            // 获取 SysOperationRecord 列表

	}
}

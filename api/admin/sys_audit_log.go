package admin

import (
	"achilles/global"
	"achilles/models/common/request"
	"achilles/models/common/response"
	"achilles/models/system"
	systemReq "achilles/models/system/request"
	"achilles/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SysAuditLog struct{}

// @Tags SysAuditLog
// @Summary 创建 SysAuditLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysAuditLog true "创建 SysAuditLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysOperationRecord/createSysOperationRecord [post]
func (s *SysAuditLog) CreateSysOperationRecord(c *gin.Context) {
	var sysOperationRecord system.SysAuditLog
	_ = c.ShouldBindJSON(&sysOperationRecord)
	if err := operationRecordService.CreateSysOperationRecord(sysOperationRecord); err != nil {
		global.GVA_LOG.Error("创建失败！", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags SysAuditLog
// @Summary 删除 SysAuditLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysAuditLog true "SysAuditLog 模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysOperationRecord/deleteSysOperationRecord [delete]
func (s *SysAuditLog) DeleteSysOperationRecord(c *gin.Context) {
	var sysOperationRecord system.SysAuditLog
	_ = c.ShouldBindJSON(&sysOperationRecord)
	if err := operationRecordService.DeleteSysOperationRecord(sysOperationRecord); err != nil {
		global.GVA_LOG.Error("删除失败！", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags SysAuditLog
// @Summary 批量删除 SysAuditLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除 SysAuditLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /sysOperationRecord/deleteSysOperationRecordByIds [delete]
func (s *SysAuditLog) DeleteSysOperationRecordByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := operationRecordService.DeleteSysOperationRecordByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败！", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags SysAuditLog
// @Summary 用 id 查询 SysAuditLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query system.SysAuditLog true "Id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysOperationRecord/findSysOperationRecord [get]
func (s *SysAuditLog) FindSysOperationRecord(c *gin.Context) {
	var sysOperationRecord system.SysAuditLog
	_ = c.ShouldBindQuery(&sysOperationRecord)
	if err := utils.Verify(sysOperationRecord, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, resysOperationRecord := operationRecordService.GetSysOperationRecord(sysOperationRecord.ID); err != nil {
		global.GVA_LOG.Error("查询失败！", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithDetailed(gin.H{"resysOperationRecord": resysOperationRecord}, "查询成功", c)
	}
}

// @Tags SysAuditLog
// @Summary 分页获取 SysAuditLog 列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.SysOperationRecordSearch true "页码，每页大小，搜索条件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysOperationRecord/getSysOperationRecordList [get]
func (s *SysAuditLog) GetSysOperationRecordList(c *gin.Context) {
	var pageInfo systemReq.SysOperationRecordSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := operationRecordService.GetSysOperationRecordInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败！", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

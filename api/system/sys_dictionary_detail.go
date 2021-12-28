package system

import (
	"achilles/global"
	"achilles/models/common/response"
	"achilles/models/system"
	"achilles/models/system/request"
	"achilles/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DictionaryDetailApi struct{}

// @Tags SysDictItem
// @Summary 创建 SysDictItem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysDictItem true "SysDictItem 模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /sysDictionaryDetail/createSysDictionaryDetail [post]
func (s *DictionaryDetailApi) CreateSysDictItem(c *gin.Context) {
	var detail system.SysDictItem
	_ = c.ShouldBindJSON(&detail)
	if err := dictionaryDetailService.CreateSysDictionaryDetail(detail); err != nil {
		global.GVA_LOG.Error("创建失败！", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags SysDictItem
// @Summary 删除 SysDictItem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysDictItem true "SysDictItem 模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysDictionaryDetail/deleteSysDictionaryDetail [delete]
func (s *DictionaryDetailApi) DeleteSysDictItem(c *gin.Context) {
	var detail system.SysDictItem
	_ = c.ShouldBindJSON(&detail)
	if err := dictionaryDetailService.DeleteSysDictionaryDetail(detail); err != nil {
		global.GVA_LOG.Error("删除失败！", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags SysDictItem
// @Summary 更新 SysDictItem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysDictItem true "更新 SysDictItem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysDictionaryDetail/updateSysDictionaryDetail [put]
func (s *DictionaryDetailApi) UpdateSysDictItem(c *gin.Context) {
	var detail system.SysDictItem
	_ = c.ShouldBindJSON(&detail)
	if err := dictionaryDetailService.UpdateSysDictionaryDetail(&detail); err != nil {
		global.GVA_LOG.Error("更新失败！", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags SysDictItem
// @Summary 用 id 查询 SysDictItem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query system.SysDictItem true "用 id 查询 SysDictItem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysDictionaryDetail/findSysDictionaryDetail [get]
func (s *DictionaryDetailApi) FindSysDictItem(c *gin.Context) {
	var detail system.SysDictItem
	_ = c.ShouldBindQuery(&detail)
	if err := utils.Verify(detail, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, resysDictionaryDetail := dictionaryDetailService.GetSysDictionaryDetail(detail.ID); err != nil {
		global.GVA_LOG.Error("查询失败！", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithDetailed(gin.H{"resysDictionaryDetail": resysDictionaryDetail}, "查询成功", c)
	}
}

// @Tags SysDictItem
// @Summary 分页获取 SysDictItem 列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.SysDictionaryDetailSearch true "页码，每页大小，搜索条件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysDictionaryDetail/getSysDictionaryDetailList [get]
func (s *DictionaryDetailApi) GetSysDictItem(c *gin.Context) {
	var pageInfo request.SysDictionaryDetailSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := dictionaryDetailService.GetSysDictionaryDetailInfoList(pageInfo); err != nil {
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

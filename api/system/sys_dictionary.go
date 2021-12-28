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

type DictionaryApi struct{}

// @Tags SysDict
// @Summary 创建 SysDict
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysDict true "SysDict 模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /SysDict/createSysDict [post]
func (s *DictionaryApi) CreateSysDict(c *gin.Context) {
	var dictionary system.SysDict
	_ = c.ShouldBindJSON(&dictionary)
	if err := dictionaryService.CreateSysDict(dictionary); err != nil {
		global.GVA_LOG.Error("创建失败！", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags SysDict
// @Summary 删除 SysDict
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysDict true "SysDict 模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /SysDict/deleteSysDict [delete]
func (s *DictionaryApi) DeleteSysDict(c *gin.Context) {
	var dictionary system.SysDict
	_ = c.ShouldBindJSON(&dictionary)
	if err := dictionaryService.DeleteSysDict(dictionary); err != nil {
		global.GVA_LOG.Error("删除失败！", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags SysDict
// @Summary 更新 SysDict
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysDict true "SysDict 模型"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /SysDict/updateSysDict [put]
func (s *DictionaryApi) UpdateSysDict(c *gin.Context) {
	var dictionary system.SysDict
	_ = c.ShouldBindJSON(&dictionary)
	if err := dictionaryService.UpdateSysDict(&dictionary); err != nil {
		global.GVA_LOG.Error("更新失败！", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags SysDict
// @Summary 用 id 查询 SysDict
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query system.SysDict true "ID 或字典英名"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /SysDict/findSysDict [get]
func (s *DictionaryApi) FindSysDict(c *gin.Context) {
	var dictionary system.SysDict
	_ = c.ShouldBindQuery(&dictionary)
	if err, SysDict := dictionaryService.GetSysDict(dictionary.Type, dictionary.ID); err != nil {
		global.GVA_LOG.Error("查询失败！", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithDetailed(gin.H{"reSysDict": SysDict}, "查询成功", c)
	}
}

// @Tags SysDict
// @Summary 分页获取 SysDict 列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.SysDictSearch true "页码，每页大小，搜索条件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /SysDict/getSysDictList [get]
func (s *DictionaryApi) GetSysDictList(c *gin.Context) {
	var pageInfo request.SysDictSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err := utils.Verify(pageInfo.PageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, list, total := dictionaryService.GetSysDictInfoList(pageInfo); err != nil {
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

package exampleView

import (
	"achilles/global"
	"achilles/models/common/response"
	"achilles/models/example"
	"achilles/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ExcelApi struct{}

// /excel/importExcel 接口，与 upload 接口作用类似，只是把文件存到 resource/excel 目录下，用于导入 Excel 时存放 Excel 文件 (ExcelImport.xlsx)
// /excel/loadExcel 接口，用于读取 resource/excel 目录下的文件 ((ExcelImport.xlsx) 并加载为 []model.SysBaseMenu 类型的示例数据
// /excel/exportExcel 接口，用于读取前端传来的 tableData，生成 Excel 文件并返回
// /excel/downloadTemplate 接口，用于下载 resource/excel 目录下的 ExcelTemplate.xlsx 文件，作为导入的模板

// @Tags excel
// @Summary 导出 Excel
// @Security ApiKeyAuth
// @accept application/json
// @Produce  application/octet-stream
// @Param data body example.ExcelInfo true "导出 Excel 文件信息"
// @Success 200
// @Router /excel/exportExcel [post]
func (e *ExcelApi) ExportExcel(c *gin.Context) {
	var excelInfo example.ExcelInfo
	_ = c.ShouldBindJSON(&excelInfo)
	filePath := global.GVA_CONFIG.Excel.Dir + excelInfo.FileName
	err := excelService.ParseInfoList2Excel(excelInfo.InfoList, filePath)
	if err != nil {
		global.GVA_LOG.Error("转换 Excel 失败！", zap.Error(err))
		response.FailWithMessage("转换 Excel 失败", c)
		return
	}
	c.Writer.Header().Add("success", "true")
	c.File(filePath)
}

// @Tags excel
// @Summary 导入 Excel 文件
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "导入 Excel 文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"导入成功"}"
// @Router /excel/importExcel [post]
func (e *ExcelApi) ImportExcel(c *gin.Context) {
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		global.GVA_LOG.Error("接收文件失败！", zap.Error(err))
		response.FailWithMessage("接收文件失败", c)
		return
	}
	_ = c.SaveUploadedFile(header, global.GVA_CONFIG.Excel.Dir+"ExcelImport.xlsx")
	response.OkWithMessage("导入成功", c)
}

// @Tags excel
// @Summary 加载 Excel 数据
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"加载数据成功"}"
// @Router /excel/loadExcel [get]
func (e *ExcelApi) LoadExcel(c *gin.Context) {
	menus, err := excelService.ParseExcel2InfoList()
	if err != nil {
		global.GVA_LOG.Error("加载数据失败！", zap.Error(err))
		response.FailWithMessage("加载数据失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     menus,
		Total:    int64(len(menus)),
		Page:     1,
		PageSize: 999,
	}, "加载数据成功", c)
}

// @Tags excel
// @Summary 下载模板
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param fileName query string true "模板名称"
// @Success 200
// @Router /excel/downloadTemplate [get]
func (e *ExcelApi) DownloadTemplate(c *gin.Context) {
	fileName := c.Query("fileName")
	filePath := global.GVA_CONFIG.Excel.Dir + fileName
	ok, err := utils.PathExists(filePath)
	if !ok || err != nil {
		global.GVA_LOG.Error("文件不存在！", zap.Error(err))
		response.FailWithMessage("文件不存在", c)
		return
	}
	c.Writer.Header().Add("success", "true")
	c.File(filePath)
}

package example

import (
	"achilles/apps/api/v1"

	"github.com/gin-gonic/gin"
)

type ExcelRouter struct{}

func (e *ExcelRouter) InitExcelRouter(Router *gin.RouterGroup) {
	excelRouter := Router.Group("excel")
	exaExcelApi := v1.ApiGroupApp.ExampleApiGroup.ExcelApi
	{
		excelRouter.POST("importExcel", exaExcelApi.ImportExcel)          // 导入 Excel
		excelRouter.GET("loadExcel", exaExcelApi.LoadExcel)               // 加载 Excel 数据
		excelRouter.POST("exportExcel", exaExcelApi.ExportExcel)          // 导出 Excel
		excelRouter.GET("downloadTemplate", exaExcelApi.DownloadTemplate) // 下载模板文件
	}
}

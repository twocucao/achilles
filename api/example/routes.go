package example

import (
	"achilles/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(group *gin.RouterGroup) {
	initCustomerRouter(group)
	initExcelRouter(group)
	initFileUploader(group)
}

func initFileUploader(group *gin.RouterGroup) {
	r := group.Group("fileUploadAndDownload")
	api := new(FileUploadAndDownloadApi)
	{
		r.POST("upload", api.UploadFile)                                 // 上传文件
		r.POST("getFileList", api.GetFileList)                           // 获取上传文件列表
		r.POST("deleteFile", api.DeleteFile)                             // 删除指定文件
		r.POST("breakpointContinue", api.BreakpointContinue)             // 断点续传
		r.GET("findFile", api.FindFile)                                  // 查询当前文件成功的切片
		r.POST("breakpointContinueFinish", api.BreakpointContinueFinish) // 查询当前文件成功的切片
		r.POST("removeChunk", api.RemoveChunk)                           // 查询当前文件成功的切片
	}
}

func initExcelRouter(group *gin.RouterGroup) {
	excelRouter := group.Group("excel")
	exaExcelApi := new(ExcelApi)
	{
		excelRouter.POST("importExcel", exaExcelApi.ImportExcel)          // 导入 Excel
		excelRouter.GET("loadExcel", exaExcelApi.LoadExcel)               // 加载 Excel 数据
		excelRouter.POST("exportExcel", exaExcelApi.ExportExcel)          // 导出 Excel
		excelRouter.GET("downloadTemplate", exaExcelApi.DownloadTemplate) // 下载模板文件
	}
}

func initCustomerRouter(group *gin.RouterGroup) {
	r := group.Group("/customer").Use(middleware.SysAuditLog())
	rNoAudit := group.Group("/customer")
	api := new(CustomerApi)
	{
		r.POST("/create", api.CreateExaCustomer)       // 创建客户
		r.POST("/update", api.UpdateExaCustomer)       // 更新客户
		r.POST("/delete", api.DeleteExaCustomer)       // 删除客户
		rNoAudit.POST("/detail", api.GetExaCustomer)   // 获取单一客户信息
		rNoAudit.POST("/list", api.GetExaCustomerList) // 获取客户列表
	}
}

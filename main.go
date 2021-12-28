package main

import (
	"achilles/core"
	"achilles/global"
	"achilles/initialize"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title Swagger VjExample API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {
	global.GVA_VP = core.Viper()      // 初始化 Viper
	global.GVA_LOG = core.Zap()       // 初始化 zap 日志库
	global.GVA_DB = initialize.Gorm() // gorm 连接数据库
	initialize.Timer()
	initialize.DBList()
	if global.GVA_DB != nil {
		initialize.RegisterTables(global.GVA_DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}
	core.RunWindowsServer()
}

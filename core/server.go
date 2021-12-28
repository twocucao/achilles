package core

import (
	"fmt"
	"time"

	"achilles/global"
	"achilles/initialize"
	"achilles/service/system"

	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.GVA_CONFIG.System.UseMultipoint {
		// 初始化 redis 服务
		initialize.Redis()
	}

	// 从 db 加载 jwt 数据
	if global.GVA_DB != nil {
		system.LoadAll()
	}

	Router := initialize.Routers()

	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := initServer(address, Router)
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.GVA_LOG.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`
====================

  欢迎使用 achilles

====================
`, address)
	global.GVA_LOG.Error(s.ListenAndServe().Error())
}

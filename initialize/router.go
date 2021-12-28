package initialize

import (
	"achilles/api/admin"
	"achilles/api/autocode"
	"achilles/api/example"
	_ "achilles/docs"
	"achilles/global"
	"achilles/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// 初始化总路由

func Routers() *gin.Engine {
	r := gin.Default()
	r.StaticFS(global.GVA_CONFIG.Local.Path, http.Dir(global.GVA_CONFIG.Local.Path)) // 为用户头像和文件提供静态地址
	global.GVA_LOG.Info("use middleware logger")
	r.Use(middleware.Cors()) // cors allow all
	// r.Use(middleware.CorsByRules()) // cors allow by rules
	global.GVA_LOG.Info("use middleware cors")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.GVA_LOG.Info("register swagger handler")
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, "ok")
	})
	admin.InitRouter(r.Group("/api/admin"))
	autocode.InitRouter(r.Group("/api/autocode"))
	example.InitRouter(r.Group("/api/example"))
	global.GVA_LOG.Info("router register success")
	return r
}

package initialize

import (
	"achilles/global"

	email "github.com/flipped-aurora/gva-plugins/email" // 在线仓库模式 go

	//"achilles/plugin/email" // 本地插件仓库地址模式
	"achilles/plugin/example_plugin"
	"achilles/utils/plugin"

	"github.com/gin-gonic/gin"
)

func PluginInit(group *gin.RouterGroup, Plugin ...plugin.Plugin) {
	for i := range Plugin {
		PluginGroup := group.Group(Plugin[i].RouterPath())
		Plugin[i].Register(PluginGroup)
	}
}

func InstallPlugin(PublicGroup *gin.RouterGroup, PrivateGroup *gin.RouterGroup) {
	//  添加开放权限的插件 示例
	PluginInit(PublicGroup,
		example_plugin.ExamplePlugin)

	//  添加跟角色挂钩权限的插件 示例 本地示例模式于在线仓库模式注意上方的 import 可以自行切换 效果相同
	PluginInit(PrivateGroup, email.CreateEmailPlug(
		global.GVA_CONFIG.Email.To,
		global.GVA_CONFIG.Email.From,
		global.GVA_CONFIG.Email.Host,
		global.GVA_CONFIG.Email.Secret,
		global.GVA_CONFIG.Email.Nickname,
		global.GVA_CONFIG.Email.Port,
		global.GVA_CONFIG.Email.IsSSL,
	))
}

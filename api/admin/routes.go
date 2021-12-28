package admin

import (
	"achilles/api/system"
	"achilles/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(group *gin.RouterGroup) {
	initBaseRouter(group)
	initApiRouter(group)
	initAuthorityRouter(group)
	initAutoCoder(group)
	initCasbinRouter(group)
	initDictRouter(group)
	initDBRouter(group)
	initJwtRouter(group)
	initMenuRouter(group)
	initSystemRouter(group)
	initUserRouter(group)
	initOperationRecordRouter(group)

}

func initOperationRecordRouter(group *gin.RouterGroup) {
	r := group.Group("/sysAuditLog")
	api := ApiGroupApp.SystemApiGroup.OperationRecordApi
	{
		r.POST("/create", api.CreateSysOperationRecord)           // 新建 SysOperationRecord
		r.POST("/delete", api.DeleteSysOperationRecord)           // 删除 SysOperationRecord
		r.POST("/bulk_delete", api.DeleteSysOperationRecordByIds) // 批量删除 SysOperationRecord
		r.POST("/detail", api.FindSysOperationRecord)             // 根据 ID 获取 SysOperationRecord
		r.POST("/list", api.GetSysOperationRecordList)            // 获取 SysOperationRecord 列表
	}
}

func initUserRouter(group *gin.RouterGroup) {
	r := group.Group("/user").Use(middleware.OperationRecord())
	rNoAudit := group.Group("/user")
	baseApi := ApiGroupApp.SystemApiGroup.BaseApi
	{
		r.POST("/register", baseApi.Register)                     // 用户注册账号
		r.POST("/changePassword", baseApi.ChangePassword)         // 用户修改密码
		r.POST("/setUserAuthority", baseApi.SetUserAuthority)     // 设置用户权限
		r.POST("/delete", baseApi.DeleteUser)                     // 删除用户
		r.POST("/setUserInfo", baseApi.SetUserInfo)               // 设置用户信息
		r.POST("/setUserAuthorities", baseApi.SetUserAuthorities) // 设置用户权限组
		r.POST("/resetPassword", baseApi.ResetPassword)           // 设置用户权限组
		rNoAudit.POST("/list", baseApi.GetUserList)               // 分页获取用户列表
		rNoAudit.GET("/detail", baseApi.GetUserInfo)              // 获取自身信息
	}
}

func initSystemRouter(group *gin.RouterGroup) {
	r := group.Group("/system").Use(middleware.OperationRecord())
	api := ApiGroupApp.SystemApiGroup.SystemApi
	{
		r.GET("/getConfig", api.GetSystemConfig)    // 获取配置文件内容
		r.POST("/setConfig", api.SetSystemConfig)   // 设置配置文件内容
		r.POST("/getServerInfo", api.GetServerInfo) // 获取服务器信息
		r.POST("/reloadSystem", api.ReloadSystem)   // 重启服务
	}
}

func initMenuRouter(group *gin.RouterGroup) system.AuthorityMenuApi {
	r := group.Group("/menu").Use(middleware.OperationRecord())
	rNoAudit := group.Group("/menu")
	api := ApiGroupApp.SystemApiGroup.AuthorityMenuApi
	{
		r.POST("/addBaseMenu", api.AddBaseMenu)                  // 新增菜单
		r.POST("/addMenuAuthority", api.AddMenuAuthority)        // 增加 menu 和角色关联关系
		r.POST("/deleteBaseMenu", api.DeleteBaseMenu)            // 删除菜单
		r.POST("/updateBaseMenu", api.UpdateBaseMenu)            // 更新菜单
		rNoAudit.POST("/getMenu", api.GetMenu)                   // 获取菜单树
		rNoAudit.POST("/getMenuList", api.GetMenuList)           // 分页获取基础 menu 列表
		rNoAudit.POST("/getBaseMenuTree", api.GetBaseMenuTree)   // 获取用户动态路由
		rNoAudit.POST("/getMenuAuthority", api.GetMenuAuthority) // 获取指定角色 menu
		rNoAudit.POST("/getBaseMenuById", api.GetBaseMenuById)   // 根据 id 获取菜单
	}
	return api
}

func initBaseRouter(group *gin.RouterGroup) {
	r := group.Group("")
	api := ApiGroupApp.SystemApiGroup.BaseApi
	{
		r.POST("/login", api.Login)
		r.POST("/captcha", api.Captcha)
	}
}

func initCasbinRouter(group *gin.RouterGroup) {
	r := group.Group("/casbin").Use(middleware.OperationRecord())
	rNoAudit := group.Group("/casbin")
	api := ApiGroupApp.SystemApiGroup.CasbinApi
	{
		r.POST("/update", api.UpdateCasbin)
		rNoAudit.POST("/getPolicyPathByAuthorityId", api.GetPolicyPathByAuthorityId)
	}
}

func initDictRouter(group *gin.RouterGroup) {
	{
		r := group.Group("/sysDict").Use(middleware.OperationRecord())
		rNoAudit := group.Group("/sysDict")
		sysDictionaryApi := ApiGroupApp.SystemApiGroup.DictionaryApi
		{
			r.POST("/create", sysDictionaryApi.CreateSysDict)      // 新建 SysDict
			r.POST("/delete", sysDictionaryApi.DeleteSysDict)      // 删除 SysDict
			r.POST("/update", sysDictionaryApi.UpdateSysDict)      // 更新 SysDict
			rNoAudit.GET("/detail", sysDictionaryApi.FindSysDict)  // 根据 ID 获取 SysDict
			rNoAudit.GET("/list", sysDictionaryApi.GetSysDictList) // 获取 SysDict 列表
		}
		api := ApiGroupApp.SystemApiGroup.DictionaryDetailApi
		{
			r.POST("/item/create", api.CreateSysDictItem)     // 新建 SysDictItem
			r.PUT("/item/update", api.UpdateSysDictItem)      // 更新 SysDictItem
			r.DELETE("/item/delete", api.DeleteSysDictItem)   // 删除 SysDictItem
			rNoAudit.GET("/item/detail", api.FindSysDictItem) // 根据 ID 获取 SysDictItem
			rNoAudit.GET("/item/list", api.GetSysDictItem)    // 获取 SysDictItem 列表
		}
	}

}

func initAutoCoder(group *gin.RouterGroup) {
	r := group.Group("/autoCode")
	api := ApiGroupApp.SystemApiGroup.AutoCodeApi
	apiHistory := ApiGroupApp.SystemApiGroup.AutoCodeHistoryApi
	{
		r.GET("/getDB", api.GetDB)                   // 获取数据库
		r.GET("/getTables", api.GetTables)           // 获取对应数据库的表
		r.GET("/getColumn", api.GetColumn)           // 获取指定表所有字段信息
		r.POST("/preview", api.PreviewTemp)          // 获取自动创建代码预览
		r.POST("/createTemp", api.CreateTemp)        // 创建自动化代码
		r.POST("/getMeta", apiHistory.First)         // 根据 id 获取 meta 信息
		r.POST("/rollback", apiHistory.RollBack)     // 回滚
		r.POST("/delSysHistory", apiHistory.Delete)  // 删除回滚记录
		r.POST("/getSysHistory", apiHistory.GetList) // 获取回滚记录分页
	}
}

func initAuthorityRouter(group *gin.RouterGroup) {
	r := group.Group("/authority").Use(middleware.OperationRecord())
	rNotAudit := group.Group("/authority")
	api := ApiGroupApp.SystemApiGroup.AuthorityApi
	{
		r.POST("/create", api.CreateAuthority)        // 创建角色
		r.POST("/delete", api.DeleteAuthority)        // 删除角色
		r.POST("/update", api.UpdateAuthority)        // 更新角色
		r.POST("/copy", api.CopyAuthority)            // 拷贝角色
		r.POST("/setData", api.SetDataAuthority)      // 设置角色资源权限
		rNotAudit.POST("/list", api.GetAuthorityList) // 获取角色列表
	}
}

func initApiRouter(group *gin.RouterGroup) {
	r := group.Group("/api").Use(middleware.OperationRecord())
	rNotAudit := group.Group("/api")
	api := ApiGroupApp.SystemApiGroup.SystemApiApi
	{
		r.POST("/create", api.CreateApi)            // 创建 Api
		r.POST("/delete", api.DeleteApi)            // 删除 Api
		r.POST("/detail", api.GetApiById)           // 获取单条 Api 消息
		r.POST("/update", api.UpdateApi)            // 更新 api
		r.POST("/bulk_delete", api.DeleteApisByIds) // 删除选中 api
	}
	{
		rNotAudit.POST("/getAllApis", api.GetAllApis) // 获取所有 api
		rNotAudit.POST("/getApiList", api.GetApiList) // 获取 Api 列表
	}
}

func initDBRouter(group *gin.RouterGroup) {
	r := group.Group("/init")
	api := ApiGroupApp.SystemApiGroup.DBApi
	{
		r.POST("/initdb", api.InitDB)   // 创建 Api
		r.POST("/checkdb", api.CheckDB) // 创建 Api
	}
}

func initJwtRouter(group *gin.RouterGroup) {
	r := group.Group("/jwt")
	api := ApiGroupApp.SystemApiGroup.JwtApi
	{
		r.POST("/deny", api.DenyJWT)
		r.POST("/expire", api.ExpireJWT)
	}
}

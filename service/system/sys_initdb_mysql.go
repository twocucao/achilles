package system

import (
	"achilles/initialize/source/example"
	system2 "achilles/initialize/source/system"
	"fmt"
	"path/filepath"

	"achilles/config"
	"achilles/global"
	model "achilles/models/system"
	"achilles/models/system/request"
	"achilles/utils"

	uuid "github.com/satori/go.uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// writeMysqlConfig mysql 回写配置
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [songzhibin97](https://github.com/songzhibin97)
func (initDBService *InitDBService) writeMysqlConfig(mysql config.Mysql) error {
	global.GVA_CONFIG.Mysql = mysql
	cs := utils.StructToMap(global.GVA_CONFIG)
	for k, v := range cs {
		global.GVA_VP.Set(k, v)
	}
	global.GVA_VP.Set("jwt.signing-key", uuid.NewV4().String())
	return global.GVA_VP.WriteConfig()
}

// initMsqlDB 创建数据库并初始化 mysql
// Author [piexlmax](https://github.com/piexlmax)
// Author [SliverHorn](https://github.com/SliverHorn)
// Author: [songzhibin97](https://github.com/songzhibin97)
func (initDBService *InitDBService) initMsqlDB(conf request.InitDB) error {
	dsn := conf.MysqlEmptyDsn()
	createSql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;", conf.DBName)
	if err := initDBService.createDatabase(dsn, "mysql", createSql); err != nil {
		return err
	} // 创建数据库

	mysqlConfig := conf.ToMysqlConfig()
	if mysqlConfig.Dbname == "" {
		return nil
	} // 如果没有数据库名，则跳出初始化数据

	if db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       mysqlConfig.Dsn(), // DSN data source name
		DefaultStringSize:         191,               // string 类型字段的默认长度
		SkipInitializeWithVersion: true,              // 根据版本自动配置
	}), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}); err != nil {
		return nil
	} else {
		global.GVA_DB = db
	}

	if err := initDBService.initTables(); err != nil {
		global.GVA_DB = nil
		return err
	}

	if err := initDBService.initMysqlData(); err != nil {
		global.GVA_DB = nil
		return err
	}

	if err := initDBService.writeMysqlConfig(mysqlConfig); err != nil {
		return err
	}

	global.GVA_CONFIG.AutoCode.Root, _ = filepath.Abs("..")
	return nil
}

// initData mysql 初始化数据
// Author [SliverHorn](https://github.com/SliverHorn)
// Author: [songzhibin97](https://github.com/songzhibin97)
func (initDBService *InitDBService) initMysqlData() error {
	return model.MysqlDataInitialize(
		system2.Api,
		system2.User,
		system2.Casbin,
		system2.BaseMenu,
		system2.Authority,
		system2.Dictionary,
		system2.UserAuthority,
		system2.DataAuthorities,
		system2.AuthoritiesMenus,
		system2.DictionaryDetail,
		system2.ViewAuthorityMenuMysql,
		example.FileMysql,
	)
}

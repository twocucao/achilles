package system

import (
	"achilles/initialize/source/example"
	system2 "achilles/initialize/source/system"
	"path/filepath"

	"achilles/config"
	"achilles/global"
	model "achilles/models/system"
	"achilles/models/system/request"
	"achilles/utils"

	uuid "github.com/satori/go.uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// writePgsqlConfig pgsql 回写配置
// Author [SliverHorn](https://github.com/SliverHorn)
func (initDBService *InitDBService) writePgsqlConfig(pgsql config.Pgsql) error {
	global.GVA_CONFIG.System.DbType = "pgsql"
	global.GVA_CONFIG.Pgsql = pgsql
	cs := utils.StructToMap(global.GVA_CONFIG)
	for k, v := range cs {
		global.GVA_VP.Set(k, v)
	}
	global.GVA_VP.Set("jwt.signing-key", uuid.NewV4().String())
	return global.GVA_VP.WriteConfig()
}

func (initDBService *InitDBService) initPgsqlDB(conf request.InitDB) error {
	dsn := conf.PgsqlEmptyDsn()
	createSql := "CREATE DATABASE " + conf.DBName
	if err := initDBService.createDatabase(dsn, "pgx", createSql); err != nil {
		return err
	} // 创建数据库

	pgsqlConfig := conf.ToPgsqlConfig()
	if pgsqlConfig.Dbname == "" {
		return nil
	} // 如果没有数据库名，则跳出初始化数据

	if db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  pgsqlConfig.Dsn(), // DSN data source name
		PreferSimpleProtocol: false,
	}), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}); err != nil {
		return nil
	} else {
		global.GVA_DB = db
	}

	if err := initDBService.initTables(); err != nil {
		global.GVA_DB = nil
		return err
	}

	if err := initDBService.initPgsqlData(); err != nil {
		global.GVA_DB = nil
		return err
	}

	if err := initDBService.writePgsqlConfig(pgsqlConfig); err != nil {
		return err
	}

	global.GVA_CONFIG.AutoCode.Root, _ = filepath.Abs("..")
	return nil
}

// initPgsqlData pgsql 初始化数据
// Author [SliverHorn](https://github.com/SliverHorn)
func (initDBService *InitDBService) initPgsqlData() error {
	return model.PgsqlDataInitialize(
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
		system2.ViewAuthorityMenuPostgres,

		example.FilePgsql,
	)
}

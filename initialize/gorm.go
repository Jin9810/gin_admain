package initialize

import (
	"gin-vue-admin-STL/entity"
	"gin-vue-admin-STL/global"
	"gorm.io/gorm"
	"os"
)

// Gorm todo:适配所有数据库
func Gorm() *gorm.DB {
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}

// RegisterTables 初始化表
// todo: 适配所有功能
func RegisterTables() {
	db := global.GVA_DB
	err := db.AutoMigrate(
		entity.User{},
	)
	if err != nil {
		os.Exit(0)
	}
}

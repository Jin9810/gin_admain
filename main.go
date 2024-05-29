package main

import (
	"gin-vue-admin-STL/core"
	"gin-vue-admin-STL/global"
	"gin-vue-admin-STL/initialize"
)

func main() {
	global.GVA_VP = core.Viper()
	// todo: 如何配置 config
	global.GVA_DB = initialize.Gorm()
	if global.GVA_DB != nil {
		initialize.RegisterTables()
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}
}

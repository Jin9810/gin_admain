package initialize

import (
	"gin-vue-admin-STL/global"
	"gorm.io/gorm/logger"

	// 短划线表名不使用这个包的任何函数和变量
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GormMysql() *gorm.DB {
	m := global.GVA_CONFIG.Mysql
	if m.Dbname == "" {
		return nil
	}
	// 这个Config用的是 gorm.io/driver/mysql 中的 config
	mysqlConfig := mysql.Config{
		DSN: m.Dsn(),
	}

	// todo: 配置项没设置
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}); err != nil {
		return nil
	} else {
		// todo: 没有设置字符集
		// 建表时的额外设置
		db.InstanceSet("gorm:table_options", "ENGINE="+m.Engine)
		// todo: 设置表的最大闲置和最大连接数
		return db
	}
}

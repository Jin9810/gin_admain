package global

import (
	"gin-vue-admin-STL/config"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	GVA_DB     *gorm.DB
	GVA_VP     *viper.Viper
	GVA_CONFIG config.Config
)

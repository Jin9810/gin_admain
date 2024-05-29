package config

// DsnProvider 为了上层可以调用 Dsn 方法
//type DsnProvider interface {
//	Dsn() string
//}

type GeneralDB struct {
	Config   string `mapstructure:"config"`
	Port     string `mapstructure:"port"`
	Dbname   string `mapstructure:"db-name"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Engine   string `mapstructure:"engine"` // 数据库默认引擎 InnoDB
}

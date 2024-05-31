package config

type System struct {
	DbType string `mapstructure:"db-type"`
	Addr   int    `mapstructure:"addr"`
}

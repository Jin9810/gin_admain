package config

type Config struct {
	System System `mapstructure:"system"`
	// gorm
	Mysql Mysql `mapstructure:"mysql"`
}

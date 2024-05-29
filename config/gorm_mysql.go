package config

// Mysql 为了保持层级的一致，所以只内嵌了类型声明
type Mysql struct {
	GeneralDB `mapstructure:",squash"` // 将嵌入的结构体字段展平
}

// Dsn 连接数据库的所有信息，想要 Mysql 类型的结构体能够使用这个方法，需要在内嵌的类型里面实现一个接口
// Dsn 连接数据库的所有信息
// username:password@tcp(host:port)/databaseName?parameter=value
func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
}

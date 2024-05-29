package internal

// Gorm 这是为什么？方便解耦？
var Gorm = new(_gorm)

type _gorm struct{}

//func (g *_gorm) Config(prefix string, singular bool) *gorm.Config {
//	var general config.GeneralDB
//	switch global.GVA_CONFIG.System.Dbtype {
//	case "mysql": general = global.GVA_CONFIG.Mysql.GeneralDB
//	default:
//		general = global.GVA_CONFIG.Mysql.GeneralDB
//	}
//	return &gorm.Config{
//		Logger: logger.New(NewWriter(general)),
//	}
//}

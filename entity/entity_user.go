package entity

import (
	"gin-vue-admin-STL/global"
	"github.com/gofrs/uuid/v5"
)

// todo: 完善用户列表
type User struct {
	global.GVA_MODEL
	UUID      uuid.UUID `json:"uuid" gorm:"index"`                                                       // 使用名称空间和名称来生成唯一的 UUID
	Username  string    `json:"userName" gorm:"index"`                                                   // 用户登录名
	Password  string    `json:"-"`                                                                       // 用户密码
	NickName  string    `json:"nickName" gorm:"default:系统用户"`                                            // 用户昵称
	HeaderIMG string    `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg"` // 用户头像
	Phone     string    `json:"phone"`                                                                   // 用户电话
	Email     string    `json:"email"`                                                                   // 邮箱
}

func (User) TableName() string {
	return "users"
}

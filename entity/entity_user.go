package entity

import (
	"gin-vue-admin-STL/global"
	"github.com/gofrs/uuid/v5"
)

// todo: 完善用户列表
type User struct {
	global.GVA_MODEL
	UUID      uuid.UUID // 使用名称空间和名称来生成唯一的 UUID
	Username  string    // 用户登录名
	Password  string    // 用户密码
	NickName  string    // 用户昵称
	HeaderIMG string    // 用户头像
	Phone     string    // 用户电话
	Email     string    // 邮箱
}

func (User) TableName() string {
	return "users"
}

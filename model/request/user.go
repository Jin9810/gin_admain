package request

// Register 注册
type Register struct {
	Username  string `json:"username" example:"用户名"`
	Password  string `json:"password" example:"密码"`
	NickName  string `json:"nickname" example:"昵称"`
	HeaderImg string `json:"headerImg" example:"头像链接"`
	Phone     string `json:"phone" example:"电话号码"`
	Email     string `json:"email" example:"电子邮箱"`
}

// Login 登录
type Login struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
}

// ChangePassword 修改密码
type ChangePassword struct {
	Phone       string `json:"phone"`
	Password    string `json:"password"`    // 密码
	NewPassword string `json:"newPassword"` // 新密码
}

// ChangeUserInfo 改变用户信息
type ChangeUserInfo struct {
	ID        uint   `gorm:"primarykey"`                                                                           // 主键ID
	NickName  string `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                            // 用户昵称
	Phone     string `json:"phone"  gorm:"comment:用户手机号"`                                                          // 用户手机号
	Email     string `json:"email"  gorm:"comment:用户邮箱"`                                                           // 用户邮箱
	HeaderImg string `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"` // 用户头像
}

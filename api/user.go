package api

import (
	"gin-vue-admin-STL/entity"
	"gin-vue-admin-STL/model/system/request"
	"gin-vue-admin-STL/model/system/response"
	"gin-vue-admin-STL/service"
	"github.com/gin-gonic/gin"
)

type BaseApi struct{}

// Register
// @Router  /user/admin_register [post]
// 注册
func (b *BaseApi) Register(c *gin.Context) {
	var r request.Register
	err := c.ShouldBindBodyWithJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	user := &entity.User{Username: r.Username, NickName: r.NickName, Password: r.Password, HeaderIMG: r.HeaderImg, Phone: r.Phone, Email: r.Email}

	userReturn, err := service.ServiceGroupApp.Register(*user)
	if err != nil {
		response.FailWithDetailed(response.UserResponse{User: *userReturn}, "注册失败", c)
		return
	}
	response.OkWithDetailed(response.UserResponse{User: *userReturn}, "注册成功", c)
}

// ChangePassword
// @Router  /user/changePassword [post]
// 修改密码
func (b *BaseApi) ChangePassword(c *gin.Context) {
	var cp request.ChangePassword
	err := c.ShouldBindBodyWithJSON(&cp)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	u := &entity.User{Phone: cp.Phone, Password: cp.Password}
	_, err = service.ServiceGroupApp.ChangePassWord(u, cp.NewPassword)
	if err != nil {
		response.FailWithMessage("修改失败，原密码不正确", c)
	}
	response.OkWithMessage("修改成功", c)
}

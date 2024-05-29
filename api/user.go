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
		response.FailWithDetailed(response.UserResponse{User: userReturn}, "注册失败", c)
		return
	}
	response.FailWithDetailed(response.UserResponse{User: userReturn}, "注册成功", c)
}

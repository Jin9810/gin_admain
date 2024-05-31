package api

import (
	"gin-vue-admin-STL/entity"
	"gin-vue-admin-STL/global"
	"gin-vue-admin-STL/model/request"
	"gin-vue-admin-STL/model/response"
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

	user := entity.User{Username: r.Username, NickName: r.NickName, Password: r.Password, HeaderIMG: r.HeaderImg, Phone: r.Phone, Email: r.Email}

	userReturn, err := service.ServiceGroupApp.Register(user)
	if err != nil {
		response.FailWithDetailed(response.UserResponse{User: userReturn}, "注册失败", c)
		return
	}
	response.OkWithDetailed(response.UserResponse{User: userReturn}, "注册成功", c)
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

// DeleteUser
// @Router /user/deleteUser
// 删除用户
func (b *BaseApi) DeleteUser(c *gin.Context) {
	var id request.GetById
	err := c.ShouldBindJSON(&id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = service.ServiceGroupApp.DeleteUser(id.ID)
	if err != nil {
		response.FailWithMessage("删除失败", c)
	}
	response.OkWithMessage("删除成功", c)
}

// SetUserInfo
// @Router /user/setUserInfo
// 设置用户信息
func (b *BaseApi) SetUserInfo(c *gin.Context) {
	var userinfo request.ChangeUserInfo
	err := c.ShouldBindJSON(&userinfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = service.ServiceGroupApp.SetUserInfo(entity.User{
		GVA_MODEL: global.GVA_MODEL{
			ID: userinfo.ID,
		},
		NickName:  userinfo.NickName,
		HeaderIMG: userinfo.HeaderImg,
		Phone:     userinfo.Phone,
		Email:     userinfo.Email,
	})
	if err != nil {
		response.FailWithMessage("设置失败", c)
		return
	}
	response.OkWithMessage("设置成功", c)
}

// ResetPassword
// @Router /user/resetPassWord
// 重置密码
func (b *BaseApi) ResetPassword(c *gin.Context) {
	var id request.GetById
	err := c.ShouldBindJSON(&id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = service.ServiceGroupApp.ResetPassword(id.ID)
	if err != nil {
		response.FailWithMessage("重置失败！"+err.Error(), c)
		return
	}
	response.OkWithMessage("重置成功", c)
}

// GetUserList
// @Router /user/getUserList
// 获取用户列表
func (b *BaseApi) GetUserList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := service.ServiceGroupApp.GetUserInfoList(pageInfo)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

func (b *BaseApi) GetUserInfo(c *gin.Context) {
	var id request.GetById
	err := c.ShouldBindJSON(&id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	User, err := service.ServiceGroupApp.GetUserInfo(id.ID)
	if err != nil {
		response.FailWithMessage("获取失败", c)
	}
	response.OkWithDetailed(gin.H{"userInfo": User}, "获取成功", c)
}

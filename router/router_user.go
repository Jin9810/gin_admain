package router

import (
	"gin-vue-admin-STL/api"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	baseApi := api.GroupApp.BaseApi
	UserRouter.POST("admin_register", baseApi.Register)
	UserRouter.PUT("changePassword", baseApi.ChangePassword)
	UserRouter.PUT("setUserInfo", baseApi.SetUserInfo)
	UserRouter.PUT("resetPassword", baseApi.ResetPassword)
	UserRouter.GET("getUserList", baseApi.GetUserList)
	UserRouter.GET("getUserInfo", baseApi.GetUserInfo)
	UserRouter.DELETE("deleteUser", baseApi.DeleteUser)
}

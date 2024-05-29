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
}

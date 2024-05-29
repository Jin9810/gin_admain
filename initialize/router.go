package initialize

import (
	"gin-vue-admin-STL/router"
	"github.com/gin-gonic/gin"
)

// Routers 初始化总路由，gin.Engine 管理所有路由
func Routers() *gin.Engine {
	Router := gin.New()

	// 内置中间件，捕获任何在 HTTP 请求期间发生的 panic
	Router.Use(gin.Recovery())

	Group := Router.Group("")

	router.GroupApp.UserRouter.InitUserRouter(Group)

	return Router
}

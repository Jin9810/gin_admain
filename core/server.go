package core

import (
	"fmt"
	"gin-vue-admin-STL/global"
	"gin-vue-admin-STL/initialize"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() error {
	// 初始化路由
	Router := initialize.Routers()
	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := initServer(address, Router)
	return s.ListenAndServe()
}

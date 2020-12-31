package router

import (
	"github.com/gin-gonic/gin"
	"github.com/marx88/gin-admin/config"
	"github.com/marx88/gin-admin/gan/route"
)

// Root 根路由组
var Root = route.NewGroup()

func init() {
	// 生成路由组根路由
	Root.Init = func(router gin.IRouter) gin.IRouter {
		initRootMiddlewares(router) // 全局中间件必须挂在*gin.Eninge下
		return router.Group("")
	}
}

// initRootMiddlewares 全局中间件
func initRootMiddlewares(router gin.IRouter) {
	router.Use(gin.Recovery())
	if config.App.Debug {
		router.Use(gin.Logger())
	}
}

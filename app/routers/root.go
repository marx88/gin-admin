package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/marx88/gin-admin/app/routers/admin"
	"github.com/marx88/gin-admin/app/routers/index"
	"github.com/marx88/gin-admin/config"
	"github.com/marx88/gin-admin/gan/lib/router"
)

// Root 根路由
var Root = router.NewGroup("")

func init() {
	// 添加中间件
	if config.App.Debug {
		Root.AddMiddleware(gin.Logger())
	}
	Root.AddMiddleware(gin.Recovery())

	// 添加子模块路由组
	Root.AddSubGroup(admin.Group, index.Group)
}

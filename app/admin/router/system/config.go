package system

import (
	"github.com/gin-gonic/gin"
	"github.com/marx88/gin-admin/app/admin/controller/system/controller"
	"github.com/marx88/gin-admin/app/admin/middleware"
	"github.com/marx88/gin-admin/common/enum/businesstype"
)

func init() {
	Root.Put(configRouters)
}

// configRouters 配置路由
func configRouters(i gin.IRouter) {
	r := i.Group("/system/config")
	c := new(controller.ConfigController)
	r.GET(
		"/list",
		middleware.Auth("system:config:list"),
		c.List,
	)
	r.POST(
		"",
		middleware.Auth("system:config:add"),
		middleware.BusinessLog("参数管理", businesstype.INSERT),
		c.Add,
	)
	r.GET(
		"",
		middleware.Auth("system:config:query"),
		c.Query,
	)
	r.PUT(
		"",
		middleware.Auth("system:config:edit"),
		middleware.BusinessLog("参数管理", businesstype.UPDATE),
		c.Edit,
	)
	r.DELETE(
		"",
		middleware.Auth("system:config:remove"),
		middleware.BusinessLog("参数管理", businesstype.DELETE),
		c.Remove,
	)
	r.GET(
		"/export",
		middleware.Auth("system:config:export"),
		middleware.BusinessLog("参数管理", businesstype.EXPORT),
		c.Export,
	)
	r.GET(
		"/configKey",
		c.GetConfigKey,
	)
	r.DELETE(
		"/clearCache",
		middleware.Auth("system:config:remove"),
		middleware.BusinessLog("参数管理", businesstype.CLEAN),
		c.Remove,
	)
}

package system

import (
	"github.com/gin-gonic/gin"
	"github.com/marx88/gin-admin/app/admin/controller/system/controller"
	"github.com/marx88/gin-admin/app/admin/middleware"
	"github.com/marx88/gin-admin/common/enum/businesstype"
)

func init() {
	Root.Put(userRouters)
}

// userRouters 用户路由
func userRouters(i gin.IRouter) {
	r := i.Group("/system/user")
	c := new(controller.UserController)
	r.GET(
		"/list",
		middleware.Auth("system:user:list"),
		c.List,
	)
	r.POST(
		"",
		middleware.Auth("system:user:add"),
		middleware.BusinessLog("用户管理", businesstype.INSERT),
		c.Add,
	)
	r.GET(
		"",
		middleware.Auth("system:user:query"),
		c.Query,
	)
	r.PUT(
		"",
		middleware.Auth("system:user:edit"),
		middleware.BusinessLog("用户管理", businesstype.UPDATE),
		c.Edit,
	)
	r.DELETE(
		"",
		middleware.Auth("system:user:remove"),
		middleware.BusinessLog("用户管理", businesstype.DELETE),
		c.Remove,
	)
	r.GET(
		"/export",
		middleware.Auth("system:user:export"),
		middleware.BusinessLog("用户管理", businesstype.EXPORT),
		c.Export,
	)
	r.GET(
		"importTemplate",
		c.ImportTemplate,
	)
	r.POST(
		"/importData",
		middleware.Auth("system:user:import"),
		middleware.BusinessLog("用户管理", businesstype.IMPORT),
		c.Import,
	)
	r.PUT(
		"/resetPwd",
		middleware.Auth("system:user:resetPwd"),
		middleware.BusinessLog("用户管理", businesstype.UPDATE),
		c.ResetPwd,
	)
	r.PUT(
		"/changeStatus",
		middleware.Auth("system:user:edit"),
		middleware.BusinessLog("用户管理", businesstype.UPDATE),
		c.ChangeStatus,
	)
}

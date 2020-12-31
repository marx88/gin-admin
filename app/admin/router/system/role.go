package system

import (
	"github.com/gin-gonic/gin"
	"github.com/marx88/gin-admin/app/admin/controller/system/controller"
	"github.com/marx88/gin-admin/app/admin/middleware"
	"github.com/marx88/gin-admin/common/enum/businesstype"
)

func init() {
	Root.Put(roleRouters)
}

// roleRouters 角色路由
func roleRouters(i gin.IRouter) {
	r := i.Group("/system/role")
	c := new(controller.RoleController)
	r.GET(
		"/list",
		middleware.Auth("system:role:list"),
		c.List,
	)
	r.POST(
		"",
		middleware.Auth("system:role:add"),
		middleware.BusinessLog("角色管理", businesstype.INSERT),
		c.Add,
	)
	r.GET(
		"",
		middleware.Auth("system:role:query"),
		c.Query,
	)
	r.PUT(
		"",
		middleware.Auth("system:role:edit"),
		middleware.BusinessLog("角色管理", businesstype.UPDATE),
		c.Edit,
	)
	r.DELETE(
		"",
		middleware.Auth("system:role:remove"),
		middleware.BusinessLog("角色管理", businesstype.DELETE),
		c.Remove,
	)
	r.GET(
		"/export",
		middleware.Auth("system:role:export"),
		middleware.BusinessLog("角色管理", businesstype.EXPORT),
		c.Export,
	)
	r.PUT(
		"/dataScope",
		middleware.Auth("system:role:edit"),
		middleware.BusinessLog("角色管理", businesstype.UPDATE),
		c.DataScope,
	)
	r.PUT(
		"/changeStatus",
		middleware.Auth("system:role:edit"),
		middleware.BusinessLog("角色管理", businesstype.UPDATE),
		c.ChangeStatus,
	)
	r.PUT(
		"/optionselect",
		middleware.Auth("system:role:query"),
		c.OptionSelect,
	)
}

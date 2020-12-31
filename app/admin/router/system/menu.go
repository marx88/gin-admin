package system

import (
	"github.com/gin-gonic/gin"
	"github.com/marx88/gin-admin/app/admin/controller/system/controller"
	"github.com/marx88/gin-admin/app/admin/middleware"
	"github.com/marx88/gin-admin/common/enum/businesstype"
)

func init() {
	Root.Put(menuRouters)
}

// menuRouters 菜单路由
func menuRouters(i gin.IRouter) {
	r := i.Group("/system/menu")
	c := new(controller.MenuController)
	r.GET(
		"/list",
		middleware.Auth("system:menu:list"),
		c.List,
	)
	r.POST(
		"",
		middleware.Auth("system:menu:add"),
		middleware.BusinessLog("菜单管理", businesstype.INSERT),
		c.Add,
	)
	r.GET(
		"",
		middleware.Auth("system:menu:query"),
		c.Query,
	)
	r.PUT(
		"",
		middleware.Auth("system:menu:edit"),
		middleware.BusinessLog("菜单管理", businesstype.UPDATE),
		c.Edit,
	)
	r.DELETE(
		"",
		middleware.Auth("system:menu:remove"),
		middleware.BusinessLog("菜单管理", businesstype.DELETE),
		c.Remove,
	)
	r.GET(
		"/treeselect",
		c.TreeSelect,
	)
	r.GET(
		"/roleMenuTreeselect",
		c.RoleMenuTreeselect,
	)
}

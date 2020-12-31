package system

import (
	"github.com/gin-gonic/gin"
	"github.com/marx88/gin-admin/app/admin/controller/system/controller"
	"github.com/marx88/gin-admin/app/admin/middleware"
	"github.com/marx88/gin-admin/common/enum/businesstype"
)

func init() {
	Root.Put(deptRouters)
}

// deptRouters 部门路由
func deptRouters(i gin.IRouter) {
	r := i.Group("/system/dept")
	c := new(controller.DeptController)
	r.GET(
		"/list",
		middleware.Auth("system:dept:list"),
		c.List,
	)
	r.POST(
		"",
		middleware.Auth("system:dept:add"),
		middleware.BusinessLog("部门管理", businesstype.INSERT),
		c.Add,
	)
	r.GET(
		"",
		middleware.Auth("system:dept:query"),
		c.Query,
	)
	r.PUT(
		"",
		middleware.Auth("system:dept:edit"),
		middleware.BusinessLog("部门管理", businesstype.UPDATE),
		c.Edit,
	)
	r.DELETE(
		"",
		middleware.Auth("system:dept:remove"),
		middleware.BusinessLog("部门管理", businesstype.DELETE),
		c.Remove,
	)
	r.GET(
		"/list/exclude",
		middleware.Auth("system:dept:list"),
		c.ExcludeChild,
	)
	r.GET(
		"/treeselect",
		c.TreeSelect,
	)
	r.GET(
		"/roleDeptTreeselect",
		c.RoleDeptTreeSelect,
	)
}

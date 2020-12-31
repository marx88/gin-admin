package monitor

import (
	"github.com/gin-gonic/gin"
	"github.com/marx88/gin-admin/app/admin/controller/monitor/controller"
	"github.com/marx88/gin-admin/app/admin/middleware"
	"github.com/marx88/gin-admin/common/enum/businesstype"
)

func init() {
	Root.Put(monitorRouters)
}

// monitorRouters 操作日志路由
func monitorRouters(i gin.IRouter) {
	r := i.Group("/monitor/operlog")
	c := new(controller.OperLogController)
	r.GET(
		"/list",
		middleware.Auth("system:operlog:list"),
		c.List,
	)
	r.DELETE(
		"",
		middleware.Auth("system:operlog:remove"),
		c.Remove,
	)
	r.GET(
		"/export",
		middleware.Auth("system:operlog:export"),
		middleware.BusinessLog("操作日志", businesstype.EXPORT),
		c.Export,
	)
	r.DELETE(
		"clean",
		middleware.Auth("system:operlog:remove"),
		middleware.BusinessLog("操作日志", businesstype.CLEAN),
		c.Clean,
	)
}

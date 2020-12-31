package system

import (
	"github.com/gin-gonic/gin"
	"github.com/marx88/gin-admin/app/admin/controller/system/controller"
	"github.com/marx88/gin-admin/app/admin/middleware"
	"github.com/marx88/gin-admin/common/enum/businesstype"
)

func init() {
	Root.Put(noticeRouters)
}

// noticeRouters 通知公告路由
func noticeRouters(i gin.IRouter) {
	r := i.Group("/system/notice")
	c := new(controller.NoticeController)
	r.GET(
		"/list",
		middleware.Auth("system:notice:list"),
		c.List,
	)
	r.POST(
		"",
		middleware.Auth("system:notice:add"),
		middleware.BusinessLog("通知公告", businesstype.INSERT),
		c.Add,
	)
	r.GET(
		"",
		middleware.Auth("system:notice:query"),
		c.Query,
	)
	r.PUT(
		"",
		middleware.Auth("system:notice:edit"),
		middleware.BusinessLog("通知公告", businesstype.UPDATE),
		c.Edit,
	)
	r.DELETE(
		"",
		middleware.Auth("system:notice:remove"),
		middleware.BusinessLog("通知公告", businesstype.DELETE),
		c.Remove,
	)
}

package system

import (
	"github.com/gin-gonic/gin"
	"github.com/marx88/gin-admin/app/admin/controller/system/controller"
	"github.com/marx88/gin-admin/app/admin/middleware"
	"github.com/marx88/gin-admin/common/enum/businesstype"
)

func init() {
	Root.Put(postRouters)
}

// postRouters 岗位管理路由
func postRouters(i gin.IRouter) {
	r := i.Group("/system/post")
	c := new(controller.PostController)
	r.GET(
		"/list",
		middleware.Auth("system:post:list"),
		c.List,
	)
	r.POST(
		"",
		middleware.Auth("system:post:add"),
		middleware.BusinessLog("岗位管理", businesstype.INSERT),
		c.Add,
	)
	r.GET(
		"",
		middleware.Auth("system:post:query"),
		c.Query,
	)
	r.PUT(
		"",
		middleware.Auth("system:post:edit"),
		middleware.BusinessLog("岗位管理", businesstype.UPDATE),
		c.Edit,
	)
	r.DELETE(
		"",
		middleware.Auth("system:post:remove"),
		middleware.BusinessLog("岗位管理", businesstype.DELETE),
		c.Remove,
	)
	r.GET(
		"/export",
		middleware.Auth("system:post:export"),
		middleware.BusinessLog("岗位管理", businesstype.EXPORT),
		c.Export,
	)
	r.GET(
		"/optionselect",
		c.OptionSelect,
	)
}

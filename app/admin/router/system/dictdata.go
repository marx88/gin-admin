package system

import (
	"github.com/gin-gonic/gin"
	"github.com/marx88/gin-admin/app/admin/controller/system/controller"
	"github.com/marx88/gin-admin/app/admin/middleware"
	"github.com/marx88/gin-admin/common/enum/businesstype"
)

func init() {
	Root.Put(dictDataRouters)
}

// dictDataRouters 字典数据路由
func dictDataRouters(i gin.IRouter) {
	r := i.Group("/system/dict/data")
	c := new(controller.DictDataController)
	r.GET(
		"/list",
		middleware.Auth("system:dict:list"),
		c.List,
	)
	r.POST(
		"",
		middleware.Auth("system:dict:add"),
		middleware.BusinessLog("字典数据", businesstype.INSERT),
		c.Add,
	)
	r.GET(
		"",
		middleware.Auth("system:dict:query"),
		c.Query,
	)
	r.PUT(
		"",
		middleware.Auth("system:dict:edit"),
		middleware.BusinessLog("字典数据", businesstype.UPDATE),
		c.Edit,
	)
	r.DELETE(
		"",
		middleware.Auth("system:dict:remove"),
		middleware.BusinessLog("字典数据", businesstype.DELETE),
		c.Remove,
	)
	r.GET(
		"/export",
		middleware.Auth("system:dict:export"),
		middleware.BusinessLog("字典数据", businesstype.EXPORT),
		c.Export,
	)
	r.GET(
		"/type",
		c.DictType,
	)
}

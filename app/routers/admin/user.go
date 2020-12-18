package admin

import (
	"github.com/gin-gonic/gin"
	controller "github.com/marx88/gin-admin/app/api/admin"
	"github.com/marx88/gin-admin/app/middleware"
	"github.com/marx88/gin-admin/gan/lib/router"
)

var userGroup = router.NewGroup("user")

func init() {
	userGroup.AddRegister(func(r gin.IRouter) {
		c := new(controller.UserController)
		r.GET("index", c.Index)
		r.GET("create", c.Create)
		r.POST("save", c.Save)
		r.GET("edit", c.Edit)
		r.POST("update", c.Update)
		r.POST("delete", c.Delete)
	})
	userGroup.AddMiddleware(middleware.Logger("user控制器"))
	Group.AddSubGroup(userGroup)
}

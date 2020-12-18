package index

import (
	"github.com/gin-gonic/gin"
	controller "github.com/marx88/gin-admin/app/api/index"
	"github.com/marx88/gin-admin/gan/lib/router"
)

var administratorGroup = router.NewGroup("admin")

func init() {
	administratorGroup.AddRegister(func(r gin.IRouter) {
		c := new(controller.AdministratorController)
		r.GET("index", c.Index)
		r.GET("create", c.Create)
		r.POST("save", c.Save)
		r.GET("edit", c.Edit)
		r.POST("update", c.Update)
		r.POST("delete", c.Delete)
	})
	Group.AddSubGroup(administratorGroup)
}

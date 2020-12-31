package system

import (
	"github.com/gin-gonic/gin"
	"github.com/marx88/gin-admin/app/admin/controller/system/controller"
)

func init() {
	Root.Put(loginRouters)
}

// loginRouters 登录路由
func loginRouters(i gin.IRouter) {
	r := i.Group("")
	c := new(controller.LoginController)
	r.POST(
		"/login",
		c.Login,
	)
	r.GET(
		"/getInfo",
		c.Info,
	)
	r.GET(
		"/getRouters",
		c.Routers,
	)
}

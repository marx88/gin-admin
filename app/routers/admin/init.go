package admin

import (
	"github.com/marx88/gin-admin/app/middleware"
	"github.com/marx88/gin-admin/gan/lib/router"
)

// Group admin路由组
var Group = router.NewGroup("admin")

func init() {
	Group.AddMiddleware(middleware.Logger("admin模块"))
}

package index

import (
	"github.com/marx88/gin-admin/app/middleware"
	"github.com/marx88/gin-admin/gan/lib/router"
)

// Group index路由组
var Group = router.NewGroup("index")

func init() {
	Group.AddMiddleware(middleware.Logger("index模块"))
}

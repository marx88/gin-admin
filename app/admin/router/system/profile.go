package system

import (
	"github.com/gin-gonic/gin"
	"github.com/marx88/gin-admin/app/admin/controller/system/controller"
	"github.com/marx88/gin-admin/app/admin/middleware"
	"github.com/marx88/gin-admin/common/enum/businesstype"
)

func init() {
	Root.Put(profileRouters)
}

// profileRouters 个人信息路由
func profileRouters(i gin.IRouter) {
	r := i.Group("/system/user/profile")
	c := new(controller.ProfileController)
	r.GET(
		"",
		c.Profile,
	)
	r.PUT(
		"",
		middleware.BusinessLog("个人信息", businesstype.UPDATE),
		c.UpdateProfile,
	)
	r.PUT(
		"/updatePwd",
		middleware.BusinessLog("个人信息", businesstype.UPDATE),
		c.UpdatePwd,
	)
	r.POST(
		"/avatar",
		middleware.BusinessLog("用户头像", businesstype.UPDATE),
		c.Avatar,
	)
}

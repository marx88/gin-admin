package common

import (
	"github.com/gin-gonic/gin"
	"github.com/marx88/gin-admin/app/admin/controller/common/controller"
)

func init() {
	Root.Put(commonRouters)
}

// commonRouters 通用路由
func commonRouters(i gin.IRouter) {
	commonGroup := i.Group("/common")
	commonController := new(controller.CommonController)
	commonGroup.GET("/download", commonController.Download)
	commonGroup.POST("/upload", commonController.Upload)
	commonGroup.GET("/download/resource", commonController.Resource)
}

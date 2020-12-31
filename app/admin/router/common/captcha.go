package common

import (
	"github.com/gin-gonic/gin"
	"github.com/marx88/gin-admin/app/admin/controller/common/controller"
)

func init() {
	Root.Put(captchaRouters)
}

// captchaRouters 验证码
func captchaRouters(i gin.IRouter) {
	captchaController := new(controller.CaptchaController)
	i.GET("/captchaImage", captchaController.Generate)
}

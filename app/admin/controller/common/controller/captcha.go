package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/marx88/gin-admin/gan/captcha"
	"github.com/marx88/gin-admin/gan/response"
)

// CaptchaController 验证码控制器
type CaptchaController struct{}

// Generate 生成验证码
func (cp *CaptchaController) Generate(c *gin.Context) {
	resp := response.New()

	uuid, img, err := captcha.Get()
	if err != nil {
		c.Error(err)
		resp.ErrorMsg(c, err.Error())
		return
	}

	resp.Put("uuid", uuid)
	resp.Put("img", img)
	resp.Success(c)
}

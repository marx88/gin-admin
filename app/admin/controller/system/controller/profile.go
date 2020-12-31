package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/marx88/gin-admin/gan/response"
)

// ProfileController 个人信息控制器
type ProfileController struct{}

// Profile 个人信息
func (cp *ProfileController) Profile(c *gin.Context) {
	response.Success(c)
}

// UpdateProfile 修改用户
func (cp *ProfileController) UpdateProfile(c *gin.Context) {
	response.Success(c)
}

// UpdatePwd 重置密码
func (cp *ProfileController) UpdatePwd(c *gin.Context) {
	response.Success(c)
}

// Avatar 头像上传
func (cp *ProfileController) Avatar(c *gin.Context) {
	response.Success(c)
}

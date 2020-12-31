package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/marx88/gin-admin/gan/response"
)

// UserController 用户控制器
type UserController struct{}

// List 用户列表
func (cp *UserController) List(c *gin.Context) {
	response.SuccessPage(c, 0, nil)
}

// Add 新增用户
func (cp *UserController) Add(c *gin.Context) {
	response.Success(c)
}

// Query 用户详情
func (cp *UserController) Query(c *gin.Context) {
	response.Success(c)
}

// Edit 更新用户
func (cp *UserController) Edit(c *gin.Context) {
	response.Success(c)
}

// Remove 删除用户
func (cp *UserController) Remove(c *gin.Context) {
	response.Success(c)
}

// Export 导出用户
func (cp *UserController) Export(c *gin.Context) {
	response.Success(c)
}

// ImportTemplate 导入数据的模板
func (cp *UserController) ImportTemplate(c *gin.Context) {
	response.Success(c)
}

// Import 导入数据
func (cp *UserController) Import(c *gin.Context) {
	response.Success(c)
}

// ResetPwd 重置密码
func (cp *UserController) ResetPwd(c *gin.Context) {
	response.Success(c)
}

// ChangeStatus 状态修改
func (cp *UserController) ChangeStatus(c *gin.Context) {
	response.Success(c)
}

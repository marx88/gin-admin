package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/marx88/gin-admin/gan/response"
)

// RoleController 角色控制器
type RoleController struct{}

// List 角色列表
func (cp *RoleController) List(c *gin.Context) {
	response.SuccessPage(c, 0, nil)
}

// Add 新增角色
func (cp *RoleController) Add(c *gin.Context) {
	response.Success(c)
}

// Query 角色详情
func (cp *RoleController) Query(c *gin.Context) {
	response.Success(c)
}

// Edit 更新角色
func (cp *RoleController) Edit(c *gin.Context) {
	response.Success(c)
}

// Remove 删除角色
func (cp *RoleController) Remove(c *gin.Context) {
	response.Success(c)
}

// Export 导出角色
func (cp *RoleController) Export(c *gin.Context) {
	response.Success(c)
}

// DataScope 修改保存数据权限
func (cp *RoleController) DataScope(c *gin.Context) {
	response.Success(c)
}

// ChangeStatus 状态修改
func (cp *RoleController) ChangeStatus(c *gin.Context) {
	response.Success(c)
}

// OptionSelect 获取角色选择框列表
func (cp *RoleController) OptionSelect(c *gin.Context) {
	response.Success(c)
}

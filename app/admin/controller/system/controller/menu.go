package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/marx88/gin-admin/gan/response"
)

// MenuController 菜单控制器
type MenuController struct{}

// List 菜单列表
func (cp *MenuController) List(c *gin.Context) {
	response.SuccessPage(c, 0, nil)
}

// Add 新增菜单
func (cp *MenuController) Add(c *gin.Context) {
	response.Success(c)
}

// Query 菜单详情
func (cp *MenuController) Query(c *gin.Context) {
	response.Success(c)
}

// Edit 更新菜单
func (cp *MenuController) Edit(c *gin.Context) {
	response.Success(c)
}

// Remove 删除菜单
func (cp *MenuController) Remove(c *gin.Context) {
	response.Success(c)
}

// TreeSelect 获取菜单下拉树列表
func (cp *MenuController) TreeSelect(c *gin.Context) {
	response.Success(c)
}

// RoleMenuTreeselect 加载对应角色菜单列表树
func (cp *MenuController) RoleMenuTreeselect(c *gin.Context) {
	response.Success(c)
}

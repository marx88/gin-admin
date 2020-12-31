package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/marx88/gin-admin/gan/response"
)

// DeptController 部门控制器
type DeptController struct{}

// List 部门列表
func (cp *DeptController) List(c *gin.Context) {
	response.SuccessPage(c, 0, nil)
}

// Add 新增部门
func (cp *DeptController) Add(c *gin.Context) {
	response.Success(c)
}

// Query 部门详情
func (cp *DeptController) Query(c *gin.Context) {
	response.Success(c)
}

// Edit 更新部门
func (cp *DeptController) Edit(c *gin.Context) {
	response.Success(c)
}

// Remove 删除部门
func (cp *DeptController) Remove(c *gin.Context) {
	response.Success(c)
}

// ExcludeChild 查询部门列表（排除节点）
func (cp *DeptController) ExcludeChild(c *gin.Context) {
	response.Success(c)
}

// TreeSelect 获取部门下拉树列表
func (cp *DeptController) TreeSelect(c *gin.Context) {
	response.Success(c)
}

// RoleDeptTreeSelect 加载对应角色部门列表树
func (cp *DeptController) RoleDeptTreeSelect(c *gin.Context) {
	response.Success(c)
}

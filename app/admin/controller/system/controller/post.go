package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/marx88/gin-admin/gan/response"
)

// PostController 岗位控制器
type PostController struct{}

// List 岗位列表
func (cp *PostController) List(c *gin.Context) {
	response.SuccessPage(c, 0, nil)
}

// Add 新增岗位
func (cp *PostController) Add(c *gin.Context) {
	response.Success(c)
}

// Query 岗位详情
func (cp *PostController) Query(c *gin.Context) {
	response.Success(c)
}

// Edit 更新岗位
func (cp *PostController) Edit(c *gin.Context) {
	response.Success(c)
}

// Remove 删除岗位
func (cp *PostController) Remove(c *gin.Context) {
	response.Success(c)
}

// Export 导出岗位
func (cp *PostController) Export(c *gin.Context) {
	response.Success(c)
}

// OptionSelect 获取岗位选择框列表
func (cp *PostController) OptionSelect(c *gin.Context) {
	response.Success(c)
}

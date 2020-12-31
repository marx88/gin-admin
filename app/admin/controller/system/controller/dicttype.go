package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/marx88/gin-admin/gan/response"
)

// DictTypeController 字典类型控制器
type DictTypeController struct{}

// List 字典类型列表
func (cp *DictTypeController) List(c *gin.Context) {
	response.SuccessPage(c, 0, nil)
}

// Add 新增字典类型
func (cp *DictTypeController) Add(c *gin.Context) {
	response.Success(c)
}

// Query 字典类型详情
func (cp *DictTypeController) Query(c *gin.Context) {
	response.Success(c)
}

// Edit 更新字典类型
func (cp *DictTypeController) Edit(c *gin.Context) {
	response.Success(c)
}

// Remove 删除字典类型
func (cp *DictTypeController) Remove(c *gin.Context) {
	response.Success(c)
}

// Export 导出字典类型
func (cp *DictTypeController) Export(c *gin.Context) {
	response.Success(c)
}

// ClearCache 清空缓存
func (cp *DictTypeController) ClearCache(c *gin.Context) {
	response.Success(c)
}

// OptionSelect 获取字典选择框列表
func (cp *DictTypeController) OptionSelect(c *gin.Context) {
	response.Success(c)
}

package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/marx88/gin-admin/gan/excel"
	"github.com/marx88/gin-admin/gan/response"
)

// ConfigController 配置控制器
type ConfigController struct{}

// List 配置列表
func (cp *ConfigController) List(c *gin.Context) {
	response.SuccessPage(c, 0, nil)
}

// Add 新增配置
func (cp *ConfigController) Add(c *gin.Context) {
	response.Success(c)
}

// Query 配置详情
func (cp *ConfigController) Query(c *gin.Context) {
	response.Success(c)
}

// Edit 更新配置
func (cp *ConfigController) Edit(c *gin.Context) {
	response.Success(c)
}

// Remove 删除配置
func (cp *ConfigController) Remove(c *gin.Context) {
	response.Success(c)
}

// Export 导出配置
func (cp *ConfigController) Export(c *gin.Context) {
	path, err := excel.Export(nil, "参数数据")
	if err != nil {
		c.Error(err)
		response.ErrorMsg(c, err.Error())
		return
	}

	response.SuccessMsg(c, path)
}

// GetConfigKey 根据参数键名查询参数值
func (cp *ConfigController) GetConfigKey(c *gin.Context) {
	response.Success(c)
}

// ClearCache 清空缓存
func (cp *ConfigController) ClearCache(c *gin.Context) {
	response.Success(c)
}

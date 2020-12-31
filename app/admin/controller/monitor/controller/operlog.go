package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/marx88/gin-admin/gan/excel"
	"github.com/marx88/gin-admin/gan/response"
)

// OperLogController 操作日志控制器
type OperLogController struct{}

// List 操作日志列表
func (cp *OperLogController) List(c *gin.Context) {
	response.SuccessPage(c, 0, nil)
}

// Remove 删除操作日志
func (cp *OperLogController) Remove(c *gin.Context) {
	response.Success(c)
}

// Export 导出操作日志
func (cp *OperLogController) Export(c *gin.Context) {
	path, err := excel.Export(nil, "参数数据")
	if err != nil {
		c.Error(err)
		response.ErrorMsg(c, err.Error())
		return
	}

	response.SuccessMsg(c, path)
}

// Clean 清空日志
func (cp *OperLogController) Clean(c *gin.Context) {
	response.Success(c)
}

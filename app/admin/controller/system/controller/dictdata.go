package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/marx88/gin-admin/gan/response"
)

// DictDataController 字典数据控制器
type DictDataController struct{}

// List 字典数据列表
func (cp *DictDataController) List(c *gin.Context) {
	response.SuccessPage(c, 0, nil)
}

// Add 新增字典数据
func (cp *DictDataController) Add(c *gin.Context) {
	response.Success(c)
}

// Query 字典数据详情
func (cp *DictDataController) Query(c *gin.Context) {
	response.Success(c)
}

// Edit 更新字典数据
func (cp *DictDataController) Edit(c *gin.Context) {
	response.Success(c)
}

// Remove 删除字典数据
func (cp *DictDataController) Remove(c *gin.Context) {
	response.Success(c)
}

// Export 导出字典
func (cp *DictDataController) Export(c *gin.Context) {
	response.Success(c)
}

// DictType 根据字典类型查询字典数据信息
func (cp *DictDataController) DictType(c *gin.Context) {
	response.SuccessData(c, []gin.H{{
		"createBy":    "admin",
		"createTime":  "2020-11-20 19:29:50",
		"cssClass":    "",
		"default":     true,
		"dictCode":    12,
		"dictLabel":   "是",
		"dictSort":    1,
		"dictType":    "sys_yes_no",
		"dictValue":   "Y",
		"isDefault":   "Y",
		"listClass":   "primary",
		"params":      gin.H{},
		"remark":      "系统默认是",
		"searchValue": nil,
		"status":      "0",
		"updateBy":    nil,
		"updateTime":  nil,
	}})
}

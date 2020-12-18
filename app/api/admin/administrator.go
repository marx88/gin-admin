package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AdministratorController 管理员控制器
type AdministratorController struct{}

// Index 管理员列表
func (ac *AdministratorController) Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": nil,
	})
}

// Create 创建管理员
func (ac *AdministratorController) Create(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": nil,
	})
}

// Save 保存管理员
func (ac *AdministratorController) Save(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": nil,
	})
}

// Edit 编辑管理员
func (ac *AdministratorController) Edit(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": nil,
	})
}

// Update 更新管理员
func (ac *AdministratorController) Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": nil,
	})
}

// Delete 删除管理员
func (ac *AdministratorController) Delete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": nil,
	})
}

package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserController 用户控制器
type UserController struct {
	num int
}

// Index 用户列表
func (uc *UserController) Index(c *gin.Context) {
	uc.num++
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": uc.num,
	})
}

// Create 创建用户
func (uc *UserController) Create(c *gin.Context) {
	uc.num--
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": nil,
	})
}

// Save 保存用户
func (uc *UserController) Save(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": nil,
	})
}

// Edit 编辑用户
func (uc *UserController) Edit(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": nil,
	})
}

// Update 更新用户
func (uc *UserController) Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": nil,
	})
}

// Delete 删除用户
func (uc *UserController) Delete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": nil,
	})
}

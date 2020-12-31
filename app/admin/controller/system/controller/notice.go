package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/marx88/gin-admin/gan/response"
)

// NoticeController 公告控制器
type NoticeController struct{}

// List 公告列表
func (cp *NoticeController) List(c *gin.Context) {
	response.SuccessPage(c, 0, nil)
}

// Add 新增公告
func (cp *NoticeController) Add(c *gin.Context) {
	response.Success(c)
}

// Query 公告详情
func (cp *NoticeController) Query(c *gin.Context) {
	response.Success(c)
}

// Edit 更新公告
func (cp *NoticeController) Edit(c *gin.Context) {
	response.Success(c)
}

// Remove 删除公告
func (cp *NoticeController) Remove(c *gin.Context) {
	response.Success(c)
}

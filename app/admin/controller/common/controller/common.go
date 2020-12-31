package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/marx88/gin-admin/gan/response"
)

// CommonController 通用控制器
type CommonController struct{}

// Download 下载
func (cp *CommonController) Download(c *gin.Context) {
	// TODO::获取文件名、校验文件名是否合法、判断文件是否存在
	c.Header("Content-Disposition", "attachment; filename=文件名.txt")
	c.Header("Content-Type", "application/octet-stream")
	c.File("./file/a.txt")
}

// Upload 上传
func (cp *CommonController) Upload(c *gin.Context) {
	resp := response.New()

	// 获取文件句柄
	fileHeader, err := c.FormFile("file")
	if err != nil {
		resp.ErrorMsg(c, "获取文件信息失败："+err.Error())
		return
	}

	// 保存文件
	// TODO::定义上传完整路径、web访问的url
	fileName := "上传完整路径"
	url := "web访问的url"
	if err := c.SaveUploadedFile(fileHeader, fileName); err != nil {
		resp.ErrorMsg(c, "获取文件信息失败："+err.Error())
		return
	}

	resp.Put("fileName", fileName)
	resp.Put("url", url)
	resp.Success(c)
}

// Resource 本地资源下载
func (cp *CommonController) Resource(c *gin.Context) {
	// TODO::本地资源路径
	localPath := "本地资源路径"

	// 数据库资源地址
	// TODO::去掉前缀的文件名
	downloadPath := localPath + "去掉前缀的文件名"

	// TODO下载名称
	downloadName := "下载名称"

	c.Header("Content-Disposition", "attachment; filename="+downloadName)
	c.Header("Content-Type", "application/octet-stream")
	c.File(downloadPath)
}

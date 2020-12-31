package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/marx88/gin-admin/common/enum/businesstype"
)

// BusinessLog 业务日志中间件
func BusinessLog(title string, businessType businesstype.Enum) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		log.Println("标题：", title, "，业务类型：", businessType.String())
	}
}

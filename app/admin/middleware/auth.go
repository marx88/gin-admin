package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

// Auth 权限校验
func Auth(perms string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO::根据perms和c.GetInt("admin_id")查询是否拥有权限
		// if 没有权限 {
		//     c.AbortWithStatus(http.StatusBadRequest)
		// }
		log.Println("权限校验通过：", perms)
		c.Next()
	}
}

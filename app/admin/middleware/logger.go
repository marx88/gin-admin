package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

// Logger 日志中间件
func Logger(module string) gin.HandlerFunc {
	return func(c *gin.Context) {
		go func() {
			log.Printf("进入了%s\n", module)
		}()
		c.Next()
		go func() {
			log.Printf("离开了%s\n", module)
		}()
	}
}

package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/marx88/gin-admin/config"
)

// initGin 配置gin
func initGin() *gin.Engine {
	gin.SetMode(config.App.Mode)
	engine := gin.New()
	return engine
}

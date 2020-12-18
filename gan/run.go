package gan

import (
	"github.com/gin-gonic/gin"
	"github.com/marx88/gin-admin/config"
)

// run 运行并监听
func run(router *gin.Engine) {
	if config.App.Port == "" {
		panic("配置项【app['port']】不能为空")
	}

	if err := router.Run(config.App.Host + ":" + config.App.Port); err != nil {
		panic(err)
	}
}

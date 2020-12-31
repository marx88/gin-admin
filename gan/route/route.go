package route

import "github.com/gin-gonic/gin"

// IMount 挂载接口
type IMount interface {
	mount(gin.IRouter)
}

// Mount 挂载路由
func Mount(router gin.IRouter, mouter IMount) {
	mouter.mount(router)
}

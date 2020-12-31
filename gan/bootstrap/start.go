package bootstrap

import (
	"github.com/marx88/gin-admin/common/router"
	_ "github.com/marx88/gin-admin/gan/conf" // 加载配置文件
	"github.com/marx88/gin-admin/gan/route"
)

// Start 开启服务
func Start() {
	// 配置gin框架
	engine := initGin()

	// 配置gorm框架
	initGorm()

	// 加载注册的路由
	route.Mount(engine, router.Root)

	// 启动应用
	run(engine)
}

package gan

import (
	"github.com/marx88/gin-admin/app/routers"
	_ "github.com/marx88/gin-admin/gan/lib/conf" // 加载配置文件
)

// Start 开启服务
func Start() {
	// 配置gin框架
	engine := initGin()

	// 配置gorm框架
	initGorm()

	// 加载注册的路由
	routers.Root.Mount(engine)

	// 启动应用
	run(engine)
}

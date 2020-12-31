package router

import (
	"github.com/marx88/gin-admin/app/admin/router/common"
	"github.com/marx88/gin-admin/app/admin/router/monitor"
	"github.com/marx88/gin-admin/app/admin/router/system"
)

func init() {
	Root.SubGroups = append(
		Root.SubGroups,
		// 从这里添加子模块路由组 也可以在此目录添加个模块文件
		common.Root,
		system.Root,
		monitor.Root,
	)
}

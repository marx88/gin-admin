package route

import "github.com/gin-gonic/gin"

// Group 路由组
type Group struct {
	registers []func(gin.IRouter)
	SubGroups []*Group
	router    gin.IRouter
	Init      func(gin.IRouter) gin.IRouter // 用于生成路由组根路由
}

// NewGroup 创建Group实例
func NewGroup() *Group {
	return new(Group)
}

// Put append注册器
func (g *Group) Put(registers ...func(gin.IRouter)) {
	g.registers = append(g.registers, registers...)
}

func (g *Group) mount(parentRouter gin.IRouter) {
	g.generateGroupRouter(parentRouter)
	g.runRegisters()
	g.mountSubGroups()
}

func (g *Group) generateGroupRouter(parentRouter gin.IRouter) {
	if g.Init != nil {
		g.router = g.Init(parentRouter)
	} else {
		g.router = parentRouter.Group("")
	}
}

func (g *Group) runRegisters() {
	for _, r := range g.registers {
		r(g.router)
	}
}

func (g *Group) mountSubGroups() {
	for _, sg := range g.SubGroups {
		Mount(g.router, sg)
	}
}

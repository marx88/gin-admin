package controller

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/marx88/gin-admin/gan/response"
)

// LoginController 登录控制器
type LoginController struct{}

// Login 登录方法
func (cp *LoginController) Login(c *gin.Context) {
	resp := response.New()
	resp.Put("token", "hello world!")
	resp.Success(c)
}

// Info 获取用户信息
func (cp *LoginController) Info(c *gin.Context) {
	resp := response.New()
	resp.Put("permissions", []string{"*:*:*"})
	resp.Put("roles", []string{"admin"})
	resp.Put("user", gin.H{
		"admin":      true,
		"avatar":     "",
		"createBy":   "admin",
		"createTime": "2020-11-20 19:29:42",
		"delFlag":    "0",
		"dept": gin.H{
			"ancestors":   nil,
			"children":    []int{},
			"createBy":    nil,
			"createTime":  nil,
			"delFlag":     nil,
			"deptId":      103,
			"deptName":    "研发部门",
			"email":       nil,
			"leader":      "若依",
			"orderNum":    "1",
			"params":      gin.H{},
			"parentId":    101,
			"parentName":  nil,
			"phone":       nil,
			"remark":      nil,
			"searchValue": nil,
			"status":      "0",
			"updateBy":    nil,
			"updateTime":  nil,
		},
		"deptId":      103,
		"email":       "ry@163.com",
		"loginDate":   "2020-11-20T19:29:42.000+0800",
		"loginIp":     "127.0.0.1",
		"nickName":    "若依",
		"params":      gin.H{},
		"phonenumber": "15888888888",
		"postIds":     nil,
		"remark":      "管理员",
		"roleIds":     nil,
		"roles": []gin.H{
			{
				"admin":             true,
				"createBy":          nil,
				"createTime":        nil,
				"dataScope":         "1",
				"delFlag":           nil,
				"deptCheckStrictly": false,
				"deptIds":           nil,
				"flag":              false,
				"menuCheckStrictly": false,
				"menuIds":           nil,
				"params":            gin.H{},
				"remark":            nil,
				"roleId":            1,
				"roleKey":           "admin",
				"roleName":          "超级管理员",
				"roleSort":          "1",
				"searchValue":       nil,
				"status":            "0",
				"updateBy":          nil,
				"updateTime":        nil,
			},
		},
		"salt":        nil,
		"searchValue": nil,
		"sex":         "1",
		"status":      "0",
		"updateBy":    nil,
		"updateTime":  nil,
		"userId":      1,
		"userName":    "admin",
	})
	resp.Success(c)
}

// Routers 获取路由信息
func (cp *LoginController) Routers(c *gin.Context) {
	j := `[{"name":"System","path":"/system","hidden":false,"redirect":"noRedirect","component":"Layout","alwaysShow":true,"meta":{"title":"系统管理","icon":"system","noCache":false},"children":[{"name":"User","path":"user","hidden":false,"component":"system/user/index","meta":{"title":"用户管理","icon":"user","noCache":false}},{"name":"Role","path":"role","hidden":false,"component":"system/role/index","meta":{"title":"角色管理","icon":"peoples","noCache":false}},{"name":"Menu","path":"menu","hidden":false,"component":"system/menu/index","meta":{"title":"菜单管理","icon":"tree-table","noCache":false}},{"name":"Dept","path":"dept","hidden":false,"component":"system/dept/index","meta":{"title":"部门管理","icon":"tree","noCache":false}},{"name":"Post","path":"post","hidden":false,"component":"system/post/index","meta":{"title":"岗位管理","icon":"post","noCache":false}},{"name":"Dict","path":"dict","hidden":false,"component":"system/dict/index","meta":{"title":"字典管理","icon":"dict","noCache":false}},{"name":"Config","path":"config","hidden":false,"component":"system/config/index","meta":{"title":"参数设置","icon":"edit","noCache":false}},{"name":"Notice","path":"notice","hidden":false,"component":"system/notice/index","meta":{"title":"通知公告","icon":"message","noCache":false}},{"name":"Log","path":"log","hidden":false,"redirect":"noRedirect","component":"ParentView","alwaysShow":true,"meta":{"title":"日志管理","icon":"log","noCache":false},"children":[{"name":"Operlog","path":"operlog","hidden":false,"component":"monitor/operlog/index","meta":{"title":"操作日志","icon":"form","noCache":false}},{"name":"Logininfor","path":"logininfor","hidden":false,"component":"monitor/logininfor/index","meta":{"title":"登录日志","icon":"logininfor","noCache":false}}]}]},{"name":"Monitor","path":"/monitor","hidden":false,"redirect":"noRedirect","component":"Layout","alwaysShow":true,"meta":{"title":"系统监控","icon":"monitor","noCache":false},"children":[{"name":"Online","path":"online","hidden":false,"component":"monitor/online/index","meta":{"title":"在线用户","icon":"online","noCache":false}},{"name":"Job","path":"job","hidden":false,"component":"monitor/job/index","meta":{"title":"定时任务","icon":"job","noCache":false}},{"name":"Druid","path":"druid","hidden":false,"component":"monitor/druid/index","meta":{"title":"数据监控","icon":"druid","noCache":false}},{"name":"Server","path":"server","hidden":false,"component":"monitor/server/index","meta":{"title":"服务监控","icon":"server","noCache":false}},{"name":"Cache","path":"cache","hidden":false,"component":"monitor/cache/index","meta":{"title":"缓存监控","icon":"redis","noCache":false}}]},{"name":"Tool","path":"/tool","hidden":false,"redirect":"noRedirect","component":"Layout","alwaysShow":true,"meta":{"title":"系统工具","icon":"tool","noCache":false},"children":[{"name":"Build","path":"build","hidden":false,"component":"tool/build/index","meta":{"title":"表单构建","icon":"build","noCache":false}},{"name":"Gen","path":"gen","hidden":false,"component":"tool/gen/index","meta":{"title":"代码生成","icon":"code","noCache":false}},{"name":"Swagger","path":"swagger","hidden":false,"component":"tool/swagger/index","meta":{"title":"系统接口","icon":"swagger","noCache":false}}]},{"name":"Http://ruoyi.vip","path":"http://ruoyi.vip","hidden":false,"component":"Layout","meta":{"title":"若依官网","icon":"guide","noCache":false}}]`
	a := []interface{}{}
	json.Unmarshal([]byte(j), &a)
	response.SuccessData(c, a)
}

package config

import (
	"github.com/marx88/gin-admin/gan/lib/conf"
)

type app struct {
	Name  string `yaml:"name"`
	Host  string `yaml:"host"`
	Port  string `yaml:"port"`
	Mode  string `yaml:"mode"`
	Debug bool   `yaml:"debug"`
}

// App app配置
var App = new(app)

func init() {
	conf.New("app", App)
}

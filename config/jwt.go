package config

import "github.com/marx88/gin-admin/gan/conf"

type jwt struct {
	Secret  string `yaml:"secret"`
	Timeout uint32 `yaml:"timeout"`
}

// Jwt jwt配置
var Jwt = new(jwt)

func init() {
	conf.New("jwt", Jwt)
}

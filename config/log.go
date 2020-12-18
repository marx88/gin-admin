package config

import "github.com/marx88/gin-admin/gan/lib/conf"

type log struct {
	Path       string `yaml:"path"`
	Level      string `yaml:"level"`
	Stdout     bool   `yaml:"stdout"`
	EnabledBus bool   `yaml:"enabled_bus"`
	EnabledReq bool   `yaml:"enabled_req"`
}

// Log logger配置
var Log = new(log)

func init() {
	conf.New("log", Log)
}

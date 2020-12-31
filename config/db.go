package config

import "github.com/marx88/gin-admin/gan/conf"

type db struct {
	Type     string `yaml:"type"`
	Host     string `yaml:"host"`
	Post     string `yaml:"port"`
	Database string `yaml:"database"`
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
}

// Db db配置
var Db = new(db)

func init() {
	conf.New("db", Db)
}

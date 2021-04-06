package model

import (
	"github.com/marx88/gin-admin/gan/orm"
)

// Config 配置表
type Config struct {
	ConfigID    int32  `gorm:"type:int;primarykey;comment:参数主键" json:"configId"`
	ConfigName  string `gorm:"type:varchar;size:100;not null;comment:参数名称" json:"configName"`
	ConfigKey   string `gorm:"type:varchar;size:100;not nullcomment:参数键名" json:"configKey"`
	ConfigValue string `gorm:"type:varchar;size:500;comment:参数键值" json:"configValue"`
	ConfigType  string `gorm:"type:char;size:1;not null;default:N;comment:系统内置(Y是 N否)" json:"configType"`
	Remark      string `gorm:"type:varchar;size:500;comment:备注" json:"remark"`
	Base
}

// TableName 表名
func (Config) TableName() string {
	return "sys_config"
}

// ConfigList 配置列表
func ConfigList() (list []Config, total int64, err error) {
	result := orm.Db.Find(&list)
	total = result.RowsAffected
	err = result.Error
	return
}

package model

import (
	"github.com/marx88/gin-admin/gan/orm"
)

// Base 根模型
type Base struct {
	CreateBy   orm.String `gorm:"size:64;comment:创建者" json:"createBy"`
	CreateTime orm.Time   `gorm:"comment:创建时间" json:"createTime"`
	UpdateBy   orm.String `gorm:"size:64;comment:更新者" json:"updateBy"`
	UpdateTime orm.Time   `gorm:"comment:更新时间" json:"updateTime"`
}

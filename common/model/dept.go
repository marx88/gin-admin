package model

import "github.com/marx88/gin-admin/gan/orm"

// Dept 部门表
type Dept struct {
	DeptID    uint   `gorm:"primarykey;comment:部门ID" json:"deptId"`
	ParentID  int64  `gorm:"default:0;comment:父部门ID" json:"parentId"`
	Ancestors string `gorm:"size:50;comment:祖级列表" json:"ancestors"`
	DeptName  string `gorm:"size:30;comment:部门名称" json:"deptName"`
	OrderNum  int16  `gorm:"comment:显示顺序" json:"orderNum"`
	Leader    string `gorm:"size:20;comment:显示顺序" json:"leader"`
	Phone     string `gorm:"size:11;comment:联系电话" json:"phone"`
	Email     string `gorm:"size:50;comment:邮箱" json:"email"`
	Status    int8   `gorm:"type:char;size:1;default:0;comment:部门状态（0正常 1停用）" json:"status"`
	DelFlag   int8   `gorm:"type:char;size:1;default:0;comment:删除标志（0代表存在 2代表删除）" json:"delFlag"`
	Base
}

// TableName 表名
func (Dept) TableName() string {
	return "sys_dept"
}

// DeptList 部门列表
func DeptList() (list []Dept, total int64, err error) {
	result := orm.Db.Find(&list)
	total = result.RowsAffected
	err = result.Error
	return
}

package model

import "github.com/marx88/gin-admin/gan/orm"

// DictData 配置表
type DictData struct {
	// DictDataID    uint   `gorm:"primarykey;comment:参数主键" json:"dict_code"`
	// DictDataName  string `gorm:"size:100;not null;comment:参数名称" json:"configName"`
	// DictDataKey   string `gorm:"size:100;not nullcomment:参数键名" json:"configKey"`
	// DictDataValue string `gorm:"size:500;comment:参数键值" json:"configValue"`
	// DictDataType  string `gorm:"type:char;size:1;not null;default:N;comment:系统内置(Y是 N否)" json:"configType"`

	DictCode  int64  `gorm:"comment:;" json:"dict_code"`
	DictSort  int64  `gorm:"comment:;" json:"dict_sort"`
	DictLabel int64  `gorm:"comment:;" json:"dict_label"`
	DictValue int64  `gorm:"comment:;" json:"dict_value"`
	DictType  int64  `gorm:"comment:;" json:"dict_type"`
	CSSClass  int64  `gorm:"comment:;" json:"css_class"`
	ListClass int64  `gorm:"comment:;" json:"list_class"`
	IsDefault int64  `gorm:"comment:;" json:"is_default"`
	Status    int64  `gorm:"comment:;" json:"status"`
	Remark    string `gorm:"size:500;comment:备注" json:"remark"`
	Base
}

// TableName 表名
func (DictData) TableName() string {
	return "sys_dict_data"
}

// DictDataList 字典数据列表
func DictDataList() (list []DictData, total int64, err error) {
	result := orm.Db.Find(&list)
	total = result.RowsAffected
	err = result.Error
	return
}

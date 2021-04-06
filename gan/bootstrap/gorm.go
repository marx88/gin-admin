package bootstrap

import (
	"fmt"
	"strings"

	"github.com/marx88/gin-admin/config"
	"github.com/marx88/gin-admin/gan/orm"
)

// initGorm 初始化Gorm库
func initGorm() {
	if strings.ToLower(config.Db.Type) != "mysql" {
		panic("不支持的数据库类型：" + config.Db.Type)
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		config.Db.UserName,
		config.Db.Password,
		config.Db.Host,
		config.Db.Port,
		config.Db.Database,
		config.Db.Charset,
	)
	if err := orm.InitMysql(dsn); err != nil {
		panic(err)
	}
}

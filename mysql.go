package mysql

import (
	"fmt"
	"github.com/0xJacky/cosy/settings"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Open(dbs *settings.DataBase) gorm.Dialector {
	return mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbs.User, dbs.Password, dbs.Host, dbs.Port, dbs.Name))
}

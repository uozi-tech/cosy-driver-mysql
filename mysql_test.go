package mysql

import (
	"github.com/0xJacky/cosy/settings"
	"gorm.io/driver/mysql"
	"testing"
)

func TestOpen(t *testing.T) {
	dbs := &settings.DataBase{
		User:     "cosy",
		Password: "cosy",
		Host:     "127.0.0.1",
		Port:     3306,
		Name:     "cosy",
	}

	dialector := Open(dbs)

	d, ok := dialector.(*mysql.Dialector)
	if !ok {
		t.Fatal("dialector is not *Dialector")
	}
	if d.DSN == "" {
		t.Error("dialector.DSN is empty")
	}
}

package dao

import (
	"djforgo/system"
	l4g "github.com/alecthomas/log4go"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB_Instance *gorm.DB

func DB_Init() error {
	var err error
	DB_Instance, err = gorm.Open(system.SysConfig.DB.Drivername, system.SysConfig.DB.DataSource)
	if err != nil {
		l4g.Error(err)
	}

	return err
}

func DB_Destroy() {
	if DB_Instance != nil {
		DB_Instance.Close()
	}
}

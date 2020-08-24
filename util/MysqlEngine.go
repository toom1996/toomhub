package util

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB
var err error

func MysqlInit() {
	config := GetConfig()
	DB, err = gorm.Open("mysql", config.Database.Username+":"+config.Database.Password+"@tcp("+config.Database.Host+":"+config.Database.Port+")/"+config.Database.Dbname+"?charset="+config.Database.Charset+"&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println(err)
	}
	DB.LogMode(true)
	DB.SingularTable(true)

}

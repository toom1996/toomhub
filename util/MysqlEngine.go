package util

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB
var err error

func MysqlInit() {
	DB, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/toom?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println(err)
	}
	DB.LogMode(true)
	DB.SingularTable(true)

}

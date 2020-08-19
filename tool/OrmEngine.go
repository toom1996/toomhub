package tool

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB
var err error

func init() {
	DB, err = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/toomhub?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println(err)
	}
	DB.LogMode(true)
	DB.SingularTable(true)

}

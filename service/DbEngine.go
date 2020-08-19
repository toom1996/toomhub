// @Description
// @Author    2020/8/18 8:55
package service

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB
var err error

func init() {
	DB, err = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/toomhub?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println(err)
	}
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
	DB.DB().SetConnMaxLifetime(1)
	DB.LogMode(true)
	DB.SingularTable(true)
}

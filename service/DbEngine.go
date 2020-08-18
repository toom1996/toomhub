// @Description
// @Author    2020/8/18 8:55
package service

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB
var err error

func init() {
	DB, err = gorm.Open("mysql", "root:root@tcp(192.168.10.113:3306)/toom?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println(err)
	}
	DB.LogMode(true)
	DB.SingularTable(true)
}

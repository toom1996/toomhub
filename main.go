package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"toom/tool"
)

func init() {

}

func main() {
	//db, err := gorm.Open("mysql", "root:root@tcp(192.168.10.113:3306)/toom?charset=utf8mb4&parseTime=True&loc=Local")
	//if err != nil {
	//	fmt.Println("666666")
	//	fmt.Println(err)
	//}
	//db.LogMode(true)
	//db.SingularTable(true)
	//fmt.Println(db)
	//tt, _ := db.Raw("SELECT * FROM toomhub_user_mini WHERE 1").Rows()
	//fmt.Println("1111")
	//fmt.Println(tt)
	//for tt.Next() {
	//	var mini_id int
	//	var openid string
	//	var created_at string
	//	err = tt.Scan(&mini_id, &openid, &created_at)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	fmt.Println(mini_id, openid, created_at)
	//}
	//
	//defer db.Close()

	////初始化配置
	config, _ := tool.Load("./config/app.json")

	app := gin.Default()
	//注册路由
	tool.RegisterRoutes(app)
	_ = app.Run(config.AppHost + ":" + config.AppPort)
}

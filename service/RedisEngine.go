// @Description
// @Author    2020/8/18 9:06
package service

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()

// 声明一个全局的rdb变量
var Rdb *redis.Client

func init() {
	//config := util.GetConfig()
	//// 实例化一个redis客户端
	//Rdb = redis.NewClient(&redis.Options{
	//	Addr:     config.Redis.Host + ":" + config.Redis.Port, // ip:port
	//	Password: config.Redis.Password,                       // redis连接密码
	//	DB:       config.Redis.Database,                       // 选择的redis库
	//	PoolSize: 20,                                          // 设置连接数,默认是10个连接
	//})
}

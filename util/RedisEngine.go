// @Description
// @Author    2020/8/18 9:06
package util

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()

// 声明一个全局的rdb变量
var Rdb *redis.Client

func RedisInit() {
	config := GetConfig()
	// 实例化一个redis客户端
	Rdb = redis.NewClient(&redis.Options{
		Addr:     config.Redis.Host + ":" + config.Redis.Port, // ip:port
		Password: config.Redis.Password,                       // redis连接密码
		DB:       0,                                           // 选择的redis库
		PoolSize: 20,                                          // 设置连接数,默认是10个连接
	})
}

// pipeline 批量HMGET
func RedisMulti(fields []string, keys ...string) ([]interface{}, error) {
	pipe := Rdb.Pipeline()
	var commands []*redis.SliceCmd
	var slice []interface{}
	for _, item := range keys {
		commands = append(commands, pipe.HMGet(Ctx, item, fields...))
	}
	_, err = pipe.Exec(Ctx)
	if err != nil {
		return []interface{}{}, errors.New(err.Error())
	}
	for _, cmd := range commands {
		result, _ := cmd.Result()
		slice = append(slice, result)
	}
	return slice, nil
}

//
func RedisGetUserInfo(id string, fields ...string) ([]interface{}, error) {
	r, err := Rdb.HMGet(Ctx, UserCacheKey+id, fields...).Result()
	if err != nil {
		return nil, err
	}
	return r, nil
}

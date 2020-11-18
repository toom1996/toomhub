// @Description
// @Author: toom1996 <1023150697@qq.com>
// @dateTime: 2020/11/18 14:36
package util

import (
	"fmt"
	"github.com/go-redis/redis/v8"
)

func GetUserInfoForRedis(column []string, keys ...string) {

	pipe := Rdb.Pipeline()
	var commands []*redis.SliceCmd
	var array []interface{}
	for _, item := range keys {
		fmt.Println(item)
		commands = append(commands, pipe.HMGet(Ctx, item, column...))
	}

	_, _ = pipe.Exec(Ctx)

	for _, cmd := range commands {
		result, _ := cmd.Result()
		array = append(array, result)
	}

	fmt.Println(array)
}

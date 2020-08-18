// @Description
// @Author    2020/8/18 9:06
package service

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

const REDIS_KEY_USER = "miniuser:id:"

func init() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()
}

func gerUser() {

}

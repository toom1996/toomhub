// @Description
// @Author    2020/8/20 11:49
package util

import (
	"math/rand"
	"strconv"
	"time"
)

func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func ToInt(str string) int64 {
	res, _ := strconv.ParseInt(str, 10, 0)
	return res
}

// @Description
// @Author: toom1996 <1023150697@qq.com>
// @dateTime: 2021/2/23 13:58
package util

import (
	"math/rand"
	"time"
)

// 生成指定范围的随机数
func GenerateRandomInt(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(max - min)
	randNum = randNum + min
	return randNum
}

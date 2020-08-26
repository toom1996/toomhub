// @Description
// @Author    2020/8/26 10:51
package ServiceMiniV1

import (
	"fmt"
	validatorMiniprogramV1 "toomhub/validator/miniprogram/v1"
)

// @title
func GetSquareIndex(validator *validatorMiniprogramV1.SquareIndex) (interface{}, error) {
	fmt.Println(validator)
	return map[string]interface{}{
		"created_at":     "2020: 01 :08",
		"content":        "测试的那个",
		"created_by":     "admin",
		"likes_count":    "555",
		"argument_count": "111",
		"collect_count":  "88",
		"share_count":    "100",
	}, nil
}

func SetSquarePost() {

}

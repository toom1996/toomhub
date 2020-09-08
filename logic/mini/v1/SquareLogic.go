// @Description
// @Author    2020/8/26 10:23
package LogicMiniV1

import (
	"encoding/json"
	"fmt"
	"github.com/goinggo/mapstructure"
	ServiceMiniV1 "toomhub/service/mini/v1"
	validatorMiniprogramV1 "toomhub/validator/miniprogram/v1"
)

type SquareLogic struct {
}

// @title	获取图片广场信息
func (logic *SquareLogic) SquareIndex(validator *validatorMiniprogramV1.SquareIndex) (interface{}, error) {

	query, err := ServiceMiniV1.GetSquareIndex(validator)

	if err != nil {
		fmt.Println("000000")
	}

	return query, nil
}

type ImageList struct {
	RequestHost
	Size
	Name
	Param
}

func (logic *SquareLogic) SquareCreate(validator *validatorMiniprogramV1.SquareCreate) (bool, error) {

	fmt.Println(validator)

	var dat map[string]interface{}
	err := json.Unmarshal([]byte(validator.ImageList), &dat)
	if err != nil {
		fmt.Println(err)
	}

	for k, value := range dat {
		fmt.Println(k)
		fmt.Println(value)
	}

	mapstructure.Decode()
	return true, nil
}

// @Description
// @Author    2020/8/26 10:01
package validatorMiniprogramV1

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
)

type SquareIndex struct {
	LastId int `form:"last_id" binding:"required"`
	Page   int `form:"page" binding:"required"`
}

// @description	label是输出错误信息时显示的字段中文名
type SquareCreate struct {
	Content   string `json:"content" form:"content" binding:"required" label:"内容"`
	ImageList string `form:"image_list" binding:"required,countValidate"  label:"图片列表"`
	Tag       string `form:"tag"`
}

type CreateImageInfo struct {
	Param       string `json:"param"`
	RequestHost int    `json:"request_host"`
	Name        int    `json:"name"`
	Size        int    `json:"size"`
	Extension   int    `json:"extension"`
}

type LikeValidator struct {
	Id int64 `json:"id" form:"id" binding:"required"`
}

//图片数量验证
func CountValidate(fl validator.FieldLevel) bool {

	df := fl.Field().String()
	//json转化成map
	var dat map[string]interface{}
	err := json.Unmarshal([]byte(df), &dat)
	if err != nil {
		fmt.Println(err)
	}

	l := len(dat)

	if l == 0 || l > 9 {
		return false
	}

	return true
}

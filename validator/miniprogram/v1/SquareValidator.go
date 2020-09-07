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

type createImageInfo struct {
	Url  string `json:"url"`
	Size int    `json:"size"`
}

func CountValidate(fl validator.FieldLevel) bool {

	var create createImageInfo
	aa := `{"deletable":true,"status":"done","message":"0%","url":"http://toomhub.image.23cm.cn/006APoFYly1fowt3eeuk6g306o08g4q3.gif?imageMogr2/auto-orient/format/webp","extension":".jpg","size":51405}`
	err := json.Unmarshal([]byte(aa), &create)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(create)
	//fmt.Println(fl.Field().String())
	return false
}

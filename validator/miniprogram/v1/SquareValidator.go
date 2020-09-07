// @Description
// @Author    2020/8/26 10:01
package validatorMiniprogramV1

type SquareIndex struct {
	LastId int `form:"last_id" binding:"required"`
	Page   int `form:"page" binding:"required"`
}

// @description	label是输出错误信息时显示的字段中文名
type SquareCreate struct {
	Content   string            `json:"content" form:"content" binding:"required" label:"内容"`
	ImageList map[string]string `json:"image_list" form:"image_list" binding:"len=2,min=2"  label:"图片列表"`
	Tag       string            `form:"tag"`
}

type ImageInfo struct {
	Url string `json:"url" binding:"required"`
}

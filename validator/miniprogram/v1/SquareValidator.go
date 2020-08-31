// @Description
// @Author    2020/8/26 10:01
package validatorMiniprogramV1

type SquareIndex struct {
	LastId int `form:"last_id" binding:"required"`
	Page   int `form:"page" binding:"required"`
}

type SquareCreate struct {
	Content   string `from:"content" binding:"required"`
	ImageList string `from:"image_list" binding:"required"`
	Tag       string `from:"tag" binding:"required"`
}

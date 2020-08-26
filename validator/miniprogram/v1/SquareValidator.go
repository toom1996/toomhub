// @Description
// @Author    2020/8/26 10:01
package validatorMiniprogramV1

type SquareIndex struct {
	LastId int `form:"last_id" binding:"required"`
	Page   int `form:"page" binding:"required"`
}

// @Description
// @Author    2020/8/17 9:45
package model

//数据表toomhub_user_mini 结构体
type ToomhubUserMini struct {
	MiniId    int
	Openid    string
	CreatedAt int
}

type UserMiniCreate struct {
	Code string `json:"code" form:"code" binding:"required"`
}

func UserMiniCreated(Openid int) {

}

// @Description
// @Author    2020/8/17 9:45
package model

//数据表toomhub_user_mini 结构体
type ToomhubUserMini struct {
	MiniId    int
	Openid    string
	CreatedAt int
}

type V1MiniUserInfo struct {
	Nickname  string `json:"nickname" form:"nickname" binding:"required"`
	Gender    int    `json:"gender" form:"gender" binding:"required"`
	City      string `json:"city" form:"city" binding:"required"`
	Province  string `json:"province" form:"province" binding:"required"`
	Country   string `json:"country" form:"country" binding:"required"`
	AvatarUrl string `json:"avatarUrl" form:"avatarUrl" binding:"required"`
}

type LoginByV1Model struct {
	Code     string           `form:"code" binding:"required"`
	UserInfo []V1MiniUserInfo `form:"userInfo"`
}

func UserMiniCreated(Openid int) {

}

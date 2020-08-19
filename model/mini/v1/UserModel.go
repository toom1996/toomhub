// @Description
// @Author    2020/8/19 16:16
package ModelMiniV1

//表 toomhub_user_mini
type ToomhubUserMini struct {
	MiniId    int `json:"mini_id"`
	OpenId    string
	CreatedAt int64 `gorm:"default:''"`
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

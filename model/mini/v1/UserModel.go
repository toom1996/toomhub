// @Description
// @Author    2020/8/19 16:16
package ModelMiniV1

//数据表 toomhub_user_mini 结构体
type ToomhubUserMini struct {
	MiniId    int    `json:"mini_id"`
	OpenId    string `json:"open_id"`
	CreatedAt int64  `json:"created_at"`
}

//数据表 toomhub_user_mini 结构体
type ToomhubUserMiniProfile struct {
	Id        int
	MiniId    int
	NickName  string
	Gender    int8
	City      string
	Province  string
	Country   string
	AvatarUrl string
}

//数据表 toomhub_user_mini_token 结构体
type ToomhubUserMiniToken struct {
	MiniId       int
	AccessToken  string
	RefreshToken string
	CreatedAt    int64
	UpdatedAt    int64
}

type V1MiniUserInfo struct {
	Nickname  string `json:"nickname" form:"nickname"`
	Gender    int8   `json:"gender" form:"gender"`
	City      string `json:"city" form:"city"`
	Province  string `json:"province" form:"province"`
	Country   string `json:"country" form:"country"`
	AvatarUrl string `json:"avatarUrl" form:"avatarUrl" binding:"required"`
}

type LoginByV1Model struct {
	Code     string         `form:"code" binding:"required"`
	UserInfo V1MiniUserInfo `form:"userInfo"`
}

func UserMiniCreated(Openid int) {

}

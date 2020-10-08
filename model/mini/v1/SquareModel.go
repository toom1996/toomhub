// @Description
// @Author    2020/9/1 15:37
package ModelMiniV1

// TABLE `toomhub_square` 结构体
type ToomhubSquare struct {
	Id            int64  `json:"id"`
	Content       string `json:"content"`          //内容
	CreatedBy     int64  `json:"created_by"`       //创建时间
	CreatedAt     int64  `json:"created_at"`       //创建人
	LikesCount    int64  `json:"likes_count"`      //点赞数
	ArgumentCount int64  `json:"argument_count"`   //评论数
	CollectCount  int64  `json:"collection_count"` //收藏数
	ShareCount    int64  `json:"share_count"`      //分享数
	Tag           string `json:"tag"`
}

// TABLE `toomhub_square_image` 结构体
type ToomhubSquareImage struct {
	Rid   int64
	Host  string `json:"host"`
	Size  int64  `json:"size"`
	Ext   string
	Param string
	Name  string
}

type ToomhubSquareTag struct {
	Id    int64
	Tag   string
	Count int64
}

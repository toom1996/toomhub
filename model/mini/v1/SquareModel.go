// @Description
// @Author    2020/9/1 15:37
package ModelMiniV1

// TABLE `toomhub_square` 结构体
type Square struct {
	Content       string `json:"content"`          //内容
	CreatedBy     int64  `json:"created_by"`       //创建时间
	CreatedAt     int64  `json:"created_at"`       //创建人
	LikesCount    int64  `json:"likes_count"`      //点赞数
	ArgumentCount int64  `json:"argument_count"`   //评论数
	CollectCount  int64  `json:"collection_count"` //收藏数
	ShareCount    int64  `json:"share_count"`      //分享数
}

// TABLE `toomhub_square_image` 结构体
type SquareImage struct {
	ImageUrl       string //图片地址
	ImageExtension string //图片后缀
	ImageSize      int64  //图片大小
	SquareId       int    //关联的square
}

// @Description 发布消息逻辑层
// @Author: toom1996 <1023150697@qq.com>
// @dateTime: 2021/3/10 15:49
package logic

import (
	rules "toomhub/rules/user/v1"
	"toomhub/service"
)

type PostLogic struct {
}

func (l *PostLogic) PublishPost(v *rules.V1PostPublishPost) {
	s := service.PostService{}
	// TODO 判断用户状态

	//TODO 存储发布内容
	_, _ = s.Create(v)
}

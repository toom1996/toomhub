package service

import (
	"fmt"
	"toomhub/model"
	rules "toomhub/rules/user/v1"
	"toomhub/util"
)

type PostService struct {
}

func (p *PostService) Create(v *rules.V1PostPublishPost) (bool, error) {
	info := model.ZawazawaPost{
		Content: "ooooo",
	}
	e := model.ZawazawaPostMgr(util.DB).Create(&info).Error
	fmt.Println(e)

	return true, nil
}

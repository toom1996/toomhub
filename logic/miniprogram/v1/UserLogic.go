// @Description
// @Author    2020/8/14 17:03
package v1

import (
	"fmt"
	"toom/model"
	"toom/service"
)

type UserLogic struct {
}

func (logic *UserLogic) Login(validator *model.UserMiniCreate) bool {
	if service.V1MiniUserHasUser(validator.Code) == false {
		fmt.Println("false")
	}
	fmt.Println("true")
	return false
}

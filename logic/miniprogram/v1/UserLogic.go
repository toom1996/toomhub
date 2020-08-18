// @Description
// @Author    2020/8/14 17:03
package v1

import (
	"errors"
	"fmt"
	"github.com/medivhzhan/weapp/v2"
	"toom/model"
)

type UserLogic struct {
}

func (logic *UserLogic) Login(validator *model.LoginByV1Model) (bool, error) {
	res, err := weapp.Login("wxa9a7f53bff2fc937", "971869511ee44662c56bcf8a833bd679", validator.Code)
	if err != nil {
		fmt.Println(err)
	}
	if res.ErrCode != 0 {
		return false, errors.New(res.ErrMSG)
	}
	fmt.Println(res.ErrCode)
	//if service.V1MiniUserHasUser(res.OpenID) == false {
	fmt.Println("00")
	//}
	//fmt.Println("true")
	return false, err
}

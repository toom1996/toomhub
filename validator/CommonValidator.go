// @Description
// @Author    2020/8/31 17:35
package validator

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type CommonValidator struct {
}

func (c *CommonValidator) TransError(err validator.ValidationErrors) string {

	fmt.Println(err)
	return "参数错误"
}

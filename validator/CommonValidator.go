// @Description
// @Author    2020/8/31 17:35
package validator

import "github.com/go-playground/validator/v10"

type CommonValidator struct {
}

func (c *CommonValidator) TransError(err validator.ValidationErrors) string {

	return "参数错误"
}

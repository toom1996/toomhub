// @Description
// @Author    2020/8/31 17:35
package validatorRules

import (
	"github.com/go-playground/validator/v10"
)

type CommonValidator struct {
}

func (c *CommonValidator) TransError(err validator.ValidationErrors) string {
	var firstError = err[0]
	return firstError.Field() + GetBindingLabel(firstError.Tag())
}

// @title	翻译验证规则标签
func GetBindingLabel(tag string) string {
	var label = map[string]string{
		"required": "不能为空",
	}

	if _, ok := label[tag]; ok {
		return label[tag]
	}

	return "有误"
}

package rules

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	zhongwen "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"regexp"
)

//var v *validator.Validate
//
var trans ut.Translator

func InitVali() {
	// 中文翻译
	zh := zhongwen.New()
	uni := ut.New(zh, zh)
	trans, _ = uni.GetTranslator("zh")

	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		// 自定义验证方法
		_ = v.RegisterValidation("mobileValidator", mobileValidator)

		_ = v.RegisterValidation("checkMobileForV1UserRegister", checkMobileForV1UserRegister)
		// 验证器注册翻译器
		_ = zh_translations.RegisterDefaultTranslations(v, trans)

		//注册一个函数，获取struct tag里自定义的label作为字段名
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := fld.Tag.Get("label")
			return name
		})

		// 添加额外翻译
		_ = v.RegisterTranslation("mobileValidator", trans, func(ut ut.Translator) error {
			_ = ut.Add("checkMobileForV1UserRegister", "{0}手机号已被注册", true)
			return ut.Add("mobileValidator", "{0}格式错误", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("mobileValidator", fe.Field())
			return t
		})
	}
}

//只显示第一条错误
func Translate(errs validator.ValidationErrors) string {
	fmt.Println(errs[0])
	//var errList []string
	//for _, e := range errs {
	//	// can translate each error one at a time.
	//	errList = append(errList, e.Translate(trans))
	//}
	return errs[0].Translate(trans)
	//return strings.Join(errList, "|")
}

func mobileValidator(fl validator.FieldLevel) bool {
	mobile := fl.Field().String()
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobile)
}

//验证手机号码是否被注册过
func checkMobileForV1UserRegister(fl validator.FieldLevel) bool {
	return true
}

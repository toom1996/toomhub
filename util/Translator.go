package util

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	zhongwen "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"regexp"
	"toomhub/model"
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

		_ = v.RegisterValidation("checkPublishImage", checkPublishImage)
		// 验证器注册翻译器
		_ = zh_translations.RegisterDefaultTranslations(v, trans)

		//注册一个函数，获取struct tag里自定义的label作为字段名
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := fld.Tag.Get("label")
			return name
		})

		// 添加额外翻译
		_ = v.RegisterTranslation("mobileValidator", trans, func(ut ut.Translator) error {
			return ut.Add("mobileValidator", "{0}格式错误", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("mobileValidator", fe.Field())
			return t
		})

		// 添加额外翻译
		_ = v.RegisterTranslation("checkMobileForV1UserRegister", trans, func(ut ut.Translator) error {
			return ut.Add("checkMobileForV1UserRegister", "{0}手机号已被注册", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("checkMobileForV1UserRegister", fe.Field())
			return t
		})

		// 添加额外翻译
		_ = v.RegisterTranslation("checkPublishImage", trans, func(ut ut.Translator) error {
			return ut.Add("checkPublishImage", "{0}验证失败", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("checkPublishImage", fe.Field())
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

//验证手机号码格式是否正确
func mobileValidator(fl validator.FieldLevel) bool {
	mobile := fl.Field().String()
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobile)
}

//验证手机号码是否被注册过
func checkMobileForV1UserRegister(fl validator.FieldLevel) bool {
	mobile := fl.Field().String()
	s := model.ZawazawaUser{}
	query := model.ZawazawaUserMgr(DB).Select([]string{"mobile"}).Where(&model.ZawazawaUser{
		Mobile: mobile,
	}).Take(&s)
	fmt.Println("checkMobileForV1UserRegister")
	if query.Error != nil {
		return true
	}
	return false
}

type Image struct {
	Key string `json:"key"`
}

//验证上传图片
func checkPublishImage(fl validator.FieldLevel) bool {
	formData := fl.Field().String()
	image := Image{}

	err := json.Unmarshal([]byte(formData), &image)
	if err != nil {
		fmt.Println(err)
		return false
	}

	fmt.Println(image)
	return true
}

package rules

//用户注册验证规则
type V1UserRegister struct {
	Mobile string `form:"mobile" binding:"required,mobileValidator" label:"手机号码"`
	Code   string `form:"code" binding:"required" label:"验证码"`
}

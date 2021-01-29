package rules

//用户注册验证规则
type V1UserRegister struct {
	Mobile string `form:"mobile" binding:"required,mobileValidator,checkMobileForV1UserRegister" label:"手机号码"`
	Code   string `form:"code" binding:"required" label:"验证码"`
}

type V1UserSmsSend struct {
	Mobile string `form:"mobile" binding:"required,mobileValidator,checkMobileForV1UserRegister" label:"手机号码"`
}

//发布规则验证
type V1PostPublishPost struct {
	Content string `form:"content" binding:"required" label:"zawa内容"`
	Image   string `form:"image" binding:"required,checkPublishImage" label:"图片"`
}

//githubOauth 回调
type V1UserGithubOAuth struct {
	Code string `form:"code" binding:"required" label:"OAuthCode"`
}

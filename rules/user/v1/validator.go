package rules

//用户注册验证规则
type V1UserRegister struct {
	Mobile string `form:"mobile" binding:"required,mobileValidator" label:"手机号码"`
	Code   string `form:"code" binding:"required" label:"验证码"`
}

type V1UserSmsSend struct {
	Mobile string `form:"mobile" binding:"required,mobileValidator" label:"手机号码"`
}

//发布规则验证
type V1PostPublishPost struct {
	Content string `form:"content" binding:"required" label:"发布内容"`
	Attach  string `form:"attach" binding:"required,checkPublishImage" label:"附件"`
}

//githubOauth 回调
type V1UserGithubOAuth struct {
	Code  string `form:"code" binding:"required" label:"OAuthCode"`
	State string `form:"state" binding:"required" label:"State"`
}

//获取授权用户信息
type V1UserGetOAuthInfo struct {
	Code string `form:"code" binding:"required" label:"OAuthCode"`
}

// 刷新token验证规则
type V1UserRefreshToken struct {
	RefreshToken string `form:"refresh_token" binding:"required" label:"refreshToken"`
}

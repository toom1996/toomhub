//// @Description
//// @Author    2020/8/19 15:59
package controllers

import (
	"github.com/gin-gonic/gin"
	"toomhub/util"
	"toomhub/validatorRules"
)

//
type UserController struct {
}

//当前控制器注册的路由
func (u *UserController) Register(engine *gin.Engine) {
	//user := engine.Group("/v1/mini/user")
	//{
	//小程序用户登陆接口
	//user.POST("/login", u.Login)
	//user.POST("/token-checker", u.tokenChecker)
	//user.GET("/get-session", u.GetSessionKey)

	//}

	group := engine.Group("/api/v1/user")
	{
		group.POST("/register", u.actionRegister)
	}
	//user.Use(middleware.CheckIdentity())
	//{
	//	user.GET("/get-info", u.refreshInfo)
	//}
}

//
//func (u *UserController) GetSessionKey(Context *gin.Context) {
//	config := util.GetConfig()
//	code, _ := Context.GetQuery("code")
//
//	if code == "" {
//		Context.JSON(200, gin.H{
//			"code":    400,
//			"message": "can't find code",
//		})
//		return
//	}
//	wechatLogin, _ := weapp.Login(config.Mini.AppId, config.Mini.AppSecret, code)
//
//	if wechatLogin.ErrCode != 0 {
//		Context.JSON(200, gin.H{
//			"code":    400,
//			"message": wechatLogin.ErrMSG,
//		})
//		return
//	}
//
//	wechatLoginJson, err := json.Marshal(wechatLogin)
//
//	if err != nil {
//		Context.JSON(200, gin.H{
//			"code":    400,
//			"message": err.Error(),
//		})
//		return
//	}
//
//	authKey := util.GetRandomString(10) + fmt.Sprintf("%d", time.Now().Unix())
//	_, err = util.Rdb.Set(util.Ctx, authKey, wechatLoginJson, time.Second*7150).Result()
//
//	if err != nil {
//		Context.JSON(200, gin.H{
//			"code":    400,
//			"message": err.Error(),
//		})
//		return
//	}
//
//	fmt.Println(code)
//	Context.JSON(200, map[string]interface{}{
//		"code":    200,
//		"message": "OK",
//		"errcode": 0,
//		"data": map[string]interface{}{
//			"authKey": authKey,
//		},
//	})
//	return
//}
//
//// @url 	localhost:8080/mini/login	POST
//// @title    小程序用户登陆接口
//// @description   初次登陆的用户将会入库并返回信息, 非初次登陆的用户将会返回用户信息
//// @auth	toom <1023150697@qq.com>
//// @param     Context	*gin.Context
//// @return
//func (u *UserController) Login(Context *gin.Context) {
//	//validator验证
//	formValidator := validatorRules.Login{}
//	err := Context.ShouldBind(&formValidator)
//	if err != nil {
//		Context.String(http.StatusOK, "参数错误:%s", err.Error())
//		return
//	}
//
//	logic := LogicMiniV1.UserLogic{}
//	query, err := logic.Login(&formValidator)
//
//	if err != nil {
//		Context.JSON(http.StatusOK, gin.H{
//			"code":    400,
//			"message": err.Error(),
//			"data":    "",
//		})
//		return
//	}
//
//	Context.JSON(http.StatusOK, gin.H{
//		"code":    200,
//		"message": "登陆成功",
//		"data":    query,
//	})
//}
//
//// @title	检查token接口
//func (u *UserController) tokenChecker(Context *gin.Context) {
//	var commonValidator validatorRules.CommonValidator
//
//	formValidator := validatorRules.Refresh{}
//	err := Context.ShouldBind(&formValidator)
//	if err != nil {
//		Context.JSON(http.StatusOK, gin.H{
//			"code": 400,
//			"msg":  commonValidator.TransError(err.(validator.ValidationErrors)),
//		})
//		return
//	}
//
//	formLogic := LogicMiniV1.UserLogic{}
//	token, err := formLogic.Check(&formValidator, Context)
//	if err != nil {
//		Context.JSON(http.StatusOK, gin.H{
//			"code": 401,
//			"msg":  err.Error(),
//		})
//		return
//	}
//	Context.JSON(http.StatusOK, gin.H{
//		"code": 200,
//		"msg":  token,
//	})
//}
//
//// @title 取用户最新的信息
//func (u *UserController) refreshInfo(Context *gin.Context) {
//
//	r, _ := util.Rdb.HMGet(util.Ctx, util.UserCacheKey+strconv.FormatInt(util.GetIdentity().MiniId, 10), []string{"likes_count", "fans_count", "follow_count", "exp"}...).Result()
//
//	for index, _ := range r {
//		fmt.Println(r[index])
//		if r[index] != nil {
//			r[index], _ = strconv.Atoi(r[index].(string))
//		} else {
//			r[index] = 0
//		}
//	}
//
//	tagInfo := util.GetLevelTag(r[3].(int))
//	Context.JSON(200, gin.H{
//		"code":    200,
//		"message": "OK",
//		"data": map[string]interface{}{
//			"likes_count":  r[0],
//			"fans_count":   r[1],
//			"follow_count": r[2],
//			"exp":          r[3],
//			"tag": map[string]interface{}{
//				"text":             tagInfo[0],
//				"background_color": tagInfo[1],
//				"text-color":       tagInfo[2],
//			},
//		},
//	})
//}
//

type ReleaseTemplateAdd struct {
	Name               string `json:"name"`
	DeployEnv          string `json:"deploy_env"`
	GitlabType         int    `json:"gitlab_type"`
	GitlabBranchName   string `json:"gitlab_branch_name"`
	IsAutoRelease      int    `json:"is_auto_release"`
	Description        string `json:"description"`
	GitlabCITemplateID int32  `json:"gitlab_ci_template_id"`
	GitlabID           uint32 `json:"gitlab_id"`
}

// @Summary 用户注册接口
// @title Swagger Example API
// @Tags  用户类接口
// @description  用户注册接口
// @Produce  json
// @Param mobile body string true "13502127317"
// @Param code body	ReleaseTemplateAdd true "JSON数据"
// @Success 200 {string} json "{"code":200,"data":"name","msg":"ok"}"
// @Router /api/v1/user/register [post]
func (u *UserController) actionRegister(context *gin.Context) {
	var formValidator validatorRules.V1UserRegister
	err := context.ShouldBind(&formValidator)
	if err != nil {
		util.ResponseError(context, err)
		return
	}

	//_, _ = service.V1PcRegisterUser()

	util.ResponseOk(context, "test", "test")
	//_, err = util.SendRegisterSms("13502127317", 123456)
	//
	//fmt.Println(err)

}

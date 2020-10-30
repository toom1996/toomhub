// @Description
// @Author    2020/8/19 15:59
package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/medivhzhan/weapp/v2"
	"net/http"
	"time"
	LogicMiniV1 "toomhub/logic/mini/v1"
	"toomhub/middware"
	"toomhub/util"
	validator2 "toomhub/validatorRules"
)

type UserController struct {
	Name string `validate:"required"`
	Age  int    `validate:"gte=0,lte=100"`
}

//当前控制器注册的路由
func (u *UserController) Register(engine *gin.Engine) {
	user := engine.Group("/v1/mini/user")
	{
		//小程序用户登陆接口
		user.POST("/login", u.Login)
		user.POST("/token-checker", u.tokenChecker)
		user.GET("/get-session", u.GetSessionKey)
	}
	user.Use(middware.CheckIdentity())
	{
		user.GET("/get-info", u.refreshInfo)
	}
}

func (u *UserController) GetSessionKey(Context *gin.Context) {
	config := util.GetConfig()
	code, _ := Context.GetQuery("code")

	if code == "" {
		Context.JSON(200, gin.H{
			"code":    400,
			"message": "can't find code",
		})
		return
	}
	wechatLogin, _ := weapp.Login(config.Mini.AppId, config.Mini.AppSecret, code)

	if wechatLogin.ErrCode != 0 {
		Context.JSON(200, gin.H{
			"code":    400,
			"message": wechatLogin.ErrMSG,
		})
		return
	}

	wechatLoginJson, err := json.Marshal(wechatLogin)

	if err != nil {
		Context.JSON(200, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	authKey := util.GetRandomString(10) + fmt.Sprintf("%d", time.Now().Unix())
	_, err = util.Rdb.Set(util.Ctx, authKey, wechatLoginJson, time.Second*7150).Result()

	if err != nil {
		Context.JSON(200, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	fmt.Println(code)
	Context.JSON(200, map[string]interface{}{
		"code":    200,
		"message": "OK",
		"errcode": 0,
		"data": map[string]interface{}{
			"authKey": authKey,
		},
	})
	return
}

// @url 	localhost:8080/mini/login	POST
// @title    小程序用户登陆接口
// @description   初次登陆的用户将会入库并返回信息, 非初次登陆的用户将会返回用户信息
// @auth	toom <1023150697@qq.com>
// @param     Context	*gin.Context
// @return
func (u *UserController) Login(Context *gin.Context) {
	//validator验证
	formValidator := validator2.Login{}
	err := Context.ShouldBind(&formValidator)
	if err != nil {
		Context.String(http.StatusOK, "参数错误:%s", err.Error())
		return
	}

	logic := LogicMiniV1.UserLogic{}
	query, err := logic.Login(&formValidator)

	if err != nil {
		Context.JSON(http.StatusOK, gin.H{
			"code":    400,
			"message": err.Error(),
			"data":    "",
		})
		return
	}

	Context.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "登陆成功",
		"data":    query,
	})
}

// @title	检查token接口
func (u *UserController) tokenChecker(Context *gin.Context) {
	var commonValidator validator2.CommonValidator

	formValidator := validator2.Refresh{}
	err := Context.ShouldBind(&formValidator)
	if err != nil {
		Context.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  commonValidator.TransError(err.(validator.ValidationErrors)),
		})
		return
	}

	formLogic := LogicMiniV1.UserLogic{}
	token, err := formLogic.Check(&formValidator, Context)
	if err != nil {
		Context.JSON(http.StatusOK, gin.H{
			"code": 401,
			"msg":  err.Error(),
		})
		return
	}
	Context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  token,
	})
}

// @title 取用户最新的信息
func (u *UserController) refreshInfo(Context *gin.Context) {
	r, _ := util.Rdb.HMGet(util.Ctx, util.UserCacheKey+fmt.Sprintf("%d", util.GetIdentity().MiniId), []string{"likes_count", "fans_count", "follow_count"}...).Result()
	Context.JSON(200, gin.H{
		"code":    200,
		"message": "OK",
		"data": map[string]interface{}{
			"likes_count":  r[0],
			"fans_count":   r[1],
			"follow_count": r[1],
		},
	})
}

// @Description
// @Author: toom1996 <1023150697@qq.com>
// @dateTime: 2021/1/29 15:46
package logic

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
	rules "toomhub/rules/user/v1"
	"toomhub/service"
	"toomhub/setting"
	"toomhub/util"
)

type UserLogic struct {
}

type token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"` // 这个字段没用到
	Scope       string `json:"scope"`      // 这个字段也没用到
}

func (l *UserLogic) GithubOAuthLogic(validator *rules.V1UserGithubOAuth) (interface{}, error) {
	var ser service.UserService
	var saveInfo map[string]interface{}
	//获取github信息
	info, err := ser.GetGithubOAuthInfo(validator)
	if err != nil {
		return nil, err
	}

	//判断是否存在此用户
	if dbPointer, isNew := ser.IsNewUser(info.ID); isNew == false {
		//存在,更新
		saveInfo, err = ser.UpdateGithubOAuthInfo(dbPointer, &info)
		return map[string]interface{}{
			"avatar":   saveInfo["avatar"],
			"username": saveInfo["username"],
			"token":    "tokentest1234",
		}, err
	} else {
		//不存在,新增
		saveInfo, err = ser.SaveGithubOAuthInfo(&info)
		if err != nil {
			return nil, err
		}

	}
	return map[string]interface{}{
		"avatar":   saveInfo["avatar"],
		"username": saveInfo["username"],
		"token":    "tokentest1234",
	}, nil
}

// 注册、登陆逻辑层
func (l *UserLogic) Register(validator *rules.V1UserRegister) (interface{}, error) {
	r, err := util.Rdb.Get(util.Ctx, validator.Mobile+util.RedisMobileKey).Result()

	fmt.Println(err)
	if err != nil {
		return false, errors.New("验证码已过期")
	}

	if validator.Code != r {
		return false, errors.New("验证码有误")
	}

	var ser service.UserService
	isNew := ser.IsRegister(validator.Mobile)

	fmt.Println(isNew)
	// 未注册过
	if isNew == false {
		fmt.Println("un register")
		r, _ := ser.SaveMobileUser(validator)
		return r, nil
	}

	// 不是新用户
	s, _ := ser.GetMobileUser(isNew.(uint))
	return s, nil

}

// 发送短信逻辑层
func (l *UserLogic) SmsSend(validator *rules.V1UserSmsSend) (bool, error) {
	recordKey := validator.Mobile + "s"
	mobileKey := validator.Mobile + util.RedisMobileKey
	// 验证短信验证码发送时间间隔
	r, err := util.Rdb.PTTL(util.Ctx, recordKey).Result()
	if err != nil {
		return false, err
	}
	if r.Milliseconds() > 0 {
		return false, errors.New("快TM休息一会吧")
	}

	// 验证上次发送的短信验证码是否存在
	r, err = util.Rdb.PTTL(util.Ctx, mobileKey).Result()
	if err != nil {
		return false, err
	}

	var code string
	// key 不存在， 重新生成新的验证码
	if r.Milliseconds() == 0 {
		// 发送验证码
		code = strconv.Itoa(util.GenerateRandomInt(100000, 999999))
		// 验证码存十五分钟
		_, err = util.Rdb.Set(util.Ctx, mobileKey, code, 890*time.Second).Result()
		if err != nil {
			return false, err
		}
	} else {
		code, err = util.Rdb.Get(util.Ctx, mobileKey).Result()
		if err != nil {
			return false, err
		}
	}
	fmt.Println(code)
	// 记录短信验证码信息
	//_, err = util.SendRegisterSms(validator.Mobile, code)
	//if err != nil {
	//	return false, err
	//}
	_, err = util.Rdb.Set(util.Ctx, recordKey, 1, 50*time.Second).Result()
	if err != nil {
		return false, err
	}

	//var ser service.UserService
	//// 判断是否为新用户
	//_, _ = ser.IsRegister(validator.Mobile)
	return true, nil
}

// 刷新token逻辑层
func (l *UserLogic) RefreshToken(validator *rules.V1UserRefreshToken, context *gin.Context) (interface{}, error) {
	expire, _ := strconv.Atoi(setting.ZConfig.Jwt.JwtExpire)
	_, err := util.ParseToken(validator.RefreshToken, context)
	if err != nil {
		return false, err
	}

	// 判断是否为refresh_token
	t, err := util.GetIdentity(context)
	if err != nil {
		return nil, err
	}

	fmt.Println(t)

	xxxx, _ := context.Get("xxxx")
	fmt.Println("xxxx -> ", xxxx)
	fmt.Println(t)
	id, err := strconv.Atoi(t.Id)

	if t.Type != "refresh_token" {
		return false, errors.New("error")
	}
	token, _ := util.GenerateToken(uint(id))
	rt, _ := util.GenerateRefreshToken(uint(id))
	return map[string]interface{}{
		"expire":        expire,
		"issuing_time":  time.Now().Unix(),
		"token":         token,
		"refresh_token": rt,
	}, nil
}

// @Description
// @Author: toom1996 <1023150697@qq.com>
// @dateTime: 2020/12/14 12:41
package util

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"strconv"
	"toomhub/setting"
)

type Sms struct {
	PhoneNumbers  string //电话号码
	SignName      string //模板名称
	TemplateCode  string //模板ID
	TemplateParam int    //模板参数
}

// @param: PhoneNumbers 电话号码
// @param: TemplateParam 验证码数值
func SendRegisterSms(PhoneNumbers string, TemplateParam int) (bool, error) {
	_, err := send(Sms{
		PhoneNumbers:  PhoneNumbers,
		SignName:      "咋哇咋哇",
		TemplateCode:  "SMS_206741473",
		TemplateParam: TemplateParam,
	})

	if err != nil {
		return false, err
	}
	return true, nil
}

//send aliSms
func send(param Sms) (bool, error) {
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", setting.ZConfig.AliSms.AccessKeyId, setting.ZConfig.AliSms.AccessKeySecret)

	if err != nil {
		Debug(fmt.Sprintf("%s", err))
		fmt.Printf("err is %#v\n", err)
	}
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	request.PhoneNumbers = param.PhoneNumbers
	request.SignName = param.SignName
	request.TemplateCode = param.TemplateCode
	request.TemplateParam = "{\"code\":\"" + strconv.Itoa(param.TemplateParam) + "\"}"

	response, err := client.SendSms(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	if response.Code != "200" {
		Debug(fmt.Sprintf("%s", response))
		return false, err
	}
	return true, nil
}

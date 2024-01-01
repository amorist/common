/*
 * Created: 2020-10-09 09:32:31
 * Author : Amor
 * Email : amor90s.zhang@gmail.com
 * -----
 * Description: Technology is how you get to the solution, it is not the solution.
 */

package sms

import (
	"encoding/json"
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/amorist/common/utils/password"
)

const (
	// SignName 短信签名
	SignName string = "云数享"
	// CommonCaptchaTemplateCode 通用验证码模板.
	CommonCaptchaTemplateCode string = "SMS_204190227"
)

type aliyun struct {
	client *dysmsapi.Client
}

// 通用验证码模板参数
type commonCaptchaTemplateParam struct {
	Code string `json:"code"`
}

// NewAliyunSms .
func NewAliyunSms(regionID, accessKeyID, accessKeySecret string) ISMS {
	aliyun := new(aliyun)
	client, err := dysmsapi.NewClientWithAccessKey(regionID, accessKeyID, accessKeySecret)
	if err != nil {
		fmt.Println(err)
	}
	aliyun.client = client
	return aliyun
}

// SendCaptcha 发送验证码
func (aliyun *aliyun) SendCaptcha(kind CodeKind, mobile string) (code string, err error) {
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.PhoneNumbers = mobile
	request.SignName = SignName
	request.TemplateCode = CommonCaptchaTemplateCode
	code = password.RandomString(6, "0")
	templateParam := commonCaptchaTemplateParam{Code: code}
	tp, err := json.Marshal(templateParam)
	if err != nil {
		fmt.Println(err)
		return
	}
	request.TemplateParam = string(tp)
	_, err = aliyun.client.SendSms(request)
	return
}

// SendMsg 发送其他消息
func (aliyun *aliyun) SendMsg(kind MsgKind, mobile string) (msg string, err error) {
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.PhoneNumbers = mobile
	request.SignName = SignName
	// request.TemplateCode = CommonCaptchaTemplateCode
	// code = password.RandomString(6, "A0")
	// templateParam := commonCaptchaTemplateParam{Code: code}
	// tp, err := json.Marshal(templateParam)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// request.TemplateParam = string(tp)
	// _, err = aliyun.client.SendSms(request)
	return
}

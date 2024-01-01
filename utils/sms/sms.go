/*
 * Created: 2020-07-17 18:00:46
 * Author : Amor
 * Email : amor90s.zhang@gmail.com
 * -----
 * Description: Technology is how you get to the solution, it is not the solution.
 */

// Package sms provides
package sms

import "os"

// CodeKind 验证码类型.
type CodeKind uint

// MsgKind 消息类型.
type MsgKind uint

const (
	// LoginCode 登录验证码 .
	LoginCode CodeKind = 1
	// RegisterCode 注册验证码 .
	RegisterCode CodeKind = 2
	// PhoneBindCode 手机号绑定验证码 .
	PhoneBindCode CodeKind = 3
)

// ISMS .
type ISMS interface {
	SendCaptcha(kind CodeKind, mobile string) (code string, err error)
	SendMsg(kind MsgKind, mobile string) (msg string, err error)
}

// NewSms .
func NewSms() ISMS {
	provider := os.Getenv("SMS_PROVIDER")
	var sms ISMS
	switch provider {
	case "aliyun":
		sms = NewAliyunSms(os.Getenv("SMS_REGION_ID"), os.Getenv("ALIYUN_ACCESS_KEY_ID"), os.Getenv("ALIYUN_ACCESS_KEY_SECRET"))
	default:
		sms = NewAliyunSms(os.Getenv("SMS_REGION_ID"), os.Getenv("ALIYUN_ACCESS_KEY_ID"), os.Getenv("ALIYUN_ACCESS_KEY_SECRET"))
	}
	return sms
}

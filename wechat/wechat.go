package wechat

import (
	gowechat "github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
	"github.com/silenceper/wechat/v2/officialaccount"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
)

// Wechat .
type Wechat struct {
	wx *gowechat.Wechat
}

// MiniConfig options.
type MiniConfig struct {
	AppID     string `json:"app_id"`     // appid
	AppSecret string `json:"app_secret"` // appSecret
	Cache     cache.Cache
}

// OffConfig .config for 微信公众号
type OffConfig struct {
	AppID          string `json:"app_id"`           // appid
	AppSecret      string `json:"app_secret"`       // appsecret
	Token          string `json:"token"`            // token
	EncodingAESKey string `json:"encoding_aes_key"` // EncodingAESKey
	Cache          cache.Cache
}

// New Wechat .
func New() *Wechat {
	wx := gowechat.NewWechat()
	return &Wechat{
		wx: wx,
	}
}

// OfficialAccount 获取微信公众号.
func (srv *Wechat) OfficialAccount(config *OffConfig) (officialAccount *officialaccount.OfficialAccount) {
	officialAccount = srv.wx.GetOfficialAccount(&offConfig.Config{
		AppID:          config.AppID,
		AppSecret:      config.AppSecret,
		Token:          config.Token,
		EncodingAESKey: config.EncodingAESKey,
		Cache:          config.Cache,
	})
	return
}

// GetMini 获取微信小程序.
func (srv *Wechat) GetMini(cfg *MiniConfig) (miniProgram *miniprogram.MiniProgram) {
	miniProgram = srv.wx.GetMiniProgram(&miniConfig.Config{
		AppID:     cfg.AppID,
		AppSecret: cfg.AppSecret,
		Cache:     cfg.Cache,
	})
	return
}

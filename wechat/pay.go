package wechat

import (
	"context"
	"fmt"
	"net/http"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/downloader"
	"github.com/wechatpay-apiv3/wechatpay-go/core/notify"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

// PayConfig .
type PayConfig struct {
	AppID            string
	MchID            string
	Key              string
	NotifyURL        string
	CertPath         string
	KeyPath          string
	MchCertificateSN string
}

// Pay .
type Pay struct {
	Config *PayConfig
	Client *core.Client
}

// NewPay .
func NewPay(config *PayConfig) *Pay {
	ctx := context.Background()
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath(config.KeyPath)
	if err != nil {
		fmt.Println("load merchant private key error", err.Error())
		panic(err)
	}
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(
			config.MchID,
			config.MchCertificateSN,
			mchPrivateKey,
			config.Key,
		),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		fmt.Println("new wechat pay client err", err.Error())
		panic(err)
	}
	return &Pay{
		Config: config,
		Client: client,
	}
}

// JSAPI .
// desc: 描述
// amount: 金额 分
func (srv *Pay) JSAPI(desc string, amount int64, openid, orderNo string) (resp *jsapi.PrepayWithRequestPaymentResponse, err error) {
	jsAPIService := jsapi.JsapiApiService{Client: srv.Client}
	ctx := context.Background()
	resp, _, err = jsAPIService.PrepayWithRequestPayment(ctx,
		jsapi.PrepayRequest{
			Appid:       core.String(srv.Config.AppID),
			Mchid:       core.String(srv.Config.MchID),
			Description: core.String(desc),
			OutTradeNo:  core.String(orderNo), // 商户系统内部订单号，只能是数字、大小写字母_-*且在同一个商户号下唯一
			NotifyUrl:   core.String(srv.Config.NotifyURL),
			Amount: &jsapi.Amount{
				Total:    &amount,
				Currency: core.String("CNY"),
			},
			Payer: &jsapi.Payer{
				Openid: core.String(openid),
			},
		},
	)
	return
}

// Notify .
func (srv *Pay) Notify(request *http.Request) (transaction *payments.Transaction, err error) {
	ctx := context.Background()
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath(srv.Config.KeyPath)
	if err != nil {
		return
	}
	// 1. 使用 `RegisterDownloaderWithPrivateKey` 注册下载器
	err = downloader.MgrInstance().RegisterDownloaderWithPrivateKey(ctx, mchPrivateKey, srv.Config.MchCertificateSN, srv.Config.MchID, srv.Config.Key)
	if err != nil {
		return
	}
	// 2. 获取商户号对应的微信支付平台证书访问器
	certificateVisitor := downloader.MgrInstance().GetCertificateVisitor(srv.Config.MchID)
	// 3. 使用证书访问器初始化 `notify.Handler`
	handler := notify.NewNotifyHandler(srv.Config.Key, verifiers.NewSHA256WithRSAVerifier(certificateVisitor))
	transaction = new(payments.Transaction)
	_, err = handler.ParseNotifyRequest(context.Background(), request, transaction)
	// 如果验签未通过，或者解密失败
	if err != nil {
		return
	}
	return
}

// QueryOrderByTransactionID 订单查询.
func (srv *Pay) QueryOrderByTransactionID(rransactionID, mchID string) (resp *payments.Transaction, err error) {
	jsAPIService := jsapi.JsapiApiService{Client: srv.Client}
	ctx := context.Background()
	resp, _, err = jsAPIService.QueryOrderById(ctx,
		jsapi.QueryOrderByIdRequest{
			TransactionId: core.String(rransactionID),
			Mchid:         core.String(mchID),
		},
	)
	if err != nil {
		return
	}
	return
}

// QueryOrderByOutTradeNo 订单查询.
func (srv *Pay) QueryOrderByOutTradeNo(outTradeNo, mchID string) (resp *payments.Transaction, err error) {
	jsAPIService := jsapi.JsapiApiService{Client: srv.Client}
	ctx := context.Background()
	resp, _, err = jsAPIService.QueryOrderByOutTradeNo(ctx,
		jsapi.QueryOrderByOutTradeNoRequest{
			OutTradeNo: core.String(outTradeNo),
			Mchid:      core.String(mchID),
		},
	)
	if err != nil {
		return
	}
	return
}

// refund 退款.
func (srv *Pay) refund() {

}

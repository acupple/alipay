package alipay

import (
	"crypto/rsa"
	"encoding/pem"
	"errors"
	"crypto/x509"
	"github.com/acupple/alipay/enums"
)

const (
	GATEWAY = "https://mapi.alipay.com/gateway.do?"
)

type Alipay struct {
	/**
	 * 签约的支付宝账号对应的支付宝唯一用户号，以2088开头的16位纯数字组成。
	 */
	MerchantId string

	/**
	 * 商户密钥
	 */
	Secret string

	/**
	 * 商户邮箱帐号
	 */
	Email string

	/**
	 * 商户网站使用的编码格式，如utf-8、gbk、gb2312等
	 */
	InputCharset string

	/**
	 * 支付类型
	 */
	PaymentType enums.PaymentType

	/**
	 * 默认支付方式
	 */
	PayMethod enums.PayMethod

	/**
	 * 支付超时时间
	 */
	Expired string

	/**
	 * APP RSA私钥
	 */
	AppPriKey string

	/**
	 * APP RSA公钥
	 */
	AppPubKey string

	/**
	 * 二维码 RSA私钥
	 */
	QrPriKey string

	/**
	 * 二维码 RSA公钥
	 */
	QrPubKey string

	/**
	 * 支付配置参数，不需每次请求都生成
	 */
	PayConfig map[string]string

	/**
	 * 退款配置参数，不需每次请求都生成
	 */
	RefundConfig map[string]string

	Pays Pays

	Verifies Verifies

	Refunds Refunds

	PrivateApiKey *rsa.PrivateKey

	PublicApiKey *rsa.PublicKey
}

func NewAlipay(merchantId string, secret string) Alipay {
	return Alipay{
		MerchantId:   merchantId,
		Secret:       secret,
		InputCharset: "UTF-8",
		PaymentType:  enums.BUY,
		PayMethod:    enums.DIRECT_PAY,
		Expired:      "1h",
	}
}

func (p *Alipay) Start() error {
	err := p.initConfig()
	if err != nil {
		return err
	}
	p.Pays = NewPays(p)
	p.Verifies = NewVerifies(p)
	p.Refunds = NewRefunds(p)

	return nil
}

func (p *Alipay) initConfig() error {
	p.PayConfig = make(map[string]string)
	p.PayConfig[enums.PARTNER] = p.MerchantId
	p.PayConfig[enums.SELLER_ID] = p.MerchantId
	p.PayConfig[enums.PAYMENT_TYPE] = p.PaymentType
	p.PayConfig[enums.IT_B_PAY] = p.Expired
	p.PayConfig[enums.INPUT_CHARSET] = p.InputCharset

	p.RefundConfig = make(map[string]string)
	p.RefundConfig[enums.PARTNER] = p.MerchantId
	p.RefundConfig[enums.INPUT_CHARSET] = p.InputCharset
	p.RefundConfig[enums.SELLER_USER_ID] = p.MerchantId

	if err := p.loadPrivateApiKey(); err != nil {
		return err
	}

	if err := p.loadPublicApiKey(); err != nil {
		return err
	}

	return nil
}

func (p *Alipay)loadPrivateApiKey() error {
	block, _ := pem.Decode(p.AppPriKey)
	if block == nil {
		return errors.New("Malformed PEM file")
	}
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return err
	}

	p.PrivateApiKey = key

	return
}

func (p *Alipay) loadPublicApiKey() error {
	block, _ := pem.Decode(p.PublicApiKey)
	if block == nil {
		return errors.New("Malformed PEM file")
	}

	rawKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}

	p.PublicApiKey = rawKey.(*rsa.PublicKey)

	return
}

func (p *Alipay) ToString() string {
	return "Alipay{" +
		"merchantId='" + p.MerchantId + '\'' +
		", secret='" + p.Secret + '\'' +
		", inputCharset='" + p.InputCharset + '\'' +
		", paymentType='" + p.PaymentType + '\'' +
		", payMethod='" + p.PayMethod + '\'' +
		", expired='" + p.Expired + '\'' +
		", appPriKey='" + p.AppPriKey + '\'' +
		", appPubKey='" + p.AppPubKey + '\'' +
		", qrPriKey='" + p.QrPriKey + '\'' +
		", qrPubKey='" + p.QrPubKey + '\'' +
		'}'
}

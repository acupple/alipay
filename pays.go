package alipay

import (
	"bytes"
	"errors"
	"github.com/acupple/alipay/models/pay"
	"github.com/acupple/alipay/enums"
)

type Pays struct {
	Component
}

func NewPays(alipay Alipay) Pays {
	return Pays{
		Component.Alipay: alipay,
	}
}

func (p *Pays) WebPay(detail pay.WabPayDetail)  (string, error) {
	payParams, err := p.buildWebPayParams(detail)
	if err != nil {
		return nil, err
	}

	return p.buildPayForm(payParams), nil
}

func (p *Pays) buildWebPayParams(detail pay.WabPayDetail) (map[string]string, error) {
	payParams, err := p.buildPayParams(pay.PayDetail(detail), enums.WEB_PAY)
	if err != nil {
		return nil, err
	}

	if detail.ExterInvokeIp != "" {
		payParams[enums.EXTER_INVOKE_IP] = detail.ExterInvokeIp
	}

	if detail.ErrorNotifyUrl != "" {
		payParams[enums.ERROR_NOTIFY_URL] = detail.ErrorNotifyUrl
	}

	//MD5签名
	p.BuildMD5SignParams(payParams)

	return payParams, nil
}

func (p *Pays)WapPay(detail pay.WapPayDetail) (string, error) {
	payParams, err := p.buildWapPayParams(detail)
	if err != nil {
		return nil, err
	}

	return p.buildPayForm(payParams), nil
}

func (p *Pays)buildWapPayParams(detail pay.WapPayDetail) (map[string]string, error) {
	payParams, err := p.buildPayParams(pay.PayDetail(detail), enums.WAP_PAY)
	if err != nil {
		return nil, err
	}

	if detail.ShowUrl != "" {
		payParams[enums.SHOW_URL] = detail.ShowUrl
	}

	if detail.RnCheck != "" {
		payParams[enums.RN_CHECK] = detail.RnCheck
	}

	if detail.AirTicket != "" {
		payParams[enums.AIR_TICKET] = detail.AirTicket
	}

	if detail.BuyerCertNo != "" {
		payParams[enums.BUYER_CERT_NO] = detail.BuyerCertNo
	}

	if detail.BuyerRealName != "" {
		payParams[enums.BUYER_REAL_NAME] = detail.BuyerRealName
	}

	if detail.ExternToken != "" {
		payParams[enums.EXTERN_TOKEN] = detail.ExternToken
	}

	if detail.OtherFee != "" {
		payParams[enums.OTHER_FEE] = detail.OtherFee
	}

	if detail.Scene != "" {
		payParams[enums.SCENE] = detail.Scene
	}

	p.BuildMD5SignParams(payParams)

	return payParams
}

func (p *Pays)AppPay(detail pay.AppPayDetail) (string, error) {
	if p.Alipay.PrivateApiKey == nil {
		return errors.New("app private key not load correctly")
	}

	payParams, err := p.buildAppPayParams(detail)
	if err != nil {
		return nil, err
	}

	return p.BuildSignStringWithWrapChar(payParams, "\"")
}

func (p *Pays)buildAppPayParams(detail pay.AppPayDetail) (map[string]string, error) {
	payParams, err := p.buildPayParams(pay.PayDetail(detail), enums.APP_PAY)
	if err != nil {
		return nil, err
	}

	//APP支付无return_url
	delete(payParams, enums.RETURN_URL)

	if detail.Body == "" {
		return nil, errors.New("body required")
	}

	payParams[enums.BODY] = detail.Body
	if detail.AppId != "" {
		payParams[enums.APP_ID] = detail.AppId
	}

	if detail.Appenv != "" {
		payParams[enums.APPENV] = detail.Appenv
	}

	if detail.ExternToken != "" {
		payParams[enums.EXTERN_TOKEN] = detail.ExternToken
	}

	if detail.OutContext != "" {
		payParams[enums.OUT_CONTEXT] = detail.OutContext
	}

	if detail.RnCheck != "" {
		payParams[enums.RN_CHECK] = detail.RnCheck
	}

	if detail.GoodsType != "" {
		payParams[enums.GOODS_TYPE] = detail.GoodsType
	}

	return payParams
}

func (p *Pays)buildPayParams(detail pay.PayDetail, service enums.ServiceName) (map[string]string, error) {
	if detail.OutTradeNo == "" || detail.OrderName == "" || detail.TotalFee == "" {
		return nil, errors.New("OutTradeNo, OrderName or TotalFee is empty")
	}
	payParams := make(map[string]string)
	for key, value := range Alipay.PayConfig {
		payParams[key] = value
	}

	payParams[enums.SERVICE] = service
	payParams[enums.OUT_TRADE_NO] = detail.OutTradeNo
	payParams[enums.ORDER_NAME] = detail.OrderName
	payParams[enums.TOTAL_FEE] = detail.TotalFee

	if detail.NotifyUrl != "" {
		payParams[enums.NOTIFY_URL] = detail.NotifyUrl
	}

	if detail.ReturnUrl != "" {
		payParams[enums.RETURN_URL] = detail.ReturnUrl
	}

	return payParams, nil
}

func (p *Pays)buildPayForm(payParams map[string]string) string {
	var buffer bytes.Buffer
	buffer.WriteString("<form id=\"pay_form\" name=\"pay_form\"")
	buffer.WriteString(" action=\"" + GATEWAY)
	buffer.WriteString(enums.INPUT_CHARSET + "=" + Alipay.InputCharset)
	buffer.WriteString("\" method=\"POST\">")

	for key, value := range payParams {
		buffer.WriteString("<input type=\"hidden\" name=\"" + key + "\" value=\"" + value + "\" />")
	}

	buffer.WriteString("<input type=\"submit\" value=\"去支付\" style=\"display:none;\" />")
	buffer.WriteString("</form>")
	buffer.WriteString("<script>document.forms['pay_form'].submit();</script>")

	return buffer.String()
}

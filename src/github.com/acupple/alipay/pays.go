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
		return nil
	}

	return p.buildPayForm(payParams), nil
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

	return payParams, nil
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

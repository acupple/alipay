package alipay

import (
	"time"
	"errors"
	"net/http"
	"net/url"
	"io/ioutil"
	"github.com/acupple/alipay/enums"
	"github.com/acupple/alipay/models/refund"
)

type Refunds struct {
	Component
}

func NewRefunds(alipay Alipay) Refunds {
	return Pays{
		Component.Alipay: alipay,
	}
}

func (r *Refunds)Refund(detail refund.RefundDetail) (bool, error) {
	refundPrams, err := r.buildRefundParams(detail)

	if err != nil {
		return false, err
	}
	url,_ := url.Parse(GATEWAY + "_input_charset=" + r.Alipay.InputCharset)
	q := url.Query()

	for key, value := range refundPrams {
		q.Set(key, value)
	}

	url.RawQuery = q.Encode()
	resp, err := http.Get(url.String());
	if err != nil {
		return false, err
	}

	if resp.StatusCode != http.StatusOK {
		return false, errors.New("Bad refund request")
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	if string(bytes) == "T" {
		return true, nil
	}

	return false, nil
}

func (r *Refunds)buildRefundParams(detail refund.RefundDetail) (map[string]string, error){
	if detail.BatchNo == "" {
		return nil, errors.New("Refund Batch No. is empty")
	}

	if detail.DetailDatas == nil || len(detail.DetailDatas) <= 0 {
		return nil, errors.New("Refund detail datas are empty")
	}
	refundPrams := make(map[string]string)
	for key, value := range r.Alipay.RefundConfig {
		refundPrams[key] = value
	}

	refundPrams[enums.SERVICE] = enums.REFUND_NO_PWD
	if r.Alipay.Email != "" {
		refundPrams[enums.SELLER_EMAIL] = r.Alipay.Email
	}

	if detail.NotifyUrl != "" {
		refundPrams[enums.NOTIFY_URL] = detail.NotifyUrl
	}

	refundPrams[enums.SELLER_USER_ID] = r.Alipay.MerchantId
	refundPrams[enums.REFUND_DATE] = time.Now().Format("yyyy-MM-dd HH:mm:ss")

	refundPrams[enums.BATCH_NO] = detail.BatchNo
	refundPrams[enums.BATCH_NUM] = len(detail.DetailDatas)

	refundPrams[enums.DETAIL_DATA] = detail.FormatDetailDatas()

	r.BuildMD5SignParams(refundPrams)

	return refundPrams, nil
}

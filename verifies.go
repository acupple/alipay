package alipay

import (
	"net/http"
	"io/ioutil"
	"errors"
	"strings"
	"github.com/acupple/alipay/enums"
)

type Verifies struct {
	Component
}

func NewVerifies(alipay Alipay) Verifies {
	return Pays{
		Component.Alipay: alipay,
	}
}

func (v *Verifies) MD5 (notifyParams map[string]string) bool {
	if notifyParams == nil || notifyParams[enums.SIGN] == "" {
		return false
	}
	validParams := v.filterSigningParams(notifyParams)
	signing := v.BuildSignString(validParams)
	signed := v.md5(signing)

	return notifyParams[enums.SIGN] == signed
}

func (v *Verifies) filterSigningParams(signingParams map[string]string) (map[string]string) {
	validParams := make(map[string]string)
	for key, value := range signingParams {
		if value != "" || key == enums.SIGN || key == enums.SIGN_TYPE {
			continue
		}

		validParams[key] = value
	}

	return validParams
}

func (v *Verifies) RSA(notifyParams map[string]string) (bool, error) {
	if notifyParams == nil || notifyParams[enums.SIGN] == "" {
		return false
	}

	if v.Alipay.PublicApiKey == nil {
		return false, errors.New("Public key not load correctly")
	}

	validParams := v.filterSigningParams(notifyParams)
	signing := v.BuildSignString(validParams)

	return v.rsaVerify(signing, notifyParams[enums.SIGN])
}

func (v *Verifies)NotifyId(notifyId string) (bool, error) {
	url := GATEWAY + "&" +
		enums.SERVICE + "=" + enums.NOTIFY_VERIFY + "&" +
		enums.PARTNER + "=" + v.Alipay.MerchantId + "&" +
		enums.NOTIFY_ID + "=" + notifyId

	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		return false, errors.New("Request failed")
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	return strings.ToLower(string((bytes))) == "true"
}

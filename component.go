package alipay

import (
	"sort"
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"crypto/rsa"
	"crypto/rand"
	"crypto/sha1"
	"crypto"
	"encoding/base64"
	"github.com/acupple/alipay/enums"
)

type Component struct {
	Alipay Alipay
}

func (c *Component) BuildMD5SignParams(payParams map[string]string) error {
	payString := c.BuildSignString(payParams)

	payParams[enums.SIGN] = c.md5(payString)
	payParams[enums.SIGN_TYPE] = enums.MD5

	return nil
}

func (c *Component) BuildRsaPayString(payParams map[string]string) error {
	payString := c.BuildSignStringWithWrapChar(payParams, "\"")
	sign, err := c.rsa(payString)
	if err != nil {
		return err
	}

	payParams[enums.SIGN] = sign
	payParams[enums.SIGN_TYPE] = enums.RSA

	return nil
}

func (c *Component)md5(payString string) string {
	hasher := md5.New()
	hasher.Write([]byte(payString))
	return hex.EncodeToString(hasher.Sum(nil))
}

func (c *Component)rsa(payString string) (string, error) {
	hashSum := sha1.Sum([]byte(payString))
	signed, err := rsa.SignPKCS1v15(rand.Reader, c.Alipay.PrivateApiKey, crypto.SHA1, hashSum[:])
	if err != nil {
		return nil, err
	}

	return base64.StdEncoding.EncodeToString(signed)
}

func (c *Component)rsaVerify(signing string, signed string) (bool, error) {
	hashSum := sha1.Sum([]byte(signing))

	err := rsa.VerifyPKCS1v15(c.Alipay.PrivateApiKey, crypto.SHA1, hashSum[:], []byte(signed))

	if err != nil {
		return false, err
	}

	return false
}

func (c *Component)BuildSignStringWithWrapChar(payParams map[string]string, wrapChar string) string {
	keys := make([]string, len(payParams))
	i :=0
	for key,  _ := range payParams {
		keys[i] = key
		i ++
	}
	sort.Strings(keys)

	var buffer bytes.Buffer
	for i := 0; i< len(keys) ; i ++ {
		value := payParams[keys[i]]
		if i == len(keys) - 1 {
			buffer.WriteString(keys[i] + "=" + wrapChar + value + wrapChar)
		} else {
			buffer.WriteString(keys[i] + "=" + wrapChar + value + wrapChar + "&")
		}
	}

	return buffer.String()
}

func (c *Component)BuildSignString(payParams map[string]string) string {
	return c.BuildSignStringWithWrapChar(payParams, "")
}

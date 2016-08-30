package pay

import "github.com/acupple/alipay/enums"

type AppPayDetail struct {
	PayDetail
	/**
	 * 客户端号，标识客户端
	 */
	AppId string

	/**
	 * 客户端来源
	 */
	Appenv string

	/**
	 * 是否发起实名校验
	 */
	RnCheck string

	/**
	 * 授权令牌(32)
	 */
	ExternToken string

	/**
	 * 商户业务扩展参数
	 */
	OutContext string

	/**
	 * 商品详情
	 */
	Body string

	/**
	 * 商品类型
	 */
	GoodsType enums.GoodsType
}

func (m *AppPayDetail) NewAppPayDetail(outTradeNo string, orderName string, totalFee string, body string) AppPayDetail {
	return AppPayDetail{
		PayDetail.OutTradeNo: outTradeNo,
		PayDetail.OrderName:  orderName,
		PayDetail.TotalFee:   totalFee,
		Body:                 body,
	}
}

func (m *AppPayDetail) ToString() string {
	return "AppPayDetail{" +
		"appId='" + m.AppId + '\'' +
		", appenv='" + m.Appenv + '\'' +
		", rnCheck='" + m.RnCheck + '\'' +
		", externToken='" + m.ExternToken + '\'' +
		", outContext='" + m.OutContext + '\'' +
		", body='" + m.Body + '\'' +
		", goodsType=" + m.GoodsType +
		"} " + m.PayDetail.ToString()
}

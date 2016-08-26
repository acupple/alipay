package pay

type PayDetail struct {
	/**
	 * 我方唯一订单号
	 */
	OutTradeNo string

	/**
	 * 商品名称
	 */
	OrderName string

	/**
	 * 商品金额(元)
	 */
	TotalFee string

	/**
	 * 支付宝后置通知url，若为空，则使用Alipay类中的notifyUrl
	 */
	NotifyUrl string

	/**
	 * 支付宝前端跳转url，若为空，则使用Alipay类中的returnUrl
	 */
	ReturnUrl string
}

func (m *PayDetail) ToString() string {
	return "PayFields{" +
		"outTradeNo='" + m.OutTradeNo + '\'' +
		", orderName='" + m.OrderName + '\'' +
		", totalFee='" + m.TotalFee + '\'' +
		", notifyUrl='" + m.NotifyUrl + '\'' +
		", returnUrl='" + m.ReturnUrl + '\'' +
		'}'
}

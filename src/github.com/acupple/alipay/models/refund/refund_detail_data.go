package refund

type RefundDetailData struct {
	/**
	* 支付宝交易号
	 */
	TradeNo string

	/**
	 * 退款金额(元)
	 */
	Fee string

	/**
	 * 退款原因
	 */
	Reason string
}

/**
* 格式化为支付宝需要的格式: tradeNo^fee^reason
 */
func (m *RefundDetailData) Format() string {
	return m.TradeNo + "^" + m.Fee + "^" + m.Reason
}

func (m *RefundDetailData) ToString() {
	return "RefundDetailData: {" +
		"tradeNo='" + m.TradeNo + '\'' +
		", fee='" + m.Fee + '\'' +
		", reason='" + m.Reason + '\'' +
		'}'
}

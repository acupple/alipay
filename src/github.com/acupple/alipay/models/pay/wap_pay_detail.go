package pay

type WapPayDetail struct {
	PayDetail
	/**
	 * 商品展示网址
	 */
	ShowUrl string

	/**
	 * 手机支付宝token
	 */
	ExternToken string

	/**
	 * 航旅订单其它费用
	 */
	OtherFee string

	/**
	 * 航旅订单金额
	 */
	AirTicket string

	/**
	 * 是否发起实名校验
	 */
	RnCheck string

	/**
	 * 买家证件号码
	 */
	BuyerCertNo string

	/**
	 * 买家真实姓名
	 */
	BuyerRealName string

	/**
	 * 收单场景
	 */
	Scene string
}

func (m *WapPayDetail) ToString() {
	return "WapPayFields{" +
		"showUrl='" + m.ShowUrl + '\'' +
		", externToken='" + m.ExternToken + '\'' +
		", otherFee='" + m.OtherFee + '\'' +
		", airTicket='" + m.AirTicket + '\'' +
		", rnCheck='" + m.RnCheck + '\'' +
		", buyerCertNo='" + m.BuyerCertNo + '\'' +
		", buyerRealName='" + m.BuyerRealName + '\'' +
		", scene='" + m.Scene + '\'' +
		"} " + m.PayDetail.ToString()
}

package pay

type WabPayDetail struct {
	PayDetail
	/**
	 * 客服端IP
	 */
	ExterInvokeIp string

	/**
	 * 支付宝错误通知跳转
	 */
	ErrorNotifyUrl string
}

func (m *WabPayDetail) ToString() {
	return "WebPayFields{" +
		"exterInvokeIp='" + m.ExterInvokeIp + '\'' +
		", errorNotifyUrl='" + m.ErrorNotifyUrl + '\'' +
		"} " + m.PayDetail.ToString()
}

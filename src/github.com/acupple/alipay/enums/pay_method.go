package enums

type PayMethod string

const (
	/**
	* 信用支付
	 */
	CREDIT_PAY PayMethod = "creditPay"

	/**
	 * 余额支付
	 */
	DIRECT_PAY PayMethod = "directPay"
)

func (p PayMethod) Humanize() string {
	switch p {
	case CREDIT_PAY:
		return "信用卡支付"
	case DIRECT_PAY:
		return "余额支付"
	default:
		return ""
	}
}

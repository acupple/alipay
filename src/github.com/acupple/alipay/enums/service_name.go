package enums

type ServiceName string

const (
	/**
	* WEB支付
	 */
	WEB_PAY ServiceName = "create_direct_pay_by_user"

	/**
	 * WAP支付
	 */
	WAP_PAY ServiceName = "alipay.wap.create.direct.pay.by.user"

	/**
	 * APP支付
	 */
	APP_PAY ServiceName = "mobile.securitypay.pay"

	/**
	 * 无密退款
	 */
	REFUND_NO_PWD ServiceName = "refund_fastpay_by_platform_nopwd"

	/**
	 * 支付宝通知校验
	 */
	NOTIFY_VERIFY ServiceName = "notify_verify"
)

func (p ServiceName) Humanize() string {
	switch p {
	case WEB_PAY:
		return "WEB支付"
	case WAP_PAY:
		return "WAP支付"
	case APP_PAY:
		return "APP支付"
	case REFUND_NO_PWD:
		return "无密退款"
	case NOTIFY_VERIFY:
		return "支付宝通知校验"
	default:
		return ""
	}
}

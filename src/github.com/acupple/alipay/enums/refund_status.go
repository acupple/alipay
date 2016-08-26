package enums

type RefundStatus string

const (
	/**
	退款成功:
		全额退款情况：trade_status = TRADE_CLOSED，而refund_status=REFUND_SUCCESS；
		非全额退款情况：trade_status = TRADE_SUCCESS，而refund_status=REFUND_SUCCESS
	*/
	REFUND_SUCCESS RefundStatus = "REFUND_SUCCESS"

	/**
	 * 退款关闭
	 */
	REFUND_CLOSED RefundStatus = "REFUND_CLOSED"
)

func (p RefundStatus) Humanize() string {
	switch p {
	case REFUND_SUCCESS:
		return "退款成功"
	case REFUND_CLOSED:
		return "退款关闭"
	default:
		return ""
	}
}

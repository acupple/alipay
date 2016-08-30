package enums

type PaymentType string

const (
	BUY         PaymentType = "1"
	DONATE      PaymentType = "47"
	CARD_COUPON PaymentType = "47"
)

func (p PaymentType) Humanize() string {
	switch p {
	case BUY:
		return "购买商品"
	case DONATE:
		return "捐赠"
	case CARD_COUPON:
		return "充值卡或抵扣券"
	default:
		return ""
	}
}

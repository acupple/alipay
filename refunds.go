package alipay

type Refunds struct {
	Component
}

func NewRefunds(alipay Alipay) Refunds {
	return Pays{
		Component.Alipay: alipay,
	}
}

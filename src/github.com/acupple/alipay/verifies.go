package alipay

type Verifies struct {
	Component
}

func NewVerifies(alipay Alipay) Verifies {
	return Pays{
		Component.Alipay: alipay,
	}
}

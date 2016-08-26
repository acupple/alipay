package enums

type GoodsType string

const (
	/**
	 * 虚拟物品
	 */
	VIRTUAL GoodsType = "0"

	/**
	 * 实物
	 */
	REAL GoodsType = "1"
)

func (p GoodsType) Humanize() string {
	switch p {
	case VIRTUAL:
		return "虚拟物品"
	case REAL:
		return "实物"
	default:
		return ""
	}
}

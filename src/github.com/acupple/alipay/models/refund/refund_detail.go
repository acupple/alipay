package refund

type RefundDetail struct {
	NotifyUrl string

	/**
	 * 退款批次号
	 */
	BatchNo string

	/**
	 * 单笔数据集
	 */
	DetailDatas []RefundDetailData
}

func (m *RefundDetail) FormatDetailDatas() string {
	first := true
	details := ""
	for _, data := range m.DetailDatas {
		if first {
			details += data.Format()
			first = false
		} else {
			details += "#" + data.Format()
		}
	}

	return details
}

func (m *RefundDetail) ToString() string {
	return "RefundDetail{" +
		"notifyUrl='" + m.NotifyUrl + '\'' +
		", batchNo='" + m.BatchNo + '\'' +
		", detailDatas=" + m.DetailDatas +
		'}'
}

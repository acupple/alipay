package enums

type AlipayField string

/**
https://doc.open.alipay.com/docs/doc.htm?spm=a219a.7629140.0.0.HgAv4j&treeId=62&articleId=104743&docType=1
*/
const (
	PARTNER    AlipayField = "partner"
	PAY_METHOD AlipayField = "paymethod"

	IT_B_PAY AlipayField = "it_b_pay"

	INPUT_CHARSET AlipayField = "_input_charset"

	NOTIFY_URL AlipayField = "notify_url"

	RETURN_URL AlipayField = "return_url"

	/**
	 * 设置请求出错时的通知页面路径(该功能需要联系支付宝开通)
	 * 当商户通过该接口发起请求时，如果出现提示报错，
	 * 支付宝会根据请求出错时的通知错误码通过异步的方式发送通知给商户。
	 */
	ERROR_NOTIFY_URL AlipayField = "error_notify_url"

	/**
	 * 用户在创建交易时，该用户当前所使用机器的IP。如果商户申请后台开通防钓鱼IP地址检查选项，此字段必填，校验用。
	 */
	EXTER_INVOKE_IP AlipayField = "exter_invoke_ip"

	/**
	 * 收银台页面上，商品展示的超链接
	 */
	SHOW_URL AlipayField = "show_url"

	ORDER_NAME AlipayField = "order_name"

	/**
	 * T发起，F不发
	 */
	RN_CHECK AlipayField = "rn_check"

	/**
	 * 航旅订单金额描述，由四项或两项构成，各项之间由“|”分隔，
	 * 每项包含金额与描述，金额与描述间用“^”分隔，票面价之外的价格之和必须与otherfee相等。
	 */
	AIR_TICKET AlipayField = "air_ticket"

	/**
	 * 航旅订单中除去票面价之外的费用，单位为RMB-Yuan。
	 *  取值范围为[0.01,100000000.00]，精确到小数点后两位。
	 */
	OTHER_FEE AlipayField = "other_fee"

	/**
	 * 接入极简版wap收银台时支持。当商户请求是来自手机支付宝，在手机支付宝登录后，
	 *  有生成登录信息token时，使用该参数传入token将可以实现信任登录收银台，不需要再次登录。
	 *  注意：登录后用户还是有入口可以切换账户，不能使用该参数锁定用户。
	 */
	EXTERN_TOKEN AlipayField = "extern_token"

	/**
	 * 需要与支付宝实名认证时所填写的证件号码一致
	 */
	BUYER_CERT_NO AlipayField = "buyer_cert_no"

	/**
	 * 需要与支付宝实名认证时所填写的证件号码一致
	 */
	BUYER_REAL_NAME AlipayField = "buyer_real_name"

	/**
	 * 如需使用该字段，需向支付宝申请开通，否则传入无效。
	 */
	SCENE AlipayField = "scene"

	SIGN AlipayField = "sign"

	SIGN_TYPE AlipayField = "sign_type"

	OUT_TRADE_NO AlipayField = "out_trade_no"

	SUBJECT AlipayField = "subject"

	PAYMENT_TYPE AlipayField = "payment_type"

	/**
	 * 有效性一般为1分钟
	 */
	NOTIFY_ID AlipayField = "notify_id"

	NOTIFY_TIYE AlipayField = "notify_type"

	/**
	 * yyyy-MM-dd HH:mm:ss
	 */
	NOTIFY_TIME AlipayField = "notify_time"

	TRADE_NO AlipayField = "trade_no"

	TRADE_STATUS AlipayField = "trade_status"

	SELLER_ID AlipayField = "seller_id"

	/**
	 * 卖家支付宝账号对应的支付宝唯一用户号。以2088开头的纯16位数字。
	 * 登录时，seller_email和seller_user_id两者必填一个。如果两者都填，以seller_user_id为准。
	 */
	SELLER_USER_ID AlipayField = "seller_user_id"

	SELLER_EMAIL AlipayField = "seller_email"

	BUYER_ID AlipayField = "buyer_id"

	BUYER_EMAIL AlipayField = "buyer_email"

	/**
	 * 取值范围为[0.01，100000000.00]，精确到小数点后两位
	 */
	TOTAL_FEE AlipayField = "total_fee"

	GMT_CREATE AlipayField = "gmt_create"

	GMT_PAYMENT AlipayField = "gmt_payment"

	GMT_CLOSE AlipayField = "gmt_close"

	GMT_REFUND AlipayField = "gmt_refund"

	REFUND_STATUS AlipayField = "refund_status"

	/**
	 * yyyy-MM-dd HH:mm:ss
	 */
	REFUND_DATE AlipayField = "refund_date"

	/**
	 * 如果请求时使用的是total_fee，那么price等于total_fee；如果请求时使用的是price，那么对应请求时的price参数，原样通知回来。
	 */
	PRICE AlipayField = "price"

	/**
	 * 如果请求时使用的是total_fee，那么quantity等于1；如果请求时使用的是quantity，那么对应请求时的quantity参数，原样通知回来。
	 */
	QUANTITY AlipayField = "quantity"

	/**
	 * 该笔订单的备注、描述、明细等。对应请求时的body参数，原样通知回来。
	 */
	BODY AlipayField = "body"

	/**
	 * 支付宝系统会把discount的值加到交易金额上，如果需要折扣，本参数为负数。
	 */
	DISCOUNT AlipayField = "discount"

	IS_TOTAL_FEE_ADJUST AlipayField = "is_total_fee_adjust"

	USE_COUPON AlipayField = "use_coupon"

	/**
	 * 该笔交易所使用的支付渠道。格式为：渠道1|渠道2|…，如果有多个渠道，用“|”隔开
	 */
	OUT_CHANNEL_TYPE AlipayField = "out_channel_type"

	/**
	 * 该笔交易通过使用各支付渠道所支付的金额。格式为：金额1|金额2|…，如果有多个支付渠道，各渠道所支付金额用“|”隔开。
	 */
	OUT_CHANNEL_AMOUNT AlipayField = "out_channel_amount"

	/**
	 * 该交易支付时实际使用的银行渠道。格式为：支付渠道1|支付渠道2|…，如果有多个支付渠道，用“|”隔开
	 */
	OUT_CHANNEL_INST AlipayField = "out_channel_inst"

	/**
	 * 回传给商户此标识为qrpay时，表示对应交易为扫码支付。目前只有qrpay一种回传值。非扫码支付方式下，目前不会返回该参数。
	 */
	BUSINESS_SCENE AlipayField = "business_scene"

	/**
	 * 用于商户回传参数
	 *  该值不能包含等号、与号等特殊字符。如果用户请求时传递了该参数，则返回给商户时会回传该参数。
	 */
	EXTRA_COMMON_PARAM AlipayField = "extra_common_param"

	/**
	 * 表示接口调用是否成功，并不表明业务处理结果，如T
	 */
	IS_SUCCESS AlipayField = "is_success"

	/**
	 * 标志调用哪个接口返回的链接，如create_direct_pay_by_user
	 */
	EXTERFACE AlipayField = "exterface"

	/**
	 * 标志调用哪个接口返回的链接，如alipay.wap.create.direct.pay.by.user
	 */
	SERVICE AlipayField = "service"

	/**
	 * 本参数用于信用支付。它代表执行支付操作的操作员账号所属的代理人的支付宝唯一用户号。以2088开头的纯16位数字。
	 */
	AGENT_USER_ID AlipayField = "agent_user_id"

	ERROR_CODE AlipayField = "error_code"

	APP_ID AlipayField = "app_id"

	APPENV AlipayField = "appenv"

	/**
	 * 业务扩展参数，支付宝特定的业务需要添加该字段，json格式。 商户接入时和支付宝协商确定。
	 */
	OUT_CONTEXT AlipayField = "out_context"

	/**
	 * 进行一次即时到账批量退款，都需要提供一个批次号，通过该批次号可以查询这一批次的退款交易记录，
	 * 对于每一个合作伙伴，传递的每一个批次号都必须保证唯一性。格式为：退款日期（8位）+流水号（3～24位）。
	 * 不可重复，且退款日期必须是当天日期。流水号可以接受数字或英文字符，建议使用数字，但不可接受“000”。
	 */
	BATCH_NO AlipayField = "batch_no"

	/**
	 * 即参数detail_data的值中，“#”字符出现的数量加1，最大支持1000笔（即“#”字符出现的最大数量为999个）
	 */
	BATCH_NUM AlipayField = "batch_num"

	/**
	 * 退款请求的明细数据:
	 * 1. 数据集格式: 第一笔#第二笔#...#第N笔
	 * 2. 单笔格式: 原付款支付宝交易号^退款总金额^退款理由
	 * 3. 不支持退分润功能
	 */
	DETAIL_DATA AlipayField = "detail_data"

	/**
	 * 0 小于 success_num 小于等于 总退款笔数(batch_num)
	 */
	SUCCESS_NUM AlipayField = "success_num"

	/**
	 * 1. 退手续费结果返回格式:
	 *  交易号^退款金额^处理结果$退费账号^退费账户ID^退费金额^处理结果
	 * 2. 不退手续费结果返回格式：
	 *  交易号^退款金额^处理结果
	 * 3. 若退款申请提交成功，处理结果会返回“SUCCESS”。若提交失败，退款的处理结果中会有报错码
	 */
	RESULT_DETAILS AlipayField = "result_details"

	GOODS_TYPE AlipayField = "goods_type"
)

func (p AlipayField) Humanize() string {
	switch p {
	case PARTNER:
		return "商户ID"

	case PAY_METHOD:
		return "支付类型"

	case IT_B_PAY:
		return "超时设置"

	case INPUT_CHARSET:
		return "字符集设置"

	case NOTIFY_URL:
		return "支付宝服务器通知地址"

	case RETURN_URL:
		return "支付宝页面跳转地址"

	case ERROR_NOTIFY_URL:
		return "支付宝错误通知地址"

	case EXTER_INVOKE_IP:
		return "客户端IP"

	case SHOW_URL:
		return "商品展示链接"

	case ORDER_NAME:
		return "商品名称"

	case RN_CHECK:
		return "是否发起实名校验"

	case AIR_TICKET:
		return "航旅订单金额"

	case OTHER_FEE:
		return "航旅订单其它费用"

	case EXTERN_TOKEN:
		return "手机支付宝token"

	case BUYER_CERT_NO:
		return "买家证件号码"

	case BUYER_REAL_NAME:
		return "买家真实姓名"

	case SCENE:
		return "收单场景"

	case SIGN:
		return "签名"

	case SIGN_TYPE:
		return "签名类型"

	case OUT_TRADE_NO:
		return "我方唯一订单号"

	case SUBJECT:
		return "商品名称"

	case PAYMENT_TYPE:
		return "支付类型"

	case NOTIFY_ID:
		return "通知校验ID"

	case NOTIFY_TIYE:
		return "通知类型"

	case NOTIFY_TIME:
		return "通知时间"

	case TRADE_NO:
		return "支付宝交易号"

	case TRADE_STATUS:
		return "交易状态"

	case SELLER_ID:
		return "卖家支付宝账户号"

	case SELLER_USER_ID:
		return "卖家支付宝账户号"

	case SELLER_EMAIL:
		return "卖家支付宝帐号"

	case BUYER_ID:
		return "买家支付宝账户号"

	case BUYER_EMAIL:
		return "买家支付宝帐号"

	case TOTAL_FEE:
		return "订单的总金额"

	case GMT_CREATE:
		return "交易创建时间"

	case GMT_PAYMENT:
		return "交易付款时间"

	case GMT_CLOSE:
		return "交易关闭时间"

	case GMT_REFUND:
		return "退款时间"

	case REFUND_STATUS:
		return "退款状态"

	case REFUND_DATE:
		return "退款请求时间"

	case PRICE:
		return "商品单价"

	case QUANTITY:
		return "购买数量"

	case BODY:
		return "商品描述"

	case DISCOUNT:
		return "折扣"

	case IS_TOTAL_FEE_ADJUST:
		return "该交易是否调整过价格"

	case USE_COUPON:
		return "买家是否使用红包"

	case OUT_CHANNEL_TYPE:
		return "支付渠道组合信息"

	case OUT_CHANNEL_AMOUNT:
		return "支付金额组合信息"

	case OUT_CHANNEL_INST:
		return "实际支付渠道"

	case BUSINESS_SCENE:
		return "是否扫码支付"

	case EXTRA_COMMON_PARAM:
		return "公用回传参数"

	case IS_SUCCESS:
		return "成功标识"

	case EXTERFACE:
		return "服务接口名称"

	case SERVICE:
		"接口名称"
		return

	case AGENT_USER_ID:
		return "信用支付购票员的代理人ID"

	case ERROR_CODE:
		return "错误码"

	case APP_ID:
		return "客户端号"

	case APPENV:
		return "客户端来源"

	case OUT_CONTEXT:
		return "商户业务扩展参数"

	case BATCH_NO:
		return "退款批次号"

	case BATCH_NUM:
		return "总笔数"

	case DETAIL_DATA:
		return "单笔数据集"

	case SUCCESS_NUM:
		return "退款成功总数"

	case RESULT_DETAILS:
		return "退款结果明细"

	case GOODS_TYPE:
		return "商品类型"
	default:
		return ""
	}
}

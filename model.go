package alipay

/*
tip:
response:all response member is  string,not float64
request:request is ok
*/

type ReqCustomerDto struct {
	PriKey string `json:"pri_key,omitempty" query:"pri_key,omitempty"`
	PubKey string `json:"pub_key,omitempty" query:"pub_key,omitempty"`
}

type ReqBaseDto struct {
	AppId    string `json:"app_id,omitempty" query:"app_id,omitempty"`       //支付宝分配给开发者的应用ID
	Method   string `json:"method,omitempty" query:"method,omitempty"`       //接口名称
	Format   string `json:"format,omitempty" query:"format,omitempty"`       //仅支持JSON
	Charset  string `json:"charset,omitempty" query:"charset,omitempty"`     //请求使用的编码格式，如utf-8,gbk,gb2312等
	SignType string `json:"sign_type,omitempty" query:"sign_type,omitempty"` //商户生成签名字符串所使用的签名算法类型，目前支持RSA2和RSA，推荐使用RSA2

	Sign         string `json:"sign,omitempty" query:"sign,omitempty"`                     //商户请求参数的签名串，详见签名
	Timestamp    string `json:"timestamp,omitempty" query:"timestamp,omitempty"`           //发送请求的时间，格式"yyyy-MM-dd HH:mm:ss"
	Version      string `json:"version,omitempty" query:"version,omitempty"`               //调用的接口版本，固定为：1.0
	AppAuthToken string `json:"app_auth_token,omitempty" query:"app_auth_token,omitempty"` //支付宝服务器主动通知商户服务器里指定的页面http/https路径。
	NotifyUrl    string `json:"notify_url,omitempty" query:"notify_url,omitempty"`         //服务商账号时 配置

	BizContent string `json:"biz_content" query:"biz_content,omitempty"` //请求参数的集合，最大长度不限，除公共参数外所有请求参数都必须放在这个参数中传递，具体参照各产品快速接入文档
}

type ReqPayDto struct {
	*ReqBaseDto `json:"-"`

	OutTradeNo  string `json:"out_trade_no,omitempty" query:"out_trade_no,omitempty"` //商户订单号,64个字符以内、可包含字母、数字、下划线；需保证在商户端不重复
	Scene       string `json:"scene,omitempty" query:"scene,omitempty"`               //条码支付，取值：bar_code ,声波支付，取值：wave_code
	AuthCode    string `json:"auth_code,omitempty" query:"auth_code,omitempty"`       //支付授权码，25~30开头的长度为16~24位的数字，实际字符串长度以开发者获取的付款码长度为准
	ProductCode string `json:"product_code,omitempty" query:"product_code,omitempty"` //销售产品码
	Subject     string `json:"subject,omitempty" query:"subject,omitempty"`           //订单标题

	BuyerId        string  `json:"buyer_id,omitempty" query:"buyer_id,omitempty"`               //买家的支付宝用户id，如果为空，会从传入了码值信息中获取买家ID
	SellerId       string  `json:"seller_id,omitempty" query:"seller_id,omitempty"`             //如果该值为空，则默认为商户签约账号对应的支付宝用户ID
	TotalAmount    float64 `json:"total_amount,omitempty" query:"total_amount,omitempty"`       //float64 订单总金额，单位为元
	TransCurrency  string  `json:"trans_currency,omitempty" query:"trans_currency,omitempty"`   //标价币种, total_amount 对应的币种单位。支持英镑：GBP、港币：HKD、美元：USD
	SettleCurrency string  `json:"settle_currency,omitempty" query:"settle_currency,omitempty"` //商户指定的结算币种，支持英镑：GBP、港币：HKD、美元：USD

	DiscountableAmount float64        `json:"discountable_amount,omitempty" query:"discountable_amount,omitempty"` //float64
	Body               string         `json:"body,omitempty" query:"body,omitempty"`
	GoodsDetails       []*GoodsDetail `json:"goods_detail,omitempty" query:"goods_detail,omitempty"`
	OperatorId         string         `json:"operator_id,omitempty" query:"operator_id,omitempty"`
	StoreId            string         `json:"store_id,omitempty" query:"store_id,omitempty"`

	TerminalId      string         `json:"terminal_id,omitempty" query:"terminal_id,omitempty"`
	ExtendParams    *ExtendParams  `json:"extend_params,omitempty" query:"extend_params,omitempty"`
	TimeoutExpress  string         `json:"timeout_express,omitempty" query:"timeout_express,omitempty"`
	AuthConfirmMode string         `json:"auth_confirm_mode,omitempty" query:"auth_confirm_mode,omitempty"` //预授权确认模式，授权转交易请求中传入，适用于预授权转交易业务使用，目前只支持PRE_AUTH(预授权产品码) COMPLETE：转交易支付完成结束预授权，解冻剩余金额; NOT_COMPLETE：转交易支付完成不结束预授权，不解冻剩余金额
	TerminalParams  string         `json:"terminal_params,omitempty" query:"terminal_params,omitempty"`     //商户传入终端设备相关信息，具体值要和支付宝约定
	PromoParams     *PromoParamDto `json:"promo_params,omitempty" query:"promo_params,omitempty"`           //优惠明细参数，通过此属性补充营销参数

}

type ReqQueryDto struct {
	*ReqBaseDto `json:"-"`

	OutTradeNo string `json:"out_trade_no,omitempty" query:"out_trade_no,omitempty"`
	TradeNo    string `json:"trade_no,omitempty" query:"trade_no,omitempty"`
	OrgPid     string `json:"org_pid,omitempty" query:"org_pid,omitempty"`
}

type ReqRefundDto struct {
	*ReqBaseDto `json:"-"`

	OutTradeNo     string  `json:"out_trade_no,omitempty" query:"out_trade_no,omitempty"`
	TradeNo        string  `json:"trade_no,omitempty" query:"trade_no,omitempty"`
	RefundAmount   float64 `json:"refund_amount,omitempty" query:"refund_amount,omitempty"` //float64
	RefundCurrency string  `json:"refund_currency,omitempty" query:"refund_currency,omitempty"`
	RefundReason   string  `json:"refund_reason,omitempty" query:"refund_reason,omitempty"`
	OutRequestNo   string  `json:"out_request_no,omitempty" query:"out_request_no,omitempty"`

	OperatorId              string                          `json:"opreator_id,omitempty" query:"opreator_id,omitempty"`
	StoreId                 string                          `json:"store_id,omitempty" query:"store_id,omitempty"`
	TerminalId              string                          `json:"terminal_id,omitempty" query:"terminal_id,omitempty"`
	GoodsDetails            []*GoodsDetail                  `json:"goods_detail,omitempty" query:"goods_detail,omitempty"`
	RefundRoyaltyParameters []*OpenApiRoyaltyDetailInfoPojo `json:"refund_royalty_parameters,omitempty" query:"refund_royalty_parameters,omitempty"`
	OrgPid                  string                          `json:"org_pid,omitempty" query:"org_pid,omitempty"`
}

type OpenApiRoyaltyDetailInfoPojo struct {
	RoyaltyType  string `json:"royalty_type,omitempty" query:"royalty_type,omitempty"`
	TransOut     string `json:"trans_out,omitempty" query:"trans_out,omitempty"`
	TransOutType string `json:"trans_out_type,omitempty" query:"trans_out_type,omitempty"`
	TransInType  string `json:"trans_in_type,omitempty" query:"trans_in_type,omitempty"`
	TransIn      string `json:"trans_in,omitempty" query:"trans_in,omitempty"`

	Amount           float64 `json:"amount,omitempty" query:"amount,omitempty"`
	AmountPercentage int64   `json:"amount_percentage,omitempty" query:"amount_percentage,omitempty"`
	Desc             string  `json:"desc,omitempty" query:"desc,omitempty"`
}

type ReqReverseDto struct {
	*ReqBaseDto `json:"-"`

	OutTradeNo string `json:"out_trade_no,omitempty" query:"out_trade_no,omitempty"`
	TradeNo    string `json:"trade_no,omitempty" query:"trade_no,omitempty"`
}

type ReqPrepayDto struct {
	*ReqBaseDto `json:"-"`

	OutTradeNo         string  `json:"out_trade_no,omitempty" query:"out_trade_no,omitempty"`
	SellerId           string  `json:"seller_id,omitempty" query:"seller_id,omitempty"`
	TotalAmount        float64 `json:"total_amount,omitempty" query:"total_amount,omitempty"`               //float64
	DiscountableAmount float64 `json:"discountable_amount,omitempty" query:"discountable_amount,omitempty"` //float64
	Subject            string  `json:"subject,omitempty" query:"subject,omitempty"`

	GoodsDetail        []*GoodsDetail `json:"goods_detail,omitempty" query:"goods_detail,omitempty"`
	Body               string         `json:"body,omitempty" query:"body,omitempty"`
	OperatorId         string         `json:"operator_id,omitempty" query:"operator_id,omitempty"`
	StoreId            string         `json:"store_id,omitempty" query:"store_id,omitempty"`
	DisablePayChannels string         `json:"disable_pay_channels,omitempty" query:"disable_pay_channels,omitempty"`

	EnablePayChannels    string        `json:"enable_pay_channels,omitempty" query:"enable_pay_channels,omitempty"`
	TerminalId           string        `json:"terminal_id,omitempty" query:"terminal_id,omitempty"`
	ExtendParams         *ExtendParams `json:"extend_params,omitempty" query:"extend_params,omitempty"`
	TimeoutExpress       string        `json:"timeout_express,omitempty" query:"timeout_express,omitempty"`
	SettleInfo           *SettleInfo   `json:"settle_info,omitempty" query:"settle_info,omitempty"`
	BusinessParams       string        `json:"business_params,omitempty" query:"business_params,omitempty"`
	QrCodeTimeoutExpress string        `json:"qr_code_timeout_express,omitempty" query:"qr_code_timeout_express,omitempty"`
}

type ReqNotifyDto struct {
	NotifyTime string `json:"notify_time" query:"notify_time,omitempty"`
	NotifyType string `json:"notify_type" query:"notify_type,omitempty"`
	NotifyId   string `json:"notify_id" query:"notify_id,omitempty"`
	SignType   string `json:"sign_type" query:"sign_type,omitempty"`
	Sign       string `json:"sign" query:"sign,omitempty"`

	TradeNo    string `json:"trade_no" query:"trade_no,omitempty"`
	AppId      string `json:"app_id" query:"app_id,omitempty"`
	OutTradeNo string `json:"out_trade_no" query:"out_trade_no,omitempty"`
	OutBizNo   string `json:"out_biz_no" query:"out_biz_no,omitempty"`
	BuyerId    string `json:"buyer_id" query:"buyer_id,omitempty"`

	BuyerLogonId string  `json:"buyer_logon_id" query:"buyer_logon_id,omitempty"`
	SellerId     string  `json:"seller_id" query:"seller_id,omitempty"`
	SellerEmail  string  `json:"seller_email" query:"seller_email,omitempty"`
	TradeStatus  string  `json:"trade_status" query:"trade_status,omitempty"`
	TotalAmount  float64 `json:"total_amount" query:"total_amount,omitempty"`

	ReceiptAmount  float64 `json:"receipt_amount" query:"receipt_amount,omitempty"`
	InvoiceAmount  float64 `json:"invoice_amount" query:"invoice_amount,omitempty"`
	BuyerPayAmount float64 `json:"buyer_pay_amount" query:"buyer_pay_amount,omitempty"`
	PointAmount    float64 `json:"point_amount" query:"point_amount,omitempty"`
	RefundFee      float64 `json:"refund_fee" query:"refund_fee,omitempty"`

	SendBackFee float64 `json:"send_back_fee" query:"send_back_fee,omitempty"`
	Subject     string  `json:"subject" query:"subject,omitempty"`
	Body        string  `json:"body" query:"body,omitempty"`
	GmtCreate   string  `json:"gmt_create" query:"gmt_create,omitempty"`
	GmtPayment  string  `json:"gmt_payment" query:"gmt_payment,omitempty"`

	GmtRefund    string `json:"gmt_refund" query:"gmt_refund,omitempty"`
	GmtClose     string `json:"gmt_close" query:"gmt_close,omitempty"`
	FundBillList string `json:"fund_bill_list" query:"fund_bill_list,omitempty"`
}

//resp
type RespBaseDto struct {
	Code    string `json:"code,omitempty" query:"code,omitempty"`
	Msg     string `json:"msg,omitempty" query:"msg,omitempty"`
	SubCode string `json:"sub_code,omitempty" query:"sub_code,omitempty"`
	SubMsg  string `json:"sub_msg,omitempty" query:"sub_msg,omitempty"`
	Sign    string `json:"sign,omitempty" query:"sign,omitempty"`
}

type RespPayDto struct {
	*RespBaseDto

	TradeNo         string `json:"trade_no,omitempty" query:"trade_no,omitempty"`
	OutTradeNo      string `json:"out_trade_no,omitempty" query:"out_trade_no,omitempty"`
	BuyerLogonId    string `json:"buyer_logon_id,omitempty" query:"buyer_logon_id,omitempty"`
	PayAmount       string `json:"pay_amount,omitempty" query:"pay_amount,omitempty"`
	SettleTransRate string `json:"settle_trans_rate,omitempty" query:"settle_trans_rate,omitempty"`

	TransPayRate   string `json:"trans_pay_rate,omitempty" query:"trans_pay_rate,omitempty"`
	TotalAmount    string `json:"total_amount,omitempty" query:"total_amount,omitempty"` //float64
	TransCurrency  string `json:"trans_currency,omitempty" query:"trans_currency,omitempty"`
	SettleCurrency string `json:"settle_currency,omitempty" query:"settle_currency,omitempty"`
	SettleAmount   string `json:"settle_amount,omitempty" query:"settle_amount,omitempty"`

	PayCurrency    string `json:"pay_currency,omitempty" query:"pay_currency,omitempty"`
	ReceiptAmount  string `json:"receipt_amount,omitempty" query:"receipt_amount,omitempty"`
	BuyerPayAmount string `json:"buyer_pay_amount,omitempty" query:"buyer_pay_amount,omitempty"` //float64
	PointAmount    string `json:"point_amount,omitempty" query:"point_amount,omitempty"`         //float64
	InvoiceAmount  string `json:"invoice_amount,omitempty" query:"invoice_amount,omitempty"`     //float64

	GmtPayment   string      `json:"gmt_payment,omitempty" query:"gmt_payment,omitempty"` //time.Time
	FundBillList []*FundBill `json:"fund_bill_list,omitempty" query:"fund_bill_list,omitempty"`
	CardBalance  string      `json:"card_balance,omitempty" query:"card_balance,omitempty"` //float64
	StoreName    string      `json:"store_name,omitempty" query:"store_name,omitempty"`
	BuyerUserId  string      `json:"buyer_user_id,omitempty" query:"buyer_user_id,omitempty"`

	DiscountGoodsDetail string           `json:"discount_goods_detail,omitempty" query:"discount_goods_detail,omitempty"`
	VoucherDetailList   []*VoucherDetail `json:"voucher_detail_list,omitempty" query:"voucher_detail_list,omitempty"`
	AuthTradePayMode    string           `json:"auth_trade_pay_mode,omitempty" query:"auth_trade_pay_mode,omitempty"`
	BusinessParam       string           `json:"business_param,omitempty" query:"business_param,omitempty"`
	BuyerUserType       string           `json:"buyer_user_type,omitempty" query:"buyer_user_type,omitempty"`

	MdiscountAmount string `json:"mdiscount_amount,omitempty" query:"mdiscount_amount,omitempty"`
	DiscountAmount  string `json:"discount_amount,omitempty" query:"discount_amount,omitempty"`
}
type RespQueryDto struct {
	*RespBaseDto

	TradeNo      string `json:"trade_no,omitempty" query:"trade_no,omitempty"`
	OutTradeNo   string `json:"out_trade_no,omitempty" query:"out_trade_no,omitempty"`
	BuyerLogonId string `json:"buyer_logon_id,omitempty" query:"buyer_logon_id,omitempty"`
	TradeStatus  string `json:"trade_status,omitempty" query:"trade_status,omitempty"`
	TotalAmount  string `json:"total_amount,omitempty" query:"total_amount,omitempty"` //float64

	TransCurrency  string `json:"trans_currency,omitempty" query:"trans_currency,omitempty"`   //
	SettleCurrency string `json:"settle_currency,omitempty" query:"settle_currency,omitempty"` //
	SettleAmount   string `json:"settle_amount,omitempty" query:"settle_amount,omitempty"`     //
	PayCurrency    string `json:"pay_currency,omitempty" query:"pay_currency,omitempty"`       //
	PayAmount      string `json:"pay_amount,omitempty" query:"pay_amount,omitempty"`           //

	SettleTransRate string `json:"settle_trans_rate,omitempty" query:"settle_trans_rate,omitempty"` //
	TransPayRate    string `json:"trans_pay_rate,omitempty" query:"trans_pay_rate,omitempty"`       //
	BuyerPayAmount  string `json:"buyer_pay_amount,omitempty" query:"buyer_pay_amount,omitempty"`   //float64
	PointAmount     string `json:"point_amount,omitempty" query:"point_amount,omitempty"`           //float64
	InvoiceAmount   string `json:"invoice_amount,omitempty" query:"invoice_amount,omitempty"`       //float64

	SendPayDate   string      `json:"send_pay_date,omitempty" query:"send_pay_date,omitempty"` //time.Time
	ReceiptAmount string      `json:"receipt_amount,omitempty" query:"receipt_amount,omitempty"`
	StoreId       string      `json:"store_id,omitempty" query:"store_id,omitempty"`
	TerminalId    string      `json:"terminal_id,omitempty" query:"terminal_id,omitempty"`
	FundBillList  []*FundBill `json:"fund_bill_list,omitempty" query:"fund_bill_list,omitempty"`

	StoreName        string `json:"store_name,omitempty" query:"store_name,omitempty"`
	BuyerUserId      string `json:"buyer_user_id,omitempty" query:"buyer_user_id,omitempty"`
	AuthTradePayMode string `json:"auth_trade_pay_mode,omitempty" query:"auth_trade_pay_mode,omitempty"`
	BuyerUserType    string `json:"buyer_user_type,omitempty" query:"buyer_user_type,omitempty"`
	MdiscountAmount  string `json:"mdiscount_amount,omitempty" query:"mdiscount_amount,omitempty"`

	DiscountAmount string `json:"discount_amount,omitempty" query:"discount_amount,omitempty"`
}
type RespRefundDto struct {
	*RespBaseDto

	TradeNo      string `json:"trade_no,omitempty" query:"trade_no,omitempty"`
	OutTradeNo   string `json:"out_trade_no,omitempty" query:"out_trade_no,omitempty"`
	BuyerLogonId string `json:"buyer_logon_id,omitempty" query:"buyer_logon_id,omitempty"`
	FundChange   string `json:"fund_change,omitempty" query:"fund_change,omitempty"`
	RefundFee    string `json:"refund_fee,omitempty" query:"refund_fee,omitempty"` //float64

	GmtRefundPay         string              `json:"gmt_refund_pay,omitempty" query:"gmt_refund_pay,omitempty"` //time.Time
	RefundDetailItemList []*RefundDetailItem `json:"refund_detail_item_list,omitempty" query:"refund_detail_item_list,omitempty"`
	StoreName            string              `json:"store_name,omitempty" query:"store_name,omitempty"`
	BuyerUserId          string              `json:"buyer_user_id,omitempty" query:"buyer_user_id,omitempty"`

	SendBackFee string `json:"send_back_fee,omitempty" query:"send_back_fee,omitempty"` //float64
}
type RespReverseDto struct {
	*RespBaseDto

	TradeNo    string `json:"trade_no,omitempty" query:"trade_no,omitempty"`
	OutTradeNo string `json:"out_trade_no,omitempty" query:"out_trade_no,omitempty"`
	RetryFlag  string `json:"retry_flag,omitempty" query:"retry_flag,omitempty"`
	Action     string `json:"action,omitempty" query:"action,omitempty"`
}

type RespPrepayDto struct {
	*RespBaseDto

	OutTradeNo string `json:"out_trade_no,omitempty" query:"out_trade_no,omitempty"`
	QrCode     string `json:"qr_code,omitempty" query:"qr_code,omitempty"`
}

type RefundDetailItem struct {
	FundChannel string `json:"fund_channel,omitempty" query:"fund_channel,omitempty"`
	Amount      string `json:"amount,omitempty" query:"amount,omitempty"`           //float64
	RealAmount  string `json:"real_amount,omitempty" query:"real_amount,omitempty"` //float64
	FundType    string `json:"fund_type,omitempty" query:"fund_type,omitempty"`
}

type FundBill struct {
	FundChannel string `json:"fund_channel,omitempty" query:"fund_channel,omitempty"`
	Amount      string `json:"amount,omitempty" query:"amount,omitempty"`           //float64
	RealAmount  string `json:"real_amount,omitempty" query:"real_amount,omitempty"` //float64
}

type VoucherDetail struct {
	Id                         string `json:"id,omitempty" query:"id,omitempty"`
	Name                       string `json:"name,omitempty" query:"name,omitempty"`
	Type                       string `json:"type,omitempty" query:"type,omitempty"`
	Amount                     string `json:"amount,omitempty" query:"amount,omitempty"`                           //float64
	MerchantContribute         string `json:"merchant_contribute,omitempty" query:"merchant_contribute,omitempty"` //float64
	Othercontribute            string `json:"other_contribute,omitempty" query:"other_contribute,omitempty"`       //float64
	Memo                       string `json:"memo,omitempty" query:"memo,omitempty"`
	TemplateId                 string `json:"template_id,omitempty" query:"template_id,omitempty"`
	PurchaseBuyerContribute    string `json:"purchase_buyer_contribute,omitempty" query:"purchase_buyer_contribute,omitempty"` //float64
	PurchaseMerchantContribute string `json:"purchase_buyer_contribute,omitempty" query:"purchase_buyer_contribute,omitempty"` //float64
	PurchaseAntContribute      string `json:"purchase_ant_contribute,omitempty" query:"purchase_ant_contribute,omitempty"`     //float64
}

//req
type GoodsDetail struct {
	GoodsId       string  `json:"goods_id,omitempty" query:"goods_id,omitempty"`
	GoodsName     string  `json:"goods_name,omitempty" query:"goods_name,omitempty"`
	Quantity      int64   `json:"quantity,omitempty" query:"quantity,omitempty"`
	Price         float64 `json:"price,omitempty" query:"price,omitempty"`
	GoodsCategory string  `json:"goods_category,omitempty" query:"goods_category,omitempty"`
	Body          string  `json:"body,omitempty" query:"body,omitempty"`
	ShowUrl       string  `json:"show_url,omitempty" query:"show_url,omitempty"`
}

type ExtendParams struct {
	SysServiceProviderId string `json:"sys_service_provider_id,omitempty" query:"sys_service_provider_id,omitempty"`
}
type PromoParamDto struct {
	ActualOrderTime string `json:"actual_order_time,omitempty" query:"actual_order_time,omitempty"` //存在延迟扣款这一类的场景，用这个时间表明用户发生交易的时间，比如说，在公交地铁场景，用户刷码出站的时间，和商户上送交易的时间是不一样的。
}

type SettleInfo struct {
	SettleDetailInfos []*SettleDetailInfo `json:"settle_detail_infos,omitempty" query:"settle_detail_infos,omitempty"`
	MerchantType      string              `json:"merchant_type,omitempty" query:"merchant_type,omitempty"`
}

type SettleDetailInfo struct {
	TransInType      string `json:"trans_in_type,omitempty" query:"trans_in_type,omitempty"`
	TransIn          string `json:"trans_in,omitempty" query:"trans_in,omitempty"`
	SummaryDimension string `json:"summary_dimension,omitempty" query:"summary_dimension,omitempty"`
	SettleEntityId   string `json:"settle_entity_id,omitempty" query:"settle_entity_id,omitempty"`
	SettleEntityType string `json:"settle_entity_type,omitempty" query:"settle_entity_type,omitempty"`

	Amount float64 `json:"amount,omitempty" query:"amount,omitempty"`
}

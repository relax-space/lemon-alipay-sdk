package alipay

/*
tip:
response:all response member is  string,not float64
request:request is ok
*/

type ReqCustomerDto struct {
	PriKey string `json:"pri_key,omitempty" query:"pri_key"`
	PubKey string `json:"pub_key,omitempty" query:"pub_key"`
}

type ReqBaseDto struct {
	AppId    string `json:"app_id,omitempty" query:"app_id"`       //支付宝分配给开发者的应用ID
	Method   string `json:"method,omitempty" query:"method"`       //接口名称
	Format   string `json:"format,omitempty" query:"format"`       //仅支持JSON
	Charset  string `json:"charset,omitempty" query:"charset"`     //请求使用的编码格式，如utf-8,gbk,gb2312等
	SignType string `json:"sign_type,omitempty" query:"sign_type"` //商户生成签名字符串所使用的签名算法类型，目前支持RSA2和RSA，推荐使用RSA2

	Sign         string `json:"sign,omitempty" query:"sign"`                     //商户请求参数的签名串，详见签名
	Timestamp    string `json:"timestamp,omitempty" query:"timestamp"`           //发送请求的时间，格式"yyyy-MM-dd HH:mm:ss"
	Version      string `json:"version,omitempty" query:"version"`               //调用的接口版本，固定为：1.0
	AppAuthToken string `json:"app_auth_token,omitempty" query:"app_auth_token"` //支付宝服务器主动通知商户服务器里指定的页面http/https路径。
	NotifyUrl    string `json:"notify_url,omitempty" query:"notify_url"`         //服务商账号时 配置

	BizContent string `json:"biz_content" query:"biz_content"` //请求参数的集合，最大长度不限，除公共参数外所有请求参数都必须放在这个参数中传递，具体参照各产品快速接入文档
}

type ReqPayDto struct {
	*ReqBaseDto `json:"-"`

	OutTradeNo  string `json:"out_trade_no,omitempty" query:"out_trade_no"` //商户订单号,64个字符以内、可包含字母、数字、下划线；需保证在商户端不重复
	Scene       string `json:"scene,omitempty" query:"scene"`               //条码支付，取值：bar_code ,声波支付，取值：wave_code
	AuthCode    string `json:"auth_code,omitempty" query:"auth_code"`       //支付授权码，25~30开头的长度为16~24位的数字，实际字符串长度以开发者获取的付款码长度为准
	ProductCode string `json:"product_code,omitempty" query:"product_code"` //销售产品码
	Subject     string `json:"subject,omitempty" query:"subject"`           //订单标题

	BuyerId        string  `json:"buyer_id,omitempty" query:"buyer_id"`               //买家的支付宝用户id，如果为空，会从传入了码值信息中获取买家ID
	SellerId       string  `json:"seller_id,omitempty" query:"seller_id"`             //如果该值为空，则默认为商户签约账号对应的支付宝用户ID
	TotalAmount    float64 `json:"total_amount,omitempty" query:"total_amount"`       //float64 订单总金额，单位为元
	TransCurrency  string  `json:"trans_currency,omitempty" query:"trans_currency"`   //标价币种, total_amount 对应的币种单位。支持英镑：GBP、港币：HKD、美元：USD
	SettleCurrency string  `json:"settle_currency,omitempty" query:"settle_currency"` //商户指定的结算币种，支持英镑：GBP、港币：HKD、美元：USD

	DiscountableAmount float64        `json:"discountable_amount,omitempty" query:"discountable_amount"` //float64
	Body               string         `json:"body,omitempty" query:"body"`
	GoodsDetails       []*GoodsDetail `json:"goods_detail,omitempty" query:"goods_detail"`
	OperatorId         string         `json:"operator_id,omitempty" query:"operator_id"`
	StoreId            string         `json:"store_id,omitempty" query:"store_id"`

	TerminalId      string         `json:"terminal_id,omitempty" query:"terminal_id"`
	ExtendParams    *ExtendParams  `json:"extend_params,omitempty" query:"extend_params"`
	TimeoutExpress  string         `json:"timeout_express,omitempty" query:"timeout_express"`
	AuthConfirmMode string         `json:"auth_confirm_mode,omitempty" query:"auth_confirm_mode"` //预授权确认模式，授权转交易请求中传入，适用于预授权转交易业务使用，目前只支持PRE_AUTH(预授权产品码) COMPLETE：转交易支付完成结束预授权，解冻剩余金额; NOT_COMPLETE：转交易支付完成不结束预授权，不解冻剩余金额
	TerminalParams  string         `json:"terminal_params,omitempty" query:"terminal_params"`     //商户传入终端设备相关信息，具体值要和支付宝约定
	PromoParams     *PromoParamDto `json:"promo_params,omitempty" query:"promo_params"`           //优惠明细参数，通过此属性补充营销参数

}

type ReqQueryDto struct {
	*ReqBaseDto `json:"-"`

	OutTradeNo string `json:"out_trade_no,omitempty" query:"out_trade_no"`
	TradeNo    string `json:"trade_no,omitempty" query:"trade_no"`
	OrgPid     string `json:"org_pid,omitempty" query:"org_pid"`
}

type ReqRefundDto struct {
	*ReqBaseDto `json:"-"`

	OutTradeNo     string  `json:"out_trade_no,omitempty" query:"out_trade_no"`
	TradeNo        string  `json:"trade_no,omitempty" query:"trade_no"`
	RefundAmount   float64 `json:"refund_amount,omitempty" query:"refund_amount"` //float64
	RefundCurrency string  `json:"refund_currency,omitempty" query:"refund_currency"`
	RefundReason   string  `json:"refund_reason,omitempty" query:"refund_reason"`
	OutRequestNo   string  `json:"out_request_no,omitempty" query:"out_request_no"`

	OperatorId              string                          `json:"opreator_id,omitempty" query:"opreator_id"`
	StoreId                 string                          `json:"store_id,omitempty" query:"store_id"`
	TerminalId              string                          `json:"terminal_id,omitempty" query:"terminal_id"`
	GoodsDetails            []*GoodsDetail                  `json:"goods_detail,omitempty" query:"goods_detail"`
	RefundRoyaltyParameters []*OpenApiRoyaltyDetailInfoPojo `json:"refund_royalty_parameters,omitempty" query:"refund_royalty_parameters"`
	OrgPid                  string                          `json:"org_pid,omitempty" query:"org_pid"`
}

type ReqReverseDto struct {
	*ReqBaseDto `json:"-"`

	OutTradeNo string `json:"out_trade_no,omitempty" query:"out_trade_no"`
	TradeNo    string `json:"trade_no,omitempty" query:"trade_no"`
}

type ReqPrepayDto struct {
	*ReqBaseDto `json:"-"`

	OutTradeNo         string  `json:"out_trade_no,omitempty" query:"out_trade_no"`
	SellerId           string  `json:"seller_id,omitempty" query:"seller_id"`
	TotalAmount        float64 `json:"total_amount,omitempty" query:"total_amount"`               //float64
	DiscountableAmount float64 `json:"discountable_amount,omitempty" query:"discountable_amount"` //float64
	Subject            string  `json:"subject,omitempty" query:"subject"`

	GoodsDetail        []*GoodsDetail `json:"goods_detail,omitempty" query:"goods_detail"`
	Body               string         `json:"body,omitempty" query:"body"`
	OperatorId         string         `json:"operator_id,omitempty" query:"operator_id"`
	StoreId            string         `json:"store_id,omitempty" query:"store_id"`
	DisablePayChannels string         `json:"disable_pay_channels,omitempty" query:"disable_pay_channels"`

	EnablePayChannels    string        `json:"enable_pay_channels,omitempty" query:"enable_pay_channels"`
	TerminalId           string        `json:"terminal_id,omitempty" query:"terminal_id"`
	ExtendParams         *ExtendParams `json:"extend_params,omitempty" query:"extend_params"`
	TimeoutExpress       string        `json:"timeout_express,omitempty" query:"timeout_express"`
	SettleInfo           *SettleInfo   `json:"settle_info,omitempty" query:"settle_info"`
	BusinessParams       string        `json:"business_params,omitempty" query:"business_params"`
	QrCodeTimeoutExpress string        `json:"qr_code_timeout_express,omitempty" query:"qr_code_timeout_express"`
}

type ReqNotifyDto struct {
	NotifyTime string `json:"notify_time" query:"notify_time"`
	NotifyType string `json:"notify_type" query:"notify_type"`
	NotifyId   string `json:"notify_id" query:"notify_id"`
	SignType   string `json:"sign_type" query:"sign_type"`
	Sign       string `json:"sign" query:"sign"`

	TradeNo    string `json:"trade_no" query:"trade_no"`
	AppId      string `json:"app_id" query:"app_id"`
	OutTradeNo string `json:"out_trade_no" query:"out_trade_no"`
	OutBizNo   string `json:"out_biz_no" query:"out_biz_no"`
	BuyerId    string `json:"buyer_id" query:"buyer_id"`

	BuyerLogonId string  `json:"buyer_logon_id" query:"buyer_logon_id"`
	SellerId     string  `json:"seller_id" query:"seller_id"`
	SellerEmail  string  `json:"seller_email" query:"seller_email"`
	TradeStatus  string  `json:"trade_status" query:"trade_status"`
	TotalAmount  float64 `json:"total_amount" query:"total_amount"`

	ReceiptAmount  float64 `json:"receipt_amount" query:"receipt_amount"`
	InvoiceAmount  float64 `json:"invoice_amount" query:"invoice_amount"`
	BuyerPayAmount float64 `json:"buyer_pay_amount" query:"buyer_pay_amount"`
	PointAmount    float64 `json:"point_amount" query:"point_amount"`
	RefundFee      float64 `json:"refund_fee" query:"refund_fee"`

	SendBackFee float64 `json:"send_back_fee" query:"send_back_fee"`
	Subject     string  `json:"subject" query:"subject"`
	Body        string  `json:"body" query:"body"`
	GmtCreate   string  `json:"gmt_create" query:"gmt_create"`
	GmtPayment  string  `json:"gmt_payment" query:"gmt_payment"`

	GmtRefund    string `json:"gmt_refund" query:"gmt_refund"`
	GmtClose     string `json:"gmt_close" query:"gmt_close"`
	FundBillList string `json:"fund_bill_list" query:"fund_bill_list"`
}

type ReqBillDto struct {
	*ReqBaseDto `json:"-"`
	BillType    string `json:"bill_type,omitempty" query:"bill_type"`
	BillDate    string `json:"bill_date,omitempty" query:"bill_date"`
}

//resp
type RespBaseDto struct {
	Code    string `json:"code,omitempty" query:"code"`
	Msg     string `json:"msg,omitempty" query:"msg"`
	SubCode string `json:"sub_code,omitempty" query:"sub_code"`
	SubMsg  string `json:"sub_msg,omitempty" query:"sub_msg"`
	Sign    string `json:"sign,omitempty" query:"sign"`
}

type RespPayDto struct {
	*RespBaseDto

	TradeNo         string `json:"trade_no,omitempty" query:"trade_no"`
	OutTradeNo      string `json:"out_trade_no,omitempty" query:"out_trade_no"`
	BuyerLogonId    string `json:"buyer_logon_id,omitempty" query:"buyer_logon_id"`
	PayAmount       string `json:"pay_amount,omitempty" query:"pay_amount"`
	SettleTransRate string `json:"settle_trans_rate,omitempty" query:"settle_trans_rate"`

	TransPayRate   string `json:"trans_pay_rate,omitempty" query:"trans_pay_rate"`
	TotalAmount    string `json:"total_amount,omitempty" query:"total_amount"` //float64
	TransCurrency  string `json:"trans_currency,omitempty" query:"trans_currency"`
	SettleCurrency string `json:"settle_currency,omitempty" query:"settle_currency"`
	SettleAmount   string `json:"settle_amount,omitempty" query:"settle_amount"`

	PayCurrency    string `json:"pay_currency,omitempty" query:"pay_currency"`
	ReceiptAmount  string `json:"receipt_amount,omitempty" query:"receipt_amount"`
	BuyerPayAmount string `json:"buyer_pay_amount,omitempty" query:"buyer_pay_amount"` //float64
	PointAmount    string `json:"point_amount,omitempty" query:"point_amount"`         //float64
	InvoiceAmount  string `json:"invoice_amount,omitempty" query:"invoice_amount"`     //float64

	GmtPayment   string      `json:"gmt_payment,omitempty" query:"gmt_payment"` //time.Time
	FundBillList []*FundBill `json:"fund_bill_list,omitempty" query:"fund_bill_list"`
	CardBalance  string      `json:"card_balance,omitempty" query:"card_balance"` //float64
	StoreName    string      `json:"store_name,omitempty" query:"store_name"`
	BuyerUserId  string      `json:"buyer_user_id,omitempty" query:"buyer_user_id"`

	DiscountGoodsDetail string           `json:"discount_goods_detail,omitempty" query:"discount_goods_detail"`
	VoucherDetailList   []*VoucherDetail `json:"voucher_detail_list,omitempty" query:"voucher_detail_list"`
	AuthTradePayMode    string           `json:"auth_trade_pay_mode,omitempty" query:"auth_trade_pay_mode"`
	BusinessParam       string           `json:"business_param,omitempty" query:"business_param"`
	BuyerUserType       string           `json:"buyer_user_type,omitempty" query:"buyer_user_type"`

	MdiscountAmount string `json:"mdiscount_amount,omitempty" query:"mdiscount_amount"`
	DiscountAmount  string `json:"discount_amount,omitempty" query:"discount_amount"`
}
type RespQueryDto struct {
	*RespBaseDto

	TradeNo      string `json:"trade_no,omitempty" query:"trade_no"`
	OutTradeNo   string `json:"out_trade_no,omitempty" query:"out_trade_no"`
	BuyerLogonId string `json:"buyer_logon_id,omitempty" query:"buyer_logon_id"`
	TradeStatus  string `json:"trade_status,omitempty" query:"trade_status"`
	TotalAmount  string `json:"total_amount,omitempty" query:"total_amount"` //float64

	TransCurrency  string `json:"trans_currency,omitempty" query:"trans_currency"`   //
	SettleCurrency string `json:"settle_currency,omitempty" query:"settle_currency"` //
	SettleAmount   string `json:"settle_amount,omitempty" query:"settle_amount"`     //
	PayCurrency    string `json:"pay_currency,omitempty" query:"pay_currency"`       //
	PayAmount      string `json:"pay_amount,omitempty" query:"pay_amount"`           //

	SettleTransRate string `json:"settle_trans_rate,omitempty" query:"settle_trans_rate"` //
	TransPayRate    string `json:"trans_pay_rate,omitempty" query:"trans_pay_rate"`       //
	BuyerPayAmount  string `json:"buyer_pay_amount,omitempty" query:"buyer_pay_amount"`   //float64
	PointAmount     string `json:"point_amount,omitempty" query:"point_amount"`           //float64
	InvoiceAmount   string `json:"invoice_amount,omitempty" query:"invoice_amount"`       //float64

	SendPayDate   string      `json:"send_pay_date,omitempty" query:"send_pay_date"` //time.Time
	ReceiptAmount string      `json:"receipt_amount,omitempty" query:"receipt_amount"`
	StoreId       string      `json:"store_id,omitempty" query:"store_id"`
	TerminalId    string      `json:"terminal_id,omitempty" query:"terminal_id"`
	FundBillList  []*FundBill `json:"fund_bill_list,omitempty" query:"fund_bill_list"`

	StoreName        string `json:"store_name,omitempty" query:"store_name"`
	BuyerUserId      string `json:"buyer_user_id,omitempty" query:"buyer_user_id"`
	AuthTradePayMode string `json:"auth_trade_pay_mode,omitempty" query:"auth_trade_pay_mode"`
	BuyerUserType    string `json:"buyer_user_type,omitempty" query:"buyer_user_type"`
	MdiscountAmount  string `json:"mdiscount_amount,omitempty" query:"mdiscount_amount"`

	DiscountAmount string `json:"discount_amount,omitempty" query:"discount_amount"`
}
type RespRefundDto struct {
	*RespBaseDto

	TradeNo      string `json:"trade_no,omitempty" query:"trade_no"`
	OutTradeNo   string `json:"out_trade_no,omitempty" query:"out_trade_no"`
	BuyerLogonId string `json:"buyer_logon_id,omitempty" query:"buyer_logon_id"`
	FundChange   string `json:"fund_change,omitempty" query:"fund_change"`
	RefundFee    string `json:"refund_fee,omitempty" query:"refund_fee"` //float64

	GmtRefundPay         string              `json:"gmt_refund_pay,omitempty" query:"gmt_refund_pay"` //time.Time
	RefundDetailItemList []*RefundDetailItem `json:"refund_detail_item_list,omitempty" query:"refund_detail_item_list"`
	StoreName            string              `json:"store_name,omitempty" query:"store_name"`
	BuyerUserId          string              `json:"buyer_user_id,omitempty" query:"buyer_user_id"`

	SendBackFee string `json:"send_back_fee,omitempty" query:"send_back_fee"` //float64
}
type RespReverseDto struct {
	*RespBaseDto

	TradeNo    string `json:"trade_no,omitempty" query:"trade_no"`
	OutTradeNo string `json:"out_trade_no,omitempty" query:"out_trade_no"`
	RetryFlag  string `json:"retry_flag,omitempty" query:"retry_flag"`
	Action     string `json:"action,omitempty" query:"action"`
}

type RespPrepayDto struct {
	*RespBaseDto

	OutTradeNo string `json:"out_trade_no,omitempty" query:"out_trade_no"`
	QrCode     string `json:"qr_code,omitempty" query:"qr_code"`
}

type RespBillDto struct {
	*RespBaseDto
	BillDownloadUrl string `json:"bill_download_url,omitempty" query:"bill_download_url"`
}

type RefundDetailItem struct {
	FundChannel string `json:"fund_channel,omitempty" query:"fund_channel"`
	Amount      string `json:"amount,omitempty" query:"amount"`           //float64
	RealAmount  string `json:"real_amount,omitempty" query:"real_amount"` //float64
	FundType    string `json:"fund_type,omitempty" query:"fund_type"`
}

type FundBill struct {
	FundChannel string `json:"fund_channel,omitempty" query:"fund_channel"`
	Amount      string `json:"amount,omitempty" query:"amount"`           //float64
	RealAmount  string `json:"real_amount,omitempty" query:"real_amount"` //float64
}

type VoucherDetail struct {
	Id                         string `json:"id,omitempty" query:"id"`
	Name                       string `json:"name,omitempty" query:"name"`
	Type                       string `json:"type,omitempty" query:"type"`
	Amount                     string `json:"amount,omitempty" query:"amount"`                           //float64
	MerchantContribute         string `json:"merchant_contribute,omitempty" query:"merchant_contribute"` //float64
	Othercontribute            string `json:"other_contribute,omitempty" query:"other_contribute"`       //float64
	Memo                       string `json:"memo,omitempty" query:"memo"`
	TemplateId                 string `json:"template_id,omitempty" query:"template_id"`
	PurchaseBuyerContribute    string `json:"purchase_buyer_contribute,omitempty" query:"purchase_buyer_contribute"` //float64
	PurchaseMerchantContribute string `json:"purchase_buyer_contribute,omitempty" query:"purchase_buyer_contribute"` //float64
	PurchaseAntContribute      string `json:"purchase_ant_contribute,omitempty" query:"purchase_ant_contribute"`     //float64
}

//req
type GoodsDetail struct {
	GoodsId       string  `json:"goods_id,omitempty" query:"goods_id"`
	GoodsName     string  `json:"goods_name,omitempty" query:"goods_name"`
	Quantity      int64   `json:"quantity,omitempty" query:"quantity"`
	Price         float64 `json:"price,omitempty" query:"price"`
	GoodsCategory string  `json:"goods_category,omitempty" query:"goods_category"`
	Body          string  `json:"body,omitempty" query:"body"`
	ShowUrl       string  `json:"show_url,omitempty" query:"show_url"`
}

type ExtendParams struct {
	SysServiceProviderId string `json:"sys_service_provider_id,omitempty" query:"sys_service_provider_id"`
}
type PromoParamDto struct {
	ActualOrderTime string `json:"actual_order_time,omitempty" query:"actual_order_time"` //存在延迟扣款这一类的场景，用这个时间表明用户发生交易的时间，比如说，在公交地铁场景，用户刷码出站的时间，和商户上送交易的时间是不一样的。
}

type SettleInfo struct {
	SettleDetailInfos []*SettleDetailInfo `json:"settle_detail_infos,omitempty" query:"settle_detail_infos"`
	MerchantType      string              `json:"merchant_type,omitempty" query:"merchant_type"`
}

type SettleDetailInfo struct {
	TransInType      string `json:"trans_in_type,omitempty" query:"trans_in_type"`
	TransIn          string `json:"trans_in,omitempty" query:"trans_in"`
	SummaryDimension string `json:"summary_dimension,omitempty" query:"summary_dimension"`
	SettleEntityId   string `json:"settle_entity_id,omitempty" query:"settle_entity_id"`
	SettleEntityType string `json:"settle_entity_type,omitempty" query:"settle_entity_type"`

	Amount float64 `json:"amount,omitempty" query:"amount"`
}

type OpenApiRoyaltyDetailInfoPojo struct {
	RoyaltyType  string `json:"royalty_type,omitempty" query:"royalty_type"`
	TransOut     string `json:"trans_out,omitempty" query:"trans_out"`
	TransOutType string `json:"trans_out_type,omitempty" query:"trans_out_type"`
	TransInType  string `json:"trans_in_type,omitempty" query:"trans_in_type"`
	TransIn      string `json:"trans_in,omitempty" query:"trans_in"`

	Amount           float64 `json:"amount,omitempty" query:"amount"`
	AmountPercentage int64   `json:"amount_percentage,omitempty" query:"amount_percentage"`
	Desc             string  `json:"desc,omitempty" query:"desc"`
}

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
	AppId    string `json:"app_id,omitempty" query:"app_id,omitempty"`
	Method   string `json:"method,omitempty" query:"method,omitempty"`
	Format   string `json:"format,omitempty" query:"format,omitempty"`
	Charset  string `json:"charset,omitempty" query:"charset,omitempty"`
	SignType string `json:"sign_type,omitempty" query:"sign_type,omitempty"`

	Sign         string `json:"sign,omitempty" query:"sign,omitempty"`
	Timestamp    string `json:"timestamp,omitempty" query:"timestamp,omitempty"`
	Version      string `json:"version,omitempty" query:"version,omitempty"`
	AppAuthToken string `json:"app_auth_token,omitempty" query:"app_auth_token,omitempty"`
	NotifyUrl    string `json:"notify_url,omitempty" query:"notify_url,omitempty"`

	BizContent string `json:"biz_content" query:"biz_content,omitempty"`
}

type ReqPayDto struct {
	*ReqBaseDto `json:"-"`

	NotifyUrl string `json:"notify_url,omitempty" query:"notify_url,omitempty"`

	OutTradeNo  string `json:"out_trade_no,omitempty" query:"out_trade_no,omitempty"`
	Scene       string `json:"scene,omitempty" query:"scene,omitempty"`
	AuthCode    string `json:"auth_code,omitempty" query:"auth_code,omitempty"`
	ProductCode string `json:"product_code,omitempty" query:"product_code,omitempty"`
	Subject     string `json:"subject,omitempty" query:"subject,omitempty"`

	BuyerId            string  `json:"buyer_id,omitempty" query:"buyer_id,omitempty"`
	SellerId           string  `json:"seller_id,omitempty" query:"seller_id,omitempty"`
	TotalAmount        float64 `json:"total_amount,omitempty" query:"total_amount,omitempty"`               //float64
	DiscountableAmount float64 `json:"discountable_amount,omitempty" query:"discountable_amount,omitempty"` //float64
	Body               string  `json:"body,omitempty" query:"body,omitempty"`

	GoodsDetails *[]GoodsDetail `json:"goods_detail,omitempty" query:"goods_detail,omitempty"`
	OperatorId   string         `json:"operator_id,omitempty" query:"operator_id,omitempty"`
	StoreId      string         `json:"store_id,omitempty" query:"store_id,omitempty"`
	TerminalId   string         `json:"terminal_id,omitempty" query:"terminal_id,omitempty"`
	ExtendParams *ExtendParams  `json:"extend_params,omitempty" query:"extend_params,omitempty"`

	TimeoutExpress string `json:"timeout_express,omitempty" query:"timeout_express,omitempty"`
}

type ReqQueryDto struct {
	*ReqBaseDto `json:"-"`

	OutTradeNo string `json:"out_trade_no,omitempty" query:"out_trade_no,omitempty"`
	TradeNo    string `json:"trade_no,omitempty" query:"trade_no,omitempty"`
}

type ReqRefundDto struct {
	*ReqBaseDto `json:"-"`

	OutTradeNo   string  `json:"out_trade_no,omitempty" query:"out_trade_no,omitempty"`
	TradeNo      string  `json:"trade_no,omitempty" query:"trade_no,omitempty"`
	RefundAmount float64 `json:"refund_amount,omitempty" query:"refund_amount,omitempty"` //float64
	RefundReason string  `json:"refund_reason,omitempty" query:"refund_reason,omitempty"`
	OutRequestNo string  `json:"out_request_no,omitempty" query:"out_request_no,omitempty"`

	OperatorId string `json:"opreator_id,omitempty" query:"opreator_id,omitempty"`
	StoreId    string `json:"store_id,omitempty" query:"store_id,omitempty"`
	TerminalId string `json:"terminal_id,omitempty" query:"terminal_id,omitempty"`
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

	GoodsDetail        *[]GoodsDetail `json:"goods_detail,omitempty" query:"goods_detail,omitempty"`
	Body               string         `json:"body,omitempty" query:"body,omitempty"`
	OperatorId         string         `json:"operator_id,omitempty" query:"operator_id,omitempty"`
	StoreId            string         `json:"store_id,omitempty" query:"store_id,omitempty"`
	DisablePayChannels string         `json:"disable_pay_channels,omitempty" query:"disable_pay_channels,omitempty"`

	EnablePayChannels string        `json:"enable_pay_channels,omitempty" query:"enable_pay_channels,omitempty"`
	TerminalId        string        `json:"terminal_id,omitempty" query:"terminal_id,omitempty"`
	ExtendParams      *ExtendParams `json:"extend_params,omitempty" query:"extend_params,omitempty"`
	TimeoutExpress    string        `json:"timeout_express,omitempty" query:"timeout_express,omitempty"`
	BusinessParams    string        `json:"business_params,omitempty" query:"business_params,omitempty"`
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

	TradeNo       string `json:"trade_no,omitempty" query:"trade_no,omitempty"`
	OutTradeNo    string `json:"out_trade_no,omitempty" query:"out_trade_no,omitempty"`
	BuyerLogonId  string `json:"buyer_logon_id,omitempty" query:"buyer_logon_id,omitempty"`
	TotalAmount   string `json:"total_amount,omitempty" query:"total_amount,omitempty"` //float64
	ReceiptAmount string `json:"receipt_amount,omitempty" query:"receipt_amount,omitempty"`

	BuyerPayAmount string      `json:"buyer_pay_amount,omitempty" query:"buyer_pay_amount,omitempty"` //float64
	PointAmount    string      `json:"point_amount,omitempty" query:"point_amount,omitempty"`         //float64
	InvoiceAmount  string      `json:"invoice_amount,omitempty" query:"invoice_amount,omitempty"`     //float64
	GmtPayment     string      `json:"gmt_payment,omitempty" query:"gmt_payment,omitempty"`           //time.Time
	FundBillList   *[]FundBill `json:"fund_bill_list,omitempty" query:"fund_bill_list,omitempty"`

	CardBalance         string           `json:"card_balance,omitempty" query:"card_balance,omitempty"` //float64
	StoreName           string           `json:"store_name,omitempty" query:"store_name,omitempty"`
	BuyerUserId         string           `json:"buyer_user_id,omitempty" query:"buyer_user_id,omitempty"`
	DiscountGoodsDetail string           `json:"discount_goods_detail,omitempty" query:"discount_goods_detail,omitempty"`
	VoucherDetailList   *[]VoucherDetail `json:"voucher_detail_list,omitempty" query:"voucher_detail_list,omitempty"`

	BusinessParam string `json:"business_param,omitempty" query:"business_param,omitempty"`
	BuyerUserType string `json:"buyer_user_type,omitempty" query:"buyer_user_type,omitempty"`
}
type RespQueryDto struct {
	*RespBaseDto

	TradeNo      string `json:"trade_no,omitempty" query:"trade_no,omitempty"`
	OutTradeNo   string `json:"out_trade_no,omitempty" query:"out_trade_no,omitempty"`
	BuyerLogonId string `json:"buyer_logon_id,omitempty" query:"buyer_logon_id,omitempty"`
	TradeStatus  string `json:"trade_status,omitempty" query:"trade_status,omitempty"`
	TotalAmount  string `json:"total_amount,omitempty" query:"total_amount,omitempty"` //float64

	ReceiptAmount  string `json:"receipt_amount,omitempty" query:"receipt_amount,omitempty"`
	BuyerPayAmount string `json:"buyer_pay_amount,omitempty" query:"buyer_pay_amount,omitempty"` //float64
	PointAmount    string `json:"point_amount,omitempty" query:"point_amount,omitempty"`         //float64
	InvoiceAmount  string `json:"invoice_amount,omitempty" query:"invoice_amount,omitempty"`     //float64
	SendPayDate    string `json:"send_pay_date,omitempty" query:"send_pay_date,omitempty"`       //time.Time

	StoreId      string      `json:"store_id,omitempty" query:"store_id,omitempty"`
	TerminalId   string      `json:"terminal_id,omitempty" query:"terminal_id,omitempty"`
	FundBillList *[]FundBill `json:"fund_bill_list,omitempty" query:"fund_bill_list,omitempty"`
	StoreName    string      `json:"store_name,omitempty" query:"store_name,omitempty"`
	BuyerUserId  string      `json:"buyer_user_id,omitempty" query:"buyer_user_id,omitempty"`

	BuyerUserType string `json:"buyer_user_type,omitempty" query:"buyer_user_type,omitempty"`
}
type RespRefundDto struct {
	*RespBaseDto

	TradeNo      string `json:"trade_no,omitempty" query:"trade_no,omitempty"`
	OutTradeNo   string `json:"out_trade_no,omitempty" query:"out_trade_no,omitempty"`
	BuyerLogonId string `json:"buyer_logon_id,omitempty" query:"buyer_logon_id,omitempty"`
	FundChange   string `json:"fund_change,omitempty" query:"fund_change,omitempty"`
	RefundFee    string `json:"refund_fee,omitempty" query:"refund_fee,omitempty"` //float64

	GmtRefundPay         string              `json:"gmt_refund_pay,omitempty" query:"gmt_refund_pay,omitempty"` //time.Time
	RefundDetailItemList *[]RefundDetailItem `json:"refund_detail_item_list,omitempty" query:"refund_detail_item_list,omitempty"`
	StoreName            string              `json:"store_name,omitempty" query:"store_name,omitempty"`
	BuyerUserId          string              `json:"buyer_user_id,omitempty" query:"buyer_user_id,omitempty"`
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

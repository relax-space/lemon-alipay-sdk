package alipay

/*
tip:
response:all response member is  string,not float64
request:request is ok
*/

type ReqCustomerDto struct {
	PriKey string `json:"pri_key,omitempty"`
	PubKey string `json:"pub_key,omitempty"`
}

type ReqBaseDto struct {
	AppId    string `json:"app_id,omitempty"`
	Method   string `json:"method,omitempty"`
	Format   string `json:"format,omitempty"`
	Charset  string `json:"charset,omitempty"`
	SignType string `json:"sign_type,omitempty"`

	Sign         string `json:"sign,omitempty"`
	Timestamp    string `json:"timestamp,omitempty"`
	Version      string `json:"version,omitempty"`
	AppAuthToken string `json:"app_auth_token,omitempty"`
	BizContent   string `json:"biz_content"`
}

type ReqPayDto struct {
	*ReqBaseDto `json:"-"`

	NotifyUrl string `json:"notify_url,omitempty"`

	OutTradeNo  string `json:"out_trade_no,omitempty"`
	Scene       string `json:"scene,omitempty"`
	AuthCode    string `json:"auth_code,omitempty"`
	ProductCode string `json:"product_code,omitempty"`
	Subject     string `json:"subject,omitempty"`

	BuyerId            string  `json:"buyer_id,omitempty"`
	SellerId           string  `json:"seller_id,omitempty"`
	TotalAmount        float64 `json:"total_amount,omitempty"`        //float64
	DiscountableAmount float64 `json:"discountable_amount,omitempty"` //float64
	Body               string  `json:"body,omitempty"`

	GoodsDetails *[]GoodsDetail `json:"goods_detail,omitempty"`
	OperatorId   string         `json:"operator_id,omitempty"`
	StoreId      string         `json:"store_id,omitempty"`
	TerminalId   string         `json:"terminal_id,omitempty"`
	ExtendParams *ExtendParams  `json:"extend_params,omitempty"`

	TimeoutExpress string `json:"timeout_express,omitempty"`
}

type ReqQueryDto struct {
	*ReqBaseDto `json:"-"`

	OutTradeNo string `json:"out_trade_no,omitempty"`
	TradeNo    string `json:"trade_no,omitempty"`
}

type ReqRefundDto struct {
	*ReqBaseDto `json:"-"`

	OutTradeNo   string  `json:"out_trade_no,omitempty"`
	TradeNo      string  `json:"trade_no,omitempty"`
	RefundAmount float64 `json:"refund_amount,omitempty"` //float64
	RefundReason string  `json:"refund_reason,omitempty"`
	OutRequestNo string  `json:"out_request_no,omitempty"`

	OperatorId string `json:"opreator_id,omitempty"`
	StoreId    string `json:"store_id,omitempty"`
	TerminalId string `json:"terminal_id,omitempty"`
}

type ReqReverseDto struct {
	*ReqBaseDto `json:"-"`

	OutTradeNo string `json:"out_trade_no,omitempty"`
	TradeNo    string `json:"trade_no,omitempty"`
}

type ReqPrepayDto struct {
	*ReqBaseDto `json:"-"`

	NotifyUrl string `json:"notify_url,omitempty"`

	OutTradeNo         string  `json:"out_trade_no,omitempty"`
	SellerId           string  `json:"seller_id,omitempty"`
	TotalAmount        float64 `json:"total_amount,omitempty"`        //float64
	DiscountableAmount float64 `json:"discountable_amount,omitempty"` //float64
	Subject            string  `json:"subject,omitempty"`

	GoodsDetail        *[]GoodsDetail `json:"goods_detail,omitempty"`
	Body               string         `json:"body,omitempty"`
	OperatorId         string         `json:"operator_id,omitempty"`
	StoreId            string         `json:"store_id,omitempty"`
	DisablePayChannels string         `json:"disable_pay_channels,omitempty"`

	EnablePayChannels string        `json:"enable_pay_channels,omitempty"`
	TerminalId        string        `json:"terminal_id,omitempty"`
	ExtendParams      *ExtendParams `json:"extend_params,omitempty"`
	TimeoutExpress    string        `json:"timeout_express,omitempty"`
	BusinessParams    string        `json:"business_params,omitempty"`
}

type RespBaseDto struct {
	Code    string `json:"code,omitempty"`
	Msg     string `json:"msg,omitempty"`
	SubCode string `json:"sub_code,omitempty"`
	SubMsg  string `json:"sub_msg,omitempty"`
	Sign    string `json:"sign,omitempty"`
}

type RespPayDto struct {
	*RespBaseDto

	TradeNo       string `json:"trade_no,omitempty"`
	OutTradeNo    string `json:"out_trade_no,omitempty"`
	BuyerLogonId  string `json:"buyer_logon_id,omitempty"`
	TotalAmount   string `json:"total_amount,omitempty"` //float64
	ReceiptAmount string `json:"receipt_amount,omitempty"`

	BuyerPayAmount string      `json:"buyer_pay_amount,omitempty"` //float64
	PointAmount    string      `json:"point_amount,omitempty"`     //float64
	InvoiceAmount  string      `json:"invoice_amount,omitempty"`   //float64
	GmtPayment     string      `json:"gmt_payment,omitempty"`      //time.Time
	FundBillList   *[]FundBill `json:"fund_bill_list,omitempty"`

	CardBalance         string           `json:"card_balance,omitempty"` //float64
	StoreName           string           `json:"store_name,omitempty"`
	BuyerUserId         string           `json:"buyer_user_id,omitempty"`
	DiscountGoodsDetail string           `json:"discount_goods_detail,omitempty"`
	VoucherDetailList   *[]VoucherDetail `json:"voucher_detail_list,omitempty"`

	BusinessParam string `json:"business_param,omitempty"`
	BuyerUserType string `json:"buyer_user_type,omitempty"`
}
type RespQueryDto struct {
	*RespBaseDto

	TradeNo      string `json:"trade_no,omitempty"`
	OutTradeNo   string `json:"out_trade_no,omitempty"`
	BuyerLogonId string `json:"buyer_logon_id,omitempty"`
	TradeStatus  string `json:"trade_status,omitempty"`
	TotalAmount  string `json:"total_amount,omitempty"` //float64

	ReceiptAmount  string `json:"receipt_amount,omitempty"`
	BuyerPayAmount string `json:"buyer_pay_amount,omitempty"` //float64
	PointAmount    string `json:"point_amount,omitempty"`     //float64
	InvoiceAmount  string `json:"invoice_amount,omitempty"`   //float64
	SendPayDate    string `json:"send_pay_date,omitempty"`    //time.Time

	StoreId      string      `json:"store_id,omitempty"`
	TerminalId   string      `json:"terminal_id,omitempty"`
	FundBillList *[]FundBill `json:"fund_bill_list,omitempty"`
	StoreName    string      `json:"store_name,omitempty"`
	BuyerUserId  string      `json:"buyer_user_id,omitempty"`

	BuyerUserType string `json:"buyer_user_type,omitempty"`
}
type RespRefundDto struct {
	*RespBaseDto

	TradeNo      string  `json:"trade_no,omitempty"`
	OutTradeNo   string  `json:"out_trade_no,omitempty"`
	BuyerLogonId string  `json:"buyer_logon_id,omitempty"`
	FundChange   string  `json:"fund_change,omitempty"`
	RefundFee    float64 `json:"refund_fee,omitempty"`

	GmtRefundPay         string                `json:"gmt_refund_pay,omitempty"` //time.Time
	RefundDetailItemList *RefundDetailItemList `json:"refund_detail_item_list,omitempty"`
	StoreName            string                `json:"store_name,omitempty"`
	BuyerUserId          string                `json:"buyer_user_id,omitempty"`
}
type RespReverseDto struct {
	*RespBaseDto

	TradeNo    string `json:"trade_no,omitempty"`
	OutTradeNo string `json:"out_trade_no,omitempty"`
	RetryFlag  string `json:"retry_flag,omitempty"`
	Action     string `json:"action,omitempty"`
}

type RespPrepayDto struct {
	*RespBaseDto

	OutTradeNo string `json:"out_trade_no,omitempty"`
	QrCode     string `json:"qr_code,omitempty"`
}

type RefundDetailItemList struct {
	FundChannel string  `json:"fund_channel,omitempty"`
	Amount      float64 `json:"amount,omitempty"`
	RealAmount  float64 `json:"real_amount,omitempty"`
	FundType    string  `json:"fund_type,omitempty"`
}

type FundBill struct {
	FundChannel string `json:"fund_channel,omitempty"`
	Amount      string `json:"amount,omitempty"`      //float64
	RealAmount  string `json:"real_amount,omitempty"` //float64
}

type VoucherDetail struct {
	Id                         string `json:"id,omitempty"`
	Name                       string `json:"name,omitempty"`
	Type                       string `json:"type,omitempty"`
	Amount                     string `json:"amount,omitempty"`              //float64
	MerchantContribute         string `json:"merchant_contribute,omitempty"` //float64
	Othercontribute            string `json:"other_contribute,omitempty"`    //float64
	Memo                       string `json:"memo,omitempty"`
	TemplateId                 string `json:"template_id,omitempty"`
	PurchaseBuyerContribute    string `json:"purchase_buyer_contribute,omitempty"` //float64
	PurchaseMerchantContribute string `json:"purchase_buyer_contribute,omitempty"` //float64
	PurchaseAntContribute      string `json:"purchase_ant_contribute,omitempty"`   //float64
}

//req
type GoodsDetail struct {
	GoodsId       string  `json:"goods_id,omitempty"`
	GoodsName     string  `json:"goods_name,omitempty"`
	Quantity      int64   `json:"quantity,omitempty"`
	Price         float64 `json:"price,omitempty"`
	GoodsCategory string  `json:"goods_category,omitempty"`
	Body          string  `json:"body,omitempty"`
	ShowUrl       string  `json:"show_url,omitempty"`
}

type ExtendParams struct {
	SysServiceProviderId string `json:"sys_service_provider_id,omitempty"`
}

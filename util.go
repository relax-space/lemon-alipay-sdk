package alipay

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/relax-space/go-kit/sign"
)

const (
	OPENAPIURL = "https://openapi.alipay.com/gateway.do"
)

const (
	REQUEST_METHOD_PAY       = "alipay.trade.pay"
	REQUEST_METHOD_QUERY     = "alipay.trade.query"
	REQUEST_METHOD_REFUND    = "alipay.trade.refund"
	REQUEST_METHOD_CANCEL    = "alipay.trade.cancel"
	REQUEST_METHOD_PRECREATE = "alipay.trade.precreate"

	RESPONSE_METHOD_PAY       = "alipay_trade_pay_response"
	RESPONSE_METHOD_QUERY     = "alipay_trade_query_response"
	RESPONSE_METHOD_REFUND    = "alipay_trade_refund_response"
	RESPONSE_METHOD_CANCEL    = "alipay_trade_cancel_response"
	RESPONSE_METHOD_PRECREATE = "alipay_trade_precreate_response"
)

const (
	MESSAGE_PAYING   = "MESSAGE_PAYING"
	MESSAGE_ALIPAY   = "MESSAGE_ALIPAY"
	MESSAGE_OVERTIME = "MESSAGE_OVERTIME"
)

const (
	PRE_OUTTRADENO    = "11"
	PRE_OUTREFUNDNO   = "12"
	PRE_PREOUTTRADENO = "13"
)

func BuildCommonparam(appId, appAuthToken, method, notifyUrl string) (baseDto *ReqBaseDto) {
	baseDto = &ReqBaseDto{
		AppId:    appId,
		Method:   method,
		Format:   "JSON",
		Charset:  "utf-8",
		SignType: "RSA",

		Timestamp:    time.Now().Format("2006-01-02 15:04:05"),
		Version:      "1.0",
		AppAuthToken: appAuthToken,
		NotifyUrl:    notifyUrl,
	}
	return
}

func ValidResponse(respBaseDto *RespBaseDto, body []byte, signStr, respMethod, pubKey string) (code string, err error) {
	code, err = ValidSign(body, signStr, respMethod, pubKey)
	if err != nil {
		return
	}
	switch respBaseDto.Code {
	case "10000":
		code = SUC
		return
	case "10003":
		err = errors.New(MESSAGE_PAYING)
		code = E03
		return
	default:
		err = fmt.Errorf("\nvalidate response failure:code:%v,message:%v,subCode:%v,subMessage:%v",
			respBaseDto.Code, respBaseDto.Msg,
			respBaseDto.SubCode, respBaseDto.SubMsg,
		)
		code = E03
		//errors.New("validate response failure:")
		return
	}
	return
}

func ValidSign(body []byte, signStr, respMethod, pubKey string) (code string, err error) {
	enc := strings.TrimPrefix(string(body), `{"`+respMethod+`":{`)
	idx := strings.Index(enc, `},"sign":`)
	if idx == -1 { //when no appId,sign node is not exist
		err = errors.New(string(body))
		code = E03
		return
	}
	enc = "{" + enc[:idx] + "}"
	if isValid := sign.CheckSha1Sign(enc, signStr, pubKey); isValid != true {
		err = errors.New("Signature verification failed")
		code = E04
		return
	}
	return
}

func MovePayData(respQueryDto *RespQueryDto, respPayDto *RespPayDto) {
	respPayDto.TradeNo = respQueryDto.TradeNo
	respPayDto.OutTradeNo = respQueryDto.OutTradeNo
	respPayDto.BuyerLogonId = respQueryDto.BuyerLogonId
	respPayDto.PayAmount = respQueryDto.PayAmount
	respPayDto.SettleTransRate = respQueryDto.SettleTransRate

	respPayDto.TransPayRate = respQueryDto.TransPayRate
	respPayDto.TotalAmount = respQueryDto.TotalAmount
	respPayDto.TransCurrency = respQueryDto.TransCurrency
	respPayDto.SettleCurrency = respQueryDto.SettleCurrency
	respPayDto.SettleAmount = respQueryDto.SettleAmount

	respPayDto.PayCurrency = respQueryDto.PayCurrency
	respPayDto.ReceiptAmount = respQueryDto.ReceiptAmount
	respPayDto.BuyerPayAmount = respQueryDto.BuyerPayAmount
	respPayDto.PointAmount = respQueryDto.PointAmount
	respPayDto.InvoiceAmount = respQueryDto.InvoiceAmount

	respPayDto.GmtPayment = respQueryDto.SendPayDate
	respPayDto.FundBillList = respQueryDto.FundBillList
	//respPayDto.CardBalance = respQueryDto.CardBalance
	respPayDto.StoreName = respQueryDto.StoreName
	respPayDto.BuyerUserId = respQueryDto.BuyerUserId

	//respPayDto.DiscountGoodsDetail = respQueryDto.DiscountGoodsDetail
	//respPayDto.VoucherDetailList = respQueryDto.VoucherDetailList
	respPayDto.AuthTradePayMode = respQueryDto.AuthTradePayMode
	//respPayDto.BusinessParam = respQueryDto.BusinessParam
	respPayDto.BuyerUserType = respQueryDto.BuyerUserType

	respPayDto.MdiscountAmount = respQueryDto.MdiscountAmount
	respPayDto.DiscountAmount = respQueryDto.DiscountAmount

}

const (
	SUC = "SUC" //success
	E01 = "E01" //system error,can re-try
	E02 = "E02" //bad request format
	E03 = "E03" //message from alipay
	E04 = "E04" //bad response format
)

package alipay

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/relax-space/go-kit/base"

	"github.com/fatih/structs"
	"github.com/relax-space/go-kit/httpreq"
	"github.com/relax-space/go-kit/sign"

	"github.com/relax-space/go-kitt/random"
)

func LoopQuery(reqDto *ReqQueryDto, custDto *ReqCustomerDto, limit, interval int) (statusCode int, code string, result *RespPayDto, err error) {
	count := limit / interval
	waitTime := time.Duration(interval) * time.Second //2s
	for index := 0; index < count; index++ {
		var respQueryDto *RespQueryDto
		statusCode, code, respQueryDto, err = Query(reqDto, custDto)
		if err == nil { // 1. request wechat query api success
			if len(respQueryDto.TradeStatus) == 0 { //1.1 wechat query api response result is exception
				time.Sleep(waitTime)
				continue
			}
			switch respQueryDto.TradeStatus {
			case "TRADE_SUCCESS": //1.2 pay success
				result = &RespPayDto{}
				MovePayData(respQueryDto, result)
				code = SUC
				return
			case "TRADE_CLOSED", "TRADE_FINISHED":
				err = errors.New("alipay pay failure")
				code = E03
				return //1.3 pay failure
			case "WAIT_BUYER_PAY": //1.4 pay unknown
				if index < count {
					time.Sleep(waitTime)
					continue
				}
			}
		} else { //2. request wechat query api failure
			time.Sleep(waitTime)
			continue
		}
	}
	err = errors.New(MESSAGE_OVERTIME)
	code = E03
	return
}

func Pay(reqDto *ReqPayDto, custDto *ReqCustomerDto) (statusCode int, code string, result *RespPayDto, err error) {

	reqMethod := REQUEST_METHOD_PAY
	respMethod := RESPONSE_METHOD_PAY
	reqDto.ReqBaseDto = BuildCommonparam(reqDto.AppId, reqDto.AppAuthToken, reqMethod, "")

	if len(reqDto.OutTradeNo) == 0 {
		reqDto.OutTradeNo = random.NewUuid(PRE_OUTTRADENO)
	}
	if len(reqDto.Scene) == 0 {
		reqDto.Scene = "bar_code"
	}

	baseDto := structs.New(reqDto.ReqBaseDto)
	baseDto.TagName = "json"
	baseMap := baseDto.Map()

	b, err := json.Marshal(reqDto)
	if err != nil {
		code = E02
		return
	}
	baseMap["biz_content"] = string(b)
	signStr := base.JoinMapObject(baseMap)
	baseMap["sign"], err = sign.MakeSha1Sign(signStr, custDto.PriKey)
	if err != nil {
		err = errors.New("Signature create failed")
		code = E02
		return
	}
	resp, body, err := httpreq.NewPost(OPENAPIURL, []byte(base.JoinMapObjectEncode(baseMap)),
		&httpreq.Header{ContentType: httpreq.MIMEApplicationFormUTF8}, nil)
	if err != nil {
		code = E01
		return
	}
	statusCode = resp.StatusCode
	var respDto struct {
		Content *RespPayDto `json:"alipay_trade_pay_response"`
		Sign    string      `json:"sign"`
	}
	err = json.Unmarshal(body, &respDto)
	if err != nil {
		code = E04
		return
	}
	if respDto.Content == nil {
		err = errors.New("alipay response data format is wrong.")
		code = E04
		return
	}
	code, err = ValidResponse(respDto.Content.RespBaseDto, body, respDto.Sign, respMethod, custDto.PubKey)
	if err != nil {
		if err.Error() == MESSAGE_PAYING { //customer will query pay result by calling this method:LoopQuery
			result = &RespPayDto{OutTradeNo: reqDto.OutTradeNo}
		}
		return
	}
	result = respDto.Content
	return
}

func Query(reqDto *ReqQueryDto, custDto *ReqCustomerDto) (statusCode int, code string, result *RespQueryDto, err error) {
	reqMethod := REQUEST_METHOD_QUERY
	respMethod := RESPONSE_METHOD_QUERY

	reqDto.ReqBaseDto = BuildCommonparam(reqDto.AppId, reqDto.AppAuthToken, reqMethod, "")
	baseDto := structs.New(reqDto.ReqBaseDto)
	baseDto.TagName = "json"
	baseMap := baseDto.Map()

	b, err := json.Marshal(reqDto)
	if err != nil {
		code = E02
		return
	}
	baseMap["biz_content"] = string(b)
	signStr := base.JoinMapObject(baseMap)
	baseMap["sign"], err = sign.MakeSha1Sign(signStr, custDto.PriKey)
	if err != nil {
		err = errors.New("Signature create failed")
		code = E02
		return
	}

	resp, body, err := httpreq.NewPost(OPENAPIURL, []byte(base.JoinMapObjectEncode(baseMap)),
		&httpreq.Header{ContentType: httpreq.MIMEApplicationFormUTF8}, nil)
	if err != nil {
		code = E01
		return
	}
	statusCode = resp.StatusCode
	var respDto struct {
		Content *RespQueryDto `json:"alipay_trade_query_response"`
		Sign    string        `json:"sign"`
	}
	err = json.Unmarshal(body, &respDto)
	if err != nil {
		code = E04
		return
	}
	if respDto.Content == nil {
		err = errors.New("alipay response data format is wrong.")
		code = E04
		return
	}
	code, err = ValidResponse(respDto.Content.RespBaseDto, body, respDto.Sign, respMethod, custDto.PubKey)
	if err != nil {
		return
	}
	result = respDto.Content
	return
}

func Refund(reqDto *ReqRefundDto, custDto *ReqCustomerDto) (statusCode int, code string, result *RespRefundDto, err error) {
	reqMethod := REQUEST_METHOD_REFUND
	respMethod := RESPONSE_METHOD_REFUND

	if len(reqDto.OutRequestNo) == 0 {
		reqDto.OutRequestNo = random.NewUuid(PRE_OUTREFUNDNO)
	}

	reqDto.ReqBaseDto = BuildCommonparam(reqDto.AppId, reqDto.AppAuthToken, reqMethod, "")
	baseDto := structs.New(reqDto.ReqBaseDto)
	baseDto.TagName = "json"
	baseMap := baseDto.Map()

	b, err := json.Marshal(reqDto)
	if err != nil {
		code = E02
		return
	}
	baseMap["biz_content"] = string(b)
	signStr := base.JoinMapObject(baseMap)
	baseMap["sign"], err = sign.MakeSha1Sign(signStr, custDto.PriKey)
	if err != nil {
		err = errors.New("Signature create failed")
		code = E02
		return
	}

	resp, body, err := httpreq.NewPost(OPENAPIURL, []byte(base.JoinMapObjectEncode(baseMap)),
		&httpreq.Header{ContentType: httpreq.MIMEApplicationFormUTF8}, nil)
	if err != nil {
		code = E01
		return
	}
	statusCode = resp.StatusCode
	var respDto struct {
		Content *RespRefundDto `json:"alipay_trade_refund_response"`
		Sign    string         `json:"sign"`
	}
	err = json.Unmarshal(body, &respDto)
	if err != nil {
		code = E04
		return
	}
	if respDto.Content == nil {
		err = errors.New("alipay response data format is wrong.")
		code = E04
		return
	}
	code, err = ValidResponse(respDto.Content.RespBaseDto, body, respDto.Sign, respMethod, custDto.PubKey)
	if err != nil {
		return
	}
	result = respDto.Content
	return
}

func Reverse(reqDto *ReqReverseDto, custDto *ReqCustomerDto, count int, interval int) (statusCode int, code string, result *RespReverseDto, err error) {
	if count <= 0 {
		err = errors.New("reverse failure,please re-try")
		code = E01
		return
	}

	reqMethod := REQUEST_METHOD_CANCEL
	respMethod := RESPONSE_METHOD_CANCEL

	reqDto.ReqBaseDto = BuildCommonparam(reqDto.AppId, reqDto.AppAuthToken, reqMethod, "")
	baseDto := structs.New(reqDto.ReqBaseDto)
	baseDto.TagName = "json"
	baseMap := baseDto.Map()

	b, err := json.Marshal(reqDto)
	if err != nil {
		code = E02
		return
	}
	baseMap["biz_content"] = string(b)
	signStr := base.JoinMapObject(baseMap)
	baseMap["sign"], err = sign.MakeSha1Sign(signStr, custDto.PriKey)
	if err != nil {
		err = errors.New("Signature create failed")
		code = E02
		return
	}

	resp, body, err := httpreq.NewPost(OPENAPIURL, []byte(base.JoinMapObjectEncode(baseMap)),
		&httpreq.Header{ContentType: httpreq.MIMEApplicationFormUTF8}, nil)
	if err != nil {
		code = E01
		return
	}
	statusCode = resp.StatusCode
	var respDto struct {
		Content *RespReverseDto `json:"alipay_trade_cancel_response"`
		Sign    string          `json:"sign"`
	}
	err = json.Unmarshal(body, &respDto)
	if err != nil {
		code = E04
		return
	}
	if respDto.Content == nil {
		err = errors.New("alipay response data format is wrong.")
		code = E04
		return
	}
	code, err = ValidResponse(respDto.Content.RespBaseDto, body, respDto.Sign, respMethod, custDto.PubKey)
	if err != nil {
		return
	}
	if respDto.Content.RetryFlag == "N" {
		result = respDto.Content
		return
	} else {
		time.Sleep(time.Duration(interval) * time.Second) //10s
		count = count - 1
		return Reverse(reqDto, custDto, count, interval)
	}
}

func Prepay(reqDto *ReqPrepayDto, custDto *ReqCustomerDto) (statusCode int, code string, result *RespPrepayDto, err error) {
	reqMethod := REQUEST_METHOD_PRECREATE
	respMethod := RESPONSE_METHOD_PRECREATE

	if len(reqDto.OutTradeNo) == 0 {
		reqDto.OutTradeNo = random.NewUuid(PRE_PREOUTTRADENO)
	}
	reqDto.ReqBaseDto = BuildCommonparam(reqDto.AppId, reqDto.AppAuthToken, reqMethod, reqDto.NotifyUrl)
	baseDto := structs.New(reqDto.ReqBaseDto)
	baseDto.TagName = "json"
	baseMap := baseDto.Map()

	b, err := json.Marshal(reqDto)
	if err != nil {
		code = E02
		return
	}
	baseMap["biz_content"] = string(b)

	signStr := base.JoinMapObject(baseMap)
	baseMap["sign"], err = sign.MakeSha1Sign(signStr, custDto.PriKey)
	if err != nil {
		err = errors.New("Signature create failed")
		code = E02
		return
	}
	resp, body, err := httpreq.NewPost(OPENAPIURL, []byte(base.JoinMapObjectEncode(baseMap)),
		&httpreq.Header{ContentType: httpreq.MIMEApplicationFormUTF8}, nil)
	if err != nil {
		code = E01
		return
	}
	statusCode = resp.StatusCode
	var respDto struct {
		Content *RespPrepayDto `json:"alipay_trade_precreate_response"`
		Sign    string         `json:"sign"`
	}
	err = json.Unmarshal(body, &respDto)
	if err != nil {
		code = E04
		return
	}
	if respDto.Content == nil {
		err = errors.New("alipay response data format is wrong.")
		code = E04
		return
	}
	code, err = ValidResponse(respDto.Content.RespBaseDto, body, respDto.Sign, respMethod, custDto.PubKey)
	if err != nil {
		return
	}
	result = respDto.Content
	return
}

func Bill(reqDto *ReqBillDto, custDto *ReqCustomerDto) (statusCode int, code string, result *RespBillDto, err error) {
	reqMethod := REQUEST_METHOD_BILL
	respMethod := RESPONSE_METHOD_BILL

	reqDto.ReqBaseDto = BuildCommonparam(reqDto.AppId, reqDto.AppAuthToken, reqMethod, "")
	baseDto := structs.New(reqDto.ReqBaseDto)
	baseDto.TagName = "json"
	baseMap := baseDto.Map()

	b, err := json.Marshal(reqDto)
	if err != nil {
		code = E02
		return
	}
	baseMap["biz_content"] = string(b)
	signStr := base.JoinMapObject(baseMap)
	baseMap["sign"], err = sign.MakeSha1Sign(signStr, custDto.PriKey)
	if err != nil {
		err = errors.New("Signature create failed")
		code = E02
		return
	}

	resp, body, err := httpreq.NewPost(OPENAPIURL, []byte(base.JoinMapObjectEncode(baseMap)),
		&httpreq.Header{ContentType: httpreq.MIMEApplicationFormUTF8}, nil)
	if err != nil {
		code = E01
		return
	}
	statusCode = resp.StatusCode
	var respDto struct {
		Content *RespBillDto `json:"alipay_data_dataservice_bill_downloadurl_query_response"`
		Sign    string       `json:"sign"`
	}
	err = json.Unmarshal(body, &respDto)
	if err != nil {
		code = E04
		return
	}
	if respDto.Content == nil {
		err = errors.New("alipay response data format is wrong.")
		code = E04
		return
	}
	code, err = ValidResponse(respDto.Content.RespBaseDto, body, respDto.Sign, respMethod, custDto.PubKey)
	if err != nil {
		return
	}
	result = respDto.Content
	return
}

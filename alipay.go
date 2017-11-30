package alipay

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/relax-space/go-kit/base"

	"github.com/fatih/structs"
	"github.com/relax-space/go-kit/httpreq"
	"github.com/relax-space/go-kit/sign"

	"github.com/relax-space/go-kitt/random"
)

func LoopQuery(reqDto *ReqQueryDto, custDto *ReqCustomerDto, limit, interval int) (result *RespPayDto, err error) {
	count := limit / interval
	waitTime := time.Duration(interval) * time.Second //2s
	for index := 0; index < count; index++ {
		var respQueryDto *RespQueryDto
		respQueryDto, err = Query(reqDto, custDto)
		if err == nil { // 1. request wechat query api success
			if len(respQueryDto.TradeStatus) == 0 { //1.1 wechat query api response result is exception
				time.Sleep(waitTime)
				continue
			}
			switch respQueryDto.TradeStatus {
			case "TRADE_SUCCESS": //1.2 pay success
				MovePayData(respQueryDto, result)
				return
			case "TRADE_CLOSED", "TRADE_FINISHED":
				err = errors.New("alipay pay failure")
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
	return
}

func Pay(reqDto *ReqPayDto, custDto *ReqCustomerDto) (result *RespPayDto, err error) {

	reqMethod := REQUEST_METHOD_PAY
	respMethod := RESPONSE_METHOD_PAY
	var respDto struct {
		Content *RespPayDto `json:"alipay_trade_pay_response"`
		Sign    string      `json:"sign"`
	}

	reqDto.ReqBaseDto = BuildCommonparam(reqDto.AppId, reqDto.AppAuthToken, reqMethod)

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
	baseMap["biz_content"] = string(b)
	signStr := base.JoinMapObject(baseMap)
	baseMap["sign"], err = sign.MakeSha1Sign(signStr, custDto.PriKey)
	if err != nil {
		err = errors.New("Signature create failed")
		return
	}

	_, body, err := httpreq.NewPost(OPENAPIURL, []byte(base.JoinMapObjectEncode(baseMap)),
		&httpreq.Header{ContentType: httpreq.MIMEApplicationFormUTF8}, nil)
	if err != nil {
		err = fmt.Errorf("%v:%v", MESSAGE_ALIPAY, err)
		return
	}

	err = json.Unmarshal(body, &respDto)
	if err != nil {
		return
	}
	if respDto.Content == nil {
		err = errors.New("validate response failure.")
		return
	}
	err = ValidResponse(respDto.Content.RespBaseDto, body, respDto.Sign, respMethod, custDto.PubKey)
	if err != nil {
		if err.Error() == MESSAGE_PAYING {
			result = &RespPayDto{OutTradeNo: reqDto.OutTradeNo}
			result.OutTradeNo = reqDto.OutTradeNo
		}
		return
	}
	result = respDto.Content
	return
}

func Query(reqDto *ReqQueryDto, custDto *ReqCustomerDto) (result *RespQueryDto, err error) {
	reqMethod := REQUEST_METHOD_QUERY
	respMethod := RESPONSE_METHOD_QUERY
	var respDto struct {
		Content *RespQueryDto `json:"alipay_trade_query_response"`
		Sign    string        `json:"sign"`
	}

	reqDto.ReqBaseDto = BuildCommonparam(reqDto.AppId, reqDto.AppAuthToken, reqMethod)
	baseDto := structs.New(reqDto.ReqBaseDto)
	baseDto.TagName = "json"
	baseMap := baseDto.Map()

	b, err := json.Marshal(reqDto)
	baseMap["biz_content"] = string(b)
	signStr := base.JoinMapObject(baseMap)
	baseMap["sign"], err = sign.MakeSha1Sign(signStr, custDto.PriKey)
	if err != nil {
		err = errors.New("Signature create failed")
		return
	}

	_, body, err := httpreq.NewPost(OPENAPIURL, []byte(base.JoinMapObjectEncode(baseMap)),
		&httpreq.Header{ContentType: httpreq.MIMEApplicationFormUTF8}, nil)
	if err != nil {
		err = fmt.Errorf("%v:%v", MESSAGE_ALIPAY, err)
		return
	}

	err = json.Unmarshal(body, &respDto)
	if err != nil {
		return
	}
	if respDto.Content == nil {
		err = errors.New("validate response failure.")
		return
	}
	err = ValidResponse(respDto.Content.RespBaseDto, body, respDto.Sign, respMethod, custDto.PubKey)
	if err != nil {
		return
	}
	result = respDto.Content
	return
}

func Refund(reqDto *ReqRefundDto, custDto *ReqCustomerDto) (result *RespRefundDto, err error) {
	reqMethod := REQUEST_METHOD_REFUND
	respMethod := RESPONSE_METHOD_REFUND
	var respDto struct {
		Content *RespRefundDto `json:"alipay_trade_refund_response"`
		Sign    string         `json:"sign"`
	}

	if len(reqDto.OutTradeNo) == 0 {
		reqDto.OutRequestNo = random.NewUuid(PRE_OUTREFUNDNO)
	}

	reqDto.ReqBaseDto = BuildCommonparam(reqDto.AppId, reqDto.AppAuthToken, reqMethod)
	baseDto := structs.New(reqDto.ReqBaseDto)
	baseDto.TagName = "json"
	baseMap := baseDto.Map()

	b, err := json.Marshal(reqDto)
	baseMap["biz_content"] = string(b)

	signStr := base.JoinMapObject(baseMap)
	baseMap["sign"], err = sign.MakeSha1Sign(signStr, custDto.PriKey)
	if err != nil {
		err = errors.New("Signature create failed")
		return
	}

	_, body, err := httpreq.NewPost(OPENAPIURL, []byte(base.JoinMapObjectEncode(baseMap)),
		&httpreq.Header{ContentType: httpreq.MIMEApplicationFormUTF8}, nil)
	if err != nil {
		err = fmt.Errorf("%v:%v", MESSAGE_ALIPAY, err)
		return
	}

	err = json.Unmarshal(body, &respDto)
	if err != nil {
		return
	}
	if respDto.Content == nil {
		err = errors.New("validate response failure.")
		return
	}
	err = ValidResponse(respDto.Content.RespBaseDto, body, respDto.Sign, respMethod, custDto.PubKey)
	if err != nil {
		return
	}
	result = respDto.Content
	return
}

func Reverse(reqDto *ReqReverseDto, custDto *ReqCustomerDto, count int, interval int) (result *RespReverseDto, err error) {
	if count <= 0 {
		err = errors.New("The count of reverse must be greater than 0")
		return
	}

	reqMethod := REQUEST_METHOD_CANCEL
	respMethod := RESPONSE_METHOD_CANCEL
	var respDto struct {
		Content *RespReverseDto `json:"alipay_trade_cancel_response"`
		Sign    string          `json:"sign"`
	}

	reqDto.ReqBaseDto = BuildCommonparam(reqDto.AppId, reqDto.AppAuthToken, reqMethod)
	baseDto := structs.New(reqDto.ReqBaseDto)
	baseDto.TagName = "json"
	baseMap := baseDto.Map()

	b, err := json.Marshal(reqDto)
	baseMap["biz_content"] = string(b)

	signStr := base.JoinMapObject(baseMap)
	baseMap["sign"], err = sign.MakeSha1Sign(signStr, custDto.PriKey)
	if err != nil {
		err = errors.New("Signature create failed")
		return
	}

	_, body, err := httpreq.NewPost(OPENAPIURL, []byte(base.JoinMapObjectEncode(baseMap)),
		&httpreq.Header{ContentType: httpreq.MIMEApplicationFormUTF8}, nil)
	if err != nil {
		err = fmt.Errorf("%v:%v", MESSAGE_ALIPAY, err)
		return
	}

	err = json.Unmarshal(body, &respDto)
	if err != nil {
		return
	}
	if respDto.Content == nil {
		err = errors.New("validate response failure.")
		return
	}
	err = ValidResponse(respDto.Content.RespBaseDto, body, respDto.Sign, respMethod, custDto.PubKey)
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

	return
}

func Prepay(reqDto *ReqPrepayDto, custDto *ReqCustomerDto) (result *RespPrepayDto, err error) {
	reqMethod := REQUEST_METHOD_PRECREATE
	respMethod := RESPONSE_METHOD_PRECREATE
	var respDto struct {
		Content *RespPrepayDto `json:"alipay_trade_precreate_response"`
		Sign    string         `json:"sign"`
	}

	if len(reqDto.OutTradeNo) == 0 {
		reqDto.OutTradeNo = random.NewUuid(PRE_PREOUTTRADENO)
	}
	reqDto.ReqBaseDto = BuildCommonparam(reqDto.AppId, reqDto.AppAuthToken, reqMethod)
	baseDto := structs.New(reqDto.ReqBaseDto)
	baseDto.TagName = "json"
	baseMap := baseDto.Map()

	b, err := json.Marshal(reqDto)
	baseMap["biz_content"] = string(b)

	signStr := base.JoinMapObject(baseMap)
	baseMap["sign"], err = sign.MakeSha1Sign(signStr, custDto.PriKey)
	if err != nil {
		err = errors.New("Signature create failed")
		return
	}

	_, body, err := httpreq.NewPost(OPENAPIURL, []byte(base.JoinMapObjectEncode(baseMap)),
		&httpreq.Header{ContentType: httpreq.MIMEApplicationFormUTF8}, nil)
	if err != nil {
		err = fmt.Errorf("%v:%v", MESSAGE_ALIPAY, err)
		return
	}

	err = json.Unmarshal(body, &respDto)
	if err != nil {
		return
	}
	if respDto.Content == nil {
		err = errors.New("validate response failure.")
		return
	}
	err = ValidResponse(respDto.Content.RespBaseDto, body, respDto.Sign, respMethod, custDto.PubKey)
	if err != nil {
		return
	}
	result = respDto.Content
	return
}

func CheckNotifySign(reqDto *ReqNotifyDto, custDto *ReqCustomerDto) (err error) {

	baseStruts := structs.New(reqDto)
	baseStruts.TagName = "json"
	baseMap := baseStruts.Map()

	rawSignb, err := base64.StdEncoding.DecodeString(reqDto.Sign)
	if err != nil {
		err = errors.New("Signature verification failed")
		return
	}
	delete(baseMap, "sign")
	delete(baseMap, "sign_type")
	signStr := base.JoinMapObjectEncode(baseMap)

	if !sign.CheckSha1Sign(signStr, string(rawSignb), custDto.PubKey) {
		err = errors.New("Signature verification failed")
		return
	}

	return
}

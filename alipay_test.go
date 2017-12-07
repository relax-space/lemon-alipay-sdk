package alipay

import (
	"flag"
	"fmt"
	"kit/test"
	"os"
	"testing"
)

var (
	appId  = flag.String("ALIPAY_APPID", os.Getenv("ALIPAY_APPID"), "ALIPAY_APPID")
	priKey = flag.String("ALIPAY_PRIKEY", os.Getenv("ALIPAY_PRIKEY"), "ALIPAY_PRIKEY")
	pubKey = flag.String("ALIPAY_PUBKEY", os.Getenv("ALIPAY_PUBKEY"), "ALIPAY_PUBKEY")
)

func Test_Pay(t *testing.T) {
	reqDto := ReqPayDto{
		ReqBaseDto: &ReqBaseDto{
			AppId: *appId,
		},
		AuthCode:    "283663915849522991",
		Subject:     "xinmiao test ali",
		TotalAmount: 0.01,
	}
	customDto := ReqCustomerDto{
		PriKey: *priKey,
		PubKey: *pubKey,
	}
	result, err := Pay(&reqDto, &customDto)
	if err != nil {
		if err.Error() == MESSAGE_PAYING {
			dto := ReqQueryDto{
				ReqBaseDto: reqDto.ReqBaseDto,
				OutTradeNo: result.OutTradeNo,
			}
			respPayDto, err := LoopQuery(&dto, &customDto, 40, 2)
			fmt.Printf("%+v", respPayDto)
			test.Ok(t, err)
			return
		}
		test.Ok(t, err)
	}
	fmt.Printf("%+v", result)
	test.Ok(t, err)

}

func Test_Query(t *testing.T) {
	reqDto := ReqQueryDto{
		ReqBaseDto: &ReqBaseDto{
			AppId: *appId,
		},
		OutTradeNo: "11593651266244657670",
	}
	custDto := ReqCustomerDto{
		PriKey: *priKey,
		PubKey: *pubKey,
	}
	result, err := Query(&reqDto, &custDto)
	fmt.Printf("%+v", result)
	test.Ok(t, err)
}

func Test_Refund(t *testing.T) {

	reqDto := ReqRefundDto{
		ReqBaseDto: &ReqBaseDto{
			AppId: *appId,
		},
		OutTradeNo:   "112911611299552014581743175398",
		RefundAmount: 0.01,
	}
	custDto := ReqCustomerDto{
		PriKey: *priKey,
		PubKey: *pubKey,
	}
	result, err := Refund(&reqDto, &custDto)
	fmt.Printf("%+v", result)
	test.Ok(t, err)
}

func Test_Reverse(t *testing.T) {

	reqDto := ReqReverseDto{
		ReqBaseDto: &ReqBaseDto{
			AppId: *appId,
		},
		OutTradeNo: "11593651266244657670",
	}
	custDto := ReqCustomerDto{
		PriKey: *priKey,
		PubKey: *pubKey,
	}
	result, err := Reverse(&reqDto, &custDto, 10, 10)
	fmt.Printf("%+v", result)
	test.Ok(t, err)
}

func Test_Prepay(t *testing.T) {
	reqDto := ReqPrepayDto{
		ReqBaseDto: &ReqBaseDto{
			AppId: *appId,
		},
		Subject:     "xinmiao test ali",
		TotalAmount: 0.01,
		NotifyUrl:   "https://staging.p2shop.cn/ipay/v3/al/notify",
	}
	custDto := ReqCustomerDto{
		PriKey: *priKey,
		PubKey: *pubKey,
	}
	result, err := Prepay(&reqDto, &custDto)
	fmt.Printf("%+v", result)
	test.Ok(t, err)
}

package alipay

import (
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/relax-space/go-kit/test"
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
		AuthCode:    "284724726900638072",
		Subject:     "xinmiao test ali",
		TotalAmount: 0.01,
	}
	customDto := ReqCustomerDto{
		PriKey: *priKey,
		PubKey: *pubKey,
	}
	statusCode, code, result, err := Pay(&reqDto, &customDto)
	if err != nil {
		if err.Error() == MESSAGE_PAYING {
			dto := ReqQueryDto{
				ReqBaseDto: reqDto.ReqBaseDto,
				OutTradeNo: result.OutTradeNo,
			}
			statusCode, code, respPayDto, err := LoopQuery(&dto, &customDto, 2, 2)
			if err != nil && err.Error() == MESSAGE_OVERTIME {
				reqDto := ReqReverseDto{
					ReqBaseDto: &ReqBaseDto{
						AppId: *appId,
					},
					OutTradeNo: result.OutTradeNo,
				}
				custDto := ReqCustomerDto{
					PriKey: *priKey,
					PubKey: *pubKey,
				}
				statusCode, code, result, err := Reverse(&reqDto, &custDto, 10, 10)
				fmt.Printf("Reverse code:%+v\n", code)
				fmt.Printf("Reverse status code:%+v\n", statusCode)
				fmt.Printf("Reverse result:%+v\n", result)
				test.Ok(t, err)
				return
			}
			fmt.Printf("LoopQuery code:%+v\n", code)
			fmt.Printf("LoopQuery status code:%+v\n", statusCode)
			fmt.Printf("LoopQuery respPayDto:%+v\n", respPayDto)
			test.Ok(t, err)
			return
		}
		test.Ok(t, err)
	}
	fmt.Printf("Pay code:%+v\n", code)
	fmt.Printf("Pay status code:%+v\n", statusCode)
	fmt.Printf("Pay result:%+v\n", result)
	test.Ok(t, err)

}

func Test_Query(t *testing.T) {
	reqDto := ReqQueryDto{
		ReqBaseDto: &ReqBaseDto{
			AppId: *appId,
		},
		OutTradeNo: "111810178362336733927457547",
	}
	custDto := ReqCustomerDto{
		PriKey: *priKey,
		PubKey: *pubKey,
	}
	statusCode, code, result, err := Query(&reqDto, &custDto)
	fmt.Printf("code:%+v", code)
	fmt.Printf("status code:%+v", statusCode)
	fmt.Printf("result:%+v", result)
	test.Ok(t, err)
}

func Test_Refund(t *testing.T) {

	reqDto := ReqRefundDto{
		ReqBaseDto: &ReqBaseDto{
			AppId: *appId,
		},
		OutTradeNo:   "111810179304049813706983983",
		RefundAmount: 0.01,
	}
	custDto := ReqCustomerDto{
		PriKey: *priKey,
		PubKey: *pubKey,
	}
	statusCode, code, result, err := Refund(&reqDto, &custDto)
	fmt.Printf("code:%+v", code)
	fmt.Printf("status code:%+v", statusCode)
	fmt.Printf("respPayDto:%+v", result)
	test.Ok(t, err)
}

func Test_Reverse(t *testing.T) {

	reqDto := ReqReverseDto{
		ReqBaseDto: &ReqBaseDto{
			AppId: *appId,
		},
		OutTradeNo: "111810178362336733927457547",
	}
	custDto := ReqCustomerDto{
		PriKey: *priKey,
		PubKey: *pubKey,
	}
	statusCode, code, result, err := Reverse(&reqDto, &custDto, 10, 10)
	fmt.Printf("code:%+v", code)
	fmt.Printf("status code:%+v", statusCode)
	fmt.Printf("respPayDto:%+v", result)
	test.Ok(t, err)
}

func Test_Prepay(t *testing.T) {
	reqDto := ReqPrepayDto{
		ReqBaseDto: &ReqBaseDto{
			AppId:     *appId,
			NotifyUrl: "https://staging.p2shop.cn/ipay/v3/al/notify",
		},
		Subject:     "xinmiao test ali",
		TotalAmount: 0.01,
	}
	custDto := ReqCustomerDto{
		PriKey: *priKey,
		PubKey: *pubKey,
	}
	statusCode, code, result, err := Prepay(&reqDto, &custDto)
	fmt.Printf("code:%+v", code)
	fmt.Printf("status code:%+v", statusCode)
	fmt.Printf("respPayDto:%+v", result)
	test.Ok(t, err)
}

func Test_Bill(t *testing.T) {
	reqDto := ReqBillDto{
		ReqBaseDto: &ReqBaseDto{
			AppId:     *appId,
			NotifyUrl: "https://staging.p2shop.cn/ipay/v3/al/notify",
		},
		BillType: "trade",
		BillDate: "2018-12-17", //2018-12-23
	}
	custDto := ReqCustomerDto{
		PriKey: *priKey,
		PubKey: *pubKey,
	}
	statusCode, code, result, err := Bill(&reqDto, &custDto)
	fmt.Printf("code:%+v", code)
	fmt.Printf("status code:%+v", statusCode)
	fmt.Printf("respPayDto:%+v", result)
	test.Ok(t, err)
}

# lemon-alipay-sdk
alipay pay go-sdk

## Installation
```
go get github.com/relax-space/lemon-alipay-sdk
```

## Usage
> pay
> query
> refund
> reverse

### pay
```go
    reqDto := ReqPayDto{
		ReqBaseDto: &ReqBaseDto{
			AppId: "******",
		},
		AuthCode:    "283663915849522991",
		Subject:     "xinmiao test ali",
		TotalAmount: 0.01,
	}
	customDto := ReqCustomerDto{
		PriKey:  "******",
		PubKey:  "******",
	}
	result, err := Pay(&reqDto, &customDto)
    fmt.Printf("%+v,%v", result,err)
```

### query
```go
	reqDto := ReqQueryDto{
		ReqBaseDto: &ReqBaseDto{
			AppId: "******",
		},
		OutTradeNo: "11593651266244657670",
	}
	custDto := ReqCustomerDto{
		PriKey: "******",
		PubKey: "******",
	}
	result, err := Query(&reqDto, &custDto)
    fmt.Printf("%+v,%v", result,err)
```

### refund
```go
	reqDto := ReqRefundDto{
		ReqBaseDto: &ReqBaseDto{
			AppId: "******",
		},
		OutTradeNo:   "112911611299552014581743175398",
		RefundAmount: 0.01,
	}
	custDto := ReqCustomerDto{
		PriKey: "******",
		PubKey: "******",
	}
	result, err := Refund(&reqDto, &custDto)
    fmt.Printf("%+v,%v", result,err)
```

### reverse
```go
	reqDto := ReqReverseDto{
		ReqBaseDto: &ReqBaseDto{
			AppId: "******",
		},
		OutTradeNo: "11593651266244657670",
	}
	custDto := ReqCustomerDto{
		PriKey: "******",
		PubKey: "******",
	}
	result, err := Reverse(&reqDto, &custDto, 10, 10)
    fmt.Printf("%+v,%v", result,err)
```


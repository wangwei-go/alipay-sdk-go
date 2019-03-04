package test

import (
	"testing"

	"github.com/wangwei-go/alipay-sdk-go/request"
)

func TestAlipayFundCouponOrderAppPay(t *testing.T) {
	client, _ := NewClient()
	data := &request.AlipayFundCouponOrderAppPayRequest{
		OutOrderNo:   "20180901000000AFCOAPP00001",
		OutRequestNo: "20180901000000AFCOAPP00001",
		OrderTitle:   "发送红包",
		Amount:       "9.99",
	}
	_, _ = client.SendRequest(request.AlipayFundCouponOrderAppPayMethod, data)
}

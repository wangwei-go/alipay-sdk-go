package test

import (
	"testing"

	"github.com/wangwei-go/alipay-sdk-go/request"
)

func TestAlipayTradeRefund(t *testing.T) {
	client, _ := NewClient()
	data := request.AlipayTradeRefundRequest{
		OutTradeNo:   "20180901000000ATWP00001",
		RefundAmount: "9.99",
		RefundReason: "退款测试",
	}
	_, _ = client.SendRequest(request.AlipayTradeRefundMethod, data)
}

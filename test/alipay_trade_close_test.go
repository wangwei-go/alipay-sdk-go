package test

import (
	"testing"

	"github.com/wangwei-go/alipay-sdk-go/request"
)

func TestAlipayTradeClose(t *testing.T) {
	client, _ := NewClient()
	data := request.AlipayTradeCloseRequest{
		OutTradeNo: "20180901000000ATWP00001",
	}
	_, _ = client.SendRequest(request.AlipayTradeCloseMethod, data)
}

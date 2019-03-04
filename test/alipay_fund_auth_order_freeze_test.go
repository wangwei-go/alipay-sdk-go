package test

import (
	"testing"

	"github.com/wangwei-go/alipay-sdk-go/request"
)

func TestAlipayFundAuthOrderFreeze(t *testing.T) {
	client, _ := NewClient()
	data := &request.AlipayFundAuthOrderFreezeRequest{
		AuthCode:     "281442445657444252",
		AuthCodeType: "bar_code",
		OutOrderNo:   "20180901000000AFAOF00001",
		OutRequestNo: "20180901000000AFAOF00001",
		OrderTitle:   "预授权冻结",
		Amount:       "9.99",
		PayeeLogonId: "dlnhdb4422@sandbox.com",
	}
	_, _ = client.SendRequest(request.AlipayFundAuthOrderFreezeMethod, data)
}

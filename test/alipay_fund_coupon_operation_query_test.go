package test

import (
	"testing"

	"github.com/wangwei-go/alipay-sdk-go/request"
)

func TestAlipayFundCouponOperationQuery(t *testing.T) {
	client, _ := NewClient()
	data := &request.AlipayFundCouponOperationQueryRequest{
		AuthNo:       "",
		OutOrderNo:   "20180901000000AFCOPP00001",
		OperationId:  "",
		OutRequestNo: "20180901000000AFCOPP00001",
	}
	_, _ = client.SendRequest(request.AlipayFundCouponOperationQueryMethod, data)
}

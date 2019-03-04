package test

import (
	"testing"

	"github.com/wangwei-go/alipay-sdk-go/request"
)

func TestAlipayDataDataserviceBillDownloadurlQuery(t *testing.T) {
	client, _ := NewClient()
	data := request.AlipayDataDataserviceBillDownloadurlQueryRequest{
		BillType: "trade",
		BillDate: "2018-07",
	}
	_, _ = client.SendRequest(request.AlipayDataDataserviceBillDownloadurlQueryMethod, data)
}

package test

import (
	"testing"

	"github.com/wangwei-go/alipay-sdk-go/request"
)

func TestAlipayOpenAuthTokenApp(t *testing.T) {
	client, _ := NewClient()
	data := &request.AlipayOpenAuthTokenAppRequest{
		GrantType: "authorization_code",
		Code:      "0c41e956be0949449339c1b1e8e40X96",
	}
	_, _ = client.SendRequest(request.AlipayOpenAuthTokenAppMethod, data)
}

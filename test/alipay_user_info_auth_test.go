package test

import (
	"testing"

	"github.com/wangwei-go/alipay-sdk-go/request"
)

func TestAlipayUserInfoAuth(t *testing.T) {
	client, config := NewClient()
	client.SetReturnUrl(config.ReturnUrl)
	data := &request.AlipayUserInfoAuthRequest{
		Scopes: []string{"auth_user"},
		State:  "12345678",
	}
	_, _ = client.RequestUrl(request.AlipayUserInfoAuthMethod, data)
}

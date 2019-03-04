package test

import (
	"testing"

	"github.com/wangwei-go/alipay-sdk-go/request"
)

func TestAlipayMarketingCampaignCashDetailQuery(t *testing.T) {
	client, _ := NewClient()
	data := &request.AlipayMarketingCampaignCashDetailQueryRequest{
		CrowdNo: "5PZx2Y5c55NlJV_FXl0V0_Wve9z3gpyqu-HzZaTrTFTMnSZ96O-zxUfKlHp5cxmx",
	}
	_, _ = client.SendRequest(request.AlipayMarketingCampaignCashDetailQueryMethod, data)
}

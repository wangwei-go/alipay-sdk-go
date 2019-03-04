package test

import (
	"testing"

	"github.com/wangwei-go/alipay-sdk-go/request"
)

func TestAlipayMarketingCampaignCashStatusModify(t *testing.T) {
	client, _ := NewClient()
	data := &request.AlipayMarketingCampaignCashStatusModifyRequest{
		CrowdNo:    "5PZx2Y5c55NlJV_FXl0V0_Wve9z3gpyqu-HzZaTrTFTMnSZ96O-zxUfKlHp5cxmx",
		CampStatus: "PAUSE",
	}
	_, _ = client.SendRequest(request.AlipayMarketingCampaignCashStatusModifyMethod, data)
}

package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// UpdatePrePaidBandwidthResponse Response Object
type UpdatePrePaidBandwidthResponse struct {
	Bandwidth *BandwidthResp `json:"bandwidth,omitempty"`

	// 订单号（包周期场景返回该字段）
	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdatePrePaidBandwidthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePrePaidBandwidthResponse struct{}"
	}

	return strings.Join([]string{"UpdatePrePaidBandwidthResponse", string(data)}, " ")
}

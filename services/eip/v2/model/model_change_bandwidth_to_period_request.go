package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ChangeBandwidthToPeriodRequest Request Object
type ChangeBandwidthToPeriodRequest struct {
	Body *BwChangeToPeriodReq `json:"body,omitempty"`
}

func (o ChangeBandwidthToPeriodRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeBandwidthToPeriodRequest struct{}"
	}

	return strings.Join([]string{"ChangeBandwidthToPeriodRequest", string(data)}, " ")
}

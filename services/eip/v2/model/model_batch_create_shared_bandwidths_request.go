package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// BatchCreateSharedBandwidthsRequest Request Object
type BatchCreateSharedBandwidthsRequest struct {
	Body *BatchCreateBandwidthRequestBody `json:"body,omitempty"`
}

func (o BatchCreateSharedBandwidthsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateSharedBandwidthsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateSharedBandwidthsRequest", string(data)}, " ")
}

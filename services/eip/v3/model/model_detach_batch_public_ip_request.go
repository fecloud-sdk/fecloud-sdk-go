package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// DetachBatchPublicIpRequest Request Object
type DetachBatchPublicIpRequest struct {
	Body *DetachBatchSharedbwReq `json:"body,omitempty"`
}

func (o DetachBatchPublicIpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetachBatchPublicIpRequest struct{}"
	}

	return strings.Join([]string{"DetachBatchPublicIpRequest", string(data)}, " ")
}

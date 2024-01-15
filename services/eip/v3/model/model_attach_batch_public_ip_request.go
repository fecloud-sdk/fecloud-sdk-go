package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// AttachBatchPublicIpRequest Request Object
type AttachBatchPublicIpRequest struct {
	Body *BatchAttachSharebwReq `json:"body,omitempty"`
}

func (o AttachBatchPublicIpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachBatchPublicIpRequest struct{}"
	}

	return strings.Join([]string{"AttachBatchPublicIpRequest", string(data)}, " ")
}

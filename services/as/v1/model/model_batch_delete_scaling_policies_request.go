package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// BatchDeleteScalingPoliciesRequest Request Object
type BatchDeleteScalingPoliciesRequest struct {
	Body *BatchDeleteScalingPoliciesOption `json:"body,omitempty"`
}

func (o BatchDeleteScalingPoliciesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteScalingPoliciesRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteScalingPoliciesRequest", string(data)}, " ")
}

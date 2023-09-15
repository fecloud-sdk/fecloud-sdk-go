package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// BatchResumeScalingPoliciesRequest Request Object
type BatchResumeScalingPoliciesRequest struct {
	Body *BatchResumeScalingPoliciesOption `json:"body,omitempty"`
}

func (o BatchResumeScalingPoliciesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchResumeScalingPoliciesRequest struct{}"
	}

	return strings.Join([]string{"BatchResumeScalingPoliciesRequest", string(data)}, " ")
}

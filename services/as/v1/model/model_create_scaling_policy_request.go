package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// CreateScalingPolicyRequest Request Object
type CreateScalingPolicyRequest struct {
	Body *CreateScalingPolicyOption `json:"body,omitempty"`
}

func (o CreateScalingPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScalingPolicyRequest struct{}"
	}

	return strings.Join([]string{"CreateScalingPolicyRequest", string(data)}, " ")
}

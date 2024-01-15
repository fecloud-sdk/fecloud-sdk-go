package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowScalingPolicyRequest struct {
	ScalingPolicyId string `json:"scaling_policy_id"`
}

func (o ShowScalingPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowScalingPolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowScalingPolicyRequest", string(data)}, " ")
}

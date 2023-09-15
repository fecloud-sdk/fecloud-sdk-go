package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ShowScalingPolicyRequest Request Object
type ShowScalingPolicyRequest struct {

	// 伸缩组ID。
	ScalingPolicyId string `json:"scaling_policy_id"`
}

func (o ShowScalingPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowScalingPolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowScalingPolicyRequest", string(data)}, " ")
}

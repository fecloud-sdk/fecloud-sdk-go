package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateScalingPolicyResponse struct {
	ScalingPolicyId *string `json:"scaling_policy_id,omitempty"`
	HttpStatusCode  int     `json:"-"`
}

func (o UpdateScalingPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateScalingPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateScalingPolicyResponse", string(data)}, " ")
}

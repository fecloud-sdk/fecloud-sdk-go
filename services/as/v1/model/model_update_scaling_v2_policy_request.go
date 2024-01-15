package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type UpdateScalingV2PolicyRequest struct {
	ScalingPolicyId string `json:"scaling_policy_id"`

	Body *UpdateScalingV2PolicyOption `json:"body,omitempty"`
}

func (o UpdateScalingV2PolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateScalingV2PolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdateScalingV2PolicyRequest", string(data)}, " ")
}

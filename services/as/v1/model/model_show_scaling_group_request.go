package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowScalingGroupRequest struct {
	ScalingGroupId string `json:"scaling_group_id"`
}

func (o ShowScalingGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowScalingGroupRequest struct{}"
	}

	return strings.Join([]string{"ShowScalingGroupRequest", string(data)}, " ")
}

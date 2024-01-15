package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type BatchProtectScalingInstancesRequest struct {
	ScalingGroupId string `json:"scaling_group_id"`

	Body *BatchProtectInstancesOption `json:"body,omitempty"`
}

func (o BatchProtectScalingInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchProtectScalingInstancesRequest struct{}"
	}

	return strings.Join([]string{"BatchProtectScalingInstancesRequest", string(data)}, " ")
}

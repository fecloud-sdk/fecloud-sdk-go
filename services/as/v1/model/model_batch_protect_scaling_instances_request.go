package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// BatchProtectScalingInstancesRequest Request Object
type BatchProtectScalingInstancesRequest struct {

	// 实例ID。
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

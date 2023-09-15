package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// PauseScalingGroupRequest Request Object
type PauseScalingGroupRequest struct {

	// 伸缩组ID
	ScalingGroupId string `json:"scaling_group_id"`

	Body *PauseScalingGroupOption `json:"body,omitempty"`
}

func (o PauseScalingGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PauseScalingGroupRequest struct{}"
	}

	return strings.Join([]string{"PauseScalingGroupRequest", string(data)}, " ")
}

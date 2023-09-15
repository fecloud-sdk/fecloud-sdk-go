package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// UpdateScalingGroupResponse Response Object
type UpdateScalingGroupResponse struct {

	// 伸缩组ID
	ScalingGroupId *string `json:"scaling_group_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateScalingGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateScalingGroupResponse struct{}"
	}

	return strings.Join([]string{"UpdateScalingGroupResponse", string(data)}, " ")
}

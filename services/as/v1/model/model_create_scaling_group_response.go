package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// CreateScalingGroupResponse Response Object
type CreateScalingGroupResponse struct {

	// 伸缩组ID
	ScalingGroupId *string `json:"scaling_group_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateScalingGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScalingGroupResponse struct{}"
	}

	return strings.Join([]string{"CreateScalingGroupResponse", string(data)}, " ")
}

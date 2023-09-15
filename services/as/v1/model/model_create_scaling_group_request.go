package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// CreateScalingGroupRequest Request Object
type CreateScalingGroupRequest struct {
	Body *CreateScalingGroupOption `json:"body,omitempty"`
}

func (o CreateScalingGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScalingGroupRequest struct{}"
	}

	return strings.Join([]string{"CreateScalingGroupRequest", string(data)}, " ")
}

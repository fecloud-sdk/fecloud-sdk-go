package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type NodePoolNodeAutoscaling struct {
	Enable *bool `json:"enable,omitempty"`

	MinNodeCount *int32 `json:"minNodeCount,omitempty"`

	MaxNodeCount *int32 `json:"maxNodeCount,omitempty"`

	ScaleDownCooldownTime *int32 `json:"scaleDownCooldownTime,omitempty"`

	Priority *int32 `json:"priority,omitempty"`
}

func (o NodePoolNodeAutoscaling) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodePoolNodeAutoscaling struct{}"
	}

	return strings.Join([]string{"NodePoolNodeAutoscaling", string(data)}, " ")
}

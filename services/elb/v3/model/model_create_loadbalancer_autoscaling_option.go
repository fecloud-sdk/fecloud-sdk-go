package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type CreateLoadbalancerAutoscalingOption struct {
	Enable bool `json:"enable"`

	MinL7FlavorId *string `json:"min_l7_flavor_id,omitempty"`
}

func (o CreateLoadbalancerAutoscalingOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLoadbalancerAutoscalingOption struct{}"
	}

	return strings.Join([]string{"CreateLoadbalancerAutoscalingOption", string(data)}, " ")
}

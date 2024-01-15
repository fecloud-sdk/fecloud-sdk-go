package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type AutoScalingPolicy struct {
	AutoScalingEnable bool `json:"auto_scaling_enable"`

	MinCapacity int32 `json:"min_capacity"`

	MaxCapacity int32 `json:"max_capacity"`

	ResourcesPlans *[]ResourcesPlan `json:"resources_plans,omitempty"`

	ExecScripts *[]ScaleScript `json:"exec_scripts,omitempty"`

	Rules *[]Rule `json:"rules,omitempty"`
}

func (o AutoScalingPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoScalingPolicy struct{}"
	}

	return strings.Join([]string{"AutoScalingPolicy", string(data)}, " ")
}

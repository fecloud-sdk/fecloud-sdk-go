package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ClusterScalingReq struct {
	ServiceId *string `json:"service_id,omitempty"`

	PlanId *string `json:"plan_id,omitempty"`

	Parameters *ClusterScalingParams `json:"parameters"`

	PreviousValues map[string]string `json:"previous_values,omitempty"`
}

func (o ClusterScalingReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterScalingReq struct{}"
	}

	return strings.Join([]string{"ClusterScalingReq", string(data)}, " ")
}

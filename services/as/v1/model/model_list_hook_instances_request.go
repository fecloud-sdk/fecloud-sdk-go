package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListHookInstancesRequest struct {
	ScalingGroupId string `json:"scaling_group_id"`

	InstanceId *string `json:"instance_id,omitempty"`
}

func (o ListHookInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHookInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListHookInstancesRequest", string(data)}, " ")
}

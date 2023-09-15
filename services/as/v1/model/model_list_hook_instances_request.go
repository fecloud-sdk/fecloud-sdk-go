package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// ListHookInstancesRequest Request Object
type ListHookInstancesRequest struct {

	// 伸缩组ID。
	ScalingGroupId string `json:"scaling_group_id"`

	// 伸缩实例ID。
	InstanceId *string `json:"instance_id,omitempty"`
}

func (o ListHookInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHookInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListHookInstancesRequest", string(data)}, " ")
}

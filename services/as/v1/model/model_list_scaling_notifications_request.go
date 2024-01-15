package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListScalingNotificationsRequest struct {
	ScalingGroupId string `json:"scaling_group_id"`
}

func (o ListScalingNotificationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScalingNotificationsRequest struct{}"
	}

	return strings.Join([]string{"ListScalingNotificationsRequest", string(data)}, " ")
}

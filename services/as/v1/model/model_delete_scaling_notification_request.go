package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type DeleteScalingNotificationRequest struct {
	ScalingGroupId string `json:"scaling_group_id"`

	TopicUrn string `json:"topic_urn"`
}

func (o DeleteScalingNotificationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteScalingNotificationRequest struct{}"
	}

	return strings.Join([]string{"DeleteScalingNotificationRequest", string(data)}, " ")
}

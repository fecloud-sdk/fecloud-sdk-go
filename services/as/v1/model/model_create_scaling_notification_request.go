package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// CreateScalingNotificationRequest Request Object
type CreateScalingNotificationRequest struct {

	// 伸缩组标识。
	ScalingGroupId string `json:"scaling_group_id"`

	Body *CreateNotificationOption `json:"body,omitempty"`
}

func (o CreateScalingNotificationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScalingNotificationRequest struct{}"
	}

	return strings.Join([]string{"CreateScalingNotificationRequest", string(data)}, " ")
}
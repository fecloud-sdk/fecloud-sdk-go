package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowBackgroundTaskRequest struct {
	InstanceId string `json:"instance_id"`

	TaskId string `json:"task_id"`
}

func (o ShowBackgroundTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBackgroundTaskRequest struct{}"
	}

	return strings.Join([]string{"ShowBackgroundTaskRequest", string(data)}, " ")
}

package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

// DeleteBackgroundTaskRequest Request Object
type DeleteBackgroundTaskRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 后台任务ID
	TaskId string `json:"task_id"`
}

func (o DeleteBackgroundTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBackgroundTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteBackgroundTaskRequest", string(data)}, " ")
}

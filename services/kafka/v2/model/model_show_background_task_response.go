package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ShowBackgroundTaskResponse struct {
	TaskCount *string `json:"task_count,omitempty"`

	Tasks          *[]ListBackgroundTasksRespTasks `json:"tasks,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o ShowBackgroundTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBackgroundTaskResponse struct{}"
	}

	return strings.Join([]string{"ShowBackgroundTaskResponse", string(data)}, " ")
}

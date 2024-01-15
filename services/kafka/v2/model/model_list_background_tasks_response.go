package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListBackgroundTasksResponse struct {
	TaskCount *string `json:"task_count,omitempty"`

	Tasks          *[]ListBackgroundTasksRespTasks `json:"tasks,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o ListBackgroundTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBackgroundTasksResponse struct{}"
	}

	return strings.Join([]string{"ListBackgroundTasksResponse", string(data)}, " ")
}

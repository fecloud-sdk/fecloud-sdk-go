package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListSinkTasksRespTasks struct {
	TaskId *string `json:"task_id,omitempty"`

	TaskName *string `json:"task_name,omitempty"`

	DestinationType *string `json:"destination_type,omitempty"`

	CreateTime *int64 `json:"create_time,omitempty"`

	Status *string `json:"status,omitempty"`

	Topics *string `json:"topics,omitempty"`
}

func (o ListSinkTasksRespTasks) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSinkTasksRespTasks struct{}"
	}

	return strings.Join([]string{"ListSinkTasksRespTasks", string(data)}, " ")
}

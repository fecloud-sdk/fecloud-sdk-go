package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListSinkTasksResponse struct {
	Tasks *[]ListSinkTasksRespTasks `json:"tasks,omitempty"`

	TotalNumber *int32 `json:"total_number,omitempty"`

	MaxTasks *int32 `json:"max_tasks,omitempty"`

	QuotaTasks     *int32 `json:"quota_tasks,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSinkTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSinkTasksResponse struct{}"
	}

	return strings.Join([]string{"ListSinkTasksResponse", string(data)}, " ")
}

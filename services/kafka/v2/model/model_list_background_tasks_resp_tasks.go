package model

import (
	"github.com/fecloud-sdk/fecloud-sdk-go/core/utils"

	"strings"
)

type ListBackgroundTasksRespTasks struct {
	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	UserName *string `json:"user_name,omitempty"`

	UserId *string `json:"user_id,omitempty"`

	Params *string `json:"params,omitempty"`

	Status *string `json:"status,omitempty"`

	CreatedAt *string `json:"created_at,omitempty"`

	UpdatedAt *string `json:"updated_at,omitempty"`
}

func (o ListBackgroundTasksRespTasks) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBackgroundTasksRespTasks struct{}"
	}

	return strings.Join([]string{"ListBackgroundTasksRespTasks", string(data)}, " ")
}
